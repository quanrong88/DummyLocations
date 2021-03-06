// Code generated by protoc-gen-go. DO NOT EDIT.
// source: locations.proto

/*
Package locations is a generated protocol buffer package.

It is generated from these files:
	locations.proto

It has these top-level messages:
	LocationsRequest
	Location
	LocationsResponse
	LocationDetailRequest
	LocationDetailResponse
*/
package locations

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type LocationsRequest struct {
}

func (m *LocationsRequest) Reset()                    { *m = LocationsRequest{} }
func (m *LocationsRequest) String() string            { return proto.CompactTextString(m) }
func (*LocationsRequest) ProtoMessage()               {}
func (*LocationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Location struct {
	Id       int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Subtitle string  `protobuf:"bytes,3,opt,name=subtitle" json:"subtitle,omitempty"`
	Lat      float32 `protobuf:"fixed32,4,opt,name=lat" json:"lat,omitempty"`
	Long     float32 `protobuf:"fixed32,5,opt,name=long" json:"long,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Location) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Location) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Location) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Location) GetLong() float32 {
	if m != nil {
		return m.Long
	}
	return 0
}

type LocationsResponse struct {
	Locations []*Location `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
}

func (m *LocationsResponse) Reset()                    { *m = LocationsResponse{} }
func (m *LocationsResponse) String() string            { return proto.CompactTextString(m) }
func (*LocationsResponse) ProtoMessage()               {}
func (*LocationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LocationsResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

type LocationDetailRequest struct {
	LocationId int32 `protobuf:"varint,1,opt,name=locationId" json:"locationId,omitempty"`
}

func (m *LocationDetailRequest) Reset()                    { *m = LocationDetailRequest{} }
func (m *LocationDetailRequest) String() string            { return proto.CompactTextString(m) }
func (*LocationDetailRequest) ProtoMessage()               {}
func (*LocationDetailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LocationDetailRequest) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

type LocationDetailResponse struct {
	Location *Location                         `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Comment  []*LocationDetailResponse_Comment `protobuf:"bytes,2,rep,name=comment" json:"comment,omitempty"`
}

func (m *LocationDetailResponse) Reset()                    { *m = LocationDetailResponse{} }
func (m *LocationDetailResponse) String() string            { return proto.CompactTextString(m) }
func (*LocationDetailResponse) ProtoMessage()               {}
func (*LocationDetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LocationDetailResponse) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *LocationDetailResponse) GetComment() []*LocationDetailResponse_Comment {
	if m != nil {
		return m.Comment
	}
	return nil
}

type LocationDetailResponse_Comment struct {
	Content  string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=imageUrl" json:"imageUrl,omitempty"`
}

func (m *LocationDetailResponse_Comment) Reset()         { *m = LocationDetailResponse_Comment{} }
func (m *LocationDetailResponse_Comment) String() string { return proto.CompactTextString(m) }
func (*LocationDetailResponse_Comment) ProtoMessage()    {}
func (*LocationDetailResponse_Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *LocationDetailResponse_Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *LocationDetailResponse_Comment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LocationDetailResponse_Comment) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationsRequest)(nil), "LocationsRequest")
	proto.RegisterType((*Location)(nil), "Location")
	proto.RegisterType((*LocationsResponse)(nil), "LocationsResponse")
	proto.RegisterType((*LocationDetailRequest)(nil), "LocationDetailRequest")
	proto.RegisterType((*LocationDetailResponse)(nil), "LocationDetailResponse")
	proto.RegisterType((*LocationDetailResponse_Comment)(nil), "LocationDetailResponse.Comment")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Locations service

type LocationsClient interface {
	RequestLocations(ctx context.Context, in *LocationsRequest, opts ...grpc.CallOption) (*LocationsResponse, error)
	RequestLocationDetail(ctx context.Context, in *LocationDetailRequest, opts ...grpc.CallOption) (*LocationDetailResponse, error)
}

type locationsClient struct {
	cc *grpc.ClientConn
}

func NewLocationsClient(cc *grpc.ClientConn) LocationsClient {
	return &locationsClient{cc}
}

func (c *locationsClient) RequestLocations(ctx context.Context, in *LocationsRequest, opts ...grpc.CallOption) (*LocationsResponse, error) {
	out := new(LocationsResponse)
	err := grpc.Invoke(ctx, "/Locations/RequestLocations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) RequestLocationDetail(ctx context.Context, in *LocationDetailRequest, opts ...grpc.CallOption) (*LocationDetailResponse, error) {
	out := new(LocationDetailResponse)
	err := grpc.Invoke(ctx, "/Locations/RequestLocationDetail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Locations service

type LocationsServer interface {
	RequestLocations(context.Context, *LocationsRequest) (*LocationsResponse, error)
	RequestLocationDetail(context.Context, *LocationDetailRequest) (*LocationDetailResponse, error)
}

func RegisterLocationsServer(s *grpc.Server, srv LocationsServer) {
	s.RegisterService(&_Locations_serviceDesc, srv)
}

func _Locations_RequestLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).RequestLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Locations/RequestLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).RequestLocations(ctx, req.(*LocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_RequestLocationDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).RequestLocationDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Locations/RequestLocationDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).RequestLocationDetail(ctx, req.(*LocationDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Locations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestLocations",
			Handler:    _Locations_RequestLocations_Handler,
		},
		{
			MethodName: "RequestLocationDetail",
			Handler:    _Locations_RequestLocationDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "locations.proto",
}

func init() { proto.RegisterFile("locations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xd9, 0xb4, 0xfd, 0xda, 0xcc, 0x27, 0xda, 0x0e, 0xb6, 0x8d, 0x41, 0x6c, 0x08, 0x88,
	0x39, 0x25, 0x50, 0x0f, 0x22, 0x78, 0xd3, 0x8b, 0xa0, 0x97, 0x88, 0x0f, 0x90, 0xb6, 0x4b, 0x58,
	0xd8, 0xec, 0xd6, 0x66, 0x7b, 0x12, 0x2f, 0xbe, 0x82, 0xcf, 0xe4, 0xd9, 0x83, 0xaf, 0xe0, 0x83,
	0x48, 0x36, 0xd9, 0xa4, 0xd6, 0x7a, 0x9b, 0xf9, 0xcf, 0xcc, 0x7f, 0x77, 0x7e, 0xbb, 0x70, 0xc0,
	0xe5, 0x3c, 0x51, 0x4c, 0x8a, 0x3c, 0x5c, 0xae, 0xa4, 0x92, 0xee, 0x71, 0x2a, 0x65, 0xca, 0x69,
	0x94, 0x2c, 0x59, 0x94, 0x08, 0x21, 0xd5, 0x66, 0xd5, 0x47, 0xe8, 0xdf, 0x99, 0x81, 0x98, 0x3e,
	0xad, 0x69, 0xae, 0xfc, 0x15, 0xf4, 0x8c, 0x86, 0xfb, 0x60, 0xb1, 0x85, 0x43, 0x3c, 0x12, 0x74,
	0x62, 0x8b, 0x2d, 0xf0, 0x10, 0x3a, 0x8a, 0x29, 0x4e, 0x1d, 0xcb, 0x23, 0x81, 0x1d, 0x97, 0x09,
	0xba, 0xd0, 0xcb, 0xd7, 0xb3, 0xb2, 0xd0, 0xd2, 0x85, 0x3a, 0xc7, 0x3e, 0xb4, 0x78, 0xa2, 0x9c,
	0xb6, 0x47, 0x02, 0x2b, 0x2e, 0x42, 0x44, 0x68, 0x73, 0x29, 0x52, 0xa7, 0xa3, 0x25, 0x1d, 0xfb,
	0x57, 0x30, 0xd8, 0xb8, 0x47, 0xbe, 0x94, 0x22, 0xa7, 0x78, 0x06, 0x76, 0xbd, 0x8d, 0x43, 0xbc,
	0x56, 0xf0, 0x7f, 0x6a, 0x87, 0xa6, 0x2d, 0x6e, 0x6a, 0xfe, 0x05, 0x0c, 0x8d, 0x7c, 0x43, 0x55,
	0xc2, 0x78, 0xb5, 0x0a, 0x9e, 0x00, 0x98, 0xae, 0x5b, 0xb3, 0xc6, 0x86, 0xe2, 0x7f, 0x10, 0x18,
	0x6d, 0x4f, 0x56, 0x87, 0x9f, 0x42, 0xcf, 0x34, 0xea, 0xc1, 0x1f, 0x67, 0xd7, 0x25, 0xbc, 0x84,
	0xee, 0x5c, 0x66, 0x19, 0x15, 0xca, 0xb1, 0xf4, 0x0d, 0x27, 0xe1, 0x6e, 0xc3, 0xf0, 0xba, 0x6c,
	0x8b, 0x4d, 0xbf, 0xfb, 0x00, 0xdd, 0x4a, 0x43, 0xa7, 0x70, 0x11, 0xaa, 0x70, 0x21, 0x9a, 0x9f,
	0x49, 0x0b, 0x58, 0x22, 0xc9, 0x0c, 0x6f, 0x1d, 0x17, 0xb8, 0x59, 0x96, 0xa4, 0xf4, 0x71, 0xc5,
	0x0d, 0x6e, 0x93, 0x4f, 0xdf, 0x09, 0xd8, 0x35, 0x49, 0xbc, 0x87, 0x7e, 0x85, 0xa2, 0xd1, 0x06,
	0xe1, 0xf6, 0x8b, 0xbb, 0x18, 0xfe, 0x82, 0xef, 0xe3, 0xeb, 0xe7, 0xd7, 0x9b, 0xb5, 0x87, 0x10,
	0xd5, 0x9c, 0x91, 0xc1, 0x70, 0xcb, 0xae, 0xdc, 0x11, 0x47, 0xe1, 0x4e, 0xfe, 0xee, 0xf8, 0x0f,
	0x18, 0xfe, 0x44, 0xbb, 0x1f, 0xe1, 0xb8, 0x71, 0x8f, 0x9e, 0x9b, 0x87, 0x79, 0x99, 0xfd, 0xd3,
	0xff, 0xf3, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x34, 0xd2, 0xb3, 0xd0, 0x02, 0x00, 0x00,
}
