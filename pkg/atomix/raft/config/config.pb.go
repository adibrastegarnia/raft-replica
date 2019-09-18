// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/raft/config/config.proto

package config

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StorageLevel int32

const (
	StorageLevel_DISK   StorageLevel = 0
	StorageLevel_MAPPED StorageLevel = 1
)

var StorageLevel_name = map[int32]string{
	0: "DISK",
	1: "MAPPED",
}

var StorageLevel_value = map[string]int32{
	"DISK":   0,
	"MAPPED": 1,
}

func (x StorageLevel) String() string {
	return proto.EnumName(StorageLevel_name, int32(x))
}

func (StorageLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e09be49defe43eb0, []int{0}
}

type ProtocolConfig struct {
	ElectionTimeout   *time.Duration    `protobuf:"bytes,1,opt,name=election_timeout,json=electionTimeout,proto3,stdduration" json:"election_timeout,omitempty"`
	HeartbeatInterval *time.Duration    `protobuf:"bytes,2,opt,name=heartbeat_interval,json=heartbeatInterval,proto3,stdduration" json:"heartbeat_interval,omitempty"`
	Storage           *StorageConfig    `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Compaction        *CompactionConfig `protobuf:"bytes,4,opt,name=compaction,proto3" json:"compaction,omitempty"`
}

func (m *ProtocolConfig) Reset()         { *m = ProtocolConfig{} }
func (m *ProtocolConfig) String() string { return proto.CompactTextString(m) }
func (*ProtocolConfig) ProtoMessage()    {}
func (*ProtocolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e09be49defe43eb0, []int{0}
}
func (m *ProtocolConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolConfig.Merge(m, src)
}
func (m *ProtocolConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolConfig proto.InternalMessageInfo

func (m *ProtocolConfig) GetElectionTimeout() *time.Duration {
	if m != nil {
		return m.ElectionTimeout
	}
	return nil
}

func (m *ProtocolConfig) GetHeartbeatInterval() *time.Duration {
	if m != nil {
		return m.HeartbeatInterval
	}
	return nil
}

func (m *ProtocolConfig) GetStorage() *StorageConfig {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ProtocolConfig) GetCompaction() *CompactionConfig {
	if m != nil {
		return m.Compaction
	}
	return nil
}

type StorageConfig struct {
	Directory     string       `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	Level         StorageLevel `protobuf:"varint,2,opt,name=level,proto3,enum=atomix.raft.config.StorageLevel" json:"level,omitempty"`
	MaxEntrySize  uint32       `protobuf:"varint,3,opt,name=max_entry_size,json=maxEntrySize,proto3" json:"max_entry_size,omitempty"`
	SegmentSize   uint32       `protobuf:"varint,4,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	FlushOnCommit bool         `protobuf:"varint,5,opt,name=flush_on_commit,json=flushOnCommit,proto3" json:"flush_on_commit,omitempty"`
}

func (m *StorageConfig) Reset()         { *m = StorageConfig{} }
func (m *StorageConfig) String() string { return proto.CompactTextString(m) }
func (*StorageConfig) ProtoMessage()    {}
func (*StorageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e09be49defe43eb0, []int{1}
}
func (m *StorageConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StorageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StorageConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StorageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageConfig.Merge(m, src)
}
func (m *StorageConfig) XXX_Size() int {
	return m.Size()
}
func (m *StorageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StorageConfig proto.InternalMessageInfo

func (m *StorageConfig) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *StorageConfig) GetLevel() StorageLevel {
	if m != nil {
		return m.Level
	}
	return StorageLevel_DISK
}

func (m *StorageConfig) GetMaxEntrySize() uint32 {
	if m != nil {
		return m.MaxEntrySize
	}
	return 0
}

func (m *StorageConfig) GetSegmentSize() uint32 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *StorageConfig) GetFlushOnCommit() bool {
	if m != nil {
		return m.FlushOnCommit
	}
	return false
}

type CompactionConfig struct {
	Dynamic          bool    `protobuf:"varint,1,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	FreeDiskBuffer   float32 `protobuf:"fixed32,2,opt,name=free_disk_buffer,json=freeDiskBuffer,proto3" json:"free_disk_buffer,omitempty"`
	FreeMemoryBuffer float32 `protobuf:"fixed32,3,opt,name=free_memory_buffer,json=freeMemoryBuffer,proto3" json:"free_memory_buffer,omitempty"`
}

func (m *CompactionConfig) Reset()         { *m = CompactionConfig{} }
func (m *CompactionConfig) String() string { return proto.CompactTextString(m) }
func (*CompactionConfig) ProtoMessage()    {}
func (*CompactionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e09be49defe43eb0, []int{2}
}
func (m *CompactionConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompactionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactionConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompactionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionConfig.Merge(m, src)
}
func (m *CompactionConfig) XXX_Size() int {
	return m.Size()
}
func (m *CompactionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionConfig proto.InternalMessageInfo

func (m *CompactionConfig) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *CompactionConfig) GetFreeDiskBuffer() float32 {
	if m != nil {
		return m.FreeDiskBuffer
	}
	return 0
}

func (m *CompactionConfig) GetFreeMemoryBuffer() float32 {
	if m != nil {
		return m.FreeMemoryBuffer
	}
	return 0
}

func init() {
	proto.RegisterEnum("atomix.raft.config.StorageLevel", StorageLevel_name, StorageLevel_value)
	proto.RegisterType((*ProtocolConfig)(nil), "atomix.raft.config.ProtocolConfig")
	proto.RegisterType((*StorageConfig)(nil), "atomix.raft.config.StorageConfig")
	proto.RegisterType((*CompactionConfig)(nil), "atomix.raft.config.CompactionConfig")
}

func init() { proto.RegisterFile("atomix/raft/config/config.proto", fileDescriptor_e09be49defe43eb0) }

var fileDescriptor_e09be49defe43eb0 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0x12, 0x41,
	0x1c, 0xc6, 0x19, 0x4a, 0x5b, 0xfa, 0x2f, 0x50, 0x9c, 0x78, 0x58, 0x1b, 0x33, 0x50, 0x42, 0x0c,
	0x31, 0x66, 0x49, 0x6a, 0xe2, 0xc5, 0x93, 0x40, 0x0f, 0x55, 0xab, 0x64, 0xf1, 0xbe, 0x19, 0x96,
	0xd9, 0xed, 0xa4, 0x3b, 0x3b, 0x64, 0x76, 0x68, 0xa0, 0x67, 0x1f, 0xc0, 0xa3, 0x8f, 0xe0, 0xcd,
	0xab, 0x8f, 0xe0, 0xb1, 0x27, 0xe3, 0x4d, 0x85, 0x97, 0xf0, 0x68, 0x76, 0x66, 0xa9, 0xb5, 0x1a,
	0xd3, 0x13, 0xc3, 0x37, 0xbf, 0xef, 0xfb, 0xcf, 0xf7, 0xcf, 0x42, 0x83, 0x6a, 0x29, 0xf8, 0xbc,
	0xab, 0x68, 0xa8, 0xbb, 0x81, 0x4c, 0x42, 0x1e, 0xe5, 0x3f, 0xee, 0x54, 0x49, 0x2d, 0x31, 0xb6,
	0x80, 0x9b, 0x01, 0xae, 0xbd, 0xd9, 0x27, 0x91, 0x94, 0x51, 0xcc, 0xba, 0x86, 0x18, 0xcf, 0xc2,
	0xee, 0x64, 0xa6, 0xa8, 0xe6, 0x32, 0xb1, 0x9e, 0xfd, 0xc6, 0xcd, 0x7b, 0xcd, 0x05, 0x4b, 0x35,
	0x15, 0xd3, 0x1c, 0xb8, 0x1b, 0xc9, 0x48, 0x9a, 0x63, 0x37, 0x3b, 0x59, 0xb5, 0xf5, 0xb1, 0x08,
	0xb5, 0x61, 0x76, 0x0a, 0x64, 0xdc, 0x37, 0x93, 0xf0, 0x73, 0xa8, 0xb3, 0x98, 0x05, 0x59, 0xb6,
	0x9f, 0x85, 0xc8, 0x99, 0x76, 0x50, 0x13, 0x75, 0x76, 0x0f, 0xef, 0xb9, 0x76, 0x88, 0xbb, 0x1e,
	0xe2, 0x0e, 0xf2, 0x47, 0xf4, 0x4a, 0xef, 0xbf, 0x35, 0x90, 0xb7, 0xb7, 0x36, 0xbe, 0xb1, 0x3e,
	0xfc, 0x0a, 0xf0, 0x29, 0xa3, 0x4a, 0x8f, 0x19, 0xd5, 0x3e, 0x4f, 0x34, 0x53, 0xe7, 0x34, 0x76,
	0x8a, 0xb7, 0x4b, 0xbb, 0x73, 0x65, 0x3d, 0xce, 0x9d, 0xf8, 0x29, 0x6c, 0xa7, 0x5a, 0x2a, 0x1a,
	0x31, 0x67, 0xc3, 0x84, 0x1c, 0xb8, 0x7f, 0xef, 0xca, 0x1d, 0x59, 0xc4, 0xf6, 0xf1, 0xd6, 0x0e,
	0x3c, 0x00, 0x08, 0xa4, 0x98, 0x52, 0xf3, 0x42, 0xa7, 0x64, 0xfc, 0xed, 0x7f, 0xf9, 0xfb, 0x57,
	0x54, 0x1e, 0x71, 0xcd, 0xd7, 0xfa, 0x82, 0xa0, 0xfa, 0xc7, 0x00, 0x7c, 0x1f, 0x76, 0x26, 0x5c,
	0xb1, 0x40, 0x4b, 0xb5, 0x30, 0x9b, 0xda, 0xf1, 0x7e, 0x0b, 0xf8, 0x09, 0x6c, 0xc6, 0xec, 0x9c,
	0xd9, 0xd6, 0xb5, 0xc3, 0xe6, 0x7f, 0x1e, 0xfc, 0x32, 0xe3, 0x3c, 0x8b, 0xe3, 0x36, 0xd4, 0x04,
	0x9d, 0xfb, 0x2c, 0xd1, 0x6a, 0xe1, 0xa7, 0xfc, 0xc2, 0x36, 0xae, 0x7a, 0x15, 0x41, 0xe7, 0x47,
	0x99, 0x38, 0xe2, 0x17, 0x0c, 0x1f, 0x40, 0x25, 0x65, 0x91, 0x60, 0x89, 0xb6, 0x4c, 0xc9, 0x30,
	0xbb, 0xb9, 0x66, 0x90, 0x07, 0xb0, 0x17, 0xc6, 0xb3, 0xf4, 0xd4, 0x97, 0x89, 0x1f, 0x48, 0x21,
	0xb8, 0x76, 0x36, 0x9b, 0xa8, 0x53, 0xf6, 0xaa, 0x46, 0x7e, 0x9d, 0xf4, 0x8d, 0xd8, 0x7a, 0x8b,
	0xa0, 0x7e, 0xb3, 0x39, 0x76, 0x60, 0x7b, 0xb2, 0x48, 0xa8, 0xe0, 0x81, 0x69, 0x56, 0xf6, 0xd6,
	0x7f, 0x71, 0x07, 0xea, 0xa1, 0x62, 0xcc, 0x9f, 0xf0, 0xf4, 0xcc, 0x1f, 0xcf, 0xc2, 0x90, 0x29,
	0x53, 0xb1, 0xe8, 0xd5, 0x32, 0x7d, 0xc0, 0xd3, 0xb3, 0x9e, 0x51, 0xf1, 0x23, 0xc0, 0x86, 0x14,
	0x4c, 0x48, 0xb5, 0x58, 0xb3, 0x1b, 0x86, 0x35, 0x19, 0x27, 0xe6, 0xc2, 0xd2, 0x0f, 0xdb, 0x50,
	0xb9, 0xbe, 0x0e, 0x5c, 0x86, 0xd2, 0xe0, 0x78, 0xf4, 0xa2, 0x5e, 0xc0, 0x00, 0x5b, 0x27, 0xcf,
	0x86, 0xc3, 0xa3, 0x41, 0x1d, 0xf5, 0xda, 0x3f, 0x7f, 0x10, 0xf4, 0x61, 0x49, 0xd0, 0xa7, 0x25,
	0x41, 0x9f, 0x97, 0x04, 0x5d, 0x2e, 0x09, 0xfa, 0xbe, 0x24, 0xe8, 0xdd, 0x8a, 0x14, 0x2e, 0x57,
	0xa4, 0xf0, 0x75, 0x45, 0x0a, 0xe3, 0x2d, 0xf3, 0x69, 0x3d, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x60, 0xee, 0xb8, 0x34, 0x72, 0x03, 0x00, 0x00,
}

func (this *ProtocolConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProtocolConfig)
	if !ok {
		that2, ok := that.(ProtocolConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ElectionTimeout != nil && that1.ElectionTimeout != nil {
		if *this.ElectionTimeout != *that1.ElectionTimeout {
			return false
		}
	} else if this.ElectionTimeout != nil {
		return false
	} else if that1.ElectionTimeout != nil {
		return false
	}
	if this.HeartbeatInterval != nil && that1.HeartbeatInterval != nil {
		if *this.HeartbeatInterval != *that1.HeartbeatInterval {
			return false
		}
	} else if this.HeartbeatInterval != nil {
		return false
	} else if that1.HeartbeatInterval != nil {
		return false
	}
	if !this.Storage.Equal(that1.Storage) {
		return false
	}
	if !this.Compaction.Equal(that1.Compaction) {
		return false
	}
	return true
}
func (this *StorageConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StorageConfig)
	if !ok {
		that2, ok := that.(StorageConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Directory != that1.Directory {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if this.MaxEntrySize != that1.MaxEntrySize {
		return false
	}
	if this.SegmentSize != that1.SegmentSize {
		return false
	}
	if this.FlushOnCommit != that1.FlushOnCommit {
		return false
	}
	return true
}
func (this *CompactionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompactionConfig)
	if !ok {
		that2, ok := that.(CompactionConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Dynamic != that1.Dynamic {
		return false
	}
	if this.FreeDiskBuffer != that1.FreeDiskBuffer {
		return false
	}
	if this.FreeMemoryBuffer != that1.FreeMemoryBuffer {
		return false
	}
	return true
}
func (m *ProtocolConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Compaction != nil {
		{
			size, err := m.Compaction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Storage != nil {
		{
			size, err := m.Storage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HeartbeatInterval != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.HeartbeatInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.HeartbeatInterval):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintConfig(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x12
	}
	if m.ElectionTimeout != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.ElectionTimeout, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ElectionTimeout):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintConfig(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StorageConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StorageConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FlushOnCommit {
		i--
		if m.FlushOnCommit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.SegmentSize != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.SegmentSize))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxEntrySize != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxEntrySize))
		i--
		dAtA[i] = 0x18
	}
	if m.Level != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Directory) > 0 {
		i -= len(m.Directory)
		copy(dAtA[i:], m.Directory)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Directory)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CompactionConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactionConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactionConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FreeMemoryBuffer != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.FreeMemoryBuffer))))
		i--
		dAtA[i] = 0x1d
	}
	if m.FreeDiskBuffer != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.FreeDiskBuffer))))
		i--
		dAtA[i] = 0x15
	}
	if m.Dynamic {
		i--
		if m.Dynamic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedProtocolConfig(r randyConfig, easy bool) *ProtocolConfig {
	this := &ProtocolConfig{}
	if r.Intn(5) != 0 {
		this.ElectionTimeout = github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	if r.Intn(5) != 0 {
		this.HeartbeatInterval = github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Storage = NewPopulatedStorageConfig(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Compaction = NewPopulatedCompactionConfig(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStorageConfig(r randyConfig, easy bool) *StorageConfig {
	this := &StorageConfig{}
	this.Directory = string(randStringConfig(r))
	this.Level = StorageLevel([]int32{0, 1}[r.Intn(2)])
	this.MaxEntrySize = uint32(r.Uint32())
	this.SegmentSize = uint32(r.Uint32())
	this.FlushOnCommit = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCompactionConfig(r randyConfig, easy bool) *CompactionConfig {
	this := &CompactionConfig{}
	this.Dynamic = bool(bool(r.Intn(2) == 0))
	this.FreeDiskBuffer = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.FreeDiskBuffer *= -1
	}
	this.FreeMemoryBuffer = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.FreeMemoryBuffer *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyConfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneConfig(r randyConfig) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringConfig(r randyConfig) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneConfig(r)
	}
	return string(tmps)
}
func randUnrecognizedConfig(r randyConfig, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldConfig(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldConfig(dAtA []byte, r randyConfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateConfig(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ProtocolConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ElectionTimeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ElectionTimeout)
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.HeartbeatInterval != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.HeartbeatInterval)
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Storage != nil {
		l = m.Storage.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Compaction != nil {
		l = m.Compaction.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *StorageConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Directory)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovConfig(uint64(m.Level))
	}
	if m.MaxEntrySize != 0 {
		n += 1 + sovConfig(uint64(m.MaxEntrySize))
	}
	if m.SegmentSize != 0 {
		n += 1 + sovConfig(uint64(m.SegmentSize))
	}
	if m.FlushOnCommit {
		n += 2
	}
	return n
}

func (m *CompactionConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dynamic {
		n += 2
	}
	if m.FreeDiskBuffer != 0 {
		n += 5
	}
	if m.FreeMemoryBuffer != 0 {
		n += 5
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtocolConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElectionTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ElectionTimeout == nil {
				m.ElectionTimeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.ElectionTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HeartbeatInterval == nil {
				m.HeartbeatInterval = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.HeartbeatInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Storage == nil {
				m.Storage = &StorageConfig{}
			}
			if err := m.Storage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compaction == nil {
				m.Compaction = &CompactionConfig{}
			}
			if err := m.Compaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StorageConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StorageConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Directory", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Directory = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= StorageLevel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEntrySize", wireType)
			}
			m.MaxEntrySize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxEntrySize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentSize", wireType)
			}
			m.SegmentSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SegmentSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushOnCommit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FlushOnCommit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompactionConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CompactionConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactionConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dynamic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Dynamic = bool(v != 0)
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreeDiskBuffer", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.FreeDiskBuffer = float32(math.Float32frombits(v))
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreeMemoryBuffer", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.FreeMemoryBuffer = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)