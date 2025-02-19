// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package blockdb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transaction_Types int32

const (
	Transaction_UNKNOWN  Transaction_Types = 0
	Transaction_GET      Transaction_Types = 1
	Transaction_PUT      Transaction_Types = 2
	Transaction_DEPOSIT  Transaction_Types = 3
	Transaction_WITHDRAW Transaction_Types = 4
	Transaction_TRANSFER Transaction_Types = 5
)

var Transaction_Types_name = map[int32]string{
	0: "UNKNOWN",
	1: "GET",
	2: "PUT",
	3: "DEPOSIT",
	4: "WITHDRAW",
	5: "TRANSFER",
}

var Transaction_Types_value = map[string]int32{
	"UNKNOWN":  0,
	"GET":      1,
	"PUT":      2,
	"DEPOSIT":  3,
	"WITHDRAW": 4,
	"TRANSFER": 5,
}

func (x Transaction_Types) String() string {
	return proto.EnumName(Transaction_Types_name, int32(x))
}

func (Transaction_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{6, 0}
}

type GetRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetResponse struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Request struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Request) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BooleanResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BooleanResponse) Reset()         { *m = BooleanResponse{} }
func (m *BooleanResponse) String() string { return proto.CompactTextString(m) }
func (*BooleanResponse) ProtoMessage()    {}
func (*BooleanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{3}
}

func (m *BooleanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BooleanResponse.Unmarshal(m, b)
}
func (m *BooleanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BooleanResponse.Marshal(b, m, deterministic)
}
func (m *BooleanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooleanResponse.Merge(m, src)
}
func (m *BooleanResponse) XXX_Size() int {
	return xxx_messageInfo_BooleanResponse.Size(m)
}
func (m *BooleanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BooleanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BooleanResponse proto.InternalMessageInfo

func (m *BooleanResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type TransferRequest struct {
	FromID               string   `protobuf:"bytes,1,opt,name=FromID,proto3" json:"FromID,omitempty"`
	ToID                 string   `protobuf:"bytes,2,opt,name=ToID,proto3" json:"ToID,omitempty"`
	Value                int32    `protobuf:"varint,3,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{4}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetFromID() string {
	if m != nil {
		return m.FromID
	}
	return ""
}

func (m *TransferRequest) GetToID() string {
	if m != nil {
		return m.ToID
	}
	return ""
}

func (m *TransferRequest) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{5}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type Transaction struct {
	Type                 Transaction_Types `protobuf:"varint,1,opt,name=Type,proto3,enum=blockdb.Transaction_Types" json:"Type,omitempty"`
	UserID               string            `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	FromID               string            `protobuf:"bytes,3,opt,name=FromID,proto3" json:"FromID,omitempty"`
	ToID                 string            `protobuf:"bytes,4,opt,name=ToID,proto3" json:"ToID,omitempty"`
	Value                int32             `protobuf:"varint,5,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{6}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetType() Transaction_Types {
	if m != nil {
		return m.Type
	}
	return Transaction_UNKNOWN
}

func (m *Transaction) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Transaction) GetFromID() string {
	if m != nil {
		return m.FromID
	}
	return ""
}

func (m *Transaction) GetToID() string {
	if m != nil {
		return m.ToID
	}
	return ""
}

func (m *Transaction) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Block struct {
	BlockID              int32          `protobuf:"varint,1,opt,name=BlockID,proto3" json:"BlockID,omitempty"`
	PrevHash             string         `protobuf:"bytes,2,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,3,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
	Nonce                string         `protobuf:"bytes,4,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }// initialize
func (m *Block) String() string { return proto.CompactTextString(m) } // write to 1.json files
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{7}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetBlockID() int32 {
	if m != nil {
		return m.BlockID
	}
	return 0
}

func (m *Block) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterEnum("blockdb.Transaction_Types", Transaction_Types_name, Transaction_Types_value)
	proto.RegisterType((*GetRequest)(nil), "blockdb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "blockdb.GetResponse")
	proto.RegisterType((*Request)(nil), "blockdb.Request")
	proto.RegisterType((*BooleanResponse)(nil), "blockdb.BooleanResponse")
	proto.RegisterType((*TransferRequest)(nil), "blockdb.TransferRequest")
	proto.RegisterType((*Null)(nil), "blockdb.Null")
	proto.RegisterType((*Transaction)(nil), "blockdb.Transaction")
	proto.RegisterType((*Block)(nil), "blockdb.Block")
}

func init() { proto.RegisterFile("db.proto", fileDescriptor_8817812184a13374) }

var fileDescriptor_8817812184a13374 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x35, 0xc6, 0x18, 0x3c, 0x4e, 0x1a, 0xb4, 0xb5, 0x5a, 0xcb, 0xa7, 0x88, 0xe6, 0x10, 0xa9,
	0x12, 0x07, 0xe7, 0x90, 0x1c, 0x1b, 0x84, 0xe3, 0x58, 0x8d, 0x30, 0x5a, 0xe3, 0xfa, 0x0c, 0x78,
	0x1b, 0xa3, 0x12, 0xd6, 0x65, 0x97, 0x56, 0xf9, 0x11, 0xfd, 0x91, 0xbd, 0xf5, 0x67, 0x54, 0x2c,
	0x9f, 0xae, 0xdc, 0x56, 0xb9, 0xed, 0xdb, 0x7d, 0x33, 0x6f, 0xde, 0xd3, 0x68, 0x41, 0xdb, 0x06,
	0xe6, 0x3e, 0xa5, 0x9c, 0x22, 0x35, 0x88, 0x69, 0xf8, 0x65, 0x1b, 0x18, 0x17, 0x00, 0x73, 0xc2,
	0x31, 0xf9, 0x9a, 0x11, 0xc6, 0xd1, 0x1b, 0xe8, 0xaf, 0x19, 0x49, 0x17, 0xf6, 0x58, 0x3a, 0x97,
	0x2e, 0x07, 0xb8, 0x44, 0xc6, 0x3b, 0x18, 0x0a, 0x16, 0xdb, 0xd3, 0x84, 0x11, 0x34, 0x02, 0xe5,
	0x93, 0x1f, 0x67, 0x44, 0xb0, 0x14, 0x5c, 0x00, 0xe3, 0x1a, 0xd4, 0xff, 0xf4, 0x69, 0x0a, 0xbb,
	0xed, 0xc2, 0xf7, 0x70, 0x66, 0x51, 0x1a, 0x13, 0x3f, 0xa9, 0x15, 0xc6, 0xa0, 0xae, 0xb2, 0x30,
	0x24, 0x8c, 0x89, 0x0e, 0x1a, 0xae, 0xa0, 0xb1, 0x82, 0x33, 0x2f, 0xf5, 0x13, 0xf6, 0x99, 0xa4,
	0x2d, 0xb5, 0xbb, 0x94, 0x3e, 0x35, 0x6a, 0x05, 0x42, 0x08, 0x7a, 0x1e, 0x5d, 0xd8, 0x42, 0x6c,
	0x80, 0xc5, 0xb9, 0x99, 0x40, 0x6e, 0x4f, 0xd0, 0x87, 0x9e, 0x93, 0xc5, 0xb1, 0xf1, 0x4b, 0x82,
	0xa1, 0xe8, 0xee, 0x87, 0x3c, 0xa2, 0x09, 0x32, 0xa1, 0xe7, 0x3d, 0xef, 0x0b, 0x9f, 0xaf, 0xa6,
	0x13, 0xb3, 0x4c, 0xcd, 0x6c, 0x71, 0xcc, 0x9c, 0xc0, 0xb0, 0xe0, 0xb5, 0x7c, 0x77, 0x0f, 0x7c,
	0x37, 0x13, 0xca, 0x47, 0x27, 0xec, 0x1d, 0x9b, 0x50, 0x69, 0x4f, 0xb8, 0x04, 0x45, 0x08, 0xa1,
	0x21, 0xa8, 0x6b, 0xe7, 0xa3, 0xb3, 0xdc, 0x38, 0x7a, 0x07, 0xa9, 0x20, 0xcf, 0x67, 0x9e, 0x2e,
	0xe5, 0x07, 0x77, 0xed, 0xe9, 0xdd, 0xfc, 0xd9, 0x9e, 0xb9, 0xcb, 0xd5, 0xc2, 0xd3, 0x65, 0x74,
	0x02, 0xda, 0x66, 0xe1, 0xdd, 0xdb, 0xf8, 0x76, 0xa3, 0xf7, 0x72, 0xe4, 0xe1, 0x5b, 0x67, 0x75,
	0x37, 0xc3, 0xba, 0x62, 0xfc, 0x90, 0x40, 0xb1, 0x72, 0x3b, 0x79, 0xd6, 0xe2, 0x50, 0xe6, 0xa7,
	0xe0, 0x0a, 0xa2, 0x09, 0x68, 0x6e, 0x4a, 0xbe, 0xdd, 0xfb, 0x6c, 0x57, 0x1a, 0xaa, 0x31, 0xba,
	0x81, 0x93, 0x56, 0x0a, 0x6c, 0x2c, 0x9f, 0xcb, 0x97, 0xc3, 0xe9, 0xe8, 0x58, 0x44, 0xf8, 0x80,
	0x99, 0x1b, 0x74, 0x68, 0x12, 0x92, 0xd2, 0x75, 0x01, 0xa6, 0x3f, 0xbb, 0x70, 0x2a, 0x74, 0x6d,
	0x9f, 0xfb, 0x81, 0xcf, 0x08, 0x9a, 0x82, 0x3c, 0x27, 0x1c, 0xbd, 0xae, 0x5b, 0x36, 0x8b, 0x3a,
	0x19, 0x1d, 0x5e, 0x16, 0x5b, 0x63, 0x74, 0xd0, 0x15, 0xc8, 0x6e, 0xc6, 0x91, 0x5e, 0x3f, 0x57,
	0x05, 0xe3, 0xfa, 0xe6, 0x8f, 0x55, 0x33, 0x3a, 0xe8, 0x06, 0xb4, 0x4d, 0xc4, 0x77, 0xdb, 0xd4,
	0xff, 0xfe, 0xc2, 0xca, 0x6b, 0x50, 0x6d, 0xb2, 0xa7, 0x2c, 0x7a, 0xa9, 0xe4, 0x07, 0xd0, 0xaa,
	0x2d, 0x46, 0xe3, 0xc3, 0xcc, 0x9a, 0xc5, 0xfe, 0x67, 0x87, 0x29, 0x0c, 0x1e, 0xe8, 0xe3, 0x03,
	0x49, 0x1e, 0xf9, 0x0e, 0x9d, 0xd6, 0xc4, 0x7c, 0x8d, 0xff, 0x96, 0x8e, 0x75, 0x01, 0x6f, 0xa3,
	0x28, 0x62, 0x26, 0x7b, 0x66, 0x9c, 0x3c, 0x31, 0x93, 0xb2, 0x8a, 0x68, 0xa9, 0xb6, 0xe5, 0xe6,
	0x3f, 0x83, 0x2b, 0x05, 0x7d, 0xf1, 0x45, 0x5c, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xda, 0x69,
	0x57, 0x6d, 0x2e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlockDatabaseClient is the client API for BlockDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockDatabaseClient interface {
	// Return db[UserID]
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Set db[UserID]=Value
	Put(ctx context.Context, in *Request, opts ...grpc.CallOption) (*BooleanResponse, error)
	// Perform db[UserID]+=Value
	Withdraw(ctx context.Context, in *Request, opts ...grpc.CallOption) (*BooleanResponse, error)
	// Perform db[UserID]-=Value
	// Return Success=false if balance is insufficient
	Deposit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*BooleanResponse, error)
	// Perform db[FromID]-=Value and db[ToID]+=Value
	// Return Success=false if FromID is same as ToID or balance of FromID is insufficient
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*BooleanResponse, error)
	// Return the length of transient (non-block) log on disk
	LogLength(ctx context.Context, in *Null, opts ...grpc.CallOption) (*GetResponse, error)
}

type blockDatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockDatabaseClient(cc grpc.ClientConnInterface) BlockDatabaseClient {
	return &blockDatabaseClient{cc}
}

func (c *blockDatabaseClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/blockdb.BlockDatabase/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockDatabaseClient) Put(ctx context.Context, in *Request, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/blockdb.BlockDatabase/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockDatabaseClient) Withdraw(ctx context.Context, in *Request, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/blockdb.BlockDatabase/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockDatabaseClient) Deposit(ctx context.Context, in *Request, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/blockdb.BlockDatabase/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockDatabaseClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/blockdb.BlockDatabase/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockDatabaseClient) LogLength(ctx context.Context, in *Null, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/blockdb.BlockDatabase/LogLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockDatabaseServer is the server API for BlockDatabase service.
type BlockDatabaseServer interface {
	// Return db[UserID]
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Set db[UserID]=Value
	Put(context.Context, *Request) (*BooleanResponse, error)
	// Perform db[UserID]+=Value
	Withdraw(context.Context, *Request) (*BooleanResponse, error)
	// Perform db[UserID]-=Value
	// Return Success=false if balance is insufficient
	Deposit(context.Context, *Request) (*BooleanResponse, error)
	// Perform db[FromID]-=Value and db[ToID]+=Value
	// Return Success=false if FromID is same as ToID or balance of FromID is insufficient
	Transfer(context.Context, *TransferRequest) (*BooleanResponse, error)
	// Return the length of transient (non-block) log on disk
	LogLength(context.Context, *Null) (*GetResponse, error)
}

// UnimplementedBlockDatabaseServer can be embedded to have forward compatible implementations.
type UnimplementedBlockDatabaseServer struct {
}

func (*UnimplementedBlockDatabaseServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBlockDatabaseServer) Put(ctx context.Context, req *Request) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedBlockDatabaseServer) Withdraw(ctx context.Context, req *Request) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (*UnimplementedBlockDatabaseServer) Deposit(ctx context.Context, req *Request) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedBlockDatabaseServer) Transfer(ctx context.Context, req *TransferRequest) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (*UnimplementedBlockDatabaseServer) LogLength(ctx context.Context, req *Null) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogLength not implemented")
}

func RegisterBlockDatabaseServer(s *grpc.Server, srv BlockDatabaseServer) {
	s.RegisterService(&_BlockDatabase_serviceDesc, srv)
}

func _BlockDatabase_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockDatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockdb.BlockDatabase/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockDatabaseServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockDatabase_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockDatabaseServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockdb.BlockDatabase/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockDatabaseServer).Put(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockDatabase_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockDatabaseServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockdb.BlockDatabase/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockDatabaseServer).Withdraw(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockDatabase_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockDatabaseServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockdb.BlockDatabase/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockDatabaseServer).Deposit(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockDatabase_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockDatabaseServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockdb.BlockDatabase/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockDatabaseServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockDatabase_LogLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockDatabaseServer).LogLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockdb.BlockDatabase/LogLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockDatabaseServer).LogLength(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockDatabase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blockdb.BlockDatabase",
	HandlerType: (*BlockDatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BlockDatabase_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _BlockDatabase_Put_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _BlockDatabase_Withdraw_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _BlockDatabase_Deposit_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _BlockDatabase_Transfer_Handler,
		},
		{
			MethodName: "LogLength",
			Handler:    _BlockDatabase_LogLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}
