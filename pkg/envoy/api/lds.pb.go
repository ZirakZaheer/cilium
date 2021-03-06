// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/lds.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf4 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Listener_DrainType int32

const (
	// Drain in response to calling /healthcheck/fail admin endpoint (along with the health check
	// filter), listener removal/modification, and hot restart.
	Listener_DEFAULT Listener_DrainType = 0
	// Drain in response to listener removal/modification and hot restart. This setting does not
	// include /healthcheck/fail. This setting may be desirable if Envoy is hosting both ingress
	// and egress listeners.
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}
var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}
func (Listener_DrainType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type Filter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	Config       *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	DeprecatedV1 *Filter_DeprecatedV1     `protobuf:"bytes,3,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetConfig() *google_protobuf4.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Filter) GetDeprecatedV1() *Filter_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

type Filter_DeprecatedV1 struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Filter_DeprecatedV1) Reset()                    { *m = Filter_DeprecatedV1{} }
func (m *Filter_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*Filter_DeprecatedV1) ProtoMessage()               {}
func (*Filter_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Filter_DeprecatedV1) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// Specifies the match criteria for selecting a specific filter chain for a
// listener [V2-API-DIFF].
type FilterChainMatch struct {
	// If non-empty, the SNI domains to consider. May contain a wildcard prefix,
	// e.g. ``*.example.com``.
	SniDomains []string `protobuf:"bytes,1,rep,name=sni_domains,json=sniDomains" json:"sni_domains,omitempty"`
	// If non-empty, an IP address and prefix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	PrefixRanges []*CidrRange `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges" json:"prefix_ranges,omitempty"`
	// If non-empty, an IP address and suffix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	AddressSuffix string                       `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix" json:"address_suffix,omitempty"`
	SuffixLen     *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen" json:"suffix_len,omitempty"`
	// The criteria is satisfied if the source IP address of the downstream
	// connection is contained in at least one of the specified subnets. If the
	// parameter is not specified or the list is empty, the source IP address is
	// ignored.
	SourcePrefixRanges []*CidrRange `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges" json:"source_prefix_ranges,omitempty"`
	// The criteria is satisfied if the source port of the downstream connection
	// is contained in at least one of the specified ports. If the parameter is
	// not specified, the source port is ignored.
	SourcePorts []*google_protobuf.UInt32Value `protobuf:"bytes,7,rep,name=source_ports,json=sourcePorts" json:"source_ports,omitempty"`
	// Optional destination port to consider when use_original_dst is set on the
	// listener in determining a filter chain match.
	DestinationPort *google_protobuf.UInt32Value `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort" json:"destination_port,omitempty"`
}

func (m *FilterChainMatch) Reset()                    { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string            { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()               {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FilterChainMatch) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *google_protobuf.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []*google_protobuf.UInt32Value {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetDestinationPort() *google_protobuf.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

// Grouping of FilterChainMatch criteria, DownstreamTlsContext, the actual filter chain
// and related parameters.
type FilterChain struct {
	FilterChainMatch *FilterChainMatch     `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch" json:"filter_chain_match,omitempty"`
	TlsContext       *DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext" json:"tls_context,omitempty"`
	// A list of individual network filters that make up the filter chain for
	// connections established with the listener. Order matters as the filters are
	// processed sequentially as connection events happen.  Note: If the filter
	// list is empty, the connection will close by default.
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters" json:"filters,omitempty"`
	// Whether the listener should expect a PROXY protocol V1 header on new
	// connections. If this option is enabled, the listener will assume that that
	// remote address of the connection is the one specified in the header. Some
	// load balancers including the AWS ELB support this option. If the option is
	// absent or set to false, Envoy will use the physical peer address of the
	// connection as the remote address.
	UseProxyProto *google_protobuf.BoolValue `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto" json:"use_proxy_proto,omitempty"`
	// See base.Metadata description.
	Metadata *Metadata `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	// See base.TransportSocket description.
	TransportSocket *TransportSocket `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket" json:"transport_socket,omitempty"`
}

func (m *FilterChain) Reset()                    { *m = FilterChain{} }
func (m *FilterChain) String() string            { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()               {}
func (*FilterChain) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetTlsContext() *DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

type Listener struct {
	// The unique name of the listener. If no name is provided, Envoy will generate a
	// UUID for internal use. The name is used for dynamic listener update and removal
	// via the LDS APIs.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The address that the listener should listen on.
	Address *Address `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// A list of filter chains to consider for this listener. The FilterChain with
	// the most specific FilterChainMatch criteria is used on a connection. The
	// algorithm works as follows:
	//
	// 1. If SNI information is presented at connection time, only the
	//    FilterChains matching the SNI are considered. Otherwise, only
	//    FilterChains with no SNI domains are considered.
	// 2. Of the FilterChains from step 1, the longest prefix match on the
	//    bound destination address is used to select the next set of
	//    FilterChains. This may be one FilterChain or multiple if there is
	//    a tie.
	// 3. The longest suffix match on the bound destination address is used to
	//    select the FilterChain from step 2 that is used.
	FilterChains []*FilterChain `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains" json:"filter_chains,omitempty"`
	// If a connection is redirected using iptables, the port on which the proxy
	// receives it might be different from the original destination port. When
	// this flag is set to true, the listener uses the original destination
	// address and port during FilterChain matching. Default is false.
	UseOriginalDst *google_protobuf.BoolValue `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst" json:"use_original_dst,omitempty"`
	// Soft limit on size of the listener’s new connection read and write buffers.
	// If unspecified, an implementation defined default is applied (1MiB).
	PerConnectionBufferLimitBytes *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes" json:"per_connection_buffer_limit_bytes,omitempty"`
	// See base.Metadata description.
	Metadata     *Metadata              `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	DeprecatedV1 *Listener_DeprecatedV1 `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
	// The type of draining to perform at a listener-wide level.
	DrainType Listener_DrainType `protobuf:"varint,8,opt,name=drain_type,json=drainType,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	// A list of individual listener filters that make up the filter chain for
	// sockets accepted with the listener. These filters are run before
	// any in the 'filter_chains', and these filters have the opportunity
	// to manipulate and augment the connection metadata that is used in
	// connection filter chain matching.  Order matters as the filters are
	// processed sequentially right after a socket has been accepted by
	// the listener, and before a connection is created.
	ListenerFilterChain []*Filter `protobuf:"bytes,9,rep,name=listener_filter_chain,json=listenerFilterChain" json:"listener_filter_chain,omitempty"`
}

func (m *Listener) Reset()                    { *m = Listener{} }
func (m *Listener) String() string            { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()               {}
func (*Listener) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

func (m *Listener) GetUseOriginalDst() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *google_protobuf.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilterChain() []*Filter {
	if m != nil {
		return m.ListenerFilterChain
	}
	return nil
}

type Listener_DeprecatedV1 struct {
	// Whether the listener should bind to the port. A listener that doesn’t
	// bind can only receive connections redirected from other listeners that
	// set use_original_dst parameter to true. Default is true.
	//
	// [V2-API-DIFF] This is deprecated in v2, all Listeners will bind to their
	// port. An additional filter chain must be created for every original
	// destination port this listener may redirect to in v2, with the original
	// port specified in the FilterChainMatch destination_port field.
	BindToPort *google_protobuf.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort" json:"bind_to_port,omitempty"`
}

func (m *Listener_DeprecatedV1) Reset()                    { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()               {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

func (m *Listener_DeprecatedV1) GetBindToPort() *google_protobuf.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "envoy.api.v2.Filter")
	proto.RegisterType((*Filter_DeprecatedV1)(nil), "envoy.api.v2.Filter.DeprecatedV1")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v2.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v2.FilterChain")
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ListenerDiscoveryService service

type ListenerDiscoveryServiceClient interface {
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ListenerDiscoveryService service

type ListenerDiscoveryServiceServer interface {
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/lds.proto",
}

func init() { proto.RegisterFile("api/lds.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 919 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x92, 0xdb, 0x34,
	0x14, 0xae, 0x9b, 0x36, 0xbb, 0x39, 0xf9, 0x45, 0x6d, 0xa9, 0xc9, 0x6c, 0xdb, 0x34, 0x0c, 0x33,
	0x0b, 0x17, 0x0e, 0x4d, 0xef, 0x4a, 0x87, 0x4e, 0x37, 0x21, 0x74, 0x67, 0xb2, 0xec, 0x8e, 0x93,
	0x76, 0xe8, 0x95, 0x47, 0xb1, 0xe5, 0xac, 0x06, 0x47, 0x32, 0x92, 0x9c, 0x6e, 0x6e, 0x79, 0x04,
	0x78, 0x0b, 0xde, 0x80, 0x2b, 0x1e, 0x82, 0x57, 0xe0, 0x8e, 0x97, 0x60, 0x24, 0xdb, 0xa9, 0x9d,
	0x76, 0x97, 0xbd, 0xe0, 0x4e, 0xd2, 0xf9, 0xbe, 0x4f, 0xc7, 0xe7, 0x3b, 0x3a, 0x09, 0x34, 0x71,
	0x4c, 0x07, 0x51, 0x20, 0x9d, 0x58, 0x70, 0xc5, 0x51, 0x83, 0xb0, 0x35, 0xdf, 0x38, 0x38, 0xa6,
	0xce, 0x7a, 0xd8, 0xfd, 0x44, 0x07, 0x71, 0x10, 0x08, 0x22, 0x33, 0x40, 0xb7, 0xa5, 0x8f, 0x16,
	0x58, 0x92, 0x6c, 0x7f, 0x47, 0xef, 0x03, 0x2a, 0x7d, 0xbe, 0x26, 0x62, 0x93, 0x1d, 0x1a, 0x51,
	0x99, 0x8b, 0x76, 0x0f, 0x96, 0x9c, 0x2f, 0x23, 0x32, 0x30, 0x6a, 0x8c, 0x71, 0x85, 0x15, 0xe5,
	0x6c, 0x37, 0x6a, 0x76, 0x8b, 0x24, 0x1c, 0x48, 0x25, 0x12, 0x5f, 0x65, 0xd1, 0x87, 0xbb, 0xd1,
	0x77, 0x02, 0xc7, 0x31, 0x11, 0x19, 0xbb, 0xff, 0x87, 0x05, 0xd5, 0x09, 0x8d, 0x14, 0x11, 0x08,
	0xc1, 0x2d, 0x86, 0x57, 0xc4, 0xb6, 0x7a, 0xd6, 0x61, 0xcd, 0x35, 0x6b, 0x34, 0x80, 0xaa, 0xcf,
	0x59, 0x48, 0x97, 0xf6, 0xcd, 0x9e, 0x75, 0x58, 0x1f, 0xde, 0x77, 0x52, 0x3d, 0x27, 0xd7, 0x73,
	0x66, 0xe6, 0x36, 0x37, 0x83, 0xa1, 0x09, 0x34, 0x03, 0x12, 0x0b, 0xe2, 0x63, 0x45, 0x02, 0x6f,
	0xfd, 0xc4, 0xae, 0x18, 0xde, 0x63, 0xa7, 0x58, 0x18, 0x27, 0xbd, 0xd1, 0x19, 0x6f, 0x91, 0x6f,
	0x9e, 0xb8, 0x8d, 0xa0, 0xb0, 0xeb, 0xf6, 0xa1, 0x51, 0x8c, 0xea, 0xe4, 0xd4, 0x26, 0xde, 0x26,
	0xa7, 0xd7, 0xfd, 0xdf, 0x2b, 0xd0, 0x49, 0x95, 0x46, 0xe7, 0x98, 0xb2, 0x13, 0xac, 0xfc, 0x73,
	0xf4, 0x08, 0xea, 0x92, 0x51, 0x2f, 0xe0, 0x2b, 0x4c, 0x99, 0xb4, 0xad, 0x5e, 0xe5, 0xb0, 0xe6,
	0x82, 0x64, 0x74, 0x9c, 0x9e, 0xa0, 0xe7, 0xd0, 0x8c, 0x05, 0x09, 0xe9, 0x85, 0x27, 0x30, 0x5b,
	0x12, 0x69, 0x57, 0x7a, 0x15, 0xf3, 0x65, 0xa5, 0x0c, 0x47, 0x34, 0x10, 0xae, 0x8e, 0xbb, 0x8d,
	0x14, 0x6d, 0x36, 0x12, 0x7d, 0x01, 0xad, 0xcc, 0x50, 0x4f, 0x26, 0x61, 0x48, 0x2f, 0xec, 0x5b,
	0x26, 0xa3, 0x66, 0x76, 0x3a, 0x33, 0x87, 0xe8, 0x1b, 0x80, 0x34, 0xec, 0x45, 0x84, 0xd9, 0xb7,
	0x4d, 0x0d, 0x0e, 0x3e, 0xa8, 0xdd, 0xeb, 0x63, 0xa6, 0x9e, 0x0e, 0xdf, 0xe0, 0x28, 0x21, 0x6e,
	0x2d, 0xc5, 0x4f, 0x09, 0x43, 0xc7, 0x70, 0x57, 0xf2, 0x44, 0xf8, 0xc4, 0x2b, 0x27, 0x5a, 0xbd,
	0x3a, 0x51, 0x94, 0x92, 0xce, 0x8a, 0xe9, 0xbe, 0x80, 0x46, 0x2e, 0xc5, 0x85, 0x92, 0xf6, 0x9e,
	0x91, 0xb8, 0x3a, 0x93, 0x7a, 0xa6, 0xa3, 0x09, 0xe8, 0x7b, 0xe8, 0x04, 0x44, 0x2a, 0xca, 0x4c,
	0xcf, 0x19, 0x15, 0x7b, 0xff, 0x1a, 0x9f, 0xd3, 0x2e, 0xb0, 0xb4, 0x52, 0xff, 0xd7, 0x0a, 0xd4,
	0x0b, 0x66, 0xa1, 0x29, 0xa0, 0xd0, 0x6c, 0x3d, 0x5f, 0xef, 0xbd, 0x95, 0x76, 0xcf, 0xd8, 0x5b,
	0x1f, 0x3e, 0xfc, 0x58, 0xb7, 0xbc, 0xf7, 0xd8, 0xed, 0x84, 0xbb, 0xae, 0x8f, 0xa0, 0xae, 0x22,
	0xe9, 0xf9, 0x9c, 0x29, 0x72, 0xa1, 0xb2, 0x66, 0xed, 0x97, 0x65, 0xc6, 0xfc, 0x1d, 0x93, 0x4a,
	0x10, 0xbc, 0x9a, 0x47, 0x72, 0x94, 0x22, 0x5d, 0x50, 0xdb, 0x35, 0x72, 0x60, 0x2f, 0x15, 0xce,
	0x7b, 0xe2, 0xee, 0xc7, 0xf2, 0x70, 0x73, 0x10, 0x3a, 0x82, 0x76, 0x22, 0xb5, 0x49, 0xfc, 0x62,
	0xe3, 0x99, 0x2a, 0x98, 0x66, 0xa8, 0x0f, 0xbb, 0x1f, 0x94, 0xe6, 0x88, 0xf3, 0x28, 0x2d, 0x4c,
	0x33, 0x91, 0xe4, 0x4c, 0x33, 0xce, 0xcc, 0xc0, 0x18, 0xc2, 0xfe, 0x8a, 0x28, 0x1c, 0x60, 0x85,
	0xb3, 0x36, 0xf9, 0xb4, 0x7c, 0xe9, 0x49, 0x16, 0x75, 0xb7, 0x38, 0xf4, 0x0a, 0x3a, 0x4a, 0x60,
	0x26, 0xb5, 0x19, 0x9e, 0xe4, 0xfe, 0x4f, 0x44, 0xd9, 0x55, 0xc3, 0x7d, 0x50, 0xe6, 0xce, 0x73,
	0xd4, 0xcc, 0x80, 0xdc, 0xb6, 0x2a, 0x1f, 0xf4, 0xff, 0xbc, 0x0d, 0xfb, 0x53, 0x2a, 0x15, 0x61,
	0x97, 0xbe, 0xff, 0xbd, 0xac, 0xb1, 0xb3, 0x9a, 0xde, 0x2b, 0xdf, 0xf0, 0x32, 0x0d, 0xba, 0x39,
	0x0a, 0x7d, 0x0b, 0xcd, 0xa2, 0xad, 0x79, 0x25, 0x3f, 0xbb, 0xd4, 0x51, 0xb7, 0x51, 0x30, 0x53,
	0xa2, 0x31, 0x74, 0x74, 0x4d, 0xb9, 0xa0, 0x4b, 0xca, 0x70, 0xe4, 0x05, 0x52, 0x5d, 0xa3, 0xa8,
	0xad, 0x44, 0x92, 0xd3, 0x8c, 0x32, 0x96, 0x0a, 0x85, 0xf0, 0x38, 0xd6, 0x29, 0x70, 0xc6, 0x88,
	0x6f, 0x1a, 0x77, 0x91, 0x84, 0x21, 0x11, 0x5e, 0x44, 0x57, 0x54, 0x79, 0x8b, 0x8d, 0x22, 0xf2,
	0x5a, 0xaf, 0xf2, 0x41, 0x4c, 0xc4, 0x68, 0xab, 0x72, 0x64, 0x44, 0xa6, 0x5a, 0xe3, 0x48, 0x4b,
	0x94, 0xdc, 0xab, 0x5e, 0xdb, 0xbd, 0x9d, 0x09, 0xb9, 0x67, 0x88, 0x9f, 0x97, 0x89, 0xb9, 0x2b,
	0x57, 0xcc, 0x48, 0xf4, 0x02, 0x20, 0x10, 0xfa, 0xed, 0x98, 0xc9, 0xa8, 0x5f, 0x65, 0x6b, 0xd8,
	0xbb, 0x4c, 0x46, 0x03, 0xe7, 0x9b, 0x98, 0xb8, 0xb5, 0x20, 0x5f, 0xa2, 0x57, 0x70, 0x2f, 0xca,
	0x00, 0x5e, 0xd1, 0x35, 0xbb, 0x76, 0x45, 0xfb, 0xdf, 0xc9, 0x29, 0x05, 0x13, 0xbb, 0xd3, 0x9d,
	0x71, 0xfd, 0x1c, 0x1a, 0x0b, 0xca, 0x02, 0x4f, 0xf1, 0x74, 0x64, 0x58, 0xff, 0x69, 0x21, 0x68,
	0xfc, 0x9c, 0x9b, 0x59, 0xf1, 0x25, 0xd4, 0xb6, 0xf9, 0xa2, 0x3a, 0xec, 0x8d, 0xbf, 0x9b, 0xbc,
	0x7c, 0x3d, 0x9d, 0x77, 0x6e, 0xa0, 0x36, 0xd4, 0x4f, 0x4e, 0xc7, 0xc7, 0x93, 0xb7, 0xde, 0xe9,
	0x0f, 0xd3, 0xb7, 0x1d, 0x6b, 0xf8, 0x8f, 0x05, 0x76, 0xfe, 0x91, 0xe3, 0xfc, 0x67, 0x74, 0x46,
	0xc4, 0x9a, 0xfa, 0x04, 0xfd, 0x08, 0xed, 0x99, 0x79, 0xf0, 0x39, 0x42, 0xa2, 0x9d, 0xd1, 0xb2,
	0xa5, 0xb8, 0xe4, 0xe7, 0x84, 0x48, 0xd5, 0x7d, 0x74, 0x69, 0x5c, 0xc6, 0x9c, 0x49, 0xd2, 0xbf,
	0x71, 0x68, 0x7d, 0x6d, 0xa1, 0x04, 0x5a, 0x13, 0xa2, 0xfc, 0xf3, 0xff, 0x51, 0xb8, 0xff, 0xcb,
	0x5f, 0x7f, 0xff, 0x76, 0xf3, 0xa0, 0x7f, 0x7f, 0xb0, 0x1e, 0xbe, 0xff, 0x47, 0xf0, 0x2c, 0xaf,
	0xb6, 0x7c, 0x66, 0x7d, 0xb5, 0xa8, 0x9a, 0xc2, 0x3d, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0x5f, 0x5a, 0x02, 0x76, 0x08, 0x00, 0x00,
}
