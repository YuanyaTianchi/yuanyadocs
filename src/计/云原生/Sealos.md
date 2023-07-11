

+++

title = "Title.sub_title"
description = "Title sub_title quick_start"
tags = ["techn", "computer", "tagx", "Title", "_sub_title", "__quick_start(content)"]

+++



# Sealos

> [官网](https://docs.sealos.io/zh-Hans/)；[文档](https://docs.sealos.io/zh-Hans/docs/Intro)；



## ---快开---



## 安装

> [releases](https://github.com/labring/sealos/releases)；

### 二进制

```shell
wget https://github.com/labring/sealos/releases/download/v4.1.7/sealos_4.1.7_linux_amd64.tar.gz

tar zxvf sealos_4.1.7_linux_amd64.tar.gz sealos && mv sealos /usr/local/bin
```

### 编译

## 一键启动集群

> [集群安装 Kuberentes](https://docs.sealos.io/zh-Hans/docs/getting-started/kuberentes-life-cycle#%E9%9B%86%E7%BE%A4%E5%AE%89%E8%A3%85-kuberentes)；

```shell
# sealos 支持使用 containerd，使用前关闭docker 或者直接删除
$ systemctl disable --now docker
$ apt autoremove -y docker-ce

$ sealos run \
    labring/kubernetes:v1.25.0 \
    labring/calico:v3.24.1 \
    labring/helm:v3.8.2 \
    --cluster containerd-calico-01
    --masters 192.168.11.21,192.168.11.22,192.168.11.23 \
    --nodes 192.168.11.31,192.168.11.32 -p [your-ssh-passwd]
```



## Clusterfile 自定义集群配置

> [自定义配置安装](https://docs.sealos.io/docs/lifecycle-management/operations/run-cluster/gen-apply-cluster)；

### 配置详解

- [sealos Cluster(v1beta1)](https://docs.sealos.io/docs/lifecycle-management/operations/run-cluster/gen-apply-cluster)；

- kubernetes [Configuration APIs](https://kubernetes.io/docs/reference/config-api/)：

    - [kubeadm Configuration (v1beta3)](https://kubernetes.io/docs/reference/config-api/kubeadm-config.v1beta3/)：
        - [ClusterConfiguration](https://kubernetes.io/docs/reference/config-api/kubeadm-config.v1beta3/#kubeadm-k8s-io-v1beta3-ClusterConfiguration)；
        - [InitConfiguration](https://kubernetes.io/docs/reference/config-api/kubeadm-config.v1beta3/#kubeadm-k8s-io-v1beta3-InitConfiguration)；
        - [JoinConfiguration](https://kubernetes.io/docs/reference/config-api/kubeadm-config.v1beta3/#kubeadm-k8s-io-v1beta3-JoinConfiguration)；

    - [kube-proxy Configuration (v1alpha1)](https://kubernetes.io/docs/reference/config-api/kube-proxy-config.v1alpha1/)；

    - [Kubelet Configuration (v1alpha1)](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1alpha1/)；

### 配置生成样例

以 crio 环境和 criu 需求为例。

```shell
$ sealos gen \
    labring/kubernetes-crio:v1.26.5-4.2.1 \
    labring/cilium:v1.12.11 \
    --cluster crio-cilium-01 \
    --masters 192.168.11.21,192.168.11.22,192.168.11.23 \
    --nodes 192.168.11.31,192.168.11.32 \
    -p [your-ssh-passwd] \
    -o .sealos/crio-cilium-01.Clusterfile
$ sealos apply --config Clusterfile
```

修改 CRISocket；添加 featureGates。

```yaml
---
apiVersion: kubeadm.k8s.io/v1beta3
kind: ClusterConfiguration
Networking:
  PodSubnet: 100.11.0.0/16 # 修改 pod cidr

---
apiVersion: kubeadm.k8s.io/v1beta3
kind: InitConfiguration
NodeRegistration:
  CRISocket: unix:///var/run/crio/crio.sock # 修改为crio。注意添加`unix://`前缀，否则会有 Warning

---
apiVersion: kubeadm.k8s.io/v1beta3
kind: JoinConfiguration
NodeRegistration:
  CRISocket: unix:///var/run/crio/crio.sock # 修改为crio

---
apiVersion: kubelet.config.k8s.io/v1beta1
featureGates: # 功能门
  ContainerCheckpoint: true # 启用容器检查点功能
```





## ---基操---

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

