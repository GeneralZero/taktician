// Code generated by protoc-gen-go. DO NOT EDIT.
// source: position.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Position struct {
	Day      string  `protobuf:"bytes,1,opt,name=day" json:"day,omitempty"`
	Id       int32   `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Ply      int32   `protobuf:"varint,3,opt,name=ply" json:"ply,omitempty"`
	Tps      string  `protobuf:"bytes,4,opt,name=tps" json:"tps,omitempty"`
	Move     string  `protobuf:"bytes,5,opt,name=move" json:"move,omitempty"`
	Value    float32 `protobuf:"fixed32,6,opt,name=value" json:"value,omitempty"`
	Plies    int32   `protobuf:"varint,7,opt,name=plies" json:"plies,omitempty"`
	Features []int64 `protobuf:"varint,8,rep,packed,name=features" json:"features,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Position) GetDay() string {
	if m != nil {
		return m.Day
	}
	return ""
}

func (m *Position) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Position) GetPly() int32 {
	if m != nil {
		return m.Ply
	}
	return 0
}

func (m *Position) GetTps() string {
	if m != nil {
		return m.Tps
	}
	return ""
}

func (m *Position) GetMove() string {
	if m != nil {
		return m.Move
	}
	return ""
}

func (m *Position) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Position) GetPlies() int32 {
	if m != nil {
		return m.Plies
	}
	return 0
}

func (m *Position) GetFeatures() []int64 {
	if m != nil {
		return m.Features
	}
	return nil
}

func init() {
	proto.RegisterType((*Position)(nil), "tak.proto.Position")
}

func init() { proto.RegisterFile("position.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0x0e, 0xc2, 0x20,
	0x10, 0x40, 0x03, 0xb4, 0xb5, 0x65, 0x68, 0x0c, 0x71, 0xb8, 0x38, 0x11, 0x27, 0x26, 0x17, 0xff,
	0xc0, 0x2f, 0x30, 0x8c, 0x6e, 0x34, 0xc5, 0x84, 0x58, 0xe5, 0x52, 0x68, 0x93, 0xfe, 0x93, 0x1f,
	0x69, 0x00, 0xe3, 0xf6, 0xde, 0x83, 0xbb, 0x1c, 0xef, 0xd1, 0x07, 0x17, 0x9d, 0x7f, 0x9f, 0x71,
	0xf6, 0xd1, 0x8b, 0x2e, 0x9a, 0x67, 0xc1, 0xd3, 0x87, 0xf0, 0xf6, 0xf6, 0x7b, 0x15, 0x7b, 0xce,
	0x46, 0xb3, 0x01, 0x91, 0x44, 0x75, 0x3a, 0xa1, 0xe8, 0x39, 0x75, 0x23, 0x50, 0x49, 0x54, 0xad,
	0xa9, 0x1b, 0xd3, 0x0f, 0x9c, 0x36, 0x60, 0x39, 0x24, 0x4c, 0x25, 0x62, 0x80, 0xaa, 0xcc, 0x44,
	0x0c, 0x42, 0xf0, 0xea, 0xe5, 0x57, 0x0b, 0x75, 0x4e, 0x99, 0xc5, 0x81, 0xd7, 0xab, 0x99, 0x16,
	0x0b, 0x8d, 0x24, 0x8a, 0xea, 0x22, 0xa9, 0xe2, 0xe4, 0x6c, 0x80, 0x5d, 0xde, 0x57, 0x44, 0x1c,
	0x79, 0xfb, 0xb0, 0x26, 0x2e, 0xb3, 0x0d, 0xd0, 0x4a, 0xa6, 0x98, 0xfe, 0xfb, 0xb5, 0xba, 0x53,
	0x1c, 0x86, 0x26, 0xdf, 0x7e, 0xf9, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xed, 0x21, 0x82, 0xd8,
	0x00, 0x00, 0x00,
}
