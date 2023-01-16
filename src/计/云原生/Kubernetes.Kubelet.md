

+++

title = "Kubelet"
description = "it.cloud.kubernetes.Kubelet"
tags = ["it","cloud","kubernetes"]

+++



# Kubelet



## 运行时

### 查看

https://kubernetes.io/zh/docs/tasks/administer-cluster/migrating-from-dockershim/find-out-runtime-you-use/

### 指定

配置文件 `/var/lib/kubelet/kubeadm-flags.env`，`--container-runtime-endpoint=unix:///run/containerd/containerd.sock`指定运行时。

### 迁移

参考：[将节点上的容器运行时从 Docker 引擎更改为 Containerd](https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/change-runtime-containerd/)。

安装并启动新的容器运行时

停止调度节点，修改配置文件指定容器运行时

```shell
kubectl drain <node-name> --ignore-daemonsets
```

如果使用了 kubeadm，kubeadm 工具将每个主机的 CRI socket 作为 annotation 存储在该主机的 Node 对象中，通过 `kubectl edit node <node-name>` 修改为新的 cri-socket。

```yaml
kubeadm.alpha.kubernetes.io/cri-socket: unix:///run/containerd/containerd.sock
```

重启 kubelet 并恢复调度

```shell
systemctl restart kubelet
```

恢复调度

```shell
kubectl uncordon <node-name>
```



## 节点清空

https://kubernetes.io/zh/docs/tasks/administer-cluster/safely-drain-node/
