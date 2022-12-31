#!/bin/bash

# 输出目录
GO_PUT_PATH='./'

# 定义服务数组
services=(
'./errors/*.proto'
'./paging/*.proto'
'./dictionary/*.proto'
'./model/*.proto'
)

for item in "${services[@]}" ; do
protoc --go_out=paths=source_relative:$GO_PUT_PATH $item
done