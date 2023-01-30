

+++

title = "Linux"
description = "Linux"
tags = ["it","os","Linux"]

+++



# Linux.快速开始

- Linux
  - Linus Benedict Torvalds（Linux之父）编写的开源操作系统的内核
  - 广义上的基于 Linux内核 的 Linux操作系统
- 内核版本：https://www.kernel.org/
  - 内核版本分为三个部分：主版本号、次版本号、末版本号
  - 次版本号是奇数为开发版，偶数为稳定版。但是实际上在2.6以后不这么区分了
- 发行版本：RedHat Enterprise Linux、Fedora、CentOS、Debian、Ubuntu...
- 终端：图形终端，命令行终端，远程终端（SSH、VNC）





## VirtualBox

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

8. 虚拟机与宿主机 共享粘贴板、拖拽文件

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

## ubuntu

> [ubuntu阿里云镜像](https://mirrors.aliyun.com/ubuntu-releases/)；



### root密码重设

```shell
sudo passwd

# 登录root
su root
```

### root图形界面登录

```shell
# ubuntu桌面版不允许图形界面登录root用户，须注释`/etc/pam.d/gdm-password`&`/etc/pam.d/gdm-autologin`文件中 auth required pam_succeed_if.so user != root quiet_success 所在行
sed -i 's/auth.*required.*pam_succeed_if.so user != root quiet_success/# auth\trequired\tpam_succeed_if.so user != root quiet_success/g' /etc/pam.d/gdm-password
sed -i 's/auth.*required.*pam_succeed_if.so user != root quiet_success/# auth\trequired\tpam_succeed_if.so user != root quiet_success/g' /etc/pam.d/gdm-autologin
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

### ssh服务启用

```shell
# 安装
apt install -y openssh-server
# 默认禁止远程登录，需 /etc/ssh/sshd_config 中 PermitRootLogin 字段
sed -i '/#PermitRootLogin prohibit-password/aPermitRootLogin yes' /etc/ssh/sshd_config
# 重启
systemctl restart sshd
# 查看状态
systemctl status sshd
```

### 网卡配置

系统会读取`/etc/netplan/*.yaml`

```shell
cat >> /etc/netplan/02.yaml << EOF
network:
  version: 2
  renderer: NetworkManager
  ethernets:
    enp0s3:
      dhcp4: false
      # gateway4: 192.168.31.1
      routes:
      - to: default
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

### https访问

```shell
apt install ca-certificates -y
```



### apt源

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





## centos

> [centos阿里云镜像](https://mirrors.aliyun.com/ubuntu-releases/)；



## 目录

- 目录
  - /：根目录
  - /root：root用户的家目录
  - /home/username：普通用户的家目录
  - /etc：配置文件目录
  - /bin：命令目录
  - /sbin：管理命令目录
  - /usr/bin、/usr/sbin：系统预装的其他命令



## 命令行

- 命令（command）：通过使用命令调用对应的命令程序文件，如使用`ls`命令将调用`/bin/ls`程序文件
- 参数（option选项）：

### 语法

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

### 简写

- 多命令：命令之间`;`分隔，顺序执行

```sh
cd / ; ls
```

- 多参数简写

```shell
ls -l -r -t -R
ls -lrtR #简写
```

- 当前目录简写

```shell
ls ./
ls .   #简写。具体到文件（或目录）时无法使用"."作简写
ls     #简写

cd ../
cd ..  #简写

cat ./config.yml
cat config.yml #简写
```

### 用户

```shell
clear     #清屏。或者快捷键CTRL+L
su - root #切换到root用户
exit      #退出当前系统用户
init 0    #关机
```

### 帮助命令

- 为什么要学习帮助命令：Linux的基本操作方式是命令行，海量的命令不适合“死记硬背”，你要升级你的大脑
- 使用网络资源（搜索引擎和官方文档)

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





### 环境变量

- PATH：指定命令的搜索路径
  - PATH=$PAHT:<PATH 1>:<PATH 2>:--------:< PATH n >
  - export PATH

临时：仅对当前 shell(bash) 或其子 shell(bash) 生效

```sh
export PATH=$PATH:/usr/local/go/bin
```

用户

```sh
# 编辑文件 ~/.bash_profile
vim ~/.bash_profile
# 写入环境变量
export PATH=$PATH:/usr/local/go/bin
# 刷新环境变量使生效
source ~/.bash_profile
```

全局

```sh
# 编辑文件 /etc/profile
vim /etc/profile
# 写入环境变量
export PATH=$PATH:/usr/local/go/bin
# 刷新环境变量使生效
source /etc/profile
```





## grub配置

- grub：centos7使用的是grub2，此前是grub1，grub1中所有的配置文件需要手动去编辑，要像网卡配置文件一样记住每个项是什么功能，grub2则可以用命令即可进行修改

- grub配置文件：/boot/grub2/grub.cfg，一般不要直接编辑该文件，而是修改/etc/default/grub文件（一些基本配置），如果还有更详细的配置，可以修改/etc/grub.d/下的文件
  - grub2-mkconfig -o /boot/grub2/grub.cfg：产生新的grub配置文件
  - /etc/default/grub
    
    - GRUB_DEFAULT=saved。表示系统默认引导的版本内核。通过命令grub2-editenv list查看默认引导的版本内核
    
      - grep ^menu /boot/grub2/grub.cfg：找到文本文件当中包含关键字的一行，^表示以什么开头，这里即在grub配置文件去找到以menu开头的，能够看到内核的列表，按索引以0开始
      - grub2-set-default 0：选择第一个内核作为linux启动时的默认引导
      - grub2-editenv list：查看当前默认引导内核已经改变了
    
    - GRUB_CMDLINE_LINUX：确认引导时对linux内核增加什么参数
    
      - rhgb：引导时为图形界面
    
      - quiet：静默模式，引导时只打印必要的消息，启动出现异常时可以去除quiet和rhgb以显示更多信息
    
      - readhat7重置root密码
    
        - 引导界面时，选择要引导的内核，按E进入设置信息，找到linux16开头的一段，可以发现刚刚quiet、rhgb等信息，可以直接键盘输入添加更多项，在该行末尾添加rd.break，ctrl+x启动
    
        - 进入后是内存中的虚拟的一个文件系统，而真正的根目录是/sysroot（输入命令mount可以发现根为/sysroot），在这里所做的操作是不会进行保存的，并且是只读方式的挂载，不能写，防止修复时损坏原有文件
    
        - `mount -o remount,rw /sysroot`，重新挂载到根目录并且是要可读写的，之后mount，发现有了r,w权限
    
        - chroot /sysroot 选择根，即设置根为/sysroot目录；
    
          - echo 123456|passwd --stdin root 修改root密码为redhat，echo正常情况是打印到终端，这里通过管道符发送给password命令，--stdin是password命令的参数，正常情况是通过终端输入，这里即表示通过标准输入进行输入，并传递给root用户。或者password 123456；或者输入passwd，交互修改；
    
          - SELinux安全组件，叫做强制访问控制，会对etc/password和/etc/shadow进行校验，如果这两个文件不是在系统进行标准修改的，会导致无法进入系统。`vim /etc/selinux/config`中可以通过设置SELINUX=disabled关闭SELinux，即使生产环境也多半会关掉它
          - 或者修改/etc/shadow文件，touch /.autorelabel，这句是为了selinux生效
          - 注意备份
    
        - exit退出根回到虚拟的root中，然后reboot



## 快捷键

- TAB快捷键补全文件（目录）名
- CTRL+L清屏
- CTRL+C终止（命令）程序



## 别名



## sysctl

可以查看和修改系统参数



## grubby

可以修改内核参数



## 图形界面黑屏

centos7上，使用root用户登录，会出现图形界面黑屏问题，在输入用户名密码之前是有图形界面的，但是输入用户名密码之后桌面出现一秒钟之后转为黑屏

解决问题：通过 ctrl+alt+F2 切换到命令行界面，或者通过其它终端连接该机器，通过 `startx` 命令重启图形界面，可以发现缺少（损失）了某些文件，因为系统启动时会检查用户家目录下所有文件，因缺少相应文件导致图形界面不可用，通过 `startx` 重启图形界面过程中系统会自动重新创建 /root/.Xauthority 文件，图形界面就可以正常使用了。但是重启发现还是会黑屏，仍然需要通过 `startx` 重开图形界面，但是仅提示 `/root/.serverauth.xxx` 文件不存在，

```sh
$ startx
xauth:  file /root/.serverauth.3218 does not exist
xauth:  file /root/.Xauthority does not exist
xauth:  file /root/.Xauthority does not exist


X.Org X Server 1.20.4
X Protocol Version 11, Revision 0
Build Operating System:  3.10.0-957.1.3.el7.x86_64 
Current Operating System: Linux MiWiFi-R3A-srv 3.10.0-1160.24.1.el7.x86_64 #1 SMP Thu Apr 8 19:51:47 UTC 2021 x86_64
Kernel command line: BOOT_IMAGE=/vmlinuz-3.10.0-1160.24.1.el7.x86_64 root=/dev/mapper/centos-root ro crashkernel=auto rd.lvm.lv=centos/root rd.lvm.lv=centos/swap rhgb quiet LANG=zh_CN.UTF-8 user_namespace.enable=1
Build Date: 24 February 2021  09:09:20PM
Build ID: xorg-x11-server 1.20.4-15.el7_9 
Current version of pixman: 0.34.0
        Before reporting problems, check http://wiki.x.org
        to make sure that you have the latest version.
Markers: (--) probed, (**) from config file, (==) default setting,
        (++) from command line, (!!) notice, (II) informational,
        (WW) warning, (EE) error, (NI) not implemented, (??) unknown.
(==) Log file: "/var/log/Xorg.1.log", Time: Wed Apr 21 00:16:03 2021
(==) Using config directory: "/etc/X11/xorg.conf.d"
(==) Using system config directory "/usr/share/X11/xorg.conf.d"
VMware: No 3D enabled (0, Success).
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





## 关闭swap分区

```sh
# 删除 swap 区所有内容，即临时关闭
swapoff -a
# 删除 swap 挂载，注释 swap 相关行即可，这样系统下次启动不会再挂载 swap
vim /etc/fstab
# 重启
reboot
# 查看。swap 一行应该全部是 0
free -h
```



## 修改静态主机名hostname

```sh
vim /etc/hostname
# 或者
hostnamectl set-hostname xxx
# 查看
hostname
hostnamectl
```



## ubuntu查看内核

```shell
# x86_64,x64,AMD64基本上是同一个东西
arch
```

## ubuntu18.04配置apt源

`/etc/apt/sources.list`

```
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
```



****

ubuntu20.04配置apt源
