# WSL2

当前版本[linux-msft-wsl-5.15.90.1](https://github.com/microsoft/WSL2-Linux-Kernel/archive/refs/tags/linux-msft-wsl-5.15.90.1.tar.gz)。

## 准备

以管理员身份运行 Windows PowerShell

```powershell
# 启用：Windows功能-适用于Linux的Windows子系统
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart

# 启用：Windows功能-虚拟机平台
# 或启用：Windows功能-Hyper-V，其包含 Windows功能-虚拟机平台，相比之下将获得额外的管理功能等
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart

# 重启 Windows
```

## 安装

Win11 [wsl 安装](https://learn.microsoft.com/zh-cn/windows/wsl/install)默认版本即为 wsl2

```powershell
# 查看可用发行版
wsl -l -o

# 命令行安装：选择 Ubuntu 22.04 LTS（包含 gui 功能）
wsl --install -d Ubuntu-22.04

# 上面我域名解析失败所以我通过微软商店安装
# 微软商店安装：选择 Ubuntu 22.04 LTS，功能一样

# 更新 wsl，否则报错 0x800701bc
wsl --update
```

然后可以打开 ubuntu，在 windows terminal 中也可以选择



## vscode插件

搜索 WSL，由微软官方提供的远程连接插件



## KVM

用于云原生学习

https://www.bilibili.com/video/av201187993

https://boxofcables.dev/accelerated-kvm-guests-on-wsl-2/

https://boxofcables.dev/kvm-optimized-custom-kernel-wsl2-2022/

默认安装的 wsl2 不含 kvm 模块，需要重新[下载内核](https://github.com/microsoft/WSL2-Linux-Kernel)编译。

更新、升级、安装依赖、编译内核

```shell
sudo apt update && sudo apt -y upgrade

# 依赖
sudo apt -y install build-essential libncurses-dev bison flex libssl-dev libelf-dev cpu-checker qemu-kvm aria2 bc dwarves

# 下载
aria2c -x 10 https://github.com/microsoft/WSL2-Linux-Kernel/archive/refs/tags/linux-msft-wsl-5.15.90.1.tar.gz
# 解压
tar -zxf WSL2-Linux-Kernel-linux-msft-wsl-5.15.90.1.tar.gz

cd WSL2-Linux-Kernel-linux-msft-wsl-5.15.90.1/

# 拷贝官方提供的 WSL2 内核配置
cp Microsoft/config-wsl .config

# 进入内核配置界面
make menuconfig
# 配置内核：
# Virtualization 中关闭 Intel 支持（我是 AMD cpu）。可以将AMD支持从 * 改为M ，使 kvm 不直接编译到内核中，而是采用内核模块的方式，使用内核模块每次重启需要手动加载模块，但可以定制参数，具体方式请自行选择
# Processor type and features 中的 Linux guest support 中开启 KVM Guest support (including kvmclock)


# 编译，使用16个线程，完成后显示如下
$ sudo make -j 16
Kernel: arch/x86/boot/bzImage is ready

# 拷贝 bzImage 到win下
cp arch/x86/boot/bzImage /mnt/c/Users/<user_name>/

# 如果通过 M 模块方式安装，完成后显示如下
$ sudo make modules_install
INSTALL /lib/modules/5.15.90.1-microsoft-standard-WSL2/kernel/arch/x86/kvm/kvm-amd.ko
```

在win用户目录下创建并编辑[.wslconfig](https://learn.microsoft.com/en-us/windows/wsl/wsl-config)，以启用嵌套虚拟化，和使用我们新编译的内核

```toml
[wsl2]
nestedVirtualization=true
kernel=C:\\Users\\<user_name>\\bzImage
```

在win上重启wsl

```shell
wsl --shutdown

# 检查 KVM
kvm-ok

# 检查嵌套 KVM
 cat /sys/module/kvm_amd/parameters/nested
```

之后到wsl上 `uanme -a` 查看内核编译时间确认一下即可



使用 M 模块加载模式使用如下操作

```shell
# kvm 模块参数，注意平台更改，我是amd
sudo cat >> /etc/modprobe.d/kvm-nested.conf << EOF
options kvm-amd nested=1
options kvm-amd enable_shadow_vmcs=1
options kvm-amd enable_apicv=1
options kvm-amd ept=1
EOF

# 加载内核模块
sudo modprobe kvm-amd

# 选择模块方式可以查看
$ lsmod 
Module                  Size  Used by
kvm_amd               114688  0
```

配置启动时自动[加载内核模块](https://blog.51cto.com/myunix/1320782)。没找到 /etc/rc.sysinit，但是找到 /etc/rcS.d/S01kmod 发现类似的内容，并参考[官方文档](https://manpages.ubuntu.com/manpages/bionic/zh_CN/man5/modules-load.d.5.html#%E4%BE%8B%E5%AD%90)，选择在 /etc/modules-load.d 中添加 kvm-amd。但是失败了，重新编译内核更换了编译到内核的模式

```shell
sudo touch /etc/modules-load.d/kvm-amd.conf
# 无权限就直接 sudo vim 改
sudo cat >> /etc/modules-load.d/kvm-amd.conf << EOF
kvm-amd
EOF
```



[官方支持的启用systemd](https://www.cnblogs.com/wswind/p/wsl2-official-systemd.html)；

```shell
# 启用
echo -e "[boot]\nsystemd=true" | sudo tee -a /etc/wsl.conf
# 检查
ps --no-headers -o comm 1
```



[通过 genie 启用 systemd](https://www.ddupan.top/posts/wsl2-kvm/)（新版 wsl 不再需要了）；[22.04 .NET 安装文档](https://learn.microsoft.com/en-us/dotnet/core/install/linux-ubuntu#2204-microsoft-package-feed)。



[安装 virt-manager](https://www.myfreax.com/how-to-install-kvm-on-ubuntu-20-04/)。

```shell
# 安装 virt-manager，依赖于 libvirt
sudo apt install -y qemu qemu-kvm libvirt-daemon libvirt-clients bridge-utils virt-manager
# 已有 systemd，此时 libvirt-daemon 可以正常运行，未启动重启 wsl 以启动
```





## 其他

```shell
sudo apt install -y lsb-core
```





## Nvidia显卡驱动

用于ai学习