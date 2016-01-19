// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/datamodel/vantagepoint.proto
// DO NOT EDIT!

package datamodel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VantagePoint struct {
	Hostname     string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Ip           uint32 `protobuf:"varint,2,opt,name=ip" json:"ip,omitempty"`
	Sshable      bool   `protobuf:"varint,3,opt,name=sshable" json:"sshable,omitempty"`
	Timestamp    bool   `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	RecordRoute  bool   `protobuf:"varint,5,opt,name=record_route" json:"record_route,omitempty"`
	LastUpdated  int64  `protobuf:"varint,6,opt,name=last_updated" json:"last_updated,omitempty"`
	CanSpoof     bool   `protobuf:"varint,7,opt,name=can_spoof" json:"can_spoof,omitempty"`
	Controller   uint32 `protobuf:"varint,8,opt,name=controller" json:"controller,omitempty"`
	ReceiveSpoof bool   `protobuf:"varint,9,opt,name=receive_spoof" json:"receive_spoof,omitempty"`
	Site         string `protobuf:"bytes,10,opt,name=site" json:"site,omitempty"`
	SpoofChecked int64  `protobuf:"varint,11,opt,name=spoof_checked" json:"spoof_checked,omitempty"`
	Port         uint32 `protobuf:"varint,12,opt,name=port" json:"port,omitempty"`
}

func (m *VantagePoint) Reset()                    { *m = VantagePoint{} }
func (m *VantagePoint) String() string            { return proto.CompactTextString(m) }
func (*VantagePoint) ProtoMessage()               {}
func (*VantagePoint) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type VPRequest struct {
}

func (m *VPRequest) Reset()                    { *m = VPRequest{} }
func (m *VPRequest) String() string            { return proto.CompactTextString(m) }
func (*VPRequest) ProtoMessage()               {}
func (*VPRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type VPReturn struct {
	Vps []*VantagePoint `protobuf:"bytes,1,rep,name=vps" json:"vps,omitempty"`
}

func (m *VPReturn) Reset()                    { *m = VPReturn{} }
func (m *VPReturn) String() string            { return proto.CompactTextString(m) }
func (*VPReturn) ProtoMessage()               {}
func (*VPReturn) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *VPReturn) GetVps() []*VantagePoint {
	if m != nil {
		return m.Vps
	}
	return nil
}

func init() {
	proto.RegisterType((*VantagePoint)(nil), "datamodel.VantagePoint")
	proto.RegisterType((*VPRequest)(nil), "datamodel.VPRequest")
	proto.RegisterType((*VPReturn)(nil), "datamodel.VPReturn")
}

var fileDescriptor8 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4e, 0xf3, 0x30,
	0x10, 0x84, 0xd5, 0xa6, 0x7f, 0x9b, 0x6c, 0x52, 0xfd, 0x10, 0x81, 0xf0, 0x11, 0x55, 0x1c, 0x7a,
	0x21, 0x41, 0xf0, 0x04, 0x20, 0x71, 0xad, 0xaa, 0x16, 0x7a, 0xe0, 0x12, 0xb9, 0xc9, 0xd2, 0x44,
	0x24, 0x59, 0x63, 0x3b, 0x79, 0x6b, 0xde, 0x81, 0x8d, 0x81, 0xc2, 0xcd, 0x9e, 0xfd, 0xbc, 0x33,
	0x1e, 0xb8, 0x3f, 0x54, 0xb6, 0xec, 0xf6, 0x49, 0x4e, 0x4d, 0xba, 0x7a, 0x7c, 0xbe, 0xde, 0xae,
	0xb6, 0xe9, 0x06, 0x7b, 0xd4, 0x06, 0x9f, 0xb4, 0xcc, 0x51, 0x53, 0x67, 0x31, 0x2d, 0xa4, 0x95,
	0x0d, 0x15, 0x58, 0xa7, 0xbd, 0x6c, 0xad, 0x3c, 0xa0, 0xa2, 0xaa, 0xb5, 0x89, 0xd2, 0x64, 0x29,
	0x0e, 0x8e, 0xd3, 0xc5, 0xc7, 0x08, 0xa2, 0xdd, 0x17, 0xb1, 0x1e, 0x88, 0xf8, 0x04, 0xfc, 0x92,
	0x8c, 0x6d, 0x65, 0x83, 0x62, 0x74, 0x39, 0x5a, 0x06, 0x31, 0xc0, 0xb8, 0x52, 0x62, 0xcc, 0xe7,
	0x79, 0xfc, 0x1f, 0x66, 0xc6, 0x94, 0x72, 0x5f, 0xa3, 0xf0, 0x58, 0xf0, 0xe3, 0x53, 0x08, 0x6c,
	0xd5, 0xa0, 0xe1, 0x7d, 0x4a, 0x4c, 0x9c, 0x74, 0x06, 0x91, 0xc6, 0x9c, 0x74, 0x91, 0xb9, 0x28,
	0xe2, 0xdf, 0x8f, 0x5a, 0x4b, 0x63, 0xb3, 0x4e, 0xb1, 0x39, 0x16, 0x62, 0xca, 0xaa, 0x37, 0x3c,
	0xcf, 0x65, 0x9b, 0x19, 0x45, 0xf4, 0x2a, 0x66, 0x0e, 0x64, 0xbf, 0x9c, 0x5a, 0xab, 0xa9, 0xae,
	0x51, 0x0b, 0xdf, 0xd9, 0x9e, 0xc3, 0x9c, 0x57, 0x62, 0xd5, 0xe3, 0x37, 0x1a, 0x38, 0x34, 0x82,
	0x89, 0xa9, 0xd8, 0x01, 0x5c, 0x4e, 0x86, 0xdc, 0x30, 0xcb, 0x4b, 0xcc, 0xdf, 0xd8, 0x22, 0x74,
	0x16, 0x0c, 0x29, 0xd2, 0x56, 0x44, 0xc3, 0xa6, 0x45, 0x08, 0xc1, 0x6e, 0xbd, 0xc1, 0xf7, 0x8e,
	0x23, 0x2f, 0x6e, 0xc0, 0x1f, 0x2e, 0xb6, 0xd3, 0x6d, 0x7c, 0x05, 0x5e, 0xaf, 0x0c, 0x7f, 0xd9,
	0x5b, 0x86, 0xb7, 0x17, 0xc9, 0xb1, 0xa1, 0xe4, 0x6f, 0x3b, 0x0f, 0xe1, 0xcb, 0x6f, 0x77, 0xfb,
	0xa9, 0x6b, 0xf3, 0xee, 0x33, 0x00, 0x00, 0xff, 0xff, 0x52, 0x77, 0x8b, 0x4b, 0x92, 0x01, 0x00,
	0x00,
}
