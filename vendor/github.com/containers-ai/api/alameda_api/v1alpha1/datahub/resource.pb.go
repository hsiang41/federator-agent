// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/resource.proto

package containers_ai_alameda_v1alpha1_datahub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
// Represents kubernetes resource kind
//
type Kind int32

const (
	Kind_POD              Kind = 0
	Kind_DEPLOYMENT       Kind = 1
	Kind_DEPLOYMENTCONFIG Kind = 2
	Kind_ALAMEDASCALER    Kind = 3
	Kind_STATEFULSET      Kind = 4
)

var Kind_name = map[int32]string{
	0: "POD",
	1: "DEPLOYMENT",
	2: "DEPLOYMENTCONFIG",
	3: "ALAMEDASCALER",
	4: "STATEFULSET",
}

var Kind_value = map[string]int32{
	"POD":              0,
	"DEPLOYMENT":       1,
	"DEPLOYMENTCONFIG": 2,
	"ALAMEDASCALER":    3,
	"STATEFULSET":      4,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}

func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{0}
}

//*
// Represents a container and its containing limit and requeset configurations
//
type Container struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitResource        []*MetricData    `protobuf:"bytes,2,rep,name=limit_resource,json=limitResource,proto3" json:"limit_resource,omitempty"`
	RequestResource      []*MetricData    `protobuf:"bytes,3,rep,name=request_resource,json=requestResource,proto3" json:"request_resource,omitempty"`
	Status               *ContainerStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{0}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetLimitResource() []*MetricData {
	if m != nil {
		return m.LimitResource
	}
	return nil
}

func (m *Container) GetRequestResource() []*MetricData {
	if m != nil {
		return m.RequestResource
	}
	return nil
}

func (m *Container) GetStatus() *ContainerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

//*
// Represents a Kubernetes pod
//
type Pod struct {
	NamespacedName         *NamespacedName       `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ResourceLink           string                `protobuf:"bytes,2,opt,name=resource_link,json=resourceLink,proto3" json:"resource_link,omitempty"`
	Containers             []*Container          `protobuf:"bytes,3,rep,name=containers,proto3" json:"containers,omitempty"`
	IsAlameda              bool                  `protobuf:"varint,4,opt,name=is_alameda,json=isAlameda,proto3" json:"is_alameda,omitempty"`
	AlamedaScaler          *NamespacedName       `protobuf:"bytes,5,opt,name=alameda_scaler,json=alamedaScaler,proto3" json:"alameda_scaler,omitempty"`
	NodeName               string                `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	StartTime              *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Policy                 RecommendationPolicy  `protobuf:"varint,8,opt,name=policy,proto3,enum=containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy" json:"policy,omitempty"`
	TopController          *TopController        `protobuf:"bytes,9,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	UsedRecommendationId   string                `protobuf:"bytes,10,opt,name=used_recommendation_id,json=usedRecommendationId,proto3" json:"used_recommendation_id,omitempty"`
	Status                 *PodStatus            `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Enable_VPA             bool                  `protobuf:"varint,12,opt,name=enable_VPA,json=enableVPA,proto3" json:"enable_VPA,omitempty"`
	Enable_HPA             bool                  `protobuf:"varint,13,opt,name=enable_HPA,json=enableHPA,proto3" json:"enable_HPA,omitempty"`
	AppName                string                `protobuf:"bytes,14,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppPartOf              string                `protobuf:"bytes,15,opt,name=app_part_of,json=appPartOf,proto3" json:"app_part_of,omitempty"`
	AlamedaScalerResources *ResourceRequirements `protobuf:"bytes,16,opt,name=alameda_scaler_resources,json=alamedaScalerResources,proto3" json:"alameda_scaler_resources,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{1}
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

func (m *Pod) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *Pod) GetResourceLink() string {
	if m != nil {
		return m.ResourceLink
	}
	return ""
}

func (m *Pod) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Pod) GetIsAlameda() bool {
	if m != nil {
		return m.IsAlameda
	}
	return false
}

func (m *Pod) GetAlamedaScaler() *NamespacedName {
	if m != nil {
		return m.AlamedaScaler
	}
	return nil
}

func (m *Pod) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Pod) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Pod) GetPolicy() RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED
}

func (m *Pod) GetTopController() *TopController {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *Pod) GetUsedRecommendationId() string {
	if m != nil {
		return m.UsedRecommendationId
	}
	return ""
}

func (m *Pod) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Pod) GetEnable_VPA() bool {
	if m != nil {
		return m.Enable_VPA
	}
	return false
}

func (m *Pod) GetEnable_HPA() bool {
	if m != nil {
		return m.Enable_HPA
	}
	return false
}

func (m *Pod) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Pod) GetAppPartOf() string {
	if m != nil {
		return m.AppPartOf
	}
	return ""
}

func (m *Pod) GetAlamedaScalerResources() *ResourceRequirements {
	if m != nil {
		return m.AlamedaScalerResources
	}
	return nil
}

//*
// Represents the capacity of a Kubernetes node
//
type Capacity struct {
	CpuCores                 int64    `protobuf:"varint,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	MemoryBytes              int64    `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	NetwotkMegabitsPerSecond int64    `protobuf:"varint,3,opt,name=netwotk_megabits_per_second,json=netwotkMegabitsPerSecond,proto3" json:"netwotk_megabits_per_second,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Capacity) Reset()         { *m = Capacity{} }
func (m *Capacity) String() string { return proto.CompactTextString(m) }
func (*Capacity) ProtoMessage()    {}
func (*Capacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{2}
}

func (m *Capacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capacity.Unmarshal(m, b)
}
func (m *Capacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capacity.Marshal(b, m, deterministic)
}
func (m *Capacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capacity.Merge(m, src)
}
func (m *Capacity) XXX_Size() int {
	return xxx_messageInfo_Capacity.Size(m)
}
func (m *Capacity) XXX_DiscardUnknown() {
	xxx_messageInfo_Capacity.DiscardUnknown(m)
}

var xxx_messageInfo_Capacity proto.InternalMessageInfo

func (m *Capacity) GetCpuCores() int64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Capacity) GetMemoryBytes() int64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

func (m *Capacity) GetNetwotkMegabitsPerSecond() int64 {
	if m != nil {
		return m.NetwotkMegabitsPerSecond
	}
	return 0
}

//*
// Represents a Kubernetes node
//
type Node struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             *Capacity            `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Provider             *Provider            `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{3}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetCapacity() *Capacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *Node) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Node) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

//*
// Represents top controller of the pod
//
type TopController struct {
	NamespacedName       *NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 Kind            `protobuf:"varint,2,opt,name=kind,proto3,enum=containers_ai.alameda.v1alpha1.datahub.Kind" json:"kind,omitempty"`
	Replicas             int32           `protobuf:"varint,3,opt,name=Replicas,proto3" json:"Replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TopController) Reset()         { *m = TopController{} }
func (m *TopController) String() string { return proto.CompactTextString(m) }
func (*TopController) ProtoMessage()    {}
func (*TopController) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{4}
}

func (m *TopController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopController.Unmarshal(m, b)
}
func (m *TopController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopController.Marshal(b, m, deterministic)
}
func (m *TopController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopController.Merge(m, src)
}
func (m *TopController) XXX_Size() int {
	return xxx_messageInfo_TopController.Size(m)
}
func (m *TopController) XXX_DiscardUnknown() {
	xxx_messageInfo_TopController.DiscardUnknown(m)
}

var xxx_messageInfo_TopController proto.InternalMessageInfo

func (m *TopController) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *TopController) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_POD
}

func (m *TopController) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

type ResourceInfo struct {
	NamespacedName       *NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 Kind            `protobuf:"varint,2,opt,name=kind,proto3,enum=containers_ai.alameda.v1alpha1.datahub.Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResourceInfo) Reset()         { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()    {}
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{5}
}

func (m *ResourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceInfo.Unmarshal(m, b)
}
func (m *ResourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceInfo.Marshal(b, m, deterministic)
}
func (m *ResourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceInfo.Merge(m, src)
}
func (m *ResourceInfo) XXX_Size() int {
	return xxx_messageInfo_ResourceInfo.Size(m)
}
func (m *ResourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceInfo proto.InternalMessageInfo

func (m *ResourceInfo) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ResourceInfo) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_POD
}

type Controller struct {
	ControllerInfo                *ResourceInfo        `protobuf:"bytes,1,opt,name=controller_info,json=controllerInfo,proto3" json:"controller_info,omitempty"`
	OwnerInfo                     []*ResourceInfo      `protobuf:"bytes,2,rep,name=owner_info,json=ownerInfo,proto3" json:"owner_info,omitempty"`
	Replicas                      int32                `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	EnableRecommendationExecution bool                 `protobuf:"varint,4,opt,name=enable_recommendation_execution,json=enableRecommendationExecution,proto3" json:"enable_recommendation_execution,omitempty"`
	Policy                        RecommendationPolicy `protobuf:"varint,5,opt,name=policy,proto3,enum=containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy" json:"policy,omitempty"`
	SpecReplicas                  int32                `protobuf:"varint,6,opt,name=spec_replicas,json=specReplicas,proto3" json:"spec_replicas,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}             `json:"-"`
	XXX_unrecognized              []byte               `json:"-"`
	XXX_sizecache                 int32                `json:"-"`
}

func (m *Controller) Reset()         { *m = Controller{} }
func (m *Controller) String() string { return proto.CompactTextString(m) }
func (*Controller) ProtoMessage()    {}
func (*Controller) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{6}
}

func (m *Controller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Controller.Unmarshal(m, b)
}
func (m *Controller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Controller.Marshal(b, m, deterministic)
}
func (m *Controller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Controller.Merge(m, src)
}
func (m *Controller) XXX_Size() int {
	return xxx_messageInfo_Controller.Size(m)
}
func (m *Controller) XXX_DiscardUnknown() {
	xxx_messageInfo_Controller.DiscardUnknown(m)
}

var xxx_messageInfo_Controller proto.InternalMessageInfo

func (m *Controller) GetControllerInfo() *ResourceInfo {
	if m != nil {
		return m.ControllerInfo
	}
	return nil
}

func (m *Controller) GetOwnerInfo() []*ResourceInfo {
	if m != nil {
		return m.OwnerInfo
	}
	return nil
}

func (m *Controller) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Controller) GetEnableRecommendationExecution() bool {
	if m != nil {
		return m.EnableRecommendationExecution
	}
	return false
}

func (m *Controller) GetPolicy() RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED
}

func (m *Controller) GetSpecReplicas() int32 {
	if m != nil {
		return m.SpecReplicas
	}
	return 0
}

type Provider struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	InstanceType         string   `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Os                   string   `protobuf:"bytes,5,opt,name=os,proto3" json:"os,omitempty"`
	Role                 string   `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	InstanceId           string   `protobuf:"bytes,7,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	StorageSize          int64    `protobuf:"varint,8,opt,name=storage_size,json=storageSize,proto3" json:"storage_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{7}
}

func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Provider) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *Provider) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Provider) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Provider) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *Provider) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Provider) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Provider) GetStorageSize() int64 {
	if m != nil {
		return m.StorageSize
	}
	return 0
}

func init() {
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.Kind", Kind_name, Kind_value)
	proto.RegisterType((*Container)(nil), "containers_ai.alameda.v1alpha1.datahub.Container")
	proto.RegisterType((*Pod)(nil), "containers_ai.alameda.v1alpha1.datahub.Pod")
	proto.RegisterType((*Capacity)(nil), "containers_ai.alameda.v1alpha1.datahub.Capacity")
	proto.RegisterType((*Node)(nil), "containers_ai.alameda.v1alpha1.datahub.Node")
	proto.RegisterType((*TopController)(nil), "containers_ai.alameda.v1alpha1.datahub.TopController")
	proto.RegisterType((*ResourceInfo)(nil), "containers_ai.alameda.v1alpha1.datahub.ResourceInfo")
	proto.RegisterType((*Controller)(nil), "containers_ai.alameda.v1alpha1.datahub.Controller")
	proto.RegisterType((*Provider)(nil), "containers_ai.alameda.v1alpha1.datahub.Provider")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/resource.proto", fileDescriptor_e45e09c80210b55a)
}

var fileDescriptor_e45e09c80210b55a = []byte{
	// 1107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xef, 0x6e, 0x23, 0x35,
	0x10, 0x27, 0x7f, 0x9a, 0x26, 0x93, 0x26, 0x0d, 0xd6, 0xa9, 0x5a, 0x7a, 0x3a, 0x5a, 0x72, 0x12,
	0x2a, 0x07, 0x4a, 0x69, 0x38, 0x40, 0x48, 0x20, 0x11, 0xd2, 0x94, 0xab, 0x48, 0xdb, 0xe0, 0x84,
	0x93, 0x2a, 0x51, 0xad, 0x9c, 0x5d, 0xb7, 0x67, 0x35, 0xbb, 0xf6, 0xd9, 0x4e, 0x8f, 0xf4, 0x0d,
	0x78, 0x0a, 0xbe, 0xf3, 0x8d, 0x17, 0xe1, 0x01, 0xf8, 0xc2, 0x33, 0xf0, 0x06, 0xc8, 0x5e, 0xef,
	0x36, 0xa9, 0x50, 0x49, 0x0f, 0xf8, 0xc0, 0x37, 0xcf, 0x78, 0xe6, 0xe7, 0x99, 0x9f, 0x67, 0xc6,
	0x86, 0xf7, 0xc9, 0x84, 0x44, 0x34, 0x24, 0x3e, 0x11, 0x6c, 0xf7, 0x6a, 0x8f, 0x4c, 0xc4, 0x0b,
	0xb2, 0xb7, 0x1b, 0x12, 0x4d, 0x5e, 0x4c, 0xc7, 0xbb, 0x92, 0x2a, 0x3e, 0x95, 0x01, 0x6d, 0x09,
	0xc9, 0x35, 0x47, 0xef, 0x06, 0x3c, 0xd6, 0x84, 0xc5, 0x54, 0x2a, 0x9f, 0xb0, 0x96, 0x73, 0x6d,
	0xa5, 0x6e, 0x2d, 0xe7, 0xb6, 0xf9, 0xe4, 0x4e, 0x50, 0x21, 0x69, 0xc8, 0x02, 0x9d, 0x60, 0x6e,
	0x6e, 0x5d, 0x70, 0x7e, 0x31, 0xa1, 0xbb, 0x56, 0x1a, 0x4f, 0xcf, 0x77, 0x35, 0x8b, 0xa8, 0xd2,
	0x24, 0x12, 0xce, 0xe0, 0xee, 0x08, 0x23, 0xaa, 0x89, 0x59, 0x3b, 0xe3, 0xf7, 0xfe, 0xce, 0x58,
	0xb2, 0xc0, 0x99, 0xee, 0xdc, 0x69, 0xaa, 0x67, 0x82, 0xaa, 0xc4, 0xb2, 0xf9, 0x73, 0x1e, 0x2a,
	0xdd, 0x34, 0x73, 0x84, 0xa0, 0x18, 0x93, 0x88, 0x7a, 0xb9, 0xed, 0xdc, 0x4e, 0x05, 0xdb, 0x35,
	0x3a, 0x85, 0xfa, 0x84, 0x45, 0x4c, 0xfb, 0x29, 0x61, 0x5e, 0x7e, 0xbb, 0xb0, 0x53, 0x6d, 0xb7,
	0x5b, 0xcb, 0x31, 0xd6, 0x3a, 0xb2, 0x91, 0xed, 0x13, 0x4d, 0x70, 0xcd, 0x22, 0x61, 0x07, 0x84,
	0xce, 0xa0, 0x21, 0xe9, 0xcb, 0x29, 0x55, 0x73, 0xe0, 0x85, 0xd7, 0x06, 0x5f, 0x77, 0x58, 0x19,
	0xfc, 0x09, 0x94, 0x94, 0x26, 0x7a, 0xaa, 0xbc, 0xe2, 0x76, 0x6e, 0xa7, 0xda, 0xfe, 0x74, 0x59,
	0xd0, 0x8c, 0x90, 0xa1, 0x75, 0xc7, 0x0e, 0xa6, 0xf9, 0xdb, 0x2a, 0x14, 0x06, 0x3c, 0x44, 0x3e,
	0xac, 0x1b, 0x6a, 0x94, 0x20, 0x01, 0x0d, 0xfd, 0x8c, 0xb1, 0x6a, 0xfb, 0x93, 0x65, 0x4f, 0x38,
	0xce, 0xdc, 0xcd, 0x0a, 0xd7, 0xe3, 0x05, 0x19, 0x3d, 0x86, 0x5a, 0x4a, 0x88, 0x3f, 0x61, 0xf1,
	0xa5, 0x97, 0xb7, 0x17, 0xb2, 0x96, 0x2a, 0xfb, 0x2c, 0xbe, 0x44, 0xdf, 0x02, 0xdc, 0x9c, 0xe6,
	0x78, 0xdb, 0xbb, 0x77, 0x8a, 0x78, 0x0e, 0x04, 0x3d, 0x02, 0x60, 0xca, 0x77, 0x4e, 0x96, 0xb5,
	0x32, 0xae, 0x30, 0xd5, 0x49, 0x14, 0xe8, 0x0c, 0xea, 0x69, 0x61, 0xa9, 0x80, 0x4c, 0xa8, 0xf4,
	0x56, 0xfe, 0x51, 0xda, 0x35, 0x67, 0x38, 0xb4, 0x60, 0xe8, 0x21, 0x54, 0x62, 0x1e, 0xd2, 0x84,
	0xd0, 0x92, 0xcd, 0xb8, 0x6c, 0x14, 0x96, 0x92, 0xcf, 0x00, 0x94, 0x26, 0x52, 0xfb, 0xa6, 0x87,
	0xbc, 0x55, 0x7b, 0xee, 0x66, 0x2b, 0x69, 0xb0, 0x56, 0xda, 0x60, 0xad, 0x51, 0xda, 0x60, 0xb8,
	0x62, 0xad, 0x8d, 0x8c, 0x46, 0x50, 0x12, 0x7c, 0xc2, 0x82, 0x99, 0x57, 0xde, 0xce, 0xed, 0xd4,
	0xdb, 0x9f, 0x2f, 0x1b, 0x2e, 0xa6, 0x01, 0x8f, 0x22, 0x1a, 0x87, 0x44, 0x33, 0x1e, 0x0f, 0x2c,
	0x06, 0x76, 0x58, 0xe8, 0x7b, 0xa8, 0x6b, 0x2e, 0x7c, 0x03, 0x25, 0xf9, 0xc4, 0x90, 0x51, 0xb1,
	0x41, 0x7d, 0xbc, 0x2c, 0xfa, 0x88, 0x8b, 0x6e, 0xe6, 0x8c, 0x6b, 0x7a, 0x5e, 0x44, 0x4f, 0x61,
	0x63, 0xaa, 0x68, 0xe8, 0xcb, 0x85, 0x10, 0x7c, 0x16, 0x7a, 0x60, 0x89, 0x79, 0x60, 0x76, 0x17,
	0xe3, 0x3b, 0x0c, 0xd1, 0x61, 0x56, 0xf1, 0x55, 0x1b, 0xcb, 0xd2, 0xe5, 0x30, 0xe0, 0xe1, 0x62,
	0xad, 0x9b, 0x52, 0xa0, 0x31, 0x19, 0x4f, 0xa8, 0xff, 0x7c, 0xd0, 0xf1, 0xd6, 0x92, 0x52, 0x48,
	0x34, 0xcf, 0x07, 0x9d, 0xb9, 0xed, 0x67, 0x83, 0x8e, 0x57, 0x9b, 0xdf, 0x7e, 0x36, 0xe8, 0xa0,
	0xb7, 0xa0, 0x4c, 0x84, 0x48, 0x6e, 0xb2, 0x6e, 0x03, 0x5e, 0x25, 0x42, 0xd8, 0x8b, 0x7c, 0x1b,
	0xaa, 0x66, 0x4b, 0x98, 0xbb, 0xe4, 0xe7, 0xde, 0xba, 0xdd, 0xad, 0x10, 0x21, 0x06, 0x44, 0xea,
	0x93, 0x73, 0x74, 0x05, 0xde, 0x62, 0x91, 0x65, 0xb3, 0x41, 0x79, 0x0d, 0x9b, 0xd5, 0x3d, 0xee,
	0x2f, 0x71, 0xc4, 0xf4, 0xe5, 0x94, 0x49, 0x1a, 0xd1, 0x58, 0x2b, 0xbc, 0xb1, 0x50, 0x74, 0xa9,
	0x89, 0x6a, 0xfe, 0x98, 0x83, 0x72, 0x97, 0x08, 0x12, 0x30, 0x3d, 0x33, 0xa5, 0x18, 0x88, 0xa9,
	0x1f, 0x70, 0x49, 0x95, 0xed, 0xed, 0x02, 0x2e, 0x07, 0x62, 0xda, 0x35, 0x32, 0x7a, 0x07, 0xd6,
	0x22, 0x1a, 0x71, 0x39, 0xf3, 0xc7, 0x33, 0x4d, 0x95, 0x6d, 0xce, 0x02, 0xae, 0x26, 0xba, 0xaf,
	0x8c, 0x0a, 0x7d, 0x01, 0x0f, 0x63, 0xaa, 0x5f, 0x71, 0x7d, 0xe9, 0x47, 0xf4, 0x82, 0x8c, 0x99,
	0x56, 0xbe, 0xa0, 0xd2, 0x57, 0x34, 0xe0, 0x71, 0xe8, 0x15, 0xac, 0x87, 0xe7, 0x4c, 0x8e, 0x9c,
	0xc5, 0x80, 0xca, 0xa1, 0xdd, 0x6f, 0xfe, 0x91, 0x83, 0xe2, 0x31, 0x0f, 0xe9, 0x5f, 0x0e, 0xe4,
	0x3e, 0x94, 0x03, 0x17, 0xa7, 0x3d, 0xba, 0xda, 0xfe, 0x70, 0xe9, 0xae, 0x77, 0x7e, 0x38, 0x43,
	0xb8, 0xd5, 0x57, 0x85, 0xfb, 0xf4, 0x55, 0x1f, 0xca, 0x42, 0xf2, 0x2b, 0x16, 0x52, 0xe9, 0x26,
	0xec, 0xd2, 0x81, 0x0c, 0x9c, 0x1f, 0xce, 0x10, 0x9a, 0xbf, 0xe6, 0xa0, 0xb6, 0xd0, 0x12, 0xff,
	0xfd, 0x98, 0xfd, 0x12, 0x8a, 0x97, 0x2c, 0x0e, 0x2d, 0x8b, 0xf5, 0xf6, 0x07, 0xcb, 0xa2, 0x7e,
	0xc3, 0xe2, 0x10, 0x5b, 0x4f, 0xb4, 0x09, 0x65, 0x4c, 0xc5, 0x84, 0x05, 0x44, 0x59, 0xee, 0x56,
	0x70, 0x26, 0x37, 0x7f, 0xc9, 0xc1, 0x5a, 0x5a, 0x5e, 0x87, 0xf1, 0x39, 0xff, 0x1f, 0xe4, 0xd3,
	0xfc, 0xa9, 0x00, 0x30, 0x77, 0x03, 0x67, 0xb0, 0x7e, 0x33, 0xdf, 0x7c, 0x16, 0x9f, 0x73, 0x17,
	0xf1, 0xd3, 0xfb, 0xb6, 0xa0, 0x21, 0x00, 0xd7, 0x6f, 0xc0, 0x2c, 0x21, 0x43, 0x00, 0xfe, 0x2a,
	0x4e, 0x91, 0x93, 0x6f, 0xc5, 0xeb, 0x21, 0x57, 0x2c, 0x8e, 0x05, 0xdd, 0x84, 0xb2, 0xbc, 0x75,
	0x25, 0xa9, 0x8c, 0x0e, 0x60, 0xcb, 0x4d, 0xad, 0x5b, 0x73, 0x95, 0xfe, 0x40, 0x83, 0xa9, 0x59,
	0xb9, 0x47, 0xef, 0x51, 0x62, 0xb6, 0x38, 0x60, 0x7b, 0xa9, 0xd1, 0xdc, 0x8b, 0xb2, 0xf2, 0x2f,
	0xbe, 0x28, 0x8f, 0xa1, 0xa6, 0x04, 0x0d, 0xfc, 0x2c, 0xfc, 0x92, 0x0d, 0x7f, 0xcd, 0x28, 0xb3,
	0xaa, 0xfa, 0x3d, 0x07, 0xe5, 0xb4, 0x7b, 0x4c, 0xae, 0x59, 0x07, 0x26, 0x23, 0x22, 0x93, 0x0d,
	0x1a, 0x8b, 0x95, 0x26, 0x71, 0x40, 0x7d, 0xf3, 0xe3, 0x4b, 0xff, 0x10, 0xa9, 0x72, 0x34, 0x13,
	0x14, 0x6d, 0x40, 0x49, 0xd2, 0x0b, 0x93, 0x77, 0xc1, 0xee, 0x3a, 0xc9, 0xcc, 0x9d, 0x6b, 0x1e,
	0x53, 0xcb, 0x46, 0x05, 0xdb, 0x35, 0xaa, 0x43, 0x9e, 0x2b, 0x9b, 0x70, 0x05, 0xe7, 0xb9, 0x32,
	0x36, 0x92, 0x4f, 0xd2, 0x97, 0xda, 0xae, 0xd1, 0x16, 0x54, 0xb3, 0x43, 0x59, 0x68, 0x9f, 0xe9,
	0x0a, 0x86, 0x54, 0x75, 0x18, 0x9a, 0xd9, 0xa9, 0x34, 0x97, 0xe4, 0x82, 0xfa, 0x8a, 0x5d, 0x53,
	0xfb, 0x22, 0x17, 0x70, 0xd5, 0xe9, 0x86, 0xec, 0x9a, 0x3e, 0x39, 0x85, 0xa2, 0xa9, 0x48, 0x64,
	0x3e, 0x5b, 0x27, 0xfb, 0x8d, 0x37, 0x50, 0x1d, 0x60, 0xbf, 0x37, 0xe8, 0x9f, 0x9c, 0x1e, 0xf5,
	0x8e, 0x47, 0x8d, 0x1c, 0x7a, 0x00, 0x8d, 0x1b, 0xb9, 0x7b, 0x72, 0x7c, 0x70, 0xf8, 0x75, 0x23,
	0x8f, 0xde, 0x84, 0x5a, 0xa7, 0xdf, 0x39, 0xea, 0xed, 0x77, 0x86, 0xdd, 0x4e, 0xbf, 0x87, 0x1b,
	0x05, 0xb4, 0x0e, 0xd5, 0xe1, 0xa8, 0x33, 0xea, 0x1d, 0x7c, 0xd7, 0x1f, 0xf6, 0x46, 0x8d, 0xe2,
	0xb8, 0x64, 0x07, 0xda, 0x47, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x48, 0xc8, 0xb5, 0xde, 0x1a,
	0x0c, 0x00, 0x00,
}
