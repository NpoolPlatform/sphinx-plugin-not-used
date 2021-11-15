// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0--rc1
// source: npool/plugin.proto

package agents

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

// GetSignInfo 参数
type GetSignInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // 发送方钱包地址
}

func (x *GetSignInfoRequest) Reset() {
	*x = GetSignInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignInfoRequest) ProtoMessage() {}

func (x *GetSignInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignInfoRequest.ProtoReflect.Descriptor instead.
func (*GetSignInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *GetSignInfoRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// GetSignInfo 返回
type SignInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"` // 需要的预签名信息
}

func (x *SignInfo) Reset() {
	*x = SignInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInfo) ProtoMessage() {}

func (x *SignInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInfo.ProtoReflect.Descriptor instead.
func (*SignInfo) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *SignInfo) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

// GetBalance 参数
type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId       int32  `protobuf:"varint,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                                // 查询的钱包地址
	TimestampUtc int64  `protobuf:"varint,3,opt,name=timestamp_utc,json=timestampUtc,proto3" json:"timestamp_utc,omitempty"` // 长整型时间戳
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceRequest) GetCoinId() int32 {
	if x != nil {
		return x.CoinId
	}
	return 0
}

func (x *GetBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetBalanceRequest) GetTimestampUtc() int64 {
	if x != nil {
		return x.TimestampUtc
	}
	return 0
}

// GetBalance 返回
type AccountBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinId       int32  `protobuf:"varint,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                                // 查询的钱包地址
	TimestampUtc int64  `protobuf:"varint,3,opt,name=timestamp_utc,json=timestampUtc,proto3" json:"timestamp_utc,omitempty"` // 长整型时间戳
	AmountInt    int64  `protobuf:"varint,4,opt,name=amount_int,json=amountInt,proto3" json:"amount_int,omitempty"`          // 金额整数
	AmountDigits int32  `protobuf:"varint,5,opt,name=amount_digits,json=amountDigits,proto3" json:"amount_digits,omitempty"` // 金额*了10的^n，默认为9
	AmountString string `protobuf:"bytes,6,opt,name=amount_string,json=amountString,proto3" json:"amount_string,omitempty"`  // 金额字符串，"123.45678901"
}

func (x *AccountBalance) Reset() {
	*x = AccountBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBalance) ProtoMessage() {}

func (x *AccountBalance) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBalance.ProtoReflect.Descriptor instead.
func (*AccountBalance) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *AccountBalance) GetCoinId() int32 {
	if x != nil {
		return x.CoinId
	}
	return 0
}

func (x *AccountBalance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AccountBalance) GetTimestampUtc() int64 {
	if x != nil {
		return x.TimestampUtc
	}
	return 0
}

func (x *AccountBalance) GetAmountInt() int64 {
	if x != nil {
		return x.AmountInt
	}
	return 0
}

func (x *AccountBalance) GetAmountDigits() int32 {
	if x != nil {
		return x.AmountDigits
	}
	return 0
}

func (x *AccountBalance) GetAmountString() string {
	if x != nil {
		return x.AmountString
	}
	return ""
}

// BroadcastScript 参数
type BroadcastScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionScript string `protobuf:"bytes,1,opt,name=transaction_script,json=transactionScript,proto3" json:"transaction_script,omitempty"`
}

func (x *BroadcastScriptRequest) Reset() {
	*x = BroadcastScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastScriptRequest) ProtoMessage() {}

func (x *BroadcastScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastScriptRequest.ProtoReflect.Descriptor instead.
func (*BroadcastScriptRequest) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastScriptRequest) GetTransactionScript() string {
	if x != nil {
		return x.TransactionScript
	}
	return ""
}

// BroadcastScript 返回
type BroadcastScriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionIdChain string `protobuf:"bytes,1,opt,name=transaction_id_chain,json=transactionIdChain,proto3" json:"transaction_id_chain,omitempty"`
}

func (x *BroadcastScriptResponse) Reset() {
	*x = BroadcastScriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastScriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastScriptResponse) ProtoMessage() {}

func (x *BroadcastScriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastScriptResponse.ProtoReflect.Descriptor instead.
func (*BroadcastScriptResponse) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *BroadcastScriptResponse) GetTransactionIdChain() string {
	if x != nil {
		return x.TransactionIdChain
	}
	return ""
}

// GetTxStatus 参数
type GetTxStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionIdChain string `protobuf:"bytes,1,opt,name=transaction_id_chain,json=transactionIdChain,proto3" json:"transaction_id_chain,omitempty"`
}

func (x *GetTxStatusRequest) Reset() {
	*x = GetTxStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxStatusRequest) ProtoMessage() {}

func (x *GetTxStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxStatusRequest.ProtoReflect.Descriptor instead.
func (*GetTxStatusRequest) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *GetTxStatusRequest) GetTransactionIdChain() string {
	if x != nil {
		return x.TransactionIdChain
	}
	return ""
}

// GetTxStatus 返回
type GetTxStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmountInt          int64  `protobuf:"varint,1,opt,name=amount_int,json=amountInt,proto3" json:"amount_int,omitempty"`                               // 放大后的金额整数
	AmountDigits       int32  `protobuf:"varint,2,opt,name=amount_digits,json=amountDigits,proto3" json:"amount_digits,omitempty"`                      // amount_int == amount*10^n
	AmountString       string `protobuf:"bytes,3,opt,name=amount_string,json=amountString,proto3" json:"amount_string,omitempty"`                       // 便于验证，数据库里不存
	AddressFrom        string `protobuf:"bytes,4,opt,name=address_from,json=addressFrom,proto3" json:"address_from,omitempty"`                          // 发送方
	AddressTo          string `protobuf:"bytes,5,opt,name=address_to,json=addressTo,proto3" json:"address_to,omitempty"`                                // 接收方
	TransactionIdChain string `protobuf:"bytes,6,opt,name=transaction_id_chain,json=transactionIdChain,proto3" json:"transaction_id_chain,omitempty"`   // 公链交易ID
	CreatetimeUtc      int64  `protobuf:"varint,11,opt,name=createtime_utc,json=createtimeUtc,proto3" json:"createtime_utc,omitempty"`                  // 创建时间
	UpdatetimeUtc      int64  `protobuf:"varint,12,opt,name=updatetime_utc,json=updatetimeUtc,proto3" json:"updatetime_utc,omitempty"`                  // 上次更新时间
	IsSuccess          bool   `protobuf:"varint,13,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`                              // 便于调用方判断
	IsFailed           bool   `protobuf:"varint,14,opt,name=is_failed,json=isFailed,proto3" json:"is_failed,omitempty"`                                 // 不success不fail就是pending了
	NumBlocksConfirmed int32  `protobuf:"varint,15,opt,name=num_blocks_confirmed,json=numBlocksConfirmed,proto3" json:"num_blocks_confirmed,omitempty"` // 已确认区块数
}

func (x *GetTxStatusResponse) Reset() {
	*x = GetTxStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxStatusResponse) ProtoMessage() {}

func (x *GetTxStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxStatusResponse.ProtoReflect.Descriptor instead.
func (*GetTxStatusResponse) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *GetTxStatusResponse) GetAmountInt() int64 {
	if x != nil {
		return x.AmountInt
	}
	return 0
}

func (x *GetTxStatusResponse) GetAmountDigits() int32 {
	if x != nil {
		return x.AmountDigits
	}
	return 0
}

func (x *GetTxStatusResponse) GetAmountString() string {
	if x != nil {
		return x.AmountString
	}
	return ""
}

func (x *GetTxStatusResponse) GetAddressFrom() string {
	if x != nil {
		return x.AddressFrom
	}
	return ""
}

func (x *GetTxStatusResponse) GetAddressTo() string {
	if x != nil {
		return x.AddressTo
	}
	return ""
}

func (x *GetTxStatusResponse) GetTransactionIdChain() string {
	if x != nil {
		return x.TransactionIdChain
	}
	return ""
}

func (x *GetTxStatusResponse) GetCreatetimeUtc() int64 {
	if x != nil {
		return x.CreatetimeUtc
	}
	return 0
}

func (x *GetTxStatusResponse) GetUpdatetimeUtc() int64 {
	if x != nil {
		return x.UpdatetimeUtc
	}
	return 0
}

func (x *GetTxStatusResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *GetTxStatusResponse) GetIsFailed() bool {
	if x != nil {
		return x.IsFailed
	}
	return false
}

func (x *GetTxStatusResponse) GetNumBlocksConfirmed() int32 {
	if x != nil {
		return x.NumBlocksConfirmed
	}
	return 0
}

// GetTxJSONRequest 参数
type GetTxJSONRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 继承钱包节点基础功能，预留
	CoinId      int32  `protobuf:"varint,1,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                             // 要查询的钱包地址
	TimefromUtc uint64 `protobuf:"varint,3,opt,name=timefrom_utc,json=timefromUtc,proto3" json:"timefrom_utc,omitempty"` // 开始时间
	TimetillUtc uint64 `protobuf:"varint,4,opt,name=timetill_utc,json=timetillUtc,proto3" json:"timetill_utc,omitempty"` // 结束时间
	LimitN      int32  `protobuf:"varint,5,opt,name=limit_n,json=limitN,proto3" json:"limit_n,omitempty"`                // 服务端限制返回条数
}

func (x *GetTxJSONRequest) Reset() {
	*x = GetTxJSONRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxJSONRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxJSONRequest) ProtoMessage() {}

func (x *GetTxJSONRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxJSONRequest.ProtoReflect.Descriptor instead.
func (*GetTxJSONRequest) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{8}
}

func (x *GetTxJSONRequest) GetCoinId() int32 {
	if x != nil {
		return x.CoinId
	}
	return 0
}

func (x *GetTxJSONRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetTxJSONRequest) GetTimefromUtc() uint64 {
	if x != nil {
		return x.TimefromUtc
	}
	return 0
}

func (x *GetTxJSONRequest) GetTimetillUtc() uint64 {
	if x != nil {
		return x.TimetillUtc
	}
	return 0
}

func (x *GetTxJSONRequest) GetLimitN() int32 {
	if x != nil {
		return x.LimitN
	}
	return 0
}

// GetTxJSONRequest 返回
type AccountTxJSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *AccountTxJSON) Reset() {
	*x = AccountTxJSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_plugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountTxJSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountTxJSON) ProtoMessage() {}

func (x *AccountTxJSON) ProtoReflect() protoreflect.Message {
	mi := &file_npool_plugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountTxJSON.ProtoReflect.Descriptor instead.
func (*AccountTxJSON) Descriptor() ([]byte, []int) {
	return file_npool_plugin_proto_rawDescGZIP(), []int{9}
}

func (x *AccountTxJSON) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

var File_npool_plugin_proto protoreflect.FileDescriptor

var file_npool_plugin_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31, 0x22,
	0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x1e, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22,
	0x6b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x55, 0x74, 0x63, 0x22, 0xd1, 0x01, 0x0a,
	0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f,
	0x75, 0x74, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x55, 0x74, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x22, 0x47, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x4b, 0x0a, 0x17, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x78, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x22, 0xae,
	0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x74, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x74,
	0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a,
	0x14, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6e, 0x75, 0x6d,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22,
	0xa4, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x78, 0x4a, 0x53, 0x4f, 0x4e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74,
	0x69, 0x6d, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x74, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x74, 0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x69, 0x6c, 0x6c, 0x55, 0x74, 0x63, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4e, 0x22, 0x23, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x78, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x88, 0x03, 0x0a, 0x06,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x70, 0x68, 0x69,
	0x6e, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x68,
	0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x78, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1b, 0x2e,
	0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x4a,
	0x53, 0x4f, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x70, 0x68,
	0x69, 0x6e, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x78,
	0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_plugin_proto_rawDescOnce sync.Once
	file_npool_plugin_proto_rawDescData = file_npool_plugin_proto_rawDesc
)

func file_npool_plugin_proto_rawDescGZIP() []byte {
	file_npool_plugin_proto_rawDescOnce.Do(func() {
		file_npool_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_plugin_proto_rawDescData)
	})
	return file_npool_plugin_proto_rawDescData
}

var file_npool_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_plugin_proto_goTypes = []interface{}{
	(*GetSignInfoRequest)(nil),      // 0: sphinx.v1.GetSignInfoRequest
	(*SignInfo)(nil),                // 1: sphinx.v1.SignInfo
	(*GetBalanceRequest)(nil),       // 2: sphinx.v1.GetBalanceRequest
	(*AccountBalance)(nil),          // 3: sphinx.v1.AccountBalance
	(*BroadcastScriptRequest)(nil),  // 4: sphinx.v1.BroadcastScriptRequest
	(*BroadcastScriptResponse)(nil), // 5: sphinx.v1.BroadcastScriptResponse
	(*GetTxStatusRequest)(nil),      // 6: sphinx.v1.GetTxStatusRequest
	(*GetTxStatusResponse)(nil),     // 7: sphinx.v1.GetTxStatusResponse
	(*GetTxJSONRequest)(nil),        // 8: sphinx.v1.GetTxJSONRequest
	(*AccountTxJSON)(nil),           // 9: sphinx.v1.AccountTxJSON
}
var file_npool_plugin_proto_depIdxs = []int32{
	0, // 0: sphinx.v1.Plugin.GetSignInfo:input_type -> sphinx.v1.GetSignInfoRequest
	2, // 1: sphinx.v1.Plugin.GetBalance:input_type -> sphinx.v1.GetBalanceRequest
	4, // 2: sphinx.v1.Plugin.BroadcastScript:input_type -> sphinx.v1.BroadcastScriptRequest
	6, // 3: sphinx.v1.Plugin.GetTxStatus:input_type -> sphinx.v1.GetTxStatusRequest
	8, // 4: sphinx.v1.Plugin.GetTxJSON:input_type -> sphinx.v1.GetTxJSONRequest
	1, // 5: sphinx.v1.Plugin.GetSignInfo:output_type -> sphinx.v1.SignInfo
	3, // 6: sphinx.v1.Plugin.GetBalance:output_type -> sphinx.v1.AccountBalance
	5, // 7: sphinx.v1.Plugin.BroadcastScript:output_type -> sphinx.v1.BroadcastScriptResponse
	7, // 8: sphinx.v1.Plugin.GetTxStatus:output_type -> sphinx.v1.GetTxStatusResponse
	9, // 9: sphinx.v1.Plugin.GetTxJSON:output_type -> sphinx.v1.AccountTxJSON
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_plugin_proto_init() }
func file_npool_plugin_proto_init() {
	if File_npool_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignInfoRequest); i {
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
		file_npool_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInfo); i {
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
		file_npool_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_npool_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBalance); i {
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
		file_npool_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastScriptRequest); i {
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
		file_npool_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastScriptResponse); i {
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
		file_npool_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxStatusRequest); i {
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
		file_npool_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxStatusResponse); i {
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
		file_npool_plugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxJSONRequest); i {
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
		file_npool_plugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountTxJSON); i {
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
			RawDescriptor: file_npool_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_plugin_proto_goTypes,
		DependencyIndexes: file_npool_plugin_proto_depIdxs,
		MessageInfos:      file_npool_plugin_proto_msgTypes,
	}.Build()
	File_npool_plugin_proto = out.File
	file_npool_plugin_proto_rawDesc = nil
	file_npool_plugin_proto_goTypes = nil
	file_npool_plugin_proto_depIdxs = nil
}
