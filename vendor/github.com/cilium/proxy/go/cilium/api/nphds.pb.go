// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/api/nphds.proto

package cilium

import (
	context "context"
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The mapping of a network policy identifier to the IP addresses of all the
// hosts on which the network policy is enforced.
// A host may be associated only with one network policy.
type NetworkPolicyHosts struct {
	// The unique identifier of the network policy enforced on the hosts.
	Policy uint64 `protobuf:"varint,1,opt,name=policy,proto3" json:"policy,omitempty"`
	// The set of IP addresses of the hosts on which the network policy is enforced.
	// Optional. May be empty.
	HostAddresses        []string `protobuf:"bytes,2,rep,name=host_addresses,json=hostAddresses,proto3" json:"host_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkPolicyHosts) Reset()         { *m = NetworkPolicyHosts{} }
func (m *NetworkPolicyHosts) String() string { return proto.CompactTextString(m) }
func (*NetworkPolicyHosts) ProtoMessage()    {}
func (*NetworkPolicyHosts) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09d37d0b8c67603, []int{0}
}

func (m *NetworkPolicyHosts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPolicyHosts.Unmarshal(m, b)
}
func (m *NetworkPolicyHosts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPolicyHosts.Marshal(b, m, deterministic)
}
func (m *NetworkPolicyHosts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPolicyHosts.Merge(m, src)
}
func (m *NetworkPolicyHosts) XXX_Size() int {
	return xxx_messageInfo_NetworkPolicyHosts.Size(m)
}
func (m *NetworkPolicyHosts) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPolicyHosts.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPolicyHosts proto.InternalMessageInfo

func (m *NetworkPolicyHosts) GetPolicy() uint64 {
	if m != nil {
		return m.Policy
	}
	return 0
}

func (m *NetworkPolicyHosts) GetHostAddresses() []string {
	if m != nil {
		return m.HostAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkPolicyHosts)(nil), "cilium.NetworkPolicyHosts")
}

func init() { proto.RegisterFile("cilium/api/nphds.proto", fileDescriptor_b09d37d0b8c67603) }

var fileDescriptor_b09d37d0b8c67603 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0xff, 0xcc, 0x5f, 0x07, 0x1a, 0xd0, 0xc5, 0x20, 0xed, 0x30, 0x14, 0x2d, 0xb3, 0xb1,
	0x08, 0xce, 0x48, 0xdd, 0xd5, 0x95, 0x45, 0xc4, 0x95, 0x48, 0xfb, 0x00, 0x25, 0xce, 0x5c, 0xda,
	0xe0, 0x34, 0x37, 0xe6, 0xa6, 0x91, 0x6e, 0xdd, 0xb9, 0x6d, 0x1f, 0xcd, 0x57, 0xf0, 0x29, 0xba,
	0x92, 0xce, 0x74, 0x04, 0x51, 0x97, 0xee, 0x72, 0xf3, 0x9d, 0xe4, 0x24, 0xe7, 0xf0, 0x56, 0x26,
	0x0b, 0xb9, 0x98, 0xa7, 0x42, 0xcb, 0x54, 0xe9, 0x59, 0x4e, 0x89, 0x36, 0x68, 0x31, 0xf0, 0xab,
	0xfd, 0xa8, 0x03, 0xca, 0xe1, 0xb2, 0xc4, 0xae, 0x9f, 0xe6, 0x92, 0x32, 0x74, 0x60, 0x96, 0x95,
	0x2a, 0xea, 0x4c, 0x11, 0xa7, 0x05, 0x94, 0x58, 0x28, 0x85, 0x56, 0x58, 0x89, 0x6a, 0x77, 0x47,
	0xd4, 0x76, 0xa2, 0x90, 0xb9, 0xb0, 0x90, 0xd6, 0x8b, 0x0a, 0xc4, 0x92, 0x07, 0x77, 0x60, 0x9f,
	0xd1, 0x3c, 0xde, 0x63, 0x21, 0xb3, 0xe5, 0x2d, 0x92, 0xa5, 0xa0, 0xc5, 0x7d, 0x5d, 0x8e, 0x21,
	0xeb, 0xb2, 0x5e, 0x63, 0xb4, 0x9b, 0x82, 0x4b, 0x7e, 0x30, 0x43, 0xb2, 0x13, 0x91, 0xe7, 0x06,
	0x88, 0x80, 0x42, 0xaf, 0xfb, 0xbf, 0xd7, 0x1c, 0x1e, 0x6e, 0x86, 0x7b, 0x2b, 0xe6, 0x85, 0x6c,
	0x33, 0x6c, 0xae, 0x98, 0x1f, 0x37, 0x8c, 0xd7, 0x65, 0xa3, 0xfd, 0xad, 0xf6, 0xaa, 0x96, 0xf6,
	0xd7, 0x1e, 0x8f, 0xbf, 0x7b, 0x5d, 0xd7, 0xff, 0x18, 0x83, 0x71, 0x32, 0x83, 0x40, 0xf0, 0x70,
	0x6c, 0x0d, 0x88, 0xf9, 0x0f, 0xef, 0x3a, 0x4a, 0xca, 0x0c, 0x12, 0xa1, 0x65, 0xe2, 0xfa, 0xc9,
	0xe7, 0xd9, 0x11, 0x3c, 0x2d, 0x80, 0x6c, 0x74, 0xfc, 0x2b, 0x27, 0x8d, 0x8a, 0x20, 0xfe, 0xd7,
	0x63, 0xe7, 0x2c, 0x78, 0x65, 0xbc, 0x7d, 0x03, 0x36, 0x9b, 0xfd, 0x85, 0xc5, 0xd9, 0xcb, 0xdb,
	0xfb, 0xda, 0x3b, 0x89, 0xe3, 0x2f, 0x2d, 0x0d, 0x54, 0x65, 0x35, 0xa9, 0x92, 0x9c, 0x6c, 0xa3,
	0xa1, 0x01, 0x3b, 0x7d, 0xf0, 0xcb, 0x1e, 0x2e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x02, 0x2e,
	0xc1, 0xfc, 0xfe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkPolicyHostsDiscoveryServiceClient is the client API for NetworkPolicyHostsDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkPolicyHostsDiscoveryServiceClient interface {
	StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error)
	FetchNetworkPolicyHosts(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type networkPolicyHostsDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkPolicyHostsDiscoveryServiceClient(cc *grpc.ClientConn) NetworkPolicyHostsDiscoveryServiceClient {
	return &networkPolicyHostsDiscoveryServiceClient{cc}
}

func (c *networkPolicyHostsDiscoveryServiceClient) StreamNetworkPolicyHosts(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkPolicyHostsDiscoveryService_serviceDesc.Streams[0], "/cilium.NetworkPolicyHostsDiscoveryService/StreamNetworkPolicyHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient{stream}
	return x, nil
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient struct {
	grpc.ClientStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyHostsDiscoveryServiceClient) FetchNetworkPolicyHosts(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkPolicyHostsDiscoveryServiceServer is the server API for NetworkPolicyHostsDiscoveryService service.
type NetworkPolicyHostsDiscoveryServiceServer interface {
	StreamNetworkPolicyHosts(NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error
	FetchNetworkPolicyHosts(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedNetworkPolicyHostsDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkPolicyHostsDiscoveryServiceServer struct {
}

func (*UnimplementedNetworkPolicyHostsDiscoveryServiceServer) StreamNetworkPolicyHosts(srv NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNetworkPolicyHosts not implemented")
}
func (*UnimplementedNetworkPolicyHostsDiscoveryServiceServer) FetchNetworkPolicyHosts(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchNetworkPolicyHosts not implemented")
}

func RegisterNetworkPolicyHostsDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyHostsDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyHostsDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyHostsDiscoveryServiceServer).StreamNetworkPolicyHosts(&networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer{stream})
}

type NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHostsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer struct {
	grpc.ServerStream
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyHostsDiscoveryServiceStreamNetworkPolicyHostsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cilium.NetworkPolicyHostsDiscoveryService/FetchNetworkPolicyHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyHostsDiscoveryServiceServer).FetchNetworkPolicyHosts(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkPolicyHostsDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cilium.NetworkPolicyHostsDiscoveryService",
	HandlerType: (*NetworkPolicyHostsDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchNetworkPolicyHosts",
			Handler:    _NetworkPolicyHostsDiscoveryService_FetchNetworkPolicyHosts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNetworkPolicyHosts",
			Handler:       _NetworkPolicyHostsDiscoveryService_StreamNetworkPolicyHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cilium/api/nphds.proto",
}
