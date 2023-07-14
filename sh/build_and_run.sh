#!/bin/bash

# 第一个参数是服务的名称
service_name=$1

sh sh/clean.sh
cd ${service_name}
go build
./${service_name}
