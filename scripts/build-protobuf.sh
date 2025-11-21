#!/usr/bin/env bash

global_proto_path="$(find ./module/global/proto -iname "*.proto")"

for proto in $global_proto_path ; do
    protoc \
      --go_out=./module/global/protogen \
      --go-grpc_out=./module/global/protogen \
      -I=. \
      $proto
done


supply_proto_path="$(find ./module/supply/proto -iname "*.proto")"

for proto in $supply_proto_path ; do
    protoc \
      --go_out=./module/supply/protogen \
      --go-grpc_out=./module/supply/protogen \
      -I=. \
      $proto
done
