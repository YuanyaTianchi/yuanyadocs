

+++

title = "Title.sub_title"
description = "Title sub_title quick_start"
tags = ["techn", "computer", "tagx", "Title", "_sub_title", "__quick_start(content)"]

+++



# Sealos

> [官网](https://docs.sealos.io/zh-Hans/)；[文档](https://docs.sealos.io/zh-Hans/docs/Intro)；



# 快开



## 安装

> [releases](https://github.com/labring/sealos/releases)；

### 二进制

```shell
wget https://github.com/labring/sealos/releases/download/v4.1.7/sealos_4.1.7_linux_amd64.tar.gz

tar zxvf sealos_4.1.7_linux_amd64.tar.gz sealos && mv sealos /usr/local/bin
```



## 一键启动集群

> [集群安装 Kuberentes](https://docs.sealos.io/zh-Hans/docs/getting-started/kuberentes-life-cycle#%E9%9B%86%E7%BE%A4%E5%AE%89%E8%A3%85-kuberentes)；

```shell
# sealos 支持使用 containerd，使用前关闭docker 或者直接删除
$ systemctl disable --now docker
$ apt autoremove -y docker-ce

$ sealos run \
    labring/kubernetes:v1.25.0 \
    labring/helm:v3.8.2 \
    labring/calico:v3.24.1 \
    --masters 192.168.11.21,192.168.11.22,192.168.11.23 \
    --nodes 192.168.11.31,192.168.11.32 -p [your-ssh-passwd]
```



# 常用

### 使用 sealos pull 从本地 dockerd 拉取镜像

```shell
$ sealos pull -h
...
Examples:
  sealos pull imagename
  sealos pull docker-daemon:imagename:imagetag
  sealos pull myregistry/myrepository/imagename:imagetag
...

$ sealos pull docker-daemon:labring/kubernetes:v1.26
```

