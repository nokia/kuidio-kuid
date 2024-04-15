/*
Copyright 2024 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/ // Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kuidio/kuid/apis/common/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"

	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (m *ClaimLabels) Reset()      { *m = ClaimLabels{} }
func (*ClaimLabels) ProtoMessage() {}
func (*ClaimLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff64f0e0dd0fa4cf, []int{0}
}
func (m *ClaimLabels) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClaimLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimLabels.Merge(m, src)
}
func (m *ClaimLabels) XXX_Size() int {
	return m.Size()
}
func (m *ClaimLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimLabels.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimLabels proto.InternalMessageInfo

func (m *OwnerReference) Reset()      { *m = OwnerReference{} }
func (*OwnerReference) ProtoMessage() {}
func (*OwnerReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff64f0e0dd0fa4cf, []int{1}
}
func (m *OwnerReference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnerReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *OwnerReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerReference.Merge(m, src)
}
func (m *OwnerReference) XXX_Size() int {
	return m.Size()
}
func (m *OwnerReference) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerReference.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerReference proto.InternalMessageInfo

func (m *UserDefinedLabels) Reset()      { *m = UserDefinedLabels{} }
func (*UserDefinedLabels) ProtoMessage() {}
func (*UserDefinedLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff64f0e0dd0fa4cf, []int{2}
}
func (m *UserDefinedLabels) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserDefinedLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *UserDefinedLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDefinedLabels.Merge(m, src)
}
func (m *UserDefinedLabels) XXX_Size() int {
	return m.Size()
}
func (m *UserDefinedLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDefinedLabels.DiscardUnknown(m)
}

var xxx_messageInfo_UserDefinedLabels proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClaimLabels)(nil), "github.com.kuidio.kuid.apis.common.v1alpha1.ClaimLabels")
	proto.RegisterType((*OwnerReference)(nil), "github.com.kuidio.kuid.apis.common.v1alpha1.OwnerReference")
	proto.RegisterType((*UserDefinedLabels)(nil), "github.com.kuidio.kuid.apis.common.v1alpha1.UserDefinedLabels")
	proto.RegisterMapType((map[string]string)(nil), "github.com.kuidio.kuid.apis.common.v1alpha1.UserDefinedLabels.LabelsEntry")
}

func init() {
	proto.RegisterFile("github.com/kuidio/kuid/apis/common/v1alpha1/generated.proto", fileDescriptor_ff64f0e0dd0fa4cf)
}

var fileDescriptor_ff64f0e0dd0fa4cf = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xbf, 0x6f, 0xd3, 0x4e,
	0x14, 0xcf, 0x35, 0x49, 0xdb, 0x5c, 0xfa, 0xed, 0x97, 0x9c, 0x18, 0x42, 0x06, 0xa7, 0x0a, 0x4b,
	0x51, 0xc5, 0x59, 0x29, 0x0c, 0x05, 0x24, 0x06, 0x03, 0x42, 0x02, 0x04, 0xe2, 0x10, 0x0c, 0x48,
	0x0c, 0x17, 0xe7, 0xc5, 0x39, 0xd9, 0xbe, 0xb3, 0xce, 0x76, 0x50, 0x36, 0x36, 0x56, 0xfe, 0x26,
	0xa6, 0x8c, 0x5d, 0x90, 0x3a, 0x45, 0xc4, 0xfc, 0x0f, 0xcc, 0xc8, 0x77, 0x0e, 0x0d, 0x64, 0xaa,
	0x98, 0x7c, 0xef, 0x3e, 0x3f, 0xde, 0xe7, 0xbd, 0x93, 0xf1, 0x83, 0x40, 0x64, 0xd3, 0x7c, 0x44,
	0x7d, 0x15, 0xbb, 0x61, 0x2e, 0xc6, 0x42, 0x99, 0x8f, 0xcb, 0x13, 0x91, 0xba, 0xbe, 0x8a, 0x63,
	0x25, 0xdd, 0xd9, 0x90, 0x47, 0xc9, 0x94, 0x0f, 0xdd, 0x00, 0x24, 0x68, 0x9e, 0xc1, 0x98, 0x26,
	0x5a, 0x65, 0x8a, 0x9c, 0x5c, 0x8a, 0xa9, 0x15, 0x9b, 0x0f, 0x2d, 0xc5, 0xd4, 0x8a, 0xe9, 0x5a,
	0xdc, 0xbb, 0xbd, 0xd1, 0x29, 0x50, 0x81, 0x72, 0x8d, 0xc7, 0x28, 0x9f, 0x98, 0xca, 0x14, 0xe6,
	0x64, 0xbd, 0x7b, 0x77, 0xc3, 0xb3, 0x94, 0x0a, 0x55, 0x06, 0x89, 0xb9, 0x3f, 0x15, 0x12, 0xf4,
	0xdc, 0x4d, 0xc2, 0xc0, 0x26, 0x8b, 0x21, 0xe3, 0xee, 0x6c, 0x2b, 0xd1, 0xe0, 0x27, 0xc2, 0xed,
	0x47, 0x11, 0x17, 0xf1, 0x0b, 0x3e, 0x82, 0x28, 0x25, 0x9f, 0x11, 0xee, 0xe4, 0x29, 0xe8, 0xc7,
	0x30, 0x11, 0x12, 0xc6, 0xf6, 0xb6, 0x8b, 0x8e, 0xd0, 0x71, 0xfb, 0xf4, 0x21, 0xbd, 0x42, 0x7c,
	0xfa, 0xf6, 0x6f, 0x17, 0xef, 0xc6, 0x62, 0xd9, 0xaf, 0x15, 0xcb, 0x7e, 0x67, 0x0b, 0x62, 0xdb,
	0x3d, 0xc9, 0x07, 0xbc, 0x9f, 0x42, 0x04, 0x7e, 0xa6, 0x74, 0x77, 0xc7, 0xf4, 0xbf, 0x43, 0xed,
	0x88, 0x74, 0x73, 0x44, 0x9a, 0x84, 0x81, 0x0d, 0x50, 0x8e, 0x48, 0x67, 0x43, 0x6a, 0xf4, 0x6f,
	0x2a, 0xa9, 0x77, 0x50, 0x2c, 0xfb, 0xfb, 0xeb, 0x8a, 0xfd, 0xb6, 0x1c, 0x7c, 0x43, 0xf8, 0xf0,
	0xd5, 0x47, 0x09, 0x9a, 0xc1, 0x04, 0x34, 0x48, 0x1f, 0xc8, 0x4d, 0xdc, 0x0c, 0xb4, 0xca, 0x13,
	0x33, 0x6e, 0xcb, 0xfb, 0xaf, 0x8a, 0xdb, 0x7c, 0x5a, 0x5e, 0x32, 0x8b, 0x91, 0x5b, 0x78, 0x6f,
	0x06, 0x3a, 0x15, 0x4a, 0x9a, 0x54, 0x2d, 0xef, 0xff, 0x8a, 0xb6, 0xf7, 0xce, 0x5e, 0xb3, 0x35,
	0x4e, 0x8e, 0x70, 0x23, 0x14, 0x72, 0xdc, 0xad, 0x1b, 0xde, 0x41, 0xc5, 0x6b, 0x3c, 0x17, 0x72,
	0xcc, 0x0c, 0x42, 0x5c, 0xdc, 0x92, 0x3c, 0x86, 0x34, 0xe1, 0x3e, 0x74, 0x1b, 0x86, 0xd6, 0xa9,
	0x68, 0xad, 0x97, 0x6b, 0x80, 0x5d, 0x72, 0x4a, 0xcb, 0xb2, 0xe8, 0x36, 0xff, 0xb4, 0x2c, 0xb9,
	0xcc, 0x20, 0x83, 0xaf, 0x08, 0x6f, 0xef, 0x97, 0x68, 0xbc, 0x1b, 0xad, 0x9f, 0xb2, 0x7e, 0xdc,
	0x3e, 0x7d, 0xf6, 0x6f, 0x4f, 0x69, 0x77, 0x9c, 0x3e, 0x91, 0x99, 0x9e, 0x7b, 0x87, 0x55, 0x8a,
	0xdd, 0xea, 0x2d, 0xab, 0x4e, 0xbd, 0x7b, 0xb8, 0xbd, 0x41, 0x23, 0xd7, 0x70, 0x3d, 0x84, 0xb9,
	0xdd, 0x2d, 0x2b, 0x8f, 0xe4, 0x3a, 0x6e, 0xce, 0x78, 0x94, 0x83, 0x5d, 0x24, 0xb3, 0xc5, 0xfd,
	0x9d, 0x33, 0xe4, 0xbd, 0x5e, 0xac, 0x9c, 0xda, 0xf9, 0xca, 0xa9, 0x5d, 0xac, 0x9c, 0xda, 0xa7,
	0xc2, 0x41, 0x8b, 0xc2, 0x41, 0xe7, 0x85, 0x83, 0x2e, 0x0a, 0x07, 0x7d, 0x2f, 0x1c, 0xf4, 0xe5,
	0x87, 0x53, 0x7b, 0x7f, 0x72, 0x85, 0x5f, 0xf1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x02,
	0xae, 0x5d, 0xb8, 0x03, 0x00, 0x00,
}

func (m *ClaimLabels) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimLabels) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimLabels) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Selector != nil {
		{
			size, err := m.Selector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.UserDefinedLabels.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OwnerReference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnerReference) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnerReference) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0x2a
	i -= len(m.Namespace)
	copy(dAtA[i:], m.Namespace)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Namespace)))
	i--
	dAtA[i] = 0x22
	i -= len(m.Kind)
	copy(dAtA[i:], m.Kind)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Kind)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.Version)
	copy(dAtA[i:], m.Version)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Version)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Group)
	copy(dAtA[i:], m.Group)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Group)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *UserDefinedLabels) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserDefinedLabels) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserDefinedLabels) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		keysForLabels := make([]string, 0, len(m.Labels))
		for k := range m.Labels {
			keysForLabels = append(keysForLabels, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
		for iNdEx := len(keysForLabels) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Labels[string(keysForLabels[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(keysForLabels[iNdEx])
			copy(dAtA[i:], keysForLabels[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForLabels[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimLabels) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.UserDefinedLabels.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *OwnerReference) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Version)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *UserDefinedLabels) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClaimLabels) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClaimLabels{`,
		`UserDefinedLabels:` + strings.Replace(strings.Replace(this.UserDefinedLabels.String(), "UserDefinedLabels", "UserDefinedLabels", 1), `&`, ``, 1) + `,`,
		`Selector:` + strings.Replace(fmt.Sprintf("%v", this.Selector), "LabelSelector", "v1.LabelSelector", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *OwnerReference) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OwnerReference{`,
		`Group:` + fmt.Sprintf("%v", this.Group) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UserDefinedLabels) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&UserDefinedLabels{`,
		`Labels:` + mapStringForLabels + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClaimLabels) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ClaimLabels: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimLabels: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserDefinedLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserDefinedLabels.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *OwnerReference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: OwnerReference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerReference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *UserDefinedLabels) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: UserDefinedLabels: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserDefinedLabels: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
							return ErrIntOverflowGenerated
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
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
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
							return ErrIntOverflowGenerated
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
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
