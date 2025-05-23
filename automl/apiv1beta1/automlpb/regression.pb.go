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
// source: google/cloud/automl/v1beta1/regression.proto

package automlpb

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

// Metrics for regression problems.
type RegressionEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Root Mean Squared Error (RMSE).
	RootMeanSquaredError float32 `protobuf:"fixed32,1,opt,name=root_mean_squared_error,json=rootMeanSquaredError,proto3" json:"root_mean_squared_error,omitempty"`
	// Output only. Mean Absolute Error (MAE).
	MeanAbsoluteError float32 `protobuf:"fixed32,2,opt,name=mean_absolute_error,json=meanAbsoluteError,proto3" json:"mean_absolute_error,omitempty"`
	// Output only. Mean absolute percentage error. Only set if all ground truth
	// values are are positive.
	MeanAbsolutePercentageError float32 `protobuf:"fixed32,3,opt,name=mean_absolute_percentage_error,json=meanAbsolutePercentageError,proto3" json:"mean_absolute_percentage_error,omitempty"`
	// Output only. R squared.
	RSquared float32 `protobuf:"fixed32,4,opt,name=r_squared,json=rSquared,proto3" json:"r_squared,omitempty"`
	// Output only. Root mean squared log error.
	RootMeanSquaredLogError float32 `protobuf:"fixed32,5,opt,name=root_mean_squared_log_error,json=rootMeanSquaredLogError,proto3" json:"root_mean_squared_log_error,omitempty"`
}

func (x *RegressionEvaluationMetrics) Reset() {
	*x = RegressionEvaluationMetrics{}
	mi := &file_google_cloud_automl_v1beta1_regression_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegressionEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegressionEvaluationMetrics) ProtoMessage() {}

func (x *RegressionEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_regression_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegressionEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*RegressionEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_regression_proto_rawDescGZIP(), []int{0}
}

func (x *RegressionEvaluationMetrics) GetRootMeanSquaredError() float32 {
	if x != nil {
		return x.RootMeanSquaredError
	}
	return 0
}

func (x *RegressionEvaluationMetrics) GetMeanAbsoluteError() float32 {
	if x != nil {
		return x.MeanAbsoluteError
	}
	return 0
}

func (x *RegressionEvaluationMetrics) GetMeanAbsolutePercentageError() float32 {
	if x != nil {
		return x.MeanAbsolutePercentageError
	}
	return 0
}

func (x *RegressionEvaluationMetrics) GetRSquared() float32 {
	if x != nil {
		return x.RSquared
	}
	return 0
}

func (x *RegressionEvaluationMetrics) GetRootMeanSquaredLogError() float32 {
	if x != nil {
		return x.RootMeanSquaredLogError
	}
	return 0
}

var File_google_cloud_automl_v1beta1_regression_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_regression_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0xa4, 0x02, 0x0a, 0x1b,
	0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x72,
	0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x72, 0x6f,
	0x6f, 0x74, 0x4d, 0x65, 0x61, 0x6e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x11, 0x6d, 0x65, 0x61, 0x6e, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x43, 0x0a, 0x1e, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1b, 0x6d, 0x65, 0x61, 0x6e,
	0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x5f, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x72, 0x53, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x61,
	0x6e, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x17, 0x72, 0x6f, 0x6f, 0x74, 0x4d,
	0x65, 0x61, 0x6e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0xaa, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62,
	0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_regression_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_regression_proto_rawDescData = file_google_cloud_automl_v1beta1_regression_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_regression_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_regression_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_regression_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_regression_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_regression_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_regression_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_automl_v1beta1_regression_proto_goTypes = []any{
	(*RegressionEvaluationMetrics)(nil), // 0: google.cloud.automl.v1beta1.RegressionEvaluationMetrics
}
var file_google_cloud_automl_v1beta1_regression_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_regression_proto_init() }
func file_google_cloud_automl_v1beta1_regression_proto_init() {
	if File_google_cloud_automl_v1beta1_regression_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1beta1_regression_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_regression_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_regression_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_regression_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_regression_proto = out.File
	file_google_cloud_automl_v1beta1_regression_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_regression_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_regression_proto_depIdxs = nil
}
