
package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	mi := &file_proto_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Product       string                 `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderReply) Reset() {
	*x = OrderReply{}
	mi := &file_proto_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReply) ProtoMessage() {}

func (x *OrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*OrderReply) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderReply) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderReply) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *OrderReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type OrderUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderUpdate) Reset() {
	*x = OrderUpdate{}
	mi := &file_proto_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdate) ProtoMessage() {}

func (x *OrderUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*OrderUpdate) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderUpdate) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderUpdate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_order_proto protoreflect.FileDescriptor

const file_proto_order_proto_rawDesc = "" +
	"\n" +
	"\x11proto/order.proto\x12\x05order\")\n" +
	"\fOrderRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"Y\n" +
	"\n" +
	"OrderReply\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x18\n" +
	"\aproduct\x18\x02 \x01(\tR\aproduct\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\"@\n" +
	"\vOrderUpdate\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status2{\n" +
	"\fOrderService\x122\n" +
	"\bGetOrder\x12\x13.order.OrderRequest\x1a\x11.order.OrderReply\x127\n" +
	"\n" +
	"TrackOrder\x12\x13.order.OrderRequest\x1a\x12.order.OrderUpdate0\x01B)Z'example.com/ordertracking/proto;orderpbb\x06proto3"

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData []byte
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)))
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)

var file_proto_order_proto_goTypes = []any{
	(*OrderRequest)(nil),
	(*OrderReply)(nil),
	(*OrderUpdate)(nil),
}

var file_proto_order_proto_depIdxs = []int32{
	0,
	0,
	1,
	2,
	2,
	0,
	0,
	0,
	0,
}

func init() { file_proto_order_proto_init() }

func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
