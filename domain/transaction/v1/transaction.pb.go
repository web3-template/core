// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: transaction/v1/transaction.proto

package transactionv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Transaction_Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	Contract string                `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Network  uint32                `protobuf:"varint,3,opt,name=network,proto3" json:"network,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Transaction_SafeTransferFrom
	//	*Transaction_SafeBatchTransferFrom
	//	*Transaction_SafeApprovalForAll
	Data isTransaction_Data `protobuf_oneof:"data"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_v1_transaction_proto_msgTypes[0]
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
	return file_transaction_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetMetadata() *Transaction_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Transaction) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *Transaction) GetNetwork() uint32 {
	if x != nil {
		return x.Network
	}
	return 0
}

func (m *Transaction) GetData() isTransaction_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Transaction) GetSafeTransferFrom() *Transaction_ERC1155_SafeTransferFrom {
	if x, ok := x.GetData().(*Transaction_SafeTransferFrom); ok {
		return x.SafeTransferFrom
	}
	return nil
}

func (x *Transaction) GetSafeBatchTransferFrom() *Transaction_ERC1155_SafeBatchTransferFrom {
	if x, ok := x.GetData().(*Transaction_SafeBatchTransferFrom); ok {
		return x.SafeBatchTransferFrom
	}
	return nil
}

func (x *Transaction) GetSafeApprovalForAll() *Transaction_ERC1155_SetApprovalForAll {
	if x, ok := x.GetData().(*Transaction_SafeApprovalForAll); ok {
		return x.SafeApprovalForAll
	}
	return nil
}

type isTransaction_Data interface {
	isTransaction_Data()
}

type Transaction_SafeTransferFrom struct {
	SafeTransferFrom *Transaction_ERC1155_SafeTransferFrom `protobuf:"bytes,4,opt,name=safe_transfer_from,json=safeTransferFrom,proto3,oneof"`
}

type Transaction_SafeBatchTransferFrom struct {
	SafeBatchTransferFrom *Transaction_ERC1155_SafeBatchTransferFrom `protobuf:"bytes,5,opt,name=safe_batch_transfer_from,json=safeBatchTransferFrom,proto3,oneof"`
}

type Transaction_SafeApprovalForAll struct {
	SafeApprovalForAll *Transaction_ERC1155_SetApprovalForAll `protobuf:"bytes,6,opt,name=safe_approval_for_all,json=safeApprovalForAll,proto3,oneof"`
}

func (*Transaction_SafeTransferFrom) isTransaction_Data() {}

func (*Transaction_SafeBatchTransferFrom) isTransaction_Data() {}

func (*Transaction_SafeApprovalForAll) isTransaction_Data() {}

type Transaction_ERC1155 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Transaction_ERC1155) Reset() {
	*x = Transaction_ERC1155{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_v1_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC1155) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC1155) ProtoMessage() {}

func (x *Transaction_ERC1155) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_v1_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC1155.ProtoReflect.Descriptor instead.
func (*Transaction_ERC1155) Descriptor() ([]byte, []int) {
	return file_transaction_v1_transaction_proto_rawDescGZIP(), []int{0, 0}
}

type Transaction_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     string                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Date     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Reverted *string                `protobuf:"bytes,3,opt,name=reverted,proto3,oneof" json:"reverted,omitempty"`
}

func (x *Transaction_Metadata) Reset() {
	*x = Transaction_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_v1_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_Metadata) ProtoMessage() {}

func (x *Transaction_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_v1_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_Metadata.ProtoReflect.Descriptor instead.
func (*Transaction_Metadata) Descriptor() ([]byte, []int) {
	return file_transaction_v1_transaction_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Transaction_Metadata) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction_Metadata) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Transaction_Metadata) GetReverted() string {
	if x != nil && x.Reverted != nil {
		return *x.Reverted
	}
	return ""
}

type Transaction_ERC1155_SafeTransferFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Id    uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Value uint64 `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	Data  []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction_ERC1155_SafeTransferFrom) Reset() {
	*x = Transaction_ERC1155_SafeTransferFrom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_v1_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC1155_SafeTransferFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC1155_SafeTransferFrom) ProtoMessage() {}

func (x *Transaction_ERC1155_SafeTransferFrom) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_v1_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC1155_SafeTransferFrom.ProtoReflect.Descriptor instead.
func (*Transaction_ERC1155_SafeTransferFrom) Descriptor() ([]byte, []int) {
	return file_transaction_v1_transaction_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Transaction_ERC1155_SafeTransferFrom) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction_ERC1155_SafeTransferFrom) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction_ERC1155_SafeTransferFrom) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction_ERC1155_SafeTransferFrom) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Transaction_ERC1155_SafeTransferFrom) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Transaction_ERC1155_SafeBatchTransferFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To     string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Ids    []uint64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Values []uint64 `protobuf:"varint,4,rep,packed,name=values,proto3" json:"values,omitempty"`
	Data   []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) Reset() {
	*x = Transaction_ERC1155_SafeBatchTransferFrom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_v1_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC1155_SafeBatchTransferFrom) ProtoMessage() {}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_v1_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC1155_SafeBatchTransferFrom.ProtoReflect.Descriptor instead.
func (*Transaction_ERC1155_SafeBatchTransferFrom) Descriptor() ([]byte, []int) {
	return file_transaction_v1_transaction_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) GetValues() []uint64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Transaction_ERC1155_SafeBatchTransferFrom) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Transaction_ERC1155_SetApprovalForAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Approved bool   `protobuf:"varint,2,opt,name=approved,proto3" json:"approved,omitempty"`
}

func (x *Transaction_ERC1155_SetApprovalForAll) Reset() {
	*x = Transaction_ERC1155_SetApprovalForAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_v1_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_ERC1155_SetApprovalForAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_ERC1155_SetApprovalForAll) ProtoMessage() {}

func (x *Transaction_ERC1155_SetApprovalForAll) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_v1_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_ERC1155_SetApprovalForAll.ProtoReflect.Descriptor instead.
func (*Transaction_ERC1155_SetApprovalForAll) Descriptor() ([]byte, []int) {
	return file_transaction_v1_transaction_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *Transaction_ERC1155_SetApprovalForAll) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *Transaction_ERC1155_SetApprovalForAll) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

var File_transaction_v1_transaction_proto protoreflect.FileDescriptor

var file_transaction_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x07, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x64, 0x0a, 0x12, 0x73, 0x61, 0x66, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35,
	0x35, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x48, 0x00, 0x52, 0x10, 0x73, 0x61, 0x66, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x74, 0x0a, 0x18, 0x73, 0x61, 0x66, 0x65, 0x5f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x2e, 0x53, 0x61,
	0x66, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46,
	0x72, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x15, 0x73, 0x61, 0x66, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x6a, 0x0a, 0x15,
	0x73, 0x61, 0x66, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x66, 0x6f,
	0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35,
	0x2e, 0x53, 0x65, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x46, 0x6f, 0x72, 0x41,
	0x6c, 0x6c, 0x48, 0x00, 0x52, 0x12, 0x73, 0x61, 0x66, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x1a, 0xc3, 0x02, 0x0a, 0x07, 0x45, 0x52, 0x43,
	0x31, 0x31, 0x35, 0x35, 0x1a, 0x70, 0x0a, 0x10, 0x53, 0x61, 0x66, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x79, 0x0a, 0x15, 0x53, 0x61, 0x66, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x4b, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x7c,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0xc2, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x2d, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_v1_transaction_proto_rawDescOnce sync.Once
	file_transaction_v1_transaction_proto_rawDescData = file_transaction_v1_transaction_proto_rawDesc
)

func file_transaction_v1_transaction_proto_rawDescGZIP() []byte {
	file_transaction_v1_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_v1_transaction_proto_rawDescData)
	})
	return file_transaction_v1_transaction_proto_rawDescData
}

var file_transaction_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transaction_v1_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),                               // 0: transaction.v1.Transaction
	(*Transaction_ERC1155)(nil),                       // 1: transaction.v1.Transaction.ERC1155
	(*Transaction_Metadata)(nil),                      // 2: transaction.v1.Transaction.Metadata
	(*Transaction_ERC1155_SafeTransferFrom)(nil),      // 3: transaction.v1.Transaction.ERC1155.SafeTransferFrom
	(*Transaction_ERC1155_SafeBatchTransferFrom)(nil), // 4: transaction.v1.Transaction.ERC1155.SafeBatchTransferFrom
	(*Transaction_ERC1155_SetApprovalForAll)(nil),     // 5: transaction.v1.Transaction.ERC1155.SetApprovalForAll
	(*timestamppb.Timestamp)(nil),                     // 6: google.protobuf.Timestamp
}
var file_transaction_v1_transaction_proto_depIdxs = []int32{
	2, // 0: transaction.v1.Transaction.metadata:type_name -> transaction.v1.Transaction.Metadata
	3, // 1: transaction.v1.Transaction.safe_transfer_from:type_name -> transaction.v1.Transaction.ERC1155.SafeTransferFrom
	4, // 2: transaction.v1.Transaction.safe_batch_transfer_from:type_name -> transaction.v1.Transaction.ERC1155.SafeBatchTransferFrom
	5, // 3: transaction.v1.Transaction.safe_approval_for_all:type_name -> transaction.v1.Transaction.ERC1155.SetApprovalForAll
	6, // 4: transaction.v1.Transaction.Metadata.date:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transaction_v1_transaction_proto_init() }
func file_transaction_v1_transaction_proto_init() {
	if File_transaction_v1_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_v1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_transaction_v1_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC1155); i {
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
		file_transaction_v1_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_Metadata); i {
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
		file_transaction_v1_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC1155_SafeTransferFrom); i {
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
		file_transaction_v1_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC1155_SafeBatchTransferFrom); i {
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
		file_transaction_v1_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_ERC1155_SetApprovalForAll); i {
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
	file_transaction_v1_transaction_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Transaction_SafeTransferFrom)(nil),
		(*Transaction_SafeBatchTransferFrom)(nil),
		(*Transaction_SafeApprovalForAll)(nil),
	}
	file_transaction_v1_transaction_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transaction_v1_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_v1_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_v1_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_v1_transaction_proto_msgTypes,
	}.Build()
	File_transaction_v1_transaction_proto = out.File
	file_transaction_v1_transaction_proto_rawDesc = nil
	file_transaction_v1_transaction_proto_goTypes = nil
	file_transaction_v1_transaction_proto_depIdxs = nil
}
