// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plan.proto

package planning

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type EncodedQueryPlan struct {
	TimeRange EncodedQueryTimeRange `protobuf:"bytes,1,opt,name=timeRange,proto3" json:"timeRange"`
	Nodes     []*EncodedNode        `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	RootNode  int64                 `protobuf:"varint,3,opt,name=rootNode,proto3" json:"rootNode,omitempty"`
	// The original PromQL expression for this query.
	// May not accurately represent the query being executed if this query was built from a query plan representing a subexpression of a query.
	OriginalExpression string `protobuf:"bytes,4,opt,name=originalExpression,proto3" json:"originalExpression,omitempty"`
}

func (m *EncodedQueryPlan) Reset()      { *m = EncodedQueryPlan{} }
func (*EncodedQueryPlan) ProtoMessage() {}
func (*EncodedQueryPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{0}
}
func (m *EncodedQueryPlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncodedQueryPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncodedQueryPlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncodedQueryPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedQueryPlan.Merge(m, src)
}
func (m *EncodedQueryPlan) XXX_Size() int {
	return m.Size()
}
func (m *EncodedQueryPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedQueryPlan.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedQueryPlan proto.InternalMessageInfo

func (m *EncodedQueryPlan) GetTimeRange() EncodedQueryTimeRange {
	if m != nil {
		return m.TimeRange
	}
	return EncodedQueryTimeRange{}
}

func (m *EncodedQueryPlan) GetNodes() []*EncodedNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *EncodedQueryPlan) GetRootNode() int64 {
	if m != nil {
		return m.RootNode
	}
	return 0
}

func (m *EncodedQueryPlan) GetOriginalExpression() string {
	if m != nil {
		return m.OriginalExpression
	}
	return ""
}

type EncodedQueryTimeRange struct {
	StartT               int64 `protobuf:"varint,1,opt,name=startT,proto3" json:"startT,omitempty"`
	EndT                 int64 `protobuf:"varint,2,opt,name=endT,proto3" json:"endT,omitempty"`
	IntervalMilliseconds int64 `protobuf:"varint,3,opt,name=intervalMilliseconds,proto3" json:"intervalMilliseconds,omitempty"`
	IsInstant            bool  `protobuf:"varint,4,opt,name=isInstant,proto3" json:"isInstant,omitempty"`
}

func (m *EncodedQueryTimeRange) Reset()      { *m = EncodedQueryTimeRange{} }
func (*EncodedQueryTimeRange) ProtoMessage() {}
func (*EncodedQueryTimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{1}
}
func (m *EncodedQueryTimeRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncodedQueryTimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncodedQueryTimeRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncodedQueryTimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedQueryTimeRange.Merge(m, src)
}
func (m *EncodedQueryTimeRange) XXX_Size() int {
	return m.Size()
}
func (m *EncodedQueryTimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedQueryTimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedQueryTimeRange proto.InternalMessageInfo

func (m *EncodedQueryTimeRange) GetStartT() int64 {
	if m != nil {
		return m.StartT
	}
	return 0
}

func (m *EncodedQueryTimeRange) GetEndT() int64 {
	if m != nil {
		return m.EndT
	}
	return 0
}

func (m *EncodedQueryTimeRange) GetIntervalMilliseconds() int64 {
	if m != nil {
		return m.IntervalMilliseconds
	}
	return 0
}

func (m *EncodedQueryTimeRange) GetIsInstant() bool {
	if m != nil {
		return m.IsInstant
	}
	return false
}

type EncodedNode struct {
	// Why use Any rather than a oneof?
	// This allows us to define nodes in multiple packages without creating dependencies between them,
	// and allows us to define new kinds of nodes in other repositories (eg. the one used for Grafana Cloud Metrics).
	// TODO: is this the best type to use? Any will encode a string for the type of message, which seems expensive - could we use an integer perhaps?
	Details  *types.Any `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	Children []int64    `protobuf:"varint,2,rep,packed,name=children,proto3" json:"children,omitempty"`
	// The following fields will only be populated if this plan was generated with descriptions enabled.
	Type           string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Description    string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ChildrenLabels []string `protobuf:"bytes,5,rep,name=childrenLabels,proto3" json:"childrenLabels,omitempty"`
}

func (m *EncodedNode) Reset()      { *m = EncodedNode{} }
func (*EncodedNode) ProtoMessage() {}
func (*EncodedNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{2}
}
func (m *EncodedNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncodedNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncodedNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncodedNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedNode.Merge(m, src)
}
func (m *EncodedNode) XXX_Size() int {
	return m.Size()
}
func (m *EncodedNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedNode.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedNode proto.InternalMessageInfo

func (m *EncodedNode) GetDetails() *types.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *EncodedNode) GetChildren() []int64 {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *EncodedNode) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EncodedNode) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EncodedNode) GetChildrenLabels() []string {
	if m != nil {
		return m.ChildrenLabels
	}
	return nil
}

func init() {
	proto.RegisterType((*EncodedQueryPlan)(nil), "planning.EncodedQueryPlan")
	proto.RegisterType((*EncodedQueryTimeRange)(nil), "planning.EncodedQueryTimeRange")
	proto.RegisterType((*EncodedNode)(nil), "planning.EncodedNode")
}

func init() { proto.RegisterFile("plan.proto", fileDescriptor_2d655ab2f7683c23) }

var fileDescriptor_2d655ab2f7683c23 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x6e, 0x13, 0x31,
	0x14, 0x1d, 0x77, 0xd2, 0x92, 0x71, 0x24, 0x84, 0xac, 0x16, 0x0d, 0x11, 0x72, 0x47, 0x59, 0xa0,
	0x91, 0x10, 0x53, 0x29, 0x70, 0x01, 0x8a, 0xba, 0x40, 0x02, 0x04, 0x56, 0x2e, 0xe0, 0x8c, 0xcd,
	0xd4, 0x92, 0xfb, 0x3d, 0xb2, 0x1d, 0x44, 0x76, 0x1c, 0x81, 0x0d, 0x77, 0xe0, 0x06, 0x5c, 0xa1,
	0xcb, 0x2c, 0xbb, 0x42, 0x64, 0xb2, 0x61, 0xd9, 0x23, 0xa0, 0x78, 0x32, 0x4d, 0x44, 0xb3, 0x7b,
	0xef, 0xfd, 0x67, 0xfb, 0xbf, 0x67, 0x8c, 0x6b, 0xcd, 0xa1, 0xa8, 0xad, 0xf1, 0x86, 0xf4, 0xd7,
	0x18, 0x14, 0x54, 0xc3, 0x17, 0x95, 0xf2, 0x97, 0xb3, 0x69, 0x51, 0x9a, 0xab, 0xb3, 0xca, 0x54,
	0xe6, 0x2c, 0x18, 0xa6, 0xb3, 0xcf, 0x81, 0x05, 0x12, 0x50, 0x7b, 0x70, 0xf8, 0xa4, 0x32, 0xa6,
	0xd2, 0x72, 0xeb, 0xe2, 0x30, 0x6f, 0x47, 0xa3, 0x05, 0xc2, 0x8f, 0x2e, 0xa0, 0x34, 0x42, 0x8a,
	0x4f, 0x33, 0x69, 0xe7, 0x1f, 0x35, 0x07, 0xf2, 0x06, 0x27, 0x5e, 0x5d, 0x49, 0xc6, 0xa1, 0x92,
	0x29, 0xca, 0x50, 0x3e, 0x18, 0x9f, 0x16, 0xdd, 0xe3, 0xc5, 0xae, 0x7d, 0xd2, 0xd9, 0xce, 0x7b,
	0xd7, 0xbf, 0x4f, 0x23, 0xb6, 0x3d, 0x47, 0x9e, 0xe3, 0x43, 0x30, 0x42, 0xba, 0xf4, 0x20, 0x8b,
	0xf3, 0xc1, 0xf8, 0xe4, 0xde, 0x05, 0x1f, 0x8c, 0x90, 0xac, 0xf5, 0x90, 0x21, 0xee, 0x5b, 0x63,
	0xfc, 0x5a, 0x4a, 0xe3, 0x0c, 0xe5, 0x31, 0xbb, 0xe3, 0xa4, 0xc0, 0xc4, 0x58, 0x55, 0x29, 0xe0,
	0xfa, 0xe2, 0x6b, 0x6d, 0xa5, 0x73, 0xca, 0x40, 0xda, 0xcb, 0x50, 0x9e, 0xb0, 0x3d, 0x93, 0xd1,
	0x0f, 0x84, 0x4f, 0xf6, 0xee, 0x48, 0x1e, 0xe3, 0x23, 0xe7, 0xb9, 0xf5, 0x93, 0x10, 0x2a, 0x66,
	0x1b, 0x46, 0x08, 0xee, 0x49, 0x10, 0x93, 0xf4, 0x20, 0xa8, 0x01, 0x93, 0x31, 0x3e, 0x56, 0xe0,
	0xa5, 0xfd, 0xc2, 0xf5, 0x7b, 0xa5, 0xb5, 0x72, 0xb2, 0x34, 0x20, 0xdc, 0x66, 0xbb, 0xbd, 0x33,
	0xf2, 0x14, 0x27, 0xca, 0xbd, 0x05, 0xe7, 0x39, 0xf8, 0xb0, 0x60, 0x9f, 0x6d, 0x85, 0xd1, 0x2f,
	0x84, 0x07, 0x3b, 0xd1, 0x49, 0x81, 0x1f, 0x08, 0xe9, 0xb9, 0xd2, 0x6e, 0xd3, 0xf1, 0x71, 0xd1,
	0xfe, 0x53, 0xd1, 0xfd, 0x53, 0xf1, 0x1a, 0xe6, 0xac, 0x33, 0xad, 0x3b, 0x2a, 0x2f, 0x95, 0x16,
	0x56, 0x42, 0xe8, 0x34, 0x66, 0x77, 0x7c, 0x9d, 0xc0, 0xcf, 0xeb, 0xb6, 0xbb, 0x84, 0x05, 0x4c,
	0x32, 0x3c, 0x10, 0xd2, 0x95, 0x56, 0xd5, 0x7e, 0x5b, 0xd8, 0xae, 0x44, 0x9e, 0xe1, 0x87, 0xdd,
	0x0d, 0xef, 0xf8, 0x54, 0x6a, 0x97, 0x1e, 0x66, 0x71, 0x9e, 0xb0, 0xff, 0xd4, 0xf3, 0x57, 0x8b,
	0x25, 0x8d, 0x6e, 0x96, 0x34, 0xba, 0x5d, 0x52, 0xf4, 0xad, 0xa1, 0xe8, 0x67, 0x43, 0xd1, 0x75,
	0x43, 0xd1, 0xa2, 0xa1, 0xe8, 0x4f, 0x43, 0xd1, 0xdf, 0x86, 0x46, 0xb7, 0x0d, 0x45, 0xdf, 0x57,
	0x34, 0x5a, 0xac, 0x68, 0x74, 0xb3, 0xa2, 0xd1, 0xf4, 0x28, 0xc4, 0x78, 0xf9, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x6b, 0xae, 0xc9, 0xc3, 0x02, 0x00, 0x00,
}

func (this *EncodedQueryPlan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EncodedQueryPlan)
	if !ok {
		that2, ok := that.(EncodedQueryPlan)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.TimeRange.Equal(&that1.TimeRange) {
		return false
	}
	if len(this.Nodes) != len(that1.Nodes) {
		return false
	}
	for i := range this.Nodes {
		if !this.Nodes[i].Equal(that1.Nodes[i]) {
			return false
		}
	}
	if this.RootNode != that1.RootNode {
		return false
	}
	if this.OriginalExpression != that1.OriginalExpression {
		return false
	}
	return true
}
func (this *EncodedQueryTimeRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EncodedQueryTimeRange)
	if !ok {
		that2, ok := that.(EncodedQueryTimeRange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.StartT != that1.StartT {
		return false
	}
	if this.EndT != that1.EndT {
		return false
	}
	if this.IntervalMilliseconds != that1.IntervalMilliseconds {
		return false
	}
	if this.IsInstant != that1.IsInstant {
		return false
	}
	return true
}
func (this *EncodedNode) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EncodedNode)
	if !ok {
		that2, ok := that.(EncodedNode)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Details.Equal(that1.Details) {
		return false
	}
	if len(this.Children) != len(that1.Children) {
		return false
	}
	for i := range this.Children {
		if this.Children[i] != that1.Children[i] {
			return false
		}
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.ChildrenLabels) != len(that1.ChildrenLabels) {
		return false
	}
	for i := range this.ChildrenLabels {
		if this.ChildrenLabels[i] != that1.ChildrenLabels[i] {
			return false
		}
	}
	return true
}
func (this *EncodedQueryPlan) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&planning.EncodedQueryPlan{")
	s = append(s, "TimeRange: "+strings.Replace(this.TimeRange.GoString(), `&`, ``, 1)+",\n")
	if this.Nodes != nil {
		s = append(s, "Nodes: "+fmt.Sprintf("%#v", this.Nodes)+",\n")
	}
	s = append(s, "RootNode: "+fmt.Sprintf("%#v", this.RootNode)+",\n")
	s = append(s, "OriginalExpression: "+fmt.Sprintf("%#v", this.OriginalExpression)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *EncodedQueryTimeRange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&planning.EncodedQueryTimeRange{")
	s = append(s, "StartT: "+fmt.Sprintf("%#v", this.StartT)+",\n")
	s = append(s, "EndT: "+fmt.Sprintf("%#v", this.EndT)+",\n")
	s = append(s, "IntervalMilliseconds: "+fmt.Sprintf("%#v", this.IntervalMilliseconds)+",\n")
	s = append(s, "IsInstant: "+fmt.Sprintf("%#v", this.IsInstant)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *EncodedNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&planning.EncodedNode{")
	if this.Details != nil {
		s = append(s, "Details: "+fmt.Sprintf("%#v", this.Details)+",\n")
	}
	s = append(s, "Children: "+fmt.Sprintf("%#v", this.Children)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "ChildrenLabels: "+fmt.Sprintf("%#v", this.ChildrenLabels)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPlan(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *EncodedQueryPlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncodedQueryPlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncodedQueryPlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OriginalExpression) > 0 {
		i -= len(m.OriginalExpression)
		copy(dAtA[i:], m.OriginalExpression)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.OriginalExpression)))
		i--
		dAtA[i] = 0x22
	}
	if m.RootNode != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.RootNode))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.TimeRange.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EncodedQueryTimeRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncodedQueryTimeRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncodedQueryTimeRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsInstant {
		i--
		if m.IsInstant {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.IntervalMilliseconds != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.IntervalMilliseconds))
		i--
		dAtA[i] = 0x18
	}
	if m.EndT != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.EndT))
		i--
		dAtA[i] = 0x10
	}
	if m.StartT != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.StartT))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EncodedNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncodedNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncodedNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChildrenLabels) > 0 {
		for iNdEx := len(m.ChildrenLabels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChildrenLabels[iNdEx])
			copy(dAtA[i:], m.ChildrenLabels[iNdEx])
			i = encodeVarintPlan(dAtA, i, uint64(len(m.ChildrenLabels[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Children) > 0 {
		dAtA3 := make([]byte, len(m.Children)*10)
		var j2 int
		for _, num1 := range m.Children {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintPlan(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x12
	}
	if m.Details != nil {
		{
			size, err := m.Details.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlan(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EncodedQueryPlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TimeRange.Size()
	n += 1 + l + sovPlan(uint64(l))
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	if m.RootNode != 0 {
		n += 1 + sovPlan(uint64(m.RootNode))
	}
	l = len(m.OriginalExpression)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	return n
}

func (m *EncodedQueryTimeRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartT != 0 {
		n += 1 + sovPlan(uint64(m.StartT))
	}
	if m.EndT != 0 {
		n += 1 + sovPlan(uint64(m.EndT))
	}
	if m.IntervalMilliseconds != 0 {
		n += 1 + sovPlan(uint64(m.IntervalMilliseconds))
	}
	if m.IsInstant {
		n += 2
	}
	return n
}

func (m *EncodedNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		l = m.Details.Size()
		n += 1 + l + sovPlan(uint64(l))
	}
	if len(m.Children) > 0 {
		l = 0
		for _, e := range m.Children {
			l += sovPlan(uint64(e))
		}
		n += 1 + sovPlan(uint64(l)) + l
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if len(m.ChildrenLabels) > 0 {
		for _, s := range m.ChildrenLabels {
			l = len(s)
			n += 1 + l + sovPlan(uint64(l))
		}
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *EncodedQueryPlan) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForNodes := "[]*EncodedNode{"
	for _, f := range this.Nodes {
		repeatedStringForNodes += strings.Replace(f.String(), "EncodedNode", "EncodedNode", 1) + ","
	}
	repeatedStringForNodes += "}"
	s := strings.Join([]string{`&EncodedQueryPlan{`,
		`TimeRange:` + strings.Replace(strings.Replace(this.TimeRange.String(), "EncodedQueryTimeRange", "EncodedQueryTimeRange", 1), `&`, ``, 1) + `,`,
		`Nodes:` + repeatedStringForNodes + `,`,
		`RootNode:` + fmt.Sprintf("%v", this.RootNode) + `,`,
		`OriginalExpression:` + fmt.Sprintf("%v", this.OriginalExpression) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EncodedQueryTimeRange) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EncodedQueryTimeRange{`,
		`StartT:` + fmt.Sprintf("%v", this.StartT) + `,`,
		`EndT:` + fmt.Sprintf("%v", this.EndT) + `,`,
		`IntervalMilliseconds:` + fmt.Sprintf("%v", this.IntervalMilliseconds) + `,`,
		`IsInstant:` + fmt.Sprintf("%v", this.IsInstant) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EncodedNode) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EncodedNode{`,
		`Details:` + strings.Replace(fmt.Sprintf("%v", this.Details), "Any", "types.Any", 1) + `,`,
		`Children:` + fmt.Sprintf("%v", this.Children) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`ChildrenLabels:` + fmt.Sprintf("%v", this.ChildrenLabels) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPlan(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *EncodedQueryPlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: EncodedQueryPlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncodedQueryPlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &EncodedNode{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootNode", wireType)
			}
			m.RootNode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RootNode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalExpression", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalExpression = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func (m *EncodedQueryTimeRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: EncodedQueryTimeRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncodedQueryTimeRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartT", wireType)
			}
			m.StartT = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartT |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndT", wireType)
			}
			m.EndT = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndT |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntervalMilliseconds", wireType)
			}
			m.IntervalMilliseconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntervalMilliseconds |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsInstant", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsInstant = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func (m *EncodedNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: EncodedNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncodedNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Details == nil {
				m.Details = &types.Any{}
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPlan
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Children = append(m.Children, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPlan
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
					return ErrInvalidLengthPlan
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPlan
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
				if elementCount != 0 && len(m.Children) == 0 {
					m.Children = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPlan
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Children = append(m.Children, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildrenLabels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildrenLabels = append(m.ChildrenLabels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
