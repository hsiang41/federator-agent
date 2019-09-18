// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/ai_service/ai_service.proto

package containers_ai_alameda_v1alpha1_ai_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

//*
//  Recommendation policy. A policy may be either stable or compact.
type RecommendationPolicy int32

const (
	RecommendationPolicy_STABLE  RecommendationPolicy = 0
	RecommendationPolicy_COMPACT RecommendationPolicy = 1
)

var RecommendationPolicy_name = map[int32]string{
	0: "STABLE",
	1: "COMPACT",
}

var RecommendationPolicy_value = map[string]int32{
	"STABLE":  0,
	"COMPACT": 1,
}

func (x RecommendationPolicy) String() string {
	return proto.EnumName(RecommendationPolicy_name, int32(x))
}

func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{0}
}

/// Types of an object
type Object_Type int32

const (
	Object_POD  Object_Type = 0
	Object_NODE Object_Type = 1
)

var Object_Type_name = map[int32]string{
	0: "POD",
	1: "NODE",
}

var Object_Type_value = map[string]int32{
	"POD":  0,
	"NODE": 1,
}

func (x Object_Type) String() string {
	return proto.EnumName(Object_Type_name, int32(x))
}

func (Object_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{1, 0}
}

//*
// Represents a Kubernetes pod.
//
type Pod struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ResourceLink         string   `protobuf:"bytes,2,opt,name=resourceLink,proto3" json:"resourceLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{0}
}

func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (m *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(m, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Pod) GetResourceLink() string {
	if m != nil {
		return m.ResourceLink
	}
	return ""
}

//*
// Represents a Kubernetes object.
type Object struct {
	Type                 Object_Type          `protobuf:"varint,1,opt,name=type,proto3,enum=containers_ai.alameda.v1alpha1.ai_service.Object_Type" json:"type,omitempty"`
	Policy               RecommendationPolicy `protobuf:"varint,2,opt,name=policy,proto3,enum=containers_ai.alameda.v1alpha1.ai_service.RecommendationPolicy" json:"policy,omitempty"`
	Uid                  string               `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Namespace            string               `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Pod                  *Pod                 `protobuf:"bytes,6,opt,name=pod,proto3" json:"pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{1}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetType() Object_Type {
	if m != nil {
		return m.Type
	}
	return Object_POD
}

func (m *Object) GetPolicy() RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return RecommendationPolicy_STABLE
}

func (m *Object) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Object) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Object) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

//*
// Represents a request for creating a list of prediction objects
type PredictionObjectListCreationRequest struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PredictionObjectListCreationRequest) Reset()         { *m = PredictionObjectListCreationRequest{} }
func (m *PredictionObjectListCreationRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionObjectListCreationRequest) ProtoMessage()    {}
func (*PredictionObjectListCreationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{2}
}

func (m *PredictionObjectListCreationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Unmarshal(m, b)
}
func (m *PredictionObjectListCreationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Marshal(b, m, deterministic)
}
func (m *PredictionObjectListCreationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionObjectListCreationRequest.Merge(m, src)
}
func (m *PredictionObjectListCreationRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Size(m)
}
func (m *PredictionObjectListCreationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionObjectListCreationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionObjectListCreationRequest proto.InternalMessageInfo

func (m *PredictionObjectListCreationRequest) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

//*
// Represents a request for removing a list of prediction objects
type PredictionObjectListDeletionRequest struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PredictionObjectListDeletionRequest) Reset()         { *m = PredictionObjectListDeletionRequest{} }
func (m *PredictionObjectListDeletionRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionObjectListDeletionRequest) ProtoMessage()    {}
func (*PredictionObjectListDeletionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{3}
}

func (m *PredictionObjectListDeletionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Unmarshal(m, b)
}
func (m *PredictionObjectListDeletionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Marshal(b, m, deterministic)
}
func (m *PredictionObjectListDeletionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionObjectListDeletionRequest.Merge(m, src)
}
func (m *PredictionObjectListDeletionRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Size(m)
}
func (m *PredictionObjectListDeletionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionObjectListDeletionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionObjectListDeletionRequest proto.InternalMessageInfo

func (m *PredictionObjectListDeletionRequest) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

//*
// Represents a reponse of a request
type RequestResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestResponse) Reset()         { *m = RequestResponse{} }
func (m *RequestResponse) String() string { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()    {}
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_facbad1e8d2a4adf, []int{4}
}

func (m *RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponse.Unmarshal(m, b)
}
func (m *RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponse.Marshal(b, m, deterministic)
}
func (m *RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponse.Merge(m, src)
}
func (m *RequestResponse) XXX_Size() int {
	return xxx_messageInfo_RequestResponse.Size(m)
}
func (m *RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponse proto.InternalMessageInfo

func (m *RequestResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.ai_service.RecommendationPolicy", RecommendationPolicy_name, RecommendationPolicy_value)
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.ai_service.Object_Type", Object_Type_name, Object_Type_value)
	proto.RegisterType((*Pod)(nil), "containers_ai.alameda.v1alpha1.ai_service.Pod")
	proto.RegisterType((*Object)(nil), "containers_ai.alameda.v1alpha1.ai_service.Object")
	proto.RegisterType((*PredictionObjectListCreationRequest)(nil), "containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest")
	proto.RegisterType((*PredictionObjectListDeletionRequest)(nil), "containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest")
	proto.RegisterType((*RequestResponse)(nil), "containers_ai.alameda.v1alpha1.ai_service.RequestResponse")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/ai_service/ai_service.proto", fileDescriptor_facbad1e8d2a4adf)
}

var fileDescriptor_facbad1e8d2a4adf = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xe3, 0x90, 0x34, 0x13, 0x54, 0xac, 0x15, 0xa2, 0x26, 0xe5, 0x10, 0x99, 0x4b, 0x00,
	0x69, 0x4d, 0x8c, 0xc4, 0x81, 0x0b, 0x84, 0x24, 0x48, 0x40, 0x48, 0x2c, 0x37, 0x12, 0xc7, 0x68,
	0x63, 0x0f, 0x61, 0xc1, 0xf6, 0x2e, 0x5e, 0xa7, 0x52, 0xce, 0x7c, 0x0c, 0xff, 0xc1, 0x07, 0xf0,
	0x4d, 0xc8, 0xeb, 0x58, 0xa5, 0x15, 0x88, 0xfa, 0xd0, 0xdb, 0xec, 0x8c, 0xe6, 0xcd, 0x7b, 0x6f,
	0x56, 0x03, 0x4f, 0x59, 0xcc, 0x12, 0x8c, 0xd8, 0x9a, 0x49, 0xee, 0x9e, 0x8f, 0x58, 0x2c, 0x3f,
	0xb3, 0x91, 0xcb, 0xf8, 0x5a, 0x61, 0x76, 0xce, 0x43, 0xfc, 0x23, 0xa4, 0x32, 0x13, 0xb9, 0x20,
	0x8f, 0x42, 0x91, 0xe6, 0x8c, 0xa7, 0x98, 0xa9, 0x35, 0xe3, 0xf4, 0xd0, 0x4f, 0xab, 0x5e, 0x7a,
	0xd1, 0xd0, 0x3f, 0xdd, 0x0a, 0xb1, 0x8d, 0xd1, 0xd5, 0x8d, 0x9b, 0xdd, 0x27, 0x17, 0x13, 0x99,
	0xef, 0x4b, 0x1c, 0xe7, 0x0d, 0x98, 0xbe, 0x88, 0xc8, 0x29, 0x74, 0x53, 0x11, 0xe1, 0x3a, 0x65,
	0x09, 0xda, 0xc6, 0xc0, 0x18, 0x76, 0x83, 0xa3, 0x22, 0xb1, 0x60, 0x09, 0x12, 0x07, 0x6e, 0x67,
	0xa8, 0xc4, 0x2e, 0x0b, 0x71, 0xce, 0xd3, 0xaf, 0x76, 0x53, 0xd7, 0x2f, 0xe5, 0x9c, 0x5f, 0x4d,
	0x68, 0x2f, 0x37, 0x5f, 0x30, 0xcc, 0xc9, 0x3b, 0x68, 0xe5, 0x7b, 0x59, 0xc2, 0x1c, 0x7b, 0xcf,
	0xe9, 0xb5, 0x99, 0xd2, 0x12, 0x80, 0xae, 0xf6, 0x12, 0x03, 0x8d, 0x41, 0x3e, 0x42, 0x5b, 0x8a,
	0x98, 0x87, 0x7b, 0x3d, 0xf4, 0xd8, 0x7b, 0x59, 0x03, 0x2d, 0xc0, 0x50, 0x24, 0x09, 0xa6, 0x11,
	0xcb, 0xb9, 0x48, 0x7d, 0x0d, 0x13, 0x1c, 0xe0, 0x88, 0x05, 0xe6, 0x8e, 0x47, 0xb6, 0xa9, 0xa5,
	0x14, 0x21, 0x79, 0x00, 0xdd, 0x42, 0xbd, 0x92, 0x2c, 0x44, 0xbb, 0xa5, 0xf3, 0x17, 0x09, 0x42,
	0xa0, 0xa5, 0xbd, 0xb9, 0xa5, 0x0b, 0x3a, 0x26, 0xaf, 0xc0, 0x94, 0x22, 0xb2, 0xdb, 0x03, 0x63,
	0xd8, 0xf3, 0x68, 0x0d, 0x66, 0xbe, 0x88, 0x82, 0xa2, 0xd5, 0xb9, 0x0f, 0xad, 0x42, 0x2c, 0xe9,
	0x80, 0xe9, 0x2f, 0xa7, 0x56, 0x83, 0x1c, 0x41, 0x6b, 0xb1, 0x9c, 0xce, 0x2c, 0xc3, 0xc9, 0xe0,
	0xa1, 0x9f, 0x61, 0xc4, 0xc3, 0x82, 0x7c, 0x69, 0xcc, 0x9c, 0xab, 0x7c, 0x92, 0xa1, 0x96, 0x13,
	0xe0, 0xb7, 0x1d, 0xaa, 0x9c, 0xbc, 0x87, 0x8e, 0xd0, 0x45, 0x65, 0x1b, 0x03, 0x73, 0xd8, 0xf3,
	0x46, 0xb5, 0xfd, 0x0e, 0x2a, 0x84, 0x7f, 0xcd, 0x9c, 0x62, 0x8c, 0x37, 0x36, 0xf3, 0x09, 0xdc,
	0x39, 0xe0, 0x06, 0xa8, 0xa4, 0x48, 0x15, 0x12, 0x1b, 0x3a, 0x09, 0x2a, 0xc5, 0xb6, 0xd5, 0x57,
	0xac, 0x9e, 0x8f, 0x5d, 0xb8, 0xfb, 0xb7, 0xad, 0x12, 0x80, 0xf6, 0xd9, 0x6a, 0xfc, 0x7a, 0x3e,
	0xb3, 0x1a, 0xa4, 0x07, 0x9d, 0xc9, 0xf2, 0x83, 0x3f, 0x9e, 0xac, 0x2c, 0xc3, 0xfb, 0xd9, 0x04,
	0x6b, 0x5c, 0xb2, 0x19, 0xbf, 0x3d, 0x2b, 0x29, 0x90, 0x1f, 0x06, 0x9c, 0x68, 0x1f, 0xf1, 0xaa,
	0x5a, 0x45, 0x16, 0x75, 0xd6, 0xf8, 0xff, 0xfd, 0xf4, 0x5f, 0xd4, 0xfa, 0xb0, 0x97, 0x7c, 0x70,
	0x1a, 0xe4, 0xbb, 0x01, 0x27, 0xda, 0xfd, 0x1b, 0x60, 0x7a, 0x65, 0xab, 0xfd, 0x7b, 0xb4, 0xbc,
	0x13, 0xb4, 0xba, 0x13, 0x74, 0x56, 0xdc, 0x09, 0xa7, 0xb1, 0x69, 0xeb, 0xcc, 0xb3, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x17, 0xea, 0xb8, 0x22, 0xa6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlamedaAIServiceClient is the client API for AlamedaAIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlamedaAIServiceClient interface {
	/// Used to create prediction objects
	CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error)
	/// Used to remove prediction objects
	DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type alamedaAIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlamedaAIServiceClient(cc *grpc.ClientConn) AlamedaAIServiceClient {
	return &alamedaAIServiceClient{cc}
}

func (c *alamedaAIServiceClient) CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error) {
	out := new(RequestResponse)
	err := c.cc.Invoke(ctx, "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alamedaAIServiceClient) DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlamedaAIServiceServer is the server API for AlamedaAIService service.
type AlamedaAIServiceServer interface {
	/// Used to create prediction objects
	CreatePredictionObjects(context.Context, *PredictionObjectListCreationRequest) (*RequestResponse, error)
	/// Used to remove prediction objects
	DeletePredictionObjects(context.Context, *PredictionObjectListDeletionRequest) (*empty.Empty, error)
}

// UnimplementedAlamedaAIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlamedaAIServiceServer struct {
}

func (*UnimplementedAlamedaAIServiceServer) CreatePredictionObjects(ctx context.Context, req *PredictionObjectListCreationRequest) (*RequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePredictionObjects not implemented")
}
func (*UnimplementedAlamedaAIServiceServer) DeletePredictionObjects(ctx context.Context, req *PredictionObjectListDeletionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePredictionObjects not implemented")
}

func RegisterAlamedaAIServiceServer(s *grpc.Server, srv AlamedaAIServiceServer) {
	s.RegisterService(&_AlamedaAIService_serviceDesc, srv)
}

func _AlamedaAIService_CreatePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamedaAIServiceServer).CreatePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamedaAIServiceServer).CreatePredictionObjects(ctx, req.(*PredictionObjectListCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlamedaAIService_DeletePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListDeletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamedaAIServiceServer).DeletePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamedaAIServiceServer).DeletePredictionObjects(ctx, req.(*PredictionObjectListDeletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlamedaAIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService",
	HandlerType: (*AlamedaAIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePredictionObjects",
			Handler:    _AlamedaAIService_CreatePredictionObjects_Handler,
		},
		{
			MethodName: "DeletePredictionObjects",
			Handler:    _AlamedaAIService_DeletePredictionObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alameda_api/v1alpha1/ai_service/ai_service.proto",
}
