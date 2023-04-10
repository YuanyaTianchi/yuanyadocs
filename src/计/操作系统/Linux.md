



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



# 快开



## VirtualBox

> [Download VirtualBox](https://www.virtualbox.org/wiki/Downloads)；

VirtualBox7.0 默认兼容 hyper-v，可以与 wsl 共存；

安装 platform packages 即可，Extension Pack 根据需要考虑

### 安装 VM

1. 新建

2. 名称：demovm → 文件夹：D:\it\vm\vms → 类型：Linux → 版本：Ubuntu/Red Hat...

3. 内存大小：2048mb

4. 现在创建虚拟硬盘 → VDI (VirtualBox 磁盘映像) → 动态分配 → 硬盘大小：32G

5. 设置 → 存储 → 控制器: IDE → 右侧光盘图标选择虚拟盘 → OK

6. 设置 → 网络 → 选择桥接网卡

7. 启动：这里根据发行商版本有所不同

   1. Install xxx → English

   2. DATE & TIME → Asia Shanghai

   3. SOFTWARE SELECTION → Minimal Install （按需选择）

   4. INSTALL ATION DESTINATION → 直接点done即可

   5. NETWORK & HOST NAME → Ethernet (enp0s3)：ON

   6. 开始安装 → 设置 ROOT PASSWORD → 安装成功后 reboot

### 启用虚拟机与宿主机共享 粘贴板/卷

进入 VM 设置

1. 设置 - 常规 - 高级 - 共享粘贴板 双向 & 拖放 双向

2. 设置 - 存储 - 控制器：SATA - √ 使用主机输入输出（I/O）缓存

3. 设置 - 存储 - 控制器：SATA - .vdi - √ 固态驱动器(s)

4. ubuntu：只需要安装gcc make perl

   ```shell
   apt install -y gcc make perl
   ```

5. centos：

   kernel-devel

   ```sh
   # 安装 kernel-devel 和 gcc
   yum install -y kernel-devel gcc
   # 更新 kernel 和 kernel-devel 到最新版本
   yum upgrade kernel kernel-devel -y
   # 重启
   reboot
   ```

   虚拟机窗口上方菜单栏 - 设备 - 安装增强功能，会挂载一个盘，如果程序没有自动执行则到挂载的盘中运行`aoturun.sh`（重新安装要先弹出iso，否则报"未能加载虚拟光盘"）

   ```shell
   # 必要依赖。安装时提示 please install the gcc make perl packages from your modules；安装过后日志里又提示 please install libelf-dev, libelf-devel or elfutils-libelf-devel 
   yum install gcc make perl elfutils-libelf-devel -y
   # 如果提示 ValueError: File context for /opt/VBoxGuestAdditions-6.1.4/other/mount.vboxsf already defined
   semanage fcontext -d /opt/VBoxGuestAdditions-6.1.4/other/mount.vboxsf
   restorecon /opt/VBoxGuestAdditions-6.1.4/other/mount.vboxsf
   # 若是以上相关软件或内核更新后，增强功能还无法正常使用，使用一次性更新所有软件（这次可以了）
   yum update
   # 重启
   reboot
   ```

   stalstalstal



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
      - to: default # default 必需
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
$ apt update && apt upgrade -y
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

Ubuntu 桌面版不允许图形界面登录 root 用户，须注释 `/etc/pam.d/gdm-password` & `/etc/pam.d/gdm-autologin` 文件中 `auth required pam_succeed_if.so user != root quiet_success` 所在行

```shell
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

### 其它软件

```shell
apt install -y vim git sshpass
```



## CentOS

> [centos阿里云镜像](https://mirrors.aliyun.com/ubuntu-releases/)；



# 常用



## 别名

后台启动且不做输出（输出到 `/dev/null`）

```shell
alias goland="nohup ~/techn/computer/lang/go/goland/bin/goland.sh >/dev/null & 2>&1"'
```

## ssh 免密

```shell
sshpass -p <password> ssh -o StrictHostKeychecking=no -o ServerAliveInterval=60 root@192.168.1.11
```



## netplan 网桥配置

> [netplan examples](https://netplan.io/examples)；
>
> [configuring-network-bridges](https://netplan.io/examples#configuring-network-bridges)；

```shell
$ cat > /etc/netplan/br.yaml << EOF
network:
  version: 2
  renderer: networkd
  ethernets:
    eth0:
      dhcp4: no
  bridges:
    br0:
      dhcp4: yes
      interfaces:
        - eth0
EOF

$ netplan apply
```

## sh -c

这个命令将权限不够，因为重定向符号 “>” 和 ">>" 也是 bash 的命令。我们使用 sudo 只是让 echo 命令具有了 root 权限

```sh
sudo echo "hahah" >> test.csv`
```

 `sh -c` 命令，它可以让 bash 将一个字串作为完整的命令来执行，这样就可以将 sudo 的影响范围扩展到整条命令

```sh
sudo sh -c echo "hahah" >> test.csv`
```



## hostname

```sh
vim /etc/hostname
# 或者
hostnamectl set-hostname xxx
# 查看
hostname
hostnamectl
```



## sysctl

可以查看和修改系统参数



## grubby

可以修改内核参数

## 快捷键

- TAB快捷键补全文件（目录）名
- CTRL+L清屏
- CTRL+C终止（命令）程序

### 帮助命令

###### help

```shell
help cd   #查看内部命令帮助
type cd   #查看命令类型。shell（命令解释器）自带的命令称为内部命令，其他的是外部命令
ls --help #查看外部命令帮助
```

###### info

```shell
info ls #查看ls命令的信息。info比help更详细，作为help的补充
```

###### man

- `man [(1)|2|3|4|5|6|7|8|9] [文件名]`：查看指定文件（程序）对应的手册（如果有的话）。这里1-9指定要查看手册的类型，因为系统中很多重名文件，所以需要分类来区分
- 按`Q`退出手册

```shell
man man   #查看man命令的手册。可以看到默认的选项"1"，表示查看 可执行程序或shell命令 这一类型的文件的手册

man ls    #查看ls命令的手册。选项"1"是默认选项，可以省略
man 1 ls
man -a ls #查看所有名为ls的文件的手册
```



## 命令行语法

- 命令 <必选参数1 | 必选参数2> [-option {必选参数1 | 必选参数2 | 必选参数3}] [可选参数...] {(默认参数) | 参数 | 参数}
- 命令行语法符号
  - 方括号**[ ]**：可选参数，在命令中根据需要加以取舍
  - 尖括号**< >**：必选参数，实际使用时应将其替换为所需要的参数
  - 大括号**{ }**：必选参数, 内部使用, 包含此处允许使用的参数
  - 小括号**()**：指明参数的默认值, 只用于{ }中
  - 管道符（竖线）**|**：分隔多个互斥参数, 含义为"或", 使用时只能选择一个
  - 省略号**...**：多个参数
  - 分号**;**：分割多个命令，命令将按顺序执行
- 注意：命令行语法（包括在 UNIX 和 Linux 平台中使用的用户名、密码和文件名）是区分大小写的，如commandline、CommandLine、COMMANDLINE 是不一样的
