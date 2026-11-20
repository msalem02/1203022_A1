package powpb

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

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkerId      string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_proto_pow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_pow_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *RegisterRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Accepted      bool                   `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	mi := &file_proto_pow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_proto_pow_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkerId      string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_proto_pow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_pow_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type HeartbeatReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatReply) Reset() {
	*x = HeartbeatReply{}
	mi := &file_proto_pow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReply) ProtoMessage() {}

func (x *HeartbeatReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*HeartbeatReply) Descriptor() ([]byte, []int) {
	return file_proto_pow_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type SearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Data          string                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Difficulty    uint32                 `protobuf:"varint,3,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Start         uint64                 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	End           uint64                 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_proto_pow_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pow_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_pow_proto_rawDescGZIP(), []int{4}
}

func (x *SearchRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *SearchRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SearchRequest) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *SearchRequest) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SearchRequest) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

type SearchUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Checked       uint64                 `protobuf:"varint,2,opt,name=checked,proto3" json:"checked,omitempty"`
	Finished      bool                   `protobuf:"varint,3,opt,name=finished,proto3" json:"finished,omitempty"`
	Found         bool                   `protobuf:"varint,4,opt,name=found,proto3" json:"found,omitempty"`
	Nonce         uint64                 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Hash          string                 `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchUpdate) Reset() {
	*x = SearchUpdate{}
	mi := &file_proto_pow_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUpdate) ProtoMessage() {}

func (x *SearchUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pow_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*SearchUpdate) Descriptor() ([]byte, []int) {
	return file_proto_pow_proto_rawDescGZIP(), []int{5}
}

func (x *SearchUpdate) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *SearchUpdate) GetChecked() uint64 {
	if x != nil {
		return x.Checked
	}
	return 0
}

func (x *SearchUpdate) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *SearchUpdate) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *SearchUpdate) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *SearchUpdate) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_proto_pow_proto protoreflect.FileDescriptor

const file_proto_pow_proto_rawDesc = "" +
	"\n" +
	"\x0fproto/pow.proto\x12\x03pow\"H\n" +
	"\x0fRegisterRequest\x12\x1b\n" +
	"\tworker_id\x18\x01 \x01(\tR\bworkerId\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\"+\n" +
	"\rRegisterReply\x12\x1a\n" +
	"\baccepted\x18\x01 \x01(\bR\baccepted\"/\n" +
	"\x10HeartbeatRequest\x12\x1b\n" +
	"\tworker_id\x18\x01 \x01(\tR\bworkerId\" \n" +
	"\x0eHeartbeatReply\x12\x0e\n" +
	"\x02ok\x18\x01 \x01(\bR\x02ok\"\x82\x01\n" +
	"\rSearchRequest\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x12\n" +
	"\x04data\x18\x02 \x01(\tR\x04data\x12\x1e\n" +
	"\n" +
	"difficulty\x18\x03 \x01(\rR\n" +
	"difficulty\x12\x14\n" +
	"\x05start\x18\x04 \x01(\x04R\x05start\x12\x10\n" +
	"\x03end\x18\x05 \x01(\x04R\x03end\"\x9b\x01\n" +
	"\fSearchUpdate\x12\x15\n" +
	"\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x18\n" +
	"\achecked\x18\x02 \x01(\x04R\achecked\x12\x1a\n" +
	"\bfinished\x18\x03 \x01(\bR\bfinished\x12\x14\n" +
	"\x05found\x18\x04 \x01(\bR\x05found\x12\x14\n" +
	"\x05nonce\x18\x05 \x01(\x04R\x05nonce\x12\x12\n" +
	"\x04hash\x18\x06 \x01(\tR\x04hash2\x89\x01\n" +
	"\x12CoordinatorService\x12:\n" +
	"\x0eRegisterWorker\x12\x14.pow.RegisterRequest\x1a\x12.pow.RegisterReply\x127\n" +
	"\tHeartbeat\x12\x15.pow.HeartbeatRequest\x1a\x13.pow.HeartbeatReply2G\n" +
	"\rWorkerService\x126\n" +
	"\vSearchRange\x12\x12.pow.SearchRequest\x1a\x11.pow.SearchUpdate0\x01B(Z&example.com/distributedpow/proto;powpbb\x06proto3"

var (
	file_proto_pow_proto_rawDescOnce sync.Once
	file_proto_pow_proto_rawDescData []byte
)

func file_proto_pow_proto_rawDescGZIP() []byte {
	file_proto_pow_proto_rawDescOnce.Do(func() {
		file_proto_pow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_pow_proto_rawDesc), len(file_proto_pow_proto_rawDesc)))
	})
	return file_proto_pow_proto_rawDescData
}

var file_proto_pow_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_pow_proto_goTypes = []any{
	(*RegisterRequest)(nil),
	(*RegisterReply)(nil),
	(*HeartbeatRequest)(nil),
	(*HeartbeatReply)(nil),
	(*SearchRequest)(nil),
	(*SearchUpdate)(nil),
}
var file_proto_pow_proto_depIdxs = []int32{
	0,
	2,
	4,
	1,
	3,
	5,
	3,
	0,
	0,
	0,
	0,
}

func init() { file_proto_pow_proto_init() }
func file_proto_pow_proto_init() {
	if File_proto_pow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pow_proto_rawDesc), len(file_proto_pow_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_pow_proto_goTypes,
		DependencyIndexes: file_proto_pow_proto_depIdxs,
		MessageInfos:      file_proto_pow_proto_msgTypes,
	}.Build()
	File_proto_pow_proto = out.File
	file_proto_pow_proto_goTypes = nil
	file_proto_pow_proto_depIdxs = nil
}
