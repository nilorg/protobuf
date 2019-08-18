#!/bin/bash

# 输出目录
GO_PUT_PATH='./'

# 定义服务数组
services=(
'./job/*.proto'
)

for item in "${services[@]}" ; do
echo ${item}
protoc \
--proto_path=$GOPATH/src:. \
--micro_out=paths=source_relative,plugins=grpc:${GO_PUT_PATH} \
--go_out=paths=source_relative,plugins=grpc:${GO_PUT_PATH} \
${item}
done