# WSL2

当前版本 [linux-msft-wsl-5.15.90.1](https://github.com/microsoft/WSL2-Linux-Kernel/archive/refs/tags/linux-msft-wsl-5.15.90.1.tar.gz)。



## 家庭版安装 hyper-v

> [Win11 家庭版/专业版开启 Hyper-V](https://zhuanlan.zhihu.com/p/577980646)；

桌面空白处右键-新建-文本文档，命名为 `hyper-v.cmd`，编辑添加如下内容

```shell
pushd "%~dp0"
dir /b %SystemRoot%\servicing\Packages\*Hyper-V*.mum >hyper-v.txt
for /f %%i in ('findstr /i . hyper-v.txt 2^>nul') do dism /online /norestart /add-package:"%SystemRoot%\servicing\Packages\%%i"
del hyper-v.txt
Dism /online /enable-feature /featurename:Microsoft-Hyper-V-All /LimitAccess /ALL
```

右击文件-以管理员身份运行，下载完成后输入 y 重启。

## 启用 windows 功能

在 `Windows 功能` 中启用：

- `适用于Linux的Windows子系统` ；
- `Windows功能-虚拟机平台`，或者包含该功能的 `Hyper-V` ，相比之下将获得额外的管理功能；

以管理员身份运行 Windows PowerShell

```powershell
# 启用：Windows功能-适用于Linux的Windows子系统
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart

# 启用：Windows功能-虚拟机平台
# 或启用：Windows功能-Hyper-V，其包含 Windows功能-虚拟机平台，相比之下将获得额外的管理功能等
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart

# 重启 Windows
```

## 安装系统

Win11 [wsl 安装](https://learn.microsoft.com/zh-cn/windows/wsl/install) 默认版本即为 wsl2

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



### WslRegisterDistribution failed with error: 0x800701bc

下载 适用于 x64 计算机的 WSL2 Linux 内核[更新包](https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi)；



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
# Virtualization 中关闭 Intel 支持（我是 AMD cpu）。（不推荐）可以将AMD支持从 * 改为M ，使 kvm 不直接编译到内核中，而是采用内核模块的方式，使用内核模块每次重启需要手动加载模块，但可以定制参数，具体方式请自行选择
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

在 win 上 `wsl --shutdown` 重启 wsl

```shell
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



## 桥接网络

> [WSL2设置桥接网络](https://www.midlane.top/wiki/pages/viewpage.action?pageId=49676341)；

配置桥接设备，需要在 `windows 功能` 中开启 `Hyper-V`，才能使用虚拟化管理工具。

创建用于桥接的设备：打开 `Hyper-V 管理器`，在右侧 `操作` 栏，连接到本地计算机虚拟服务，选择 `虚拟交换机管理器`，创建名为 `wslbr0` 的 `外部网络` 类型链接到 `有网网卡设备`（我这里是 WIFI 设备）的虚拟交换机

家目录下配置 `.wslconfig` 文件

```toml
[wsl2]
networkingMode=bridged
vmSwitch=wslbr0 # 前面创建的虚拟交换机名称
ipv6=true
```

```shell
# 重启
wsl --shutdown
```



## vscode插件

搜索 WSL，由微软官方提供的远程连接插件





## 其他

```shell
sudo apt install -y lsb-core
```





## Nvidia cuda

WSL2 上无需安装 Nvidia 驱动，Win 上安装了即可，WSL2 上仅需安装 cuda。一下方式二选一即可

### WSL 上直接安装

> [Enable NVIDIA CUDA on WSL](https://learn.microsoft.com/en-us/windows/ai/directml/gpu-cuda-in-wsl)；
>
> [CUDA 工具包下载](https://developer.nvidia.com/cuda-downloads?target_os=Linux&target_arch=x86_64&Distribution=WSL-Ubuntu&target_version=2.0&target_type=deb_local)；[cudnn 下载](https://developer.nvidia.com/rdp/cudnn-download)；
>
> [Windows10/11 WSL2 ubuntu安装nvidia-cuda驱动](https://www.jianshu.com/p/be669d9359e2)；
>
> [win10的wsl2安装cuda并配置pytorch](https://zhuanlan.zhihu.com/p/350399229)；

```shell
# 删除旧的
sudo apt autoremove -y --purge cuda cuda-12-1 cuda-toolkit-12-1

# cuda
wget https://developer.download.nvidia.cn/compute/cuda/repos/wsl-ubuntu/x86_64/cuda-wsl-ubuntu.pin
sudo mv cuda-wsl-ubuntu.pin /etc/apt/preferences.d/cuda-repository-pin-600
wget https://developer.download.nvidia.cn/compute/cuda/12.1.0/local_installers/cuda-repo-wsl-ubuntu-12-1-local_12.1.0-1_amd64.deb
sudo dpkg -i cuda-repo-wsl-ubuntu-12-1-local_12.1.0-1_amd64.deb
sudo cp /var/cuda-repo-wsl-ubuntu-12-1-local/cuda-*-keyring.gpg /usr/share/keyrings/
sudo apt update
sudo apt -y install cuda-toolkit-12-1

# cudnn，须登录下载
tar -xf cudnn-linux-x86_64-8.8.1.3_cuda12-archive.tar.xz
sudo cp /root/compute/ai/cuda/cudnn-linux-x86_64-8.8.1.3_cuda12-archive/lib/libcudnn* /usr/local/cuda-12.1/lib64/
sudo cp /root/compute/ai/cuda/cudnn-linux-x86_64-8.8.1.3_cuda12-archive/include/cudnn* /usr/local/cuda-12.1/include/
```

### nvidia-docker

> [nvidia-docker](https://github.com/NVIDIA/nvidia-docker)；
>
> [安装指南](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html#docker)；
>
> [5步搭建wsl2+cuda+docker解决windows深度学习开发问题](https://zhuanlan.zhihu.com/p/408403790)；

```shell
# 此前须先安装 docker(-ce)
distribution=$(. /etc/os-release;echo $ID$VERSION_ID) \
      && curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg \
      && curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.list | \
            sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
            sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list

sudo apt update && sudo apt install -y nvidia-container-toolkit
```

最后一行报错`/sbin/ldconfig.real: /usr/lib/wsl/lib/libcuda.so.1 is not a symbolic link`

```shell
# 软链
sudo mv /usr/lib/wsl/lib/libcuda.so.1 /usr/lib/wsl/lib/libcuda.so.1.bk && sudo ln -s /usr/lib/wsl/lib/libcuda.so.1.1 /usr/lib/wsl/lib/libcuda.so.1

sudo apt install -y nvidia-container-toolkit
```

```shell
# 配置 Docker 守护进程以识别 NVIDIA 容器运行时
sudo nvidia-ctk runtime configure --runtime=docker

sudo systemctl restart docker

# test
docker run -it --rm --gpus all --name pytorch-sd -u 1000:1000 \
-p 7860:7860 -v /mnt/adocker/stable-diffusion:/workspace/stable-diffusion \
pytorch/pytorch:2.0.0-cuda11.7-cudnn8-runtime bash
```

```shell
# 进入容器后运行一下命令查看是否支持 cuda 计算
python -c 'import torch; print(torch.cuda.is_available())'
```



```shell
# 第一次先激活 conda activate 命令
# 激活环境
source activate
# 退出环境
source deactivate
# 创建 sdwebui 虚拟环境
conda create --name sdwebui
# 激活 sdwebui
conda activate sdwebui

```

### stable-diffusion-docker

> [stable-diffusion-webui-docker Setup](https://github.com/AbdBarho/stable-diffusion-webui-docker/wiki/Setup#make-sure-you-have-the-latest-version-of-docker-and-docker-compose-installed)；
>
> [aria2 errorCode=19](https://github.com/aria2/aria2/issues/1117)；

```shell
git clone https://github.com/AbdBarho/stable-diffusion-webui-docker.git
cd stable-diffusion-webui-docker
# 运行 webui-docker-download-1 容器，完成后自动停止。将下载所有必需的模型/文件，并验证它们的完整性。有些模型下载失败无所谓，或者手动从错误上下文中下载并拷贝到对应目录
docker compose --profile download up --build
# [ui] 从 invoke | auto | auto-cpu | sygil | sygil-sl 中选择，我们选择最流行的 auto
docker compose --profile auto up --build
```



### stable-diffusion-docker-2

> [Docker版Stable Diffusion WebUI](https://zhuanlan.zhihu.com/p/614421868)；