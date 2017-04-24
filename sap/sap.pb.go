// Code generated by protoc-gen-go.
// source: sap.proto
// DO NOT EDIT!

/*
Package sap is a generated protocol buffer package.

It is generated from these files:
	sap.proto

It has these top-level messages:
	Payload
*/
package sap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Payload struct {
	Model *google_protobuf.Struct `protobuf:"bytes,1,opt,name=Model,json=model" json:"Model,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Payload) GetModel() *google_protobuf.Struct {
	if m != nil {
		return m.Model
	}
	return nil
}

func init() {
	proto.RegisterType((*Payload)(nil), "sap.Payload")
}

func init() { proto.RegisterFile("sap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4e, 0x2c, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4e, 0x2c, 0x90, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0x17, 0x97, 0x14, 0x95, 0x26, 0x97, 0x40,
	0x94, 0x28, 0x59, 0x70, 0xb1, 0x07, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x08, 0xe9, 0x72, 0xb1,
	0xfa, 0xe6, 0xa7, 0xa4, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xeb, 0x41, 0x34,
	0xea, 0xc1, 0x34, 0xea, 0x05, 0x83, 0x35, 0x06, 0xb1, 0xe6, 0x82, 0x54, 0x25, 0xb1, 0x81, 0xc5,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xd5, 0x02, 0x85, 0x70, 0x00, 0x00, 0x00,
}
