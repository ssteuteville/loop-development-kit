// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.5
// source: clipboard.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClipboardReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *ClipboardReadRequest) Reset() {
	*x = ClipboardReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardReadRequest) ProtoMessage() {}

func (x *ClipboardReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardReadRequest.ProtoReflect.Descriptor instead.
func (*ClipboardReadRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{0}
}

func (x *ClipboardReadRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type ClipboardReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ClipboardReadResponse) Reset() {
	*x = ClipboardReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardReadResponse) ProtoMessage() {}

func (x *ClipboardReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardReadResponse.ProtoReflect.Descriptor instead.
func (*ClipboardReadResponse) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{1}
}

func (x *ClipboardReadResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ClipboardReadStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *ClipboardReadStreamRequest) Reset() {
	*x = ClipboardReadStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardReadStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardReadStreamRequest) ProtoMessage() {}

func (x *ClipboardReadStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardReadStreamRequest.ProtoReflect.Descriptor instead.
func (*ClipboardReadStreamRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{2}
}

func (x *ClipboardReadStreamRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type ClipboardReadStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Error string `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ClipboardReadStreamResponse) Reset() {
	*x = ClipboardReadStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardReadStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardReadStreamResponse) ProtoMessage() {}

func (x *ClipboardReadStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardReadStreamResponse.ProtoReflect.Descriptor instead.
func (*ClipboardReadStreamResponse) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{3}
}

func (x *ClipboardReadStreamResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ClipboardReadStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ClipboardWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Text    string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ClipboardWriteRequest) Reset() {
	*x = ClipboardWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardWriteRequest) ProtoMessage() {}

func (x *ClipboardWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardWriteRequest.ProtoReflect.Descriptor instead.
func (*ClipboardWriteRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{4}
}

func (x *ClipboardWriteRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *ClipboardWriteRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_clipboard_proto protoreflect.FileDescriptor

var file_clipboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x46, 0x0a, 0x1a, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x1b, 0x43,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xff, 0x01, 0x0a, 0x09,
	0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clipboard_proto_rawDescOnce sync.Once
	file_clipboard_proto_rawDescData = file_clipboard_proto_rawDesc
)

func file_clipboard_proto_rawDescGZIP() []byte {
	file_clipboard_proto_rawDescOnce.Do(func() {
		file_clipboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_clipboard_proto_rawDescData)
	})
	return file_clipboard_proto_rawDescData
}

var file_clipboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_clipboard_proto_goTypes = []interface{}{
	(*ClipboardReadRequest)(nil),        // 0: proto.ClipboardReadRequest
	(*ClipboardReadResponse)(nil),       // 1: proto.ClipboardReadResponse
	(*ClipboardReadStreamRequest)(nil),  // 2: proto.ClipboardReadStreamRequest
	(*ClipboardReadStreamResponse)(nil), // 3: proto.ClipboardReadStreamResponse
	(*ClipboardWriteRequest)(nil),       // 4: proto.ClipboardWriteRequest
	(*Session)(nil),                     // 5: proto.Session
	(*emptypb.Empty)(nil),               // 6: google.protobuf.Empty
}
var file_clipboard_proto_depIdxs = []int32{
	5, // 0: proto.ClipboardReadRequest.session:type_name -> proto.Session
	5, // 1: proto.ClipboardReadStreamRequest.session:type_name -> proto.Session
	5, // 2: proto.ClipboardWriteRequest.session:type_name -> proto.Session
	0, // 3: proto.Clipboard.ClipboardRead:input_type -> proto.ClipboardReadRequest
	2, // 4: proto.Clipboard.ClipboardReadStream:input_type -> proto.ClipboardReadStreamRequest
	4, // 5: proto.Clipboard.ClipboardWrite:input_type -> proto.ClipboardWriteRequest
	1, // 6: proto.Clipboard.ClipboardRead:output_type -> proto.ClipboardReadResponse
	3, // 7: proto.Clipboard.ClipboardReadStream:output_type -> proto.ClipboardReadStreamResponse
	6, // 8: proto.Clipboard.ClipboardWrite:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_clipboard_proto_init() }
func file_clipboard_proto_init() {
	if File_clipboard_proto != nil {
		return
	}
	file_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_clipboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardReadRequest); i {
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
		file_clipboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardReadResponse); i {
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
		file_clipboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardReadStreamRequest); i {
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
		file_clipboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardReadStreamResponse); i {
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
		file_clipboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardWriteRequest); i {
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
			RawDescriptor: file_clipboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clipboard_proto_goTypes,
		DependencyIndexes: file_clipboard_proto_depIdxs,
		MessageInfos:      file_clipboard_proto_msgTypes,
	}.Build()
	File_clipboard_proto = out.File
	file_clipboard_proto_rawDesc = nil
	file_clipboard_proto_goTypes = nil
	file_clipboard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClipboardClient is the client API for Clipboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClipboardClient interface {
	// get the contents of the clipboard
	ClipboardRead(ctx context.Context, in *ClipboardReadRequest, opts ...grpc.CallOption) (*ClipboardReadResponse, error)
	// get the contents of the clipboard every time they change
	ClipboardReadStream(ctx context.Context, in *ClipboardReadStreamRequest, opts ...grpc.CallOption) (Clipboard_ClipboardReadStreamClient, error)
	// write to the clipboard
	ClipboardWrite(ctx context.Context, in *ClipboardWriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clipboardClient struct {
	cc grpc.ClientConnInterface
}

func NewClipboardClient(cc grpc.ClientConnInterface) ClipboardClient {
	return &clipboardClient{cc}
}

func (c *clipboardClient) ClipboardRead(ctx context.Context, in *ClipboardReadRequest, opts ...grpc.CallOption) (*ClipboardReadResponse, error) {
	out := new(ClipboardReadResponse)
	err := c.cc.Invoke(ctx, "/proto.Clipboard/ClipboardRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clipboardClient) ClipboardReadStream(ctx context.Context, in *ClipboardReadStreamRequest, opts ...grpc.CallOption) (Clipboard_ClipboardReadStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Clipboard_serviceDesc.Streams[0], "/proto.Clipboard/ClipboardReadStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &clipboardClipboardReadStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Clipboard_ClipboardReadStreamClient interface {
	Recv() (*ClipboardReadStreamResponse, error)
	grpc.ClientStream
}

type clipboardClipboardReadStreamClient struct {
	grpc.ClientStream
}

func (x *clipboardClipboardReadStreamClient) Recv() (*ClipboardReadStreamResponse, error) {
	m := new(ClipboardReadStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clipboardClient) ClipboardWrite(ctx context.Context, in *ClipboardWriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Clipboard/ClipboardWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClipboardServer is the server API for Clipboard service.
type ClipboardServer interface {
	// get the contents of the clipboard
	ClipboardRead(context.Context, *ClipboardReadRequest) (*ClipboardReadResponse, error)
	// get the contents of the clipboard every time they change
	ClipboardReadStream(*ClipboardReadStreamRequest, Clipboard_ClipboardReadStreamServer) error
	// write to the clipboard
	ClipboardWrite(context.Context, *ClipboardWriteRequest) (*emptypb.Empty, error)
}

// UnimplementedClipboardServer can be embedded to have forward compatible implementations.
type UnimplementedClipboardServer struct {
}

func (*UnimplementedClipboardServer) ClipboardRead(context.Context, *ClipboardReadRequest) (*ClipboardReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClipboardRead not implemented")
}
func (*UnimplementedClipboardServer) ClipboardReadStream(*ClipboardReadStreamRequest, Clipboard_ClipboardReadStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClipboardReadStream not implemented")
}
func (*UnimplementedClipboardServer) ClipboardWrite(context.Context, *ClipboardWriteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClipboardWrite not implemented")
}

func RegisterClipboardServer(s *grpc.Server, srv ClipboardServer) {
	s.RegisterService(&_Clipboard_serviceDesc, srv)
}

func _Clipboard_ClipboardRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClipboardReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClipboardServer).ClipboardRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clipboard/ClipboardRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClipboardServer).ClipboardRead(ctx, req.(*ClipboardReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clipboard_ClipboardReadStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClipboardReadStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClipboardServer).ClipboardReadStream(m, &clipboardClipboardReadStreamServer{stream})
}

type Clipboard_ClipboardReadStreamServer interface {
	Send(*ClipboardReadStreamResponse) error
	grpc.ServerStream
}

type clipboardClipboardReadStreamServer struct {
	grpc.ServerStream
}

func (x *clipboardClipboardReadStreamServer) Send(m *ClipboardReadStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Clipboard_ClipboardWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClipboardWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClipboardServer).ClipboardWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clipboard/ClipboardWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClipboardServer).ClipboardWrite(ctx, req.(*ClipboardWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clipboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Clipboard",
	HandlerType: (*ClipboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClipboardRead",
			Handler:    _Clipboard_ClipboardRead_Handler,
		},
		{
			MethodName: "ClipboardWrite",
			Handler:    _Clipboard_ClipboardWrite_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClipboardReadStream",
			Handler:       _Clipboard_ClipboardReadStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "clipboard.proto",
}
