// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/squash/v3alpha/squash.proto

package envoy_config_filter_http_squash_v3alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// [#proto-status: experimental]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *types.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *types.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *types.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *types.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d53a777ce9526d7, []int{0}
}
func (m *Squash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return m.Size()
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *types.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *types.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *types.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *types.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v3alpha.Squash")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/squash/v3alpha/squash.proto", fileDescriptor_6d53a777ce9526d7)
}

var fileDescriptor_6d53a777ce9526d7 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0x02, 0x31,
	0x18, 0x86, 0x53, 0x40, 0x0c, 0x25, 0xd1, 0x78, 0x1a, 0x39, 0x89, 0xb9, 0x10, 0x1d, 0x64, 0x6a,
	0x13, 0xd1, 0x3f, 0x70, 0x71, 0x60, 0xf3, 0x72, 0xb0, 0x93, 0x72, 0x14, 0xee, 0x92, 0x72, 0x3d,
	0x7a, 0x5f, 0x89, 0xfc, 0x35, 0x27, 0x47, 0x47, 0x27, 0x67, 0xc3, 0xe6, 0xbf, 0x30, 0xb4, 0x25,
	0x10, 0x1d, 0x70, 0x6b, 0xf3, 0xf6, 0x79, 0xbe, 0xb7, 0xf9, 0xf0, 0x03, 0xcf, 0x97, 0x72, 0x45,
	0x13, 0x99, 0x4f, 0xb3, 0x19, 0x9d, 0x66, 0x02, 0xb8, 0xa2, 0x29, 0x40, 0x41, 0xcb, 0x85, 0x66,
	0x65, 0x4a, 0x97, 0x3d, 0x26, 0x8a, 0x94, 0xb9, 0x2b, 0x29, 0x94, 0x04, 0xe9, 0xdd, 0x19, 0x8a,
	0x58, 0x8a, 0x58, 0x8a, 0x6c, 0x28, 0xe2, 0x9e, 0x39, 0xaa, 0x1d, 0xcc, 0xa4, 0x9c, 0x09, 0x4e,
	0x0d, 0x36, 0xd6, 0x53, 0x3a, 0xd1, 0x8a, 0x41, 0x26, 0x73, 0x2b, 0x6a, 0x5f, 0xff, 0xce, 0x4b,
	0x50, 0x3a, 0x01, 0x97, 0xb6, 0x96, 0x4c, 0x64, 0x13, 0x06, 0x9c, 0x6e, 0x0f, 0x36, 0xb8, 0xf9,
	0xac, 0xe0, 0xfa, 0xc0, 0x4c, 0xf2, 0x6e, 0xf1, 0x71, 0x22, 0x74, 0x09, 0x5c, 0xf9, 0xa8, 0x83,
	0xba, 0x8d, 0xb0, 0xf1, 0xfa, 0xfd, 0x56, 0xad, 0xa9, 0x4a, 0x07, 0xc5, 0xdb, 0xc4, 0xeb, 0xe3,
	0x73, 0x06, 0xc0, 0x92, 0x74, 0xce, 0x73, 0x18, 0x01, 0x9f, 0x17, 0x82, 0x01, 0xf7, 0x2b, 0x1d,
	0xd4, 0x6d, 0xde, 0xb7, 0x88, 0x2d, 0x41, 0xb6, 0x25, 0xc8, 0xc0, 0x94, 0x88, 0xbd, 0x1d, 0x33,
	0x74, 0x88, 0x17, 0xe2, 0x53, 0xc5, 0x17, 0x9a, 0x97, 0x30, 0x82, 0x6c, 0xce, 0xa5, 0x06, 0xbf,
	0x6a, 0x2c, 0x57, 0x7f, 0x2c, 0x4f, 0xee, 0xab, 0xf1, 0x89, 0x23, 0x86, 0x16, 0xf0, 0xfa, 0xd8,
	0xdb, 0x6f, 0xe3, 0x34, 0xb5, 0x43, 0x9a, 0xb3, 0xbd, 0x3a, 0xce, 0xf4, 0x8c, 0x2f, 0xf7, 0x4c,
	0x85, 0x14, 0x62, 0x54, 0x70, 0x95, 0xc9, 0x89, 0x7f, 0x74, 0xc8, 0x76, 0xb1, 0x03, 0x23, 0x29,
	0x44, 0x64, 0xb0, 0x30, 0x7e, 0x5f, 0x07, 0xe8, 0x63, 0x1d, 0xa0, 0xaf, 0x75, 0x80, 0xf0, 0x63,
	0x26, 0x89, 0xd9, 0x74, 0xa1, 0xe4, 0xcb, 0x8a, 0xfc, 0x73, 0xe9, 0x61, 0xd3, 0xae, 0x26, 0xda,
	0xcc, 0x8c, 0xd0, 0xb8, 0x6e, 0x86, 0xf7, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0xcd, 0x45,
	0x37, 0x6b, 0x02, 0x00, 0x00,
}

func (m *Squash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Squash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Squash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AttachmentPollPeriod != nil {
		{
			size, err := m.AttachmentPollPeriod.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.AttachmentTimeout != nil {
		{
			size, err := m.AttachmentTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.RequestTimeout != nil {
		{
			size, err := m.RequestTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AttachmentTemplate != nil {
		{
			size, err := m.AttachmentTemplate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = encodeVarintSquash(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSquash(dAtA []byte, offset int, v uint64) int {
	offset -= sovSquash(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Squash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTemplate != nil {
		l = m.AttachmentTemplate.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.RequestTimeout != nil {
		l = m.RequestTimeout.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTimeout != nil {
		l = m.AttachmentTimeout.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentPollPeriod != nil {
		l = m.AttachmentPollPeriod.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSquash(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSquash(x uint64) (n int) {
	return sovSquash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Squash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSquash
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
			return fmt.Errorf("proto: Squash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Squash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTemplate == nil {
				m.AttachmentTemplate = &types.Struct{}
			}
			if err := m.AttachmentTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestTimeout == nil {
				m.RequestTimeout = &types.Duration{}
			}
			if err := m.RequestTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTimeout == nil {
				m.AttachmentTimeout = &types.Duration{}
			}
			if err := m.AttachmentTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentPollPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentPollPeriod == nil {
				m.AttachmentPollPeriod = &types.Duration{}
			}
			if err := m.AttachmentPollPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSquash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSquash
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSquash
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSquash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSquash
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
					return 0, ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSquash
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
				return 0, ErrInvalidLengthSquash
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSquash
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSquash
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSquash(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSquash
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSquash = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSquash   = fmt.Errorf("proto: integer overflow")
)
