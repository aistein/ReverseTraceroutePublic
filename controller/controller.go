package controller

import (
	"fmt"
	da "github.com/NEU-SNS/ReverseTraceroute/dataaccess"
	dm "github.com/NEU-SNS/ReverseTraceroute/datamodel"
	"github.com/golang/glog"
	"github.com/nu7hatch/gouuid"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"strings"
	"time"
)

type controllerT struct {
	port     int
	ip       net.IP
	ptype    string
	db       da.DataAccess
	router   Router
	requests uint64
	time     time.Duration
}

var controller controllerT

func parseAddrArg(addr string) (int, net.IP, error) {
	parts := strings.Split(addr, ":")
	ip := parts[IP]
	port := parts[PORT]
	pport, err := strconv.Atoi(port)
	if err != nil {
		glog.Errorf("Failed to parse port")
		return 0, nil, err
	}
	pip := net.ParseIP(ip)
	if pip == nil {
		glog.Errorf("Invalid IP passed to Start")
		return 0, nil, ErrorInvalidIP
	}
	return pport, pip, nil
}

func Start(n, laddr string, db da.DataAccess) chan error {
	errChan := make(chan error, 1)
	if db == nil {
		glog.Fatalf("Nil db in Controller Start")
	}
	controller.ptype = n
	controller.db = db
	controller.router = createRouter()
	port, ip, err := parseAddrArg(laddr)
	if err != nil {
		glog.Errorf("Failed to start Controller")
		errChan <- err
	}
	controller.ip = ip
	controller.port = port
	go startRpc(n, laddr, errChan)
	return errChan
}

func (m MRequestError) Error() string {
	return fmt.Sprintf("Error occured while %s caused by: %v", m.cause, m.causeErr)
}

func (c ControllerApi) Register(arg int, reply *int) error {
	glog.Info("Register Called")
	*reply = 5
	return nil
}

func (c ControllerApi) Ping(arg MArg, ret *MReturn) error {
	mr, err := controller.handleMeasurement(&arg, PING)
	ret = mr
	return err
}

func (c ControllerApi) Traceroute(arg MArg, ret *MReturn) error {
	mr, err := controller.handleMeasurement(&arg, TRACEROUTE)
	ret = mr
	return err
}

func makeErrorReturn(cause MRequestState, err error) (*MReturn, error) {
	return &MReturn{Status: ERROR}, MRequestError{cause: cause, causeErr: err}
}

func (c controllerT) handleMeasurement(arg *MArg, mt dm.MType) (*MReturn, error) {
	r, err := generateRequest(arg, mt)
	if err != nil {
		glog.Errorf("Error generating request: %v", err)
		return makeErrorReturn(GenRequest, err)
	}
	rr, err := controller.routeRequest(r)
	if err != nil {
		glog.Errorf("%s: Failed to route request: %v, with error: %v", r.Id, r, err)
		return makeErrorReturn(RequestRoute, err)
	}
	result, err := rr()
	if err != nil {
		glog.Errorf("%s: Failed to execute request: %v, with error: %v", r.Id, rr, err)
		return makeErrorReturn(ExecuteRequest, err)
	}
	return result, nil
}

func (c controllerT) routeRequest(r Request) (RoutedRequest, error) {
	rr, err := c.router.RouteRequest(r)
	if err != nil {
		return nil, err
	}
	return rr, err
}

func generateRequest(marg *MArg, mt dm.MType) (Request, error) {
	id, err := uuid.NewV4()
	if err != nil {
		glog.Errorf("Failed to generate UUID: %v", err)
		return Request{}, err
	}
	r := Request{
		Id:   id,
		Args: marg,
		Key:  marg.Service,
		Type: mt}
	glog.Infof("%s: Generated Request: %v", id, r)
	return r, nil
}

func startRpc(n, laddr string, eChan chan error) {
	api := new(ControllerApi)
	server := rpc.NewServer()
	server.Register(api)
	l, e := net.Listen(n, laddr)
	if e != nil {
		glog.Fatalf("Failed to listen: %v", e)
	}
	glog.Infof("Controller started, listening on: %s", laddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			glog.Errorf("Accept failed: %v", err)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
