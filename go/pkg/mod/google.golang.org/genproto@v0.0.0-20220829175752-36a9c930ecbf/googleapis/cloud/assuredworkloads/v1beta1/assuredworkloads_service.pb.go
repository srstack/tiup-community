// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/assuredworkloads/v1beta1/assuredworkloads_service.proto

package assuredworkloads

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto protoreflect.FileDescriptor

var file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x90, 0x0b, 0x0a, 0x17, 0x41, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xf9, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x89, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x22, 0x37, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x3a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0xda, 0x41, 0x0f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0xca, 0x41,
	0x2b, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x98, 0x01, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x17,
	0xda, 0x41, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x88, 0x02, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x55, 0x22, 0x50, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0xae, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x46, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x39, 0x2a, 0x37, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x07, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xaf, 0x01, 0x0a, 0x13, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xda, 0x41, 0x0e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x95, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x3b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x09, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x1a, 0x53, 0xca, 0x41, 0x1f, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x9d, 0x02, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x1c, 0x41, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0xaa, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x73, 0x73, 0x75, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x28, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_goTypes = []interface{}{
	(*CreateWorkloadRequest)(nil),            // 0: google.cloud.assuredworkloads.v1beta1.CreateWorkloadRequest
	(*UpdateWorkloadRequest)(nil),            // 1: google.cloud.assuredworkloads.v1beta1.UpdateWorkloadRequest
	(*RestrictAllowedResourcesRequest)(nil),  // 2: google.cloud.assuredworkloads.v1beta1.RestrictAllowedResourcesRequest
	(*DeleteWorkloadRequest)(nil),            // 3: google.cloud.assuredworkloads.v1beta1.DeleteWorkloadRequest
	(*GetWorkloadRequest)(nil),               // 4: google.cloud.assuredworkloads.v1beta1.GetWorkloadRequest
	(*AnalyzeWorkloadMoveRequest)(nil),       // 5: google.cloud.assuredworkloads.v1beta1.AnalyzeWorkloadMoveRequest
	(*ListWorkloadsRequest)(nil),             // 6: google.cloud.assuredworkloads.v1beta1.ListWorkloadsRequest
	(*longrunning.Operation)(nil),            // 7: google.longrunning.Operation
	(*Workload)(nil),                         // 8: google.cloud.assuredworkloads.v1beta1.Workload
	(*RestrictAllowedResourcesResponse)(nil), // 9: google.cloud.assuredworkloads.v1beta1.RestrictAllowedResourcesResponse
	(*emptypb.Empty)(nil),                    // 10: google.protobuf.Empty
	(*AnalyzeWorkloadMoveResponse)(nil),      // 11: google.cloud.assuredworkloads.v1beta1.AnalyzeWorkloadMoveResponse
	(*ListWorkloadsResponse)(nil),            // 12: google.cloud.assuredworkloads.v1beta1.ListWorkloadsResponse
}
var file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_depIdxs = []int32{
	0,  // 0: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.CreateWorkload:input_type -> google.cloud.assuredworkloads.v1beta1.CreateWorkloadRequest
	1,  // 1: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.UpdateWorkload:input_type -> google.cloud.assuredworkloads.v1beta1.UpdateWorkloadRequest
	2,  // 2: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.RestrictAllowedResources:input_type -> google.cloud.assuredworkloads.v1beta1.RestrictAllowedResourcesRequest
	3,  // 3: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.DeleteWorkload:input_type -> google.cloud.assuredworkloads.v1beta1.DeleteWorkloadRequest
	4,  // 4: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.GetWorkload:input_type -> google.cloud.assuredworkloads.v1beta1.GetWorkloadRequest
	5,  // 5: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.AnalyzeWorkloadMove:input_type -> google.cloud.assuredworkloads.v1beta1.AnalyzeWorkloadMoveRequest
	6,  // 6: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.ListWorkloads:input_type -> google.cloud.assuredworkloads.v1beta1.ListWorkloadsRequest
	7,  // 7: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.CreateWorkload:output_type -> google.longrunning.Operation
	8,  // 8: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.UpdateWorkload:output_type -> google.cloud.assuredworkloads.v1beta1.Workload
	9,  // 9: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.RestrictAllowedResources:output_type -> google.cloud.assuredworkloads.v1beta1.RestrictAllowedResourcesResponse
	10, // 10: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.DeleteWorkload:output_type -> google.protobuf.Empty
	8,  // 11: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.GetWorkload:output_type -> google.cloud.assuredworkloads.v1beta1.Workload
	11, // 12: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.AnalyzeWorkloadMove:output_type -> google.cloud.assuredworkloads.v1beta1.AnalyzeWorkloadMoveResponse
	12, // 13: google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService.ListWorkloads:output_type -> google.cloud.assuredworkloads.v1beta1.ListWorkloadsResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_init() }
func file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_init() {
	if File_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto != nil {
		return
	}
	file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_depIdxs,
	}.Build()
	File_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto = out.File
	file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_rawDesc = nil
	file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_goTypes = nil
	file_google_cloud_assuredworkloads_v1beta1_assuredworkloads_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssuredWorkloadsServiceClient is the client API for AssuredWorkloadsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssuredWorkloadsServiceClient interface {
	// Creates Assured Workload.
	CreateWorkload(ctx context.Context, in *CreateWorkloadRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Updates an existing workload.
	// Currently allows updating of workload display_name and labels.
	// For force updates don't set etag field in the Workload.
	// Only one update operation per workload can be in progress.
	UpdateWorkload(ctx context.Context, in *UpdateWorkloadRequest, opts ...grpc.CallOption) (*Workload, error)
	// Restrict the list of resources allowed in the Workload environment.
	// The current list of allowed products can be found at
	// https://cloud.google.com/assured-workloads/docs/supported-products
	// In addition to assuredworkloads.workload.update permission, the user should
	// also have orgpolicy.policy.set permission on the folder resource
	// to use this functionality.
	RestrictAllowedResources(ctx context.Context, in *RestrictAllowedResourcesRequest, opts ...grpc.CallOption) (*RestrictAllowedResourcesResponse, error)
	// Deletes the workload. Make sure that workload's direct children are already
	// in a deleted state, otherwise the request will fail with a
	// FAILED_PRECONDITION error.
	// In addition to assuredworkloads.workload.delete permission, the user should
	// also have orgpolicy.policy.set permission on the deleted folder to remove
	// Assured Workloads OrgPolicies.
	DeleteWorkload(ctx context.Context, in *DeleteWorkloadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Gets Assured Workload associated with a CRM Node
	GetWorkload(ctx context.Context, in *GetWorkloadRequest, opts ...grpc.CallOption) (*Workload, error)
	// A request to analyze a hypothetical move of a source project or
	// project-based workload to a target (destination) folder-based workload.
	AnalyzeWorkloadMove(ctx context.Context, in *AnalyzeWorkloadMoveRequest, opts ...grpc.CallOption) (*AnalyzeWorkloadMoveResponse, error)
	// Lists Assured Workloads under a CRM Node.
	ListWorkloads(ctx context.Context, in *ListWorkloadsRequest, opts ...grpc.CallOption) (*ListWorkloadsResponse, error)
}

type assuredWorkloadsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssuredWorkloadsServiceClient(cc grpc.ClientConnInterface) AssuredWorkloadsServiceClient {
	return &assuredWorkloadsServiceClient{cc}
}

func (c *assuredWorkloadsServiceClient) CreateWorkload(ctx context.Context, in *CreateWorkloadRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/CreateWorkload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assuredWorkloadsServiceClient) UpdateWorkload(ctx context.Context, in *UpdateWorkloadRequest, opts ...grpc.CallOption) (*Workload, error) {
	out := new(Workload)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/UpdateWorkload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assuredWorkloadsServiceClient) RestrictAllowedResources(ctx context.Context, in *RestrictAllowedResourcesRequest, opts ...grpc.CallOption) (*RestrictAllowedResourcesResponse, error) {
	out := new(RestrictAllowedResourcesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/RestrictAllowedResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assuredWorkloadsServiceClient) DeleteWorkload(ctx context.Context, in *DeleteWorkloadRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/DeleteWorkload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assuredWorkloadsServiceClient) GetWorkload(ctx context.Context, in *GetWorkloadRequest, opts ...grpc.CallOption) (*Workload, error) {
	out := new(Workload)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/GetWorkload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assuredWorkloadsServiceClient) AnalyzeWorkloadMove(ctx context.Context, in *AnalyzeWorkloadMoveRequest, opts ...grpc.CallOption) (*AnalyzeWorkloadMoveResponse, error) {
	out := new(AnalyzeWorkloadMoveResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/AnalyzeWorkloadMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assuredWorkloadsServiceClient) ListWorkloads(ctx context.Context, in *ListWorkloadsRequest, opts ...grpc.CallOption) (*ListWorkloadsResponse, error) {
	out := new(ListWorkloadsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/ListWorkloads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssuredWorkloadsServiceServer is the server API for AssuredWorkloadsService service.
type AssuredWorkloadsServiceServer interface {
	// Creates Assured Workload.
	CreateWorkload(context.Context, *CreateWorkloadRequest) (*longrunning.Operation, error)
	// Updates an existing workload.
	// Currently allows updating of workload display_name and labels.
	// For force updates don't set etag field in the Workload.
	// Only one update operation per workload can be in progress.
	UpdateWorkload(context.Context, *UpdateWorkloadRequest) (*Workload, error)
	// Restrict the list of resources allowed in the Workload environment.
	// The current list of allowed products can be found at
	// https://cloud.google.com/assured-workloads/docs/supported-products
	// In addition to assuredworkloads.workload.update permission, the user should
	// also have orgpolicy.policy.set permission on the folder resource
	// to use this functionality.
	RestrictAllowedResources(context.Context, *RestrictAllowedResourcesRequest) (*RestrictAllowedResourcesResponse, error)
	// Deletes the workload. Make sure that workload's direct children are already
	// in a deleted state, otherwise the request will fail with a
	// FAILED_PRECONDITION error.
	// In addition to assuredworkloads.workload.delete permission, the user should
	// also have orgpolicy.policy.set permission on the deleted folder to remove
	// Assured Workloads OrgPolicies.
	DeleteWorkload(context.Context, *DeleteWorkloadRequest) (*emptypb.Empty, error)
	// Gets Assured Workload associated with a CRM Node
	GetWorkload(context.Context, *GetWorkloadRequest) (*Workload, error)
	// A request to analyze a hypothetical move of a source project or
	// project-based workload to a target (destination) folder-based workload.
	AnalyzeWorkloadMove(context.Context, *AnalyzeWorkloadMoveRequest) (*AnalyzeWorkloadMoveResponse, error)
	// Lists Assured Workloads under a CRM Node.
	ListWorkloads(context.Context, *ListWorkloadsRequest) (*ListWorkloadsResponse, error)
}

// UnimplementedAssuredWorkloadsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssuredWorkloadsServiceServer struct {
}

func (*UnimplementedAssuredWorkloadsServiceServer) CreateWorkload(context.Context, *CreateWorkloadRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkload not implemented")
}
func (*UnimplementedAssuredWorkloadsServiceServer) UpdateWorkload(context.Context, *UpdateWorkloadRequest) (*Workload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkload not implemented")
}
func (*UnimplementedAssuredWorkloadsServiceServer) RestrictAllowedResources(context.Context, *RestrictAllowedResourcesRequest) (*RestrictAllowedResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestrictAllowedResources not implemented")
}
func (*UnimplementedAssuredWorkloadsServiceServer) DeleteWorkload(context.Context, *DeleteWorkloadRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkload not implemented")
}
func (*UnimplementedAssuredWorkloadsServiceServer) GetWorkload(context.Context, *GetWorkloadRequest) (*Workload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkload not implemented")
}
func (*UnimplementedAssuredWorkloadsServiceServer) AnalyzeWorkloadMove(context.Context, *AnalyzeWorkloadMoveRequest) (*AnalyzeWorkloadMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeWorkloadMove not implemented")
}
func (*UnimplementedAssuredWorkloadsServiceServer) ListWorkloads(context.Context, *ListWorkloadsRequest) (*ListWorkloadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkloads not implemented")
}

func RegisterAssuredWorkloadsServiceServer(s *grpc.Server, srv AssuredWorkloadsServiceServer) {
	s.RegisterService(&_AssuredWorkloadsService_serviceDesc, srv)
}

func _AssuredWorkloadsService_CreateWorkload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).CreateWorkload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/CreateWorkload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).CreateWorkload(ctx, req.(*CreateWorkloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssuredWorkloadsService_UpdateWorkload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).UpdateWorkload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/UpdateWorkload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).UpdateWorkload(ctx, req.(*UpdateWorkloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssuredWorkloadsService_RestrictAllowedResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestrictAllowedResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).RestrictAllowedResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/RestrictAllowedResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).RestrictAllowedResources(ctx, req.(*RestrictAllowedResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssuredWorkloadsService_DeleteWorkload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).DeleteWorkload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/DeleteWorkload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).DeleteWorkload(ctx, req.(*DeleteWorkloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssuredWorkloadsService_GetWorkload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).GetWorkload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/GetWorkload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).GetWorkload(ctx, req.(*GetWorkloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssuredWorkloadsService_AnalyzeWorkloadMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeWorkloadMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).AnalyzeWorkloadMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/AnalyzeWorkloadMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).AnalyzeWorkloadMove(ctx, req.(*AnalyzeWorkloadMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssuredWorkloadsService_ListWorkloads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkloadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssuredWorkloadsServiceServer).ListWorkloads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService/ListWorkloads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssuredWorkloadsServiceServer).ListWorkloads(ctx, req.(*ListWorkloadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssuredWorkloadsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.assuredworkloads.v1beta1.AssuredWorkloadsService",
	HandlerType: (*AssuredWorkloadsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkload",
			Handler:    _AssuredWorkloadsService_CreateWorkload_Handler,
		},
		{
			MethodName: "UpdateWorkload",
			Handler:    _AssuredWorkloadsService_UpdateWorkload_Handler,
		},
		{
			MethodName: "RestrictAllowedResources",
			Handler:    _AssuredWorkloadsService_RestrictAllowedResources_Handler,
		},
		{
			MethodName: "DeleteWorkload",
			Handler:    _AssuredWorkloadsService_DeleteWorkload_Handler,
		},
		{
			MethodName: "GetWorkload",
			Handler:    _AssuredWorkloadsService_GetWorkload_Handler,
		},
		{
			MethodName: "AnalyzeWorkloadMove",
			Handler:    _AssuredWorkloadsService_AnalyzeWorkloadMove_Handler,
		},
		{
			MethodName: "ListWorkloads",
			Handler:    _AssuredWorkloadsService_ListWorkloads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/assuredworkloads/v1beta1/assuredworkloads_service.proto",
}
