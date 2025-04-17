#!/usr/bin/env bash

protoc \
  --go_out=module/global/protogen \
  --go-grpc_out=module/global/protogen \
  -I=module \
  "$(find module/global/proto -iname "*.proto")"

protoc \
  --go_out=module/supply/protogen \
  --go-grpc_out=module/supply/protogen \
  -I=module \
  "$(find module/supply/proto -iname "*.proto")"
