// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: proto/model/logstream.proto

package model

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

type CmdOP int32

const (
	CmdOP_OPEN  CmdOP = 0
	CmdOP_CLOSE CmdOP = 1
	CmdOP_READ  CmdOP = 2
)

// Enum value maps for CmdOP.
var (
	CmdOP_name = map[int32]string{
		0: "OPEN",
		1: "CLOSE",
		2: "READ",
	}
	CmdOP_value = map[string]int32{
		"OPEN":  0,
		"CLOSE": 1,
		"READ":  2,
	}
)

func (x CmdOP) Enum() *CmdOP {
	p := new(CmdOP)
	*p = x
	return p
}

func (x CmdOP) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdOP) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_model_logstream_proto_enumTypes[0].Descriptor()
}

func (CmdOP) Type() protoreflect.EnumType {
	return &file_proto_model_logstream_proto_enumTypes[0]
}

func (x CmdOP) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdOP.Descriptor instead.
func (CmdOP) EnumDescriptor() ([]byte, []int) {
	return file_proto_model_logstream_proto_rawDescGZIP(), []int{0}
}

type TargetNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *TargetNode) Reset() {
	*x = TargetNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_logstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetNode) ProtoMessage() {}

func (x *TargetNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_logstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetNode.ProtoReflect.Descriptor instead.
func (*TargetNode) Descriptor() ([]byte, []int) {
	return file_proto_model_logstream_proto_rawDescGZIP(), []int{0}
}

func (x *TargetNode) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type LogCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op CmdOP `protobuf:"varint,1,opt,name=op,proto3,enum=sbot.proto.model.CmdOP" json:"op,omitempty"`
	// the nodeId of cbot
	NodeId string `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *LogCmd) Reset() {
	*x = LogCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_logstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCmd) ProtoMessage() {}

func (x *LogCmd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_logstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCmd.ProtoReflect.Descriptor instead.
func (*LogCmd) Descriptor() ([]byte, []int) {
	return file_proto_model_logstream_proto_rawDescGZIP(), []int{1}
}

func (x *LogCmd) GetOp() CmdOP {
	if x != nil {
		return x.Op
	}
	return CmdOP_OPEN
}

func (x *LogCmd) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type OPStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//0---ok,-1 failed
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// the status description
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//nodeId
	NodeId string `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *OPStatus) Reset() {
	*x = OPStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_logstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OPStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OPStatus) ProtoMessage() {}

func (x *OPStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_logstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OPStatus.ProtoReflect.Descriptor instead.
func (*OPStatus) Descriptor() ([]byte, []int) {
	return file_proto_model_logstream_proto_rawDescGZIP(), []int{2}
}

func (x *OPStatus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OPStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OPStatus) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type LogStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//nodeId
	NodeId string `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Dsize  uint64 `protobuf:"varint,2,opt,name=dsize,proto3" json:"dsize,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LogStream) Reset() {
	*x = LogStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_logstream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStream) ProtoMessage() {}

func (x *LogStream) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_logstream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStream.ProtoReflect.Descriptor instead.
func (*LogStream) Descriptor() ([]byte, []int) {
	return file_proto_model_logstream_proto_rawDescGZIP(), []int{3}
}

func (x *LogStream) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *LogStream) GetDsize() uint64 {
	if x != nil {
		return x.Dsize
	}
	return 0
}

func (x *LogStream) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_model_logstream_proto protoreflect.FileDescriptor

var file_proto_model_logstream_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x24, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x43, 0x6d, 0x64, 0x12,
	0x27, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x62,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x6d, 0x64, 0x4f, 0x50, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x22, 0x54, 0x0a, 0x08, 0x4f, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x26, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x4f, 0x50, 0x12, 0x08,
	0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x42, 0x6e, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x62, 0x77, 0x33, 0x62, 0x61, 0x6f, 0x2e, 0x73, 0x62, 0x6f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x09, 0x4c,
	0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x01, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x62, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xaa, 0x02, 0x10, 0x53, 0x42, 0x6f, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xca, 0x02, 0x10, 0x53, 0x42, 0x6f,
	0x74, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_model_logstream_proto_rawDescOnce sync.Once
	file_proto_model_logstream_proto_rawDescData = file_proto_model_logstream_proto_rawDesc
)

func file_proto_model_logstream_proto_rawDescGZIP() []byte {
	file_proto_model_logstream_proto_rawDescOnce.Do(func() {
		file_proto_model_logstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_model_logstream_proto_rawDescData)
	})
	return file_proto_model_logstream_proto_rawDescData
}

var file_proto_model_logstream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_model_logstream_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_model_logstream_proto_goTypes = []interface{}{
	(CmdOP)(0),         // 0: sbot.proto.model.CmdOP
	(*TargetNode)(nil), // 1: sbot.proto.model.TargetNode
	(*LogCmd)(nil),     // 2: sbot.proto.model.LogCmd
	(*OPStatus)(nil),   // 3: sbot.proto.model.OPStatus
	(*LogStream)(nil),  // 4: sbot.proto.model.LogStream
}
var file_proto_model_logstream_proto_depIdxs = []int32{
	0, // 0: sbot.proto.model.LogCmd.op:type_name -> sbot.proto.model.CmdOP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_model_logstream_proto_init() }
func file_proto_model_logstream_proto_init() {
	if File_proto_model_logstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_model_logstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetNode); i {
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
		file_proto_model_logstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogCmd); i {
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
		file_proto_model_logstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OPStatus); i {
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
		file_proto_model_logstream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogStream); i {
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
			RawDescriptor: file_proto_model_logstream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_model_logstream_proto_goTypes,
		DependencyIndexes: file_proto_model_logstream_proto_depIdxs,
		EnumInfos:         file_proto_model_logstream_proto_enumTypes,
		MessageInfos:      file_proto_model_logstream_proto_msgTypes,
	}.Build()
	File_proto_model_logstream_proto = out.File
	file_proto_model_logstream_proto_rawDesc = nil
	file_proto_model_logstream_proto_goTypes = nil
	file_proto_model_logstream_proto_depIdxs = nil
}
