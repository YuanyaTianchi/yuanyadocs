

+++

title = "Title.sub_title"
description = "Title sub_title quick_start"
tags = ["techn", "computer", "tagx", "Title", "_sub_title", "__quick_start(content)"]

+++



# OVS

> [官网](https://www.openvswitch.org/)；[文档](https://docs.openvswitch.org/en/stable/)；



### 安装 ovs

> [Debian / Ubuntu](https://docs.openvswitch.org/en/stable/intro/install/distributions/#debian-ubuntu)；

核心用户空间组件软件包：`openvswitch-switch` & `openvswitch-common`。

```shell
apt install -y openvswitch-switch openvswitch-common

```

进入沙箱

```shell
git clone https://github.com/openvswitch/ovs.git
git ch v2.17.6
tutorial/ovs-sandbox
```



### 启动 faucet

```shell
mkdir inst
docker run -d --name faucet --restart=always -v $(pwd)/inst/:/etc/faucet/ -v $(pwd)/inst/:/var/log/faucet/ -p 6653:6653 -p 9302:9302 faucet/faucet
```

