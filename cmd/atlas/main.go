package main

import (
	"flag"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/NEU-SNS/ReverseTraceroute/atlas"
	"github.com/NEU-SNS/ReverseTraceroute/atlas/pb"
	"github.com/NEU-SNS/ReverseTraceroute/config"
	"github.com/NEU-SNS/ReverseTraceroute/dataaccess"
	"github.com/NEU-SNS/ReverseTraceroute/log"
)

// Config is the config for the atlas
type Config struct {
	Db       dataaccess.DbConfig
	KeyFile  string
	CertFile string
	RootCA   string
}

func init() {
	config.SetEnvPrefix("ATLAS")
	config.AddConfigPath("./atlas.config")
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		host, _, err := net.SplitHostPort(req.RemoteAddr)
		switch {
		case err != nil:
			return false, false
		case host == "localhost" || host == "127.0.0.1" || host == "::1" || host == "rhansen2.local" || host == "rhansen2.revtr.ccs.neu.edu" || host == "129.10.113.189":
			return true, true
		default:
			return false, false
		}
	}
	grpclog.SetLogger(log.GetLogger())
}

func main() {
	go sigHandle()
	conf := Config{}
	err := config.Parse(flag.CommandLine, &conf)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	da, err := dataaccess.New(conf.Db)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	cres, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	svc := atlas.NewAtlasService(da, conf.RootCA)
	serv := grpc.NewServer(grpc.Creds(cres))
	pb.RegisterAtlasServer(serv, atlas.GRPCServ{AtlasService: svc})
	ln, err := net.Listen("tcp", "0.0.0.0:55000")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer func() {
		err := ln.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	err = serv.Serve(ln)
	if err != nil {
		log.Error(err)
	}
}

func sigHandle() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGSTOP)
	for _ = range c {
		os.Exit(1)
	}
}
