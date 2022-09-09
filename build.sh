#!/bin/bash
# 编译 go 项目
base_dir=$(cd "$(dirname "$0")";pwd)
courseware_dir=$base_dir/courseware
user_dir=$base_dir/user
cd $user_dir
echo pwd is $(pwd)
echo "########## start build user  ##########"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api/apiuser api/user.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rpc/rpcuser rpc/user.go

cd $courseware_dir
echo pwd is $(pwd)
echo "########## start build courseware ##########"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api/apicourseware api/courseware.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rpc/rpccourseware rpc/courseware.go
