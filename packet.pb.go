// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: packet.proto

package decnet

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

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    string   `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Listener  int32    `protobuf:"varint,2,opt,name=Listener,proto3" json:"Listener,omitempty"`
	Username  string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	PublicKey string   `protobuf:"bytes,4,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Headers   []string `protobuf:"bytes,5,rep,name=Headers,proto3" json:"Headers,omitempty"`
	Sender    string   `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Completed bool     `protobuf:"varint,7,opt,name=completed,proto3" json:"completed,omitempty"`
	Body      []byte   `protobuf:"bytes,8,opt,name=Body,proto3" json:"Body,omitempty"`
	Created   int64    `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{0}
}

func (x *Packet) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Packet) GetListener() int32 {
	if x != nil {
		return x.Listener
	}
	return 0
}

func (x *Packet) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Packet) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Packet) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Packet) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Packet) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *Packet) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Packet) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

var File_packet_proto protoreflect.FileDescriptor

var file_packet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x65, 0x63, 0x6e, 0x65, 0x74, 0x22, 0xf4, 0x01, 0x0a, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packet_proto_rawDescOnce sync.Once
	file_packet_proto_rawDescData = file_packet_proto_rawDesc
)

func file_packet_proto_rawDescGZIP() []byte {
	file_packet_proto_rawDescOnce.Do(func() {
		file_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_proto_rawDescData)
	})
	return file_packet_proto_rawDescData
}

var file_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_packet_proto_goTypes = []interface{}{
	(*Packet)(nil), // 0: decnet.packet
}
var file_packet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_packet_proto_init() }
func file_packet_proto_init() {
	if File_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
			RawDescriptor: file_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_proto_goTypes,
		DependencyIndexes: file_packet_proto_depIdxs,
		MessageInfos:      file_packet_proto_msgTypes,
	}.Build()
	File_packet_proto = out.File
	file_packet_proto_rawDesc = nil
	file_packet_proto_goTypes = nil
	file_packet_proto_depIdxs = nil
}
