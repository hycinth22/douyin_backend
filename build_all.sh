#!/bin/bash

# 定义常量
services=("user" "feed" "relation")
cwd=$(pwd)
output_dir="$cwd/output/"

mkdir -p $output_dir

# 遍历服务
for service in "${services[@]}"
do
  echo "Building service: $service"

  go build -o $output_dir/$service "./service/$service/"
done

# 构建web
echo "Building hertz web server: web"
go build -o $output_dir/$web ./web

ls $output_dir
echo "All build files has been written into $output_dir"
