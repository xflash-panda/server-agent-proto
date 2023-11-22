// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pkg/agent.proto

package pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeType int32

const (
	NodeType_SHADOWSOCKS NodeType = 0
	NodeType_TROJAN      NodeType = 1
	NodeType_VMESS       NodeType = 2
	NodeType_HYSTERIA    NodeType = 3
	NodeType_HYSTERIA2   NodeType = 4
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "SHADOWSOCKS",
		1: "TROJAN",
		2: "VMESS",
		3: "HYSTERIA",
		4: "HYSTERIA2",
	}
	NodeType_value = map[string]int32{
		"SHADOWSOCKS": 0,
		"TROJAN":      1,
		"VMESS":       2,
		"HYSTERIA":    3,
		"HYSTERIA2":   4,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_agent_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_pkg_agent_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_agent_proto_rawDescGZIP(), []int{0}
}

type CommonParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	CoinType NodeType `protobuf:"varint,2,opt,name=coinType,proto3,enum=pkg.NodeType" json:"coinType,omitempty"`
}

func (x *CommonParams) Reset() {
	*x = CommonParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonParams) ProtoMessage() {}

func (x *CommonParams) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonParams.ProtoReflect.Descriptor instead.
func (*CommonParams) Descriptor() ([]byte, []int) {
	return file_pkg_agent_proto_rawDescGZIP(), []int{0}
}

func (x *CommonParams) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *CommonParams) GetCoinType() NodeType {
	if x != nil {
		return x.CoinType
	}
	return NodeType_SHADOWSOCKS
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *CommonParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigRequest) GetParams() *CommonParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool       `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Data   *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_pkg_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ConfigResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *CommonParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatRequest) GetParams() *CommonParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_agent_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_pkg_agent_proto protoreflect.FileDescriptor

var file_pkg_agent_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x70, 0x6b, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x52, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2a, 0x4f, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x48, 0x41, 0x44, 0x4f, 0x57, 0x53, 0x4f, 0x43, 0x4b, 0x53, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x52, 0x4f, 0x4a, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x4d,
	0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x59, 0x53, 0x54, 0x45, 0x52, 0x49,
	0x41, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x59, 0x53, 0x54, 0x45, 0x52, 0x49, 0x41, 0x32,
	0x10, 0x04, 0x32, 0x73, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x2d, 0x70, 0x61, 0x6e,
	0x64, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_agent_proto_rawDescOnce sync.Once
	file_pkg_agent_proto_rawDescData = file_pkg_agent_proto_rawDesc
)

func file_pkg_agent_proto_rawDescGZIP() []byte {
	file_pkg_agent_proto_rawDescOnce.Do(func() {
		file_pkg_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_agent_proto_rawDescData)
	})
	return file_pkg_agent_proto_rawDescData
}

var file_pkg_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_agent_proto_goTypes = []interface{}{
	(NodeType)(0),             // 0: pkg.NodeType
	(*CommonParams)(nil),      // 1: pkg.CommonParams
	(*ConfigRequest)(nil),     // 2: pkg.ConfigRequest
	(*ConfigResponse)(nil),    // 3: pkg.ConfigResponse
	(*HeartbeatRequest)(nil),  // 4: pkg.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 5: pkg.HeartbeatResponse
	(*anypb.Any)(nil),         // 6: google.protobuf.Any
}
var file_pkg_agent_proto_depIdxs = []int32{
	0, // 0: pkg.CommonParams.coinType:type_name -> pkg.NodeType
	1, // 1: pkg.ConfigRequest.params:type_name -> pkg.CommonParams
	6, // 2: pkg.ConfigResponse.data:type_name -> google.protobuf.Any
	1, // 3: pkg.HeartbeatRequest.params:type_name -> pkg.CommonParams
	2, // 4: pkg.Agent.Config:input_type -> pkg.ConfigRequest
	4, // 5: pkg.Agent.Heartbeat:input_type -> pkg.HeartbeatRequest
	3, // 6: pkg.Agent.Config:output_type -> pkg.ConfigResponse
	3, // 7: pkg.Agent.Heartbeat:output_type -> pkg.ConfigResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_agent_proto_init() }
func file_pkg_agent_proto_init() {
	if File_pkg_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonParams); i {
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
		file_pkg_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_pkg_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
		file_pkg_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_pkg_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
			RawDescriptor: file_pkg_agent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_agent_proto_goTypes,
		DependencyIndexes: file_pkg_agent_proto_depIdxs,
		EnumInfos:         file_pkg_agent_proto_enumTypes,
		MessageInfos:      file_pkg_agent_proto_msgTypes,
	}.Build()
	File_pkg_agent_proto = out.File
	file_pkg_agent_proto_rawDesc = nil
	file_pkg_agent_proto_goTypes = nil
	file_pkg_agent_proto_depIdxs = nil
}
