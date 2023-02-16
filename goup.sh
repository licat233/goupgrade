#!/bin/bash
###
 # @Author: licat
 # @Date: 2023-02-16 16:42:16
 # @LastEditors: licat
 # @LastEditTime: 2023-02-16 16:42:18
 # @Description: licat233@gmail.com
###

# 用于升级Go语言的脚本

# 定义当前Go版本和要升级到的Go版本
CURRENT_GO_VERSION=$(go version | awk '{print $3}')
UPGRADE_GO_VERSION=$1 # 如："1.17.5"

# 检查当前用户是否为root用户
if [ "$EUID" -ne 0 ]
  then echo "请使用root用户运行此脚本"
  exit
fi

# 下载Go语言安装包并解压缩
cd /tmp
curl -O https://golang.org/dl/go${UPGRADE_GO_VERSION}.linux-amd64.tar.gz
tar -C /usr/local -xzf go${UPGRADE_GO_VERSION}.linux-amd64.tar.gz

# 设置环境变量
echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
echo "export GOPATH=\$HOME/go" >> /etc/profile
source /etc/profile

# 检查Go版本是否升级成功
if [ "$CURRENT_GO_VERSION" == "$UPGRADE_GO_VERSION" ]; then
    echo "Go语言升级成功"
else
    echo "Go语言升级失败"
fi
