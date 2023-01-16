

+++

title = "Linux.工具"
description = "Linux.工具"
tags = ["it","os","Linux"]

+++

# Linux.工具

## ssh

### 密钥生成

```shell
# -t 指定算法；-C 添加注释，一般用邮箱表明身份信息
ssh-keygen -t ed25519 -C "yuanya@tianchi.com"
# 如果使用的是不支持 Ed25519 算法的旧系统，则使用 RSA 即可
ssh-keygen -t rsa -b 4096 -C "yuanya@tianchi.com"
```

### 登录认证

https://blog.csdn.net/pipisorry/article/details/52269785

登录：-p可以指定端口，默认22端口

```shell
ssh root@192.168.100.100
# 之后输入密码即可
```

密钥对，将生成到 ~/.ssh 下

```shell
# 使用ssh-keygen命令生成密钥对。
ssh-keygen -t rsa 
# ssh-copy-id会将公钥写到远程主机的 ~/.ssh/authorized_key 文件中
ssh-copy-id root@192.168.100.100
# 登录就无需密码了
ssh root@192.168.100.100
```

```shell
# ed25519算法
ssh-keygen -t ed25519 -C "yuanyatianchi@gmail.com"
# rsa算法
ssh-keygen -t rsa -b 4096 -C "yuanyatianchi@gmail.com"
```



## scp

### 文件传输

linux 下一般使用 scp 命令来通过 ssh 传输文件

```shell
# 下载远程文件到本地
scp root@192.168.100.100:/root/remotefile /root/localfile
# 上传本地文件到远程
scp /root/localfile root@192.168.100.100:/root/remotefile

# 下载远程目录到本地
scp -r root@192.168.100.100:/root/remotedir /root/localdir
# 上传本地目录到远程
scp -r /root/localdir root@192.168.100.100:/root/remotedir
```





## bat

> [bat](https://github.com/sharkdp/bat)：带有语法高亮的 cat 命令。

centos8stream 内核版本4.18 glibc版本2.28，因此选择安装 bat v0.19.0，bat v0.20.0 及以上需要更高 glibc 版本，详细参考[官方版本映射](https://github.com/sharkdp/bat#installation)。

```shell
# centos8stream 内核版本4.18 glibc版本2.28，bat 20及以上版本需要 glibc版本2.29
wget https://github.com/sharkdp/bat/releases/download/v0.19.0/bat-v0.19.0-x86_64-unknown-linux-gnu.tar.gz
tar -zxf bat-v0.19.0-x86_64-unknown-linux-gnu.tar.gz
mv bat-v0.19.0-x86_64-unknown-linux-gnu/bat /usr/local/bin/
rm -rf bat-v0.19.0-x86_64-unknown-linux-gnu*
```

bat 默认[自动分页](https://github.com/sharkdp/bat#automatic-paging)，通过`--paging=never`关闭。可直接设置别名到`.bashrc`

```shell
cat >> ~/.bashrc << EOF
# tools
alias bat ='bat --paging=never'
EOF
```

