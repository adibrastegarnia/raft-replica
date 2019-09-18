// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package roles

import (
	"context"
	"errors"
	"fmt"
	"github.com/atomix/atomix-go-node/pkg/atomix/node"
	raft "github.com/atomix/atomix-raft-node/pkg/atomix/raft/protocol"
	"github.com/atomix/atomix-raft-node/pkg/atomix/raft/state"
	"github.com/atomix/atomix-raft-node/pkg/atomix/raft/store"
	"github.com/atomix/atomix-raft-node/pkg/atomix/raft/store/log"
	"github.com/atomix/atomix-raft-node/pkg/atomix/raft/util"
	"io"
	"math"
	"time"
)

func newPassiveRole(raft raft.Raft, state state.Manager, store store.Store, log util.Logger) *PassiveRole {
	return &PassiveRole{
		raftRole: newRaftRole(raft, state, store, log),
	}
}

// PassiveRole implements a Raft follower
type PassiveRole struct {
	*raftRole
}

// Name is the name of the role
func (r *PassiveRole) Name() string {
	return "Passive"
}

// updateTermAndLeader updates the current term and leader if necessary
func (r *PassiveRole) updateTermAndLeader(term raft.Term, leader raft.MemberID) bool {
	// If the request indicates a term that is greater than the current term or no leader has been
	// set for the current term, update leader and term.
	if term > r.raft.Term() || (term == r.raft.Term() && r.raft.Leader() == "" && leader != "") {
		r.raft.SetTerm(term)
		r.raft.SetLeader(leader)
		return true
	}
	return false
}

// Append handles an append request
func (r *PassiveRole) Append(ctx context.Context, request *raft.AppendRequest) (*raft.AppendResponse, error) {
	r.log.Request("AppendRequest", request)
	r.raft.WriteLock()
	defer r.raft.WriteUnlock()
	r.updateTermAndLeader(request.Term, request.Leader)
	response, err := r.handleAppend(ctx, request)
	_ = r.log.Response("AppendResponse", response, err)
	return response, err
}

// handleAppend is a generic method for handling an AppendRequest
func (r *PassiveRole) handleAppend(ctx context.Context, request *raft.AppendRequest) (*raft.AppendResponse, error) {
	if response := r.checkTerm(request); response != nil {
		return response, nil
	}

	if response := r.checkPreviousEntry(request); response != nil {
		return response, nil
	}
	return r.appendEntries(request)
}

// checkTerm compares the given request to the current term
func (r *PassiveRole) checkTerm(request *raft.AppendRequest) *raft.AppendResponse {
	if request.Term < r.raft.Term() {
		r.log.Debug("Rejected %v: request term is less than the current term (%d)", request, r.raft.Term())
		return r.failAppend(r.store.Log().Writer().LastIndex())
	}
	return nil
}

// checkPreviousEntry compares the given request to the previous entry in the log
func (r *PassiveRole) checkPreviousEntry(request *raft.AppendRequest) *raft.AppendResponse {
	writer := r.store.Writer()
	reader := r.store.Reader()

	// If the previous term is set, validate that it matches the local log.
	// We check the previous log term since that indicates whether any entry is present in the leader's
	// log at the previous log index. It's possible that the leader can send a non-zero previous log index
	// with a zero term in the event the leader has compacted its logs and is sending the first entry.
	if request.PrevLogTerm != 0 {
		// Get the last entry written to the log.
		lastEntry := writer.LastEntry()

		// If the local log is non-empty...
		if lastEntry != nil {
			// If the previous log index is greater than the last entry index, fail the attempt.
			if request.PrevLogIndex > lastEntry.Index {
				r.log.Debug("Rejected %v: Previous index (%d) is greater than the local log's last index (%d)", request, request.PrevLogIndex, lastEntry.Index)
				return r.failAppend(lastEntry.Index)
			}

			// If the previous log index is less than the last written entry index, look up the entry.
			if request.PrevLogIndex < lastEntry.Index {
				// Reset the reader to the previous log index.
				if reader.NextIndex() != request.PrevLogIndex {
					reader.Reset(request.PrevLogIndex)
				}

				// The previous entry should exist in the log if we've gotten this far.
				previousEntry := reader.NextEntry()
				if previousEntry == nil {
					r.log.Debug("Rejected %v: Previous entry does not exist in the local log", request)
					return r.failAppend(lastEntry.Index)
				}

				// Read the previous entry and validate that the term matches the request previous log term.
				if request.PrevLogTerm != previousEntry.Entry.Term {
					r.log.Debug("Rejected %v: Previous entry term (%d) does not match local log's term for the same entry (%d)", request, request.PrevLogTerm, previousEntry.Entry.Term)
					return r.failAppend(request.PrevLogIndex - 1)
				}
				// If the previous log term doesn't equal the last entry term, fail the append, sending the prior entry.
			} else if request.PrevLogTerm != lastEntry.Entry.Term {
				r.log.Debug("Rejected %v: Previous entry term (%d) does not equal the local log's last term (%d)", request, request.PrevLogTerm, lastEntry.Entry.Term)
				return r.failAppend(request.PrevLogIndex - 1)
			}
		} else {
			// If the previous log index is set and the last entry is null, fail the append.
			if request.PrevLogIndex > 0 {
				r.log.Debug("Rejected %v: Previous index (%d) is greater than the local log's last index (0)", request, request.PrevLogIndex)
				return r.failAppend(0)
			}
		}
	}
	return nil
}

// appendEntries appends entries from the given request to the log
func (r *PassiveRole) appendEntries(request *raft.AppendRequest) (*raft.AppendResponse, error) {
	// Compute the last entry index from the previous log index and request entry count.
	lastEntryIndex := request.PrevLogIndex + raft.Index(len(request.Entries))

	// Ensure the commitIndex is not increased beyond the index of the last entry in the request.
	commitIndex := raft.Index(math.Max(float64(r.raft.CommitIndex()), math.Min(float64(request.CommitIndex), float64(lastEntryIndex))))

	// Track the last log index while entries are appended.
	index := request.PrevLogIndex

	if len(request.Entries) > 0 {
		writer := r.store.Writer()
		reader := r.store.Reader()

		// If the previous term is zero, that indicates the previous index represents the beginning of the log.
		// Reset the log to the previous index plus one.
		if request.PrevLogTerm == 0 {
			r.log.Debug("Reset first index to %d", request.PrevLogIndex+1)
			writer.Reset(request.PrevLogIndex + 1)
		}

		// Iterate through entries and append them.
		for _, entry := range request.Entries {
			index++

			// Get the last entry written to the log by the writer.
			lastEntry := writer.LastEntry()

			if lastEntry != nil {
				// If the last written entry index is greater than the next append entry index,
				// we need to validate that the entry that's already in the log matches this entry.
				if lastEntry.Index > index {
					// Reset the reader to the current entry index.
					if reader.NextIndex() != index {
						reader.Reset(index)
					}

					// If the reader does not have any next entry, that indicates an inconsistency between the reader and writer.
					existingEntry := reader.NextEntry()
					if existingEntry == nil {
						return nil, errors.New("log reader inconsistent with writer")
					}

					// If the existing entry term doesn't match the leader's term for the same entry, truncate
					// the log and append the leader's entry.
					if existingEntry.Entry.Term != entry.Term {
						writer.Truncate(index - 1)
						writer.Append(entry)
					}
					// If the last written entry is equal to the append entry index, we don't need
					// to read the entry from disk and can just compare the last entry in the writer.
				} else if lastEntry.Index == index {
					// If the last entry term doesn't match the leader's term for the same entry, truncate
					// the log and append the leader's entry.
					if lastEntry.Entry.Term != entry.Term {
						writer.Truncate(index - 1)
						indexed := writer.Append(entry)
						r.log.Trace("Appended %v", indexed)
					}
					// Otherwise, this entry is being appended at the end of the log.
				} else {
					// If the last entry index isn't the previous index, throw an exception because something crazy happened!
					if lastEntry.Index != index-1 {
						return nil, fmt.Errorf("log writer inconsistent with next append entry index %d", index)
					}

					// Append the entry and log a message.
					indexed := writer.Append(entry)
					r.log.Trace("Appended %v", indexed)
				}
				// Otherwise, if the last entry is null just append the entry and log a message.
			} else {
				indexed := writer.Append(entry)
				r.log.Trace("Appended %v", indexed)
			}
		}
	}

	// Update the context commit and global indices.
	prevCommitIndex := r.raft.SetCommitIndex(request.CommitIndex, commitIndex)
	if commitIndex > prevCommitIndex {
		r.log.Trace("Committed entries up to index %d", commitIndex)

		// If the commitIndex was for entries that were replicated in a prior request, ensure they're applied
		// to the state machine.
		if commitIndex <= request.PrevLogIndex {
			r.state.ApplyIndex(request.PrevLogIndex)
		}

		// Iterate through entries in the request and apply committed entries to the state machine.
		for i := 0; request.PrevLogIndex+raft.Index(i)+1 <= commitIndex && i < len(request.Entries); i++ {
			r.state.ApplyEntry(&log.LogEntry{
				Index: request.PrevLogIndex + raft.Index(i) + 1,
				Entry: request.Entries[i],
			}, nil)
		}
	}

	// Return a successful append response.
	return r.succeedAppend(index), nil
}

// failAppend returns a failed AppendResponse
func (r *PassiveRole) failAppend(lastIndex raft.Index) *raft.AppendResponse {
	return r.completeAppend(false, lastIndex)
}

// succeedAppend returns a successful AppendResponse
func (r *PassiveRole) succeedAppend(lastIndex raft.Index) *raft.AppendResponse {
	return r.completeAppend(true, lastIndex)
}

// completeAppend creates a new AppendResponse
func (r *PassiveRole) completeAppend(succeeded bool, lastIndex raft.Index) *raft.AppendResponse {
	return &raft.AppendResponse{
		Status:       raft.ResponseStatus_OK,
		Term:         r.raft.Term(),
		Succeeded:    succeeded,
		LastLogIndex: lastIndex,
	}
}

// Install handles an install request
func (r *PassiveRole) Install(stream raft.RaftService_InstallServer) error {
	var writer io.WriteCloser
	for {
		request, err := stream.Recv()
		r.log.Request("InstallRequest", request)

		// If the stream has ended, close the writer and respond successfully.
		if err == io.EOF {
			_ = writer.Close()
			response := &raft.InstallResponse{
				Status: raft.ResponseStatus_OK,
			}
			return r.log.Response("InstallResponse", response, stream.SendAndClose(response))
		}

		// If an error occurred, return the error.
		if err != nil {
			return r.log.Response("InstallResponse", nil, err)
		}

		// Acquire a write lock to write the snapshot.
		r.raft.WriteLock()

		// Update the term and leader
		r.updateTermAndLeader(request.Term, request.Leader)

		// If the request is for a lesser term, reject the request.
		if request.Term < r.raft.Term() {
			response := &raft.InstallResponse{
				Status: raft.ResponseStatus_ERROR,
				Error:  raft.RaftError_ILLEGAL_MEMBER_STATE,
			}
			return r.log.Response("InstallResponse", response, stream.SendAndClose(response))
		}

		if writer == nil {
			snapshot := r.store.Snapshot().NewSnapshot(request.Index, request.Timestamp)
			writer = snapshot.Writer()
		}

		_, err = writer.Write(request.Data)
		r.raft.WriteUnlock()
		if err != nil {
			response := &raft.InstallResponse{
				Status: raft.ResponseStatus_ERROR,
				Error:  raft.RaftError_PROTOCOL_ERROR,
			}
			return r.log.Response("InstallResponse", response, stream.SendAndClose(response))
		}
	}
}

// Command handles a command request
func (r *PassiveRole) Command(request *raft.CommandRequest, server raft.RaftService_CommandServer) error {
	r.log.Request("CommandRequest", request)
	r.raft.ReadLock()
	response := &raft.CommandResponse{
		Status: raft.ResponseStatus_ERROR,
		Error:  raft.RaftError_ILLEGAL_MEMBER_STATE,
		Leader: r.raft.Leader(),
		Term:   r.raft.Term(),
	}
	r.raft.ReadUnlock()
	return r.log.Response("CommandResponse", response, server.Send(response))
}

// Query handles a query request
func (r *PassiveRole) Query(request *raft.QueryRequest, server raft.RaftService_QueryServer) error {
	r.log.Request("QueryRequest", request)
	r.raft.ReadLock()
	leader := r.raft.Leader()

	// If this server has not yet applied entries up to the client's session ID, forward the
	// query to the leader. This ensures that a follower does not tell the client its session
	// doesn't exist if the follower hasn't had a chance to see the session's registration entry.
	if r.raft.Status() != raft.StatusReady {
		r.raft.ReadUnlock()
		r.log.Trace("State out of sync, forwarding query to leader")
		return r.forwardQuery(request, leader, server)
	}

	// If the session's consistency level is SEQUENTIAL, handle the request here, otherwise forward it.
	if request.ReadConsistency == raft.ReadConsistency_SEQUENTIAL {
		// If the commit index is not in the log then we've fallen too far behind the leader to perform a local query.
		// Forward the request to the leader.
		if r.store.Writer().LastIndex() < r.raft.CommitIndex() {
			r.raft.ReadUnlock()
			r.log.Trace("State out of sync, forwarding query to leader")
			return r.forwardQuery(request, leader, server)
		}

		entry := &log.LogEntry{
			Index: r.store.Writer().LastIndex(),
			Entry: &raft.RaftLogEntry{
				Term:      r.raft.Term(),
				Timestamp: time.Now(),
				Entry: &raft.RaftLogEntry_Query{
					Query: &raft.QueryEntry{
						Value: request.Value,
					},
				},
			},
		}

		// Release the read lock before applying the entry.
		r.raft.ReadUnlock()

		return r.applyQuery(entry, server)
	}
	r.raft.ReadUnlock()
	return r.forwardQuery(request, leader, server)
}

// applyQuery applies a query to the state machine
func (r *PassiveRole) applyQuery(entry *log.LogEntry, server raft.RaftService_QueryServer) error {
	// Create a result channel
	ch := make(chan node.Output)

	// Apply the entry to the state machine
	r.state.ApplyEntry(entry, ch)

	// Iterate through results and translate them into QueryResponses.
	for result := range ch {
		if result.Succeeded() {
			response := &raft.QueryResponse{
				Status: raft.ResponseStatus_OK,
				Output: result.Value,
			}
			err := r.log.Response("QueryResponse", response, server.Send(response))
			if err != nil {
				return err
			}
		} else {
			response := &raft.QueryResponse{
				Status:  raft.ResponseStatus_ERROR,
				Message: result.Error.Error(),
			}
			err := r.log.Response("QueryResponse", response, server.Send(response))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// forwardQuery forwards a query request to the leader
func (r *PassiveRole) forwardQuery(request *raft.QueryRequest, leader raft.MemberID, server raft.RaftService_QueryServer) error {
	if leader == "" {
		response := &raft.QueryResponse{
			Status: raft.ResponseStatus_ERROR,
			Error:  raft.RaftError_NO_LEADER,
		}
		return r.log.Response("QueryResponse", response, server.Send(response))
	}

	r.log.Trace("Forwarding %v", request)
	client, err := r.raft.Connect(leader)
	if err != nil {
		return r.log.Response("QueryResponse", nil, err)
	}

	stream, err := client.Query(context.Background(), request)
	if err != nil {
		return r.log.Response("QueryResponse", nil, err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return r.log.Response("QueryResponse", nil, err)
		}
		r.log.Request("QueryResponse", response)
		_ = r.log.Response("QueryResponse", response, server.Send(response))
	}
	return nil
}