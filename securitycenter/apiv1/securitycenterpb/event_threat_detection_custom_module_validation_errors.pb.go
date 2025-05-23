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
// source: google/cloud/securitycenter/v1/event_threat_detection_custom_module_validation_errors.proto

package securitycenterpb

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

// A list of zero or more errors encountered while validating the uploaded
// configuration of an Event Threat Detection Custom Module.
type CustomModuleValidationErrors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []*CustomModuleValidationError `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *CustomModuleValidationErrors) Reset() {
	*x = CustomModuleValidationErrors{}
	mi := &file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomModuleValidationErrors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomModuleValidationErrors) ProtoMessage() {}

func (x *CustomModuleValidationErrors) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomModuleValidationErrors.ProtoReflect.Descriptor instead.
func (*CustomModuleValidationErrors) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescGZIP(), []int{0}
}

func (x *CustomModuleValidationErrors) GetErrors() []*CustomModuleValidationError {
	if x != nil {
		return x.Errors
	}
	return nil
}

// An error encountered while validating the uploaded configuration of an
// Event Threat Detection Custom Module.
type CustomModuleValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A description of the error, suitable for human consumption. Required.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// The path, in RFC 8901 JSON Pointer format, to the field that failed
	// validation. This may be left empty if no specific field is affected.
	FieldPath string `protobuf:"bytes,2,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// The initial position of the error in the uploaded text version of the
	// module. This field may be omitted if no specific position applies, or if
	// one could not be computed.
	Start *Position `protobuf:"bytes,3,opt,name=start,proto3,oneof" json:"start,omitempty"`
	// The end position of the error in the uploaded text version of the
	// module. This field may be omitted if no specific position applies, or if
	// one could not be computed..
	End *Position `protobuf:"bytes,4,opt,name=end,proto3,oneof" json:"end,omitempty"`
}

func (x *CustomModuleValidationError) Reset() {
	*x = CustomModuleValidationError{}
	mi := &file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomModuleValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomModuleValidationError) ProtoMessage() {}

func (x *CustomModuleValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomModuleValidationError.ProtoReflect.Descriptor instead.
func (*CustomModuleValidationError) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescGZIP(), []int{1}
}

func (x *CustomModuleValidationError) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomModuleValidationError) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (x *CustomModuleValidationError) GetStart() *Position {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *CustomModuleValidationError) GetEnd() *Position {
	if x != nil {
		return x.End
	}
	return nil
}

// A position in the uploaded text version of a module.
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineNumber   int32 `protobuf:"varint,1,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	ColumnNumber int32 `protobuf:"varint,2,opt,name=column_number,json=columnNumber,proto3" json:"column_number,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	mi := &file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescGZIP(), []int{2}
}

func (x *Position) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *Position) GetColumnNumber() int32 {
	if x != nil {
		return x.ColumnNumber
	}
	return 0
}

var File_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDesc = []byte{
	0x0a, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x73, 0x0a,
	0x1c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x53, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x01, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x65, 0x6e, 0x64, 0x22, 0x50, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x8f, 0x02,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x35, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescData = file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_goTypes = []any{
	(*CustomModuleValidationErrors)(nil), // 0: google.cloud.securitycenter.v1.CustomModuleValidationErrors
	(*CustomModuleValidationError)(nil),  // 1: google.cloud.securitycenter.v1.CustomModuleValidationError
	(*Position)(nil),                     // 2: google.cloud.securitycenter.v1.Position
}
var file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.CustomModuleValidationErrors.errors:type_name -> google.cloud.securitycenter.v1.CustomModuleValidationError
	2, // 1: google.cloud.securitycenter.v1.CustomModuleValidationError.start:type_name -> google.cloud.securitycenter.v1.Position
	2, // 2: google.cloud.securitycenter.v1.CustomModuleValidationError.end:type_name -> google.cloud.securitycenter.v1.Position
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_init()
}
func file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_init() {
	if File_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto != nil {
		return
	}
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto = out.File
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_event_threat_detection_custom_module_validation_errors_proto_depIdxs = nil
}
