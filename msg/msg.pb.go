// Code generated by protoc-gen-go.
// source: msg/msg.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg/msg.proto

It has these top-level messages:
	Any
	Sih
*/
package msg

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Any struct {
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url" json:"type_url,omitempty"`
	Value   []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}

type Sih struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *Any   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Sih) Reset()         { *m = Sih{} }
func (m *Sih) String() string { return proto.CompactTextString(m) }
func (*Sih) ProtoMessage()    {}

func (m *Sih) GetValue() *Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
}