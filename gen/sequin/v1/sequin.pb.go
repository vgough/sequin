// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: sequin/v1/sequin.proto

package sequinv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string           `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Operation *anypb.Any       `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Metadata  *RequestMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ExecRequest) Reset() {
	*x = ExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecRequest) ProtoMessage() {}

func (x *ExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecRequest.ProtoReflect.Descriptor instead.
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{0}
}

func (x *ExecRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ExecRequest) GetOperation() *anypb.Any {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *ExecRequest) GetMetadata() *RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string           `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Operation *anypb.Any       `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Metadata  *RequestMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{1}
}

func (x *StartRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *StartRequest) GetOperation() *anypb.Any {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *StartRequest) GetMetadata() *RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{2}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results  []*anypb.Any `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Metadata *RunMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetResults() []*anypb.Any {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *GetResponse) GetMetadata() *RunMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type FuncOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []*anypb.Any `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *FuncOperation) Reset() {
	*x = FuncOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FuncOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuncOperation) ProtoMessage() {}

func (x *FuncOperation) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuncOperation.ProtoReflect.Descriptor instead.
func (*FuncOperation) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{5}
}

func (x *FuncOperation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FuncOperation) GetArgs() []*anypb.Any {
	if x != nil {
		return x.Args
	}
	return nil
}

type ExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results  []*anypb.Any `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Metadata *RunMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ExecResponse) Reset() {
	*x = ExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResponse) ProtoMessage() {}

func (x *ExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResponse.ProtoReflect.Descriptor instead.
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{6}
}

func (x *ExecResponse) GetResults() []*anypb.Any {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ExecResponse) GetMetadata() *RunMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientVersion string            `protobuf:"bytes,1,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	Labels        map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{7}
}

func (x *RequestMetadata) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *RequestMetadata) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type RunMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Labels are passed-through from the ExecRequest.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The following fields are set by the service.
	Submitter     string                 `protobuf:"bytes,2,opt,name=submitter,proto3" json:"submitter,omitempty"`
	SubmittedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	StartedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	ServerVersion string                 `protobuf:"bytes,6,opt,name=server_version,json=serverVersion,proto3" json:"server_version,omitempty"`
	Status        *status.Status         `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RunMetadata) Reset() {
	*x = RunMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequin_v1_sequin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunMetadata) ProtoMessage() {}

func (x *RunMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_sequin_v1_sequin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunMetadata.ProtoReflect.Descriptor instead.
func (*RunMetadata) Descriptor() ([]byte, []int) {
	return file_sequin_v1_sequin_proto_rawDescGZIP(), []int{8}
}

func (x *RunMetadata) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RunMetadata) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *RunMetadata) GetSubmittedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

func (x *RunMetadata) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *RunMetadata) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *RunMetadata) GetServerVersion() string {
	if x != nil {
		return x.ServerVersion
	}
	return ""
}

func (x *RunMetadata) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_sequin_v1_sequin_proto protoreflect.FileDescriptor

var file_sequin_v1_sequin_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x71, 0x75,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65,
	0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x18, 0x50, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xaa,
	0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x18, 0x50, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0f, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x18, 0x50, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x79, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65,
	0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a,
	0x0d, 0x46, 0x75, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x7a, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65,
	0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xbb, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61,
	0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb4, 0x03, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x42, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x8a, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x71, 0x75, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x1f, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71,
	0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x72,
	0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x67,
	0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x04, 0x45, 0x78,
	0x65, 0x63, 0x12, 0x1e, 0x2e, 0x61, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71,
	0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x1e, 0xca, 0x41, 0x1b, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0b, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x30, 0x01, 0x42, 0xbb, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x30,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53,
	0x65, 0x71, 0x75, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x67, 0x6f, 0x75, 0x67, 0x68, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x53, 0x58, 0xaa, 0x02, 0x11, 0x41, 0x72, 0x67, 0x30, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x65,
	0x71, 0x75, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x41, 0x72, 0x67, 0x30, 0x6e, 0x65,
	0x74, 0x5c, 0x53, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x41, 0x72,
	0x67, 0x30, 0x6e, 0x65, 0x74, 0x5c, 0x53, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x41, 0x72,
	0x67, 0x30, 0x6e, 0x65, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x71, 0x75, 0x69, 0x6e, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sequin_v1_sequin_proto_rawDescOnce sync.Once
	file_sequin_v1_sequin_proto_rawDescData = file_sequin_v1_sequin_proto_rawDesc
)

func file_sequin_v1_sequin_proto_rawDescGZIP() []byte {
	file_sequin_v1_sequin_proto_rawDescOnce.Do(func() {
		file_sequin_v1_sequin_proto_rawDescData = protoimpl.X.CompressGZIP(file_sequin_v1_sequin_proto_rawDescData)
	})
	return file_sequin_v1_sequin_proto_rawDescData
}

var file_sequin_v1_sequin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_sequin_v1_sequin_proto_goTypes = []any{
	(*ExecRequest)(nil),             // 0: arg0net.sequin.v1.ExecRequest
	(*StartRequest)(nil),            // 1: arg0net.sequin.v1.StartRequest
	(*StartResponse)(nil),           // 2: arg0net.sequin.v1.StartResponse
	(*GetRequest)(nil),              // 3: arg0net.sequin.v1.GetRequest
	(*GetResponse)(nil),             // 4: arg0net.sequin.v1.GetResponse
	(*FuncOperation)(nil),           // 5: arg0net.sequin.v1.FuncOperation
	(*ExecResponse)(nil),            // 6: arg0net.sequin.v1.ExecResponse
	(*RequestMetadata)(nil),         // 7: arg0net.sequin.v1.RequestMetadata
	(*RunMetadata)(nil),             // 8: arg0net.sequin.v1.RunMetadata
	nil,                             // 9: arg0net.sequin.v1.RequestMetadata.LabelsEntry
	nil,                             // 10: arg0net.sequin.v1.RunMetadata.LabelsEntry
	(*anypb.Any)(nil),               // 11: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),   // 12: google.protobuf.Timestamp
	(*status.Status)(nil),           // 13: google.rpc.Status
	(*longrunningpb.Operation)(nil), // 14: google.longrunning.Operation
}
var file_sequin_v1_sequin_proto_depIdxs = []int32{
	11, // 0: arg0net.sequin.v1.ExecRequest.operation:type_name -> google.protobuf.Any
	7,  // 1: arg0net.sequin.v1.ExecRequest.metadata:type_name -> arg0net.sequin.v1.RequestMetadata
	11, // 2: arg0net.sequin.v1.StartRequest.operation:type_name -> google.protobuf.Any
	7,  // 3: arg0net.sequin.v1.StartRequest.metadata:type_name -> arg0net.sequin.v1.RequestMetadata
	11, // 4: arg0net.sequin.v1.GetResponse.results:type_name -> google.protobuf.Any
	8,  // 5: arg0net.sequin.v1.GetResponse.metadata:type_name -> arg0net.sequin.v1.RunMetadata
	11, // 6: arg0net.sequin.v1.FuncOperation.args:type_name -> google.protobuf.Any
	11, // 7: arg0net.sequin.v1.ExecResponse.results:type_name -> google.protobuf.Any
	8,  // 8: arg0net.sequin.v1.ExecResponse.metadata:type_name -> arg0net.sequin.v1.RunMetadata
	9,  // 9: arg0net.sequin.v1.RequestMetadata.labels:type_name -> arg0net.sequin.v1.RequestMetadata.LabelsEntry
	10, // 10: arg0net.sequin.v1.RunMetadata.labels:type_name -> arg0net.sequin.v1.RunMetadata.LabelsEntry
	12, // 11: arg0net.sequin.v1.RunMetadata.submitted_at:type_name -> google.protobuf.Timestamp
	12, // 12: arg0net.sequin.v1.RunMetadata.started_at:type_name -> google.protobuf.Timestamp
	12, // 13: arg0net.sequin.v1.RunMetadata.finished_at:type_name -> google.protobuf.Timestamp
	13, // 14: arg0net.sequin.v1.RunMetadata.status:type_name -> google.rpc.Status
	1,  // 15: arg0net.sequin.v1.SequinService.Start:input_type -> arg0net.sequin.v1.StartRequest
	3,  // 16: arg0net.sequin.v1.SequinService.Get:input_type -> arg0net.sequin.v1.GetRequest
	0,  // 17: arg0net.sequin.v1.SequinService.Exec:input_type -> arg0net.sequin.v1.ExecRequest
	2,  // 18: arg0net.sequin.v1.SequinService.Start:output_type -> arg0net.sequin.v1.StartResponse
	4,  // 19: arg0net.sequin.v1.SequinService.Get:output_type -> arg0net.sequin.v1.GetResponse
	14, // 20: arg0net.sequin.v1.SequinService.Exec:output_type -> google.longrunning.Operation
	18, // [18:21] is the sub-list for method output_type
	15, // [15:18] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_sequin_v1_sequin_proto_init() }
func file_sequin_v1_sequin_proto_init() {
	if File_sequin_v1_sequin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sequin_v1_sequin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExecRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FuncOperation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ExecResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RequestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sequin_v1_sequin_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RunMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sequin_v1_sequin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sequin_v1_sequin_proto_goTypes,
		DependencyIndexes: file_sequin_v1_sequin_proto_depIdxs,
		MessageInfos:      file_sequin_v1_sequin_proto_msgTypes,
	}.Build()
	File_sequin_v1_sequin_proto = out.File
	file_sequin_v1_sequin_proto_rawDesc = nil
	file_sequin_v1_sequin_proto_goTypes = nil
	file_sequin_v1_sequin_proto_depIdxs = nil
}
