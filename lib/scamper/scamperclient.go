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

// Package scamper is a library to work with scamper control sockets
package scamper

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/NEU-SNS/ReverseTraceroute/lib/util"
	"github.com/golang/glog"
)

// SResponseT represents the type of responses scamper can send
type SResponseT string

const (
	// OK is the accept response from scamper
	OK SResponseT = "OK"
	// MORE is the response when more commands can be given
	MORE SResponseT = "MORE"
	// DATA represensts a data message
	DATA SResponseT = "DATA"
	// ERR is the error response from scamper
	ERR       SResponseT = "ERR"
	cancelCmd            = "halt %d"
)

var (
	// ErrorBadDataResponse is returned when the data received by scamper couldnt
	// be converted
	ErrorBadDataResponse = errors.New("Bad DATA Response")
	// ErrorBadOKResponse is returned when an OK response fails to parse
	ErrorBadOKResponse = errors.New("Bad OK Response")
	// ErrorBadResponse is the generic error when a response can't be parsed
	ErrorBadResponse = errors.New("Bad Response")
	// ErrorTimeout returned when a command times out
	ErrorTimeout = errors.New("Timeout")
)

//TODO: A possible optimization to make is to open a single connection,
// resuse it until it fails, and using the -U option to assign an id to
// each measurement and then return them to the proper caller. This
// would also involve blocking on a conn waiting for data at any time
// measurements are out.

// Response represents a response from scamper
type Response struct {
	rType SResponseT
	data  []byte
	ds    int
	id    int
}

/*
type Client struct {
	rw    *bufio.ReadWriter
	s     Socket
	cmd   Cmd
	resps []Response
	conn  *net.Conn
	id    int
}
*/

type socketMap map[string]Socket

// Client is the main object for interacting with scamper
type Client struct {
	sockets socketMap
}

// Bytes get the data as bytes from a scamper response
func (r Response) Bytes() []byte {
	return r.data
}

// WriteTo writes the response to the given io.Writer
func (r Response) WriteTo(w io.Writer) (n int64, err error) {
	glog.Infof("Writing data %v", r.data)
	c, err := w.Write(r.data)
	n = int64(c)
	glog.Infof("Wrote %d bytes", n)
	return
}

// NewClient creates a new Client
func NewClient(s Socket, c Cmd) *Client {
	return &Client{s: s, cmd: c, resps: make([]Response, 0)}
}

// GetResponses gets the raw responses from the Client
func (c *Client) GetResponses() []Response {
	return c.resps
}

func (c *Client) checkConn() error {
	if !c.connected() {
		return c.connect()
	}
	return nil
}

// CancelCmd cancels the running scamper command
func (c *Client) CancelCmd() error {
	glog.Info("Canceling command: %d", c.id)
	err := c.checkConn()
	if err != nil {
		return err
	}
	defer c.closeConnection()
	cstring := fmt.Sprintf(cancelCmd, c.id)
	_, err = c.rw.WriteString(cstring)
	if err != nil {
		return err
	}
	glog.Flush()
	return c.rw.Flush()
}

// IssueCmd runs the command on scamper
func (c *Client) IssueCmd(ec chan error, dc chan struct{}) {
	glog.Infof("Issuing command: %s", c.cmd.String())
	err := c.checkConn()
	if err != nil {
		ec <- err
		return
	}
	defer c.closeConnection()
	_, err = c.rw.WriteString(c.cmd.String())
	if err != nil {
		ec <- err
		return
	}
	c.rw.Flush()
	i := 0
	for i < 3 {
		line, err := c.rw.ReadString('\n')
		if err != nil {
			ec <- err
			return
		}
		r, err := parseResponse(line, c.rw)
		if err != nil {
			glog.Errorf("Error parsing response: %v", err)
			ec <- err
			return
		}
		switch r.rType {
		case OK:
			c.id = r.id
		case DATA:
			glog.Infof("Parsed data response")
			c.resps = append(c.resps, r)
			i++
			glog.Infof("Count of data received: %d", i)
		case ERR:
			glog.Errorf("Parsed scamper ERR return")
			ec <- fmt.Errorf("Error with scamper request: %s", c.cmd.String())
			return
		case MORE:
		}

	}
	close(dc)
	return
}

func (c *Client) connect() error {
	glog.Infof("Connecting to: %s", c.s.fname)
	conn, err := net.Dial("unix", c.s.fname)
	if err != nil {
		return err
	}
	c.conn = &conn
	c.rw = util.ConnToRW(conn)
	return nil
}

func (c *Client) closeConnection() error {
	glog.Infof("Closing connection to: %s", c.s.fname)
	if c.connected() {
		err := (*c.conn).Close()
		if err != nil {
			return err
		}
		c.conn = nil
	}
	return nil
}

func (c *Client) connected() bool {
	return c.rw != nil
}

func parseResponse(r string, rw *bufio.ReadWriter) (Response, error) {
	resp := Response{}
	glog.Infof("Parsing Response")
	switch {
	case strings.Contains(r, string(OK)):
		glog.Info("Parsed OK response")
		resp.rType = OK
		r = strings.TrimSpace(r)
		split := strings.Split(r, " ")
		idsp := strings.Split(split[1], "-")
		glog.Infof("Got id info: %v", idsp)
		id, err := strconv.Atoi(idsp[1])
		if err != nil {
			glog.Errorf("Failed to parse response: %v", r)
			return resp, ErrorBadResponse
		}
		resp.id = id
		return resp, nil
	case strings.Contains(r, string(ERR)):
		glog.Error("Parsed error response")
		resp.rType = ERR
		return resp, nil
	case strings.Contains(r, string(DATA)):
		glog.Info("Parsed DATA response")
		resp.rType = DATA
		split := strings.Split(r, " ")
		if len(split) != 2 {
			glog.Errorf("Failed to parse response: %v", r)
			return resp, ErrorBadDataResponse
		}
		n, err := strconv.Atoi(split[1][:len(split[1])-1])
		if err != nil {
			return resp, err
		}
		resp.ds = n
		buff := make([]byte, n)
		_, err = io.ReadFull(rw, buff)
		if err != nil {

			return resp, err
		}
		resp.data = buff
		glog.Infof("Parsed data response, len: %d, data: %s", n, buff)
		return resp, nil
	case strings.Contains(r, string(MORE)):
		glog.Info("Parsed MORE response")
		resp.rType = MORE
		return resp, nil
	}
	glog.Errorf("Failed to parse response: %v", r)
	return resp, ErrorBadResponse
}
