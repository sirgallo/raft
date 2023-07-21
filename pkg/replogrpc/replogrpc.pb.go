// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/replogrpc.proto

package replogrpc

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

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   int64  `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Term    int64  `protobuf:"varint,2,opt,name=Term,proto3" json:"Term,omitempty"`
	Command string `protobuf:"bytes,3,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_replogrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_replogrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_proto_replogrpc_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LogEntry) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *LogEntry) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type AppendEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term              int64       `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	LeaderId          string      `protobuf:"bytes,2,opt,name=LeaderId,proto3" json:"LeaderId,omitempty"`
	PrevLogIndex      int64       `protobuf:"varint,3,opt,name=PrevLogIndex,proto3" json:"PrevLogIndex,omitempty"`
	PrevLogTerm       int64       `protobuf:"varint,4,opt,name=PrevLogTerm,proto3" json:"PrevLogTerm,omitempty"`
	Entries           []*LogEntry `protobuf:"bytes,5,rep,name=Entries,proto3" json:"Entries,omitempty"`
	LeaderCommitIndex int64       `protobuf:"varint,6,opt,name=LeaderCommitIndex,proto3" json:"LeaderCommitIndex,omitempty"`
}

func (x *AppendEntry) Reset() {
	*x = AppendEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_replogrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntry) ProtoMessage() {}

func (x *AppendEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_replogrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntry.ProtoReflect.Descriptor instead.
func (*AppendEntry) Descriptor() ([]byte, []int) {
	return file_proto_replogrpc_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntry) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntry) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *AppendEntry) GetPrevLogIndex() int64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntry) GetPrevLogTerm() int64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntry) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntry) GetLeaderCommitIndex() int64 {
	if x != nil {
		return x.LeaderCommitIndex
	}
	return 0
}

type AppendEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *AppendEntryResponse) Reset() {
	*x = AppendEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_replogrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntryResponse) ProtoMessage() {}

func (x *AppendEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_replogrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntryResponse.ProtoReflect.Descriptor instead.
func (*AppendEntryResponse) Descriptor() ([]byte, []int) {
	return file_proto_replogrpc_proto_rawDescGZIP(), []int{2}
}

func (x *AppendEntryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_replogrpc_proto protoreflect.FileDescriptor

var file_proto_replogrpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x6f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x6f, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x4e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0xe0, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x72, 0x65,
	0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x2d, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x2f, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x5b, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x4c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x50, 0x43, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x70,
	0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_replogrpc_proto_rawDescOnce sync.Once
	file_proto_replogrpc_proto_rawDescData = file_proto_replogrpc_proto_rawDesc
)

func file_proto_replogrpc_proto_rawDescGZIP() []byte {
	file_proto_replogrpc_proto_rawDescOnce.Do(func() {
		file_proto_replogrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_replogrpc_proto_rawDescData)
	})
	return file_proto_replogrpc_proto_rawDescData
}

var file_proto_replogrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_replogrpc_proto_goTypes = []interface{}{
	(*LogEntry)(nil),            // 0: replogrpc.LogEntry
	(*AppendEntry)(nil),         // 1: replogrpc.AppendEntry
	(*AppendEntryResponse)(nil), // 2: replogrpc.AppendEntryResponse
}
var file_proto_replogrpc_proto_depIdxs = []int32{
	0, // 0: replogrpc.AppendEntry.Entries:type_name -> replogrpc.LogEntry
	1, // 1: replogrpc.RepLogService.AppendEntryRPC:input_type -> replogrpc.AppendEntry
	2, // 2: replogrpc.RepLogService.AppendEntryRPC:output_type -> replogrpc.AppendEntryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_replogrpc_proto_init() }
func file_proto_replogrpc_proto_init() {
	if File_proto_replogrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_replogrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_proto_replogrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntry); i {
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
		file_proto_replogrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntryResponse); i {
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
			RawDescriptor: file_proto_replogrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_replogrpc_proto_goTypes,
		DependencyIndexes: file_proto_replogrpc_proto_depIdxs,
		MessageInfos:      file_proto_replogrpc_proto_msgTypes,
	}.Build()
	File_proto_replogrpc_proto = out.File
	file_proto_replogrpc_proto_rawDesc = nil
	file_proto_replogrpc_proto_goTypes = nil
	file_proto_replogrpc_proto_depIdxs = nil
}
