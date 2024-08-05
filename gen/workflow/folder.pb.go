// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: workflow/folder.proto

package workflow

import (
	_ "github.com/MyWeHub/plugin-sdk/gen/pluginrunner"
	_ "github.com/MyWeHub/plugin-sdk/gen/schema"
	workflow "github.com/MyWeHub/plugin-sdk/gen/workflow"
	_ "github.com/amsokol/mongo-go-driver-protobuf/pmongo"
	_ "github.com/amsokol/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type FolderType int32

const (
	FolderType_CLIENT  FolderType = 0
	FolderType_PROJECT FolderType = 1
	FolderType_FOLDER  FolderType = 2
)

// Enum value maps for FolderType.
var (
	FolderType_name = map[int32]string{
		0: "CLIENT",
		1: "PROJECT",
		2: "FOLDER",
	}
	FolderType_value = map[string]int32{
		"CLIENT":  0,
		"PROJECT": 1,
		"FOLDER":  2,
	}
)

func (x FolderType) Enum() *FolderType {
	p := new(FolderType)
	*p = x
	return p
}

func (x FolderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FolderType) Descriptor() protoreflect.EnumDescriptor {
	return file_workflow_folder_proto_enumTypes[0].Descriptor()
}

func (FolderType) Type() protoreflect.EnumType {
	return &file_workflow_folder_proto_enumTypes[0]
}

func (x FolderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FolderType.Descriptor instead.
func (FolderType) EnumDescriptor() ([]byte, []int) {
	return file_workflow_folder_proto_rawDescGZIP(), []int{0}
}

type RemoveFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	RemoveChildren bool   `protobuf:"varint,2,opt,name=removeChildren,proto3" json:"removeChildren,omitempty"`
}

func (x *RemoveFolderRequest) Reset() {
	*x = RemoveFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_folder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFolderRequest) ProtoMessage() {}

func (x *RemoveFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_folder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFolderRequest.ProtoReflect.Descriptor instead.
func (*RemoveFolderRequest) Descriptor() ([]byte, []int) {
	return file_workflow_folder_proto_rawDescGZIP(), []int{0}
}

func (x *RemoveFolderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveFolderRequest) GetRemoveChildren() bool {
	if x != nil {
		return x.RemoveChildren
	}
	return false
}

type MoveFolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	ParentFolderId string `protobuf:"bytes,2,opt,name=parentFolderId,proto3" json:"parentFolderId,omitempty"`
}

func (x *MoveFolderRequest) Reset() {
	*x = MoveFolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_folder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveFolderRequest) ProtoMessage() {}

func (x *MoveFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_folder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveFolderRequest.ProtoReflect.Descriptor instead.
func (*MoveFolderRequest) Descriptor() ([]byte, []int) {
	return file_workflow_folder_proto_rawDescGZIP(), []int{1}
}

func (x *MoveFolderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MoveFolderRequest) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

type ListFolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folders []*Folder `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
}

func (x *ListFolderResponse) Reset() {
	*x = ListFolderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_folder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFolderResponse) ProtoMessage() {}

func (x *ListFolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_folder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFolderResponse.ProtoReflect.Descriptor instead.
func (*ListFolderResponse) Descriptor() ([]byte, []int) {
	return file_workflow_folder_proto_rawDescGZIP(), []int{2}
}

func (x *ListFolderResponse) GetFolders() []*Folder {
	if x != nil {
		return x.Folders
	}
	return nil
}

type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Resourcetype string     `protobuf:"bytes,2,opt,name=resourcetype,proto3" json:"resourcetype,omitempty"`
	Type         FolderType `protobuf:"varint,3,opt,name=type,proto3,enum=workflow.FolderType" json:"type,omitempty"`
	Parent       string     `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Name         string     `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ClientId     string     `protobuf:"bytes,6,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_folder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_folder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_workflow_folder_proto_rawDescGZIP(), []int{3}
}

func (x *Folder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Folder) GetResourcetype() string {
	if x != nil {
		return x.Resourcetype
	}
	return ""
}

func (x *Folder) GetType() FolderType {
	if x != nil {
		return x.Type
	}
	return FolderType_CLIENT
}

func (x *Folder) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Folder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Folder) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

var File_workflow_folder_proto protoreflect.FileDescriptor

var file_workflow_folder_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03,
	0x0a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x5c, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x65, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03, 0x0a, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x07,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0x9a, 0x84, 0x9e, 0x03, 0x0a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x2a, 0x31, 0x0a, 0x0a, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x02, 0x32, 0x95, 0x03, 0x0a,
	0x0d, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x1a, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x1a, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x4d, 0x6f, 0x76,
	0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x45, 0x64, 0x69, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x10, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x18, 0x5a, 0x16, 0x77, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_folder_proto_rawDescOnce sync.Once
	file_workflow_folder_proto_rawDescData = file_workflow_folder_proto_rawDesc
)

func file_workflow_folder_proto_rawDescGZIP() []byte {
	file_workflow_folder_proto_rawDescOnce.Do(func() {
		file_workflow_folder_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_folder_proto_rawDescData)
	})
	return file_workflow_folder_proto_rawDescData
}

var file_workflow_folder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_workflow_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_workflow_folder_proto_goTypes = []interface{}{
	(FolderType)(0),             // 0: workflow.FolderType
	(*RemoveFolderRequest)(nil), // 1: workflow.RemoveFolderRequest
	(*MoveFolderRequest)(nil),   // 2: workflow.MoveFolderRequest
	(*ListFolderResponse)(nil),  // 3: workflow.ListFolderResponse
	(*Folder)(nil),              // 4: workflow.Folder
	(*workflow.IdRequest)(nil),  // 5: workflow.IdRequest
	(*workflow.Empty)(nil),      // 6: workflow.Empty
}
var file_workflow_folder_proto_depIdxs = []int32{
	4, // 0: workflow.ListFolderResponse.folders:type_name -> workflow.Folder
	0, // 1: workflow.Folder.type:type_name -> workflow.FolderType
	4, // 2: workflow.FolderService.CreateFolder:input_type -> workflow.Folder
	4, // 3: workflow.FolderService.SaveFolder:input_type -> workflow.Folder
	5, // 4: workflow.FolderService.FolderByID:input_type -> workflow.IdRequest
	1, // 5: workflow.FolderService.RemoveFolder:input_type -> workflow.RemoveFolderRequest
	2, // 6: workflow.FolderService.MoveFolder:input_type -> workflow.MoveFolderRequest
	6, // 7: workflow.FolderService.ListFolders:input_type -> workflow.Empty
	4, // 8: workflow.FolderService.EditFolder:input_type -> workflow.Folder
	4, // 9: workflow.FolderService.CreateFolder:output_type -> workflow.Folder
	4, // 10: workflow.FolderService.SaveFolder:output_type -> workflow.Folder
	4, // 11: workflow.FolderService.FolderByID:output_type -> workflow.Folder
	6, // 12: workflow.FolderService.RemoveFolder:output_type -> workflow.Empty
	6, // 13: workflow.FolderService.MoveFolder:output_type -> workflow.Empty
	3, // 14: workflow.FolderService.ListFolders:output_type -> workflow.ListFolderResponse
	6, // 15: workflow.FolderService.EditFolder:output_type -> workflow.Empty
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_workflow_folder_proto_init() }
func file_workflow_folder_proto_init() {
	if File_workflow_folder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_folder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFolderRequest); i {
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
		file_workflow_folder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveFolderRequest); i {
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
		file_workflow_folder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFolderResponse); i {
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
		file_workflow_folder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
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
			RawDescriptor: file_workflow_folder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_folder_proto_goTypes,
		DependencyIndexes: file_workflow_folder_proto_depIdxs,
		EnumInfos:         file_workflow_folder_proto_enumTypes,
		MessageInfos:      file_workflow_folder_proto_msgTypes,
	}.Build()
	File_workflow_folder_proto = out.File
	file_workflow_folder_proto_rawDesc = nil
	file_workflow_folder_proto_goTypes = nil
	file_workflow_folder_proto_depIdxs = nil
}
