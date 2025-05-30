// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: internal/adapters/outbound/grpc_pb/notification.proto

package grpc_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionNotification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount        float32                `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionNotification) Reset() {
	*x = TransactionNotification{}
	mi := &file_internal_adapters_outbound_grpc_pb_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionNotification) ProtoMessage() {}

func (x *TransactionNotification) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_outbound_grpc_pb_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionNotification.ProtoReflect.Descriptor instead.
func (*TransactionNotification) Descriptor() ([]byte, []int) {
	return file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionNotification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionNotification) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionNotification) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TransactionNotification) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type NotificationAcknowledgment struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NotificationId string                 `protobuf:"bytes,2,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	Status         string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *NotificationAcknowledgment) Reset() {
	*x = NotificationAcknowledgment{}
	mi := &file_internal_adapters_outbound_grpc_pb_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationAcknowledgment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationAcknowledgment) ProtoMessage() {}

func (x *NotificationAcknowledgment) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_outbound_grpc_pb_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationAcknowledgment.ProtoReflect.Descriptor instead.
func (*NotificationAcknowledgment) Descriptor() ([]byte, []int) {
	return file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationAcknowledgment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationAcknowledgment) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *NotificationAcknowledgment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_internal_adapters_outbound_grpc_pb_notification_proto protoreflect.FileDescriptor

const file_internal_adapters_outbound_grpc_pb_notification_proto_rawDesc = "" +
	"\n" +
	"5internal/adapters/outbound/grpc_pb/notification.proto\x12\agrpc_pb\"{\n" +
	"\x17TransactionNotification\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x02R\x06amount\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"m\n" +
	"\x1aNotificationAcknowledgment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12'\n" +
	"\x0fnotification_id\x18\x02 \x01(\tR\x0enotificationId\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status2\xd2\x01\n" +
	"\x13NotificationService\x12Y\n" +
	"\x10SendNotification\x12 .grpc_pb.TransactionNotification\x1a#.grpc_pb.NotificationAcknowledgment\x12`\n" +
	"\x13StreamNotifications\x12 .grpc_pb.TransactionNotification\x1a#.grpc_pb.NotificationAcknowledgment(\x010\x01BGZEgithub.com/KKogaa/grpc-transaction/internal/adapters/outbound/grpc_pbb\x06proto3"

var (
	file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescOnce sync.Once
	file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescData []byte
)

func file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescGZIP() []byte {
	file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescOnce.Do(func() {
		file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_adapters_outbound_grpc_pb_notification_proto_rawDesc), len(file_internal_adapters_outbound_grpc_pb_notification_proto_rawDesc)))
	})
	return file_internal_adapters_outbound_grpc_pb_notification_proto_rawDescData
}

var file_internal_adapters_outbound_grpc_pb_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_adapters_outbound_grpc_pb_notification_proto_goTypes = []any{
	(*TransactionNotification)(nil),    // 0: grpc_pb.TransactionNotification
	(*NotificationAcknowledgment)(nil), // 1: grpc_pb.NotificationAcknowledgment
}
var file_internal_adapters_outbound_grpc_pb_notification_proto_depIdxs = []int32{
	0, // 0: grpc_pb.NotificationService.SendNotification:input_type -> grpc_pb.TransactionNotification
	0, // 1: grpc_pb.NotificationService.StreamNotifications:input_type -> grpc_pb.TransactionNotification
	1, // 2: grpc_pb.NotificationService.SendNotification:output_type -> grpc_pb.NotificationAcknowledgment
	1, // 3: grpc_pb.NotificationService.StreamNotifications:output_type -> grpc_pb.NotificationAcknowledgment
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_adapters_outbound_grpc_pb_notification_proto_init() }
func file_internal_adapters_outbound_grpc_pb_notification_proto_init() {
	if File_internal_adapters_outbound_grpc_pb_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_adapters_outbound_grpc_pb_notification_proto_rawDesc), len(file_internal_adapters_outbound_grpc_pb_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_adapters_outbound_grpc_pb_notification_proto_goTypes,
		DependencyIndexes: file_internal_adapters_outbound_grpc_pb_notification_proto_depIdxs,
		MessageInfos:      file_internal_adapters_outbound_grpc_pb_notification_proto_msgTypes,
	}.Build()
	File_internal_adapters_outbound_grpc_pb_notification_proto = out.File
	file_internal_adapters_outbound_grpc_pb_notification_proto_goTypes = nil
	file_internal_adapters_outbound_grpc_pb_notification_proto_depIdxs = nil
}
