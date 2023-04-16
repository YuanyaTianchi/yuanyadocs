

+++

title = "kubernetes"
description = "it.cloud.kubernetes"
tags = ["it", "cloud", "kubernetes"]

+++



# kubernetes



## - 快开 -

## 安装

> [安装 kubeadm、kubelet 和 kubectl](https://kubernetes.io/zh-cn/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#installing-kubeadm-kubelet-and-kubectl)；

遵照 [kubeadm 安装准备](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#%E5%87%86%E5%A4%87%E5%BC%80%E5%A7%8B)，禁用交换分区，启用[必需端口](https://kubernetes.io/zh/docs/reference/ports-and-protocols/)并进行[检查](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#check-required-ports)，[允许iptables检查桥接流量](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#%E5%85%81%E8%AE%B8-iptables-%E6%A3%80%E6%9F%A5%E6%A1%A5%E6%8E%A5%E6%B5%81%E9%87%8F)。

```shell
# 关闭swap分区(永久)：把加载swap分区的那行记录注释掉，重启生效
sed -ri 's/.*swap.*/#&/' /etc/fstab

# 配置br_netfilter模块
cat <<EOF | tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF
# 加载。也可以reboot以加载
modprobe  br_netfilter && lsmod | grep br_netfilter

# 允许iptables检查桥接流量
cat <<EOF | tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
# 使生效
sysctl --system
```

### ubuntu

```shell
# 关闭swap分区(临时)
swapoff -a

# 防火墙：关闭防火墙或者开放官方指南中指定的端口。ubuntu默认为关闭状态，这里仅查看
ufw status

# 必要工具
apt install -y apt-transport-https ca-certificates curl gnupg lsb-release
```

安装

```shell
# GPG密钥
curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg
# apt源
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

# 查看版本。如果出现版本错误可能需要降低尝 kubenetes 版本
apt update && apt-cache madison kubeadm | grep 1.26
# 安装
apt install -y kubeadm=1.26.3-00 kubelet=1.26.3-00 kubectl=1.26.3-00

# 开机启动并立即启动
systemctl enable --now kubelet
```

### centos

```shell
# 关闭swap分区(临时)
setenforce 0

# 防火墙：关闭防火墙或者开放官方指南中指定的端口
systemctl disable --now firewalld && firewall-cmd --state
```

使用 [kubernetes 阿里源](https://mirrors.aliyun.com/kubernetes/yum/repos/)。

```shell
cat > /etc/yum.repos.d/kubernetes.repo << EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-\$basearch
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
EOF

# 将 SELinux 设置为 permissive 模式（相当于将其禁用）。临时设置立即生效，并永久禁用
setenforce 0
sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config

# 安装。kubelet-x.x.x kubeadm-x.x.x kubectl-x.x.x 可以指定版本
yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
# 设置为开机启动并立即启动，但 kubelet 现在每隔几秒就会重启，因为它陷入了一个等待 kubeadm 指令的死循环
systemctl enable --now kubelet
# 查看 kubelet
systemctl status kubelet
```

### 二进制

#### kubectl

> [安装kubectl](https://kubernetes.io/zh-cn/docs/tasks/tools/install-kubectl-linux/)。

```shell
# kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
# 校验和
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
# 校验
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
# 移动
mv kubectl /usr/local/bin
```



## 开机启动

```shell
# 开机启动并立即启动
# kubelet 现在每隔几秒就会重启，因为当前主机还没有成为 kubernetes 节点，无法正常运作
systemctl enable --now kubelet
```



## 命令补全 & 别名

> [Linux 系统中的 bash 自动补全功能](https://kubernetes.io/zh-cn/docs/tasks/tools/included/optional-kubectl-configs-bash-linux/)。

向 `~/.bashrc` 中添加内容。如果没有 bash-completion 需要先 `apt install -y bash-completion`。

```shell
cat >> ~/.bashrc << EOF



# bash-completion
source /usr/share/bash-completion/bash_completion
# kubectl
source <(kubectl completion bash)
alias kc=kubectl
complete -o default -F __start_kubectl kc
EOF

# 更新
source .bashrc
```



## 时间同步

```shell
# 在各节点上执行如下命令
apt install -y ntpdate
ntpdate time.windows.com
```

 

## - 备份点 -



## 准备工作

### cgroup 驱动

> [dokcer 配置 systemd 作 cgroup 驱动](https://kubernetes.io/zh/docs/setup/production-environment/container-runtimes/#docker)；
>
> [kubelet 配置 systemd 作 cgroup 驱动](https://kubernetes.io/zh/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/#%E9%85%8D%E7%BD%AE-kubelet-%E7%9A%84-cgroup-%E9%A9%B1%E5%8A%A8)；

注意 docker 容器需要配置systemd作cgroup驱动；

1.22及之后`kubeadm init`默认使用 `systemd` 作为 kubelet 的 cgroup 驱动，而不是 `cgroupfs`。

### 控制平面准备

```shell
# 设置 hostname
hostnamectl set-hostname 192.168.31.11

# 主节点添加本机 DNS 映射
sed -i '$a192.168.31.11 kubecpe' /etc/hosts
```

### 数据平面准备

```shell
# 设置 hostname
hostnamectl set-hostname 192.168.31.15
```



## 集群创建

准备3台完成安装的机器

### kubeadm

[kubeadm 参考指南](https://kubernetes.io/zh/docs/reference/setup-tools/kubeadm/)。

#### 控制平面

[初始化控制平面节点](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#%E5%88%9D%E5%A7%8B%E5%8C%96%E6%8E%A7%E5%88%B6%E5%B9%B3%E9%9D%A2%E8%8A%82%E7%82%B9)。注意要再次运行 `kubeadm init`，必须首先[卸载集群](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#tear-down)。如果是非root用户

```shell
# --apiserver-advertise-address 设置 master ip
# --image-repository 指定阿里云镜像仓库地址，因为默认拉取镜像地址 k8s.gcr.io 国内无法访问；
$ kubeadm init \
--node-name 192.168.31.11 \
--control-plane-endpoint=kubecpe \
--apiserver-advertise-address=192.168.31.11 \
--pod-network-cidr=10.244.0.0/16 \
--image-repository registry.aliyuncs.com/google_containers

# 创建成功后得到如下操作提示
......
Your Kubernetes control-plane has initialized successfully!

# 如何开始使用集群（kubectl）
To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

# 如何安装 pod 网络
You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

# 如何添加控制平面节点
You can now join any number of control-plane nodes by copying certificate authorities
and service account keys on each node and then running the following as root:

  kubeadm join kubecpe:6443 --token nhn70c.ck1r8ceu4j4pq3ep \
        --discovery-token-ca-cert-hash sha256:55b9644fc941d4928f40baa30a700973cf375b826fdb8e00012797cdbe0c90c1 \
        --control-plane

# 如何添加工作节点
Then you can join any number of worker nodes by running the following on each as root:

kubeadm join kubecpe:6443 --token nhn70c.ck1r8ceu4j4pq3ep \
        --discovery-token-ca-cert-hash sha256:55b9644fc941d4928f40baa30a700973cf375b826fdb8e00012797cdbe0c90c1
```

#### 使用集群

复制 /etc/kubernetes/admin.conf 到 $HOME/.kube/config 即可使用 kubectl

```shell
mkdir ~/.kube
cp -i /etc/kubernetes/admin.conf ~/.kube/config
chown $(id -u):$(id -g) ~/.kube/config
```

#### pod 网络

[安装 pod 网络](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#pod-network)。从 [flannel github](https://github.com/flannel-io/flannel/blob/v0.15.1/Documentation/kube-flannel.yml) 上复制下来使用，使用 -f 直接指定或者使用 wget 下载后 -f 指定，都会出现 yaml 解析错误

如有需要请[配置网卡](https://huangzhongde.cn/istio/Chapter6/Chapter6-8.html)。

```shell
# 部署 kube-flannel
kubectl apply -f /mnt/share/kube-flannel.yaml
```

#### 工作节点

集群[加入节点](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#join-nodes)。如果前面提示的 token 和 hash 已经找不到了，可以在控制平面节点上查看 token 和 hash

```shell
# 查看 token 列表
kubeadm token list
# 或者创建新的 token
kubeadm token create
# 查看 CA 公钥的哈希值
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
```

然后在工作节点上通过 `kubeadm join --token <token> <control-plane-host>:<control-plane-port> --discovery-token-ca-cert-hash sha256:<hash>` 加入集群。

```shell
# 工作节点添加控制平面节点的 DNS 映射
sed -i '$a192.168.31.11 kubecpe' /etc/hosts
# 加入集群
kubeadm join kubecpe:6443 \
--node-name 192.168.31.11 \
--token st4khy.2j4qmbuge5mh1d76 \
--discovery-token-ca-cert-hash sha256:55b9644fc941d4928f40baa30a700973cf375b826fdb8e00012797cdbe0c90c1
```

#### 节点删除

> 节点清理

先在节点上重置 kubeadm 安装的状态，然后通过 kubectl 删除节点即可。主节点只需调用 `kubeadm reset` 会进行尽力而为的清理

```shell
# 使节点不调度
kubectl drain <node name> --delete-emptydir-data --force --ignore-daemonsets
kubeadm reset
kubectl delete node <node_name>
```

#### 网络卸载

重置过程 `kubeadm reset` 不会重置或清除 iptables 规则或 IPVS 表

```shell
# iptables清理
iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X
```

```shell
# IPVS 清理
ipvsadm -C
```

网卡清理

```shell
# flannel，清理完需重启 kubelet
kubectl delete -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
ifconfig cni0 down
ip link delete cni0
ifconfig flannel.1 down
ip link delete flannel.1
rm -rf /var/lib/cni/
rm -f /etc/cni/net.d/*
systemctl restart kubelet
```



### 最佳实践

#### TLS 证书和秘钥

安装 CFSSL

```shell
$ go get -u github.com/cloudflare/cfssl/cmd/...
$ ls $GOPATH/bin/cfssl*
/root/go/path/bin/cfssl         /root/go/path/bin/cfssl-certinfo  /root/go/path/bin/cfssl-newkey
/root/go/path/bin/cfssl-bundle  /root/go/path/bin/cfssljson       /root/go/path/bin/cfssl-scan
```



##### CA 证书

创建 CA 配置文件。

```shell
$ cfssl print-defaults config > config.json
$ cfssl print-defaults csr > csr.json

# 根据 config.json 文件的格式创建如下的 ca-config.json 文件，过期时间设置成了 87600h
$ cat > ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "87600h"
    },
    "profiles": {
      "kubernetes": {
        "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
        ],
        "expiry": "87600h"
      }
    }
  }
}
EOF
```

字段说明：

- `ca-config.json`：可以定义多个 profiles，分别指定不同的过期时间、使用场景等参数；后续在签名证书时使用某个 profile；
- `signing`：表示该证书可用于签名其它证书；生成的 ca.pem 证书中 `CA=TRUE`；
- `server auth`：表示client可以用该 CA 对server提供的证书进行验证；
- `client auth`：表示server可以用该CA对client提供的证书进行验证；

创建 CA 证书签名请求文件

```shell
cat > ca-csr.json << EOF
> {
>   "CN": "kubernetes",
>   "key": {
>     "algo": "rsa",
>     "size": 2048
>   },
>   "names": [
>     {
>       "C": "CN",
>       "ST": "BeiJing",
>       "L": "BeiJing",
>       "O": "k8s",
>       "OU": "System"
>     }
>   ],
>     "ca": {
>        "expiry": "87600h"
>     }
> }
> EOF
```

字段说明：

- "CN"：`Common Name`，kube-apiserver 从证书中提取该字段作为请求的用户名 (User Name)；浏览器使用该字段验证网站是否合法；
- "O"：`Organization`，kube-apiserver 从证书中提取该字段作为请求用户所属的组 (Group)；

**生成 CA 证书和私钥**。

```shell
$ cfssl gencert -initca ca-csr.json | cfssljson -bare ca
$ ls ca*
ca-config.json  ca.csr  ca-csr.json  ca-key.pem  ca.pem
```

##### kubernetes 证书

创建 kubernetes 证书签名请求文件

```shell
$ cat > kubernetes-csr.json << EOF
> {
>     "CN": "kubernetes",
>     "hosts": [
>       "127.0.0.1",
>       "172.20.0.112",
>       "172.20.0.113",
>       "172.20.0.114",
>       "172.20.0.115",
>       "10.254.0.1",
>       "kubernetes",
>       "kubernetes.default",
>       "kubernetes.default.svc",
>       "kubernetes.default.svc.cluster",
        "algo": "rsa",
>       "kubernetes.default.svc.cluster.local"
>     ],
>     "key": {
>         "algo": "rsa",
>         "size": 2048
>     },
>     "names": [
>         {
>             "C": "CN",
>             "ST": "BeiJing",
>             "L": "BeiJing",
>             "O": "k8s",
>             "OU": "System"
>         }
>     ]
> }
> EOF
```

- 如果 hosts 字段不为空则需要指定授权使用该证书的 **IP 或域名列表**，由于该证书后续被 `etcd` 集群和 `kubernetes master` 集群使用，所以上面分别指定了 `etcd` 集群、`kubernetes master` 集群的主机 IP 和 **`kubernetes` 服务的服务 IP**（一般是 `kube-apiserver` 指定的 `service-cluster-ip-range` 网段的第一个IP，如 10.254.0.1）。
- 这是最小化安装的kubernetes集群，包括一个私有镜像仓库，三个节点的kubernetes集群，以上物理节点的IP也可以更换为主机名。

**生成 kubernetes 证书和私钥**

```shell
$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kubernetes-csr.json | cfssljson -bare kubernetes
$ ls kubernetes*
kubernetes.csr  kubernetes-csr.json  kubernetes-key.pem  kubernetes.pem
```

##### admin 证书

创建 admin 证书签名请求文件

```shell
$ cat > admin-csr.json << EOF
> {
>   "CN": "admin",
>   "hosts": [],
>   "key": {
>     "algo": "rsa",
>     "size": 2048
>   },
>   "names": [
>     {
>       "C": "CN",
>       "ST": "BeiJing",
>       "L": "BeiJing",
>       "O": "system:masters",
>       "OU": "System"
>     }
>   ]
> }
> EOF
```

- 后续 `kube-apiserver` 使用 `RBAC` 对客户端(如 `kubelet`、`kube-proxy`、`Pod`)请求进行授权；
- `kube-apiserver` 预定义了一些 `RBAC` 使用的 `RoleBindings`，如 `cluster-admin` 将 Group `system:masters` 与 Role `cluster-admin` 绑定，该 Role 授予了调用`kube-apiserver` 的**所有 API**的权限；
- O 指定该证书的 Group 为 `system:masters`，`kubelet` 使用该证书访问 `kube-apiserver` 时 ，由于证书被 CA 签名，所以认证通过，同时由于证书用户组为经过预授权的 `system:masters`，所以被授予访问所有 API 的权限；

**生成 admin 证书和私钥**

```shell
$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes admin-csr.json | cfssljson -bare admin
$ ls admin*
admin.csr  admin-csr.json  admin-key.pem  admin.pem
```



**注意**：这个admin 证书，是将来生成管理员用的kube config 配置文件用的，现在我们一般建议使用RBAC 来对kubernetes 进行角色权限控制， kubernetes 将证书中的CN 字段 作为User， O 字段作为 Group（具体参考[ Kubernetes中的用户与身份认证授权](https://jimmysong.io/kubernetes-handbook/guide/authentication.html)中 X509 Client Certs 一段）。

在搭建完 kubernetes 集群后，我们可以通过命令: `kubectl get clusterrolebinding cluster-admin -o yaml` ,查看到 `clusterrolebinding cluster-admin` 的 subjects 的 kind 是 Group，name 是 `system:masters`。 `roleRef` 对象是 `ClusterRole cluster-admin`。 意思是凡是 `system:masters Group` 的 user 或者 `serviceAccount` 都拥有 `cluster-admin` 的角色。 因此我们在使用 kubectl 命令时候，才拥有整个集群的管理权限。

```yaml
$ kubectl get clusterrolebinding cluster-admin -o yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  creationTimestamp: 2017-04-11T11:20:42Z
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: cluster-admin
  resourceVersion: "52"
  selfLink: /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/cluster-admin
  uid: e61b97b2-1ea8-11e7-8cd7-f4e9d49f8ed0
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:masters
```

##### kube-proxy 证书

创建 kube-proxy 证书签名请求文件

```shell
$ cat > kube-proxy-csr.json << EOF
> {
>   "CN": "system:kube-proxy",
>   "hosts": [],
>   "key": {
>     "algo": "rsa",
>     "size": 2048
>   },
>   "names": [
>     {
>       "C": "CN",
>       "ST": "BeiJing",
>       "L": "BeiJing",
>       "O": "k8s",
>       "OU": "System"
>     }
>   ]
> }
> EOF
```

- CN 指定该证书的 User 为 `system:kube-proxy`；
- `kube-apiserver` 预定义的 RoleBinding `system:node-proxier` 将User `system:kube-proxy` 与 Role `system:node-proxier` 绑定，该 Role 授予了调用 `kube-apiserver` Proxy 相关 API 的权限；

生成 kube-proxy 客户端证书和私钥

```shell
$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes  kube-proxy-csr.json | cfssljson -bare kube-proxy
$ ls kube-proxy*
```



##### 校验证书

###### 使用 openssl

以 Kubernetes 证书为例，使用 openssl 命令读取证书信息

- 确认 `Issuer` 字段的内容和 `ca-csr.json` 一致；
- 确认 `Subject` 字段的内容和 `kubernetes-csr.json` 一致；
- 确认 `X509v3 Subject Alternative Name` 字段的内容和 `kubernetes-csr.json` 一致；
- 确认 `X509v3 Key Usage、Extended Key Usage` 字段的内容和 `ca-config.json` 中 `kubernetes` profile 一致；

```shell
$ openssl x509  -noout -text -in  kubernetes.pem
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            03:0e:b4:83:5e:6d:02:fa:bf:62:21:aa:37:47:f1:c4:88:f4:8a:0c
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = CN, ST = BeiJing, L = BeiJing, O = k8s, OU = System, CN = kubernetes
        Validity
            Not Before: Mar 12 12:04:00 2022 GMT
            Not After : Mar  9 12:04:00 2032 GMT
        Subject: C = CN, ST = BeiJing, L = BeiJing, O = k8s, OU = System, CN = kubernetes
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    ......
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Server Authentication, TLS Web Client Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier:
                B7:15:10:43:1B:E1:53:B0:6A:12:56:5B:A5:CD:EF:77:E4:A0:79:F4
            X509v3 Subject Alternative Name:
                DNS:kubernetes, DNS:kubernetes.default, DNS:kubernetes.default.svc, DNS:kubernetes.default.svc.cluster, DNS:kubernetes.default.svc.cluster.local, IP Address:127.0.0.1, IP Address:172.20.0.112, IP Address:172.20.0.113, IP Address:172.20.0.114, IP Address:172.20.0.115, IP Address:10.254.0.1
    Signature Algorithm: sha256WithRSAEncryption
         ......
```

###### 使用cfssl-certinfo

```shell
cfssl-certinfo -cert kubernetes.pem
{
  "subject": {
    "common_name": "kubernetes",
    "country": "CN",
    "organization": "k8s",
    "organizational_unit": "System",
    "locality": "BeiJing",
    "province": "BeiJing",
    "names": [
      "CN",
      "BeiJing",
      "BeiJing",
      "k8s",
      "System",
      "kubernetes"
    ]
  },
  "issuer": {
    "common_name": "kubernetes",
    "country": "CN",
    "organization": "k8s",
    "organizational_unit": "System",
    "locality": "BeiJing",
    "province": "BeiJing",
    "names": [
      "CN",
      "BeiJing",
      "BeiJing",
      "k8s",
      "System",
      "kubernetes"
    ]
  },
  "serial_number": "17454907659222183176468532411310969421540133388",
  "sans": [
    "kubernetes",
    "kubernetes.default",
    "kubernetes.default.svc",
    "kubernetes.default.svc.cluster",
    "kubernetes.default.svc.cluster.local",
    "127.0.0.1",
    "172.20.0.112",
    "172.20.0.113",
    "172.20.0.114",
    "172.20.0.115",
    "10.254.0.1"
  ],
  "not_before": "2022-03-12T12:04:00Z",
  "not_after": "2032-03-09T12:04:00Z",
  "sigalg": "SHA256WithRSA",
  "authority_key_id": "",
  "subject_key_id": "B7:15:10:43:1B:E1:53:B0:6A:12:56:5B:A5:CD:EF:77:E4:A0:79:F4",
  "pem": "-----BEGIN CERTIFICATE-----\nMIIEajCCA1KgAwIBAgIUAw60g15tAvq/YiGqN0fxxIj0igwwDQYJKoZIhvcNAQEL\nBQAwZTELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0Jl\naUppbmcxDDAKBgNVBAoTA2s4czEPMA0GA1UECxMGU3lzdGVtMRMwEQYDVQQDEwpr\ndWJlcm5ldGVzMB4XDTIyMDMxMjEyMDQwMFoXDTMyMDMwOTEyMDQwMFowZTELMAkG\nA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0JlaUppbmcxDDAK\nBgNVBAoTA2s4czEPMA0GA1UECxMGU3lzdGVtMRMwEQYDVQQDEwprdWJlcm5ldGVz\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3QrEXpJnjVNsYnFGnu4f\nJG99y/P49uL0VTEZP4+6X/NRN1Fq9clX8acDHQcYBPud6OSI2Lvah9hxKQTpmIIQ\nky2lEnA64c9BDU5x0CUP5UOzwryjddFshHfvNOg7dySR3hzOKiKac31RKCtOo3OB\nlVctX+ksm1dAsTuGoAvYXKTl2fbR6p+Ew7tkMjbtaghXS9MXVy3RN1t7G3X9Xj9U\ndHnfdTCtTVxJF6kRrOBuP3ad+YDj6QAk8jXKefkFzfezj+w1g1zPJXcTtSko/YiH\n9A0lfiKxiv+6RGoXWQRnMT1J2Or8sKP0yWb0qLkk/+jINCiNulLL+mDQrEVlSSn3\ncwIDAQABo4IBEDCCAQwwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUF\nBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBS3FRBDG+FTsGoS\nVlulze935KB59DCBrQYDVR0RBIGlMIGiggprdWJlcm5ldGVzghJrdWJlcm5ldGVz\nLmRlZmF1bHSCFmt1YmVybmV0ZXMuZGVmYXVsdC5zdmOCHmt1YmVybmV0ZXMuZGVm\nYXVsdC5zdmMuY2x1c3RlcoIka3ViZXJuZXRlcy5kZWZhdWx0LnN2Yy5jbHVzdGVy\nLmxvY2FshwR/AAABhwSsFABwhwSsFABxhwSsFAByhwSsFABzhwQK/gABMA0GCSqG\nSIb3DQEBCwUAA4IBAQAJTaKt98jCvsRo2vq9RvK732UO4tgYdhyAcQTREvsEyjdT\nYlA85xix0I9j5MlBgyfplQLCIobGjjdmkIdVcJsciPaqRdCjurHkDF5pBs4m/qlr\nfhxjEjZr36aT3/taq5xUA//TkCTb9p8Ijm+3AMdPGxovXE/TokEYcXDNasaVYsLY\naBxKDv3Bymwv9RIzPDrcNnn7llsCba6IX+L8u3TB4GIOxv7yH6XwVhqAFKh1yXUZ\n6H8Wky5z8l91virnRBfi3y4jJsl/odIznp16We8VnWCA36eTzBwOOEZiU3MvlPYc\nei7FKCGN5U3ZwwLsTMH2WzWX+Zw+dZGA97FmDYnX\n-----END CERTIFICATE-----\n"
}
```

##### 颁发证书

将生成的证书和秘钥文件（后缀名为`.pem`）拷贝到所有（需要的）机器（包括当前）的 `/etc/kubernetes/ssl` 目录下备用

```shell
mkdir -p /etc/kubernetes/ssl && cp *.pem /etc/kubernetes/ssl
```



#### kubelet

```shell
export KUBE_APISERVER="https://kubecpe:6443"

# 设置集群参数
kubectl config set-cluster kubernetes \
  --certificate-authority=/etc/kubernetes/ssl/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER}
# 设置客户端认证参数
kubectl config set-credentials admin \
  --client-certificate=/etc/kubernetes/ssl/admin.pem \
  --embed-certs=true \
  --client-key=/etc/kubernetes/ssl/admin-key.pem
# 设置上下文参数
kubectl config set-context kubernetes \
  --cluster=kubernetes \
  --user=admin
# 设置默认上下文
kubectl config use-context kubernetes
```

- `admin.pem` 证书 OU 字段值为 `system:masters`，`kube-apiserver` 预定义的 RoleBinding `cluster-admin` 将 Group `system:masters` 与 Role `cluster-admin` 绑定，该 Role 授予了调用`kube-apiserver` 相关 API 的权限；
- 生成的 kubeconfig 被保存到 `~/.kube/config` 文件；

**注意：**`~/.kube/config`文件拥有对该集群的最高权限，请妥善保管。



#### kubeconfig

##### TLS Bootstrapping Token

Token可以是任意的包含128 bit的字符串，可以使用安全的随机数发生器生成。

```shell
export BOOTSTRAP_TOKEN=$(head -c 16 /dev/urandom | od -An -t x | tr -d ' ')

cat > token.csv <<EOF
${BOOTSTRAP_TOKEN},kubelet-bootstrap,10001,"system:kubelet-bootstrap"
EOF

# 检查 token.csv 文件，确认其中的 ${BOOTSTRAP_TOKEN} 环境变量已经被真实的值替换
cat token.csv
c16e2aa67b2d10569d68ee176af17017,kubelet-bootstrap,10001,"system:kubelet-bootstrap"
```

**BOOTSTRAP_TOKEN** 将被写入到 kube-apiserver 使用的 token.csv 文件和 kubelet 使用的 `bootstrap.kubeconfig` 文件，如果后续重新生成了 BOOTSTRAP_TOKEN，则需要：

1. 更新 token.csv 文件，分发到所有机器 (master 和 node）的 /etc/kubernetes/ 目录下，分发到node节点上非必需；
2. 重新生成 bootstrap.kubeconfig 文件，分发到所有 node 机器的 /etc/kubernetes/ 目录下；
3. 重启 kube-apiserver 和 kubelet 进程；
4. 重新 approve kubelet 的 csr 请求；

```shell
cp token.csv /etc/kubernetes/
```

##### kubelet bootstrapping kubeconfig

```shell
cd /etc/kubernetes
export KUBE_APISERVER="https://kubecpe:6443"

# 设置集群参数
kubectl config set-cluster kubernetes \
  --certificate-authority=/etc/kubernetes/ssl/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=bootstrap.kubeconfig

# 设置客户端认证参数
kubectl config set-credentials kubelet-bootstrap \
  --token=${BOOTSTRAP_TOKEN} \
  --kubeconfig=bootstrap.kubeconfig

# 设置上下文参数
kubectl config set-context default \
  --cluster=kubernetes \
  --user=kubelet-bootstrap \
  --kubeconfig=bootstrap.kubeconfig

# 设置默认上下文
kubectl config use-context default --kubeconfig=bootstrap.kubeconfig
```

- `--embed-certs` 为 `true` 时表示将 `certificate-authority` 证书写入到生成的 `bootstrap.kubeconfig` 文件中；
- 设置客户端认证参数时**没有**指定秘钥和证书，后续由 `kube-apiserver` 自动生成；

##### kube-proxy kubeconfig

```shell
export KUBE_APISERVER="https://172.20.0.113:6443"
# 设置集群参数
kubectl config set-cluster kubernetes \
  --certificate-authority=/etc/kubernetes/ssl/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER} \
  --kubeconfig=kube-proxy.kubeconfig
# 设置客户端认证参数
kubectl config set-credentials kube-proxy \
  --client-certificate=/etc/kubernetes/ssl/kube-proxy.pem \
  --client-key=/etc/kubernetes/ssl/kube-proxy-key.pem \
  --embed-certs=true \
  --kubeconfig=kube-proxy.kubeconfig
# 设置上下文参数
kubectl config set-context default \
  --cluster=kubernetes \
  --user=kube-proxy \
  --kubeconfig=kube-proxy.kubeconfig
# 设置默认上下文
kubectl config use-context default --kubeconfig=kube-proxy.kubeconfig
```



##### 分发 kubeconfig

将两个 kubeconfig 文件分发到所有 Node 机器的 `/etc/kubernetes/` 目录

```shell
cp bootstrap.kubeconfig kube-proxy.kubeconfig /etc/kubernetes/
```



#### etcd

##### 准备

添加 hosts

```shell
sed -i '$a192.168.43.201 kubecpe' /etc/hosts
```



##### 安装

```shell
wget https://github.com/etcd-io/etcd/releases/download/v3.4.18/etcd-v3.4.18-linux-amd64.tar.gz
tar -xvf etcd-v3.4.18-linux-amd64.tar.gz
mv etcd-v3.1.5-linux-amd64/etcd* /usr/local/bin
```

##### 创建 etcd 的 systemd unit 文件

在/usr/lib/systemd/system/目录下创建文件etcd.service，内容如下。注意替换IP地址为你自己的etcd集群的主机IP。

```toml
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target
Documentation=https://github.com/coreos

[Service]
Type=notify
WorkingDirectory=/var/lib/etcd/
EnvironmentFile=-/etc/etcd/etcd.conf
ExecStart=/usr/local/bin/etcd \
  --name ${ETCD_NAME} \
  --cert-file=/etc/kubernetes/ssl/kubernetes.pem \
  --key-file=/etc/kubernetes/ssl/kubernetes-key.pem \
  --peer-cert-file=/etc/kubernetes/ssl/kubernetes.pem \
  --peer-key-file=/etc/kubernetes/ssl/kubernetes-key.pem \
  --trusted-ca-file=/etc/kubernetes/ssl/ca.pem \
  --peer-trusted-ca-file=/etc/kubernetes/ssl/ca.pem \
  --initial-advertise-peer-urls ${ETCD_INITIAL_ADVERTISE_PEER_URLS} \
  --listen-peer-urls ${ETCD_LISTEN_PEER_URLS} \
  --listen-client-urls ${ETCD_LISTEN_CLIENT_URLS},http://127.0.0.1:2379 \
  --advertise-client-urls ${ETCD_ADVERTISE_CLIENT_URLS} \
  --initial-cluster-token ${ETCD_INITIAL_CLUSTER_TOKEN} \
  --initial-cluster infra1=https://etce01:2380,infra2=https://etce02:2380,infra3=https://etce03:2380 \
  --initial-cluster-state new \
  --data-dir=${ETCD_DATA_DIR}
Restart=on-failure
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

- 指定 `etcd` 的工作目录为 `/var/lib/etcd`，数据目录为 `/var/lib/etcd`，需在启动服务前创建这个目录，否则启动服务的时候会报错“Failed at step CHDIR spawning /usr/bin/etcd: No such file or directory”；
- 为了保证通信安全，需要指定 etcd 的公私钥(cert-file和key-file)、Peers 通信的公私钥和 CA 证书(peer-cert-file、peer-key-file、peer-trusted-ca-file)、客户端的CA证书（trusted-ca-file）；
- 创建 `kubernetes.pem` 证书时使用的 `kubernetes-csr.json` 文件的 `hosts` 字段**包含所有 etcd 节点的IP**，否则证书校验会出错；
- `--initial-cluster-state` 值为 `new` 时，`--name` 的参数值必须位于 `--initial-cluster` 列表中；

完整 unit 文件见：[etcd.service](https://jimmysong.io/kubernetes-handbook/systemd/etcd.service)

环境变量配置文件`/etc/etcd/etcd.conf`。

```ini
# [member]
ETCD_NAME=infra1
ETCD_DATA_DIR="/var/lib/etcd"
ETCD_LISTEN_PEER_URLS="https://172.20.0.113:2380"
ETCD_LISTEN_CLIENT_URLS="https://172.20.0.113:2379"

#[cluster]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://172.20.0.113:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_ADVERTISE_CLIENT_URLS="https://172.20.0.113:2379"
```



## - 基操 -

## 强制删除

```shell
kubectl get pod | grep Terminating | awk '{print $1}' | xargs kubectl delete pod --grace-period=0 --force
```


