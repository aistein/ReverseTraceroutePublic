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
// source: github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/stats.proto
// DO NOT EDIT!

package datamodel

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Stats struct {
	StartTime  int64 `protobuf:"varint,1,opt,name=start_time" json:"start_time,omitempty"`
	UpTime     int64 `protobuf:"varint,2,opt,name=up_time" json:"up_time,omitempty"`
	Requests   int64 `protobuf:"varint,3,opt,name=requests" json:"requests,omitempty"`
	AvgReqTime int64 `protobuf:"varint,4,opt,name=avg_req_time" json:"avg_req_time,omitempty"`
	TotReqTime int64 `protobuf:"varint,5,opt,name=tot_req_time" json:"tot_req_time,omitempty"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}

type StatsArg struct {
	Service ServiceT `protobuf:"varint,1,opt,name=service,enum=datamodel.ServiceT" json:"service,omitempty"`
	Ip      string   `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
}

func (m *StatsArg) Reset()         { *m = StatsArg{} }
func (m *StatsArg) String() string { return proto.CompactTextString(m) }
func (*StatsArg) ProtoMessage()    {}

type StatsReturn struct {
	Ret   *ReturnT `protobuf:"bytes,1,opt,name=ret" json:"ret,omitempty"`
	Stats *Stats   `protobuf:"bytes,2,opt,name=stats" json:"stats,omitempty"`
}

func (m *StatsReturn) Reset()         { *m = StatsReturn{} }
func (m *StatsReturn) String() string { return proto.CompactTextString(m) }
func (*StatsReturn) ProtoMessage()    {}

func (m *StatsReturn) GetRet() *ReturnT {
	if m != nil {
		return m.Ret
	}
	return nil
}

func (m *StatsReturn) GetStats() *Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
}
