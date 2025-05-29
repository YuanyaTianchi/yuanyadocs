## 快开


## Nvidia docker on WSL2

### Nvidia 驱动

WSL2 上无需安装 Nvidia 驱动，Win 上安装了即可，WSL2 上仅需安装 cuda

### CUDA支持

> [Enable NVIDIA CUDA on WSL](https://learn.microsoft.com/en-us/windows/ai/directml/gpu-cuda-in-wsl)；
>
> [CUDA on WSL2 (nvidia.com)](https://docs.nvidia.com/cuda/wsl-user-guide/contents.html)；
>
> cuda 版本列表：[CUDA Toolkit Archive | NVIDIA Developer](https://developer.nvidia.com/cuda-toolkit-archive)；
>
> cuda 对应 nvidia 驱动版本，基本是向下兼容，不用担心：https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html#id7
>
> WSL 安装 CUDA：[CUDA Toolkit Downloads | NVIDIA Developer](https://developer.nvidia.com/cuda-downloads?target_os=Linux&target_arch=x86_64&Distribution=WSL-Ubuntu&target_version=2.0&target_type=deb_local)；

下载 cuda on wsl2（ #注意使用国内源，否则下面安装时出现E: Unable to correct problems, you have held broken packages.

```shell
# 参考引用获取最新版本
wget https://developer.download.nvidia.com/compute/cuda/repos/wsl-ubuntu/x86_64/cuda-wsl-ubuntu.pin
sudo mv cuda-wsl-ubuntu.pin /etc/apt/preferences.d/cuda-repository-pin-600
wget https://developer.download.nvidia.com/compute/cuda/12.1.1/local_installers/cuda-repo-wsl-ubuntu-12-1-local_12.1.1-1_amd64.deb
sudo dpkg -i cuda-repo-wsl-ubuntu-12-1-local_12.1.1-1_amd64.deb
sudo cp /var/cuda-repo-wsl-ubuntu-12-1-local/cuda-*-keyring.gpg /usr/share/keyrings/
sudo apt-get update
sudo apt-get -y install cuda
```

检查：nvidia-smi 显示的 CUDA Version 取决于 Driver Version 对应的版本，非 wsl 上安装的版本，不必理会

```
$ nvidia-smi
+-----------------------------------------------------------------------------------------+
| NVIDIA-SMI 560.35.02              Driver Version: 560.94         CUDA Version: 12.6     |
|-----------------------------------------+------------------------+----------------------+
| GPU  Name                 Persistence-M | Bus-Id          Disp.A | Volatile Uncorr. ECC |
| Fan  Temp   Perf          Pwr:Usage/Cap |           Memory-Usage | GPU-Util  Compute M. |
|                                         |                        |               MIG M. |
|=========================================+========================+======================|
|   0  NVIDIA GeForce RTX 4080        On  |   00000000:01:00.0  On |                  N/A |
| 30%   42C    P8             26W /  340W |    2028MiB /  16376MiB |     10%      Default |
|                                         |                        |                  N/A |
+-----------------------------------------+------------------------+----------------------+
                                                                                         
+-----------------------------------------------------------------------------------------+
| Processes:                                                                              |
|  GPU   GI   CI        PID   Type   Process name                              GPU Memory |
|        ID   ID                                                               Usage      |
|=========================================================================================|
|  No running processes found                                                             |
+-----------------------------------------------------------------------------------------+
```

卸载

> 如果使用包管理器安装则需要使用包管理器卸载：[How to uninstall cuda 12.4 - Graphics / Linux / Linux - NVIDIA Developer Forums](https://forums.developer.nvidia.com/t/how-to-uninstall-cuda-12-4/287741)；
>
> [wsl linux CUDA安装、卸载、清理、版本降级、升级过程详解_cuda降级-CSDN博客](https://blog.csdn.net/xiangshangdemayi/article/details/144918217)

```shell
sudo apt-get --purge remove 'cuda*'
sudo apt-get --purge remove 'nvidia*'
sudo rm -rf /usr/local/cuda*
sudo dpkg --purge --force-all cuda-repo-wsl-ubuntu-12-1-local
sudo rm -rf /var/cuda-repo-wsl-ubuntu-12-1-local
sudo apt-get purge 'cuda*'
sudo apt-get autoremove -y
grep -r "cuda" /etc/apt/sources.list /etc/apt/sources.list.d/

sudo rm /etc/apt/sources.list.d/cuda-wsl-ubuntu-12-1-local.list
sudo rm -rf /var/lib/apt/lists/*
```



### 容器工具包

> [NVIDIA/nvidia-container-toolkit: Build and run containers leveraging NVIDIA GPUs (github.com)](https://github.com/NVIDIA/nvidia-container-toolkit)；

安装 nvidia-container-toolkit

```shell
# 配置存储库
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | \
    sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg && \
curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | \
    sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
    sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list && \
sudo apt update

# 安装 NVIDIA Container Toolkit 软件包
sudo apt-get install -y nvidia-container-toolkit
```

配置 `/etc/docker/daemon.json`

```shell
sudo nvidia-ctk runtime configure --runtime=docker && \
sudo systemctl restart docker
```

检查配置。**如果使用 Windows Docker Desktop，docker 进程实际上不被 wsl systemctl 管理，需要手动复制配置到 Docker Desktop 的设置 Setting->Docker Engine 中，并 Apply & Restart**。

```shell
$ cat /etc/docker/daemon.json
{
    "runtimes": {
        "nvidia": {
            "args": [],
            "path": "nvidia-container-runtime"
        }
    }
}
```

检查容器

```shell
$ docker pull ubuntu && sudo docker run --rm --runtime=nvidia --gpus all ubuntu nvidia-smi
+-----------------------------------------------------------------------------+
| NVIDIA-SMI xxx.xx.xx    Driver Version: 528.24       CUDA Version: 12.0     |
|-------------------------------+----------------------+----------------------+
| GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC |
| Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. |
|                               |                      |               MIG M. |
|===============================+======================+======================|
|   0  NVIDIA GeForce ...  On   | 00000000:01:00.0  On |                  N/A |
|  0%   49C    P8    21W / 340W |   2850MiB / 16376MiB |     12%      Default |
|                               |                      |                  N/A |
+-------------------------------+----------------------+----------------------+
                                             
+-----------------------------------------------------------------------------+
| Processes:                                                                  |
|  GPU   GI   CI        PID   Type   Process name                  GPU Memory |
|        ID   ID                                                   Usage      |
|=============================================================================|
|    0   N/A  N/A        29      G   /Xwayland                       N/A      |
|    0   N/A  N/A        35      G   /Xwayland                       N/A      |
|    0   N/A  N/A        38      G   /Xwayland                       N/A      |
+-----------------------------------------------------------------------------+
```
