// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/redis/v1/backup.proto

package redis // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Description of a Redis backup. For more information, see
// the Managed Service for Redis [documentation](/docs/managed-redis/concepts/backup).
type Backup struct {
	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format
	// (i.e. when the backup operation was completed).
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the Redis cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Start timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format
	// (i.e. when the backup operation was started).
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_backup_cd1ea4fa6820b24e, []int{0}
}
func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (dst *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(dst, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Backup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Backup) GetSourceClusterId() string {
	if m != nil {
		return m.SourceClusterId
	}
	return ""
}

func (m *Backup) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "yandex.cloud.mdb.redis.v1.Backup")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/redis/v1/backup.proto", fileDescriptor_backup_cd1ea4fa6820b24e)
}

var fileDescriptor_backup_cd1ea4fa6820b24e = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xb9, 0x53, 0x8b, 0x17, 0x41, 0x31, 0xd3, 0x59, 0x07, 0x8b, 0x83, 0x14, 0xa1, 0x09,
	0xd5, 0x49, 0x9c, 0xda, 0xe2, 0xe0, 0x5a, 0x9c, 0x5c, 0x8e, 0x24, 0x5f, 0x1a, 0x83, 0x97, 0xa6,
	0x24, 0x5f, 0x8a, 0xfe, 0x52, 0xff, 0x8e, 0x90, 0x5c, 0x47, 0x71, 0x0b, 0x6f, 0x9e, 0xef, 0x7d,
	0xe0, 0x25, 0x77, 0xdf, 0x62, 0x0b, 0xfa, 0x8b, 0xab, 0xde, 0x27, 0xe0, 0x0e, 0x24, 0x0f, 0x1a,
	0x6c, 0xe4, 0xfb, 0x39, 0x97, 0x42, 0x7d, 0xa6, 0x1d, 0xdb, 0x05, 0x8f, 0x9e, 0x5e, 0x15, 0x8e,
	0x65, 0x8e, 0x39, 0x90, 0x2c, 0x73, 0x6c, 0x3f, 0x1f, 0xdf, 0x18, 0xef, 0x4d, 0xaf, 0x79, 0x06,
	0x65, 0xda, 0x70, 0xb4, 0x4e, 0x47, 0x14, 0x6e, 0xb8, 0xbd, 0xfd, 0xa9, 0xc8, 0x68, 0x99, 0xcb,
	0xe8, 0x39, 0xa9, 0x2d, 0xb4, 0xd5, 0xa4, 0x9a, 0x36, 0xeb, 0xda, 0x02, 0xbd, 0x26, 0xcd, 0xc6,
	0xf7, 0xa0, 0x43, 0x67, 0xa1, 0xad, 0x73, 0x7c, 0x5a, 0x82, 0x57, 0xa0, 0x4f, 0x84, 0xa8, 0xa0,
	0x05, 0x6a, 0xe8, 0x04, 0xb6, 0x47, 0x93, 0x6a, 0x7a, 0xf6, 0x30, 0x66, 0xc5, 0xc6, 0x0e, 0x36,
	0xf6, 0x76, 0xb0, 0xad, 0x9b, 0x81, 0x5e, 0x20, 0xbd, 0x27, 0x97, 0xd1, 0xa7, 0xa0, 0x74, 0xa7,
	0xfa, 0x14, 0xb1, 0xf4, 0x1f, 0xe7, 0xfe, 0x8b, 0xf2, 0xb1, 0x2a, 0x79, 0xd1, 0x44, 0x14, 0x61,
	0xd0, 0x9c, 0xfc, 0xaf, 0x19, 0xe8, 0x05, 0x2e, 0x5f, 0xde, 0x57, 0xc6, 0xe2, 0x47, 0x92, 0x4c,
	0x79, 0xc7, 0xcb, 0x44, 0xb3, 0x32, 0xa5, 0xf1, 0x33, 0xa3, 0xb7, 0xf9, 0x9c, 0xff, 0xb9, 0xf1,
	0x73, 0x7e, 0xc8, 0x51, 0xc6, 0x1e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x17, 0x0e, 0x6a, 0xdc,
	0x8d, 0x01, 0x00, 0x00,
}
