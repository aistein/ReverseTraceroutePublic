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
package plcontroller

import (
	dm "github.com/NEU-SNS/ReverseTraceroute/lib/datamodel"
	"github.com/golang/glog"
)

type PlControllerApi struct{}

func (c PlControllerApi) Ping(arg dm.PingArg, ret *dm.Ping) error {
	glog.Info("Ping Called")
	ping, err := plController.runPing(arg)
	if err != nil {
		return err
	}
	*ret = ping
	return nil
}

func (c PlControllerApi) Traceroute(arg dm.TracerouteArg, ret *dm.Traceroute) error {
	glog.Info("Traceroute Called")
	tr, err := plController.runTraceroute(arg)
	if err != nil {
		return err
	}
	*ret = tr
	return nil
}

func (c PlControllerApi) GetStats(arg struct{}, ret *dm.Stats) error {
	glog.Infof("GetStats Called")
	stat := plController.getStats()
	*ret = stat
	return nil
}