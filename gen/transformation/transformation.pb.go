// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: transformation/transformation.proto

package transformation

import (
	_ "github.com/MyWeHub/plugin-sdk/gen/graph"
	_ "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	_ "github.com/MyWeHub/plugin-sdk/gen/schema"
	testRunner "github.com/MyWeHub/plugin-sdk/gen/testRunner"
	_ "github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	_ "github.com/amsokol/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkflowOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytedata []byte          `protobuf:"bytes,1,opt,name=bytedata,proto3" json:"bytedata,omitempty"`
	Output   *structpb.Value `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *WorkflowOutput) Reset() {
	*x = WorkflowOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformation_transformation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowOutput) ProtoMessage() {}

func (x *WorkflowOutput) ProtoReflect() protoreflect.Message {
	mi := &file_transformation_transformation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowOutput.ProtoReflect.Descriptor instead.
func (*WorkflowOutput) Descriptor() ([]byte, []int) {
	return file_transformation_transformation_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowOutput) GetBytedata() []byte {
	if x != nil {
		return x.Bytedata
	}
	return nil
}

func (x *WorkflowOutput) GetOutput() *structpb.Value {
	if x != nil {
		return x.Output
	}
	return nil
}

type ErrorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string                 `protobuf:"bytes,1,opt,name=WorkflowId,proto3" json:"WorkflowId,omitempty"`
	ClientId   string                 `protobuf:"bytes,2,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	Time       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
	PluginName string                 `protobuf:"bytes,4,opt,name=PluginName,proto3" json:"PluginName,omitempty"`
	Error      string                 `protobuf:"bytes,5,opt,name=Error,proto3" json:"Error,omitempty"`
	Data       *structpb.Struct       `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`
	NodeId     string                 `protobuf:"bytes,7,opt,name=NodeId,proto3" json:"NodeId,omitempty"`
}

func (x *ErrorData) Reset() {
	*x = ErrorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformation_transformation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorData) ProtoMessage() {}

func (x *ErrorData) ProtoReflect() protoreflect.Message {
	mi := &file_transformation_transformation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorData.ProtoReflect.Descriptor instead.
func (*ErrorData) Descriptor() ([]byte, []int) {
	return file_transformation_transformation_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorData) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *ErrorData) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ErrorData) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ErrorData) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *ErrorData) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ErrorData) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ErrorData) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

var File_transformation_transformation_proto protoreflect.FileDescriptor

var file_transformation_transformation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74, 0x61,
	0x67, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x70, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5c, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e,
	0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xf2,
	0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x32, 0xa9, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x0b, 0x52, 0x75, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x75, 0x6e, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x4b, 0x0a, 0x0f, 0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42,
	0x1e, 0x5a, 0x1c, 0x77, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transformation_transformation_proto_rawDescOnce sync.Once
	file_transformation_transformation_proto_rawDescData = file_transformation_transformation_proto_rawDesc
)

func file_transformation_transformation_proto_rawDescGZIP() []byte {
	file_transformation_transformation_proto_rawDescOnce.Do(func() {
		file_transformation_transformation_proto_rawDescData = protoimpl.X.CompressGZIP(file_transformation_transformation_proto_rawDescData)
	})
	return file_transformation_transformation_proto_rawDescData
}

var file_transformation_transformation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transformation_transformation_proto_goTypes = []interface{}{
	(*WorkflowOutput)(nil),          // 0: transformation.WorkflowOutput
	(*ErrorData)(nil),               // 1: transformation.ErrorData
	(*structpb.Value)(nil),          // 2: google.protobuf.Value
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
	(*structpb.Struct)(nil),         // 4: google.protobuf.Struct
	(*testRunner.StartRun)(nil),     // 5: testRunner.StartRun
	(*testRunner.StartRunTest)(nil), // 6: testRunner.StartRunTest
}
var file_transformation_transformation_proto_depIdxs = []int32{
	2, // 0: transformation.WorkflowOutput.output:type_name -> google.protobuf.Value
	3, // 1: transformation.ErrorData.Time:type_name -> google.protobuf.Timestamp
	4, // 2: transformation.ErrorData.Data:type_name -> google.protobuf.Struct
	5, // 3: transformation.TransformationService.RunWorkflow:input_type -> testRunner.StartRun
	6, // 4: transformation.TransformationService.RunTestWorkflow:input_type -> testRunner.StartRunTest
	0, // 5: transformation.TransformationService.RunWorkflow:output_type -> transformation.WorkflowOutput
	0, // 6: transformation.TransformationService.RunTestWorkflow:output_type -> transformation.WorkflowOutput
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transformation_transformation_proto_init() }
func file_transformation_transformation_proto_init() {
	if File_transformation_transformation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transformation_transformation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowOutput); i {
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
		file_transformation_transformation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorData); i {
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
			RawDescriptor: file_transformation_transformation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transformation_transformation_proto_goTypes,
		DependencyIndexes: file_transformation_transformation_proto_depIdxs,
		MessageInfos:      file_transformation_transformation_proto_msgTypes,
	}.Build()
	File_transformation_transformation_proto = out.File
	file_transformation_transformation_proto_rawDesc = nil
	file_transformation_transformation_proto_goTypes = nil
	file_transformation_transformation_proto_depIdxs = nil
}
