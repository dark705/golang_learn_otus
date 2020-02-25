#!/bin/bash
protoc  --proto_path=../api ../api/protobuf.proto --go_out=plugins=grpc:../internal/protobuf