/*
Copyright (c) 2015, Northeastern University
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:
     * Redistributions of source code must retain the above copyright
       notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above copyright
       notice, this list of conditions and the following disclaimer in the
       documentation and/or other materials provided with the distribution.
     * Neither the name of the Northeastern University nor the
       names of its contributors may be used to endorse or promote products
       derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL Northeastern University BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/datamodel/ping.proto
// DO NOT EDIT!

package datamodel

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type PingMeasurement struct {
	Src        string `protobuf:"bytes,1,opt,name=src" json:"src,omitempty"`
	Dst        string `protobuf:"bytes,2,opt,name=dst" json:"dst,omitempty"`
	Host       string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	Spoof      bool   `protobuf:"varint,4,opt,name=spoof" json:"spoof,omitempty"`
	RR         bool   `protobuf:"varint,5,opt" json:"RR,omitempty"`
	SAddr      string `protobuf:"bytes,6,opt,name=s_addr" json:"s_addr,omitempty"`
	Payload    string `protobuf:"bytes,7,opt,name=payload" json:"payload,omitempty"`
	Count      string `protobuf:"bytes,8,opt,name=count" json:"count,omitempty"`
	IcmpSum    string `protobuf:"bytes,9,opt,name=icmp_sum" json:"icmp_sum,omitempty"`
	Dport      string `protobuf:"bytes,10,opt,name=dport" json:"dport,omitempty"`
	Sport      string `protobuf:"bytes,11,opt,name=sport" json:"sport,omitempty"`
	Wait       string `protobuf:"bytes,12,opt,name=wait" json:"wait,omitempty"`
	Ttl        string `protobuf:"bytes,13,opt,name=ttl" json:"ttl,omitempty"`
	Mut        string `protobuf:"bytes,14,opt,name=mut" json:"mut,omitempty"`
	ReplyCount string `protobuf:"bytes,15,opt,name=reply_count" json:"reply_count,omitempty"`
	Pattern    string `protobuf:"bytes,16,opt,name=pattern" json:"pattern,omitempty"`
	Method     string `protobuf:"bytes,17,opt,name=method" json:"method,omitempty"`
	Size       string `protobuf:"bytes,18,opt,name=size" json:"size,omitempty"`
	UserId     string `protobuf:"bytes,19,opt,name=user_id" json:"user_id,omitempty"`
	Tos        string `protobuf:"bytes,20,opt,name=tos" json:"tos,omitempty"`
	TimeStamp  string `protobuf:"bytes,21,opt,name=time_stamp" json:"time_stamp,omitempty"`
	Timeout    int64  `protobuf:"varint,22,opt,name=timeout" json:"timeout,omitempty"`
	CheckCache bool   `protobuf:"varint,23,opt,name=check_cache" json:"check_cache,omitempty"`
}

func (m *PingMeasurement) Reset()         { *m = PingMeasurement{} }
func (m *PingMeasurement) String() string { return proto.CompactTextString(m) }
func (*PingMeasurement) ProtoMessage()    {}

type PingArg struct {
	Pings []*PingMeasurement `protobuf:"bytes,1,rep,name=pings" json:"pings,omitempty"`
}

func (m *PingArg) Reset()         { *m = PingArg{} }
func (m *PingArg) String() string { return proto.CompactTextString(m) }
func (*PingArg) ProtoMessage()    {}

func (m *PingArg) GetPings() []*PingMeasurement {
	if m != nil {
		return m.Pings
	}
	return nil
}

type PingStats struct {
	Replies int32   `protobuf:"varint,1,opt,name=replies" json:"replies,omitempty"`
	Loss    int32   `protobuf:"varint,2,opt,name=loss" json:"loss,omitempty"`
	Min     float32 `protobuf:"fixed32,3,opt,name=min" json:"min,omitempty"`
	Max     float32 `protobuf:"fixed32,4,opt,name=max" json:"max,omitempty"`
	Avg     float32 `protobuf:"fixed32,5,opt,name=avg" json:"avg,omitempty"`
	Stddev  float32 `protobuf:"fixed32,6,opt,name=stddev" json:"stddev,omitempty"`
}

func (m *PingStats) Reset()         { *m = PingStats{} }
func (m *PingStats) String() string { return proto.CompactTextString(m) }
func (*PingStats) ProtoMessage()    {}

type PingResponse struct {
	From       string       `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Seq        int32        `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	ReplySize  int32        `protobuf:"varint,3,opt,name=reply_size" json:"reply_size,omitempty"`
	ReplyTtl   int32        `protobuf:"varint,4,opt,name=reply_ttl" json:"reply_ttl,omitempty"`
	ReplyProto string       `protobuf:"bytes,5,opt,name=reply_proto" json:"reply_proto,omitempty"`
	Tx         *Time        `protobuf:"bytes,6,opt,name=tx" json:"tx,omitempty"`
	Rx         *Time        `protobuf:"bytes,7,opt,name=rx" json:"rx,omitempty"`
	Rtt        float32      `protobuf:"fixed32,8,opt,name=rtt" json:"rtt,omitempty"`
	ProbeIpid  int32        `protobuf:"varint,9,opt,name=probe_ipid" json:"probe_ipid,omitempty"`
	ReplyIpid  int32        `protobuf:"varint,10,opt,name=reply_ipid" json:"reply_ipid,omitempty"`
	IcmpType   int32        `protobuf:"varint,11,opt,name=icmp_type" json:"icmp_type,omitempty"`
	IcmpCode   int32        `protobuf:"varint,12,opt,name=icmp_code" json:"icmp_code,omitempty"`
	RR         []string     `protobuf:"bytes,13,rep" json:"RR,omitempty"`
	Tsonly     []int64      `protobuf:"varint,14,rep,name=tsonly" json:"tsonly,omitempty"`
	Tsandaddr  []*TsAndAddr `protobuf:"bytes,15,rep,name=tsandaddr" json:"tsandaddr,omitempty"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}

func (m *PingResponse) GetTx() *Time {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *PingResponse) GetRx() *Time {
	if m != nil {
		return m.Rx
	}
	return nil
}

func (m *PingResponse) GetTsandaddr() []*TsAndAddr {
	if m != nil {
		return m.Tsandaddr
	}
	return nil
}

type TsAndAddr struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Ts int64  `protobuf:"varint,2,opt,name=ts" json:"ts,omitempty"`
}

func (m *TsAndAddr) Reset()         { *m = TsAndAddr{} }
func (m *TsAndAddr) String() string { return proto.CompactTextString(m) }
func (*TsAndAddr) ProtoMessage()    {}

type Ping struct {
	Version    string          `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Type       string          `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Method     string          `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Src        string          `protobuf:"bytes,4,opt,name=src" json:"src,omitempty"`
	Dst        string          `protobuf:"bytes,5,opt,name=dst" json:"dst,omitempty"`
	Start      *Time           `protobuf:"bytes,6,opt,name=start" json:"start,omitempty"`
	PingSent   int32           `protobuf:"varint,7,opt,name=ping_sent" json:"ping_sent,omitempty"`
	ProbeSize  int32           `protobuf:"varint,8,opt,name=probe_size" json:"probe_size,omitempty"`
	Userid     uint32          `protobuf:"varint,9,opt,name=userid" json:"userid,omitempty"`
	Ttl        int32           `protobuf:"varint,10,opt,name=ttl" json:"ttl,omitempty"`
	Wait       int32           `protobuf:"varint,11,opt,name=wait" json:"wait,omitempty"`
	Timeout    int32           `protobuf:"varint,12,opt,name=timeout" json:"timeout,omitempty"`
	Flags      []string        `protobuf:"bytes,13,rep,name=flags" json:"flags,omitempty"`
	Responses  []*PingResponse `protobuf:"bytes,14,rep,name=responses" json:"responses,omitempty"`
	Statistics *PingStats      `protobuf:"bytes,15,opt,name=statistics" json:"statistics,omitempty"`
	Error      string          `protobuf:"bytes,16,opt,name=error" json:"error,omitempty"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}

func (m *Ping) GetStart() *Time {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Ping) GetResponses() []*PingResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

func (m *Ping) GetStatistics() *PingStats {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type PingReturn struct {
	Ret  *ReturnT `protobuf:"bytes,1,opt,name=ret" json:"ret,omitempty"`
	Ping *Ping    `protobuf:"bytes,2,opt,name=ping" json:"ping,omitempty"`
}

func (m *PingReturn) Reset()         { *m = PingReturn{} }
func (m *PingReturn) String() string { return proto.CompactTextString(m) }
func (*PingReturn) ProtoMessage()    {}

func (m *PingReturn) GetRet() *ReturnT {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *PingReturn) GetPing() *Ping {
	if m != nil {
		return m.Ping
	}
	return nil
}

func init() {
}
