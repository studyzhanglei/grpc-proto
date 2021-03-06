// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package exception

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type SearchError int32

const (
	//初始值, 无意义
	SearchError_INIT          SearchError = 0
	SearchError_SEARCH_FAILED SearchError = 10001
)

var SearchError_name = map[int32]string{
	0:     "INIT",
	10001: "SEARCH_FAILED",
}

var SearchError_value = map[string]int32{
	"INIT":          0,
	"SEARCH_FAILED": 10001,
}

func (x SearchError) String() string {
	return proto.EnumName(SearchError_name, int32(x))
}

func (SearchError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func init() {
	proto.RegisterEnum("exception.SearchError", SearchError_name, SearchError_value)
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xad, 0x48, 0x4e, 0x2d, 0x28,
	0xc9, 0xcc, 0xcf, 0xd3, 0xd2, 0xe6, 0xe2, 0x0e, 0x06, 0x4b, 0xb9, 0x16, 0x15, 0xe5, 0x17, 0x09,
	0x71, 0x70, 0xb1, 0x78, 0xfa, 0x79, 0x86, 0x08, 0x30, 0x08, 0x09, 0x71, 0xf1, 0x06, 0xbb, 0x3a,
	0x06, 0x39, 0x7b, 0xc4, 0xbb, 0x39, 0x7a, 0xfa, 0xb8, 0xba, 0x08, 0x4c, 0xf4, 0x4b, 0x62, 0x03,
	0x6b, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x22, 0xf7, 0xa4, 0x5c, 0x4e, 0x00, 0x00, 0x00,
}
