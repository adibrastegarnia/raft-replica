// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/atomix/raft/protocol/raft.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	config "github.com/atomix/raft-replica/pkg/atomix/raft/config"
	protocol "github.com/atomix/raft-replica/pkg/atomix/raft/protocol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRaft is a mock of Raft interface
type MockRaft struct {
	ctrl     *gomock.Controller
	recorder *MockRaftMockRecorder
}

// MockRaftMockRecorder is the mock recorder for MockRaft
type MockRaftMockRecorder struct {
	mock *MockRaft
}

// NewMockRaft creates a new mock instance
func NewMockRaft(ctrl *gomock.Controller) *MockRaft {
	mock := &MockRaft{ctrl: ctrl}
	mock.recorder = &MockRaftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRaft) EXPECT() *MockRaftMockRecorder {
	return m.recorder
}

// Join mocks base method
func (m *MockRaft) Join(ctx context.Context, request *protocol.JoinRequest) (*protocol.JoinResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", ctx, request)
	ret0, _ := ret[0].(*protocol.JoinResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join
func (mr *MockRaftMockRecorder) Join(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockRaft)(nil).Join), ctx, request)
}

// Leave mocks base method
func (m *MockRaft) Leave(ctx context.Context, request *protocol.LeaveRequest) (*protocol.LeaveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave", ctx, request)
	ret0, _ := ret[0].(*protocol.LeaveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leave indicates an expected call of Leave
func (mr *MockRaftMockRecorder) Leave(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockRaft)(nil).Leave), ctx, request)
}

// Configure mocks base method
func (m *MockRaft) Configure(ctx context.Context, request *protocol.ConfigureRequest) (*protocol.ConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", ctx, request)
	ret0, _ := ret[0].(*protocol.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockRaftMockRecorder) Configure(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockRaft)(nil).Configure), ctx, request)
}

// Reconfigure mocks base method
func (m *MockRaft) Reconfigure(ctx context.Context, request *protocol.ReconfigureRequest) (*protocol.ReconfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconfigure", ctx, request)
	ret0, _ := ret[0].(*protocol.ReconfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reconfigure indicates an expected call of Reconfigure
func (mr *MockRaftMockRecorder) Reconfigure(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconfigure", reflect.TypeOf((*MockRaft)(nil).Reconfigure), ctx, request)
}

// Poll mocks base method
func (m *MockRaft) Poll(ctx context.Context, request *protocol.PollRequest) (*protocol.PollResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Poll", ctx, request)
	ret0, _ := ret[0].(*protocol.PollResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Poll indicates an expected call of Poll
func (mr *MockRaftMockRecorder) Poll(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Poll", reflect.TypeOf((*MockRaft)(nil).Poll), ctx, request)
}

// Vote mocks base method
func (m *MockRaft) Vote(ctx context.Context, request *protocol.VoteRequest) (*protocol.VoteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vote", ctx, request)
	ret0, _ := ret[0].(*protocol.VoteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vote indicates an expected call of Vote
func (mr *MockRaftMockRecorder) Vote(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockRaft)(nil).Vote), ctx, request)
}

// Transfer mocks base method
func (m *MockRaft) Transfer(ctx context.Context, request *protocol.TransferRequest) (*protocol.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, request)
	ret0, _ := ret[0].(*protocol.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer
func (mr *MockRaftMockRecorder) Transfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockRaft)(nil).Transfer), ctx, request)
}

// Append mocks base method
func (m *MockRaft) Append(ctx context.Context, request *protocol.AppendRequest) (*protocol.AppendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", ctx, request)
	ret0, _ := ret[0].(*protocol.AppendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Append indicates an expected call of Append
func (mr *MockRaftMockRecorder) Append(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockRaft)(nil).Append), ctx, request)
}

// Install mocks base method
func (m *MockRaft) Install(ch <-chan *protocol.InstallStreamRequest) (*protocol.InstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ch)
	ret0, _ := ret[0].(*protocol.InstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install
func (mr *MockRaftMockRecorder) Install(ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockRaft)(nil).Install), ch)
}

// Command mocks base method
func (m *MockRaft) Command(request *protocol.CommandRequest, ch chan<- *protocol.CommandStreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", request, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Command indicates an expected call of Command
func (mr *MockRaftMockRecorder) Command(request, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockRaft)(nil).Command), request, ch)
}

// Query mocks base method
func (m *MockRaft) Query(request *protocol.QueryRequest, ch chan<- *protocol.QueryStreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", request, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockRaftMockRecorder) Query(request, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRaft)(nil).Query), request, ch)
}

// Init mocks base method
func (m *MockRaft) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockRaftMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRaft)(nil).Init))
}

// Role mocks base method
func (m *MockRaft) Role() protocol.RoleType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Role")
	ret0, _ := ret[0].(protocol.RoleType)
	return ret0
}

// Role indicates an expected call of Role
func (mr *MockRaftMockRecorder) Role() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Role", reflect.TypeOf((*MockRaft)(nil).Role))
}

// WatchRole mocks base method
func (m *MockRaft) WatchRole(arg0 func(protocol.RoleType)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WatchRole", arg0)
}

// WatchRole indicates an expected call of WatchRole
func (mr *MockRaftMockRecorder) WatchRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRole", reflect.TypeOf((*MockRaft)(nil).WatchRole), arg0)
}

// Status mocks base method
func (m *MockRaft) Status() protocol.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(protocol.Status)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockRaftMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRaft)(nil).Status))
}

// WatchStatus mocks base method
func (m *MockRaft) WatchStatus(arg0 func(protocol.Status)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WatchStatus", arg0)
}

// WatchStatus indicates an expected call of WatchStatus
func (mr *MockRaftMockRecorder) WatchStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchStatus", reflect.TypeOf((*MockRaft)(nil).WatchStatus), arg0)
}

// Config mocks base method
func (m *MockRaft) Config() *config.ProtocolConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.ProtocolConfig)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockRaftMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockRaft)(nil).Config))
}

// Member mocks base method
func (m *MockRaft) Member() protocol.MemberID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Member")
	ret0, _ := ret[0].(protocol.MemberID)
	return ret0
}

// Member indicates an expected call of Member
func (mr *MockRaftMockRecorder) Member() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Member", reflect.TypeOf((*MockRaft)(nil).Member))
}

// Members mocks base method
func (m *MockRaft) Members() []protocol.MemberID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]protocol.MemberID)
	return ret0
}

// Members indicates an expected call of Members
func (mr *MockRaftMockRecorder) Members() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockRaft)(nil).Members))
}

// GetMember mocks base method
func (m *MockRaft) GetMember(memberID protocol.MemberID) *protocol.Member {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMember", memberID)
	ret0, _ := ret[0].(*protocol.Member)
	return ret0
}

// GetMember indicates an expected call of GetMember
func (mr *MockRaftMockRecorder) GetMember(memberID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockRaft)(nil).GetMember), memberID)
}

// Protocol mocks base method
func (m *MockRaft) Protocol() protocol.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(protocol.Client)
	return ret0
}

// Protocol indicates an expected call of Protocol
func (mr *MockRaftMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockRaft)(nil).Protocol))
}

// Term mocks base method
func (m *MockRaft) Term() protocol.Term {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Term")
	ret0, _ := ret[0].(protocol.Term)
	return ret0
}

// Term indicates an expected call of Term
func (mr *MockRaftMockRecorder) Term() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Term", reflect.TypeOf((*MockRaft)(nil).Term))
}

// SetTerm mocks base method
func (m *MockRaft) SetTerm(term protocol.Term) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTerm", term)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTerm indicates an expected call of SetTerm
func (mr *MockRaftMockRecorder) SetTerm(term interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTerm", reflect.TypeOf((*MockRaft)(nil).SetTerm), term)
}

// Leader mocks base method
func (m *MockRaft) Leader() *protocol.MemberID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leader")
	ret0, _ := ret[0].(*protocol.MemberID)
	return ret0
}

// Leader indicates an expected call of Leader
func (mr *MockRaftMockRecorder) Leader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockRaft)(nil).Leader))
}

// SetLeader mocks base method
func (m *MockRaft) SetLeader(leader *protocol.MemberID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLeader", leader)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLeader indicates an expected call of SetLeader
func (mr *MockRaftMockRecorder) SetLeader(leader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLeader", reflect.TypeOf((*MockRaft)(nil).SetLeader), leader)
}

// LastVotedFor mocks base method
func (m *MockRaft) LastVotedFor() *protocol.MemberID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastVotedFor")
	ret0, _ := ret[0].(*protocol.MemberID)
	return ret0
}

// LastVotedFor indicates an expected call of LastVotedFor
func (mr *MockRaftMockRecorder) LastVotedFor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastVotedFor", reflect.TypeOf((*MockRaft)(nil).LastVotedFor))
}

// SetLastVotedFor mocks base method
func (m *MockRaft) SetLastVotedFor(memberID protocol.MemberID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLastVotedFor", memberID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLastVotedFor indicates an expected call of SetLastVotedFor
func (mr *MockRaftMockRecorder) SetLastVotedFor(memberID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastVotedFor", reflect.TypeOf((*MockRaft)(nil).SetLastVotedFor), memberID)
}

// CommitIndex mocks base method
func (m *MockRaft) CommitIndex() protocol.Index {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitIndex")
	ret0, _ := ret[0].(protocol.Index)
	return ret0
}

// CommitIndex indicates an expected call of CommitIndex
func (mr *MockRaftMockRecorder) CommitIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitIndex", reflect.TypeOf((*MockRaft)(nil).CommitIndex))
}

// SetCommitIndex mocks base method
func (m *MockRaft) SetCommitIndex(index protocol.Index) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCommitIndex", index)
}

// SetCommitIndex indicates an expected call of SetCommitIndex
func (mr *MockRaftMockRecorder) SetCommitIndex(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCommitIndex", reflect.TypeOf((*MockRaft)(nil).SetCommitIndex), index)
}

// Commit mocks base method
func (m *MockRaft) Commit(index protocol.Index) protocol.Index {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", index)
	ret0, _ := ret[0].(protocol.Index)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockRaftMockRecorder) Commit(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockRaft)(nil).Commit), index)
}

// WriteLock mocks base method
func (m *MockRaft) WriteLock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteLock")
}

// WriteLock indicates an expected call of WriteLock
func (mr *MockRaftMockRecorder) WriteLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLock", reflect.TypeOf((*MockRaft)(nil).WriteLock))
}

// WriteUnlock mocks base method
func (m *MockRaft) WriteUnlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteUnlock")
}

// WriteUnlock indicates an expected call of WriteUnlock
func (mr *MockRaftMockRecorder) WriteUnlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUnlock", reflect.TypeOf((*MockRaft)(nil).WriteUnlock))
}

// ReadLock mocks base method
func (m *MockRaft) ReadLock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReadLock")
}

// ReadLock indicates an expected call of ReadLock
func (mr *MockRaftMockRecorder) ReadLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLock", reflect.TypeOf((*MockRaft)(nil).ReadLock))
}

// ReadUnlock mocks base method
func (m *MockRaft) ReadUnlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReadUnlock")
}

// ReadUnlock indicates an expected call of ReadUnlock
func (mr *MockRaftMockRecorder) ReadUnlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUnlock", reflect.TypeOf((*MockRaft)(nil).ReadUnlock))
}

// SetRole mocks base method
func (m *MockRaft) SetRole(role protocol.Role) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRole", role)
}

// SetRole indicates an expected call of SetRole
func (mr *MockRaftMockRecorder) SetRole(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRole", reflect.TypeOf((*MockRaft)(nil).SetRole), role)
}

// Close mocks base method
func (m *MockRaft) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRaftMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRaft)(nil).Close))
}

// MockRole is a mock of Role interface
type MockRole struct {
	ctrl     *gomock.Controller
	recorder *MockRoleMockRecorder
}

// MockRoleMockRecorder is the mock recorder for MockRole
type MockRoleMockRecorder struct {
	mock *MockRole
}

// NewMockRole creates a new mock instance
func NewMockRole(ctrl *gomock.Controller) *MockRole {
	mock := &MockRole{ctrl: ctrl}
	mock.recorder = &MockRoleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRole) EXPECT() *MockRoleMockRecorder {
	return m.recorder
}

// Join mocks base method
func (m *MockRole) Join(ctx context.Context, request *protocol.JoinRequest) (*protocol.JoinResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", ctx, request)
	ret0, _ := ret[0].(*protocol.JoinResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join
func (mr *MockRoleMockRecorder) Join(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockRole)(nil).Join), ctx, request)
}

// Leave mocks base method
func (m *MockRole) Leave(ctx context.Context, request *protocol.LeaveRequest) (*protocol.LeaveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave", ctx, request)
	ret0, _ := ret[0].(*protocol.LeaveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leave indicates an expected call of Leave
func (mr *MockRoleMockRecorder) Leave(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockRole)(nil).Leave), ctx, request)
}

// Configure mocks base method
func (m *MockRole) Configure(ctx context.Context, request *protocol.ConfigureRequest) (*protocol.ConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", ctx, request)
	ret0, _ := ret[0].(*protocol.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockRoleMockRecorder) Configure(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockRole)(nil).Configure), ctx, request)
}

// Reconfigure mocks base method
func (m *MockRole) Reconfigure(ctx context.Context, request *protocol.ReconfigureRequest) (*protocol.ReconfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconfigure", ctx, request)
	ret0, _ := ret[0].(*protocol.ReconfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reconfigure indicates an expected call of Reconfigure
func (mr *MockRoleMockRecorder) Reconfigure(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconfigure", reflect.TypeOf((*MockRole)(nil).Reconfigure), ctx, request)
}

// Poll mocks base method
func (m *MockRole) Poll(ctx context.Context, request *protocol.PollRequest) (*protocol.PollResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Poll", ctx, request)
	ret0, _ := ret[0].(*protocol.PollResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Poll indicates an expected call of Poll
func (mr *MockRoleMockRecorder) Poll(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Poll", reflect.TypeOf((*MockRole)(nil).Poll), ctx, request)
}

// Vote mocks base method
func (m *MockRole) Vote(ctx context.Context, request *protocol.VoteRequest) (*protocol.VoteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vote", ctx, request)
	ret0, _ := ret[0].(*protocol.VoteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vote indicates an expected call of Vote
func (mr *MockRoleMockRecorder) Vote(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockRole)(nil).Vote), ctx, request)
}

// Transfer mocks base method
func (m *MockRole) Transfer(ctx context.Context, request *protocol.TransferRequest) (*protocol.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, request)
	ret0, _ := ret[0].(*protocol.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer
func (mr *MockRoleMockRecorder) Transfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockRole)(nil).Transfer), ctx, request)
}

// Append mocks base method
func (m *MockRole) Append(ctx context.Context, request *protocol.AppendRequest) (*protocol.AppendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", ctx, request)
	ret0, _ := ret[0].(*protocol.AppendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Append indicates an expected call of Append
func (mr *MockRoleMockRecorder) Append(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockRole)(nil).Append), ctx, request)
}

// Install mocks base method
func (m *MockRole) Install(ch <-chan *protocol.InstallStreamRequest) (*protocol.InstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ch)
	ret0, _ := ret[0].(*protocol.InstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install
func (mr *MockRoleMockRecorder) Install(ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockRole)(nil).Install), ch)
}

// Command mocks base method
func (m *MockRole) Command(request *protocol.CommandRequest, ch chan<- *protocol.CommandStreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", request, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Command indicates an expected call of Command
func (mr *MockRoleMockRecorder) Command(request, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockRole)(nil).Command), request, ch)
}

// Query mocks base method
func (m *MockRole) Query(request *protocol.QueryRequest, ch chan<- *protocol.QueryStreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", request, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockRoleMockRecorder) Query(request, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRole)(nil).Query), request, ch)
}

// Type mocks base method
func (m *MockRole) Type() protocol.RoleType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(protocol.RoleType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockRoleMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockRole)(nil).Type))
}

// Start mocks base method
func (m *MockRole) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockRoleMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRole)(nil).Start))
}

// Stop mocks base method
func (m *MockRole) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockRoleMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRole)(nil).Stop))
}
