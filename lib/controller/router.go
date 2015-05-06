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
package controller

import (
	"fmt"
	dm "github.com/NEU-SNS/ReverseTraceroute/lib/datamodel"
	"github.com/golang/glog"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

type router struct {
	rw       sync.RWMutex
	services map[dm.ServiceT]*dm.Service
}

func createRouter() Router {
	s := make(map[dm.ServiceT]*dm.Service)
	return &router{services: s}
}

func NewRouter() Router {
	return createRouter()
}

type Router interface {
	RegisterServices(services ...*dm.Service)
	RouteRequest(req Request) (RoutedRequest, error)
	GetServices() []*dm.Service
}

func (r *router) GetServices() []*dm.Service {
	r.rw.RLock()
	serv := make([]*dm.Service, len(r.services), len(r.services))
	for _, service := range r.services {
		serv = append(serv, service)
	}
	r.rw.RUnlock()
	return serv
}

func (r *router) RegisterServices(services ...*dm.Service) {
	r.rw.Lock()
	for _, service := range services {
		r.services[service.Key] = service
		glog.Infof("Registered service: %v, with api: %v", service, service.Api)
	}
	r.rw.Unlock()
}

func (r *router) RouteRequest(req Request) (RoutedRequest, error) {
	r.rw.RLock()
	glog.Infof("Trying to get API for %s", req.Key)
	s := r.services[req.Key]
	if s == nil {
		return nil, ErrorServiceNotFound
	}
	r.rw.RUnlock()
	return wrapRequest(req, s)
}

func wrapRequest(req Request, s *dm.Service) (RoutedRequest, error) {
	glog.Infof("Wrapping request: %v", req)
	return func() (*dm.MReturn, Request, error) {
		req.Stime = time.Now()
		glog.Infof("Executing request: %v", req)
		glog.Infof("Connecting to %s, %s", s.Proto, s.FormatIp())
		c, err := jsonrpc.Dial(s.Proto, s.FormatIp())
		if err != nil {
			glog.Errorf("Failed to connect: %v, %v, with err: %v", req, s, err)
			return nil, req, err
		}
		defer c.Close()
		ret := new(dm.MReturn)
		api := s.Api[req.Type]
		sretf, ok := dm.TypeMap[api.Type]
		if !ok {
			glog.Errorf("Could not find func for apiType: %s", api.Type)
			return nil, req, fmt.Errorf("Failed to find Return type")
		}
		sret := sretf()
		err = c.Call(api.Url, req.Args, sret)
		glog.Info("%v", sret)
		req.Dur = time.Since(req.Stime)
		ret.Dur = req.Dur
		ret.SRet = sret
		glog.Infof("Finished Request: %v", req)
		if err != nil {
			return nil, req, err
		}
		return ret, req, nil
	}, nil

}