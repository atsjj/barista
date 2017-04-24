#!/bin/sh

protoc -I /go/src/app/sap -I /go/src/github.com/google/protobuf/src --go_out=plugins=grpc:/go/src/app/sap /go/src/app/sap/*.proto
