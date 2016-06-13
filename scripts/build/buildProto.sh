#!/bin/bash

set -e

GOSRC=$GOPATH/src

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/NEU-SNS/ReverseTraceroute/datamodel/*.proto 

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/NEU-SNS/ReverseTraceroute/controller/pb/*.proto 
    

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/NEU-SNS/ReverseTraceroute/plcontroller/pb/*.proto 
    

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/NEU-SNS/ReverseTraceroute/vpservice/pb/*.proto 
    

protoc  -I$GOSRC --go_out=plugins=grpc:$GOSRC \
	$GOSRC/github.com/NEU-SNS/ReverseTraceroute/atlas/pb/*.proto 
    

protoc -I/usr/local/include \
       -I.  \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
       --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:$GOSRC \
       $GOSRC/github.com/NEU-SNS/ReverseTraceroute/revtr/pb/*.proto 


protoc -I/usr/local/include \
       -I.  \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
       --grpc-gateway_out=logtostderr=true:$GOSRC \
       $GOSRC/github.com/NEU-SNS/ReverseTraceroute/revtr/pb/*.proto 
