// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package go_micro_srv_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	Response
*/
package go_micro_srv_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// define messafe types
type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created     bool         `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.consignment.ShippingService/CreateConsignment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0xfc, 0xd2, 0xff, 0x6c, 0xa4, 0x0f, 0xe1, 0x03, 0x58, 0x70, 0x20, 0x4a, 0x85, 0xd4, 0x53,
	0x90, 0xca, 0x23, 0xe4, 0x80, 0x7a, 0x75, 0x1f, 0x00, 0x8a, 0xbd, 0x4a, 0x57, 0x22, 0x76, 0x64,
	0xbb, 0xe5, 0xd9, 0x78, 0x3b, 0x14, 0xa7, 0x01, 0x4b, 0xa8, 0x88, 0xdb, 0xce, 0xce, 0x4e, 0x66,
	0x46, 0x31, 0x2c, 0x5b, 0x6b, 0xbc, 0x79, 0x90, 0x46, 0x3b, 0xaa, 0x75, 0x83, 0xda, 0xc7, 0x73,
	0x19, 0x58, 0xc6, 0x6b, 0x53, 0x36, 0x24, 0xad, 0x29, 0x9d, 0x3d, 0x96, 0x11, 0x5f, 0x7c, 0x24,
	0x90, 0x55, 0xdf, 0x98, 0xfd, 0x87, 0x11, 0x29, 0x9e, 0xe4, 0xc9, 0x2a, 0x15, 0x23, 0x52, 0x2c,
	0x87, 0x4c, 0xa1, 0x93, 0x96, 0x5a, 0x4f, 0x46, 0xf3, 0x51, 0x20, 0xe2, 0x15, 0xbb, 0x82, 0xd9,
	0x3b, 0x52, 0xbd, 0xf7, 0x7c, 0x9c, 0x27, 0xab, 0xa9, 0x38, 0x21, 0x56, 0x01, 0x48, 0xa3, 0xfd,
	0x8e, 0x34, 0x5a, 0xc7, 0x27, 0xf9, 0x78, 0x95, 0xad, 0x97, 0xe5, 0xb9, 0x20, 0x65, 0x35, 0xdc,
	0x8a, 0x48, 0xc6, 0x6e, 0x21, 0x3d, 0xa2, 0x73, 0xf8, 0xf6, 0x4c, 0x8a, 0x4f, 0x83, 0xf9, 0xa2,
	0x5f, 0x6c, 0x54, 0xd1, 0x40, 0xfa, 0xa5, 0xfa, 0x11, 0xfc, 0x0e, 0x32, 0x79, 0x70, 0xde, 0x34,
	0x68, 0x3b, 0x6d, 0x1f, 0x1c, 0x86, 0xd5, 0x46, 0x75, 0xb9, 0x8d, 0xa5, 0x9a, 0x74, 0xc8, 0x9d,
	0x8a, 0x13, 0x62, 0xd7, 0x30, 0x3f, 0xb8, 0x5e, 0x34, 0xe9, 0x89, 0x0e, 0x06, 0xbb, 0x85, 0x40,
	0xd7, 0x1a, 0xed, 0x90, 0x71, 0x98, 0x4b, 0x8b, 0x3b, 0x8f, 0xbd, 0xe5, 0x42, 0x0c, 0x90, 0x3d,
	0x41, 0x16, 0xd5, 0x0a, 0xbe, 0xd9, 0xfa, 0xfe, 0xd7, 0xde, 0xc3, 0x2c, 0x62, 0xe5, 0xda, 0xc1,
	0xc5, 0x76, 0x4f, 0x6d, 0x4b, 0xba, 0xde, 0xa2, 0x3d, 0x92, 0x44, 0xf6, 0x02, 0x97, 0x55, 0xb0,
	0x89, 0xff, 0xd8, 0xdf, 0xbe, 0x7d, 0x53, 0x9c, 0x3f, 0x1b, 0x5a, 0x15, 0xff, 0x5e, 0x67, 0xe1,
	0xbd, 0x3c, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe8, 0xcf, 0xde, 0x56, 0x02, 0x00, 0x00,
}
