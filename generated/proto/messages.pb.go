// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: proto/messages.proto

package messages

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

type Vector2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *Vector2) Reset() {
	*x = Vector2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector2) ProtoMessage() {}

func (x *Vector2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector2.ProtoReflect.Descriptor instead.
func (*Vector2) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Vector2) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector2) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type ClientConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *ClientConnectionResponse) Reset() {
	*x = ClientConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnectionResponse) ProtoMessage() {}

func (x *ClientConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnectionResponse.ProtoReflect.Descriptor instead.
func (*ClientConnectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConnectionResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type ClientInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input *Vector2 `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *ClientInputRequest) Reset() {
	*x = ClientInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInputRequest) ProtoMessage() {}

func (x *ClientInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInputRequest.ProtoReflect.Descriptor instead.
func (*ClientInputRequest) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInputRequest) GetInput() *Vector2 {
	if x != nil {
		return x.Input
	}
	return nil
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientPosition *Vector2 `protobuf:"bytes,2,opt,name=clientPosition,proto3" json:"clientPosition,omitempty"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GameState) GetClientPosition() *Vector2 {
	if x != nil {
		return x.ClientPosition
	}
	return nil
}

type WrapperMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   string `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	MessageType int32  `protobuf:"varint,2,opt,name=messageType,proto3" json:"messageType,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*WrapperMessage_ClientConnectionResponse
	//	*WrapperMessage_ClientInputRequest
	Msg isWrapperMessage_Msg `protobuf_oneof:"msg"`
}

func (x *WrapperMessage) Reset() {
	*x = WrapperMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperMessage) ProtoMessage() {}

func (x *WrapperMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperMessage.ProtoReflect.Descriptor instead.
func (*WrapperMessage) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{4}
}

func (x *WrapperMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *WrapperMessage) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (m *WrapperMessage) GetMsg() isWrapperMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *WrapperMessage) GetClientConnectionResponse() *ClientConnectionResponse {
	if x, ok := x.GetMsg().(*WrapperMessage_ClientConnectionResponse); ok {
		return x.ClientConnectionResponse
	}
	return nil
}

func (x *WrapperMessage) GetClientInputRequest() *ClientInputRequest {
	if x, ok := x.GetMsg().(*WrapperMessage_ClientInputRequest); ok {
		return x.ClientInputRequest
	}
	return nil
}

type isWrapperMessage_Msg interface {
	isWrapperMessage_Msg()
}

type WrapperMessage_ClientConnectionResponse struct {
	ClientConnectionResponse *ClientConnectionResponse `protobuf:"bytes,3,opt,name=clientConnectionResponse,proto3,oneof"`
}

type WrapperMessage_ClientInputRequest struct {
	ClientInputRequest *ClientInputRequest `protobuf:"bytes,4,opt,name=clientInputRequest,proto3,oneof"`
}

func (*WrapperMessage_ClientConnectionResponse) isWrapperMessage_Msg() {}

func (*WrapperMessage_ClientInputRequest) isWrapperMessage_Msg() {}

var File_proto_messages_proto protoreflect.FileDescriptor

var file_proto_messages_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x07, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x32, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x58, 0x12,
	0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x59, 0x22, 0x38, 0x0a,
	0x18, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x3d, 0x0a,
	0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf7, 0x01, 0x0a,
	0x0e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x57, 0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x18,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x12, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6b, 0x6b, 0x6f, 0x72, 0x79, 0x79, 0x6e, 0x61, 0x6e,
	0x65, 0x6e, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_messages_proto_rawDescOnce sync.Once
	file_proto_messages_proto_rawDescData = file_proto_messages_proto_rawDesc
)

func file_proto_messages_proto_rawDescGZIP() []byte {
	file_proto_messages_proto_rawDescOnce.Do(func() {
		file_proto_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messages_proto_rawDescData)
	})
	return file_proto_messages_proto_rawDescData
}

var file_proto_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_messages_proto_goTypes = []interface{}{
	(*Vector2)(nil),                  // 0: Vector2
	(*ClientConnectionResponse)(nil), // 1: ClientConnectionResponse
	(*ClientInputRequest)(nil),       // 2: ClientInputRequest
	(*GameState)(nil),                // 3: GameState
	(*WrapperMessage)(nil),           // 4: WrapperMessage
}
var file_proto_messages_proto_depIdxs = []int32{
	0, // 0: ClientInputRequest.input:type_name -> Vector2
	0, // 1: GameState.clientPosition:type_name -> Vector2
	1, // 2: WrapperMessage.clientConnectionResponse:type_name -> ClientConnectionResponse
	2, // 3: WrapperMessage.clientInputRequest:type_name -> ClientInputRequest
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_messages_proto_init() }
func file_proto_messages_proto_init() {
	if File_proto_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector2); i {
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
		file_proto_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConnectionResponse); i {
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
		file_proto_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInputRequest); i {
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
		file_proto_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameState); i {
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
		file_proto_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperMessage); i {
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
	file_proto_messages_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*WrapperMessage_ClientConnectionResponse)(nil),
		(*WrapperMessage_ClientInputRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_messages_proto_goTypes,
		DependencyIndexes: file_proto_messages_proto_depIdxs,
		MessageInfos:      file_proto_messages_proto_msgTypes,
	}.Build()
	File_proto_messages_proto = out.File
	file_proto_messages_proto_rawDesc = nil
	file_proto_messages_proto_goTypes = nil
	file_proto_messages_proto_depIdxs = nil
}
