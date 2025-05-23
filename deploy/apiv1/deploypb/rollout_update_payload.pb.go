// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/cloud/deploy/v1/rollout_update_payload.proto

package deploypb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RolloutUpdateType indicates the type of the rollout update.
type RolloutUpdateEvent_RolloutUpdateType int32

const (
	// Rollout update type unspecified.
	RolloutUpdateEvent_ROLLOUT_UPDATE_TYPE_UNSPECIFIED RolloutUpdateEvent_RolloutUpdateType = 0
	// rollout state updated to pending.
	RolloutUpdateEvent_PENDING RolloutUpdateEvent_RolloutUpdateType = 1
	// Rollout state updated to pending release.
	RolloutUpdateEvent_PENDING_RELEASE RolloutUpdateEvent_RolloutUpdateType = 2
	// Rollout state updated to in progress.
	RolloutUpdateEvent_IN_PROGRESS RolloutUpdateEvent_RolloutUpdateType = 3
	// Rollout state updated to cancelling.
	RolloutUpdateEvent_CANCELLING RolloutUpdateEvent_RolloutUpdateType = 4
	// Rollout state updated to cancelled.
	RolloutUpdateEvent_CANCELLED RolloutUpdateEvent_RolloutUpdateType = 5
	// Rollout state updated to halted.
	RolloutUpdateEvent_HALTED RolloutUpdateEvent_RolloutUpdateType = 6
	// Rollout state updated to succeeded.
	RolloutUpdateEvent_SUCCEEDED RolloutUpdateEvent_RolloutUpdateType = 7
	// Rollout state updated to failed.
	RolloutUpdateEvent_FAILED RolloutUpdateEvent_RolloutUpdateType = 8
	// Rollout requires approval.
	RolloutUpdateEvent_APPROVAL_REQUIRED RolloutUpdateEvent_RolloutUpdateType = 9
	// Rollout has been approved.
	RolloutUpdateEvent_APPROVED RolloutUpdateEvent_RolloutUpdateType = 10
	// Rollout has been rejected.
	RolloutUpdateEvent_REJECTED RolloutUpdateEvent_RolloutUpdateType = 11
	// Rollout requires advance to the next phase.
	RolloutUpdateEvent_ADVANCE_REQUIRED RolloutUpdateEvent_RolloutUpdateType = 12
	// Rollout has been advanced.
	RolloutUpdateEvent_ADVANCED RolloutUpdateEvent_RolloutUpdateType = 13
)

// Enum value maps for RolloutUpdateEvent_RolloutUpdateType.
var (
	RolloutUpdateEvent_RolloutUpdateType_name = map[int32]string{
		0:  "ROLLOUT_UPDATE_TYPE_UNSPECIFIED",
		1:  "PENDING",
		2:  "PENDING_RELEASE",
		3:  "IN_PROGRESS",
		4:  "CANCELLING",
		5:  "CANCELLED",
		6:  "HALTED",
		7:  "SUCCEEDED",
		8:  "FAILED",
		9:  "APPROVAL_REQUIRED",
		10: "APPROVED",
		11: "REJECTED",
		12: "ADVANCE_REQUIRED",
		13: "ADVANCED",
	}
	RolloutUpdateEvent_RolloutUpdateType_value = map[string]int32{
		"ROLLOUT_UPDATE_TYPE_UNSPECIFIED": 0,
		"PENDING":                         1,
		"PENDING_RELEASE":                 2,
		"IN_PROGRESS":                     3,
		"CANCELLING":                      4,
		"CANCELLED":                       5,
		"HALTED":                          6,
		"SUCCEEDED":                       7,
		"FAILED":                          8,
		"APPROVAL_REQUIRED":               9,
		"APPROVED":                        10,
		"REJECTED":                        11,
		"ADVANCE_REQUIRED":                12,
		"ADVANCED":                        13,
	}
)

func (x RolloutUpdateEvent_RolloutUpdateType) Enum() *RolloutUpdateEvent_RolloutUpdateType {
	p := new(RolloutUpdateEvent_RolloutUpdateType)
	*p = x
	return p
}

func (x RolloutUpdateEvent_RolloutUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RolloutUpdateEvent_RolloutUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_deploy_v1_rollout_update_payload_proto_enumTypes[0].Descriptor()
}

func (RolloutUpdateEvent_RolloutUpdateType) Type() protoreflect.EnumType {
	return &file_google_cloud_deploy_v1_rollout_update_payload_proto_enumTypes[0]
}

func (x RolloutUpdateEvent_RolloutUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RolloutUpdateEvent_RolloutUpdateType.Descriptor instead.
func (RolloutUpdateEvent_RolloutUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescGZIP(), []int{0, 0}
}

// Payload proto for "clouddeploy.googleapis.com/rollout_update"
// Platform Log event that describes the rollout update event.
type RolloutUpdateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Debug message for when a rollout update event occurs.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// Unique identifier of the pipeline.
	PipelineUid string `protobuf:"bytes,1,opt,name=pipeline_uid,json=pipelineUid,proto3" json:"pipeline_uid,omitempty"`
	// Unique identifier of the release.
	ReleaseUid string `protobuf:"bytes,2,opt,name=release_uid,json=releaseUid,proto3" json:"release_uid,omitempty"`
	// The name of the `Release`.
	Release string `protobuf:"bytes,8,opt,name=release,proto3" json:"release,omitempty"`
	// The name of the rollout.
	// rollout_uid is not in this log message because we write some of these log
	// messages at rollout creation time, before we've generated the uid.
	Rollout string `protobuf:"bytes,3,opt,name=rollout,proto3" json:"rollout,omitempty"`
	// ID of the target.
	TargetId string `protobuf:"bytes,4,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Type of this notification, e.g. for a rollout update event.
	Type Type `protobuf:"varint,7,opt,name=type,proto3,enum=google.cloud.deploy.v1.Type" json:"type,omitempty"`
	// The type of the rollout update.
	RolloutUpdateType RolloutUpdateEvent_RolloutUpdateType `protobuf:"varint,5,opt,name=rollout_update_type,json=rolloutUpdateType,proto3,enum=google.cloud.deploy.v1.RolloutUpdateEvent_RolloutUpdateType" json:"rollout_update_type,omitempty"`
}

func (x *RolloutUpdateEvent) Reset() {
	*x = RolloutUpdateEvent{}
	mi := &file_google_cloud_deploy_v1_rollout_update_payload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RolloutUpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolloutUpdateEvent) ProtoMessage() {}

func (x *RolloutUpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_deploy_v1_rollout_update_payload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolloutUpdateEvent.ProtoReflect.Descriptor instead.
func (*RolloutUpdateEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescGZIP(), []int{0}
}

func (x *RolloutUpdateEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RolloutUpdateEvent) GetPipelineUid() string {
	if x != nil {
		return x.PipelineUid
	}
	return ""
}

func (x *RolloutUpdateEvent) GetReleaseUid() string {
	if x != nil {
		return x.ReleaseUid
	}
	return ""
}

func (x *RolloutUpdateEvent) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *RolloutUpdateEvent) GetRollout() string {
	if x != nil {
		return x.Rollout
	}
	return ""
}

func (x *RolloutUpdateEvent) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *RolloutUpdateEvent) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *RolloutUpdateEvent) GetRolloutUpdateType() RolloutUpdateEvent_RolloutUpdateType {
	if x != nil {
		return x.RolloutUpdateType
	}
	return RolloutUpdateEvent_ROLLOUT_UPDATE_TYPE_UNSPECIFIED
}

var File_google_cloud_deploy_v1_rollout_update_payload_proto protoreflect.FileDescriptor

var file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x04, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x6c, 0x0a,
	0x13, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x11,
	0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x4f, 0x4c, 0x4c, 0x4f, 0x55, 0x54, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52,
	0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x41, 0x4c, 0x54,
	0x45, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45,
	0x44, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12,
	0x15, 0x0a, 0x11, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56,
	0x45, 0x44, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x44, 0x56, 0x41,
	0x4e, 0x43, 0x45, 0x44, 0x10, 0x0d, 0x42, 0x6d, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x70, 0x62, 0x3b, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescOnce sync.Once
	file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescData = file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDesc
)

func file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescGZIP() []byte {
	file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescOnce.Do(func() {
		file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescData)
	})
	return file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDescData
}

var file_google_cloud_deploy_v1_rollout_update_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_deploy_v1_rollout_update_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_deploy_v1_rollout_update_payload_proto_goTypes = []any{
	(RolloutUpdateEvent_RolloutUpdateType)(0), // 0: google.cloud.deploy.v1.RolloutUpdateEvent.RolloutUpdateType
	(*RolloutUpdateEvent)(nil),                // 1: google.cloud.deploy.v1.RolloutUpdateEvent
	(Type)(0),                                 // 2: google.cloud.deploy.v1.Type
}
var file_google_cloud_deploy_v1_rollout_update_payload_proto_depIdxs = []int32{
	2, // 0: google.cloud.deploy.v1.RolloutUpdateEvent.type:type_name -> google.cloud.deploy.v1.Type
	0, // 1: google.cloud.deploy.v1.RolloutUpdateEvent.rollout_update_type:type_name -> google.cloud.deploy.v1.RolloutUpdateEvent.RolloutUpdateType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_deploy_v1_rollout_update_payload_proto_init() }
func file_google_cloud_deploy_v1_rollout_update_payload_proto_init() {
	if File_google_cloud_deploy_v1_rollout_update_payload_proto != nil {
		return
	}
	file_google_cloud_deploy_v1_log_enums_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_deploy_v1_rollout_update_payload_proto_goTypes,
		DependencyIndexes: file_google_cloud_deploy_v1_rollout_update_payload_proto_depIdxs,
		EnumInfos:         file_google_cloud_deploy_v1_rollout_update_payload_proto_enumTypes,
		MessageInfos:      file_google_cloud_deploy_v1_rollout_update_payload_proto_msgTypes,
	}.Build()
	File_google_cloud_deploy_v1_rollout_update_payload_proto = out.File
	file_google_cloud_deploy_v1_rollout_update_payload_proto_rawDesc = nil
	file_google_cloud_deploy_v1_rollout_update_payload_proto_goTypes = nil
	file_google_cloud_deploy_v1_rollout_update_payload_proto_depIdxs = nil
}
