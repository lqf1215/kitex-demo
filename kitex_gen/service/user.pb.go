// Copyright 2021 CloudWeGo Authors
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: idl/user.proto

package service

import (
	context "context"
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

type GetUserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`         //页码
	PageSize int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"` //每页大小
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetUserListReq) Reset() {
	*x = GetUserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListReq) ProtoMessage() {}

func (x *GetUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListReq.ProtoReflect.Descriptor instead.
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return file_idl_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetUserListReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetUserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UserData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserListResp) Reset() {
	*x = GetUserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResp) ProtoMessage() {}

func (x *GetUserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResp.ProtoReflect.Descriptor instead.
func (*GetUserListResp) Descriptor() ([]byte, []int) {
	return file_idl_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserListResp) GetData() []*UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Sex      string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_idl_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_idl_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

type GetUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIdReq) Reset() {
	*x = GetUserByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdReq) ProtoMessage() {}

func (x *GetUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdReq.ProtoReflect.Descriptor instead.
func (*GetUserByIdReq) Descriptor() ([]byte, []int) {
	return file_idl_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_idl_user_proto protoreflect.FileDescriptor

var file_idl_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5c, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x48, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x22, 0x20, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8e, 0x01,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x1e,
	0x5a, 0x1c, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_user_proto_rawDescOnce sync.Once
	file_idl_user_proto_rawDescData = file_idl_user_proto_rawDesc
)

func file_idl_user_proto_rawDescGZIP() []byte {
	file_idl_user_proto_rawDescOnce.Do(func() {
		file_idl_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_user_proto_rawDescData)
	})
	return file_idl_user_proto_rawDescData
}

var file_idl_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_user_proto_goTypes = []interface{}{
	(*GetUserListReq)(nil),  // 0: service.GetUserListReq
	(*GetUserListResp)(nil), // 1: service.GetUserListResp
	(*UserData)(nil),        // 2: service.UserData
	(*GetUserByIdReq)(nil),  // 3: service.GetUserByIdReq
}
var file_idl_user_proto_depIdxs = []int32{
	2, // 0: service.GetUserListResp.data:type_name -> service.UserData
	0, // 1: service.UserService.GetUserList:input_type -> service.GetUserListReq
	3, // 2: service.UserService.GetUserById:input_type -> service.GetUserByIdReq
	1, // 3: service.UserService.GetUserList:output_type -> service.GetUserListResp
	2, // 4: service.UserService.GetUserById:output_type -> service.UserData
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_user_proto_init() }
func file_idl_user_proto_init() {
	if File_idl_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListReq); i {
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
		file_idl_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListResp); i {
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
		file_idl_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_idl_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdReq); i {
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
			RawDescriptor: file_idl_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_user_proto_goTypes,
		DependencyIndexes: file_idl_user_proto_depIdxs,
		MessageInfos:      file_idl_user_proto_msgTypes,
	}.Build()
	File_idl_user_proto = out.File
	file_idl_user_proto_rawDesc = nil
	file_idl_user_proto_goTypes = nil
	file_idl_user_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.2. DO NOT EDIT.

type UserService interface {
	GetUserList(ctx context.Context, req *GetUserListReq) (res *GetUserListResp, err error)
	GetUserById(ctx context.Context, req *GetUserByIdReq) (res *UserData, err error)
}
