// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: executor.proto

package pb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CancelBatchTasksRequest struct {
	TaskIdList []int32 `protobuf:"varint,1,rep,packed,name=task_id_list,json=taskIdList,proto3" json:"task_id_list,omitempty"`
}

func (m *CancelBatchTasksRequest) Reset()         { *m = CancelBatchTasksRequest{} }
func (m *CancelBatchTasksRequest) String() string { return proto.CompactTextString(m) }
func (*CancelBatchTasksRequest) ProtoMessage()    {}
func (*CancelBatchTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}
func (m *CancelBatchTasksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelBatchTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelBatchTasksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelBatchTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelBatchTasksRequest.Merge(m, src)
}
func (m *CancelBatchTasksRequest) XXX_Size() int {
	return m.Size()
}
func (m *CancelBatchTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelBatchTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelBatchTasksRequest proto.InternalMessageInfo

func (m *CancelBatchTasksRequest) GetTaskIdList() []int32 {
	if m != nil {
		return m.TaskIdList
	}
	return nil
}

type SubmitBatchTasksRequest struct {
	Tasks []*TaskRequest `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (m *SubmitBatchTasksRequest) Reset()         { *m = SubmitBatchTasksRequest{} }
func (m *SubmitBatchTasksRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitBatchTasksRequest) ProtoMessage()    {}
func (*SubmitBatchTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1}
}
func (m *SubmitBatchTasksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitBatchTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitBatchTasksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitBatchTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitBatchTasksRequest.Merge(m, src)
}
func (m *SubmitBatchTasksRequest) XXX_Size() int {
	return m.Size()
}
func (m *SubmitBatchTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitBatchTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitBatchTasksRequest proto.InternalMessageInfo

func (m *SubmitBatchTasksRequest) GetTasks() []*TaskRequest {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskRequest struct {
	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Inputs  []int32 `protobuf:"varint,2,rep,packed,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []int32 `protobuf:"varint,3,rep,packed,name=outputs,proto3" json:"outputs,omitempty"`
	Op      []byte  `protobuf:"bytes,4,opt,name=op,proto3" json:"op,omitempty"`
	OpTp    int32   `protobuf:"varint,5,opt,name=op_tp,json=opTp,proto3" json:"op_tp,omitempty"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{2}
}
func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return m.Size()
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskRequest) GetInputs() []int32 {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TaskRequest) GetOutputs() []int32 {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TaskRequest) GetOp() []byte {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *TaskRequest) GetOpTp() int32 {
	if m != nil {
		return m.OpTp
	}
	return 0
}

type SubmitBatchTasksResponse struct {
	Err *Error `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *SubmitBatchTasksResponse) Reset()         { *m = SubmitBatchTasksResponse{} }
func (m *SubmitBatchTasksResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitBatchTasksResponse) ProtoMessage()    {}
func (*SubmitBatchTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{3}
}
func (m *SubmitBatchTasksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmitBatchTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmitBatchTasksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmitBatchTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitBatchTasksResponse.Merge(m, src)
}
func (m *SubmitBatchTasksResponse) XXX_Size() int {
	return m.Size()
}
func (m *SubmitBatchTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitBatchTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitBatchTasksResponse proto.InternalMessageInfo

func (m *SubmitBatchTasksResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type CancelBatchTasksResponse struct {
	Err *Error `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *CancelBatchTasksResponse) Reset()         { *m = CancelBatchTasksResponse{} }
func (m *CancelBatchTasksResponse) String() string { return proto.CompactTextString(m) }
func (*CancelBatchTasksResponse) ProtoMessage()    {}
func (*CancelBatchTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{4}
}
func (m *CancelBatchTasksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelBatchTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelBatchTasksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelBatchTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelBatchTasksResponse.Merge(m, src)
}
func (m *CancelBatchTasksResponse) XXX_Size() int {
	return m.Size()
}
func (m *CancelBatchTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelBatchTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelBatchTasksResponse proto.InternalMessageInfo

func (m *CancelBatchTasksResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*CancelBatchTasksRequest)(nil), "pb.CancelBatchTasksRequest")
	proto.RegisterType((*SubmitBatchTasksRequest)(nil), "pb.SubmitBatchTasksRequest")
	proto.RegisterType((*TaskRequest)(nil), "pb.TaskRequest")
	proto.RegisterType((*SubmitBatchTasksResponse)(nil), "pb.SubmitBatchTasksResponse")
	proto.RegisterType((*CancelBatchTasksResponse)(nil), "pb.CancelBatchTasksResponse")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xcd, 0x26, 0x4d, 0xbf, 0xcf, 0x49, 0xa9, 0xb2, 0x82, 0x5d, 0x5a, 0x09, 0x21, 0x20, 0xe4,
	0xd4, 0x43, 0x3d, 0x78, 0xf0, 0x22, 0x95, 0x1e, 0x04, 0x41, 0x58, 0x7b, 0x2f, 0x49, 0xb3, 0xe0,
	0xd2, 0xda, 0x5d, 0xb3, 0x13, 0xf0, 0x67, 0xf8, 0x43, 0xfc, 0x21, 0x1e, 0x7b, 0xf4, 0x28, 0xed,
	0x1f, 0x91, 0x4d, 0x5a, 0x90, 0xd6, 0x82, 0xc7, 0x79, 0x6f, 0xde, 0x9b, 0xb7, 0x33, 0x0b, 0x6d,
	0xf1, 0x2a, 0xa6, 0x25, 0xaa, 0xa2, 0xaf, 0x0b, 0x85, 0x8a, 0xba, 0x3a, 0xeb, 0x06, 0xa2, 0x28,
	0xb6, 0x40, 0x7c, 0x0d, 0x9d, 0xdb, 0x74, 0x31, 0x15, 0xf3, 0x61, 0x8a, 0xd3, 0xa7, 0x71, 0x6a,
	0x66, 0x86, 0x8b, 0x97, 0x52, 0x18, 0xa4, 0x11, 0xb4, 0x30, 0x35, 0xb3, 0x89, 0xcc, 0x27, 0x73,
	0x69, 0x90, 0x91, 0xc8, 0x4b, 0x7c, 0x0e, 0x16, 0xbb, 0xcb, 0xef, 0xa5, 0xc1, 0xf8, 0x06, 0x3a,
	0x8f, 0x65, 0xf6, 0x2c, 0x71, 0x5f, 0x7c, 0x01, 0xbe, 0x6d, 0x34, 0xcc, 0x8d, 0xbc, 0x24, 0x18,
	0x1c, 0xf7, 0x75, 0xd6, 0xb7, 0x0d, 0x1b, 0x9e, 0xd7, 0x6c, 0x8c, 0x10, 0xfc, 0x40, 0x69, 0x1b,
	0x5c, 0x99, 0x33, 0x12, 0x91, 0xc4, 0xe7, 0xae, 0xcc, 0xe9, 0x19, 0x34, 0xe5, 0x42, 0x97, 0x58,
	0xdb, 0xf8, 0x7c, 0x53, 0x51, 0x06, 0xff, 0x54, 0x89, 0x15, 0xe1, 0x55, 0xc4, 0xb6, 0xb4, 0x0e,
	0x4a, 0xb3, 0x46, 0x44, 0x92, 0x16, 0x77, 0x95, 0xa6, 0xa7, 0xe0, 0x2b, 0x3d, 0x41, 0xcd, 0xfc,
	0xca, 0xb4, 0xa1, 0xf4, 0x58, 0xc7, 0x57, 0xc0, 0xf6, 0x73, 0x1b, 0xad, 0x16, 0x46, 0xd0, 0x1e,
	0x78, 0xa2, 0x28, 0xaa, 0x0c, 0xc1, 0xe0, 0xc8, 0xc6, 0x1e, 0xd9, 0x75, 0x71, 0x8b, 0x5a, 0xe1,
	0xfe, 0xb6, 0xfe, 0x20, 0x1c, 0xbc, 0x13, 0xf8, 0x3f, 0xda, 0x9c, 0x82, 0x3e, 0xc0, 0xc9, 0xee,
	0x78, 0xda, 0xb3, 0x82, 0x03, 0xcb, 0xec, 0x9e, 0xff, 0x4e, 0xd6, 0x83, 0x63, 0xc7, 0x1a, 0xee,
	0xc6, 0xaa, 0x0d, 0x0f, 0x9c, 0xb6, 0x36, 0x3c, 0xf4, 0x92, 0xd8, 0x19, 0xb2, 0x8f, 0x55, 0x48,
	0x96, 0xab, 0x90, 0x7c, 0xad, 0x42, 0xf2, 0xb6, 0x0e, 0x9d, 0xe5, 0x3a, 0x74, 0x3e, 0xd7, 0xa1,
	0x93, 0x35, 0xab, 0x6f, 0x73, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xce, 0x9c, 0x2f, 0x59,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	SubmitBatchTasks(ctx context.Context, in *SubmitBatchTasksRequest, opts ...grpc.CallOption) (*SubmitBatchTasksResponse, error)
	CancelBatchTasks(ctx context.Context, in *CancelBatchTasksRequest, opts ...grpc.CallOption) (*CancelBatchTasksResponse, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) SubmitBatchTasks(ctx context.Context, in *SubmitBatchTasksRequest, opts ...grpc.CallOption) (*SubmitBatchTasksResponse, error) {
	out := new(SubmitBatchTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.Executor/SubmitBatchTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) CancelBatchTasks(ctx context.Context, in *CancelBatchTasksRequest, opts ...grpc.CallOption) (*CancelBatchTasksResponse, error) {
	out := new(CancelBatchTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.Executor/CancelBatchTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	SubmitBatchTasks(context.Context, *SubmitBatchTasksRequest) (*SubmitBatchTasksResponse, error)
	CancelBatchTasks(context.Context, *CancelBatchTasksRequest) (*CancelBatchTasksResponse, error)
}

// UnimplementedExecutorServer can be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (*UnimplementedExecutorServer) SubmitBatchTasks(ctx context.Context, req *SubmitBatchTasksRequest) (*SubmitBatchTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBatchTasks not implemented")
}
func (*UnimplementedExecutorServer) CancelBatchTasks(ctx context.Context, req *CancelBatchTasksRequest) (*CancelBatchTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBatchTasks not implemented")
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_SubmitBatchTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitBatchTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).SubmitBatchTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/SubmitBatchTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).SubmitBatchTasks(ctx, req.(*SubmitBatchTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_CancelBatchTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBatchTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).CancelBatchTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/CancelBatchTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).CancelBatchTasks(ctx, req.(*CancelBatchTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitBatchTasks",
			Handler:    _Executor_SubmitBatchTasks_Handler,
		},
		{
			MethodName: "CancelBatchTasks",
			Handler:    _Executor_CancelBatchTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executor.proto",
}

func (m *CancelBatchTasksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelBatchTasksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelBatchTasksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaskIdList) > 0 {
		dAtA2 := make([]byte, len(m.TaskIdList)*10)
		var j1 int
		for _, num1 := range m.TaskIdList {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintExecutor(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubmitBatchTasksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitBatchTasksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitBatchTasksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tasks) > 0 {
		for iNdEx := len(m.Tasks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tasks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecutor(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *TaskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OpTp != 0 {
		i = encodeVarintExecutor(dAtA, i, uint64(m.OpTp))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Op) > 0 {
		i -= len(m.Op)
		copy(dAtA[i:], m.Op)
		i = encodeVarintExecutor(dAtA, i, uint64(len(m.Op)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Outputs) > 0 {
		dAtA4 := make([]byte, len(m.Outputs)*10)
		var j3 int
		for _, num1 := range m.Outputs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintExecutor(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Inputs) > 0 {
		dAtA6 := make([]byte, len(m.Inputs)*10)
		var j5 int
		for _, num1 := range m.Inputs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintExecutor(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintExecutor(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubmitBatchTasksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmitBatchTasksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmitBatchTasksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Err != nil {
		{
			size, err := m.Err.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExecutor(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CancelBatchTasksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelBatchTasksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelBatchTasksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Err != nil {
		{
			size, err := m.Err.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExecutor(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExecutor(dAtA []byte, offset int, v uint64) int {
	offset -= sovExecutor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CancelBatchTasksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TaskIdList) > 0 {
		l = 0
		for _, e := range m.TaskIdList {
			l += sovExecutor(uint64(e))
		}
		n += 1 + sovExecutor(uint64(l)) + l
	}
	return n
}

func (m *SubmitBatchTasksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovExecutor(uint64(l))
		}
	}
	return n
}

func (m *TaskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExecutor(uint64(m.Id))
	}
	if len(m.Inputs) > 0 {
		l = 0
		for _, e := range m.Inputs {
			l += sovExecutor(uint64(e))
		}
		n += 1 + sovExecutor(uint64(l)) + l
	}
	if len(m.Outputs) > 0 {
		l = 0
		for _, e := range m.Outputs {
			l += sovExecutor(uint64(e))
		}
		n += 1 + sovExecutor(uint64(l)) + l
	}
	l = len(m.Op)
	if l > 0 {
		n += 1 + l + sovExecutor(uint64(l))
	}
	if m.OpTp != 0 {
		n += 1 + sovExecutor(uint64(m.OpTp))
	}
	return n
}

func (m *SubmitBatchTasksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Err != nil {
		l = m.Err.Size()
		n += 1 + l + sovExecutor(uint64(l))
	}
	return n
}

func (m *CancelBatchTasksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Err != nil {
		l = m.Err.Size()
		n += 1 + l + sovExecutor(uint64(l))
	}
	return n
}

func sovExecutor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExecutor(x uint64) (n int) {
	return sovExecutor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CancelBatchTasksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CancelBatchTasksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelBatchTasksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExecutor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TaskIdList = append(m.TaskIdList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExecutor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthExecutor
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthExecutor
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.TaskIdList) == 0 {
					m.TaskIdList = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExecutor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TaskIdList = append(m.TaskIdList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskIdList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SubmitBatchTasksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmitBatchTasksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitBatchTasksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecutor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, &TaskRequest{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExecutor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Inputs = append(m.Inputs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExecutor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthExecutor
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthExecutor
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Inputs) == 0 {
					m.Inputs = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExecutor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Inputs = append(m.Inputs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
		case 3:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExecutor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Outputs = append(m.Outputs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExecutor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthExecutor
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthExecutor
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Outputs) == 0 {
					m.Outputs = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExecutor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Outputs = append(m.Outputs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthExecutor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Op = append(m.Op[:0], dAtA[iNdEx:postIndex]...)
			if m.Op == nil {
				m.Op = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpTp", wireType)
			}
			m.OpTp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpTp |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SubmitBatchTasksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmitBatchTasksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmitBatchTasksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecutor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Err == nil {
				m.Err = &Error{}
			}
			if err := m.Err.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CancelBatchTasksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CancelBatchTasksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelBatchTasksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecutor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecutor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Err == nil {
				m.Err = &Error{}
			}
			if err := m.Err.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExecutor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExecutor
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExecutor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthExecutor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExecutor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExecutor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExecutor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExecutor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExecutor = fmt.Errorf("proto: unexpected end of group")
)
