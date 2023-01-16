+++

title = "Kubernetes.扩展"
description = "Kubernetes.扩展"
tags = ["it","cloud","Kubernetes原理"]

+++

# Kubernetes.扩展

## kubebuilder

从根据[官方文档](https://book.kubebuilder.io/introduction.html)指导[安装 kubebuilder](https://book.kubebuilder.io/quick-start.html#installation)。

```shell
curl -L -o ~/go/path/bin kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
```

创建脚手架工程

```shell
mkdir myapp-operator
cd myapp-operator/
kubebuilder init --domain my.domain --repo my.domain/guestbook
```



[创建 API](https://book.kubebuilder.io/quick-start.html#create-an-api) 时（20220327）遇到一个[问题](https://github.com/operator-framework/operator-sdk/issues/5466)，将 golang 版本从 1.18 退回到 1.17.8 以解决





```shell

```

在 `RUN go mod download` 添加如下变量，避免网络问题

```dockerfile
ENV GOPROXY="https://goproxy.cn,direct"
RUN go env GOPROXY
```

gcr.io 无法访问，替换为 dockerhub 上搜索到的 [distroless-static:nonroot 的拷贝](https://hub.docker.com/r/katanomi/distroless-static/tags)。

```shell
$ docker search distroless
NAME                                    DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
kubeimages/distroless-static                                                            2                    
katanomi/distroless-static              Copy of gcr.io/distroless/static:nonroot        1                    
......
```

将 `FROM gcr.io/distroless/static:nonroot` 改为如下内容

```dockerfile
FROM katanomi/distroless-static:nonroot
```

然后再执行

```shell
make docker-build docker-push IMG=<some-registry>/<project-name>:tag
```



