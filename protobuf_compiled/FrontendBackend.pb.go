// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: FrontendBackend.proto

package protobuf_compiled

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

type ToFrontEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ToFrontEnd_GithubState
	//	*ToFrontEnd_AccessCode
	Type isToFrontEnd_Type `protobuf_oneof:"type"`
}

func (x *ToFrontEnd) Reset() {
	*x = ToFrontEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FrontendBackend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToFrontEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToFrontEnd) ProtoMessage() {}

func (x *ToFrontEnd) ProtoReflect() protoreflect.Message {
	mi := &file_FrontendBackend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToFrontEnd.ProtoReflect.Descriptor instead.
func (*ToFrontEnd) Descriptor() ([]byte, []int) {
	return file_FrontendBackend_proto_rawDescGZIP(), []int{0}
}

func (m *ToFrontEnd) GetType() isToFrontEnd_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ToFrontEnd) GetGithubState() *Nonce {
	if x, ok := x.GetType().(*ToFrontEnd_GithubState); ok {
		return x.GithubState
	}
	return nil
}

func (x *ToFrontEnd) GetAccessCode() *AccessCode {
	if x, ok := x.GetType().(*ToFrontEnd_AccessCode); ok {
		return x.AccessCode
	}
	return nil
}

type isToFrontEnd_Type interface {
	isToFrontEnd_Type()
}

type ToFrontEnd_GithubState struct {
	GithubState *Nonce `protobuf:"bytes,1,opt,name=github_state,json=githubState,proto3,oneof"`
}

type ToFrontEnd_AccessCode struct {
	AccessCode *AccessCode `protobuf:"bytes,2,opt,name=access_code,json=accessCode,proto3,oneof"`
}

func (*ToFrontEnd_GithubState) isToFrontEnd_Type() {}

func (*ToFrontEnd_AccessCode) isToFrontEnd_Type() {}

type Nonce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Nonce) Reset() {
	*x = Nonce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FrontendBackend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nonce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nonce) ProtoMessage() {}

func (x *Nonce) ProtoReflect() protoreflect.Message {
	mi := &file_FrontendBackend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nonce.ProtoReflect.Descriptor instead.
func (*Nonce) Descriptor() ([]byte, []int) {
	return file_FrontendBackend_proto_rawDescGZIP(), []int{1}
}

func (x *Nonce) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type AccessCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessCode string `protobuf:"bytes,1,opt,name=access_code,json=accessCode,proto3" json:"access_code,omitempty"`
	Failed     bool   `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
}

func (x *AccessCode) Reset() {
	*x = AccessCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FrontendBackend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessCode) ProtoMessage() {}

func (x *AccessCode) ProtoReflect() protoreflect.Message {
	mi := &file_FrontendBackend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessCode.ProtoReflect.Descriptor instead.
func (*AccessCode) Descriptor() ([]byte, []int) {
	return file_FrontendBackend_proto_rawDescGZIP(), []int{2}
}

func (x *AccessCode) GetAccessCode() string {
	if x != nil {
		return x.AccessCode
	}
	return ""
}

func (x *AccessCode) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

var File_FrontendBackend_proto protoreflect.FileDescriptor

var file_FrontendBackend_proto_rawDesc = []byte{
	0x0a, 0x15, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0a, 0x54, 0x6f, 0x46, 0x72, 0x6f,
	0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4e, 0x6f,
	0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x4e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_FrontendBackend_proto_rawDescOnce sync.Once
	file_FrontendBackend_proto_rawDescData = file_FrontendBackend_proto_rawDesc
)

func file_FrontendBackend_proto_rawDescGZIP() []byte {
	file_FrontendBackend_proto_rawDescOnce.Do(func() {
		file_FrontendBackend_proto_rawDescData = protoimpl.X.CompressGZIP(file_FrontendBackend_proto_rawDescData)
	})
	return file_FrontendBackend_proto_rawDescData
}

var file_FrontendBackend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_FrontendBackend_proto_goTypes = []interface{}{
	(*ToFrontEnd)(nil), // 0: ToFrontEnd
	(*Nonce)(nil),      // 1: Nonce
	(*AccessCode)(nil), // 2: AccessCode
}
var file_FrontendBackend_proto_depIdxs = []int32{
	1, // 0: ToFrontEnd.github_state:type_name -> Nonce
	2, // 1: ToFrontEnd.access_code:type_name -> AccessCode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FrontendBackend_proto_init() }
func file_FrontendBackend_proto_init() {
	if File_FrontendBackend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FrontendBackend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToFrontEnd); i {
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
		file_FrontendBackend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nonce); i {
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
		file_FrontendBackend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessCode); i {
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
	file_FrontendBackend_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ToFrontEnd_GithubState)(nil),
		(*ToFrontEnd_AccessCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FrontendBackend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FrontendBackend_proto_goTypes,
		DependencyIndexes: file_FrontendBackend_proto_depIdxs,
		MessageInfos:      file_FrontendBackend_proto_msgTypes,
	}.Build()
	File_FrontendBackend_proto = out.File
	file_FrontendBackend_proto_rawDesc = nil
	file_FrontendBackend_proto_goTypes = nil
	file_FrontendBackend_proto_depIdxs = nil
}
