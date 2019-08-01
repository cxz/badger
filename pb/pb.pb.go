// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ManifestChange_Operation int32

const (
	ManifestChange_CREATE ManifestChange_Operation = 0
	ManifestChange_DELETE ManifestChange_Operation = 1
)

var ManifestChange_Operation_name = map[int32]string{
	0: "CREATE",
	1: "DELETE",
}

var ManifestChange_Operation_value = map[string]int32{
	"CREATE": 0,
	"DELETE": 1,
}

func (x ManifestChange_Operation) String() string {
	return proto.EnumName(ManifestChange_Operation_name, int32(x))
}

func (ManifestChange_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3, 0}
}

type KV struct {
	Key       []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	UserMeta  []byte `protobuf:"bytes,3,opt,name=user_meta,json=userMeta,proto3" json:"user_meta,omitempty"`
	Version   uint64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	ExpiresAt uint64 `protobuf:"varint,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Meta      []byte `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	// Stream id is used to identify which stream the KV came from.
	StreamId             uint32   `protobuf:"varint,10,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KV.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return m.Size()
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KV) GetUserMeta() []byte {
	if m != nil {
		return m.UserMeta
	}
	return nil
}

func (m *KV) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *KV) GetExpiresAt() uint64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *KV) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *KV) GetStreamId() uint32 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

type KVList struct {
	Kv                   []*KV    `protobuf:"bytes,1,rep,name=kv,proto3" json:"kv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVList) Reset()         { *m = KVList{} }
func (m *KVList) String() string { return proto.CompactTextString(m) }
func (*KVList) ProtoMessage()    {}
func (*KVList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}
func (m *KVList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KVList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KVList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVList.Merge(m, src)
}
func (m *KVList) XXX_Size() int {
	return m.Size()
}
func (m *KVList) XXX_DiscardUnknown() {
	xxx_messageInfo_KVList.DiscardUnknown(m)
}

var xxx_messageInfo_KVList proto.InternalMessageInfo

func (m *KVList) GetKv() []*KV {
	if m != nil {
		return m.Kv
	}
	return nil
}

type ManifestChangeSet struct {
	// A set of changes that are applied atomically.
	Changes              []*ManifestChange `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ManifestChangeSet) Reset()         { *m = ManifestChangeSet{} }
func (m *ManifestChangeSet) String() string { return proto.CompactTextString(m) }
func (*ManifestChangeSet) ProtoMessage()    {}
func (*ManifestChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}
func (m *ManifestChangeSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestChangeSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestChangeSet.Merge(m, src)
}
func (m *ManifestChangeSet) XXX_Size() int {
	return m.Size()
}
func (m *ManifestChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestChangeSet proto.InternalMessageInfo

func (m *ManifestChangeSet) GetChanges() []*ManifestChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ManifestChange struct {
	Id                   uint64                   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Op                   ManifestChange_Operation `protobuf:"varint,2,opt,name=Op,proto3,enum=pb.ManifestChange_Operation" json:"Op,omitempty"`
	Level                uint32                   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Checksum             []byte                   `protobuf:"bytes,4,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ManifestChange) Reset()         { *m = ManifestChange{} }
func (m *ManifestChange) String() string { return proto.CompactTextString(m) }
func (*ManifestChange) ProtoMessage()    {}
func (*ManifestChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3}
}
func (m *ManifestChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManifestChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestChange.Merge(m, src)
}
func (m *ManifestChange) XXX_Size() int {
	return m.Size()
}
func (m *ManifestChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestChange.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestChange proto.InternalMessageInfo

func (m *ManifestChange) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ManifestChange) GetOp() ManifestChange_Operation {
	if m != nil {
		return m.Op
	}
	return ManifestChange_CREATE
}

func (m *ManifestChange) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *ManifestChange) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.ManifestChange_Operation", ManifestChange_Operation_name, ManifestChange_Operation_value)
	proto.RegisterType((*KV)(nil), "pb.KV")
	proto.RegisterType((*KVList)(nil), "pb.KVList")
	proto.RegisterType((*ManifestChangeSet)(nil), "pb.ManifestChangeSet")
	proto.RegisterType((*ManifestChange)(nil), "pb.ManifestChange")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x8a, 0xdb, 0x30,
	0x14, 0xc6, 0x47, 0x8a, 0xc7, 0xe3, 0xbc, 0xce, 0x04, 0x57, 0x94, 0x22, 0xfa, 0xc7, 0x18, 0x77,
	0xe3, 0xc5, 0xe0, 0xc5, 0xf4, 0x04, 0x69, 0xea, 0x45, 0x48, 0x42, 0x40, 0x0d, 0xd9, 0x06, 0x39,
	0x7e, 0x6d, 0x8c, 0x13, 0x5b, 0x58, 0x8a, 0x69, 0x6f, 0xd2, 0x0b, 0xf4, 0x04, 0xbd, 0x44, 0x97,
	0x3d, 0x42, 0x49, 0x2f, 0x52, 0xac, 0xfc, 0x81, 0xd0, 0xdd, 0xfb, 0xbe, 0xef, 0xbd, 0x4f, 0xf0,
	0x13, 0x78, 0x2a, 0x4b, 0x54, 0x53, 0x9b, 0x9a, 0x51, 0x95, 0x45, 0x3f, 0x09, 0xd0, 0xc9, 0x92,
	0xf9, 0xd0, 0x2b, 0xf1, 0x1b, 0x27, 0x21, 0x89, 0xef, 0x45, 0x37, 0xb2, 0x17, 0x70, 0xdb, 0xca,
	0xed, 0x1e, 0x39, 0xb5, 0xde, 0x51, 0xb0, 0xd7, 0xd0, 0xdf, 0x6b, 0x6c, 0x56, 0x3b, 0x34, 0x92,
	0xf7, 0x6c, 0xe2, 0x75, 0xc6, 0x0c, 0x8d, 0x64, 0x1c, 0xee, 0x5a, 0x6c, 0x74, 0x51, 0x57, 0xdc,
	0x09, 0x49, 0xec, 0x88, 0xb3, 0x64, 0x6f, 0x01, 0xf0, 0xab, 0x2a, 0x1a, 0xd4, 0x2b, 0x69, 0xf8,
	0xad, 0x0d, 0xfb, 0x27, 0x67, 0x68, 0x18, 0x03, 0xc7, 0x16, 0xba, 0xb6, 0xd0, 0xce, 0xdd, 0x4b,
	0xda, 0x34, 0x28, 0x77, 0xab, 0x22, 0xe7, 0x10, 0x92, 0xf8, 0x41, 0x78, 0x47, 0x63, 0x9c, 0x47,
	0x21, 0xb8, 0x93, 0xe5, 0xb4, 0xd0, 0x86, 0xbd, 0x04, 0x5a, 0xb6, 0x9c, 0x84, 0xbd, 0xf8, 0xd9,
	0x93, 0x9b, 0xa8, 0x2c, 0x99, 0x2c, 0x05, 0x2d, 0xdb, 0x68, 0x08, 0xcf, 0x67, 0xb2, 0x2a, 0x3e,
	0xa3, 0x36, 0xa3, 0x8d, 0xac, 0xbe, 0xe0, 0x27, 0x34, 0xec, 0x11, 0xee, 0xd6, 0x56, 0xe8, 0xd3,
	0x05, 0xeb, 0x2e, 0xae, 0xf7, 0xc4, 0x79, 0x25, 0xfa, 0x41, 0x60, 0x70, 0x9d, 0xb1, 0x01, 0xd0,
	0x71, 0x6e, 0x29, 0x39, 0x82, 0x8e, 0x73, 0xf6, 0x08, 0x74, 0xae, 0x2c, 0xa1, 0xc1, 0xd3, 0x9b,
	0xff, 0xbb, 0x92, 0xb9, 0xc2, 0x46, 0x9a, 0xa2, 0xae, 0x04, 0x9d, 0xab, 0x0e, 0xe9, 0x14, 0x5b,
	0xdc, 0x5a, 0x70, 0x0f, 0xe2, 0x28, 0xd8, 0x2b, 0xf0, 0x46, 0x1b, 0x5c, 0x97, 0x7a, 0xbf, 0xb3,
	0xd8, 0xee, 0xc5, 0x45, 0x47, 0xef, 0xa0, 0x7f, 0xa9, 0x60, 0x00, 0xee, 0x48, 0xa4, 0xc3, 0x45,
	0xea, 0xdf, 0x74, 0xf3, 0xc7, 0x74, 0x9a, 0x2e, 0x52, 0x9f, 0x7c, 0xf0, 0x7f, 0x1d, 0x02, 0xf2,
	0xfb, 0x10, 0x90, 0x3f, 0x87, 0x80, 0x7c, 0xff, 0x1b, 0xdc, 0x64, 0xae, 0xfd, 0xdf, 0xf7, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x28, 0x5d, 0xcf, 0xeb, 0x01, 0x00, 0x00,
}

func (m *KV) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KV) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KV) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.StreamId != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.StreamId))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x32
	}
	if m.ExpiresAt != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x28
	}
	if m.Version != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x20
	}
	if len(m.UserMeta) > 0 {
		i -= len(m.UserMeta)
		copy(dAtA[i:], m.UserMeta)
		i = encodeVarintPb(dAtA, i, uint64(len(m.UserMeta)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KVList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KVList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Kv) > 0 {
		for iNdEx := len(m.Kv) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Kv[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPb(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ManifestChangeSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChangeSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestChangeSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Changes) > 0 {
		for iNdEx := len(m.Changes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPb(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ManifestChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManifestChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x22
	}
	if m.Level != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if m.Op != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	offset -= sovPb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KV) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.UserMeta)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovPb(uint64(m.Version))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovPb(uint64(m.ExpiresAt))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.StreamId != 0 {
		n += 1 + sovPb(uint64(m.StreamId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *KVList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Kv) > 0 {
		for _, e := range m.Kv {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestChangeSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPb(uint64(m.Id))
	}
	if m.Op != 0 {
		n += 1 + sovPb(uint64(m.Op))
	}
	if m.Level != 0 {
		n += 1 + sovPb(uint64(m.Level))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KV) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: KV: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KV: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserMeta", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserMeta = append(m.UserMeta[:0], dAtA[iNdEx:postIndex]...)
			if m.UserMeta == nil {
				m.UserMeta = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = append(m.Meta[:0], dAtA[iNdEx:postIndex]...)
			if m.Meta == nil {
				m.Meta = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamId", wireType)
			}
			m.StreamId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *KVList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: KVList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kv = append(m.Kv, &KV{})
			if err := m.Kv[len(m.Kv)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestChangeSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: ManifestChangeSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChangeSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, &ManifestChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: ManifestChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= ManifestChange_Operation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
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
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
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
				return 0, ErrInvalidLengthPb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPb = fmt.Errorf("proto: unexpected end of group")
)
