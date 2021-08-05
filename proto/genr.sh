#!/bin/bash

PROJECT_PATH=`PWD`/../../../../

protoc --go_out=plugins=grpc:$PROJECT_PATH  ./common.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./coregw.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./configmap.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./api.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./customer.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./sn.proto
# protoc --go_out=plugins=grpc:$PROJECT_PATH  ./group-admin.proto

protoc --go_out=plugins=grpc:$PROJECT_PATH  ./charger/charger.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./charger/measure.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./charger/warning.proto

# protoc --go_out=plugins=grpc:$PROJECT_PATH  ./verify-code.proto
protoc --go_out=plugins=grpc:$PROJECT_PATH  ./open.proto

protoc --go_out=plugins=grpc:$PROJECT_PATH  ./pay.proto
#protoc --go_out=plugins=grpc:$PROJECT_PATH  ./monitor.proto
# protoc --go_out=plugins=grpc:$PROJECT_PATH  ./xa.proto

