// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/approved_certificates.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type ApprovedCertificates struct {
	Subject      string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId string   `protobuf:"bytes,2,opt,name=subject_key_id,json=subjectKeyId,proto3" json:"subject_key_id,omitempty"`
	Certs        []string `protobuf:"bytes,3,rep,name=certs,proto3" json:"certs,omitempty"`
}

func (m *ApprovedCertificates) Reset()         { *m = ApprovedCertificates{} }
func (m *ApprovedCertificates) String() string { return proto.CompactTextString(m) }
func (*ApprovedCertificates) ProtoMessage()    {}
func (*ApprovedCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce3f4f36edbe99d, []int{0}
}
func (m *ApprovedCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApprovedCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApprovedCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApprovedCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedCertificates.Merge(m, src)
}
func (m *ApprovedCertificates) XXX_Size() int {
	return m.Size()
}
func (m *ApprovedCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedCertificates proto.InternalMessageInfo

func (m *ApprovedCertificates) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ApprovedCertificates) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *ApprovedCertificates) GetCerts() []string {
	if m != nil {
		return m.Certs
	}
	return nil
}

func init() {
	proto.RegisterType((*ApprovedCertificates)(nil), "zigbeealliance.distributedcomplianceledger.pki.ApprovedCertificates")
}

func init() { proto.RegisterFile("pki/approved_certificates.proto", fileDescriptor_3ce3f4f36edbe99d) }

var fileDescriptor_3ce3f4f36edbe99d = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xc8, 0xce, 0xd4,
	0x4f, 0x2c, 0x28, 0x28, 0xca, 0x2f, 0x4b, 0x4d, 0x89, 0x4f, 0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb,
	0x4c, 0x4e, 0x2c, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0xab, 0xca, 0x4c,
	0x4f, 0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4b, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6,
	0xa4, 0xa7, 0x16, 0xe9, 0x15, 0x64, 0x67, 0x2a, 0xe5, 0x70, 0x89, 0x38, 0x42, 0x8d, 0x73, 0x46,
	0x32, 0x4d, 0x48, 0x82, 0x8b, 0xbd, 0xb8, 0x34, 0x29, 0x2b, 0x35, 0xb9, 0x44, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x52, 0xe1, 0xe2, 0x83, 0x32, 0xe3, 0xb3, 0x53, 0x2b, 0xe3,
	0x33, 0x53, 0x24, 0x98, 0xc0, 0x0a, 0x78, 0xa0, 0xa2, 0xde, 0xa9, 0x95, 0x9e, 0x29, 0x42, 0x22,
	0x5c, 0xac, 0x20, 0xd7, 0x15, 0x4b, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x38, 0x4e, 0x71,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xe5, 0x92, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xf1, 0x82, 0x2e, 0xcc, 0x0f, 0xfa, 0x48, 0x7e, 0xd0,
	0x45, 0x78, 0x42, 0x17, 0xe2, 0x0b, 0xfd, 0x0a, 0x7d, 0x50, 0x98, 0x94, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0x03, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x96, 0xc8, 0xf9, 0x27,
	0x01, 0x00, 0x00,
}

func (m *ApprovedCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApprovedCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApprovedCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Certs) > 0 {
		for iNdEx := len(m.Certs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Certs[iNdEx])
			copy(dAtA[i:], m.Certs[iNdEx])
			i = encodeVarintApprovedCertificates(dAtA, i, uint64(len(m.Certs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintApprovedCertificates(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintApprovedCertificates(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApprovedCertificates(dAtA []byte, offset int, v uint64) int {
	offset -= sovApprovedCertificates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ApprovedCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovApprovedCertificates(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovApprovedCertificates(uint64(l))
	}
	if len(m.Certs) > 0 {
		for _, s := range m.Certs {
			l = len(s)
			n += 1 + l + sovApprovedCertificates(uint64(l))
		}
	}
	return n
}

func sovApprovedCertificates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApprovedCertificates(x uint64) (n int) {
	return sovApprovedCertificates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ApprovedCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApprovedCertificates
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
			return fmt.Errorf("proto: ApprovedCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApprovedCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApprovedCertificates
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
				return ErrInvalidLengthApprovedCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApprovedCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApprovedCertificates
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
				return ErrInvalidLengthApprovedCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApprovedCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApprovedCertificates
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
				return ErrInvalidLengthApprovedCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApprovedCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certs = append(m.Certs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApprovedCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApprovedCertificates
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
func skipApprovedCertificates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApprovedCertificates
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
					return 0, ErrIntOverflowApprovedCertificates
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
					return 0, ErrIntOverflowApprovedCertificates
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
				return 0, ErrInvalidLengthApprovedCertificates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApprovedCertificates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApprovedCertificates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApprovedCertificates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApprovedCertificates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApprovedCertificates = fmt.Errorf("proto: unexpected end of group")
)
