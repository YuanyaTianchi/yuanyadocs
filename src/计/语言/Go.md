

+++

title = "Go"
description = "Go quick start"
tags = ["techn", "computer", "lang", "Go", "__quick_start"]

+++



# Go

> [官网](https://go.dev/)；



## go

> [下载](https://go.dev/dl/)：选择 `.linux-amd64.tar.gz`，Win 选择 .windows-amd64.zip；

安装：解压即可

```shell
tar -zxf gox.x.x.linux-amd64.tar.gz
```

环境变量：导入到系统环境变量，或者通过命令如 `go env -w GOROOT=` 写入

```shell
export GOROOT=~/compute/lang/go/go
export GOPATH=~/compute/lang/go/path
# go install 安装的程序将保存到 $GOPATH/bin 下
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
# 代理：七牛云代理 goproxy.cn,direct；阿里代理 mirrors.aliyun.com/goproxy/
export GOPROXY=https://goproxy.cn,direct

# gomod：go 1.14 后默认开启
export GO111MODULE=on
# 私有库：设置私有库以跳过代理，比如公司使用 Gitlab、Gitee 等，多私有库逗号分隔
export GOPRIVATE=*.gitlab.com,*.gitee.com
# 用于验证包的有效性的服务地址，默认 sum.golang.org。有时候如 go mod vendor 会用到，无法连接则设置可连接的地址如 sum.golang.google.cn，这是专为国内提供的 sum 验证服务地址，或者直接关闭即可。私有仓库自动忽略验证
export GOSUMDB=off

# 如写入 ~/.bashrc 则刷新
source ~/.bashrc
```



## vs code

> [为 Go 开发配置 Visual Studio Code](https://learn.microsoft.com/zh-cn/azure/developer/go/configure-visual-studio-code)；

1. `Ctrl+Shift+X` 打开扩展视图，搜索 `Go` 插件并安装
2. `Ctrl+Shift+P` 打开 vs code 命令框，搜索 `Go: Install/Update tools` 并执行



## hello world

helloworld.go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

```shell
$ go build -o ./helloworld-exec ./helloworld.go
$ ./helloworld-exec
Hello world!
```

