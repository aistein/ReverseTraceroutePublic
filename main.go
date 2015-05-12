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
package main

import (
	"flag"
	"fmt"
	dm "github.com/NEU-SNS/ReverseTraceroute/lib/datamodel"
	"net"
	"net/rpc/jsonrpc"
	"runtime"
)

const (
	PING       = "ControllerApi.Ping"
	TRACEROUTE = "ControllerApi.Traceroute"
	GETSTATS   = "ControllerApi.GetStats"
	REGISTER   = "ControllerApi.Register"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	conn, err := net.Dial("tcp", "localhost:35000")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := jsonrpc.NewClient(conn)
	args := dm.PingArg{ServiceArg: dm.ServiceArg{dm.PLANET_LAB}, Dst: "129.10.113.204",
		Host: "129.10.113.205", RR: false}
	var ret dm.PingReturn
	err = c.Call(PING, args, &ret)
	if err != nil {
		fmt.Printf("Ping failed with err: %v\n", err)
	}
	fmt.Printf("Response took: %s\n", ret.Dur.String())
	a := dm.TracerouteArg{ServiceArg: dm.ServiceArg{dm.PLANET_LAB}, Dst: "8.8.8.8",
		Host: "129.10.113.205"}
	var r dm.TracerouteReturn
	err = c.Call(TRACEROUTE, a, &r)
	if err != nil {
		fmt.Printf("Traceroute failed with err: %v\n", err)
	}
	fmt.Printf("Response took: %s\n", r.Dur.String())
	arg := dm.StatsArg{ServiceArg: dm.ServiceArg{dm.PLANET_LAB}}
	var rr dm.StatsReturn
	err = c.Call(GETSTATS, arg, &rr)
	if err != nil {
		fmt.Printf("Stats failed with err: %v\n", err)
	}
	fmt.Printf("Got back: %v\n", rr)
}
