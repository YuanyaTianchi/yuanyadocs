



+++

title = "Linux"
description = "Linux quick start"
tags = ["techn", "computer", "os", "Linux", "__quick_start"]

+++



# Linux

> [官网](https://www.kernel.org/)；
>
> [仓库](https://git.kernel.org/)；
>
> [Linus Torvalds](https://en.wikipedia.org/wiki/Linus_Torvalds)；



## Ubuntu

> [ubuntu 阿里云镜像](https://mirrors.aliyun.com/ubuntu-releases/)；



### root 密码设置

```shell
sudo passwd

# 登录root
su root
```

### vi

```shell
# vi 默认情况下上下左右、backspace等按键非正常作用，须配置`/etc/vim/vimrc.tiny`
cat >> /etc/vim/vimrc.tiny << EOF

set nocompatible
set backspace=2
EOF
```

### 快速关机

```shell
# 快速关机
sed -i '/#DefaultTimeoutStartSec=90s/aDefaultTimeoutStartSec=5s' /etc/systemd/system.conf
sed -i '/#DefaultTimeoutStopSec=90s/aDefaultTimeoutStopSec=5s' /etc/systemd/system.conf
```

### 网卡配置

```shell
# 系统会读取 `/etc/netplan/*.yaml`
cat >> /etc/netplan/02.yaml << EOF
network:
  version: 2
  renderer: NetworkManager
  ethernets:
    enp0s0:
      dhcp4: false
      # gateway4: 192.168.1.1
      routes:
      - to: default
        via: 192.168.1.1  # gateway, same as gateway4
      addresses:
      - 192.168.1.10/24
      nameservers:
        addresses: [192.168.1.1,8.8.8.8] # DNS
        search: []
    enp0s3:
      dhcp4: false
      # gateway4: 192.168.31.1
      routes:
      - to: 192.168.31.0/24
        via: 192.168.31.1  # gateway, same as gateway4
      addresses:
      - 192.168.31.10/24
      nameservers:
        addresses: [192.168.31.1,8.8.8.8] # DNS
        search: []
EOF

# 应用配置
netplan apply
```

### apt 源

```shell
$ mv /etc/apt/sources.list /etc/apt/sources.list.bk
$ cat > /etc/apt/sources.list << EOF
deb http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-proposed main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse
EOF
# 更新源
$ apt update
```

### ssh 服务启用

```shell
# 安装
apt install -y openssh-server
# 默认禁止远程登录，需 /etc/ssh/sshd_config 中 PermitRootLogin 字段
sed -i '/#PermitRootLogin prohibit-password/aPermitRootLogin yes' /etc/ssh/sshd_config
# 重启
systemctl restart sshd
```

### http 代理

### https 访问

```shell
apt install -y ca-certificates curl gnupg
```

### root 图形界面登录

```shell
# ubuntu桌面版不允许图形界面登录root用户，须注释`/etc/pam.d/gdm-password`&`/etc/pam.d/gdm-autologin`文件中 auth required pam_succeed_if.so user != root quiet_success 所在行
sed -i 's/auth.*required.*pam_succeed_if.so user != root quiet_success/# auth\trequired\tpam_succeed_if.so user != root quiet_success/g' /etc/pam.d/gdm-password
sed -i 's/auth.*required.*pam_succeed_if.so user != root quiet_success/# auth\trequired\tpam_succeed_if.so user != root quiet_success/g' /etc/pam.d/gdm-autologin
```

### 用户文件夹

指定Desktop、Downloads等

```shell
sed -i 's/XDG_DESKTOP_DIR=.*/XDG_DESKTOP_DIR="$HOME\/dirs\/Desktop"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_DOWNLOAD_DIR=.*/XDG_DOWNLOAD_DIR="$HOME\/dirs\/Downloads"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_TEMPLATES_DIR=.*/XDG_TEMPLATES_DIR="$HOME\/dirs\/Templates"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_PUBLICSHARE_DIR=.*/XDG_PUBLICSHARE_DIR="$HOME\/dirs\/Public"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_DOCUMENTS_DIR=.*/XDG_DOCUMENTS_DIR="$HOME\/dirs\/Documents"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_MUSIC_DIR=.*/XDG_MUSIC_DIR="$HOME\/dirs\/Music"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_PICTURES_DIR=.*/XDG_PICTURES_DIR="$HOME\/dirs\/Pictures"/g' ~/.config/user-dirs.dirs
sed -i 's/XDG_VIDEOS_DIR=.*/XDG_VIDEOS_DIR="$HOME\/dirs\/Videos"/g' ~/.config/user-dirs.dirs
# 现在使用
nautilus -q
```



## CentOS

> [centos阿里云镜像](https://mirrors.aliyun.com/ubuntu-releases/)；



