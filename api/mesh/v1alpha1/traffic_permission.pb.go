// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/traffic_permission.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TrafficPermission struct {
	Rules                []*TrafficPermission_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TrafficPermission) Reset()         { *m = TrafficPermission{} }
func (m *TrafficPermission) String() string { return proto.CompactTextString(m) }
func (*TrafficPermission) ProtoMessage()    {}
func (*TrafficPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_7871a84a653f4288, []int{0}
}
func (m *TrafficPermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficPermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficPermission.Merge(m, src)
}
func (m *TrafficPermission) XXX_Size() int {
	return m.Size()
}
func (m *TrafficPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficPermission.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficPermission proto.InternalMessageInfo

func (m *TrafficPermission) GetRules() []*TrafficPermission_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type TrafficPermission_Rule struct {
	Sources              []*TrafficPermission_Rule_Selector `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	Destinations         []*TrafficPermission_Rule_Selector `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *TrafficPermission_Rule) Reset()         { *m = TrafficPermission_Rule{} }
func (m *TrafficPermission_Rule) String() string { return proto.CompactTextString(m) }
func (*TrafficPermission_Rule) ProtoMessage()    {}
func (*TrafficPermission_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_7871a84a653f4288, []int{0, 0}
}
func (m *TrafficPermission_Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficPermission_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficPermission_Rule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficPermission_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficPermission_Rule.Merge(m, src)
}
func (m *TrafficPermission_Rule) XXX_Size() int {
	return m.Size()
}
func (m *TrafficPermission_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficPermission_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficPermission_Rule proto.InternalMessageInfo

func (m *TrafficPermission_Rule) GetSources() []*TrafficPermission_Rule_Selector {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *TrafficPermission_Rule) GetDestinations() []*TrafficPermission_Rule_Selector {
	if m != nil {
		return m.Destinations
	}
	return nil
}

type TrafficPermission_Rule_Selector struct {
	Match                map[string]string `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TrafficPermission_Rule_Selector) Reset()         { *m = TrafficPermission_Rule_Selector{} }
func (m *TrafficPermission_Rule_Selector) String() string { return proto.CompactTextString(m) }
func (*TrafficPermission_Rule_Selector) ProtoMessage()    {}
func (*TrafficPermission_Rule_Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_7871a84a653f4288, []int{0, 0, 0}
}
func (m *TrafficPermission_Rule_Selector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficPermission_Rule_Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficPermission_Rule_Selector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrafficPermission_Rule_Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficPermission_Rule_Selector.Merge(m, src)
}
func (m *TrafficPermission_Rule_Selector) XXX_Size() int {
	return m.Size()
}
func (m *TrafficPermission_Rule_Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficPermission_Rule_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficPermission_Rule_Selector proto.InternalMessageInfo

func (m *TrafficPermission_Rule_Selector) GetMatch() map[string]string {
	if m != nil {
		return m.Match
	}
	return nil
}

func init() {
	proto.RegisterType((*TrafficPermission)(nil), "kuma.mesh.v1alpha1.TrafficPermission")
	proto.RegisterType((*TrafficPermission_Rule)(nil), "kuma.mesh.v1alpha1.TrafficPermission.Rule")
	proto.RegisterType((*TrafficPermission_Rule_Selector)(nil), "kuma.mesh.v1alpha1.TrafficPermission.Rule.Selector")
	proto.RegisterMapType((map[string]string)(nil), "kuma.mesh.v1alpha1.TrafficPermission.Rule.Selector.MatchEntry")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/traffic_permission.proto", fileDescriptor_7871a84a653f4288)
}

var fileDescriptor_7871a84a653f4288 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4d, 0x2d, 0xce,
	0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0x4b, 0xcb,
	0x4c, 0x8e, 0x2f, 0x48, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xca, 0x2e, 0xcd, 0x4d, 0xd4, 0x03, 0x29, 0xd6, 0x83, 0x29, 0x56, 0x5a,
	0xcc, 0xcc, 0x25, 0x18, 0x02, 0xd1, 0x10, 0x00, 0x57, 0x2f, 0xe4, 0xc0, 0xc5, 0x5a, 0x54, 0x9a,
	0x93, 0x5a, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa5, 0x87, 0xa9, 0x53, 0x0f, 0x43,
	0x97, 0x5e, 0x50, 0x69, 0x4e, 0x6a, 0x10, 0x44, 0xa3, 0xd4, 0x69, 0x26, 0x2e, 0x16, 0x10, 0x5f,
	0xc8, 0x97, 0x8b, 0xbd, 0x38, 0xbf, 0xb4, 0x28, 0x19, 0x6e, 0x98, 0x31, 0xf1, 0x86, 0xe9, 0x05,
	0xa7, 0xe6, 0xa4, 0x26, 0x97, 0xe4, 0x17, 0x05, 0xc1, 0xcc, 0x10, 0x0a, 0xe7, 0xe2, 0x49, 0x49,
	0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0x60, 0x22, 0xdf, 0x4c, 0x14,
	0x83, 0xa4, 0x66, 0x31, 0x72, 0x71, 0xc0, 0xa4, 0x84, 0x42, 0xb8, 0x58, 0x73, 0x13, 0x4b, 0x92,
	0x33, 0xa0, 0x4e, 0xb6, 0x23, 0xc3, 0x78, 0x3d, 0x5f, 0x90, 0x01, 0xae, 0x79, 0x25, 0x45, 0x95,
	0x41, 0x10, 0xc3, 0xa4, 0x2c, 0xb8, 0xb8, 0x10, 0x82, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e,
	0x69, 0xaa, 0x04, 0x13, 0x58, 0x0c, 0xc2, 0xb1, 0x62, 0xb2, 0x60, 0x74, 0x12, 0x3b, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xa3, 0x38, 0x60, 0x6e, 0x48, 0x62,
	0x03, 0x47, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x82, 0x1b, 0xe4, 0xa4, 0x02, 0x02, 0x00,
	0x00,
}

func (m *TrafficPermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficPermission) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrafficPermission(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TrafficPermission_Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficPermission_Rule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sources) > 0 {
		for _, msg := range m.Sources {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrafficPermission(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Destinations) > 0 {
		for _, msg := range m.Destinations {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTrafficPermission(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TrafficPermission_Rule_Selector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficPermission_Rule_Selector) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Match) > 0 {
		for k, _ := range m.Match {
			dAtA[i] = 0xa
			i++
			v := m.Match[k]
			mapSize := 1 + len(k) + sovTrafficPermission(uint64(len(k))) + 1 + len(v) + sovTrafficPermission(uint64(len(v)))
			i = encodeVarintTrafficPermission(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTrafficPermission(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTrafficPermission(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTrafficPermission(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TrafficPermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovTrafficPermission(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TrafficPermission_Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sources) > 0 {
		for _, e := range m.Sources {
			l = e.Size()
			n += 1 + l + sovTrafficPermission(uint64(l))
		}
	}
	if len(m.Destinations) > 0 {
		for _, e := range m.Destinations {
			l = e.Size()
			n += 1 + l + sovTrafficPermission(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TrafficPermission_Rule_Selector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Match) > 0 {
		for k, v := range m.Match {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTrafficPermission(uint64(len(k))) + 1 + len(v) + sovTrafficPermission(uint64(len(v)))
			n += mapEntrySize + 1 + sovTrafficPermission(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTrafficPermission(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTrafficPermission(x uint64) (n int) {
	return sovTrafficPermission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrafficPermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficPermission
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
			return fmt.Errorf("proto: TrafficPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrafficPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficPermission
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
				return ErrInvalidLengthTrafficPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &TrafficPermission_Rule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficPermission
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
func (m *TrafficPermission_Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficPermission
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
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficPermission
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
				return ErrInvalidLengthTrafficPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sources = append(m.Sources, &TrafficPermission_Rule_Selector{})
			if err := m.Sources[len(m.Sources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficPermission
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
				return ErrInvalidLengthTrafficPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, &TrafficPermission_Rule_Selector{})
			if err := m.Destinations[len(m.Destinations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficPermission
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
func (m *TrafficPermission_Rule_Selector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficPermission
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
			return fmt.Errorf("proto: Selector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Selector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Match", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficPermission
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
				return ErrInvalidLengthTrafficPermission
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Match == nil {
				m.Match = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTrafficPermission
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTrafficPermission
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTrafficPermission
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTrafficPermission
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTrafficPermission
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTrafficPermission
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthTrafficPermission
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTrafficPermission(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTrafficPermission
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Match[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficPermission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficPermission
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficPermission
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
func skipTrafficPermission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrafficPermission
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
					return 0, ErrIntOverflowTrafficPermission
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
					return 0, ErrIntOverflowTrafficPermission
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
				return 0, ErrInvalidLengthTrafficPermission
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTrafficPermission
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrafficPermission
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
				next, err := skipTrafficPermission(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTrafficPermission
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
	ErrInvalidLengthTrafficPermission = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrafficPermission   = fmt.Errorf("proto: integer overflow")
)
