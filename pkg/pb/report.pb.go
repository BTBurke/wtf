// Code generated by protoc-gen-go. DO NOT EDIT.
// source: report.proto

package pb

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

type ReportReason int32

const (
	ReportReason_Unknown        ReportReason = 0
	ReportReason_Success        ReportReason = 1
	ReportReason_Failure        ReportReason = 2
	ReportReason_Alert          ReportReason = 3
	ReportReason_AlertRate      ReportReason = 4
	ReportReason_MemoryWarning  ReportReason = 5
	ReportReason_TimeWarning    ReportReason = 6
	ReportReason_FileNotCreated ReportReason = 7
	ReportReason_Killed         ReportReason = 8
	ReportReason_Start          ReportReason = 9
)

var ReportReason_name = map[int32]string{
	0: "Unknown",
	1: "Success",
	2: "Failure",
	3: "Alert",
	4: "AlertRate",
	5: "MemoryWarning",
	6: "TimeWarning",
	7: "FileNotCreated",
	8: "Killed",
	9: "Start",
}

var ReportReason_value = map[string]int32{
	"Unknown":        0,
	"Success":        1,
	"Failure":        2,
	"Alert":          3,
	"AlertRate":      4,
	"MemoryWarning":  5,
	"TimeWarning":    6,
	"FileNotCreated": 7,
	"Killed":         8,
	"Start":          9,
}

func (x ReportReason) String() string {
	return proto.EnumName(ReportReason_name, int32(x))
}

func (ReportReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{0}
}

type KillReason int32

const (
	KillReason_NotKilled KillReason = 0
	KillReason_Timeout   KillReason = 1
	KillReason_Memory    KillReason = 2
	KillReason_Signal    KillReason = 3
)

var KillReason_name = map[int32]string{
	0: "NotKilled",
	1: "Timeout",
	2: "Memory",
	3: "Signal",
}

var KillReason_value = map[string]int32{
	"NotKilled": 0,
	"Timeout":   1,
	"Memory":    2,
	"Signal":    3,
}

func (x KillReason) String() string {
	return proto.EnumName(KillReason_name, int32(x))
}

func (KillReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{1}
}

type Report struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname             string       `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Stdout               []string     `protobuf:"bytes,3,rep,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []string     `protobuf:"bytes,4,rep,name=stderr,proto3" json:"stderr,omitempty"`
	Success              bool         `protobuf:"varint,5,opt,name=success,proto3" json:"success,omitempty"`
	MaxMemory            uint64       `protobuf:"varint,6,opt,name=max_memory,json=maxMemory,proto3" json:"max_memory,omitempty"`
	Killed               bool         `protobuf:"varint,7,opt,name=killed,proto3" json:"killed,omitempty"`
	KillReason           KillReason   `protobuf:"varint,8,opt,name=kill_reason,json=killReason,proto3,enum=monny.monitor.KillReason" json:"kill_reason,omitempty"`
	Created              []byte       `protobuf:"bytes,9,opt,name=created,proto3" json:"created,omitempty"`
	ReportReason         ReportReason `protobuf:"varint,10,opt,name=report_reason,json=reportReason,proto3,enum=monny.monitor.ReportReason" json:"report_reason,omitempty"`
	Start                int64        `protobuf:"varint,11,opt,name=start,proto3" json:"start,omitempty"`
	Finish               int64        `protobuf:"varint,12,opt,name=finish,proto3" json:"finish,omitempty"`
	Duration             string       `protobuf:"bytes,13,opt,name=duration,proto3" json:"duration,omitempty"`
	ExitCode             int32        `protobuf:"varint,14,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	ExitCodeValid        bool         `protobuf:"varint,15,opt,name=exit_code_valid,json=exitCodeValid,proto3" json:"exit_code_valid,omitempty"`
	Messages             []string     `protobuf:"bytes,16,rep,name=messages,proto3" json:"messages,omitempty"`
	Matches              []byte       `protobuf:"bytes,17,opt,name=matches,proto3" json:"matches,omitempty"`
	UserCommand          string       `protobuf:"bytes,18,opt,name=user_command,json=userCommand,proto3" json:"user_command,omitempty"`
	Config               []byte       `protobuf:"bytes,19,opt,name=config,proto3" json:"config,omitempty"`
	CreatedAt            int64        `protobuf:"varint,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Report) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Report) GetStdout() []string {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Report) GetStderr() []string {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *Report) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Report) GetMaxMemory() uint64 {
	if m != nil {
		return m.MaxMemory
	}
	return 0
}

func (m *Report) GetKilled() bool {
	if m != nil {
		return m.Killed
	}
	return false
}

func (m *Report) GetKillReason() KillReason {
	if m != nil {
		return m.KillReason
	}
	return KillReason_NotKilled
}

func (m *Report) GetCreated() []byte {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Report) GetReportReason() ReportReason {
	if m != nil {
		return m.ReportReason
	}
	return ReportReason_Unknown
}

func (m *Report) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Report) GetFinish() int64 {
	if m != nil {
		return m.Finish
	}
	return 0
}

func (m *Report) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Report) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *Report) GetExitCodeValid() bool {
	if m != nil {
		return m.ExitCodeValid
	}
	return false
}

func (m *Report) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Report) GetMatches() []byte {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *Report) GetUserCommand() string {
	if m != nil {
		return m.UserCommand
	}
	return ""
}

func (m *Report) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Report) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type ReportAck struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportAck) Reset()         { *m = ReportAck{} }
func (m *ReportAck) String() string { return proto.CompactTextString(m) }
func (*ReportAck) ProtoMessage()    {}
func (*ReportAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{1}
}

func (m *ReportAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportAck.Unmarshal(m, b)
}
func (m *ReportAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportAck.Marshal(b, m, deterministic)
}
func (m *ReportAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportAck.Merge(m, src)
}
func (m *ReportAck) XXX_Size() int {
	return xxx_messageInfo_ReportAck.Size(m)
}
func (m *ReportAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportAck.DiscardUnknown(m)
}

var xxx_messageInfo_ReportAck proto.InternalMessageInfo

func (m *ReportAck) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("monny.monitor.ReportReason", ReportReason_name, ReportReason_value)
	proto.RegisterEnum("monny.monitor.KillReason", KillReason_name, KillReason_value)
	proto.RegisterType((*Report)(nil), "monny.monitor.Report")
	proto.RegisterType((*ReportAck)(nil), "monny.monitor.ReportAck")
}

func init() { proto.RegisterFile("report.proto", fileDescriptor_3eedb623aa6ca98c) }

var fileDescriptor_3eedb623aa6ca98c = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4b, 0x6f, 0x13, 0x31,
	0x10, 0xee, 0xe6, 0xb1, 0xc9, 0x4e, 0x1e, 0x75, 0x4d, 0x41, 0xa6, 0x15, 0xd2, 0x52, 0x09, 0xb4,
	0xea, 0x21, 0x48, 0xe5, 0x06, 0x97, 0x86, 0x4a, 0xbd, 0x54, 0xf4, 0xb0, 0xe5, 0x21, 0x71, 0x89,
	0xdc, 0x5d, 0x37, 0xb5, 0xb2, 0xb6, 0x23, 0xdb, 0x81, 0xf6, 0xcf, 0xf0, 0x0b, 0xf9, 0x11, 0xc8,
	0x8f, 0x0d, 0x2d, 0xea, 0x6d, 0xbe, 0x6f, 0xc7, 0xdf, 0xcc, 0x7c, 0x33, 0x0b, 0x63, 0xcd, 0xd6,
	0x4a, 0xdb, 0xd9, 0x5a, 0x2b, 0xab, 0xf0, 0x44, 0x28, 0x29, 0xef, 0x67, 0x42, 0x49, 0x6e, 0x95,
	0x3e, 0xfa, 0xd3, 0x83, 0xb4, 0xf4, 0xdf, 0xf1, 0x14, 0x3a, 0xbc, 0x26, 0x49, 0x9e, 0x14, 0x59,
	0xd9, 0xe1, 0x35, 0x3e, 0x80, 0xe1, 0xad, 0x32, 0x56, 0x52, 0xc1, 0x48, 0xc7, 0xb3, 0x5b, 0x8c,
	0x5f, 0x40, 0x6a, 0x6c, 0xad, 0x36, 0x96, 0x74, 0xf3, 0x6e, 0x91, 0x95, 0x11, 0x45, 0x9e, 0x69,
	0x4d, 0x7a, 0x5b, 0x9e, 0x69, 0x8d, 0x09, 0x0c, 0xcc, 0xa6, 0xaa, 0x98, 0x31, 0xa4, 0x9f, 0x27,
	0xc5, 0xb0, 0x6c, 0x21, 0x7e, 0x05, 0x20, 0xe8, 0xdd, 0x42, 0x30, 0xa1, 0xf4, 0x3d, 0x49, 0xf3,
	0xa4, 0xe8, 0x95, 0x99, 0xa0, 0x77, 0x9f, 0x3d, 0xe1, 0x04, 0x57, 0xbc, 0x69, 0x58, 0x4d, 0x06,
	0xfe, 0x5d, 0x44, 0xf8, 0x03, 0x8c, 0x5c, 0xb4, 0xd0, 0x8c, 0x1a, 0x25, 0xc9, 0x30, 0x4f, 0x8a,
	0xe9, 0xc9, 0xcb, 0xd9, 0xa3, 0xe1, 0x66, 0x17, 0xbc, 0x69, 0x4a, 0x9f, 0x50, 0xc2, 0x6a, 0x1b,
	0xbb, 0x66, 0x2a, 0xcd, 0xa8, 0x65, 0x35, 0xc9, 0xf2, 0xa4, 0x18, 0x97, 0x2d, 0xc4, 0xa7, 0x30,
	0x09, 0x66, 0xb5, 0xba, 0xe0, 0x75, 0x0f, 0xff, 0xd3, 0x0d, 0x86, 0x45, 0xe5, 0x68, 0x6f, 0xd4,
	0xde, 0x87, 0xbe, 0xb1, 0x54, 0x5b, 0x32, 0xca, 0x93, 0xa2, 0x5b, 0x06, 0xe0, 0xa6, 0xb8, 0xe1,
	0x92, 0x9b, 0x5b, 0x32, 0xf6, 0x74, 0x44, 0xce, 0xe2, 0x7a, 0xa3, 0xa9, 0xe5, 0x4a, 0x92, 0x49,
	0xb0, 0xb8, 0xc5, 0xf8, 0x10, 0x32, 0x76, 0xc7, 0xed, 0xa2, 0x52, 0x35, 0x23, 0xd3, 0x3c, 0x29,
	0xfa, 0xe5, 0xd0, 0x11, 0x67, 0xaa, 0x66, 0xf8, 0x2d, 0xec, 0x6e, 0x3f, 0x2e, 0x7e, 0xd2, 0x86,
	0xd7, 0x64, 0xd7, 0xfb, 0x33, 0x69, 0x53, 0xbe, 0x39, 0xd2, 0x15, 0x10, 0xcc, 0x18, 0xba, 0x64,
	0x86, 0x20, 0xbf, 0x91, 0x2d, 0x76, 0x36, 0x08, 0x6a, 0xab, 0x5b, 0x66, 0xc8, 0x5e, 0xb0, 0x21,
	0x42, 0xfc, 0x1a, 0xc6, 0x1b, 0xc3, 0xf4, 0xa2, 0x52, 0x42, 0x50, 0x59, 0x13, 0xec, 0x5b, 0x1b,
	0x39, 0xee, 0x2c, 0x50, 0x6e, 0xa2, 0x4a, 0xc9, 0x1b, 0xbe, 0x24, 0xcf, 0xfc, 0xdb, 0x88, 0xdc,
	0x3a, 0xa3, 0x99, 0x0b, 0x6a, 0xc9, 0xbe, 0x9f, 0x36, 0x8b, 0xcc, 0xdc, 0x1e, 0xbd, 0x81, 0x2c,
	0x98, 0x37, 0xaf, 0x56, 0x0f, 0x8f, 0x22, 0x79, 0x74, 0x14, 0xc7, 0xbf, 0x13, 0x18, 0x3f, 0x34,
	0x19, 0x8f, 0x60, 0xf0, 0x55, 0xae, 0xa4, 0xfa, 0x25, 0xd1, 0x8e, 0x03, 0x57, 0x21, 0x11, 0x25,
	0x0e, 0x9c, 0x53, 0xde, 0x6c, 0x34, 0x43, 0x1d, 0x9c, 0x41, 0x7f, 0xde, 0x30, 0x6d, 0x51, 0x17,
	0x4f, 0x20, 0xf3, 0x61, 0x49, 0x2d, 0x43, 0x3d, 0xbc, 0x07, 0x93, 0x70, 0x51, 0xdf, 0xa9, 0x96,
	0x5c, 0x2e, 0x51, 0x1f, 0xef, 0xc2, 0xe8, 0x0b, 0x17, 0xac, 0x25, 0x52, 0x8c, 0x61, 0x7a, 0xce,
	0x1b, 0x76, 0xa9, 0xec, 0x59, 0x68, 0x18, 0x0d, 0x30, 0x40, 0x7a, 0xe1, 0x2f, 0x0e, 0x0d, 0x9d,
	0xfa, 0x95, 0x5b, 0x27, 0xca, 0x8e, 0x4f, 0x01, 0xfe, 0x1d, 0x97, 0xab, 0x75, 0xa9, 0x6c, 0xcc,
	0xf3, 0xfd, 0x39, 0x61, 0xb5, 0xb1, 0x28, 0x71, 0x02, 0xa1, 0x30, 0xea, 0xb8, 0xf8, 0x8a, 0x2f,
	0x25, 0x6d, 0x50, 0xf7, 0xe4, 0x1c, 0x06, 0x61, 0x42, 0x83, 0x3f, 0x42, 0x1a, 0x0a, 0xe2, 0xe7,
	0x4f, 0x1e, 0xda, 0x01, 0x79, 0x92, 0x9e, 0x57, 0xab, 0xa3, 0x9d, 0x4f, 0xc3, 0x1f, 0xe9, 0x7a,
	0xb5, 0x7c, 0xb7, 0xbe, 0xbe, 0x4e, 0xfd, 0x0f, 0xfe, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x22, 0x9e, 0xf8, 0xf0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReportsClient is the client API for Reports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportsClient interface {
	Create(ctx context.Context, in *Report, opts ...grpc.CallOption) (*ReportAck, error)
}

type reportsClient struct {
	cc *grpc.ClientConn
}

func NewReportsClient(cc *grpc.ClientConn) ReportsClient {
	return &reportsClient{cc}
}

func (c *reportsClient) Create(ctx context.Context, in *Report, opts ...grpc.CallOption) (*ReportAck, error) {
	out := new(ReportAck)
	err := c.cc.Invoke(ctx, "/monny.monitor.Reports/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportsServer is the server API for Reports service.
type ReportsServer interface {
	Create(context.Context, *Report) (*ReportAck, error)
}

// UnimplementedReportsServer can be embedded to have forward compatible implementations.
type UnimplementedReportsServer struct {
}

func (*UnimplementedReportsServer) Create(ctx context.Context, req *Report) (*ReportAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterReportsServer(s *grpc.Server, srv ReportsServer) {
	s.RegisterService(&_Reports_serviceDesc, srv)
}

func _Reports_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Report)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monny.monitor.Reports/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServer).Create(ctx, req.(*Report))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reports_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monny.monitor.Reports",
	HandlerType: (*ReportsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Reports_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report.proto",
}
