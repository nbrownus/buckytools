// Code generated by protoc-gen-go.
// source: timeseries.proto
// DO NOT EDIT!

/*
Package metrics is a generated protocol buffer package.

It is generated from these files:
	timeseries.proto

It has these top-level messages:
	TimeSeries
*/
package metrics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type TimeSeries struct {
	Epoch    int64     `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	Interval int64     `protobuf:"varint,2,opt,name=interval" json:"interval,omitempty"`
	Values   []float64 `protobuf:"fixed64,3,rep,name=values" json:"values,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*TimeSeries)(nil), "metrics.TimeSeries")
}

var fileDescriptor0 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc9, 0xcc, 0x4d,
	0x2d, 0x4e, 0x2d, 0xca, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d,
	0x2d, 0x29, 0xca, 0x4c, 0x2e, 0x56, 0xb2, 0xe5, 0xe2, 0x0a, 0x01, 0x4a, 0x06, 0x83, 0x25, 0x85,
	0x78, 0xb9, 0x58, 0x53, 0x0b, 0xf2, 0x93, 0x33, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x85, 0x04,
	0xb8, 0x38, 0x32, 0xf3, 0x4a, 0x52, 0x8b, 0xca, 0x12, 0x73, 0x24, 0x98, 0xc0, 0x22, 0x7c, 0x5c,
	0x6c, 0x40, 0x4e, 0x69, 0x6a, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0x63, 0x12, 0x1b, 0xd8, 0x38,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xc3, 0x30, 0xc0, 0x62, 0x00, 0x00, 0x00,
}
