// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: pb/message.proto

package pb

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

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Height     uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ListenAddr string   `protobuf:"bytes,3,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	Peers      []string `protobuf:"bytes,4,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_pb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_pb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Info) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Info) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *Info) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Height     uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	PrevHash   []byte `protobuf:"bytes,3,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	Nonce      uint32 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Difficulty uint32 `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	MerkleRoot []byte `protobuf:"bytes,6,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	Timestamp  int64  `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pb_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_pb_message_proto_rawDescGZIP(), []int{1}
}

func (x *BlockHeader) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockHeader) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockHeader) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *BlockHeader) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockHeader) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *BlockHeader) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *BlockHeader) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	MinerKey     []byte         `protobuf:"bytes,3,opt,name=miner_key,json=minerKey,proto3" json:"miner_key,omitempty"`
	Signature    []byte         `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_pb_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_pb_message_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetMinerKey() []byte {
	if x != nil {
		return x.MinerKey
	}
	return nil
}

func (x *Block) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TransactionHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Sender    []byte  `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver  []byte  `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount    float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp int64   `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransactionHeader) Reset() {
	*x = TransactionHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionHeader) ProtoMessage() {}

func (x *TransactionHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pb_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionHeader.ProtoReflect.Descriptor instead.
func (*TransactionHeader) Descriptor() ([]byte, []int) {
	return file_pb_message_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionHeader) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TransactionHeader) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TransactionHeader) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *TransactionHeader) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionHeader) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *TransactionHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Signature []byte             `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_pb_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_pb_message_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetHeader() *TransactionHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_pb_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_pb_message_proto_rawDescGZIP(), []int{5}
}

var File_pb_message_proto protoreflect.FileDescriptor

var file_pb_message_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x69,
	0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x57,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x04, 0x0a, 0x02, 0x4f, 0x6b, 0x32, 0x65, 0x0a,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61,
	0x6b, 0x65, 0x12, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x26, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x03, 0x2e, 0x4f, 0x6b, 0x12, 0x1a, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a,
	0x03, 0x2e, 0x4f, 0x6b, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x70, 0x68, 0x61, 0x64, 0x61, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_message_proto_rawDescOnce sync.Once
	file_pb_message_proto_rawDescData = file_pb_message_proto_rawDesc
)

func file_pb_message_proto_rawDescGZIP() []byte {
	file_pb_message_proto_rawDescOnce.Do(func() {
		file_pb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_message_proto_rawDescData)
	})
	return file_pb_message_proto_rawDescData
}

var file_pb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_message_proto_goTypes = []interface{}{
	(*Info)(nil),              // 0: Info
	(*BlockHeader)(nil),       // 1: BlockHeader
	(*Block)(nil),             // 2: Block
	(*TransactionHeader)(nil), // 3: TransactionHeader
	(*Transaction)(nil),       // 4: Transaction
	(*Ok)(nil),                // 5: Ok
}
var file_pb_message_proto_depIdxs = []int32{
	1, // 0: Block.header:type_name -> BlockHeader
	4, // 1: Block.transactions:type_name -> Transaction
	3, // 2: Transaction.header:type_name -> TransactionHeader
	0, // 3: Node.Handshake:input_type -> Info
	4, // 4: Node.SubmitTransaction:input_type -> Transaction
	2, // 5: Node.SubmitBlock:input_type -> Block
	0, // 6: Node.Handshake:output_type -> Info
	5, // 7: Node.SubmitTransaction:output_type -> Ok
	5, // 8: Node.SubmitBlock:output_type -> Ok
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_message_proto_init() }
func file_pb_message_proto_init() {
	if File_pb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_pb_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_pb_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_pb_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionHeader); i {
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
		file_pb_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_pb_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
			RawDescriptor: file_pb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_message_proto_goTypes,
		DependencyIndexes: file_pb_message_proto_depIdxs,
		MessageInfos:      file_pb_message_proto_msgTypes,
	}.Build()
	File_pb_message_proto = out.File
	file_pb_message_proto_rawDesc = nil
	file_pb_message_proto_goTypes = nil
	file_pb_message_proto_depIdxs = nil
}
