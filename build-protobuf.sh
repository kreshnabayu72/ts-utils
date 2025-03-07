#!/usr/bin/env bash

protoc \
  --go_out=module/global/protogen \
  --go-grpc_out=module/global/protogen \
  -I=module \
  "$(find module/global/proto -iname "*.proto")"
