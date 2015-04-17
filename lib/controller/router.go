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
     * Neither the name of the University of Washington nor the
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
	dm "github.com/NEU-SNS/ReverseTraceroute/lib/datamodel"
	"net/rpc/jsonrpc"
	"time"
)

type router struct {
	services map[string]*dm.Service
}

func createRouter() Router {
	s := make(map[string]*dm.Service)
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
	serv := make([]*dm.Service, len(r.services), len(r.services))
	for _, service := range r.services {
		serv = append(serv, service)
	}
	return serv
}

func (r *router) RegisterServices(services ...*dm.Service) {
	for _, service := range services {
		r.services[service.Key] = service
	}
}

func (r *router) RouteRequest(req Request) (RoutedRequest, error) {
	s := r.services[req.Key]
	if s == nil {
		return nil, ErrorServiceNotFound
	}
	return wrapRequest(req, s)
}

func wrapRequest(req Request, s *dm.Service) (RoutedRequest, error) {
	return func() (*MReturn, error) {

		req.Stime = time.Now()
		c, err := jsonrpc.Dial(s.Proto, s.FormatIp())
		if err != nil {
			return nil, err
		}
		defer c.Close()
		err = c.Call(s.Api[req.Type], nil, nil)
		req.Dur = time.Since(req.Stime)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}, nil

}
