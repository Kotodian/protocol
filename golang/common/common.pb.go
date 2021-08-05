// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 支付类型
type PayMode int32

const (
	PayMode_PM_Unknown     PayMode = 0
	PayMode_PM_Online      PayMode = 1
	PayMode_PM_Group       PayMode = 2
	PayMode_PM_IdCard      PayMode = 3
	PayMode_PM_UserBalance PayMode = 4
	PayMode_PM_WalletCard  PayMode = 5
	PayMode_PM_Manul       PayMode = 6
	PayMode_PM_VIN         PayMode = 7
	PayMode_PM_Partner     PayMode = 1000
)

var PayMode_name = map[int32]string{
	0:    "PM_Unknown",
	1:    "PM_Online",
	2:    "PM_Group",
	3:    "PM_IdCard",
	4:    "PM_UserBalance",
	5:    "PM_WalletCard",
	6:    "PM_Manul",
	7:    "PM_VIN",
	1000: "PM_Partner",
}

var PayMode_value = map[string]int32{
	"PM_Unknown":     0,
	"PM_Online":      1,
	"PM_Group":       2,
	"PM_IdCard":      3,
	"PM_UserBalance": 4,
	"PM_WalletCard":  5,
	"PM_Manul":       6,
	"PM_VIN":         7,
	"PM_Partner":     1000,
}

func (x PayMode) String() string {
	return proto.EnumName(PayMode_name, int32(x))
}

func (PayMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// 支付渠道，一个支付类型对应多个支付渠道
type PayChannel int32

const (
	PayChannel_PC_Unknown     PayChannel = 0
	PayChannel_PC_AliPay      PayChannel = 1
	PayChannel_PC_WechatPay   PayChannel = 2
	PayChannel_PC_UserBalance PayChannel = 3
	PayChannel_PC_CardBalance PayChannel = 4
	PayChannel_PC_Group       PayChannel = 5
	PayChannel_PC_Bank        PayChannel = 6
	PayChannel_PC_WalletCard  PayChannel = 7
	PayChannel_PC_Manul       PayChannel = 8
	PayChannel_PC_VIN         PayChannel = 15
)

var PayChannel_name = map[int32]string{
	0:  "PC_Unknown",
	1:  "PC_AliPay",
	2:  "PC_WechatPay",
	3:  "PC_UserBalance",
	4:  "PC_CardBalance",
	5:  "PC_Group",
	6:  "PC_Bank",
	7:  "PC_WalletCard",
	8:  "PC_Manul",
	15: "PC_VIN",
}

var PayChannel_value = map[string]int32{
	"PC_Unknown":     0,
	"PC_AliPay":      1,
	"PC_WechatPay":   2,
	"PC_UserBalance": 3,
	"PC_CardBalance": 4,
	"PC_Group":       5,
	"PC_Bank":        6,
	"PC_WalletCard":  7,
	"PC_Manul":       8,
	"PC_VIN":         15,
}

func (x PayChannel) String() string {
	return proto.EnumName(PayChannel_name, int32(x))
}

func (PayChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type RefundMode int32

const (
	RefundMode_RM_Unknown     RefundMode = 0
	RefundMode_RM_BackTrack   RefundMode = 1
	RefundMode_RM_UserBalance RefundMode = 2
)

var RefundMode_name = map[int32]string{
	0: "RM_Unknown",
	1: "RM_BackTrack",
	2: "RM_UserBalance",
}

var RefundMode_value = map[string]int32{
	"RM_Unknown":     0,
	"RM_BackTrack":   1,
	"RM_UserBalance": 2,
}

func (x RefundMode) String() string {
	return proto.EnumName(RefundMode_name, int32(x))
}

func (RefundMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type AuthModeChannel int32

const (
	AuthModeChannel_AMC_Unknown            AuthModeChannel = 0
	AuthModeChannel_AMC_App                AuthModeChannel = 1
	AuthModeChannel_AMC_WechatService      AuthModeChannel = 2
	AuthModeChannel_AMC_ALiService         AuthModeChannel = 3
	AuthModeChannel_AMC_WechatSmallRoutine AuthModeChannel = 4
	AuthModeChannel_AMC_AliSmallRoutine    AuthModeChannel = 5
	AuthModeChannel_AMC_LocalPlug          AuthModeChannel = 6
	AuthModeChannel_AMC_LocalAdmin         AuthModeChannel = 7
	AuthModeChannel_AMC_LocalIdentityCard  AuthModeChannel = 8
	AuthModeChannel_AMC_OnlineIdentityCard AuthModeChannel = 9
	AuthModeChannel_AMC_LocalWalletCard    AuthModeChannel = 10
	AuthModeChannel_AMC_LocalVIN           AuthModeChannel = 11
	AuthModeChannel_AMC_OnlineVIN          AuthModeChannel = 12
	AuthModeChannel_AMC_Bluetooth          AuthModeChannel = 13
	AuthModeChannel_AMC_MAC                AuthModeChannel = 14
	AuthModeChannel_AMC_RemoteAdmin        AuthModeChannel = 15
	AuthModeChannel_AMC_Partner            AuthModeChannel = 1000
)

var AuthModeChannel_name = map[int32]string{
	0:    "AMC_Unknown",
	1:    "AMC_App",
	2:    "AMC_WechatService",
	3:    "AMC_ALiService",
	4:    "AMC_WechatSmallRoutine",
	5:    "AMC_AliSmallRoutine",
	6:    "AMC_LocalPlug",
	7:    "AMC_LocalAdmin",
	8:    "AMC_LocalIdentityCard",
	9:    "AMC_OnlineIdentityCard",
	10:   "AMC_LocalWalletCard",
	11:   "AMC_LocalVIN",
	12:   "AMC_OnlineVIN",
	13:   "AMC_Bluetooth",
	14:   "AMC_MAC",
	15:   "AMC_RemoteAdmin",
	1000: "AMC_Partner",
}

var AuthModeChannel_value = map[string]int32{
	"AMC_Unknown":            0,
	"AMC_App":                1,
	"AMC_WechatService":      2,
	"AMC_ALiService":         3,
	"AMC_WechatSmallRoutine": 4,
	"AMC_AliSmallRoutine":    5,
	"AMC_LocalPlug":          6,
	"AMC_LocalAdmin":         7,
	"AMC_LocalIdentityCard":  8,
	"AMC_OnlineIdentityCard": 9,
	"AMC_LocalWalletCard":    10,
	"AMC_LocalVIN":           11,
	"AMC_OnlineVIN":          12,
	"AMC_Bluetooth":          13,
	"AMC_MAC":                14,
	"AMC_RemoteAdmin":        15,
	"AMC_Partner":            1000,
}

func (x AuthModeChannel) String() string {
	return proto.EnumName(AuthModeChannel_name, int32(x))
}

func (AuthModeChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type CustomerRegChannel int32

const (
	CustomerRegChannel_CRC_Unknown            CustomerRegChannel = 0
	CustomerRegChannel_CRC_App                CustomerRegChannel = 1
	CustomerRegChannel_CRC_WechatService      CustomerRegChannel = 2
	CustomerRegChannel_CRC_ALiService         CustomerRegChannel = 3
	CustomerRegChannel_CRC_WechatSmallRoutine CustomerRegChannel = 4
	CustomerRegChannel_CRC_AliSmallRoutine    CustomerRegChannel = 5
	CustomerRegChannel_CRC_Partner            CustomerRegChannel = 1000
)

var CustomerRegChannel_name = map[int32]string{
	0:    "CRC_Unknown",
	1:    "CRC_App",
	2:    "CRC_WechatService",
	3:    "CRC_ALiService",
	4:    "CRC_WechatSmallRoutine",
	5:    "CRC_AliSmallRoutine",
	1000: "CRC_Partner",
}

var CustomerRegChannel_value = map[string]int32{
	"CRC_Unknown":            0,
	"CRC_App":                1,
	"CRC_WechatService":      2,
	"CRC_ALiService":         3,
	"CRC_WechatSmallRoutine": 4,
	"CRC_AliSmallRoutine":    5,
	"CRC_Partner":            1000,
}

func (x CustomerRegChannel) String() string {
	return proto.EnumName(CustomerRegChannel_name, int32(x))
}

func (CustomerRegChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

type Resp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Obj                  []byte   `protobuf:"bytes,3,opt,name=obj,proto3" json:"obj,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Resp) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Resp) GetObj() []byte {
	if m != nil {
		return m.Obj
	}
	return nil
}

type PageReq struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageReq) Reset()         { *m = PageReq{} }
func (m *PageReq) String() string { return proto.CompactTextString(m) }
func (*PageReq) ProtoMessage()    {}
func (*PageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *PageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageReq.Unmarshal(m, b)
}
func (m *PageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageReq.Marshal(b, m, deterministic)
}
func (m *PageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageReq.Merge(m, src)
}
func (m *PageReq) XXX_Size() int {
	return xxx_messageInfo_PageReq.Size(m)
}
func (m *PageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PageReq.DiscardUnknown(m)
}

var xxx_messageInfo_PageReq proto.InternalMessageInfo

func (m *PageReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PageReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type PageResp struct {
	TotalPage            int32    `protobuf:"varint,1,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	TotalRows            int32    `protobuf:"varint,2,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageResp) Reset()         { *m = PageResp{} }
func (m *PageResp) String() string { return proto.CompactTextString(m) }
func (*PageResp) ProtoMessage()    {}
func (*PageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *PageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageResp.Unmarshal(m, b)
}
func (m *PageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageResp.Marshal(b, m, deterministic)
}
func (m *PageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResp.Merge(m, src)
}
func (m *PageResp) XXX_Size() int {
	return xxx_messageInfo_PageResp.Size(m)
}
func (m *PageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PageResp.DiscardUnknown(m)
}

var xxx_messageInfo_PageResp proto.InternalMessageInfo

func (m *PageResp) GetTotalPage() int32 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *PageResp) GetTotalRows() int32 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

type SearchFilter struct {
	FilterName           string   `protobuf:"bytes,1,opt,name=filter_name,json=filterName,proto3" json:"filter_name,omitempty"`
	FilterVal            string   `protobuf:"bytes,2,opt,name=filter_val,json=filterVal,proto3" json:"filter_val,omitempty"`
	FilterOp             string   `protobuf:"bytes,3,opt,name=filter_op,json=filterOp,proto3" json:"filter_op,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchFilter) Reset()         { *m = SearchFilter{} }
func (m *SearchFilter) String() string { return proto.CompactTextString(m) }
func (*SearchFilter) ProtoMessage()    {}
func (*SearchFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *SearchFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchFilter.Unmarshal(m, b)
}
func (m *SearchFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchFilter.Marshal(b, m, deterministic)
}
func (m *SearchFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchFilter.Merge(m, src)
}
func (m *SearchFilter) XXX_Size() int {
	return xxx_messageInfo_SearchFilter.Size(m)
}
func (m *SearchFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchFilter.DiscardUnknown(m)
}

var xxx_messageInfo_SearchFilter proto.InternalMessageInfo

func (m *SearchFilter) GetFilterName() string {
	if m != nil {
		return m.FilterName
	}
	return ""
}

func (m *SearchFilter) GetFilterVal() string {
	if m != nil {
		return m.FilterVal
	}
	return ""
}

func (m *SearchFilter) GetFilterOp() string {
	if m != nil {
		return m.FilterOp
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.PayMode", PayMode_name, PayMode_value)
	proto.RegisterEnum("common.PayChannel", PayChannel_name, PayChannel_value)
	proto.RegisterEnum("common.RefundMode", RefundMode_name, RefundMode_value)
	proto.RegisterEnum("common.AuthModeChannel", AuthModeChannel_name, AuthModeChannel_value)
	proto.RegisterEnum("common.CustomerRegChannel", CustomerRegChannel_name, CustomerRegChannel_value)
	proto.RegisterType((*Resp)(nil), "common.Resp")
	proto.RegisterType((*PageReq)(nil), "common.PageReq")
	proto.RegisterType((*PageResp)(nil), "common.PageResp")
	proto.RegisterType((*SearchFilter)(nil), "common.SearchFilter")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdf, 0x4e, 0xdb, 0x3c,
	0x1c, 0xfd, 0xd2, 0xff, 0xfd, 0xb5, 0x50, 0x63, 0x3e, 0xbe, 0x8f, 0x6d, 0x9a, 0x86, 0x7a, 0x85,
	0x7a, 0x41, 0xa5, 0x21, 0xed, 0x72, 0x52, 0x1b, 0x69, 0x03, 0x89, 0x40, 0x64, 0x36, 0x90, 0x76,
	0x13, 0x99, 0xd4, 0xb4, 0xa1, 0x8e, 0x1d, 0x12, 0x07, 0xd4, 0x27, 0xd9, 0x0b, 0xec, 0x76, 0xef,
	0xb5, 0xeb, 0x3d, 0xc1, 0x64, 0x3b, 0x29, 0x29, 0xec, 0xce, 0xbf, 0xf3, 0x3b, 0x3e, 0x3e, 0xe7,
	0xb4, 0x0a, 0xf4, 0x43, 0x19, 0xc7, 0x52, 0x1c, 0x25, 0xa9, 0x54, 0x12, 0xb7, 0xec, 0x34, 0xfc,
	0x08, 0x0d, 0xc2, 0xb2, 0x04, 0x63, 0x68, 0x84, 0x72, 0xc6, 0xf6, 0x9d, 0x03, 0xe7, 0xb0, 0x49,
	0xcc, 0x19, 0x23, 0xa8, 0xb3, 0x34, 0xdd, 0xaf, 0x1d, 0x38, 0x87, 0x5d, 0xa2, 0x8f, 0x1a, 0x91,
	0x37, 0x77, 0xfb, 0xf5, 0x03, 0xe7, 0xb0, 0x4f, 0xf4, 0x71, 0x78, 0x0c, 0x6d, 0x9f, 0xce, 0x19,
	0x61, 0xf7, 0x5a, 0x22, 0xa1, 0xf3, 0xb5, 0x84, 0x3e, 0xe3, 0x7f, 0xa1, 0xc9, 0xa3, 0x38, 0x52,
	0x46, 0xa4, 0x49, 0xec, 0x30, 0x3c, 0x81, 0x8e, 0xbd, 0x94, 0x25, 0xf8, 0x2d, 0x80, 0x92, 0x8a,
	0xf2, 0xa0, 0x72, 0xb7, 0x6b, 0x10, 0x4d, 0x79, 0x5a, 0xa7, 0xf2, 0x31, 0x2b, 0x54, 0xec, 0x9a,
	0xc8, 0xc7, 0x6c, 0xb8, 0x84, 0xfe, 0x25, 0xa3, 0x69, 0xb8, 0xf8, 0x14, 0x71, 0xc5, 0x52, 0xfc,
	0x0e, 0x7a, 0xb7, 0xe6, 0x14, 0x08, 0x1a, 0x5b, 0xb9, 0x2e, 0x01, 0x0b, 0x9d, 0xd3, 0xd8, 0xe8,
	0x15, 0x84, 0x07, 0xca, 0x8b, 0x68, 0x5d, 0x8b, 0x5c, 0x51, 0x8e, 0xdf, 0x40, 0x31, 0x04, 0x32,
	0x31, 0x31, 0xbb, 0xa4, 0x63, 0x81, 0x8b, 0x64, 0xf4, 0xdd, 0xd1, 0x61, 0x57, 0x9e, 0xee, 0x66,
	0x1b, 0xc0, 0xf7, 0x82, 0xaf, 0x62, 0x29, 0xe4, 0xa3, 0x40, 0xff, 0xe0, 0x2d, 0xe8, 0xfa, 0x5e,
	0x70, 0x21, 0x78, 0x24, 0x18, 0x72, 0x70, 0x1f, 0x3a, 0xbe, 0x17, 0x7c, 0x4e, 0x65, 0x9e, 0xa0,
	0x5a, 0xb1, 0x3c, 0x9d, 0xb9, 0x34, 0x9d, 0xa1, 0x3a, 0xc6, 0xb0, 0xad, 0xef, 0x66, 0x2c, 0x9d,
	0x52, 0x4e, 0x45, 0xc8, 0x50, 0x03, 0xef, 0xc0, 0x96, 0xef, 0x05, 0xd7, 0x94, 0x73, 0xa6, 0x0c,
	0xad, 0x59, 0x68, 0x78, 0x54, 0xe4, 0x1c, 0xb5, 0x30, 0x40, 0xcb, 0xf7, 0x82, 0xab, 0xd3, 0x73,
	0xd4, 0xc6, 0x03, 0xf3, 0xb8, 0x4f, 0x53, 0x25, 0x58, 0x8a, 0x7e, 0xb5, 0x47, 0x3f, 0x1d, 0x00,
	0x9f, 0xae, 0xdc, 0x05, 0x15, 0x82, 0x71, 0x63, 0xce, 0x7d, 0x66, 0xce, 0x0d, 0x26, 0x3c, 0xf2,
	0xe9, 0x0a, 0x39, 0x18, 0x41, 0xdf, 0x77, 0x83, 0x6b, 0x16, 0x2e, 0xa8, 0xd2, 0x48, 0xcd, 0x38,
	0x72, 0x37, 0x1c, 0xd5, 0x0b, 0x4c, 0x7b, 0x79, 0x72, 0xa9, 0x2d, 0xb9, 0x45, 0xac, 0x26, 0xee,
	0x41, 0xdb, 0x77, 0x83, 0x29, 0x15, 0x4b, 0xd4, 0x32, 0x01, 0xdc, 0x6a, 0x80, 0x76, 0xc1, 0xb6,
	0x01, 0x3a, 0x26, 0x80, 0x6b, 0x02, 0x0c, 0x46, 0x53, 0x00, 0xc2, 0x6e, 0x73, 0x31, 0x2b, 0xbb,
	0x24, 0xd5, 0x2e, 0x11, 0xf4, 0x89, 0x17, 0x4c, 0x69, 0xb8, 0xfc, 0x92, 0xd2, 0x70, 0x89, 0x1c,
	0xed, 0x85, 0x6c, 0x36, 0x56, 0x1b, 0xfd, 0xae, 0xc1, 0x60, 0x92, 0xab, 0x85, 0x96, 0x28, 0x83,
	0x0f, 0xa0, 0x37, 0xf1, 0xaa, 0xc9, 0x7b, 0xd0, 0xd6, 0xc0, 0x24, 0x49, 0x90, 0x83, 0xf7, 0x60,
	0x47, 0x0f, 0x36, 0xf8, 0x25, 0x4b, 0x1f, 0x22, 0x2d, 0xa4, 0xc5, 0x0d, 0xe7, 0x2c, 0x2a, 0xb1,
	0x3a, 0x7e, 0x0d, 0xff, 0x55, 0xa8, 0x31, 0xe5, 0x9c, 0xc8, 0x5c, 0xe9, 0xdf, 0xb6, 0x81, 0xff,
	0x87, 0x5d, 0xc3, 0xe7, 0xd1, 0xc6, 0xa2, 0xa9, 0x2b, 0xd0, 0x8b, 0x33, 0x19, 0x52, 0xee, 0xf3,
	0x7c, 0x8e, 0x5a, 0xa5, 0xb6, 0x81, 0x26, 0xb3, 0x38, 0x12, 0xa8, 0x8d, 0x5f, 0xc1, 0xde, 0x1a,
	0x3b, 0x9d, 0x31, 0xa1, 0x22, 0xb5, 0x32, 0x8d, 0x75, 0xca, 0x67, 0xed, 0xdf, 0x68, 0x63, 0xd7,
	0x2d, 0x9f, 0x35, 0xd7, 0x2a, 0x35, 0x83, 0xae, 0x6b, 0xbd, 0xd0, 0xf5, 0xf6, 0x4a, 0x23, 0x56,
	0x46, 0x43, 0xfd, 0x12, 0x9a, 0xf2, 0x9c, 0x29, 0x29, 0xd5, 0x02, 0x6d, 0x95, 0xdd, 0x78, 0x13,
	0x17, 0x6d, 0xe3, 0x5d, 0x18, 0xe8, 0x81, 0xb0, 0x58, 0x2a, 0x66, 0x9d, 0x0e, 0x30, 0xb2, 0x75,
	0x56, 0xfe, 0x68, 0x3f, 0x1c, 0xc0, 0x6e, 0x9e, 0x29, 0x19, 0xb3, 0x94, 0xb0, 0x79, 0xa5, 0x77,
	0x97, 0x3c, 0xeb, 0x5d, 0x03, 0xeb, 0xde, 0xf5, 0xf0, 0x97, 0xde, 0x0d, 0xe7, 0x59, 0xef, 0x15,
	0xea, 0x8b, 0xde, 0x0d, 0xff, 0x45, 0xef, 0xc8, 0xbe, 0xfe, 0x64, 0x73, 0xfa, 0xe1, 0xa4, 0xfe,
	0xed, 0xfd, 0x3c, 0x52, 0x61, 0x16, 0x67, 0x47, 0x77, 0x72, 0x95, 0x49, 0x71, 0x9f, 0x47, 0xe2,
	0x28, 0x94, 0xf1, 0x58, 0x63, 0x63, 0xf3, 0x11, 0x0c, 0x25, 0x1f, 0xcf, 0x25, 0xa7, 0x62, 0x3e,
	0xb6, 0x5f, 0xc3, 0x9b, 0x96, 0xc1, 0x8f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xbf, 0xb3,
	0x20, 0x2c, 0x05, 0x00, 0x00,
}
