// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: group_rpc.proto

package group_rpc

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

type IsInGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId uint32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *IsInGroupRequest) Reset() {
	*x = IsInGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsInGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInGroupRequest) ProtoMessage() {}

func (x *IsInGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInGroupRequest.ProtoReflect.Descriptor instead.
func (*IsInGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *IsInGroupRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsInGroupRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type IsInGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsInGroup bool `protobuf:"varint,1,opt,name=is_in_group,json=isInGroup,proto3" json:"is_in_group,omitempty"`
}

func (x *IsInGroupResponse) Reset() {
	*x = IsInGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsInGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInGroupResponse) ProtoMessage() {}

func (x *IsInGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInGroupResponse.ProtoReflect.Descriptor instead.
func (*IsInGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *IsInGroupResponse) GetIsInGroup() bool {
	if x != nil {
		return x.IsInGroup
	}
	return false
}

type UserGroupSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIdList []uint32 `protobuf:"varint,1,rep,packed,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"`
	Mode       uint32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"` // 模式 1.创建群聊的个数   2.加入群聊的个数
}

func (x *UserGroupSearchRequest) Reset() {
	*x = UserGroupSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupSearchRequest) ProtoMessage() {}

func (x *UserGroupSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupSearchRequest.ProtoReflect.Descriptor instead.
func (*UserGroupSearchRequest) Descriptor() ([]byte, []int) {
	return file_group_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UserGroupSearchRequest) GetUserIdList() []uint32 {
	if x != nil {
		return x.UserIdList
	}
	return nil
}

func (x *UserGroupSearchRequest) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type UserGroupSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result map[int32]int32 `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 结果
}

func (x *UserGroupSearchResponse) Reset() {
	*x = UserGroupSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupSearchResponse) ProtoMessage() {}

func (x *UserGroupSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupSearchResponse.ProtoReflect.Descriptor instead.
func (*UserGroupSearchResponse) Descriptor() ([]byte, []int) {
	return file_group_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UserGroupSearchResponse) GetResult() map[int32]int32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_group_rpc_proto protoreflect.FileDescriptor

var file_group_rpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x46, 0x0a, 0x10,
	0x49, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x11, 0x49, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x4e, 0x0a, 0x16, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x17, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xaa, 0x01, 0x0a, 0x06, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x46, 0x0a, 0x09, 0x49, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x49,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x49, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x21,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_rpc_proto_rawDescOnce sync.Once
	file_group_rpc_proto_rawDescData = file_group_rpc_proto_rawDesc
)

func file_group_rpc_proto_rawDescGZIP() []byte {
	file_group_rpc_proto_rawDescOnce.Do(func() {
		file_group_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_rpc_proto_rawDescData)
	})
	return file_group_rpc_proto_rawDescData
}

var file_group_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_group_rpc_proto_goTypes = []interface{}{
	(*IsInGroupRequest)(nil),        // 0: group_rpc.IsInGroupRequest
	(*IsInGroupResponse)(nil),       // 1: group_rpc.IsInGroupResponse
	(*UserGroupSearchRequest)(nil),  // 2: group_rpc.UserGroupSearchRequest
	(*UserGroupSearchResponse)(nil), // 3: group_rpc.UserGroupSearchResponse
	nil,                             // 4: group_rpc.UserGroupSearchResponse.ResultEntry
}
var file_group_rpc_proto_depIdxs = []int32{
	4, // 0: group_rpc.UserGroupSearchResponse.result:type_name -> group_rpc.UserGroupSearchResponse.ResultEntry
	0, // 1: group_rpc.Groups.IsInGroup:input_type -> group_rpc.IsInGroupRequest
	2, // 2: group_rpc.Groups.UserGroupSearch:input_type -> group_rpc.UserGroupSearchRequest
	1, // 3: group_rpc.Groups.IsInGroup:output_type -> group_rpc.IsInGroupResponse
	3, // 4: group_rpc.Groups.UserGroupSearch:output_type -> group_rpc.UserGroupSearchResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_group_rpc_proto_init() }
func file_group_rpc_proto_init() {
	if File_group_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsInGroupRequest); i {
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
		file_group_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsInGroupResponse); i {
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
		file_group_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupSearchRequest); i {
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
		file_group_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupSearchResponse); i {
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
			RawDescriptor: file_group_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_rpc_proto_goTypes,
		DependencyIndexes: file_group_rpc_proto_depIdxs,
		MessageInfos:      file_group_rpc_proto_msgTypes,
	}.Build()
	File_group_rpc_proto = out.File
	file_group_rpc_proto_rawDesc = nil
	file_group_rpc_proto_goTypes = nil
	file_group_rpc_proto_depIdxs = nil
}
