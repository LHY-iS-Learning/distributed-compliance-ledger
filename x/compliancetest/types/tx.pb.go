// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliancetest/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgAddTestingResult struct {
	Signer                string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty" validate:"required"`
	Vid                   int32  `protobuf:"varint,2,opt,name=vid,proto3" json:"vid,omitempty" validate:"required,gte=0,lte=65535"`
	Pid                   int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty" validate:"required,gte=0,lte=65535"`
	SoftwareVersion       uint32 `protobuf:"varint,4,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty" validate:"required"`
	SoftwareVersionString string `protobuf:"bytes,5,opt,name=software_version_string,json=softwareVersionString,proto3" json:"software_version_string,omitempty" validate:"required"`
	TestResult            string `protobuf:"bytes,6,opt,name=test_result,json=testResult,proto3" json:"test_result,omitempty" validate:"required"`
	TestDate              string `protobuf:"bytes,7,opt,name=test_date,json=testDate,proto3" json:"test_date,omitempty" validate:"required"`
}

func (m *MsgAddTestingResult) Reset()         { *m = MsgAddTestingResult{} }
func (m *MsgAddTestingResult) String() string { return proto.CompactTextString(m) }
func (*MsgAddTestingResult) ProtoMessage()    {}
func (*MsgAddTestingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92cb3976812d254, []int{0}
}
func (m *MsgAddTestingResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddTestingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddTestingResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddTestingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddTestingResult.Merge(m, src)
}
func (m *MsgAddTestingResult) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddTestingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddTestingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddTestingResult proto.InternalMessageInfo

func (m *MsgAddTestingResult) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgAddTestingResult) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *MsgAddTestingResult) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *MsgAddTestingResult) GetSoftwareVersion() uint32 {
	if m != nil {
		return m.SoftwareVersion
	}
	return 0
}

func (m *MsgAddTestingResult) GetSoftwareVersionString() string {
	if m != nil {
		return m.SoftwareVersionString
	}
	return ""
}

func (m *MsgAddTestingResult) GetTestResult() string {
	if m != nil {
		return m.TestResult
	}
	return ""
}

func (m *MsgAddTestingResult) GetTestDate() string {
	if m != nil {
		return m.TestDate
	}
	return ""
}

type MsgAddTestingResultResponse struct {
}

func (m *MsgAddTestingResultResponse) Reset()         { *m = MsgAddTestingResultResponse{} }
func (m *MsgAddTestingResultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddTestingResultResponse) ProtoMessage()    {}
func (*MsgAddTestingResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d92cb3976812d254, []int{1}
}
func (m *MsgAddTestingResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddTestingResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddTestingResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddTestingResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddTestingResultResponse.Merge(m, src)
}
func (m *MsgAddTestingResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddTestingResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddTestingResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddTestingResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddTestingResult)(nil), "zigbeealliance.distributedcomplianceledger.compliancetest.MsgAddTestingResult")
	proto.RegisterType((*MsgAddTestingResultResponse)(nil), "zigbeealliance.distributedcomplianceledger.compliancetest.MsgAddTestingResultResponse")
}

func init() { proto.RegisterFile("compliancetest/tx.proto", fileDescriptor_d92cb3976812d254) }

var fileDescriptor_d92cb3976812d254 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x84, 0x06, 0x7a, 0x08, 0x51, 0xb9, 0x45, 0x31, 0x41, 0xb8, 0x91, 0x19, 0xc8,
	0xd0, 0xd8, 0x88, 0x12, 0x44, 0x91, 0x3a, 0xd4, 0x42, 0x62, 0x2a, 0x95, 0x0c, 0xea, 0xc0, 0x62,
	0xd9, 0xb9, 0xc7, 0x71, 0x92, 0xeb, 0x33, 0xf7, 0xce, 0xa1, 0xf0, 0x29, 0xf8, 0x30, 0x4c, 0x8c,
	0x4c, 0x8c, 0x15, 0x13, 0x2c, 0x15, 0x4a, 0xbe, 0x41, 0x3f, 0x01, 0x3a, 0x5f, 0xa2, 0x42, 0xda,
	0x46, 0x02, 0xb1, 0xd9, 0xe7, 0xf7, 0xfb, 0xfd, 0xf5, 0xde, 0xf3, 0xd1, 0xf6, 0x50, 0x1e, 0x94,
	0xb9, 0x48, 0x8b, 0x21, 0x68, 0x40, 0x1d, 0xea, 0xc3, 0xa0, 0x54, 0x52, 0x4b, 0x67, 0xeb, 0x83,
	0xe0, 0x19, 0x40, 0x9a, 0xdb, 0x8f, 0x01, 0x13, 0xa8, 0x95, 0xc8, 0x2a, 0x0d, 0xec, 0x14, 0xc9,
	0x81, 0x71, 0x50, 0xc1, 0x9f, 0x8e, 0xce, 0xad, 0xa1, 0xc4, 0x03, 0x89, 0x49, 0x2d, 0x0a, 0xed,
	0x8b, 0xb5, 0x76, 0xd6, 0xb8, 0xe4, 0xd2, 0x9e, 0x9b, 0x27, 0x7b, 0xea, 0xff, 0x68, 0xd2, 0xd5,
	0x5d, 0xe4, 0x3b, 0x8c, 0xbd, 0x04, 0xd4, 0xa2, 0xe0, 0x31, 0x60, 0x95, 0x6b, 0xe7, 0x19, 0x6d,
	0xa1, 0xe0, 0x05, 0x28, 0x97, 0x74, 0x49, 0x6f, 0x39, 0x0a, 0x4f, 0x8e, 0xd7, 0x57, 0x47, 0x69,
	0x2e, 0x58, 0xaa, 0xe1, 0x89, 0xaf, 0xe0, 0x6d, 0x25, 0x14, 0x30, 0xff, 0xdb, 0xa7, 0xfe, 0xda,
	0x34, 0x66, 0x87, 0x31, 0x05, 0x88, 0x2f, 0xb4, 0x32, 0x9a, 0x29, 0xee, 0x6c, 0xd1, 0xe6, 0x48,
	0x30, 0xf7, 0x52, 0x97, 0xf4, 0x96, 0xa2, 0x7b, 0x27, 0xc7, 0xeb, 0x77, 0xcf, 0x5a, 0x36, 0xb8,
	0x86, 0xed, 0xfb, 0x1b, 0xb9, 0x86, 0xed, 0x47, 0x83, 0xc1, 0xe6, 0xc0, 0x8f, 0x0d, 0x63, 0xd0,
	0x52, 0x30, 0xb7, 0xf9, 0x97, 0x68, 0x29, 0x98, 0x13, 0xd1, 0x15, 0x94, 0xaf, 0xf5, 0xbb, 0x54,
	0x41, 0x32, 0x02, 0x85, 0x42, 0x16, 0xee, 0xe5, 0x2e, 0xe9, 0x5d, 0x8f, 0xda, 0x17, 0x34, 0x12,
	0xdf, 0x98, 0x01, 0xfb, 0xb6, 0xde, 0xd9, 0xa3, 0xed, 0x79, 0x47, 0x82, 0x75, 0x73, 0xee, 0x52,
	0x3d, 0x93, 0x0b, 0x55, 0x37, 0xe7, 0x54, 0x76, 0x24, 0xce, 0x63, 0x7a, 0xcd, 0x2c, 0x29, 0x51,
	0xf5, 0x88, 0xdd, 0xd6, 0x62, 0x09, 0x35, 0xb5, 0xd3, 0x6d, 0x3c, 0xa4, 0xcb, 0x35, 0x69, 0x6a,
	0xdc, 0x2b, 0x8b, 0xb9, 0xab, 0xa6, 0xf2, 0x69, 0xaa, 0xc1, 0xbf, 0x43, 0x6f, 0x9f, 0xb3, 0xda,
	0x18, 0xb0, 0x94, 0x05, 0xc2, 0x83, 0x2f, 0x84, 0x36, 0x77, 0x91, 0x3b, 0x9f, 0x09, 0x5d, 0x39,
	0xb3, 0xff, 0xe7, 0xc1, 0x3f, 0xff, 0x84, 0xc1, 0x39, 0xa1, 0x9d, 0xfd, 0xff, 0xeb, 0x9b, 0x35,
	0x11, 0x89, 0xaf, 0x63, 0x8f, 0x1c, 0x8d, 0x3d, 0xf2, 0x73, 0xec, 0x91, 0x8f, 0x13, 0xaf, 0x71,
	0x34, 0xf1, 0x1a, 0xdf, 0x27, 0x5e, 0xe3, 0xd5, 0x1e, 0x17, 0xfa, 0x4d, 0x95, 0x19, 0x59, 0x68,
	0xb3, 0xfb, 0xb3, 0xf0, 0xf0, 0xb7, 0xf0, 0xfe, 0x69, 0x58, 0xdf, 0xc6, 0x87, 0x87, 0xe1, 0xfc,
	0xcd, 0x7c, 0x5f, 0x02, 0x66, 0xad, 0xfa, 0xc6, 0x6c, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x16,
	0x84, 0xb4, 0xb8, 0xb8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddTestingResult(ctx context.Context, in *MsgAddTestingResult, opts ...grpc.CallOption) (*MsgAddTestingResultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddTestingResult(ctx context.Context, in *MsgAddTestingResult, opts ...grpc.CallOption) (*MsgAddTestingResultResponse, error) {
	out := new(MsgAddTestingResultResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.compliancetest.Msg/AddTestingResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddTestingResult(context.Context, *MsgAddTestingResult) (*MsgAddTestingResultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddTestingResult(ctx context.Context, req *MsgAddTestingResult) (*MsgAddTestingResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTestingResult not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddTestingResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddTestingResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddTestingResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.compliancetest.Msg/AddTestingResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddTestingResult(ctx, req.(*MsgAddTestingResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zigbeealliance.distributedcomplianceledger.compliancetest.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTestingResult",
			Handler:    _Msg_AddTestingResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compliancetest/tx.proto",
}

func (m *MsgAddTestingResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddTestingResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddTestingResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TestDate) > 0 {
		i -= len(m.TestDate)
		copy(dAtA[i:], m.TestDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TestDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TestResult) > 0 {
		i -= len(m.TestResult)
		copy(dAtA[i:], m.TestResult)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TestResult)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SoftwareVersionString) > 0 {
		i -= len(m.SoftwareVersionString)
		copy(dAtA[i:], m.SoftwareVersionString)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SoftwareVersionString)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SoftwareVersion != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SoftwareVersion))
		i--
		dAtA[i] = 0x20
	}
	if m.Pid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x18
	}
	if m.Vid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddTestingResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddTestingResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddTestingResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddTestingResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Vid != 0 {
		n += 1 + sovTx(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovTx(uint64(m.Pid))
	}
	if m.SoftwareVersion != 0 {
		n += 1 + sovTx(uint64(m.SoftwareVersion))
	}
	l = len(m.SoftwareVersionString)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TestResult)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TestDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddTestingResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddTestingResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddTestingResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddTestingResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersion", wireType)
			}
			m.SoftwareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SoftwareVersionString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestResult", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestResult = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddTestingResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddTestingResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddTestingResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
