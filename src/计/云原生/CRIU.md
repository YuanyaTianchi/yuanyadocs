

# CRIU

> [官网](https://criu.org)；

## ---快开---

> [安装](https://criu.org/Installation)；



[安装构建依赖项](https://criu.org/Installation#Installing_build_dependencies)：deb 包为例，适用于 debian、ubuntu

1. [编译器和C库](https://criu.org/Installation#Compiler_and_C_Library)：`gcc make`；
2. [协议缓冲区](https://criu.org/Installation#Protocol_Buffers)：`libprotobuf-dev libprotobuf-c-dev protobuf-c-compiler protobuf-compiler python3-protobuf`；
3. [其他东西](https://criu.org/Installation#Other_stuff)：`pkg-config libbsd-dev iproute2 libnftables-dev libcap-dev libnet1-dev libnl-3-dev libnet-dev libgnutls28-dev python3-future`；

```shell
apt install -y \
  gcc make \
  libprotobuf-dev libprotobuf-c-dev protobuf-c-compiler protobuf-compiler python3-protobuf \
  pkg-config \
  libbsd-dev \
  iproute2 \
  libnftables-dev \
  libcap-dev \
  libnet1-dev libnl-3-dev libnet-dev \
  libgnutls28-dev \
  python3-future
```

[构建](https://criu.org/Installation#Installing)或在[容器中构建](https://criu.org/Installation#Building_the_tool)：

```shell
make
```

可以直接使用可执行文件`./criu/criu`运行 criu，或继续执行`make install`使安装到标准路径



## ---基操---