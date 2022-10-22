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
// source: google/cloud/securitycenter/v1/indicator.proto

package securitycenter

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents what's commonly known as an Indicator of compromise (IoC) in
// computer forensics. This is an artifact observed on a network or in an
// operating system that, with high confidence, indicates a computer intrusion.
// Reference: https://en.wikipedia.org/wiki/Indicator_of_compromise
type Indicator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of ip addresses associated to the Finding.
	IpAddresses []string `protobuf:"bytes,1,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	// List of domains associated to the Finding.
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	// The list of matched signatures indicating that the given
	// process is present in the environment.
	Signatures []*Indicator_ProcessSignature `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// The list of URIs associated to the Findings.
	Uris []string `protobuf:"bytes,4,rep,name=uris,proto3" json:"uris,omitempty"`
}

func (x *Indicator) Reset() {
	*x = Indicator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Indicator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Indicator) ProtoMessage() {}

func (x *Indicator) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Indicator.ProtoReflect.Descriptor instead.
func (*Indicator) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_indicator_proto_rawDescGZIP(), []int{0}
}

func (x *Indicator) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *Indicator) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *Indicator) GetSignatures() []*Indicator_ProcessSignature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

func (x *Indicator) GetUris() []string {
	if x != nil {
		return x.Uris
	}
	return nil
}

// Indicates what signature matched this process.
type Indicator_ProcessSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Signature:
	//
	//	*Indicator_ProcessSignature_MemoryHashSignature_
	//	*Indicator_ProcessSignature_YaraRuleSignature_
	Signature isIndicator_ProcessSignature_Signature `protobuf_oneof:"signature"`
}

func (x *Indicator_ProcessSignature) Reset() {
	*x = Indicator_ProcessSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Indicator_ProcessSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Indicator_ProcessSignature) ProtoMessage() {}

func (x *Indicator_ProcessSignature) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Indicator_ProcessSignature.ProtoReflect.Descriptor instead.
func (*Indicator_ProcessSignature) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_indicator_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Indicator_ProcessSignature) GetSignature() isIndicator_ProcessSignature_Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (x *Indicator_ProcessSignature) GetMemoryHashSignature() *Indicator_ProcessSignature_MemoryHashSignature {
	if x, ok := x.GetSignature().(*Indicator_ProcessSignature_MemoryHashSignature_); ok {
		return x.MemoryHashSignature
	}
	return nil
}

func (x *Indicator_ProcessSignature) GetYaraRuleSignature() *Indicator_ProcessSignature_YaraRuleSignature {
	if x, ok := x.GetSignature().(*Indicator_ProcessSignature_YaraRuleSignature_); ok {
		return x.YaraRuleSignature
	}
	return nil
}

type isIndicator_ProcessSignature_Signature interface {
	isIndicator_ProcessSignature_Signature()
}

type Indicator_ProcessSignature_MemoryHashSignature_ struct {
	// Signature indicating that a binary family was matched.
	MemoryHashSignature *Indicator_ProcessSignature_MemoryHashSignature `protobuf:"bytes,6,opt,name=memory_hash_signature,json=memoryHashSignature,proto3,oneof"`
}

type Indicator_ProcessSignature_YaraRuleSignature_ struct {
	// Signature indicating that a YARA rule was matched.
	YaraRuleSignature *Indicator_ProcessSignature_YaraRuleSignature `protobuf:"bytes,7,opt,name=yara_rule_signature,json=yaraRuleSignature,proto3,oneof"`
}

func (*Indicator_ProcessSignature_MemoryHashSignature_) isIndicator_ProcessSignature_Signature() {}

func (*Indicator_ProcessSignature_YaraRuleSignature_) isIndicator_ProcessSignature_Signature() {}

// A signature corresponding to memory page hashes.
type Indicator_ProcessSignature_MemoryHashSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The binary family.
	BinaryFamily string `protobuf:"bytes,1,opt,name=binary_family,json=binaryFamily,proto3" json:"binary_family,omitempty"`
	// The list of memory hash detections contributing to the binary family
	// match.
	Detections []*Indicator_ProcessSignature_MemoryHashSignature_Detection `protobuf:"bytes,4,rep,name=detections,proto3" json:"detections,omitempty"`
}

func (x *Indicator_ProcessSignature_MemoryHashSignature) Reset() {
	*x = Indicator_ProcessSignature_MemoryHashSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Indicator_ProcessSignature_MemoryHashSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Indicator_ProcessSignature_MemoryHashSignature) ProtoMessage() {}

func (x *Indicator_ProcessSignature_MemoryHashSignature) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Indicator_ProcessSignature_MemoryHashSignature.ProtoReflect.Descriptor instead.
func (*Indicator_ProcessSignature_MemoryHashSignature) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_indicator_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Indicator_ProcessSignature_MemoryHashSignature) GetBinaryFamily() string {
	if x != nil {
		return x.BinaryFamily
	}
	return ""
}

func (x *Indicator_ProcessSignature_MemoryHashSignature) GetDetections() []*Indicator_ProcessSignature_MemoryHashSignature_Detection {
	if x != nil {
		return x.Detections
	}
	return nil
}

// A signature corresponding to a YARA rule.
type Indicator_ProcessSignature_YaraRuleSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the YARA rule.
	YaraRule string `protobuf:"bytes,5,opt,name=yara_rule,json=yaraRule,proto3" json:"yara_rule,omitempty"`
}

func (x *Indicator_ProcessSignature_YaraRuleSignature) Reset() {
	*x = Indicator_ProcessSignature_YaraRuleSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Indicator_ProcessSignature_YaraRuleSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Indicator_ProcessSignature_YaraRuleSignature) ProtoMessage() {}

func (x *Indicator_ProcessSignature_YaraRuleSignature) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Indicator_ProcessSignature_YaraRuleSignature.ProtoReflect.Descriptor instead.
func (*Indicator_ProcessSignature_YaraRuleSignature) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_indicator_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Indicator_ProcessSignature_YaraRuleSignature) GetYaraRule() string {
	if x != nil {
		return x.YaraRule
	}
	return ""
}

// Memory hash detection contributing to the binary family match.
type Indicator_ProcessSignature_MemoryHashSignature_Detection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the binary associated with the memory hash
	// signature detection.
	Binary string `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	// The percentage of memory page hashes in the signature
	// that were matched.
	PercentPagesMatched float64 `protobuf:"fixed64,3,opt,name=percent_pages_matched,json=percentPagesMatched,proto3" json:"percent_pages_matched,omitempty"`
}

func (x *Indicator_ProcessSignature_MemoryHashSignature_Detection) Reset() {
	*x = Indicator_ProcessSignature_MemoryHashSignature_Detection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Indicator_ProcessSignature_MemoryHashSignature_Detection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Indicator_ProcessSignature_MemoryHashSignature_Detection) ProtoMessage() {}

func (x *Indicator_ProcessSignature_MemoryHashSignature_Detection) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Indicator_ProcessSignature_MemoryHashSignature_Detection.ProtoReflect.Descriptor instead.
func (*Indicator_ProcessSignature_MemoryHashSignature_Detection) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_indicator_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *Indicator_ProcessSignature_MemoryHashSignature_Detection) GetBinary() string {
	if x != nil {
		return x.Binary
	}
	return ""
}

func (x *Indicator_ProcessSignature_MemoryHashSignature_Detection) GetPercentPagesMatched() float64 {
	if x != nil {
		return x.PercentPagesMatched
	}
	return 0
}

var File_google_cloud_securitycenter_v1_indicator_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_indicator_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0xa3, 0x06, 0x0a, 0x09, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x5a, 0x0a, 0x0a, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x69, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x69, 0x73, 0x1a, 0xe8, 0x04, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x84, 0x01, 0x0a, 0x15, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x48, 0x61, 0x73, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x13, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x48, 0x61, 0x73, 0x68, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x7e, 0x0a, 0x13, 0x79, 0x61, 0x72, 0x61, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x59, 0x61, 0x72, 0x61, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x48, 0x00, 0x52, 0x11, 0x79, 0x61, 0x72, 0x61, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x8d, 0x02, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x48, 0x61, 0x73, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x12, 0x78, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x48, 0x61, 0x73, 0x68, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x57,
	0x0a, 0x09, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x13, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x1a, 0x30, 0x0a, 0x11, 0x59, 0x61, 0x72, 0x61, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x79, 0x61, 0x72, 0x61, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x79, 0x61, 0x72, 0x61, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_indicator_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_indicator_proto_rawDescData = file_google_cloud_securitycenter_v1_indicator_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_indicator_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_indicator_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_indicator_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_indicator_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_indicator_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_indicator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_securitycenter_v1_indicator_proto_goTypes = []interface{}{
	(*Indicator)(nil),                                                // 0: google.cloud.securitycenter.v1.Indicator
	(*Indicator_ProcessSignature)(nil),                               // 1: google.cloud.securitycenter.v1.Indicator.ProcessSignature
	(*Indicator_ProcessSignature_MemoryHashSignature)(nil),           // 2: google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature
	(*Indicator_ProcessSignature_YaraRuleSignature)(nil),             // 3: google.cloud.securitycenter.v1.Indicator.ProcessSignature.YaraRuleSignature
	(*Indicator_ProcessSignature_MemoryHashSignature_Detection)(nil), // 4: google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.Detection
}
var file_google_cloud_securitycenter_v1_indicator_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.Indicator.signatures:type_name -> google.cloud.securitycenter.v1.Indicator.ProcessSignature
	2, // 1: google.cloud.securitycenter.v1.Indicator.ProcessSignature.memory_hash_signature:type_name -> google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature
	3, // 2: google.cloud.securitycenter.v1.Indicator.ProcessSignature.yara_rule_signature:type_name -> google.cloud.securitycenter.v1.Indicator.ProcessSignature.YaraRuleSignature
	4, // 3: google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.detections:type_name -> google.cloud.securitycenter.v1.Indicator.ProcessSignature.MemoryHashSignature.Detection
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_indicator_proto_init() }
func file_google_cloud_securitycenter_v1_indicator_proto_init() {
	if File_google_cloud_securitycenter_v1_indicator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Indicator); i {
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
		file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Indicator_ProcessSignature); i {
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
		file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Indicator_ProcessSignature_MemoryHashSignature); i {
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
		file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Indicator_ProcessSignature_YaraRuleSignature); i {
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
		file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Indicator_ProcessSignature_MemoryHashSignature_Detection); i {
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
	file_google_cloud_securitycenter_v1_indicator_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Indicator_ProcessSignature_MemoryHashSignature_)(nil),
		(*Indicator_ProcessSignature_YaraRuleSignature_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v1_indicator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_indicator_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_indicator_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_indicator_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_indicator_proto = out.File
	file_google_cloud_securitycenter_v1_indicator_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_indicator_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_indicator_proto_depIdxs = nil
}
