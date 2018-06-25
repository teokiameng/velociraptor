// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flows.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
import any "github.com/golang/protobuf/ptypes/any"
import proto2 "www.velocidex.com/golang/velociraptor/actions/proto"
import proto1 "www.velocidex.com/golang/velociraptor/crypto/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FlowContext_State int32

const (
	FlowContext_UNSET      FlowContext_State = 0
	FlowContext_RUNNING    FlowContext_State = 1
	FlowContext_TERMINATED FlowContext_State = 2
	FlowContext_ERROR      FlowContext_State = 3
)

var FlowContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
}
var FlowContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
}

func (x FlowContext_State) String() string {
	return proto.EnumName(FlowContext_State_name, int32(x))
}
func (FlowContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{2, 0}
}

type FlowReference struct {
	FlowId               string   `protobuf:"bytes,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowReference) Reset()         { *m = FlowReference{} }
func (m *FlowReference) String() string { return proto.CompactTextString(m) }
func (*FlowReference) ProtoMessage()    {}
func (*FlowReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{0}
}
func (m *FlowReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowReference.Unmarshal(m, b)
}
func (m *FlowReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowReference.Marshal(b, m, deterministic)
}
func (dst *FlowReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowReference.Merge(dst, src)
}
func (m *FlowReference) XXX_Size() int {
	return xxx_messageInfo_FlowReference.Size(m)
}
func (m *FlowReference) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowReference.DiscardUnknown(m)
}

var xxx_messageInfo_FlowReference proto.InternalMessageInfo

func (m *FlowReference) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *FlowReference) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type StartFlowRequest struct {
	ClientId             string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Interrogate          *VInterrogateArgs `protobuf:"bytes,2,opt,name=interrogate,proto3" json:"interrogate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StartFlowRequest) Reset()         { *m = StartFlowRequest{} }
func (m *StartFlowRequest) String() string { return proto.CompactTextString(m) }
func (*StartFlowRequest) ProtoMessage()    {}
func (*StartFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{1}
}
func (m *StartFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFlowRequest.Unmarshal(m, b)
}
func (m *StartFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFlowRequest.Marshal(b, m, deterministic)
}
func (dst *StartFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFlowRequest.Merge(dst, src)
}
func (m *StartFlowRequest) XXX_Size() int {
	return xxx_messageInfo_StartFlowRequest.Size(m)
}
func (m *StartFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartFlowRequest proto.InternalMessageInfo

func (m *StartFlowRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *StartFlowRequest) GetInterrogate() *VInterrogateArgs {
	if m != nil {
		return m.Interrogate
	}
	return nil
}

// The flow context.
// Next field: 17
type FlowContext struct {
	Backtrace            string                  `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	ClientResources      *proto1.ClientResources `protobuf:"bytes,2,opt,name=client_resources,json=clientResources,proto3" json:"client_resources,omitempty"`
	CreateTime           uint64                  `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Creator              string                  `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	CurrentState         string                  `protobuf:"bytes,5,opt,name=current_state,json=currentState,proto3" json:"current_state,omitempty"`
	KillTimestamp        uint64                  `protobuf:"varint,6,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	NetworkBytesSent     uint64                  `protobuf:"varint,7,opt,name=network_bytes_sent,json=networkBytesSent,proto3" json:"network_bytes_sent,omitempty"`
	NextOutboundId       uint64                  `protobuf:"varint,8,opt,name=next_outbound_id,json=nextOutboundId,proto3" json:"next_outbound_id,omitempty"`
	NextProcessedRequest uint64                  `protobuf:"varint,9,opt,name=next_processed_request,json=nextProcessedRequest,proto3" json:"next_processed_request,omitempty"`
	//  repeated OutputPluginState output_plugins_states = 10;
	OutstandingRequests uint64 `protobuf:"varint,11,opt,name=outstanding_requests,json=outstandingRequests,proto3" json:"outstanding_requests,omitempty"`
	// DEPRECATED
	//  uint64 remaining_cpu_quota = 12;
	SessionId            string            `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	State                FlowContext_State `protobuf:"varint,14,opt,name=state,proto3,enum=proto.FlowContext_State" json:"state,omitempty"`
	Status               string            `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified         bool              `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	ActiveTime           uint64            `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FlowContext) Reset()         { *m = FlowContext{} }
func (m *FlowContext) String() string { return proto.CompactTextString(m) }
func (*FlowContext) ProtoMessage()    {}
func (*FlowContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{2}
}
func (m *FlowContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowContext.Unmarshal(m, b)
}
func (m *FlowContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowContext.Marshal(b, m, deterministic)
}
func (dst *FlowContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowContext.Merge(dst, src)
}
func (m *FlowContext) XXX_Size() int {
	return xxx_messageInfo_FlowContext.Size(m)
}
func (m *FlowContext) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowContext.DiscardUnknown(m)
}

var xxx_messageInfo_FlowContext proto.InternalMessageInfo

func (m *FlowContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *FlowContext) GetClientResources() *proto1.ClientResources {
	if m != nil {
		return m.ClientResources
	}
	return nil
}

func (m *FlowContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FlowContext) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FlowContext) GetCurrentState() string {
	if m != nil {
		return m.CurrentState
	}
	return ""
}

func (m *FlowContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *FlowContext) GetNetworkBytesSent() uint64 {
	if m != nil {
		return m.NetworkBytesSent
	}
	return 0
}

func (m *FlowContext) GetNextOutboundId() uint64 {
	if m != nil {
		return m.NextOutboundId
	}
	return 0
}

func (m *FlowContext) GetNextProcessedRequest() uint64 {
	if m != nil {
		return m.NextProcessedRequest
	}
	return 0
}

func (m *FlowContext) GetOutstandingRequests() uint64 {
	if m != nil {
		return m.OutstandingRequests
	}
	return 0
}

func (m *FlowContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *FlowContext) GetState() FlowContext_State {
	if m != nil {
		return m.State
	}
	return FlowContext_UNSET
}

func (m *FlowContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FlowContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *FlowContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

// Next field: 23
type FlowRunnerArgs struct {
	Priority          proto1.GrrMessage_Priority `protobuf:"varint,1,opt,name=priority,proto3,enum=proto.GrrMessage_Priority" json:"priority,omitempty"`
	NotifyToUser      bool                       `protobuf:"varint,2,opt,name=notify_to_user,json=notifyToUser,proto3" json:"notify_to_user,omitempty"`
	SendReplies       bool                       `protobuf:"varint,3,opt,name=send_replies,json=sendReplies,proto3" json:"send_replies,omitempty"`
	ClientId          string                     `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Queue             string                     `protobuf:"bytes,6,opt,name=queue,proto3" json:"queue,omitempty"`
	EventId           string                     `protobuf:"bytes,7,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	CpuLimit          uint64                     `protobuf:"varint,9,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	NetworkBytesLimit uint64                     `protobuf:"varint,13,opt,name=network_bytes_limit,json=networkBytesLimit,proto3" json:"network_bytes_limit,omitempty"`
	//
	// RequestState request_state = 10 [(sem_type) = {
	// description: "The request state of the parent flow.",
	// label: HIDDEN,
	// }];
	FlowName                 string            `protobuf:"bytes,11,opt,name=flow_name,json=flowName,proto3" json:"flow_name,omitempty"`
	BaseSessionId            string            `protobuf:"bytes,12,opt,name=base_session_id,json=baseSessionId,proto3" json:"base_session_id,omitempty"`
	StartTime                uint64            `protobuf:"varint,15,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	LogsCollectionUrn        string            `protobuf:"bytes,17,opt,name=logs_collection_urn,json=logsCollectionUrn,proto3" json:"logs_collection_urn,omitempty"`
	WriteIntermediateResults bool              `protobuf:"varint,18,opt,name=write_intermediate_results,json=writeIntermediateResults,proto3" json:"write_intermediate_results,omitempty"`
	RequireFastpoll          bool              `protobuf:"varint,19,opt,name=require_fastpoll,json=requireFastpoll,proto3" json:"require_fastpoll,omitempty"`
	OriginalFlow             *FlowReference    `protobuf:"bytes,22,opt,name=original_flow,json=originalFlow,proto3" json:"original_flow,omitempty"`
	Args                     *StartFlowRequest `protobuf:"bytes,23,opt,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *FlowRunnerArgs) Reset()         { *m = FlowRunnerArgs{} }
func (m *FlowRunnerArgs) String() string { return proto.CompactTextString(m) }
func (*FlowRunnerArgs) ProtoMessage()    {}
func (*FlowRunnerArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{3}
}
func (m *FlowRunnerArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowRunnerArgs.Unmarshal(m, b)
}
func (m *FlowRunnerArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowRunnerArgs.Marshal(b, m, deterministic)
}
func (dst *FlowRunnerArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowRunnerArgs.Merge(dst, src)
}
func (m *FlowRunnerArgs) XXX_Size() int {
	return xxx_messageInfo_FlowRunnerArgs.Size(m)
}
func (m *FlowRunnerArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowRunnerArgs.DiscardUnknown(m)
}

var xxx_messageInfo_FlowRunnerArgs proto.InternalMessageInfo

func (m *FlowRunnerArgs) GetPriority() proto1.GrrMessage_Priority {
	if m != nil {
		return m.Priority
	}
	return proto1.GrrMessage_LOW_PRIORITY
}

func (m *FlowRunnerArgs) GetNotifyToUser() bool {
	if m != nil {
		return m.NotifyToUser
	}
	return false
}

func (m *FlowRunnerArgs) GetSendReplies() bool {
	if m != nil {
		return m.SendReplies
	}
	return false
}

func (m *FlowRunnerArgs) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *FlowRunnerArgs) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *FlowRunnerArgs) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *FlowRunnerArgs) GetCpuLimit() uint64 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *FlowRunnerArgs) GetNetworkBytesLimit() uint64 {
	if m != nil {
		return m.NetworkBytesLimit
	}
	return 0
}

func (m *FlowRunnerArgs) GetFlowName() string {
	if m != nil {
		return m.FlowName
	}
	return ""
}

func (m *FlowRunnerArgs) GetBaseSessionId() string {
	if m != nil {
		return m.BaseSessionId
	}
	return ""
}

func (m *FlowRunnerArgs) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *FlowRunnerArgs) GetLogsCollectionUrn() string {
	if m != nil {
		return m.LogsCollectionUrn
	}
	return ""
}

func (m *FlowRunnerArgs) GetWriteIntermediateResults() bool {
	if m != nil {
		return m.WriteIntermediateResults
	}
	return false
}

func (m *FlowRunnerArgs) GetRequireFastpoll() bool {
	if m != nil {
		return m.RequireFastpoll
	}
	return false
}

func (m *FlowRunnerArgs) GetOriginalFlow() *FlowReference {
	if m != nil {
		return m.OriginalFlow
	}
	return nil
}

func (m *FlowRunnerArgs) GetArgs() *StartFlowRequest {
	if m != nil {
		return m.Args
	}
	return nil
}

// This is a short lived protobuf to hold various flow state
// information only valid between the start and end of the flow. It is
// used to keep state over multiple client round trips.
type VelociraptorFlowState struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Payload              *any.Any `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VelociraptorFlowState) Reset()         { *m = VelociraptorFlowState{} }
func (m *VelociraptorFlowState) String() string { return proto.CompactTextString(m) }
func (*VelociraptorFlowState) ProtoMessage()    {}
func (*VelociraptorFlowState) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{4}
}
func (m *VelociraptorFlowState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VelociraptorFlowState.Unmarshal(m, b)
}
func (m *VelociraptorFlowState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VelociraptorFlowState.Marshal(b, m, deterministic)
}
func (dst *VelociraptorFlowState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VelociraptorFlowState.Merge(dst, src)
}
func (m *VelociraptorFlowState) XXX_Size() int {
	return xxx_messageInfo_VelociraptorFlowState.Size(m)
}
func (m *VelociraptorFlowState) XXX_DiscardUnknown() {
	xxx_messageInfo_VelociraptorFlowState.DiscardUnknown(m)
}

var xxx_messageInfo_VelociraptorFlowState proto.InternalMessageInfo

func (m *VelociraptorFlowState) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VelociraptorFlowState) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type FlowMetaData struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowMetaData) Reset()         { *m = FlowMetaData{} }
func (m *FlowMetaData) String() string { return proto.CompactTextString(m) }
func (*FlowMetaData) ProtoMessage()    {}
func (*FlowMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{5}
}
func (m *FlowMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowMetaData.Unmarshal(m, b)
}
func (m *FlowMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowMetaData.Marshal(b, m, deterministic)
}
func (dst *FlowMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowMetaData.Merge(dst, src)
}
func (m *FlowMetaData) XXX_Size() int {
	return xxx_messageInfo_FlowMetaData.Size(m)
}
func (m *FlowMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_FlowMetaData proto.InternalMessageInfo

func (m *FlowMetaData) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

// Interrogate flow discovers information about the client.
type VInterrogateArgs struct {
	// If set a light weight version of the flow is run.
	Lightweight bool `protobuf:"varint,1,opt,name=lightweight,proto3" json:"lightweight,omitempty"`
	// Additional VQL queries to run.
	Queries              []*proto2.VQLRequest `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VInterrogateArgs) Reset()         { *m = VInterrogateArgs{} }
func (m *VInterrogateArgs) String() string { return proto.CompactTextString(m) }
func (*VInterrogateArgs) ProtoMessage()    {}
func (*VInterrogateArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_74399e6942031bcd, []int{6}
}
func (m *VInterrogateArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VInterrogateArgs.Unmarshal(m, b)
}
func (m *VInterrogateArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VInterrogateArgs.Marshal(b, m, deterministic)
}
func (dst *VInterrogateArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VInterrogateArgs.Merge(dst, src)
}
func (m *VInterrogateArgs) XXX_Size() int {
	return xxx_messageInfo_VInterrogateArgs.Size(m)
}
func (m *VInterrogateArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VInterrogateArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VInterrogateArgs proto.InternalMessageInfo

func (m *VInterrogateArgs) GetLightweight() bool {
	if m != nil {
		return m.Lightweight
	}
	return false
}

func (m *VInterrogateArgs) GetQueries() []*proto2.VQLRequest {
	if m != nil {
		return m.Queries
	}
	return nil
}

var E_FlowMetadata = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*FlowMetaData)(nil),
	Field:         65661,
	Name:          "proto.flow_metadata",
	Tag:           "bytes,65661,opt,name=flow_metadata,json=flowMetadata",
	Filename:      "flows.proto",
}

func init() {
	proto.RegisterType((*FlowReference)(nil), "proto.FlowReference")
	proto.RegisterType((*StartFlowRequest)(nil), "proto.StartFlowRequest")
	proto.RegisterType((*FlowContext)(nil), "proto.FlowContext")
	proto.RegisterType((*FlowRunnerArgs)(nil), "proto.FlowRunnerArgs")
	proto.RegisterType((*VelociraptorFlowState)(nil), "proto.VelociraptorFlowState")
	proto.RegisterType((*FlowMetaData)(nil), "proto.FlowMetaData")
	proto.RegisterType((*VInterrogateArgs)(nil), "proto.VInterrogateArgs")
	proto.RegisterEnum("proto.FlowContext_State", FlowContext_State_name, FlowContext_State_value)
	proto.RegisterExtension(E_FlowMetadata)
}

func init() { proto.RegisterFile("flows.proto", fileDescriptor_flows_74399e6942031bcd) }

var fileDescriptor_flows_74399e6942031bcd = []byte{
	// 1866 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdf, 0x6f, 0x1c, 0x49,
	0xf1, 0xff, 0x4e, 0x12, 0xc7, 0xde, 0x5e, 0xff, 0xd8, 0xb4, 0xf3, 0xcd, 0x0d, 0x01, 0x41, 0x6b,
	0x39, 0x11, 0xe7, 0x2e, 0x8c, 0x8f, 0x24, 0x3a, 0x8e, 0x1c, 0x91, 0xf0, 0xda, 0x49, 0xb4, 0xe2,
	0xe2, 0xdc, 0xcd, 0x3a, 0x91, 0xe0, 0x4e, 0xac, 0x7a, 0x67, 0x6a, 0x77, 0xfb, 0x32, 0xdb, 0x3d,
	0xe9, 0xee, 0xf1, 0x66, 0x85, 0x84, 0x78, 0x41, 0x88, 0xbf, 0x80, 0x57, 0xfe, 0x02, 0x10, 0x42,
	0xbc, 0xf1, 0xc8, 0x5f, 0x02, 0x2f, 0x88, 0x37, 0xde, 0x41, 0x42, 0x55, 0xd3, 0xe3, 0x8c, 0xcd,
	0x81, 0xe0, 0xc5, 0xeb, 0xae, 0xfe, 0xd4, 0x8f, 0xae, 0xae, 0xfa, 0x74, 0x0d, 0xeb, 0x4e, 0x0b,
	0xb3, 0x74, 0x49, 0x69, 0x8d, 0x37, 0x7c, 0x8d, 0x7e, 0x6e, 0x7e, 0x69, 0x66, 0xcc, 0xac, 0x80,
	0x7d, 0x5a, 0x4d, 0xaa, 0xe9, 0xbe, 0xd4, 0xab, 0x1a, 0x71, 0x53, 0x5c, 0xdc, 0xca, 0xc1, 0x65,
	0x56, 0x95, 0xde, 0xd8, 0x80, 0x78, 0xb0, 0x5c, 0x2e, 0x93, 0x53, 0x28, 0x4c, 0xa6, 0x72, 0x78,
	0x9d, 0x64, 0x66, 0xb1, 0x3f, 0x33, 0x85, 0xd4, 0xb3, 0xfd, 0x5a, 0x68, 0x25, 0x82, 0x6b, 0x03,
	0xfb, 0x0e, 0x16, 0x52, 0x7b, 0x95, 0x05, 0xdd, 0x87, 0xff, 0x9d, 0xae, 0xcc, 0xbc, 0x32, 0xda,
	0x05, 0x1b, 0xa7, 0xaf, 0x8a, 0xff, 0x4d, 0x3d, 0xb3, 0xab, 0xd2, 0x9b, 0xa0, 0xfd, 0xb9, 0x99,
	0x84, 0xd3, 0xf7, 0x7f, 0x17, 0xb1, 0xad, 0xc7, 0x85, 0x59, 0xa6, 0x30, 0x05, 0x0b, 0x3a, 0x03,
	0xfe, 0x84, 0xad, 0x63, 0x7a, 0xc6, 0x2a, 0x8f, 0x23, 0x11, 0xed, 0x75, 0x06, 0xc9, 0x9f, 0xfe,
	0xfe, 0xe7, 0x3f, 0x46, 0x7b, 0xfc, 0x1b, 0x27, 0x73, 0x10, 0x0e, 0x9c, 0x53, 0x46, 0x0b, 0x95,
	0x0b, 0x33, 0x15, 0x7e, 0x0e, 0xc2, 0x36, 0x9a, 0xb9, 0x40, 0xbd, 0x24, 0xbd, 0x8a, 0x3f, 0xc3,
	0x9c, 0xff, 0x80, 0x75, 0xb2, 0x42, 0x81, 0xf6, 0x68, 0xea, 0x12, 0x99, 0xfa, 0x2e, 0x99, 0x7a,
	0x9f, 0x75, 0x0e, 0x69, 0xe3, 0x79, 0x7a, 0xcc, 0x6f, 0xa3, 0xd5, 0x1a, 0x87, 0x46, 0xbf, 0xc0,
	0xa2, 0xb0, 0x52, 0x0b, 0xa3, 0x93, 0x74, 0xa3, 0x86, 0x0d, 0xf3, 0xfe, 0x6f, 0x22, 0xd6, 0x1b,
	0x79, 0x69, 0x7d, 0x1d, 0xfa, 0xab, 0x0a, 0x9c, 0xe7, 0x5f, 0x6e, 0xfb, 0xa3, 0xd0, 0xdf, 0x68,
	0xf0, 0x9f, 0xb0, 0xae, 0xd2, 0x1e, 0xac, 0x35, 0x33, 0xe9, 0x81, 0xc2, 0xe9, 0xde, 0x7d, 0xab,
	0x4e, 0x42, 0xf2, 0x62, 0xf8, 0x66, 0xeb, 0xc0, 0xce, 0xdc, 0xe0, 0x90, 0xe2, 0x7c, 0xc8, 0x3f,
	0x3c, 0x99, 0x2b, 0x57, 0x07, 0xa1, 0x9c, 0xc8, 0xc1, 0xa9, 0x99, 0x86, 0x5c, 0x78, 0x23, 0x66,
	0xd2, 0xcf, 0xc1, 0x0a, 0xa5, 0xa7, 0xc6, 0x2e, 0x24, 0x5e, 0x8d, 0x90, 0x13, 0x53, 0x79, 0x3a,
	0xc0, 0xdc, 0x38, 0x9f, 0xa4, 0x6d, 0x87, 0xfd, 0xbf, 0xad, 0xb1, 0x2e, 0x06, 0x7b, 0x68, 0xb4,
	0x87, 0xd7, 0x9e, 0x7f, 0x85, 0x75, 0x26, 0x32, 0x7b, 0xe9, 0xad, 0xcc, 0x20, 0x04, 0xfb, 0x46,
	0xc0, 0x0f, 0x58, 0x2f, 0x1c, 0xc5, 0x82, 0x33, 0x95, 0xcd, 0xc0, 0x85, 0x90, 0x6f, 0x84, 0x90,
	0xeb, 0x04, 0xa6, 0xcd, 0x6e, 0xba, 0x93, 0x9d, 0x17, 0xf0, 0xfb, 0xac, 0x9b, 0x59, 0x90, 0x1e,
	0xc6, 0x5e, 0x2d, 0x20, 0xbe, 0x2c, 0xa2, 0xbd, 0x2b, 0x83, 0x5d, 0x3a, 0xd7, 0x16, 0xeb, 0xa6,
	0x47, 0x8f, 0x8f, 0xa4, 0x07, 0xdc, 0x4a, 0x59, 0x8d, 0x3b, 0x51, 0x0b, 0xe0, 0x31, 0x5b, 0xa7,
	0x95, 0xb1, 0xf1, 0x15, 0x0a, 0xaa, 0x59, 0xf2, 0xaf, 0xb3, 0xad, 0xac, 0xb2, 0x16, 0x63, 0x72,
	0x1e, 0x53, 0xb8, 0x46, 0xfb, 0x9b, 0x41, 0x38, 0x42, 0x19, 0x7f, 0xc0, 0xb6, 0x5f, 0xaa, 0xa2,
	0x20, 0x97, 0xce, 0xcb, 0x45, 0x19, 0x5f, 0xfd, 0xf7, 0x7e, 0xb7, 0x10, 0x7a, 0xd2, 0x20, 0xf9,
	0x1d, 0xc6, 0x35, 0xf8, 0xa5, 0xb1, 0x2f, 0xc7, 0x93, 0x95, 0x07, 0x37, 0x76, 0xa0, 0x7d, 0xbc,
	0x8e, 0xfa, 0x69, 0x2f, 0xec, 0x0c, 0x70, 0x63, 0x04, 0xda, 0xf3, 0x3d, 0xd6, 0xd3, 0xf0, 0xda,
	0x8f, 0x4d, 0xe5, 0x27, 0xa6, 0xd2, 0x39, 0xde, 0xf9, 0x06, 0x61, 0xb7, 0x51, 0xfe, 0x2c, 0x88,
	0x87, 0x39, 0xbf, 0xcf, 0x6e, 0x10, 0xb2, 0xb4, 0x26, 0x03, 0xe7, 0x20, 0x1f, 0xdb, 0xba, 0x60,
	0xe2, 0x0e, 0xe1, 0xaf, 0xe3, 0xee, 0xc7, 0xcd, 0x66, 0x53, 0x4c, 0xdf, 0x62, 0xd7, 0x4d, 0xe5,
	0x9d, 0x97, 0x3a, 0x57, 0x7a, 0xd6, 0xa8, 0xb8, 0xb8, 0x4b, 0x3a, 0xbb, 0xad, 0xbd, 0xa0, 0xe1,
	0xf8, 0x7b, 0x8c, 0x85, 0xee, 0xc0, 0x60, 0xb6, 0xa8, 0xe0, 0xaf, 0xd1, 0xc1, 0xbb, 0xac, 0x33,
	0xaa, 0x77, 0x86, 0x47, 0x69, 0x27, 0x80, 0x86, 0x39, 0x4f, 0xd8, 0x5a, 0x9d, 0xcb, 0x6d, 0x11,
	0xed, 0x6d, 0xdf, 0x8d, 0xc3, 0xdd, 0xb6, 0xea, 0x24, 0xa1, 0xbc, 0xa6, 0x35, 0x8c, 0xdf, 0x60,
	0x57, 0xf1, 0x9f, 0xca, 0xc5, 0x3b, 0x94, 0xfc, 0xb0, 0xc2, 0xbb, 0xa9, 0x1c, 0xd8, 0xb1, 0x36,
	0x5e, 0x4d, 0x15, 0xe4, 0x71, 0x4f, 0x44, 0x7b, 0x1b, 0xe9, 0x26, 0x0a, 0x8f, 0x83, 0x0c, 0x0b,
	0x02, 0x39, 0xe4, 0x34, 0x14, 0xc4, 0xb5, 0xff, 0x50, 0x10, 0x35, 0x0e, 0xaf, 0xa6, 0xff, 0x80,
	0xad, 0xd5, 0x57, 0xdb, 0x61, 0x6b, 0xcf, 0x8f, 0x47, 0x8f, 0x4e, 0x7a, 0xff, 0xc7, 0xbb, 0x6c,
	0x3d, 0x7d, 0x7e, 0x7c, 0x3c, 0x3c, 0x7e, 0xd2, 0x8b, 0xf8, 0x36, 0x63, 0x27, 0x8f, 0xd2, 0xa7,
	0xc3, 0xe3, 0x83, 0x93, 0x47, 0x47, 0xbd, 0x4b, 0x88, 0x7b, 0x94, 0xa6, 0xcf, 0xd2, 0xde, 0xe5,
	0xfe, 0xcf, 0x76, 0xd8, 0x36, 0x35, 0x68, 0xa5, 0x35, 0x58, 0x6c, 0x2c, 0x3e, 0x61, 0x1b, 0xa5,
	0x55, 0xc6, 0x2a, 0xbf, 0xa2, 0xaa, 0xdf, 0xbe, 0x7b, 0x33, 0x1c, 0xfa, 0x89, 0xb5, 0x4f, 0xc1,
	0x39, 0x39, 0x83, 0xe4, 0xe3, 0x80, 0x18, 0xbc, 0x43, 0xd1, 0xbd, 0xcd, 0x05, 0x72, 0x44, 0xa3,
	0x27, 0x2a, 0x87, 0xc4, 0x60, 0xac, 0xf0, 0x4d, 0x73, 0x26, 0x71, 0x94, 0x9e, 0xd9, 0xe5, 0x8e,
	0x6d, 0x53, 0x22, 0x56, 0x63, 0x6f, 0xc6, 0x98, 0x02, 0x6a, 0x9d, 0x8d, 0xc1, 0x53, 0xb2, 0xf6,
	0x84, 0xef, 0x8f, 0xe6, 0xa6, 0x2a, 0x72, 0x21, 0x45, 0x9d, 0xaf, 0xac, 0x6e, 0xdf, 0x09, 0x92,
	0x9b, 0xf6, 0xd8, 0xdc, 0xd8, 0xc3, 0x4a, 0x2b, 0xaf, 0xb0, 0xf2, 0x93, 0xfe, 0x75, 0xca, 0xe1,
	0x4a, 0x48, 0x2f, 0x0e, 0xcd, 0xa2, 0x2c, 0x00, 0xf1, 0xe9, 0x66, 0xed, 0xe4, 0xc4, 0x3c, 0x77,
	0x60, 0xb9, 0x66, 0x9b, 0x0e, 0x34, 0xd6, 0x56, 0x59, 0x28, 0x70, 0xd4, 0x6f, 0x1b, 0x83, 0xef,
	0x93, 0xcb, 0x47, 0xfc, 0x5e, 0x70, 0x89, 0x10, 0x11, 0x20, 0x02, 0xdb, 0xbd, 0xf1, 0x56, 0x4a,
	0x6c, 0xa2, 0x9a, 0x69, 0x8c, 0xc5, 0xc0, 0x92, 0xf8, 0x52, 0x7f, 0x73, 0x84, 0xf8, 0xb4, 0xc6,
	0xa7, 0x5d, 0xd4, 0x0e, 0x0b, 0xfe, 0xa2, 0x4d, 0x76, 0xd4, 0x8a, 0x83, 0xef, 0x90, 0xb3, 0x7b,
	0x6d, 0x72, 0x7d, 0xfb, 0x22, 0xb9, 0x36, 0x6c, 0x66, 0x4a, 0xb0, 0xd2, 0x83, 0x43, 0x5e, 0x8d,
	0x2f, 0xb5, 0x78, 0x72, 0xc8, 0xd6, 0x5e, 0x55, 0x50, 0x01, 0x35, 0x6e, 0x67, 0x70, 0x8f, 0x6c,
	0x7e, 0x93, 0x5d, 0x4d, 0x8f, 0x1e, 0xa3, 0xc1, 0xaf, 0xa2, 0x41, 0x02, 0x60, 0xe4, 0x95, 0x83,
	0x70, 0x11, 0x10, 0xee, 0xe1, 0xd2, 0xdd, 0xe8, 0x71, 0x5a, 0x5b, 0xe0, 0x9f, 0xb0, 0x0d, 0x38,
	0x0d, 0x11, 0xae, 0x93, 0xb5, 0xf7, 0xc9, 0xda, 0x7b, 0xfc, 0xce, 0x81, 0x28, 0xcc, 0x6c, 0xa6,
	0xf4, 0x4c, 0x10, 0x02, 0x63, 0x43, 0x3b, 0xca, 0xb9, 0x0a, 0x85, 0xd3, 0xca, 0x12, 0xbb, 0x16,
	0x66, 0xe6, 0x30, 0xbc, 0x75, 0x42, 0x0d, 0x73, 0xfe, 0x23, 0xd6, 0xc9, 0xca, 0x6a, 0x5c, 0xa8,
	0x85, 0x0a, 0xed, 0x3b, 0x38, 0x20, 0x9b, 0x1f, 0xf2, 0xfb, 0x07, 0x82, 0xc4, 0xc2, 0x68, 0x8a,
	0x27, 0x9c, 0x3a, 0x2b, 0x2b, 0xe1, 0x20, 0x33, 0x3a, 0x77, 0x75, 0xe1, 0x4c, 0x56, 0xe7, 0xea,
	0xe6, 0x83, 0x5f, 0x7d, 0x90, 0x6e, 0x64, 0x65, 0xf5, 0x11, 0xea, 0x72, 0xc5, 0x76, 0xcf, 0x73,
	0x50, 0xed, 0x69, 0x8b, 0x3c, 0x85, 0xfc, 0xf2, 0xfd, 0x0b, 0x9e, 0xbc, 0xf1, 0xb2, 0x10, 0xde,
	0xca, 0xe9, 0x54, 0x65, 0x5f, 0xe8, 0x24, 0xbd, 0xd6, 0xe6, 0xaf, 0xda, 0xd5, 0x0f, 0x59, 0x87,
	0x9e, 0x59, 0x2d, 0x17, 0x40, 0xac, 0xd2, 0x19, 0x3c, 0x24, 0x07, 0xdf, 0xe6, 0xf7, 0x30, 0xc9,
	0x28, 0x6f, 0x9e, 0xd8, 0xac, 0x90, 0xce, 0x09, 0x85, 0xa5, 0xb7, 0x00, 0xed, 0x31, 0x45, 0x4d,
	0xca, 0xf1, 0x1a, 0x6c, 0x55, 0x5f, 0x22, 0xae, 0x8f, 0xe5, 0x02, 0xf8, 0x8f, 0xd9, 0xce, 0x44,
	0x3a, 0x18, 0xb7, 0xe8, 0x68, 0x93, 0x3c, 0x8c, 0xc8, 0xc3, 0xd3, 0xb3, 0xeb, 0x3c, 0xb8, 0xf0,
	0xa4, 0xb7, 0x2f, 0x13, 0xcd, 0x6a, 0xb0, 0x89, 0x18, 0x4e, 0xb1, 0x1e, 0x85, 0x2b, 0x21, 0x23,
	0x26, 0x11, 0x4b, 0x10, 0x0b, 0xf9, 0x12, 0x84, 0xd1, 0x80, 0x7e, 0xb7, 0xd0, 0xd7, 0xe8, 0x8c,
	0xd4, 0x7e, 0x11, 0x31, 0xe6, 0xf0, 0x6d, 0xae, 0x79, 0x66, 0x87, 0x72, 0xf7, 0x39, 0x39, 0xce,
	0xcf, 0xf1, 0x0c, 0x1f, 0x1d, 0x19, 0xb2, 0x1c, 0x98, 0xb9, 0x55, 0x9e, 0x95, 0xf6, 0xaa, 0xa8,
	0xd7, 0x88, 0x4c, 0xc4, 0xde, 0x70, 0x51, 0x37, 0xcd, 0x59, 0x74, 0xca, 0x61, 0x80, 0x42, 0xba,
	0x95, 0xce, 0xac, 0xd1, 0xa6, 0x72, 0xc5, 0x2a, 0xb9, 0x8d, 0xc9, 0xee, 0x90, 0x77, 0x7a, 0xce,
	0x16, 0x6c, 0x17, 0x6b, 0x68, 0x9c, 0x99, 0xa2, 0x00, 0x1a, 0xa0, 0xc6, 0x95, 0xd5, 0xc4, 0x7d,
	0x67, 0xe9, 0x3e, 0x4b, 0xc6, 0xbb, 0x98, 0x0c, 0x84, 0x8b, 0x37, 0x70, 0x4c, 0x6f, 0x61, 0x66,
	0xf8, 0x73, 0xa1, 0xd0, 0xd3, 0x6b, 0x08, 0x3d, 0x3c, 0x43, 0x3e, 0xb7, 0x9a, 0xff, 0x36, 0x62,
	0x37, 0x97, 0x56, 0x79, 0x18, 0xd3, 0xd3, 0xbf, 0x80, 0x5c, 0xe1, 0x03, 0x6c, 0xc1, 0x55, 0x85,
	0x77, 0x31, 0x27, 0x4e, 0x70, 0xe4, 0x76, 0xc1, 0xd5, 0x70, 0x2a, 0xbc, 0xad, 0xe0, 0x8e, 0x90,
	0x45, 0x21, 0xb2, 0xb9, 0x2a, 0x9a, 0x71, 0xa7, 0x86, 0x0b, 0x0b, 0x19, 0xa8, 0x53, 0xcc, 0xb7,
	0xf2, 0x73, 0xa2, 0x0e, 0x64, 0x8e, 0x95, 0x58, 0xaa, 0xa2, 0x40, 0xb6, 0x42, 0x4f, 0x1e, 0x74,
	0x43, 0x21, 0xa8, 0x7b, 0x0b, 0x87, 0x94, 0xa9, 0xac, 0x0a, 0xdf, 0x3a, 0x03, 0xe6, 0x24, 0xa6,
	0xb0, 0x86, 0xad, 0xa8, 0xd2, 0xda, 0x0b, 0xff, 0x4b, 0xc4, 0x7a, 0xf8, 0xba, 0x29, 0x0b, 0xe3,
	0xa9, 0x74, 0xbe, 0x34, 0x45, 0x11, 0xef, 0x52, 0xa4, 0x7f, 0x88, 0x28, 0xd4, 0xdf, 0x47, 0xfc,
	0xd7, 0x11, 0x0d, 0x42, 0xa7, 0xb2, 0xa8, 0x00, 0x33, 0x5f, 0x4a, 0x7c, 0x2b, 0x1b, 0xcf, 0xa1,
	0xd5, 0xf2, 0xca, 0x62, 0x65, 0x1e, 0xca, 0xa2, 0xa8, 0x19, 0xe8, 0x8e, 0x90, 0x9a, 0x7a, 0xa1,
	0x89, 0x89, 0x22, 0xcf, 0x24, 0xd2, 0x46, 0x4b, 0xcf, 0x1b, 0x01, 0x18, 0x98, 0x68, 0xdc, 0x0b,
	0x39, 0xf5, 0x60, 0x43, 0x55, 0x34, 0xc5, 0xbe, 0x08, 0x8f, 0x83, 0x40, 0xba, 0x75, 0xc2, 0x11,
	0xa1, 0xea, 0x5b, 0x5e, 0x2c, 0x4c, 0x8e, 0xe4, 0x4c, 0xc5, 0xe2, 0xc0, 0x63, 0x77, 0x24, 0x71,
	0xf4, 0xbd, 0x28, 0xdd, 0x09, 0xc7, 0x7a, 0x1c, 0xcc, 0xf2, 0x5f, 0x46, 0x6c, 0xcb, 0x58, 0x35,
	0x53, 0x5a, 0x16, 0x63, 0x4c, 0x57, 0x7c, 0x83, 0x66, 0xaa, 0xeb, 0xad, 0x77, 0xf7, 0x6c, 0x0e,
	0x1e, 0x7c, 0x46, 0x87, 0x7f, 0xc1, 0x53, 0xbc, 0xa7, 0xf6, 0x18, 0x28, 0x45, 0x66, 0xca, 0x15,
	0xf6, 0xa6, 0xd4, 0x86, 0x78, 0x0a, 0x77, 0xee, 0x60, 0x5b, 0x38, 0x6f, 0x2c, 0x08, 0xf9, 0x66,
	0x82, 0x6d, 0x52, 0xd4, 0xb8, 0x15, 0x73, 0xb0, 0xd4, 0x34, 0x9b, 0x8d, 0x04, 0x9d, 0xf2, 0xcf,
	0xd8, 0x15, 0x69, 0x67, 0x2e, 0x7e, 0xeb, 0xdc, 0x58, 0x7a, 0x71, 0xc2, 0x1d, 0xec, 0x53, 0x48,
	0xb7, 0xf9, 0xad, 0x13, 0x9a, 0x92, 0x49, 0xd8, 0x8a, 0x6d, 0x29, 0x9d, 0xa0, 0xda, 0x0f, 0x35,
	0x93, 0xa4, 0x64, 0xb5, 0xff, 0x29, 0xfb, 0xff, 0x17, 0xad, 0x6f, 0x01, 0xb4, 0x58, 0xbf, 0xe9,
	0x9c, 0x5d, 0xf1, 0xab, 0xb2, 0x99, 0x3f, 0xe9, 0x7f, 0x9e, 0xb0, 0xf5, 0x52, 0xae, 0x0a, 0x23,
	0xf3, 0x30, 0x71, 0x5e, 0x4f, 0xea, 0xcf, 0x9f, 0xa4, 0xf9, 0xfc, 0x49, 0x0e, 0xf4, 0x2a, 0x6d,
	0x40, 0xfd, 0x77, 0xd8, 0x26, 0x1a, 0x7c, 0x0a, 0x5e, 0x1e, 0x49, 0x2f, 0xf9, 0x4d, 0xb6, 0x91,
	0x49, 0x0f, 0x33, 0x63, 0x57, 0x67, 0x43, 0x78, 0x58, 0xf7, 0x2b, 0xd6, 0xbb, 0x38, 0x6a, 0x73,
	0xc1, 0xba, 0x85, 0x9a, 0xcd, 0xfd, 0x12, 0xf0, 0x2f, 0xa9, 0x6c, 0xa4, 0x6d, 0x11, 0x7f, 0x97,
	0xad, 0xbf, 0xaa, 0xc0, 0x2a, 0x9a, 0x81, 0x2f, 0xef, 0x75, 0xef, 0x5e, 0x6b, 0xc6, 0xf6, 0x4f,
	0x3e, 0x0a, 0x99, 0x49, 0x1b, 0xc4, 0x03, 0xfe, 0xd7, 0x9f, 0x8b, 0x6d, 0xb6, 0xd9, 0x3e, 0xef,
	0x83, 0x4f, 0xd9, 0x16, 0x51, 0xed, 0x02, 0xbc, 0xcc, 0x31, 0xc6, 0xaf, 0xfd, 0xcb, 0x91, 0xc2,
	0xe8, 0xf1, 0xac, 0xa4, 0x8f, 0xac, 0xf8, 0x1f, 0x3f, 0xbd, 0x42, 0x67, 0xdf, 0x6d, 0x55, 0x46,
	0x73, 0xc2, 0x74, 0x73, 0x1a, 0x56, 0x68, 0x6b, 0x72, 0x95, 0x10, 0xf7, 0xfe, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x1c, 0xda, 0xfc, 0x54, 0x0e, 0x00, 0x00,
}
