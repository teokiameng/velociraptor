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
	FlowContext_RUNNING    FlowContext_State = 0
	FlowContext_TERMINATED FlowContext_State = 1
	FlowContext_ERROR      FlowContext_State = 3
	// A well known flow will not queue any messages and always
	// forward messages to the worker:
	FlowContext_WELL_KNOWN FlowContext_State = 2
)

var FlowContext_State_name = map[int32]string{
	0: "RUNNING",
	1: "TERMINATED",
	3: "ERROR",
	2: "WELL_KNOWN",
}
var FlowContext_State_value = map[string]int32{
	"RUNNING":    0,
	"TERMINATED": 1,
	"ERROR":      3,
	"WELL_KNOWN": 2,
}

func (x FlowContext_State) String() string {
	return proto.EnumName(FlowContext_State_name, int32(x))
}
func (FlowContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flows_7acf84d061ec70d2, []int{1, 0}
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
	return fileDescriptor_flows_7acf84d061ec70d2, []int{0}
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
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FlowContext) Reset()         { *m = FlowContext{} }
func (m *FlowContext) String() string { return proto.CompactTextString(m) }
func (*FlowContext) ProtoMessage()    {}
func (*FlowContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_7acf84d061ec70d2, []int{1}
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
	return FlowContext_RUNNING
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
	FlowName                 string         `protobuf:"bytes,11,opt,name=flow_name,json=flowName,proto3" json:"flow_name,omitempty"`
	BaseSessionId            string         `protobuf:"bytes,12,opt,name=base_session_id,json=baseSessionId,proto3" json:"base_session_id,omitempty"`
	StartTime                uint64         `protobuf:"varint,15,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	LogsCollectionUrn        string         `protobuf:"bytes,17,opt,name=logs_collection_urn,json=logsCollectionUrn,proto3" json:"logs_collection_urn,omitempty"`
	WriteIntermediateResults bool           `protobuf:"varint,18,opt,name=write_intermediate_results,json=writeIntermediateResults,proto3" json:"write_intermediate_results,omitempty"`
	RequireFastpoll          bool           `protobuf:"varint,19,opt,name=require_fastpoll,json=requireFastpoll,proto3" json:"require_fastpoll,omitempty"`
	OriginalFlow             *FlowReference `protobuf:"bytes,22,opt,name=original_flow,json=originalFlow,proto3" json:"original_flow,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}       `json:"-"`
	XXX_unrecognized         []byte         `json:"-"`
	XXX_sizecache            int32          `json:"-"`
}

func (m *FlowRunnerArgs) Reset()         { *m = FlowRunnerArgs{} }
func (m *FlowRunnerArgs) String() string { return proto.CompactTextString(m) }
func (*FlowRunnerArgs) ProtoMessage()    {}
func (*FlowRunnerArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_flows_7acf84d061ec70d2, []int{2}
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
	return fileDescriptor_flows_7acf84d061ec70d2, []int{3}
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
	return fileDescriptor_flows_7acf84d061ec70d2, []int{4}
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
	return fileDescriptor_flows_7acf84d061ec70d2, []int{5}
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
	proto.RegisterType((*FlowContext)(nil), "proto.FlowContext")
	proto.RegisterType((*FlowRunnerArgs)(nil), "proto.FlowRunnerArgs")
	proto.RegisterType((*VelociraptorFlowState)(nil), "proto.VelociraptorFlowState")
	proto.RegisterType((*FlowMetaData)(nil), "proto.FlowMetaData")
	proto.RegisterType((*VInterrogateArgs)(nil), "proto.VInterrogateArgs")
	proto.RegisterEnum("proto.FlowContext_State", FlowContext_State_name, FlowContext_State_value)
	proto.RegisterExtension(E_FlowMetadata)
}

func init() { proto.RegisterFile("flows.proto", fileDescriptor_flows_7acf84d061ec70d2) }

var fileDescriptor_flows_7acf84d061ec70d2 = []byte{
	// 1745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xcf, 0x6f, 0x23, 0x49,
	0xf5, 0xff, 0x7a, 0x66, 0x92, 0xd8, 0x65, 0xc7, 0x71, 0x2a, 0xf9, 0x8e, 0x9a, 0x08, 0x41, 0x29,
	0xac, 0xd8, 0xec, 0x6e, 0xe8, 0x2c, 0xc9, 0x68, 0x59, 0x06, 0x56, 0x10, 0xe7, 0xc7, 0xc8, 0xda,
	0x49, 0x66, 0xb7, 0x93, 0xc9, 0x0a, 0x16, 0x61, 0x95, 0xbb, 0x9f, 0x9d, 0xda, 0x69, 0x57, 0xf5,
	0x54, 0x55, 0xc7, 0x63, 0x71, 0xe1, 0x86, 0xf8, 0x0b, 0xb8, 0x72, 0xe3, 0x86, 0x84, 0x10, 0x37,
	0x0e, 0x1c, 0xf8, 0x4b, 0xe0, 0x82, 0xf8, 0x1b, 0x40, 0x42, 0xef, 0x75, 0x75, 0xc6, 0x13, 0x16,
	0x09, 0x2e, 0x69, 0xd7, 0xab, 0xcf, 0xfb, 0x59, 0xef, 0x7d, 0xaa, 0xc2, 0xda, 0xe3, 0xdc, 0xcc,
	0x5c, 0x5c, 0x58, 0xe3, 0x0d, 0x5f, 0xa2, 0xcf, 0xd6, 0x57, 0x26, 0xc6, 0x4c, 0x72, 0xd8, 0xa3,
	0xd5, 0xa8, 0x1c, 0xef, 0x49, 0x3d, 0xaf, 0x10, 0x5b, 0xe2, 0xee, 0x56, 0x06, 0x2e, 0xb5, 0xaa,
	0xf0, 0xc6, 0x06, 0xc4, 0xe3, 0xd9, 0x6c, 0x16, 0xdf, 0x40, 0x6e, 0x52, 0x95, 0xc1, 0xab, 0x38,
	0x35, 0xd3, 0xbd, 0x89, 0xc9, 0xa5, 0x9e, 0xec, 0x55, 0x42, 0x2b, 0x11, 0x5c, 0x19, 0xd8, 0x73,
	0x30, 0x95, 0xda, 0xab, 0x34, 0xe8, 0x7e, 0xf4, 0xdf, 0xe9, 0xca, 0xd4, 0x2b, 0xa3, 0x5d, 0xb0,
	0x71, 0xf3, 0x32, 0xff, 0xdf, 0xd4, 0x53, 0x3b, 0x2f, 0xbc, 0x09, 0xda, 0x5f, 0x98, 0x51, 0xc8,
	0x7e, 0xfb, 0xf7, 0x0d, 0xb6, 0x7a, 0x9a, 0x9b, 0x59, 0x02, 0x63, 0xb0, 0xa0, 0x53, 0xe0, 0x4f,
	0xd8, 0x0a, 0x96, 0x67, 0xa8, 0xb2, 0xa8, 0x21, 0x1a, 0x3b, 0xad, 0x7e, 0xfc, 0x97, 0x7f, 0xfc,
	0xf5, 0xcf, 0x8d, 0x1d, 0xfe, 0xcd, 0xcb, 0x6b, 0x10, 0x0e, 0x9c, 0x53, 0x46, 0x0b, 0x95, 0x09,
	0x33, 0x16, 0xfe, 0x1a, 0x84, 0xad, 0x35, 0x33, 0x81, 0x7a, 0x71, 0xb2, 0x8c, 0x9f, 0x41, 0xc6,
	0x7f, 0xc4, 0x5a, 0x69, 0xae, 0x40, 0x7b, 0x34, 0x75, 0x8f, 0x4c, 0x7d, 0x9f, 0x4c, 0x7d, 0xc0,
	0x5a, 0x47, 0xb4, 0xf1, 0x3c, 0x39, 0xe7, 0xef, 0xa0, 0xd5, 0x0a, 0x87, 0x46, 0xbf, 0xc4, 0xa2,
	0xb0, 0x52, 0x0b, 0xa3, 0xe3, 0xa4, 0x59, 0xc1, 0x06, 0xd9, 0xf6, 0x9f, 0x96, 0x58, 0x1b, 0xa3,
	0x3e, 0x32, 0xda, 0xc3, 0x2b, 0xcf, 0xbf, 0xca, 0x5a, 0x23, 0x99, 0xbe, 0xf0, 0x56, 0xa6, 0x50,
	0x45, 0x9d, 0xbc, 0x16, 0xf0, 0x43, 0xd6, 0x0b, 0x81, 0x58, 0x70, 0xa6, 0xb4, 0x29, 0x38, 0x8a,
	0xa7, 0xbd, 0xff, 0xb0, 0xaa, 0x42, 0x5c, 0x85, 0x93, 0xd4, 0xbb, 0xc9, 0x5a, 0xfa, 0xa6, 0x80,
	0x3f, 0x62, 0xed, 0xd4, 0x82, 0xf4, 0x30, 0xf4, 0x6a, 0x0a, 0xd1, 0x7d, 0xd1, 0xd8, 0x79, 0xd0,
	0xdf, 0xa0, 0x6c, 0x56, 0x59, 0x3b, 0x39, 0x3e, 0x3d, 0x96, 0x1e, 0x70, 0x2b, 0x61, 0x15, 0xee,
	0x52, 0x4d, 0x81, 0x47, 0x6c, 0x85, 0x56, 0xc6, 0x46, 0x0f, 0x28, 0xa8, 0x7a, 0xc9, 0xbf, 0xc1,
	0x56, 0xd3, 0xd2, 0x5a, 0x8c, 0xc9, 0x79, 0xe9, 0x21, 0x5a, 0xa2, 0xfd, 0x4e, 0x10, 0x5e, 0xa0,
	0x8c, 0x3f, 0x66, 0xdd, 0x17, 0x2a, 0xcf, 0xc9, 0xa5, 0xf3, 0x72, 0x5a, 0x44, 0xcb, 0xff, 0xd9,
	0xef, 0x2a, 0x42, 0x2f, 0x6b, 0x24, 0xdf, 0x65, 0x5c, 0x83, 0x9f, 0x19, 0xfb, 0x62, 0x38, 0x9a,
	0x7b, 0x70, 0x43, 0x07, 0xda, 0x47, 0x2b, 0xa8, 0x9f, 0xf4, 0xc2, 0x4e, 0x1f, 0x37, 0x2e, 0x40,
	0x7b, 0xbe, 0xc3, 0x7a, 0x1a, 0x5e, 0xf9, 0xa1, 0x29, 0xfd, 0xc8, 0x94, 0x3a, 0xc3, 0x13, 0x6b,
	0x12, 0xb6, 0x8b, 0xf2, 0x67, 0x41, 0x3c, 0xc8, 0xf8, 0x23, 0xf6, 0x90, 0x90, 0x85, 0x35, 0x29,
	0x38, 0x07, 0xd9, 0xd0, 0xc2, 0xcb, 0x12, 0x9c, 0x8f, 0x5a, 0x84, 0xdf, 0xc4, 0xdd, 0x4f, 0xea,
	0xcd, 0xa4, 0xda, 0xe3, 0xdf, 0x66, 0x9b, 0xa6, 0xf4, 0xce, 0x4b, 0x9d, 0x29, 0x3d, 0xa9, 0x55,
	0x5c, 0xd4, 0x26, 0x9d, 0x8d, 0x85, 0xbd, 0xa0, 0xe1, 0xf8, 0xfb, 0x8c, 0x85, 0x5e, 0xc3, 0x60,
	0x56, 0xa9, 0x7d, 0xd6, 0x29, 0xf1, 0x36, 0x6b, 0x5d, 0x54, 0x3b, 0x83, 0xe3, 0xa4, 0x15, 0x40,
	0x83, 0x8c, 0xc7, 0x6c, 0xa9, 0xaa, 0x65, 0x57, 0x34, 0x76, 0xba, 0xfb, 0x51, 0x38, 0xdb, 0x85,
	0x3e, 0x89, 0xa9, 0xae, 0x49, 0x05, 0xe3, 0x0f, 0xd9, 0x32, 0xfe, 0x28, 0x5d, 0xb4, 0x46, 0xc5,
	0x0f, 0x2b, 0x3c, 0x9b, 0xd2, 0x81, 0x1d, 0x6a, 0xe3, 0xd5, 0x58, 0x41, 0x16, 0xf5, 0x44, 0x63,
	0xa7, 0x99, 0x74, 0x50, 0x78, 0x1e, 0x64, 0xdb, 0x3f, 0x60, 0x4b, 0xd5, 0x21, 0xb5, 0xd9, 0x4a,
	0xf2, 0xfc, 0xfc, 0x7c, 0x70, 0xfe, 0xa4, 0xf7, 0x7f, 0xbc, 0xcb, 0xd8, 0xe5, 0x49, 0x72, 0x36,
	0x38, 0x3f, 0xbc, 0x3c, 0x39, 0xee, 0x35, 0x78, 0x8b, 0x2d, 0x9d, 0x24, 0xc9, 0xb3, 0xa4, 0x77,
	0x1f, 0xb7, 0x3e, 0x3b, 0x79, 0xfa, 0x74, 0xf8, 0xf1, 0xf9, 0xb3, 0xcf, 0xce, 0x7b, 0xf7, 0xb6,
	0x7f, 0xd3, 0x65, 0x5d, 0x1a, 0xbc, 0x52, 0x6b, 0xb0, 0x87, 0x76, 0xe2, 0xf8, 0x88, 0x35, 0x0b,
	0xab, 0x8c, 0x55, 0x7e, 0x4e, 0x4d, 0xdc, 0xdd, 0xdf, 0x0a, 0x39, 0x3c, 0xb1, 0xf6, 0x0c, 0x9c,
	0x93, 0x13, 0x88, 0x3f, 0x09, 0x88, 0xfe, 0xbb, 0x54, 0x8c, 0xb7, 0xb8, 0xc0, 0x01, 0xaa, 0xf5,
	0x44, 0xe9, 0x70, 0x6a, 0x8c, 0x15, 0xfe, 0x5a, 0xb9, 0x6a, 0x20, 0xa3, 0x46, 0x72, 0x6b, 0x97,
	0x3b, 0xd6, 0xa5, 0xbc, 0xe6, 0x43, 0x6f, 0x86, 0x98, 0x11, 0x4d, 0x42, 0xb3, 0x7f, 0x46, 0xd6,
	0x9e, 0xf0, 0xbd, 0x8b, 0x6b, 0x53, 0xe6, 0x99, 0x90, 0xa2, 0x4a, 0x3f, 0x95, 0x48, 0x3b, 0x62,
	0x84, 0x93, 0xaf, 0xbd, 0xf0, 0x86, 0x26, 0x54, 0x69, 0xe5, 0x15, 0x36, 0x72, 0xbc, 0xbd, 0x49,
	0x25, 0x99, 0x0b, 0xe9, 0xc5, 0x91, 0x99, 0x16, 0x39, 0x20, 0x3e, 0xe9, 0x54, 0x4e, 0x2e, 0xcd,
	0x73, 0x07, 0x96, 0x6b, 0xd6, 0x71, 0xa0, 0xb1, 0x55, 0x8a, 0x5c, 0x81, 0xa3, 0xf1, 0x69, 0xf6,
	0x3f, 0x26, 0x97, 0x27, 0xfc, 0x20, 0xb8, 0x44, 0x88, 0x08, 0x10, 0x81, 0xd3, 0x5b, 0x7b, 0x2b,
	0x24, 0xce, 0x44, 0xc5, 0x05, 0xc6, 0x62, 0x60, 0x71, 0x74, 0x6f, 0xbb, 0x73, 0x81, 0xf8, 0xa4,
	0xc2, 0x27, 0x6d, 0xd4, 0x0e, 0x0b, 0x7e, 0xb5, 0xc8, 0x3c, 0x34, 0x59, 0xfd, 0xef, 0x92, 0xb3,
	0x83, 0x45, 0xe6, 0x79, 0xeb, 0x2e, 0xf3, 0x84, 0x82, 0x09, 0x53, 0x80, 0x95, 0x1e, 0x1c, 0x92,
	0x4e, 0x74, 0xef, 0x35, 0xed, 0xf0, 0x01, 0x5b, 0x7a, 0x59, 0x42, 0x09, 0x34, 0x87, 0xad, 0xfe,
	0x01, 0xd9, 0xfc, 0x16, 0x5b, 0x4e, 0x8e, 0x4f, 0xd1, 0xe0, 0xd7, 0xd0, 0x20, 0x01, 0x30, 0xf2,
	0xd2, 0x41, 0x38, 0x08, 0x08, 0xe7, 0x70, 0x6f, 0xbf, 0x71, 0x9a, 0x54, 0x16, 0xf8, 0xa7, 0xac,
	0x09, 0x37, 0x21, 0xc2, 0x15, 0xb2, 0xf6, 0x01, 0x59, 0x7b, 0x9f, 0xef, 0x1e, 0x8a, 0xdc, 0x4c,
	0x26, 0x4a, 0x4f, 0x04, 0x21, 0x30, 0x36, 0xb4, 0xa3, 0x9c, 0x2b, 0x51, 0x38, 0x2e, 0xad, 0xbf,
	0x06, 0x8b, 0x20, 0x87, 0xe1, 0xad, 0x10, 0x6a, 0x90, 0xf1, 0x9f, 0xb2, 0x56, 0x5a, 0x94, 0xc3,
	0x5c, 0x4d, 0x55, 0x98, 0xc6, 0xfe, 0x21, 0xd9, 0xfc, 0x1e, 0x7f, 0x74, 0x28, 0x48, 0x2c, 0x8c,
	0xa6, 0x78, 0x42, 0xd6, 0x69, 0x51, 0x0a, 0x07, 0xa9, 0xd1, 0x99, 0xab, 0x1a, 0x67, 0x34, 0x7f,
	0xa3, 0x6f, 0x3e, 0xfc, 0xf5, 0x87, 0x49, 0x33, 0x2d, 0xca, 0xa7, 0xa8, 0xcb, 0x15, 0xdb, 0x78,
	0x93, 0x52, 0x2a, 0x4f, 0xab, 0xe4, 0x29, 0xd4, 0x97, 0xef, 0xdd, 0xf1, 0xe4, 0x8d, 0x97, 0xb9,
	0xf0, 0x56, 0x8e, 0xc7, 0x2a, 0xfd, 0x52, 0x27, 0xc9, 0xfa, 0x22, 0x1d, 0x55, 0xae, 0x7e, 0xcc,
	0x5a, 0x74, 0x07, 0x69, 0x39, 0x05, 0x22, 0x89, 0x56, 0xff, 0x23, 0x72, 0xf0, 0x1d, 0x7e, 0x80,
	0x45, 0x46, 0x79, 0x7d, 0xff, 0xa4, 0xb9, 0x74, 0x4e, 0x28, 0x6c, 0xbd, 0x29, 0x68, 0x8f, 0x25,
	0xaa, 0x4b, 0x8e, 0xc7, 0x60, 0xcb, 0xea, 0x10, 0x71, 0x7d, 0x2e, 0xa7, 0xc0, 0x7f, 0xc6, 0xd6,
	0x46, 0xd2, 0xc1, 0x70, 0x81, 0x5d, 0x3a, 0xe4, 0xe1, 0x82, 0x3c, 0x9c, 0xdd, 0x1e, 0xe7, 0xe1,
	0x9d, 0xfb, 0x6e, 0xf1, 0x30, 0xd1, 0xac, 0x06, 0x1b, 0x8b, 0xc1, 0x18, 0xfb, 0x51, 0xb8, 0x02,
	0x52, 0x22, 0x06, 0x31, 0x03, 0x31, 0x95, 0x2f, 0x40, 0x18, 0x0d, 0xe8, 0x77, 0x15, 0x7d, 0x5d,
	0xdc, 0x72, 0xd4, 0x2f, 0x1b, 0x8c, 0x39, 0x2f, 0xad, 0xaf, 0xee, 0x91, 0x35, 0xaa, 0xdd, 0x17,
	0xe4, 0x38, 0x7b, 0x83, 0xcf, 0xf9, 0xc5, 0xb1, 0x21, 0xcb, 0x81, 0x68, 0x17, 0xda, 0xb3, 0xd4,
	0x5e, 0xe5, 0xd5, 0x1a, 0x91, 0xb1, 0xd8, 0x19, 0x4c, 0xab, 0xa1, 0xb9, 0x8d, 0x4e, 0x39, 0x0c,
	0x50, 0x48, 0x37, 0xd7, 0xa9, 0x35, 0xda, 0x94, 0x2e, 0x9f, 0xc7, 0xef, 0x60, 0xb1, 0x5b, 0xe4,
	0x9d, 0x6e, 0xa7, 0x29, 0xdb, 0xc0, 0x1e, 0x1a, 0xa6, 0x26, 0xcf, 0x81, 0x5e, 0x17, 0xc3, 0xd2,
	0xea, 0x68, 0x7d, 0xb1, 0xdc, 0xb7, 0xc5, 0x78, 0x0f, 0x8b, 0x81, 0x70, 0xf1, 0x1a, 0x8e, 0xe5,
	0xcd, 0xcd, 0x04, 0x3f, 0x77, 0x1a, 0x3d, 0x59, 0x47, 0xe8, 0xd1, 0x2d, 0xf2, 0xb9, 0xd5, 0xfc,
	0x77, 0x0d, 0xb6, 0x35, 0xb3, 0xca, 0xc3, 0x50, 0x69, 0x0f, 0x76, 0x0a, 0x99, 0xc2, 0xfb, 0xd4,
	0x82, 0x2b, 0x73, 0xef, 0x22, 0x4e, 0x9c, 0xe0, 0xc8, 0xed, 0x94, 0xab, 0xc1, 0x58, 0x78, 0x5b,
	0xc2, 0xae, 0x90, 0x79, 0x2e, 0xd2, 0x6b, 0x95, 0xd7, 0x6f, 0x81, 0x0a, 0x2e, 0x2c, 0xa4, 0xa0,
	0x6e, 0xb0, 0xde, 0xca, 0x5f, 0x13, 0x75, 0x20, 0x73, 0xcc, 0xc5, 0x4c, 0xe5, 0x39, 0xb2, 0x15,
	0x7a, 0xf2, 0xa0, 0x6b, 0x0a, 0x41, 0xdd, 0xb7, 0x9d, 0xc8, 0x60, 0x2c, 0xcb, 0xdc, 0x2f, 0xe4,
	0x80, 0x35, 0x89, 0x28, 0xac, 0xc1, 0x42, 0x54, 0x49, 0xe5, 0x85, 0xff, 0xad, 0xc1, 0x7a, 0x78,
	0x59, 0x29, 0x0b, 0xc3, 0xb1, 0x74, 0xbe, 0x30, 0x79, 0x1e, 0x6d, 0x50, 0xa4, 0x7f, 0x6c, 0x50,
	0xa8, 0x7f, 0x68, 0xf0, 0xdf, 0x36, 0x2e, 0xf1, 0x2c, 0x6e, 0x64, 0x5e, 0x02, 0x56, 0xbe, 0x90,
	0x78, 0xf5, 0xd5, 0x9e, 0xc3, 0xa8, 0x65, 0xa5, 0xc5, 0xce, 0x3c, 0x92, 0x79, 0x5e, 0x31, 0xd0,
	0xae, 0x90, 0x9a, 0x66, 0xa1, 0x8e, 0x89, 0x22, 0x4f, 0x25, 0xd2, 0xc6, 0x82, 0x9e, 0x37, 0x02,
	0x30, 0x30, 0x51, 0xbb, 0x17, 0x72, 0xec, 0xc1, 0x86, 0xae, 0xa8, 0x9b, 0x7d, 0x1a, 0x2e, 0x07,
	0x81, 0x74, 0xeb, 0x84, 0x23, 0x42, 0xd5, 0x6f, 0x7b, 0x31, 0x35, 0x19, 0x92, 0x33, 0x35, 0x8b,
	0x03, 0x8f, 0xd3, 0x11, 0x47, 0x8d, 0x1f, 0x36, 0x92, 0xb5, 0x90, 0xd6, 0x69, 0x30, 0xcb, 0x7f,
	0xd5, 0x60, 0xab, 0xc6, 0xaa, 0x89, 0xd2, 0x32, 0x1f, 0x62, 0xb9, 0xa2, 0x87, 0xf4, 0x44, 0xda,
	0x5c, 0xb8, 0x46, 0x6f, 0x1f, 0x89, 0xfd, 0x9f, 0x50, 0xf2, 0x57, 0x3c, 0xc1, 0x73, 0xba, 0xed,
	0x4c, 0xe5, 0x84, 0x14, 0xa9, 0x29, 0xe6, 0x38, 0x9b, 0x52, 0x1b, 0xe2, 0x29, 0xdc, 0xd9, 0xc5,
	0xb1, 0x70, 0xde, 0x58, 0x10, 0xf2, 0xf5, 0xf3, 0xae, 0x2e, 0x51, 0xed, 0x56, 0x5c, 0x83, 0xa5,
	0xa1, 0xe9, 0xd4, 0x12, 0x74, 0xba, 0xfd, 0x39, 0xfb, 0xff, 0xab, 0x85, 0xa7, 0x2c, 0xca, 0xaa,
	0xab, 0x97, 0xb3, 0x07, 0x7e, 0x5e, 0xd4, 0x0f, 0x3e, 0xfa, 0xcd, 0x63, 0xb6, 0x52, 0xc8, 0x79,
	0x6e, 0x64, 0x16, 0x9e, 0x78, 0x9b, 0x71, 0xf5, 0x7a, 0x8f, 0xeb, 0xd7, 0x7b, 0x7c, 0xa8, 0xe7,
	0x49, 0x0d, 0xda, 0x7e, 0x97, 0x75, 0xd0, 0xe0, 0x19, 0x78, 0x79, 0x2c, 0xbd, 0xe4, 0x5b, 0xac,
	0x99, 0x4a, 0x0f, 0x13, 0x63, 0xe7, 0xc1, 0xee, 0xed, 0x7a, 0xbb, 0x64, 0xbd, 0x2b, 0xea, 0x12,
	0x6b, 0x26, 0xd2, 0x03, 0xdd, 0xd9, 0x82, 0xb5, 0x73, 0x35, 0xb9, 0xf6, 0x33, 0xc0, 0xbf, 0xa4,
	0xd2, 0x4c, 0x16, 0x45, 0xfc, 0x3d, 0xb6, 0xf2, 0xb2, 0x04, 0xab, 0xe8, 0xd1, 0x79, 0x7f, 0xa7,
	0xbd, 0xbf, 0x1e, 0x2a, 0x7a, 0xf5, 0xe9, 0xd3, 0xf0, 0xda, 0x49, 0x6a, 0xc4, 0x63, 0xfe, 0xf7,
	0x5f, 0x88, 0x2e, 0xeb, 0x2c, 0xe6, 0xfb, 0xf8, 0x73, 0xb6, 0x4a, 0x64, 0x38, 0x05, 0x2f, 0x33,
	0x8c, 0xf1, 0xeb, 0xff, 0x96, 0x52, 0x78, 0x1c, 0x3c, 0x2b, 0xe8, 0x7f, 0x84, 0xe8, 0x9f, 0x3f,
	0x7f, 0x40, 0xb9, 0x6f, 0x2c, 0x9c, 0x5d, 0x9d, 0x61, 0xd2, 0x19, 0x87, 0x15, 0xda, 0x1a, 0x2d,
	0x13, 0xe2, 0xe0, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xf1, 0xc2, 0x32, 0x13, 0x0d, 0x00,
	0x00,
}