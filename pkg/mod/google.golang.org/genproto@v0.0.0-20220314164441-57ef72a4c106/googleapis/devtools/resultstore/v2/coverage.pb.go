// Copyright 2021 Google LLC
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
// source: google/devtools/resultstore/v2/coverage.proto

package resultstore

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

// Describes line coverage for a file
type LineCoverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Which source lines in the file represent the start of a statement that was
	// instrumented to detect whether it was executed by the test.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// A 1 denotes the line was instrumented.
	// A 0 denotes the line was not instrumented.
	InstrumentedLines []byte `protobuf:"bytes,1,opt,name=instrumented_lines,json=instrumentedLines,proto3" json:"instrumented_lines,omitempty"`
	// Which of the instrumented source lines were executed by the test. Should
	// include lines that were not instrumented.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// A 1 denotes the line was executed.
	// A 0 denotes the line was not executed.
	ExecutedLines []byte `protobuf:"bytes,2,opt,name=executed_lines,json=executedLines,proto3" json:"executed_lines,omitempty"`
}

func (x *LineCoverage) Reset() {
	*x = LineCoverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineCoverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineCoverage) ProtoMessage() {}

func (x *LineCoverage) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineCoverage.ProtoReflect.Descriptor instead.
func (*LineCoverage) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_coverage_proto_rawDescGZIP(), []int{0}
}

func (x *LineCoverage) GetInstrumentedLines() []byte {
	if x != nil {
		return x.InstrumentedLines
	}
	return nil
}

func (x *LineCoverage) GetExecutedLines() []byte {
	if x != nil {
		return x.ExecutedLines
	}
	return nil
}

// Describes branch coverage for a file
type BranchCoverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field branch_present denotes the lines containing at least one branch.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// A 1 denotes the line contains at least one branch.
	// A 0 denotes the line contains no branches.
	BranchPresent []byte `protobuf:"bytes,1,opt,name=branch_present,json=branchPresent,proto3" json:"branch_present,omitempty"`
	// Contains the number of branches present, only for the lines which have the
	// corresponding bit set in branch_present, in a relative order ignoring
	// lines which do not have any branches.
	BranchesInLine []int32 `protobuf:"varint,2,rep,packed,name=branches_in_line,json=branchesInLine,proto3" json:"branches_in_line,omitempty"`
	// As each branch can have any one of the following three states: not
	// executed, executed but not taken, executed and taken.
	//
	// This is a bitfield where i-th bit corresponds to the i-th branch. Divide
	// branch number by 8 to get index into byte array. Mod branch number by 8 to
	// get bit number (0 = LSB, 7 = MSB).
	//
	// i-th bit of the following two byte arrays are used to denote the above
	// mentioned states.
	//
	// not executed: i-th bit of executed == 0 && i-th bit of taken == 0
	// executed but not taken: i-th bit of executed == 1 && i-th bit of taken == 0
	// executed and taken: i-th bit of executed == 1 && i-th bit of taken == 1
	Executed []byte `protobuf:"bytes,3,opt,name=executed,proto3" json:"executed,omitempty"`
	// Described above.
	Taken []byte `protobuf:"bytes,4,opt,name=taken,proto3" json:"taken,omitempty"`
}

func (x *BranchCoverage) Reset() {
	*x = BranchCoverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchCoverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchCoverage) ProtoMessage() {}

func (x *BranchCoverage) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchCoverage.ProtoReflect.Descriptor instead.
func (*BranchCoverage) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_coverage_proto_rawDescGZIP(), []int{1}
}

func (x *BranchCoverage) GetBranchPresent() []byte {
	if x != nil {
		return x.BranchPresent
	}
	return nil
}

func (x *BranchCoverage) GetBranchesInLine() []int32 {
	if x != nil {
		return x.BranchesInLine
	}
	return nil
}

func (x *BranchCoverage) GetExecuted() []byte {
	if x != nil {
		return x.Executed
	}
	return nil
}

func (x *BranchCoverage) GetTaken() []byte {
	if x != nil {
		return x.Taken
	}
	return nil
}

// Describes code coverage for a particular file under test.
type FileCoverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path of source file within the SourceContext of this Invocation.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Details of lines in a file for calculating line coverage.
	LineCoverage *LineCoverage `protobuf:"bytes,2,opt,name=line_coverage,json=lineCoverage,proto3" json:"line_coverage,omitempty"`
	// Details of branches in a file for calculating branch coverage.
	BranchCoverage *BranchCoverage `protobuf:"bytes,3,opt,name=branch_coverage,json=branchCoverage,proto3" json:"branch_coverage,omitempty"`
}

func (x *FileCoverage) Reset() {
	*x = FileCoverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileCoverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileCoverage) ProtoMessage() {}

func (x *FileCoverage) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileCoverage.ProtoReflect.Descriptor instead.
func (*FileCoverage) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_coverage_proto_rawDescGZIP(), []int{2}
}

func (x *FileCoverage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileCoverage) GetLineCoverage() *LineCoverage {
	if x != nil {
		return x.LineCoverage
	}
	return nil
}

func (x *FileCoverage) GetBranchCoverage() *BranchCoverage {
	if x != nil {
		return x.BranchCoverage
	}
	return nil
}

// Describes code coverage for a build or test Action. This is used to store
// baseline coverage for build Actions and test coverage for test Actions.
type ActionCoverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of coverage info for all source files that the TestResult covers.
	FileCoverages []*FileCoverage `protobuf:"bytes,2,rep,name=file_coverages,json=fileCoverages,proto3" json:"file_coverages,omitempty"`
}

func (x *ActionCoverage) Reset() {
	*x = ActionCoverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionCoverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionCoverage) ProtoMessage() {}

func (x *ActionCoverage) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionCoverage.ProtoReflect.Descriptor instead.
func (*ActionCoverage) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_coverage_proto_rawDescGZIP(), []int{3}
}

func (x *ActionCoverage) GetFileCoverages() []*FileCoverage {
	if x != nil {
		return x.FileCoverages
	}
	return nil
}

// Describes aggregate code coverage for a collection of build or test Actions.
// A line or branch is covered if and only if it is covered in any of the build
// or test actions.
type AggregateCoverage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregated coverage info for all source files that the actions cover.
	FileCoverages []*FileCoverage `protobuf:"bytes,1,rep,name=file_coverages,json=fileCoverages,proto3" json:"file_coverages,omitempty"`
}

func (x *AggregateCoverage) Reset() {
	*x = AggregateCoverage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateCoverage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateCoverage) ProtoMessage() {}

func (x *AggregateCoverage) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_coverage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateCoverage.ProtoReflect.Descriptor instead.
func (*AggregateCoverage) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_coverage_proto_rawDescGZIP(), []int{4}
}

func (x *AggregateCoverage) GetFileCoverages() []*FileCoverage {
	if x != nil {
		return x.FileCoverages
	}
	return nil
}

var File_google_devtools_resultstore_v2_coverage_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_coverage_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x22,
	0x64, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x5f,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64,
	0x4c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x0e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x10, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x73, 0x49, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x22, 0xce, 0x01, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x51, 0x0a, 0x0d, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x65, 0x0a, 0x0e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x53,
	0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x68, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0d,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x42, 0x80, 0x01,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_coverage_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_coverage_proto_rawDescData = file_google_devtools_resultstore_v2_coverage_proto_rawDesc
)

func file_google_devtools_resultstore_v2_coverage_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_coverage_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_coverage_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_coverage_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_coverage_proto_rawDescData
}

var file_google_devtools_resultstore_v2_coverage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_devtools_resultstore_v2_coverage_proto_goTypes = []interface{}{
	(*LineCoverage)(nil),      // 0: google.devtools.resultstore.v2.LineCoverage
	(*BranchCoverage)(nil),    // 1: google.devtools.resultstore.v2.BranchCoverage
	(*FileCoverage)(nil),      // 2: google.devtools.resultstore.v2.FileCoverage
	(*ActionCoverage)(nil),    // 3: google.devtools.resultstore.v2.ActionCoverage
	(*AggregateCoverage)(nil), // 4: google.devtools.resultstore.v2.AggregateCoverage
}
var file_google_devtools_resultstore_v2_coverage_proto_depIdxs = []int32{
	0, // 0: google.devtools.resultstore.v2.FileCoverage.line_coverage:type_name -> google.devtools.resultstore.v2.LineCoverage
	1, // 1: google.devtools.resultstore.v2.FileCoverage.branch_coverage:type_name -> google.devtools.resultstore.v2.BranchCoverage
	2, // 2: google.devtools.resultstore.v2.ActionCoverage.file_coverages:type_name -> google.devtools.resultstore.v2.FileCoverage
	2, // 3: google.devtools.resultstore.v2.AggregateCoverage.file_coverages:type_name -> google.devtools.resultstore.v2.FileCoverage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_coverage_proto_init() }
func file_google_devtools_resultstore_v2_coverage_proto_init() {
	if File_google_devtools_resultstore_v2_coverage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_coverage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineCoverage); i {
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
		file_google_devtools_resultstore_v2_coverage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchCoverage); i {
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
		file_google_devtools_resultstore_v2_coverage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileCoverage); i {
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
		file_google_devtools_resultstore_v2_coverage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionCoverage); i {
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
		file_google_devtools_resultstore_v2_coverage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateCoverage); i {
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
			RawDescriptor: file_google_devtools_resultstore_v2_coverage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_coverage_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_coverage_proto_depIdxs,
		MessageInfos:      file_google_devtools_resultstore_v2_coverage_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_coverage_proto = out.File
	file_google_devtools_resultstore_v2_coverage_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_coverage_proto_goTypes = nil
	file_google_devtools_resultstore_v2_coverage_proto_depIdxs = nil
}
