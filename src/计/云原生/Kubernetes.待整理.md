

+++

title = "Kubernetes.hello"
description = "Kubernetes.hello"
tags = ["it", "container"]

+++



# Kubernetes

## client

语法：`kubectl [option] [type...] <name...> [flags...] `

- option：指定要对资源执行的操作
  - get获取、delete删除
- type：指定资源类型。资源类型是大小写敏感的，有单数、复数、缩略的形式
  - node(nodes)、namespace(namespaces,ns)、pod(pods)、deployment(deployments,deploy)、service(services,svc)、
- name：指定资源的名称，名称也是大小写敏感的，如果省略名称，则会显示所有的资源
- flags：指定可选的参数。
  - `-n` 指定 namespace（缺省 namespace 为 default），`-o wide` 表示查看更广泛的信息，`-f` 表示指定文件，`--all` 表示所有
  - 例如可用 -s 或者 -server参数指定Kubernetes API server的地址和端口

```sh
# 表示查看 Namespace 为 kube-system 下的 Pod & Service 信息
kubectl get pod,svc -n kube-system -o wide
# 查看单个指定 Service
kubectl get svc nginx01svc -o wide

# 根据 yaml 创建 Pod
kubectl apply -f nginx01.yaml

# 删除指定对象
kubectl delete pod,svc,statefulset nginx01,nginx01svc,nginx02svc,nginx01sts*
# 删除所有 Service,StatefulSet
kubectl delete svc,statefulset --all
# 根据 yaml 文件删除
kubectl delete -f nginx01.yaml
```

更多请查看 help

```sh
kubectl --help
```





### 获取

```shell
kubectl get pod,svc
```

#### Field Selectors

https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/

用于 kubectl get 时的筛选过滤

```shell
kubectl get pods --field-selector spec.=Running
kubectl get pods --field-selector status.phase=Running
```







获取 containerId

```shell

```



### 删除

###### 筛选删除

```shell
kubectl get pod -n prom | grep xxx | awk '{print $1}' | xargs kubectl delete pod -n prom
```

###### 强制删除·

```shell
kubectl delete pod xxx --grace-period=0 --force
```



## pod

pod作为kubernetes最小操作单位，其中可以可以运行1或多个容器，每一个Pod都有一个特殊的被称为 “根容器”的Pause容器。一个pod内容器共享网络和文件系统，可以通过进程间通信和文件共享组合完成服务。

kubernetes业务主要可以分为以下几种

- 长期伺服型：long-running
- 批处理型：batch
- 节点后台支撑型：node-daemon
- 有状态应用型：stateful application

我们知道容器通过namespace和cgroup隔离资源，同一个pod内都属于同一namespace

pod创建时首先会创建其根容器`pause`，而我们的业务容器创建时会将被添加到与 pause 相同的namespace中，共享网络资源(ip,mac,port...)



共享存储

```yaml
apiversion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers: # 这里两个容器，指定挂载到相同地址即可共享存储
  - name: write
    image: centos
    command: ["bash","-c","for i in (1..100 );do echo $i >> /data/hello;sleep 1;done"]
    volumeMounts:
    - name: data
      mountPath: /data # 挂载地址
  - name: read
    image: centos
    command: [ "bash", "-c", "tail -f l data/hello"]
    volumeMounts:
    - name: data
      mountPath: /data # 挂载地址
  volumes: # 指定data挂载作为数据卷
  - name: data
    emptyDir: {}
```

### 信息

https://blog.csdn.net/alex_yangchuansheng/article/details/107373639

根据进程名称获取 pid

```shell
pstree | grep <process_name>
```

根据 pid 根据获取 podId

```shell
cat /proc/<pid>/mountinfo | grep "etc-hosts" | awk -F / {'print $6'}
```

根据 podId 获取 pod 完整名称

```shell
crictl ps -o json | jq  '.[][].labels | select (.["io.kubernetes.pod.uid"] == "8e018a8e-4aaa-4ac6-986a-1a5133a4bcf1") | .["io.kubernetes.pod.name"]'|uniq
```

获取 pod token 信息

```shell
cat /var/lib/kubelet/pods/<pod_id>/volumes/kubernetes.io~secret/<pod_name>/token
```



### 生命周期状态

- **CrashLoopBackOff**： 奔溃循环回退。表示 Pod 不断重启和崩坏，往复循环。可能是程序错误导致的崩坏，也可能是内存超出 limit 等。
- **Evicted**： 表示 Pod 从节点被驱逐，不再在该节点上调度。多见于资源不足等情况
- **OOMKilled**： Pod 中容器使用内存超过 limit 值时，会被 kernel 因 OOM kill 掉，此时会出现短暂的 OOMKILL 状态，表示因 OOM 而被 kill 掉了，之后将尝试重新调度和启动该 Pod
  - 2021.08.28，部署dhacnpd到消费者云测试环境，就因内存超出 limit 导致程序 CrashLoopBackOff，并最终转为 OOMKilled。因为同一套配置在 2012 lab 开发环境以及其它产品线都未出现过内存不足，所以以为是代码问题，没太想到，并且消费者节点较多，node pod 数据自然更多

### QoS

https://blog.csdn.net/weixin_44729138/article/details/112602635

QoS：Quality of Service，服务质量。当 Kubernetes 创建一个 Pod 时，根据Pod中容器对 cpu/memory 的 requests/limits，会给这个 Pod 分配一个 QoS 等级，可能是以下之一：

- **Guaranteed**：保证的。Pod 中各个容器都必须设置相同的 cpu&memory 的 limit 值，且如果有一个容器设置了 request ，则其它所有容器也必须设置 request 值，且必须 request==limit（如果一个容器仅指明 limit，则 request 缺省值就等于 limit）。
- **Burstable**：爆发的。Pod 中至少有一个容器有 cpu|memory 的 requests 不满足 Guarantee 等级的要求，即 cpu|memory 的 request<limit。
- **BestEffort**：尽力的。Pod 中所有容器都没有任何 cpu&memory 的 requests&limit。

Guaranteed：将被创建到

```yaml
    spec:
      containers:
      - name: nginxGuaranteed
        image: nginx:latest
        resources:
          limits:
            cpu: 100m
            memory: 100Mi
          requests:
            cpu: 100m
            memory: 100Mi
```

Burstable

```yaml
    spec:
      containers:
      - name: nginxBurstable
        image: nginx:latest
        resources:
          limits:
            cpu: 100m
            memory: 100Mi
          requests:
            cpu: 90m
            memory: 90Mi
```

BestEffort

```yaml
    spec:
      containers:
      - name: nginx
        image: nginx:latest
```



### pod镜像拉取策略

`IfNotPresent`-镜像在宿主机上不存在则拉取(缺省策略)；`Always`-每次创建Pod都会重新拉取一次镜像；`Never`-Pod永远不会主动拉取镜像

```yaml
apiversion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: nginx
    imaqe: nainx:1.14
    imagePullPolicy: Always # 指定拉取策略
```

### pod资源限制

如果对pod进行了资源限制，则pod不会调度 资源不足其限制值的node

```yaml
apiversion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers: # 这里两个容器，指定挂载到相同地址即可共享存储
  - name: db
    image: mysql
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "password"
    resource: # 资源限制
      requests: # 调度所需资源
        memory: "64Mi"
        cpu: "250m"
      limits: # 最大所占用资源
        memory: "128Mi"
        cpu: "500m"
```

### pod重启机制

`Always`-当容器终止退出后，总是重启容器(默认策略)【nginx等，需要不断提供服务】；`OnFailure`-当容器异常退出（退出状态码非0）时，才重启容器；`Never`-当容器终止退出，从不重启容器 【批量任务】。某个容器出现问题，就会触发Pod重启机制

```yaml
apiversion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers: # 这里两个容器，指定挂载到相同地址即可共享存储
  - name: busybox
    image: busybox:1.28.4
    args:
    - /bin/sh
    - -c
    - sleep 36000
  restartPolicy: Never # 重启策略

```

### pod健康检查

```sh
# STATUS字段可以看到容器状态
$ kubectl get pod
NAME                    READY   STATUS    RESTARTS   AGE
nginx-f89759699-r59xb   1/1     Running   0          22h
```

但是也许程序出现内存溢出，容器仍然在运行，但是服务已经不能提供，此时需要使用应用层面检查，有两种检查

- livenessProbe：存活检查，如果检查失败，将杀死容器，根据Pod的restartPolicy【重启策略】来操作
- readinessProbe：就绪检查，如果检查失败，Kubernetes会把Pod从Service endpoints中剔除

```yaml
apiversion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: liveness
  image: busyboxargs:
  - /bin/sh
  - -c
  - touch /tmp/healthy;sleep 30;
  livenessProbe:
    exec:
      command:
      - cat
      - /tmp/healthy
      initialDelaySeconds: 5
      periodseconds: 5
```

Probe支持以下三种检查方式

- http Get：发送HTTP请求，返回200 - 400 范围状态码为成功
- exec：执行Shell命令返回状态码是0为成功
- tcpSocket：发起TCP Socket建立成功

#### pod调度策略

pod创建流程

- 首先创建一个pod，API Server 接收到请求，把创建信息写入在 etcd
- 然后创建 Scheduler，监控API Server是否有新的Pod，如果有的话，会通过调度算法，把pod调度到某个node上（在etcd记录分配信息）
- node上通过 `kubelet -- apiserver ` 读取 etcd，拿到分配给自己的 pod 信息，然后通过 docker 创建运行容器
- 容器正常运行之后 node 向 apiserver 发送更新pod状态的请求，由apiserver更新该pod状态到etcd ，并给node响应确认信息

###  影响Pod调度的属性

Pod资源限制对Pod的调度会有影响，因为需要找到满足资源条件的节点进行调蓄

节点选择器标签影响Pod调度：其实就是如果有多个环境，比如dev、test、prod，然后环境之间所用的资源配置不同，自然影响调度。节点环境可以通过 `kubectl label node k8snode01 env_role=prod` 设置，相当于给 k8snode01 打上了 `env_role=prod` 的标签，

```yaml
apiversion: v1
kind: Pod
metadata:
  name: my-pod
spac: 
  nodeSelector: 
    env_role: dev
```

节点亲和性 **nodeAffinity** 和 nodeSelector 相似，根据节点上标签约束来影响Pod调度到哪些节点上。支持常用操作符：in、NotIn、Exists、Gt、Lt、DoesNotExists，反亲和性就是和亲和性的一些反条件，如 NotIn、DoesNotExists 等

```yaml
apiversion: v1
kind: Pod
metadata:
  name: my-pod
spac: 
  affinity:
    nodeAffinitv:
      requiredDuringschedulingIgnoredDuringExecution: # 硬亲和性: 必须满足约束条件
        nodeselectorTerms:
        - matchExpressions:
          - key: env_role
            operator: In # 
            values:
            - dev
            - test
      preferredDuringschedulingIgnoredDuringExecution: # 软亲和性: 尝试满足条件，不保证一定满足
      - weight: 1
        preference:
          matchExpressions:
          - key: group
            operator: In
            values:
            - otherprod
  containers:
  - name: webdemo
    image: nginx
```

###  污点 & 污点容忍

nodeSelector 和 NodeAffinity，都是Prod调度到某些节点上，属于Pod的属性，是在调度的时候实现的

Taint 污点：节点不做普通分配调度，是节点属性

场景

- 专用节点【限制ip】
- 配置特定硬件的节点【固态硬盘】
- 基于Taint驱逐【在node1不放，在node2放】

```sh
# 查看 Taints
$ kubectl describe node k8smaster | grep Taints
Taints:             node-role.kubernetes.io/master:NoSchedule
```

污点值有三个

- NoSchedule：一定不被调度
- PreferNoSchedule：尽量不被调度【也有被调度的几率】
- NoExecute：不会调度，并且还会驱逐Node已有Pod

将节点添加为污点：`kubectl taint node [nodename] key=value:污点值`

```sh
$ kubectl taint node k8snode1 env_role=yes:NoSchedule
```

删除污点

```
$ kubectl taint node k8snode01 env_role:NoSchedule-
```

操作：现在创建多个Pod，查看最后分配到Node上的情况

```sh
# 创建一个nginx的pod（之前已经创建过了）
$ kubectl create deployment nginx --image=nginx
# 查看pod与节点的调度关系，可以看到该 pod 被 k8snode01 执行
$ kubectl get pods -o wide
NAME                    READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginx-f89759699-r59xb   1/1     Running   0          22h   10.244.1.2   k8snode01   <none>           <none>
# 将 pod 副本数置为 5
$ kubectl scale deployment nginx --replicas=5
deployment.apps/nginx scaled
# 再次查看，STATUS 可能是 ContainerCreating，需要一定时间
$ kubectl get pods -o wide
NAME                    READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginx-f89759699-7s6wg   1/1     Running   0          95s   10.244.2.3   k8snode02   <none>           <none>
nginx-f89759699-96x9r   1/1     Running   0          95s   10.244.2.2   k8snode02   <none>           <none>
nginx-f89759699-d5fx8   1/1     Running   0          95s   10.244.1.4   k8snode01   <none>           <none>
nginx-f89759699-r59xb   1/1     Running   0          22h   10.244.1.2   k8snode01   <none>           <none>
nginx-f89759699-s8t9k   1/1     Running   0          95s   10.244.1.3   k8snode01   <none>           <none>

# 这样可以删除所有部署名为 nginx 的 pod
$ kubectl delete deployment nginx
deployment.apps "nginx" deleted

# 现在给 k8snode01 打上污点
$ kubectl taint node k8snode01 env_role=yes:NoSchedule
node/k8snode01 tainted
$ kubectl describe node k8smaster | grep Taints
Taints:             node-role.kubernetes.io/master:NoSchedule
# 重新创建 pod，并设置5个副本，可以发现 k8snode01 将不参与调度了
$ kubectl create deployment nginx --image=nginx
$ kubectl scale deployment nginx --replicas=5
$ kubectl get pods -o wide
NAME                    READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginx-f89759699-29n24   1/1     Running   0          10m   10.244.2.5   k8snode02   <none>           <none>
nginx-f89759699-2lhfb   1/1     Running   0          10m   10.244.2.7   k8snode02   <none>           <none>
nginx-f89759699-5kltc   1/1     Running   0          11m   10.244.2.4   k8snode02   <none>           <none>
nginx-f89759699-f8559   1/1     Running   0          10m   10.244.2.8   k8snode02   <none>           <none>
nginx-f89759699-ndxxk   1/1     Running   0          10m   10.244.2.6   k8snode02   <none>           <none>
```

污点容忍：可以容忍污点被调度，但也可能被不调度

```yaml
spec:
  tolerations:
  - key: "key"
    operator: "Equal "
    value: "value"
    effect: "Noschedule"
  containers:
  - name: we bdemo
    image: nginx
```

## volume

https://kubernetes.io/zh/docs/concepts/storage/volumes/

https://www.cnblogs.com/nj-duzi/p/14105180.html

### 资源清单

```yaml
spec:
  volumes:
  - name: <string>        # 存储卷名称标识，仅可使用DNS标签格式的字符，在当前Pod中必须唯一
    VOLUME_TYPE: <Object> # 存储卷插件及具体的目标存储供给方的相关配置
  containers:
  - name: ...
    image: ...
    volumeMounts:
    - name: <string>             # 容器要挂载的 volume name,必须匹配存储卷列表中的某项定义
      mountPatch: <string>       # 容器文件系统上的挂载点路径
      readOnly: <bool>           # 是否挂载为只读模式,默认为"否"
      subPath: <string>          # 挂载存储卷上的一个子目录至指定的挂载点
      subPathExpr: <string>      # 挂载由指定的模式匹配到的存储卷的文件或目录至挂载点
      mountPropagation: <string> # 挂载卷的传播模式
```

### hostPath

https://kubernetes.io/zh/docs/concepts/storage/volumes/#hostpath

### emptyDir

https://kubernetes.io/zh/docs/concepts/storage/volumes/#emptydir

https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#emptydirvolumesource-v1-core

```go
// pkg/volume/emptydir/empty_dir.go
// EmptyDir volumes are temporary directories exposed to the pod.
// These do not persist beyond the lifetime of a pod.
type emptyDir struct {
	pod           *v1.Pod
	volName       string
	sizeLimit     *resource.Quantity
	medium        v1.StorageMedium
	mounter       mount.Interface
	mountDetector mountDetector
	plugin        *emptyDirPlugin
	volume.MetricsProvider
}

// staging/src/k8s.io/api/core/v1/types.go
// StorageMedium defines ways that storage can be allocated to a volume.
type StorageMedium string
const (
	StorageMediumDefault         StorageMedium = ""           // use whatever the default is for the node, assume anything we don't explicitly handle is this
	StorageMediumMemory          StorageMedium = "Memory"     // use memory (e.g. tmpfs on linux)
	StorageMediumHugePages       StorageMedium = "HugePages"  // use hugepages
	StorageMediumHugePagesPrefix StorageMedium = "HugePages-" // prefix for full medium notation HugePages-<size>
)
```



```shell
spec:
  volumes:
  - name: <string>
    emptyDir: <Object>
      sizeLimit: 128Mi # emptyDir 大小限制
```

emptydir 生命周期与 pod 一致，不具有持久性



```yaml
apiVersion: v1
kind: Pod
metadata:
  name: emptydir-test    #注意此处只能是小写字母或数字,大写字母会报错
spec:
  volumes:
  - name: cache-volume
    emptyDir: {}
  containers:
  - name: nginx-test
    image: nginx
    imagePullPolicy: IfNotPresent
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
```



emptydir 在宿主机上的路径

```shell
/var/lib/kubelet/pods/<PODID>/Volumes/Kubernetes.io~emptydir/<VOLUME_NAME>
```

pod volumesd 的使用

`.spec.vloumes`声明pod的volumes信息

`.spec.containers.volumesMounts`声明container如何使用pod的volumes

多个container共享一个volumes的时候,通过`.spec.containers.volumeMounts.subPath`隔离不同容器在同个volumes上数据存储的路径





## Controller

Pod是通过Controller实现应用的运维，比如弹性伸缩，滚动升级等。Pod 和 Controller之间是通过label标签来建立关系，同时Controller又被称为控制器工作负载

###  Deployment控制器应用

- Deployment控制器可以部署无状态应用
- 管理 Pod 和 ReplicaSet
- 部署，滚动升级等功能
- 应用场景：web服务，微服务

Deployment表示用户对K8S集群的一次更新操作。Deployment是一个比RS( Replica Set, RS) 应用模型更广的 API 对象，可以是创建一个新的服务，更新一个新的服务，也可以是滚动升级一个服务。滚动升级一个服务，实际是创建一个新的RS，然后逐渐将新 RS 中副本数增加到理想状态，将旧RS中的副本数减少到0的复合操作。这样一个复合操作用一个RS是不好描述的，所以用一个更通用的Deployment来描述。以K8S的发展方向，未来对所有长期伺服型的业务的管理，都会通过Deployment来管理。

###  Deployment部署应用

之前都是使用`kubectl create`创建pod。每次都需要重复操作

```sh
$ kubectl create deployment nginx01 --image=nginx
```

还有另一种方法，通过 `-o yaml > xxx.yaml` 输出一个 yaml 以方便复用，然后通过`kubectl apply -f xxx.yaml` 指定 yaml 创建 pod

```sh
# --dry-run=client 表示尝试创建，不会实际创建，以此输出一个 yaml 以方便复用
$ kubectl create deployment nginx01 --image=nginx --dry-run=client -o yaml > nginx01.yaml
$ cat nginx.yaml
```

```yaml
apiVersion: apps/v1
kind: Deployment # Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx01
  name: nginx01
spec:
  replicas: 1
  selector: # 这里的 selector 与下面的 labels 就是 Pod 和 Controller 之间建立关系的桥梁
    matchLabels:
      app: nginx01
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels: # 这里的labels
        app: nginx01
    spec:
      containers:
      - image: nginx
        name: nginx
        resources: {}
status: {}
```

```sh
# 然后就可以通过 nginx.yaml 快速创建 Pod
$ kubectl apply -f nginx01.yaml
deployment.apps/nginx01 created

# 查看 Pod & Deployment
$ kubectl get pod,deployment -o wide
NAME                           READY   STATUS    RESTARTS   AGE     IP            NODE        NOMINATED NODE   READINESS GATES
pod/nginx01-58cf646cdb-fcfhf   1/1     Running   0          2m54s   10.244.2.14   k8snode02   <none>           <none>

NAME                      READY   UP-TO-DATE   AVAILABLE   AGE     CONTAINERS   IMAGES   SELECTOR
deployment.apps/nginx01   1/1     1            1           2m54s   nginx        nginx    app=nginx01
```

```sh
# 通过暴露端口来创建 Service 以对外提供服务。--port指定内部的端口号，--target-port指定外部的端口号，--type类型，--name名称
$ kubectl expose deployment nginx01 --port=80 --target-port=80 --type=NodePort --name=nginx01svc
# 查看 Service
$ kubectl get svc
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP   10.96.0.1      <none>        443/TCP        38h
nginx01svc   NodePort    10.97.73.124   <none>        80:31065/TCP   61s

# 同样可以导出 service 配置
$ kubectl expose deployment nginx01 --port=80 --target-port=80 --type=NodePort --name=nginxsvc01 --dry-run=client -o yaml > nginx01service.yaml
$ cat nginx01service.yaml
```

```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2021-05-03T06:28:11Z"
  labels:
    app: nginx01
  managedFields:
  - apiVersion: v1
    fieldsType: FieldsV1
    fieldsV1:
      f:metadata:
        f:labels:
          .: {}
          f:app: {}
      f:spec:
        f:externalTrafficPolicy: {}
        f:ports:
          .: {}
          k:{"port":80,"protocol":"TCP"}:
            .: {}
            f:port: {}
            f:protocol: {}
            f:targetPort: {}
        f:selector:
          .: {}
          f:app: {}
        f:sessionAffinity: {}
        f:type: {}
    manager: kubectl
    operation: Update
    time: "2021-05-03T06:28:11Z"
  name: nginx01
  namespace: default
  resourceVersion: "331891"
  selfLink: /api/v1/namespaces/default/services/nginx01
  uid: da381736-cde3-4522-8684-11b817022768
spec:
  clusterIP: 10.97.73.123
  externalTrafficPolicy: Cluster
  ports:
  - nodePort: 32269
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx01
  sessionAffinity: None
  type: NodePort
status:
  loadBalancer: {}
```

```sh
# 注意，未创建 Service 的 yaml 与已创建 Seervice 的 yaml 是不一样的。这里删除刚刚创建的 Service ，再导出 yaml 查看。可以发现与
$ kubectl delete svc nginx01svc
$ kubectl expose deployment nginx01 --port=80 --target-port=80 --type=NodePort --name=nginx01 --dry-run=client -o yaml > nginx01service.yaml
$ cat nginx01service.yaml
```

```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: nginx01
  name: nginx01
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx01
status:
  loadBalancer: {}
```





### 升级 & 回滚

修改前面的 nginx.yaml 中 `- image: nginx` 为 `- image: nginx:1.14`。

```sh
# 创建配置并创建pod
$ kubectl create deployment nginx1 --image=nginx --dry-run=client -o yaml > nginx1.yaml
$ kubectl apply -f nginx1.yaml
# k8snode01 上可以看到拉取了1.14版本的nginx
$ docker images
REPOSITORY                                           TAG                 IMAGE ID            CREATED             SIZE
nginx                                                latest              62d49f9bab67        2 weeks ago         133MB
registry.aliyuncs.com/google_containers/kube-proxy   v1.18.0             43940c34f24f        13 months ago       117MB
registry.aliyuncs.com/google_containers/pause        3.2                 80d28bedfe5d        14 months ago       683kB
nginx                                                1.14                295c7be07902        2 years ago         109MB
quay.io/coreos/flannel                               v0.11.0-amd64       ff281650a721        2 years ago         52.6MB
# 回到 master 上升级 nginx 到 1.15
$ kubectl set image deployment nginx1 nginx=nginx:1.15
deployment.apps/nginx1 image updated
# 立即查看，可以看到在创建新的 nginx1 被创建运行，而旧的 nginx1 终结被替代，可以保证服务不中断
$ kubectl get pods
NAME                      READY   STATUS        RESTARTS   AGE
nginx1-756b7b9b75-vq8h8   1/1     Running       0          35s
nginx1-866bf8d976-2vplg   0/1     Terminating   0          6m35s
# 查看状态，是成功的
$ kubectl rollout status deployment nginx1
deployment "nginx1" successfully rolled out
# 查看历史版本
$ kubectl rollout history deployment nginx1
deployment.apps/nginx1 
REVISION  CHANGE-CAUSE
1         <none>
2         <none>

# 回滚，将回到版本1-nginx:1.14
$ kubectl rollout undo deployment nginx1
deployment.apps/nginx1 rolled back
# 可以看到又是一个不一样的 nginx1-866bf8d976-75rlr
$ kubectl get pods
NAME                      READY   STATUS    RESTARTS   AGE
nginx1-866bf8d976-75rlr   1/1     Running   0          30s
# 查看状态
$ kubectl rollout status deployment nginx1
deployment "nginx1" successfully rolled out
# 查看历史版本
$ kubectl rollout history deployment nginx1
deployment.apps/nginx1 
REVISION  CHANGE-CAUSE
2         <none>
3         <none>

# 指定版本回滚。版本2-nginx:1.15
$ kubectl rollout undo deployment nginx1 --to-revision=2
$ kubectl get pods
NAME                      READY   STATUS    RESTARTS   AGE
nginx1-756b7b9b75-zcm8r   1/1     Running   0          12s
$ kubectl rollout history deployment nginx1
deployment.apps/nginx1 
REVISION  CHANGE-CAUSE
3         <none>
4         <none>
```

### 弹性伸缩

```sh
# 即之前用过的 scale 操作设置 replicas
$ kubectl scale deployment web --replicas=10
```



## Service

Deployment 只是保证了支撑服务的微服务Pod的数量，但是没有解决如何访问这些服务的问题。一个Pod只是一个运行服务的实例，随时可能在一个节点上停止，在另一个节点以一个新的IP启动一个新的Pod，因此不能以确定的IP和端口号提供服务。

要稳定地提供服务发现和负载均衡能力

在K8S集群中，客户端需要访问的服务就是Service对象，每个Service会对应一个集群内部有效的虚拟IP，集群内部通过虚拟IP访问一个服务。

在K8S集群中，微服务的负载均衡是由kube-proxy实现的。kube-proxy是k8s集群内部的负载均衡器。它是一个分布式代理服务器，在K8S的每个节点上都有一个；这一设计体现了它的伸缩性优势，需要访问服务的节点越多，提供负载均衡能力的kube-proxy就越多，高可用节点也随之增多。与之相比，我们平时在服务器端使用反向代理作负载均衡，还要进一步解决反向代理的高可用问题。

Service作为注册中心提供服务发现能力，防止Pod失联。Pod每次创建都对应一个IP地址，而这个IP地址是短暂的，每次随着Pod的更新都会变化，假设当我们的前端页面有多个Pod时候，同时后端也多个Pod，这个时候，他们之间的相互访问，就需要通过注册中心，拿到Pod的IP地址，然后去访问对应的Pod



Pod 负载均衡策略。前端的Pod访问到后端的Pod，中间会通过Service一层，而Service在这里还能做负载均衡。

- 随机
- 轮询
- 响应比



 Pod 与 Service的关系：与 Controller 一样根据 label 和 selector 建立关联

当然在访问 service 的时候，其实也是需要有一个ip地址，这个ip肯定不是pod的ip地址，而是 虚拟IP `vip`，



Service类型，这里列出常用三种

- ClusterIp：集群内部访问(默认类型)
- NodePort：对外访问应用使用
- LoadBalancer：对外访问应用使用，公有云

```sh
# 比如，这里用 nginx01 导出一个yaml文件
$ kubectl expose deployment nginx01 --port=80 --target-port=80 --dry-run=client -o yaml > nginx01svc.yaml
```

```yaml
apiVersion: v1
kind: Service # 导出的 yaml 文件中 kind 为 Service，而不是 Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx01
  name: nginx01svc
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx01
  type: NodePort # 添加一个 type 字段指定 Service 类型
status:
  loadBalancer: {}
```

```sh
# 将创建 nginx01 的 NodePort 类型service
$ kubectl apply -f service.yaml
service/nginx01svc created
```



node一般是在内网进行部署，而外网一般是不能访问到的，那么如何访问的呢？

- 找到一台可以通过外网访问的机器，安装nginx，反向代理
- 手动把可以访问的节点添加到nginx中

如果我们使用LoadBalancer，就会有负载均衡的控制器，类似于nginx的功能，就不需要自己添加到nginx上



## Controller详解

###  Statefulset

Statefulset主要是用来部署有状态应用。对于StatefulSet中的Pod，每个Pod挂载自己独立的存储，如果一个Pod出现故障，从其他节点启动一个同样名字的Pod，要挂载上原来Pod的存储继续以它的状态提供服务。

- 无状态应用：我们原来使用 deployment，部署的都是无状态的应用，那什么是无状态应用？
  - 认为Pod都是一样的
  - 没有顺序要求
  - 不考虑应用在哪个node上运行
  - 能够进行随意伸缩和扩展
- 有状态应用：上述的因素都需要考虑到
  - 让每个Pod独立
  - 保持Pod启动顺序和唯一性
  - 唯一的网络标识符，持久存储
  - 有序，比如mysql中的主从

适合StatefulSet的业务包括数据库服务MySQL 和 PostgreSQL，集群化管理服务Zookeeper、etcd等有状态服务

StatefulSet的另一种典型应用场景是作为一种比普通容器更稳定可靠的模拟虚拟机的机制。传统的虚拟机正是一种有状态的宠物，运维人员需要不断地维护它，容器刚开始流行时，我们用容器来模拟虚拟机使用，所有状态都保存在容器里，而这已被证明是非常不安全、不可靠的。

使用StatefulSet，Pod仍然可以通过漂移到不同节点提供高可用，而存储也可以通过外挂的存储来提供 高可靠性，StatefulSet做的只是将确定的Pod与确定的存储关联起来保证状态的连续性。

#### 部署有状态应用

要部署StatefulSet，需要先通过`clusterIP：None`创建无头Service，整体 nginxsts.yaml 代码如下

```yaml
# 首先创建一个名为 nginx01 的 Deployment
apiVersion: apps/v1
kind: Deployment # Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx01
  name: nginx01
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx01
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx01
    spec:
      containers:
      - image: nginx
        name: nginx
        resources: {}
status: {}
---
# 然后创建一个 nginx01 的无头 Service，即 clusterIP 为 None 的 Service
apiVersion: v1
kind: Service # Service
metadata:
  name: nginx01svc
  labels:
    app: nginx01
spec:
  ports:
  - port: 80
    name: nginx01
  clusterIP: None # 无头
  selector:
    app: nginx01
---
# 最后创建 StatefulSet
apiVersion: apps/v1
kind: StatefulSet # StatefulSet
metadata:
  name: nginxsts
  namespace: default
spec:
  serviceName: nginx
  replicas: 3
  selector:
    matchLabels:
      app: nginx01
  template:
    metadata:
      labels:
        app: nginx01
    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
```

```sh
$ kubectl apply -f nginx01sts.yaml
deployment.apps/nginx01 created
service/nginx01svc created
statefulset.apps/nginxsts created

# 查看一下。可以看到每一个 pod/nginx01sts 都有唯一名字，它们是逐一创建的；service/nginx01svc 中，CLUSTER-IP 为 None，PORT中也没有外部端口映射信息，将不再对外提供服务；
$ kubectl get pod,deployment,svc,statefulset -o wide
NAME                           READY   STATUS              RESTARTS   AGE   IP            NODE        NOMINATED NODE   READINESS GATES
pod/nginx01-58cf646cdb-x7hfz   1/1     Running             0          36s   10.244.1.10   k8snode01   <none>           <none>
pod/nginx01sts-0               1/1     Running             0          36s   10.244.2.17   k8snode02   <none>           <none>
pod/nginx01sts-1               1/1     Running             0          19s   10.244.2.18   k8snode02   <none>           <none>
pod/nginx01sts-2               0/1     ContainerCreating   0          2s    <none>        k8snode01   <none>           <none>

NAME                      READY   UP-TO-DATE   AVAILABLE   AGE   CONTAINERS   IMAGES   SELECTOR
deployment.apps/nginx01   1/1     1            1           36s   nginx        nginx    app=nginx01

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE   SELECTOR
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   39h   <none>
service/nginx01svc   ClusterIP   None         <none>        80/TCP    36s   app=nginx01

NAME                          READY   AGE   CONTAINERS   IMAGES
statefulset.apps/nginx01sts   2/3     36s   nginx        nginx:latest

# 删除
$ kubectl delete statefulset --all
```

这里有状态的约定，肯定不是简简单单通过名称来进行约定，而是更加复杂的操作

- deployment：是有身份的，有唯一标识
- statefulset：根据主机名 + 按照一定规则生成域名

StatefulSet 中的每个 pod 有唯一的主机名，并且有唯一的域名。格式：主机名称.service名称.名称空间.svc.cluster.local。例如：nginx01sts-0.default.svc.cluster.local



### DaemonSet

DaemonSet 是后台支撑型服务，主要是用来部署守护进程，为每一个需要的节点提供各种支撑服务

长期伺服型 & 批处理型 的核心在于业务应用，可能有些节点运行多个同类业务的Pod，有些节点上又没有这类的Pod运行；

后台支撑型服务 的核心关注点在于K8S集群中的节点(物理机或虚拟机)，要保证每个节点上都有一个此类Pod运行。节点可能是所有集群节点，也可能是通过 nodeSelector 选定的一些特定节点。典型的后台支撑型服务包括：存储、日志和监控等。在每个节点上支撑K8S集群运行的服务。

守护进程在我们每个节点上，运行的是同一个pod，新加入的节点也同样运行在同一个pod里面



这里在每个node节点安装数据采集工具，dsdemo.yaml 代码如下

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: dsdemo
  labels:
    app: filebeat
spec:
  selector:
    matchLabels:
      app: filebeat
  template:
    metadata:
      labels:
        app: filebeat
    spec:
      containers:
      - name: logs
        image: nginx
        ports:
        - containerPort: 80
        volumeMounts:
        - name: varlog
          mountPath: /tmp/log
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
```

```sh
$ kubectl apply -f dsdemo.yaml
$ kubectl get pod -o wide
NAME                       READY   STATUS    RESTARTS   AGE    IP            NODE        NOMINATED NODE   READINESS GATES
dsdemo-fdr7k               1/1     Running   0          106s   10.244.1.12   k8snode01   <none>           <none>
dsdemo-qv826               1/1     Running   0          106s   10.244.2.19   k8snode02   <none>           <none>
# 进入某个 Pod
$ kubectl exec -it dsdemo-fdr7k bash
# 在 /tmp/log 中记录了收集的日志信息
dsdemo-fdr7k$ ls /tmp/log 
Xorg.0.log         chrony         glusterfs           pluto            speech-dispatcher    vboxadd-setup.log.2
Xorg.0.log.old     containers     grubby              pods             spooler              vboxadd-setup.log.3
Xorg.9.log         cron           grubby_prune_debug  ppp              spooler-20210502     vboxadd-setup.log.4
anaconda           cron-20210502  lastlog             qemu-ga          swtpm                vboxadd-uninstall.log
audit              cups           libvirt             rhsm             tallylog             vmware-vmtoolsd-root.log
boot.log           dmesg          maillog             sa               tuned                vmware-vmusr-root.log
boot.log-20210502  dmesg.old      maillog-20210502    samba            vboxadd-install.log  wpa_supplicant.log
btmp               firewalld      messages            secure           vboxadd-setup.log    wtmp
btmp-20210502      gdm            messages-20210502   secure-20210502  vboxadd-setup.log.1  yum.log
```



### Job

一次性任务 和 周期定时任务

Job是K8S中用来控制批处理型任务的API对象。批处理业务与长期伺服业务的主要区别就是批处理业务的运行有头有尾，而长期伺服业务在用户不停止的情况下永远运行。Job管理的Pod根据用户的设置把任务成功完成就自动退出了。成功完成的标志根据不同的 spec.completions 策略而不同：单Pod型任务有一个Pod成功就标志完成；定数成功行任务保证有N个任务全部成功；工作队列性任务根据应用确定的全局成功而标志成功。

这里创建一个计算 π 的一次性任务，代码写入 job01.yaml

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi01
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl
        command: ["perl","-Mbignum=bpi","-wle","print bpi (2000)"]
      restartPolicy: Never
  backoffLimit: 4
```

```sh
$ kubectl apply -f job01.yaml
job.batch/pi01 created
$ kubectl get job -o wide
NAME   COMPLETIONS   DURATION   AGE   CONTAINERS   IMAGES   SELECTOR
pi01   0/1           32s        32s   pi           perl     controller-uid=7a839648-9ba4-4794-8bd0-200f4a86a4be
# 计算完成后可以看到，pod 中 STATUS 为 Completed，job 中 COMPLETIONS 变为 1/1
$ kubectl get pod,job -o wide
NAME         READY   STATUS      RESTARTS   AGE    IP            NODE        NOMINATED NODE   READINESS GATES
pi01-swvwk   0/1     Completed   0          90s    10.244.1.13   k8snode01   <none>           <none>

NAME             COMPLETIONS   DURATION   AGE     CONTAINERS   IMAGES   SELECTOR
job.batch/pi01   1/1           51s        2m35s   pi           perl     controller-uid=7a839648-9ba4-4794-8bd0-200f4a86a4be

# 可以通过日志，查看一次性任务的结果 π=3.14 .....
$ kubectl logs pi01-swvwk
3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491298336733624406566430860213949463952247371907021798609437027705392171762931767523846748184676694051320005681271452635608277857713427577896091736371787214684409012249534301465495853710507922796892589235420199561121290219608640344181598136297747713099605187072113499999983729780499510597317328160963185950244594553469083026425223082533446850352619311881710100031378387528865875332083814206171776691473035982534904287554687311595628638823537875937519577818577805321712268066130019278766111959092164201989380952572010654858632788659361533818279682303019520353018529689957736225994138912497217752834791315155748572424541506959508295331168617278558890750983817546374649393192550604009277016711390098488240128583616035637076601047101819429555961989467678374494482553797747268471040475346462080466842590694912933136770289891521047521620569660240580381501935112533824300355876402474964732639141992726042699227967823547816360093417216412199245863150302861829745557067498385054945885869269956909272107975093029553211653449872027559602364806654991198818347977535663698074265425278625518184175746728909777727938000816470600161452491921732172147723501414419735685481613611573525521334757418494684385233239073941433345477624168625189835694855620992192221842725502542568876717904946016534668049886272327917860857843838279679766814541009538837863609506800642251252051173929848960841284886269456042419652850222106611863067442786220391949450471237137869609563643719172874677646575739624138908658326459958133904780275901
```

### CronJob

通过 spec.schedule 字段配置 cron 表达式指定定时任务周期，cronjob01.yaml

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: hello01
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox
            args:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure
```

```sh
$ kubectl apply -f cronjob01.yaml

# 每次执行，就会多出一个 pod/hello01，同时会从最开始的 pod 删除
$ kubectl get pod,cronjob -o wide
NAME                           READY   STATUS              RESTARTS   AGE    IP            NODE        NOMINATED NODE   READINESS GATES
pod/hello01-1620038700-j6npm   0/1     Completed           0          81s    10.244.2.20   k8snode02   <none>           <none>
pod/hello01-1620038760-b86s7   0/1     Completed           0          21s    10.244.2.22   k8snode02   <none>           <none>
pod/hello01-1620038880-t5ttf   0/1     ContainerCreating   0          84s    10.244.1.16   k8snode01   <none>           <none>

NAME                    SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE    CONTAINERS   IMAGES    SELECTOR
cronjob.batch/hello01   */1 * * * *   False     0        26s             112s   hello        busybox   <none>

# 通过 kubectl log 查看结果
$ kubectl logs hello01-1620038760-b86s7
Mon May  3 10:46:08 UTC 2021
Hello from the Kubernetes cluster
```



###  Replication Controller

Replication Controller 简称 **RC**，是K8S中的复制控制器。RC是K8S集群中最早的保证Pod高可用的API对象。通过监控运行中的Pod来保证集群中运行指定数目的Pod副本。指定的数目可以是多个也可以是1个；少于指定数目，RC就会启动新的Pod副本；多于指定数目，RC就会杀死多余的Pod副本。

即使在指定数目为1的情况下，通过RC运行Pod也比直接运行Pod更明智，因为RC也可以发挥它高可用的能力，保证永远有一个Pod在运行。RC是K8S中较早期的技术概念，只适用于长期伺服型的业务类型，比如控制Pod提供高可用的Web服务。

#### Replica Set

Replica Set 检查 RS，也就是副本集。RS是新一代的RC，提供同样高可用能力，区别主要在于RS后来居上，能够支持更多种类的匹配模式。副本集对象一般不单独使用，而是作为Deployment的理想状态参数来使用



## 配置管理

### Secret

Secret的主要作用就是加密数据，然后存在etcd里面，让Pod容器以**挂载Volume**方式进行访问。场景：用户名 和 密码进行加密；一般场景的是对某个字符串进行编码加密。这里以base64为例

```sh
$ echo -n 'root' | base64
cm9vdA==
$ echo -n 'yuanya' | base64
eXVhbnlh
```

先通过 secretdemo01.yaml 创建 Secret

```yml
apiVersion: v1
kind: Secret
metadata:
  name: secretdemo01
type: Opaque
data:
  username: cm9vdA==
  password: eXVhbnlh
```

```sh
$ kubectl create -f secretdemo01.yaml
secret/secretdemo01 created
$ kubectl get secret -o wide
NAME                  TYPE                                  DATA   AGE
default-token-s7hvt   kubernetes.io/service-account-token   3      10h
secretdemo01          Opaque                                2      62s
```

然后将 Secret 通过挂载 Volume 的方式挂载到某个 Pod 上。当然实际上还有另一种以变量形式将 Secret 添加到 Pod

####  Volume 挂载

通过 nginxpod01.yaml 创建 Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginxpod01
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - name: foo
      mountPath: "/etc/foo"
      readOnly: true
  volumes:
  - name: foo
    secret:
      secretName: secretdemo01 # 将 secretdemo01 挂载到 nginxpod01 中名为 foo 的 Volume 上
```

```sh
$ kubectl apply -f nginxpod01.yaml
pod/nginxpod01 created
$ kubectl get pod -o wide
NAME         READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginxpod01   1/1     Running   0          25s   10.244.2.6   k8snode02   <none>           <none>

# 进入容器
$ kubectl exec -it nginxpod01 bash

secretdemo01$ ls /etc/foo
password username
secretdemo01$ cat username
root
secretdemo01$ cat password
yuanya

# 如果要删除这个 pod 可以这样做
$ kubectl delete -f nginxpod01.yaml
pod "nginxpod01" deleted
```



#### 变量形式

通过 nginxpod02.yaml 创建 Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginxpod02
spec:
  containers:
  - name: nginx
    image: nginx
    env:
    - name: SECRETDEMO01_USERNAME
      valueFrom:
        secretKeyRef:
          name: secretdemo01
          key: username
    - name: SECRETDEMO01_PASSWORD
      valueFrom:
        secretKeyRef:
          name: secretdemo01
          key: password
```

```sh
$ kubectl apply -f nginxpod02.yaml
pod/nginxpod02 created
$ kubectl get pod -o wide
NAME         READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginxpod02   1/1     Running   0          26s   10.244.2.5   k8snode02   <none>           <none>

# 进入容器
$ kubectl exec -it nginxpod02 bash
# 用户名
secretdemo01$ echo $SECRETDEMO01_USERNAME
root
# 密码
secretdemo01$ echo $SECRETDEMO01_PASSWORD
yuanya

$ kubectl delete -f nginxpod02.yaml
pod "nginxpod02" deleted
```



### ConfigMap

ConfigMap 作用是存储不加密的数据到etcd中，让 Pod 以 Volume挂载 或 变量形式 让容器可以访问。应用场景：配置文件

假设有配置文件 redis.properties

```properties
redis.port=127.0.0.1
redis.port=6379
redis.password=yuanya
```

通过 kubectl create 创建 ConfigMap

```sh
# --from-file 指定根据哪个文件创建 ConfigMap
$ kubectl create configmap redis-config --from-file=redis.properties
configmap/redis-config created
# 可简写为 cm
$ kubectl get configmap
NAME           DATA   AGE
redis-config   1      50s

# 查看详细信息
$ kubectl describe cm redis-config
Name:         redis-config
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
redis.properties:
----
redis.port=127.0.0.1
redis.port=6379
redis.password=yuanya

Events:  <none>
```

####  Volume挂载

通过 busyboxpod01.yaml 创建 Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busyboxpod01
spec:
  containers:
  - name: busybox
    image: busybox
    command: ["/bin/sh","-c","cat /etc/config/redis.properties"] # 配置输出信息
    volumeMounts:
    - name: redis-cm-volume
      mountPath: /etc/config
  volumes:
  - name: redis-cm-volume
    configMap:
      name: redis-config # 指定 cm
  restartPolicy: Never
```

```sh
$ kubectl apply -f busyboxpod01.yaml
pod/busyboxpod01 created
$ kubectl get pod -o wide
NAME           READY   STATUS      RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
busyboxpod01   0/1     Completed   0          10s   10.244.1.4   k8snode01   <none>           <none>

# 通过 kubectl logs 查看结果
$ kubectl logs busyboxpod01
redis.port=127.0.0.1
redis.port=6379
redis.password=yuanya

$ kubectl delete -f busyboxpod01.yaml
pod "busyboxpod01" deleted
```

#### 变量形式

通过 mycm01.yaml 创建 ConfigMap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mycm01
  namespace: default
data:
  special.level: info
  special.type: hello
```

```sh
$ kubectl apply -f mycm01.yaml
configmap/mycm01 created
$ kubectl get cm
NAME     DATA   AGE
mycm01   2      103s
```

通过 busyboxpod02.yaml 创建 Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busyboxpod02
spec:
  containers:
  - name: busybox
    image: busybox
    command: ["/bin/sh","-c","echo $(LEVEL)$ (TYPE)"] # 配置输出信息
    env:
    - name: LEVEL # 环境变量名
      valueFrom:
        secretKeyRef:
          name: mycm01 # 指定 cm
          key: special.level # 指定变量
    - name: TYPE
      valueFrom:
        secretKeyRef:
          name: mycm01
          key: special.type
  restartPolicy: Never
```

```sh
$ kubectl apply -f busyboxpod02.yaml
pod/busyboxpod01 created
$ kubectl get pod -o wide
NAME         READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginxpod02   1/1     Running   0          26s   10.244.2.5   k8snode02   <none>           <none>

# 通过 kubectl logs 查看结果
$ kubectl logs busyboxpod02
info hello

$ kubectl delete -f busyboxpod02.yaml & kubectl delete -f mycm01.yaml
pod "busyboxpod02" deleted
configmap "mycm01" deleted
```

##  集群安全机制

当我们访问K8S集群时，需要经过三个步骤完成具体操作

- 认证
- 鉴权【授权】
- 准入控制

进行访问的时候，都需要经过 apiserver， apiserver做统一协调，比如门卫

- 访问过程中，需要证书、token、或者用户名和密码
- 如果访问pod需要serviceAccount

####  认证

保证传输安全：使用的是6443端口。对外不暴露8080端口，只能内部访问

客户端身份认证常用方式

- https证书认证，基于ca证书
- http token认证，通过token来识别用户
- http基本认证，用户名 + 密码认证

####  鉴权

基于RBAC进行鉴权操作

基于角色访问控制

####  准入控制

就是准入控制器的列表，如果列表有请求内容就通过，没有的话 就拒绝



### RBAC介绍

基于角色的访问控制，为某个角色设置访问内容，然后用户分配该角色后，就拥有该角色的访问权限

k8s中有默认的几个角色

- role：特定命名空间访问权限
- ClusterRole：所有命名空间的访问权限

角色绑定

- roleBinding：角色绑定到主体
- ClusterRoleBinding：集群角色绑定到主体

主体

- user：用户
- group：用户组
- serviceAccount：服务账号

##### 命名空间

创建命名空间

```sh
# 通过 kubectl create 创建 Namespace
$ kubectl create ns rolens01
namespace/rolens01 created
# 查看 Namespace
$ kubectl get namespace
NAME              STATUS   AGE
default           Active   12h
kube-node-lease   Active   12h
kube-public       Active   12h
kube-system       Active   12h
rolens01          Active   9s
# 在 rolens01 上运行一个 nginx
$ kubectl run nginx --image=nginx -n rolens01
pod/nginx created
# -n 指定 Namespace，默认仅查看 namespace/default
$ kubectl get pod -n rolens01 -o wide
NAME    READY   STATUS    RESTARTS   AGE   IP           NODE        NOMINATED NODE   READINESS GATES
nginx   1/1     Running   0          65s   10.244.2.7   k8snode02   <none>           <none>
```

##### 角色

通过 role01.yaml 创建 Role

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: role01
  namespace: rolens01 # 指定 Namespace
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods"]
  verbs: ["get","watch","list"] # 表示该角色只对 pods 有 get、watch、list 权限
```

```sh
$ kubectl apply -f role01.yaml
role.rbac.authorization.k8s.io/role01 created
$ kubectl get role -n roledemo
NAME     CREATED AT
role01   2021-05-04T06:02:13Z
```

##### 角色绑定

通过 rolebinding01.yaml 创建一个 RoleBinding

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rolebinding01
  namespace: rolens01
subjects:
- kind: User
  name: lucy # Name is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role # this must be Role or clusterRole
  name: role01 # this must match the name of the Role or clusterRole you wish to bind to
  apiGroup: rbac.authorization.k8s.io
```

```sh
$ kubectl apply -f rolebinding01.yaml
rolebinding.rbac.authorization.k8s.io/rolebinding01 created
$ kubectl get role,rolebinding -n rolens01
NAME                                    CREATED AT
role.rbac.authorization.k8s.io/role01   2021-05-04T06:02:13Z

NAME                                                  ROLE          AGE
rolebinding.rbac.authorization.k8s.io/rolebinding01   Role/role01   82s
```

#####  使用证书识别身份

// todo: 这里需要一些ca相关文件，后续在搞

```sh
$ mkdir mary && cd mary
$ vim rbac-user.sh
```

需要一个 rbac-user.sh 证书脚本文件，文件内容如下

```sh
cat > mary-csr.json << EOF
{
  "CN": "mary",
  "hosts": [],
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "name": [
    {
      "C": "CN",
      "L": "BeiJing",
      "ST": "BeiJing"
    }
  ]
}
EOF

cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes mary-csr.json | cfssljson -bare mary

kubectl config set-cluster kubernetes \
  --certificate-authority=ca.pem \
  --embed-certs=true \
  --server=https://192.168.31.101:6443 \ # 注意ip
  --kubeconfig=mary-kubeconfig

kubectl config set-credentials mary \
  --client-key=mary-key.pem \
  --client-certificate=mary.pem \
  --embed-certs=true \
  --kubeconfig=mary-kubeconfig

kubectl config set-context default \
  --cluster=kubernetes \
  --user=mary \
  --kubeconfig=mary-kubeconfig

kubectl config use-context default \
  --kubeconfig=mary-kubeconfig
```

```sh
# 执行
$ sh rbac-user.sh
# 测试
$ kubectl get pod,svc -n rolens01 -o wide
```



## Ingress Controller

Ingress 不是某一个具体的 Controller，它有很多实现，如 Kubernetes 官方提供的 ingress-nginx

原来我们需要将端口号对外暴露，通过 ip + 端口号就可以进行访问

原来是使用 Service 中的 NodePort 来实现

- 在每个节点上都会启动端口
- 在访问的时候通过任何节点，通过ip + 端口号就能实现访问

但是NodePort还存在一些缺陷

- 因为端口不能重复，所以每个端口只能使用一次，一个端口对应一个应用
- 实际访问中都是用域名，根据不同域名跳转到不同端口服务中

pod 和 ingress 是通过 service 进行关联的，而ingress作为统一入口，由service关联一组pod

- 首先 service 就是关联我们的 pod
- 然后 ingress 作为入口，首先需要到 service，然后发现一组pod
- 发现pod后，就可以做负载均衡等操作

 Ingress工作流程：在实际的访问中，我们都是需要维护很多域名，a.com 和 b.com。ingress作为统一入口，然后不同的域名对应的不同的Service，然后service管理不同的pod

 使用Ingress：ingress不是内置的组件，需要我们单独的安装。

- 部署ingress Controller【需要下载官方的】
- 创建ingress规则【对哪个Pod、名称空间配置规则



### 部署 Ingress

先用 Service 暴露端口

```sh
$ kubectl create deployment nginx01 --image=nginx
deployment.apps/nginx01 created
$ kubectl get pod -o wide
NAME                       READY   STATUS              RESTARTS   AGE   IP       NODE        NOMINATED NODE   READINESS GATES
nginx01-58cf646cdb-gvn6z   0/1     ContainerCreating   0          7s    <none>   k8snode01   <none>           <none>
$ kubectl expose deployment nginx01 --port=80 --target-port=80 --type=NodePort --name=nginx01svc
service/nginxsvc01 exposed
```

通过 it.cloudnative.Kubernetes.resource.ingress-controller.yaml 部署 Ingress Controller。注意文件中 `hostNetwork: true` 一项，如果是 false 一定要改为 true，否则无法访问到。

```sh
$ kubectl apply -f ingress-controller.yaml 
namespace/ingress-nginx created
configmap/nginx-configuration created
configmap/tcp-services created
configmap/udp-services created
serviceaccount/nginx-ingress-serviceaccount created
clusterrole.rbac.authorization.k8s.io/nginx-ingress-clusterrole created
role.rbac.authorization.k8s.io/nginx-ingress-role created
rolebinding.rbac.authorization.k8s.io/nginx-ingress-role-nisa-binding created
clusterrolebinding.rbac.authorization.k8s.io/nginx-ingress-clusterrole-nisa-binding created
daemonset.apps/nginx-ingress-controller created
service/ingress-nginx created

$ kubectl get pods -n ingress-nginx -o wide
NAME                             READY   STATUS    RESTARTS   AGE   IP               NODE        NOMINATED NODE   READINESS GATES
nginx-ingress-controller-g68dz   1/1     Running   0          18m   192.168.31.111   k8snode01   <none>           <none>
nginx-ingress-controller-vq9bm   1/1     Running   0          18m   192.168.31.112   k8snode02   <none>           <none>
```

通过 exampl-ingress.yml 创建 Ingress

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress # 部署 ingress-nginx 之后可以创建 Ingress Controller
metadata: 
  name: exampl-ingress # 这个名字更改过后将访问不了，不知道什么原因...
spec:
  rules: # 定制 Ingress 规则
  - host: example.ingressdemo.com
    http:
      paths:
      - path: /
        backend:
          serviceName: nginx01svc # 从这里找到前面创建的 nginx01svc
          servicePort: 80
```

```sh
$ kubectl apply -f ingress01.yml
ingress.networking.k8s.io/ingress01 created

$ kubectl get pod -n ingress-nginx -o wide && echo \ && kubectl get pod,svc,ingress -o wide
NAME                             READY   STATUS    RESTARTS   AGE   IP               NODE        NOMINATED NODE   READINESS GATES
nginx-ingress-controller-g68dz   1/1     Running   0          70m   192.168.31.111   k8snode01   <none>           <none>
nginx-ingress-controller-vq9bm   1/1     Running   0          70m   192.168.31.112   k8snode02   <none>           <none>
 
NAME                           READY   STATUS    RESTARTS   AGE    IP           NODE        NOMINATED NODE   READINESS GATES
pod/nginx01-58cf646cdb-gvn6z   1/1     Running   0          169m   10.244.1.7   k8snode01   <none>           <none>

NAME                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE   SELECTOR
service/kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP        17h   <none>
service/nginx01svc   NodePort    10.101.112.240   <none>        80:31090/TCP   38m   app=nginx01

NAME                                CLASS    HOSTS                     ADDRESS   PORTS   AGE
ingress.extensions/exampl-ingress   <none>   example.ingressdemo.com             80      4m15s
```

node节点上查看80和443端口监听状态，可以发现这两个端口存在监听

```sh
$ netstat -antp | grep:80
$ netstat -antp | grep:443
```

在 windows 的 hosts文件，添加域名访问规则【因为我们没有域名解析，所以只能这样做】，然后访问 http://example.ingressdemo.com/

```
192.168.31.111   example.ingressdemo.com
```



## Helm yaml管理

Helm就是一个包管理工具【类似于npm】



为什么需要 Helm

首先在原来项目中都是基于yaml文件来进行部署发布的，而目前项目大部分微服务化或者模块化，会分成很多个组件来部署，每个组件可能对应一个deployment.yaml、service.yaml、Ingress.yaml，还可能存在各种依赖关系，这样一个项目如果有5个组件，很可能就有15个不同的yaml文件，这些yaml分散存放，如果某天进行项目恢复的话，很难知道部署顺序，依赖关系等，而所有这些包括

- 基于yaml配置的集中存放
- 基于项目的打包
- 组件间的依赖

但是这种方式部署，会有什么问题呢？

- 如果使用之前部署单一应用，少数服务的应用，比较合适
- 但如果部署微服务项目，可能有几十个服务，每个服务都有一套 yaml 文件，需要维护大量的 yaml 文件，版本管理特别不方便

Helm的引入，就是为了解决这个问题

- 使用Helm可以把这些YAML文件作为整体管理
- 实现YAML文件高效复用
- 使用helm应用级别的版本管理



Helm是一个Kubernetes的包管理工具，就像Linux下的包管理器，如yum/apt等，可以很方便的将之前打包好的yaml文件部署到kubernetes上。

Helm有三个重要概念

- helm：一个命令行客户端工具，主要用于Kubernetes应用chart的创建、打包、发布和管理
- Chart：应用描述，一系列用于描述k8s资源相关文件的集合，相当于把yaml打包，是yaml的集合
- Release：基于Chart的部署实体，一个chart被Helm运行后将会生成对应的release，将在K8S中创建出真实的运行资源对象。也就是应用级别的版本管理
- Repository：用于发布和存储Chart的仓库



Helm采用客户端/服务端架构，有如下组件组成

- Helm CLI是Helm客户端，可以在本地执行
- Tiller是服务器端组件，在Kubernetes集群上运行，并管理Kubernetes应用程序
- Repository是Chart仓库，Helm客户端通过HTTP协议来访问仓库中Chart索引文件和压缩包



2019年11月13日，Helm团队发布了Helm v3的第一个稳定版本

该版本主要变化如下

- 架构变化
  - 最明显的变化是Tiller的删除
  - V3版本**删除Tiller**
  - relesase可以在不同命名空间重用



通过 `helm repo add 仓库名 仓库地址` 添加 helm仓库

```sh
# 配置微软源
helm repo add stable http://mirror.azure.cn/kubernetes/charts
# 配置阿里源
helm repo add aliyun https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
# 配置google源
helm repo add google https://kubernetes-charts.storage.googleapis.com/

# 更新
helm repo update

# 查看配置的仓库
helm repo list
# 搜索
helm search repo stable
# 删除
helm repo remove google
```

基本命令：`chart install`、`chart upgrade`、`chart rollback`



### 使用helm快速部署应用

helm 将帮助我们创建 chart

```sh
# 搜索 weave
$ helm search repo weave
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                                       
aliyun/weave-cloud      0.1.2                           Weave Cloud is a add-on to Kubernetes which pro...
aliyun/weave-scope      0.9.2           1.6.5           A Helm chart for the Weave Scope cluster visual...
stable/weave-cloud      0.3.9           1.4.0           DEPRECATED - Weave Cloud is a add-on to Kuberne...
stable/weave-scope      1.1.12          1.12.0          DEPRECATED - A Helm chart for the Weave Scope c...

# 安装 weave-scope，WeaveScope是一个容器监控工具。注意 aliyun/weave-scope 因为版本问题无法安装
# 第三个参数指定安装后的名字，这里叫ui，因为weave-scope就是它的带ui版本
# 安装完成之后打印了很多信息，告诉我们通过使用kubectl端口转发，您现在应该能够在web浏览器中访问Scope前端，以及如何操作
$ helm install ui stable/weave-scope
NAME: ui
LAST DEPLOYED: Tue May  4 21:20:15 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
You should now be able to access the Scope frontend in your web browser, by
using kubectl port-forward:

kubectl -n default port-forward $(kubectl -n default get endpoints \
ui-weave-scope -o jsonpath='{.subsets[0].addresses[0].targetRef.name}') 8080:4040

then browsing to http://localhost:8080/.
For more details on using Weave Scope, see the Weave Scope documentation:

https://www.weave.works/docs/scope/latest/introducing/

# 查看已安装列表
$ helm list
NAME    NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                   APP VERSION
ui      default         1               2021-05-04 21:20:15.196488373 +0800 CST deployed        weave-scope-1.1.12      1.12.0
# 查看状态，实际上跟刚刚安装后打印的信息差不多
$ helm status ui
NAME: ui
LAST DEPLOYED: Tue May  4 21:20:15 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
You should now be able to access the Scope frontend in your web browser, by
using kubectl port-forward:

kubectl -n default port-forward $(kubectl -n default get endpoints \
ui-weave-scope -o jsonpath='{.subsets[0].addresses[0].targetRef.name}') 8080:4040

then browsing to http://localhost:8080/.
For more details on using Weave Scope, see the Weave Scope documentation:

https://www.weave.works/docs/scope/latest/introducing/

# 查看，可以发现 srv 是未暴露到任何外部端口
$ kubectl get pod,svc -o wide
NAME                                                READY   STATUS    RESTARTS   AGE     IP               NODE        NOMINATED NODE   READINESS GATES
pod/weave-scope-agent-ui-c7zxv                      1/1     Running   0          16m     192.168.31.111   k8snode01   <none>           <none>
pod/weave-scope-agent-ui-nvwzx                      1/1     Running   0          16m     192.168.31.101   k8smaster   <none>           <none>
pod/weave-scope-agent-ui-zswxc                      1/1     Running   0          16m     192.168.31.112   k8snode02   <none>           <none>
pod/weave-scope-cluster-agent-ui-7498b8d4f4-d46t6   1/1     Running   0          16m     10.244.2.11      k8snode02   <none>           <none>
pod/weave-scope-frontend-ui-649c7dcd5d-stnhh        1/1     Running   0          16m     10.244.1.10      k8snode01   <none>           <none>

NAME             TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
kubernetes       ClusterIP   10.96.0.1      <none>        443/TCP   20h
ui-weave-scope   ClusterIP   10.104.49.15   <none>        80/TCP    10m

# 对其yaml进行编辑，将spec.type改为 NodePort 以暴露端口。当然也可以用 helm status ui 查看时给我们提示的方法
$ kubectl edit svc ui-weave-scope
$ kubectl get svc -o wide
NAME                     TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE   SELECTOR
service/kubernetes       ClusterIP   10.96.0.1      <none>        443/TCP        20h   <none>
service/ui-weave-scope   NodePort    10.104.49.15   <none>        80:31668/TCP   16m   app=weave-scope,component=frontend,release=ui
```

然后通过 http://192.168.31.101:31668/ 可以访问，就可以看到所有容器可视化图谱和监控信息



### 自己创建chart

```sh
$ helm create mychart
Creating mychart
$ ls mychart/
charts  Chart.yaml  templates  values.yaml
```

- templates：编写yaml文件存放到这个目录
- values.yaml：存放的是全局的yaml文件
- chart.yaml：当前chart属性配置信息

 在templates文件夹创建两个文件。这里通过下面命令创建 deployment.yaml、service.yaml

```sh
$ kubectl create deployment nginx05 --image=nginx -o yaml > /mychart/templatesdeployment.yaml
$ kubectl expose deployment nginx05 --name=nginx05svc --port=80 --target-port=80 --type=NodePort --dry-run=client -o yaml > /mychart/templatesservice.yaml
$ kubectl delete deployment nginx05

# 通过 helm install 安装
$ helm install nginx05 mychart
NAME: nginx05
LAST DEPLOYED: Tue May  4 21:53:11 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=mychart,app.kubernetes.io/instance=nginx05" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:80
# 查看
$ kubectl get pod,svc -o wide
NAME                           READY   STATUS    RESTARTS   AGE     IP            NODE        NOMINATED NODE   READINESS GATES
pod/nginx05-6fb4dfcc86-zk4sg   1/1     Running   0          2m18s   10.244.2.12   k8snode02   <none>           <none>

NAME                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE     SELECTOR
service/kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP        20h     <none>
service/nginx05svc   NodePort    10.101.244.63    <none>        80:30033/TCP   2m18s   app=nginx05

# 如果我们修改了mychart中的内容，可以通过 helm upgrade 升级
$ helm upgrade nginx05 mychart
```

###  chart模板使用

通过传递参数，动态渲染模板，yaml内容动态从传入参数生成

刚刚我们创建mychart的时候，有values.yaml文件，这个文件就是一些全局的变量，然后在templates中能取到变量的值，下面我们可以利用这个，来完成动态模板

- 在values.yaml定义变量和值
- 具体yaml文件，获取定义变量值
- yaml文件中大题有几个地方不同
  - image
  - tag
  - label
  - port
  - replicas

在values.yaml定义变量和值，注意values.yaml原有许多内容，key不要重复了，这里我为了方便全部加上前缀防止key重复

```yaml
cat >> mychart/values.yaml << EOF

myreplicas: 1
myimage: nginx
mytag: 1.16
mylabel: nginx
myport: 80
EOF
```

获取变量值：通过表达式 使用全局变量 `{{ .Values.变量名称}}`。还有一个非常常用的写法就是`{{ .Release.Name}}`，将取到你当前的发布名称，即 `helm install 发布名 mychart` 中的第三个参数，这样就可以使用同一个chart部署多个k8s对象

修改deployment.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: {{ .Values.mylabel}}
  name: {{ .Release.Name}}-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Values.mylabel}}
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: {{ .Values.mylabel}}
    spec:
      containers:
      - image: {{ .Values.myimage}}
        name: {{ .Values.myimage}}
        resources: {}
status: {}
```

修改service.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: {{ .Values.mylabel}}
  name: {{ .Release.mylabel}}-svc
spec:
  ports:
  - port: {{ .Values.myport}}
    protocol: TCP
    targetPort: 80
  selector:
    app: {{ .Values.mylabel}}
  type: NodePort
status:
  loadBalancer: {}
```

```sh
# 尝试安装，但并不真正安装。可以看到通过变量生成的文件，是没问题的
$ helm install nginx06 mychart --dry-run 
NAME: nginx06
LAST DEPLOYED: Tue May  4 22:27:46 2021
NAMESPACE: default
STATUS: pending-install
REVISION: 1
HOOKS:
---
# Source: mychart/templates/tests/test-connection.yaml
apiVersion: v1
kind: Pod
metadata:
  name: "nginx06-mychart-test-connection"
  labels:

    helm.sh/chart: mychart-0.1.0
    app.kubernetes.io/name: mychart
    app.kubernetes.io/instance: nginx06
    app.kubernetes.io/version: "1.16.0"
    app.kubernetes.io/managed-by: Helm
  annotations:
    "helm.sh/hook": test-success
spec:
  containers:
    - name: wget
      image: busybox
      command: ['wget']
      args:  ['nginx06-mychart:80']
  restartPolicy: Never
MANIFEST:
---
# Source: mychart/templates/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: nginx06-mychart
  labels:

    helm.sh/chart: mychart-0.1.0
    app.kubernetes.io/name: mychart
    app.kubernetes.io/instance: nginx06
    app.kubernetes.io/version: "1.16.0"
    app.kubernetes.io/managed-by: Helm
---
# Source: mychart/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: nginx
  name: nginx06-svc
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx
  type: NodePort
status:
  loadBalancer: {}
---
# Source: mychart/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx
  name: nginx06-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
        resources: {}
status: {}

NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=mychart,app.kubernetes.io/instance=nginx06" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:80
```



## 持久化存储

之前我们有提到数据卷：`emptydir` ，是本地存储，pod重启，数据就不存在了，需要对数据持久化存储

对于数据持久化存储【pod重启，数据还存在】，有两种方式

- nfs：网络存储【通过一台服务器来存储】

### 网络存储

准备一台nfs服务器

```sh
# 安装nfs
yum install -f nfs-utils
# 创建存放数据的目录
mkdir -p /data/nfx
# 设置挂载路径：编辑/etc/exports，设置内容 /data/nfs *(rw,no_root_squash)
vim /etc/exports
```

node节点上设置，node01 & node02

```sh
# 安装执行完成后，会自动帮我们挂载上
yum install -y nfs-utils
```

回到nfs服务器

```sh
# 启动
systemctl start nfs
```

最后我们在k8s集群上部署应用，使用nfs持久化存储

```sh
mkdir pv
cd pv
```

通过 nfs-nginx.yaml 创建 Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-dep1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: wwwroot
          mountPath: /usr/share/nginx/html
        ports:
        - containerPort: 80
      volumes:
      - name: wwwroot
        nfs:
          server: 192.168.31.121
          path: /data/nfs
```

```sh
$ kubectl apply -f nfs-nginx.yaml
$ kubectl describe pod nginx-dep1
```

通过这个方式，就挂载到了刚刚我们的nfs数据节点下的 /data/nfs 目录

最后就变成了：容器中 /usr/share/nginx/html -> 192.168.44.134/data/nfs 内容是对应的

测试：在nfs服务器上写入数据，比如一个 index.html ，然后到容器中查看

```sh
# 进入容器中查看
$ kubectl exec -it nginx-dep1 bash
nginx-dep1$ ls /usr/share/nginx/html
index.html
```

###  PV & PVC

上述的方式 ip 和端口是直接放在容器上的，管理起来可能不方便。这里就需要用到 pv 和 pvc的概念了，方便我们配置和管理我们的 ip 地址等元信息

PV：持久化存储，对存储的资源进行抽象，对外提供可以调用的地方【生产者】

PVC：用于调用，不需要关心内部实现细节【消费者】

PV 和 PVC 使得 K8S 集群具备了存储的逻辑抽象能力。使得在配置Pod的逻辑里可以忽略对实际后台存储 技术的配置，而把这项配置的工作交给PV的配置者，即集群的管理者。存储的PV和PVC的这种关系，跟 计算的Node和Pod的关系是非常类似的；PV和Node是资源的提供者，根据集群的基础设施变化而变 化，由K8s集群管理员配置；而PVC和Pod是资源的使用者，根据业务服务的需求变化而变化，由K8s集 群的使用者即服务的管理员来配置。

- PVC绑定PV
- 定义PVC
- 定义PV【数据卷定义，指定数据存储服务器的ip、路径、容量和匹配模式】

pvc-pv.yaml

```yaml
apiVersion: apps/v1
kind: Deployment # 先创建一个 Depoloyment，将设置为 pvc 调用存储
metadata:
  name: nginx-dep2
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        volumeMounts:
        - name: wwwroot
          mountPath: /usr/share/nginx/html
        ports:
        - containerPort: 80
      volumes:
      - name: wwwroot
        persistentvolumeClaim:
          claimName: my-pvc
---
apiVersion: v1
kind: PersistentVolumeClaim # 创建一个 pcv 用于调用 pv
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 5ci
---
apiVersion: v1
kind: PersistentVolume # 创建一个pv，将绑定到刚刚的nfs服务器
metadata:
 name: my-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadwriteMany
  nfs:
    path: /k8s/nfs
    server: 192.168.31.121
```

```sh
$ kubectl apply -f pvc-pv.yaml
$ kubectl get pv,pvc -o wide
$ kubect exec -it nginx-dep1 bash
nginx-dep2$ ls /usr/share/nginx.html
index.html
```



## yaml声明式编程

### 格式

```yaml
# 键值
user:
  name: Tom
  age: 18
# 键值行内表示
user: {name: Tom, age: 18}

# 数组
usernames: 
- Tom
- Jack
# 数组行内表示
usernames: [Tom, Jack]
```

### Kubernetes yaml 结构

主要分为了两部分，一个是 控制器的定义 和 被控制的对象

控制器定义

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: default
spec:
  replicas: 3
  selector:
  matchLabels:app: nginx

```

被控制的对象：包含一些 镜像，版本、端口等

```yaml
  template:    metadata:      labels:        app: nginxl    spec:      containers:      - name: nginx        image: nginx:latest        ports:        - containerPort: 80
```

| 属性名称   | 介绍       |
| :--------- | :--------- |
| apiVersion | API版本    |
| kind       | 资源类型   |
| metadata   | 资源元数据 |
| spec       | 资源规格   |
| replicas   | 副本数量   |
| selector   | 标签选择器 |
| template   | Pod模板    |
| metadata   | Pod元数据  |
| spec       | Pod规格    |
| containers | 容器配置   |



### 快速编写yaml

一般来说很少自己手写YAML文件，因为这里面涉及到了很多内容，一般都会借助工具来创建，然后再做修改

#### kubectl create

这种方式一般用于资源没有部署的时候，我们可以直接创建一个YAML配置文件

```sh
# --dry-run=client尝试运行,并不会真正的创建镜像kubectl create deployment web --image=nginx -o yaml --dry-run=client
```

或者我们可以输出到一个文件中

```sh
kubectl create deployment web --image=nginx -o yaml --dry-run > myconfig/nginx.yaml
```

#### kubectl get

kubectl get 可以导出 yaml 文件

```sh
# 查看已部署的镜像
kubectl get deploy
# 导出
kubectl get deploy nginx -o=yaml --export > k8sdemoconfig/nginx.yaml
```



## Kubernetes集群资源监控

https://gitee.com/moxi159753/LearningNotes/blob/master/K8S/17_Kubernetes%E9%9B%86%E7%BE%A4%E8%B5%84%E6%BA%90%E7%9B%91%E6%8E%A7/README.md

### 概述

#### 监控指标

一个好的系统，主要监控以下内容

- 集群监控
  - 节点资源利用率
  - 节点数
  - 运行Pods
- Pod监控
  - 容器指标
  - 应用程序【程序占用多少CPU、内存】

#### 监控平台

使用普罗米修斯【prometheus】 + Grafana 搭建监控平台

- prometheus【定时搜索被监控服务的状态】
  - 开源的
  - 监控、报警、数据库
  - 以HTTP协议周期性抓取被监控组件状态
  - 不需要复杂的集成过程，使用http接口接入即可

- Grafana
  - 开源的数据分析和可视化工具
  - 支持多种数据源



### 部署prometheus

首先需要部署一个守护进程，node-exporter.yaml

```yaml
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: node-exporter
  namespace: kube-system
  labels:
    k8s-app: node-exporter
spec:
  selector:
    matchLabels:
      k8s-app: node-exporter
  template:
    metadata:
      labels:
        k8s-app: node-exporter
    spec:
      containers:
      - image: prom/node-exporter
        name: node-exporter
        ports:
        - containerPort: 9100
          protocol: TCP
          name: http
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: node-exporter
  name: node-exporter
  namespace: kube-system
spec:
  ports:
  - name: http
    port: 9100
    nodePort: 31672
    protocol: TCP
  type: NodePort
  selector:
    k8s-app: node-exporter
```

然后执行下面命令

```bash
kubectl create -f node-exporter.yaml
```

执行完，发现会报错这是。因为版本不一致的问题，因为发布的正式版本，而这个属于测试版本，所以我们找到第一行，然后把内容修改为如下所示

```bash
# 修改前
apiVersion: extensions/v1beta1
# 修改后 【正式版本发布后，测试版本不能使用】
apiVersion: apps/v1
```



然后通过yaml的方式部署prometheus

![image-20201120083107594](images/image-20201120083107594.png)

- configmap：定义一个configmap：存储一些配置文件【不加密】
- prometheus.deploy.yaml：部署一个deployment【包括端口号，资源限制】
- prometheus.svc.yaml：对外暴露的端口
- rbac-setup.yaml：分配一些角色的权限



下面我们进入目录下，首先部署 rbac-setup.yaml

```bash
kubectl create -f rbac-setup.yaml
```

![image-20201120090002150](images/image-20201120090002150.png)

然后分别部署

```bash
# 部署configmap
kubectl create -f configmap.yaml
# 部署deployment
kubectl create -f prometheus.deploy.yml
# 部署svc
kubectl create -f prometheus.svc.yml
```

部署完成后，我们使用下面命令查看

```bash
kubectl get pods -n kube-system
```

![image-20201120093213576](images/image-20201120093213576.png)

在我们部署完成后，即可看到 prometheus 的 pod了，然后通过下面命令，能够看到对应的端口

```bash
kubectl get svc -n kube-system
```

![image-20201121091348752](images/image-20201121091348752.png)

通过这个，我们可以看到 `prometheus` 对外暴露的端口为 30003，访问页面即可对应的图形化界面

```bash
http://192.168.177.130:30003
```

![image-20201121091508851](images/image-20201121091508851.png)

在上面我们部署完prometheus后，我们还需要来部署grafana

```bash
kubectl create -f grafana-deploy.yaml
```

然后执行完后，发现下面的问题

```bash
error: unable to recognize "grafana-deploy.yaml": no matches for kind "Deployment" in version "extensions/v1beta1"
```

我们需要修改如下内容

```bash
# 修改
apiVersion: apps/v1

# 添加selector
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grafana
      component: core
```

修改完成后，我们继续执行上述代码

```bash
# 创建deployment
kubectl create -f grafana-deploy.yaml
# 创建svc
kubectl create -f grafana-svc.yaml
# 创建 ing
kubectl create -f grafana-ing.yaml
```

我们能看到，我们的grafana正在

![image-20201120110426534](images/image-20201120110426534.png)

#### 配置数据源

下面我们需要开始打开 Grafana，然后配置数据源，导入数据显示模板

```bash
kubectl get svc -n kube-system
```

![image-20201120111949197](images/image-20201120111949197.png)

我们可以通过 ip + 30431 访问我们的 grafana 图形化页面

![image-20201120112048887](images/image-20201120112048887.png)

然后输入账号和密码：admin admin

进入后，我们就需要配置 prometheus 的数据源

![image-20201121092012018](images/image-20201121092012018.png)

 和 对应的IP【这里IP是我们的ClusterIP】

![image-20201121092053215](images/image-20201121092053215.png)

#### 设置显示数据的模板

选择Dashboard，导入我们的模板

![image-20201121092312118](images/image-20201121092312118.png)

然后输入 315 号模板

![image-20201121092418180](images/image-20201121092418180.png)

然后选择 prometheus数据源 mydb，导入即可

![image-20201121092443266](images/image-20201121092443266.png)

导入后的效果如下所示

![image-20201121092610154](images/image-20201121092610154.png)



## 双 master 集群搭建

Kubelete 不是那种一个 master 带几个 slave node，多组这样的 master 组合在一起成为集群。而是所有 master 作为一组，中间通过 LoadBalancer 组件调度所有 node。master 集对外有一个统一的VIP(虚拟ip)来对外进行访问

LoadBalancer 组件作用如下：

- 负载
- 检查master节点的状态

架构：

- keepalived：配置vip，检查节点的状态
- haproxy：负载均衡服务【类似于nginx】
- controller
- apiserver
- manager
- scheduler

详细步骤

- master01(192.168.31.101)：部署Keepalived、haproxy，初始化集群，安装Doker、网络插件
- master02(192.168.31.102)：部署Keepalived、haproxy，加入集群，安装Doker、网络插件
- node01(192.168.31.111)：加入集群，安装Doker、网络插件
- VIP：192.168.121



### 机器配置 & Docker、Kubeadm、kubectl 安装

跟前面集群准备一样，不再赘述



### keepalived 安装

master01 & master02：安装

```sh
# 安装相关工具
yum install -y conntrack-tools libseccomp libtool-ltdl
# 安装keepalived
yum install -y keepalived
```

master01：Keepalived 配置

```sh
cat > /etc/keepalived/keepalived.conf <<EOF 
! Configuration File for keepalived

global_defs {
   router_id k8s
}

vrrp_script check_haproxy {
    script "killall -0 haproxy"
    interval 3
    weight -2
    fall 10
    rise 2
}

vrrp_instance VI_1 {
    state MASTER 
    interface ens33 
    virtual_router_id 51
    priority 250
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass ceb1b3ec013d66163d6ab
    }
    virtual_ipaddress {
        192.168.31.101
    }
    track_script {
        check_haproxy
    }

}
EOF
```

master02：Keepalived 配置

```sh
cat > /etc/keepalived/keepalived.conf <<EOF 
! Configuration File for keepalived

global_defs {
   router_id k8s
}

vrrp_script check_haproxy {
    script "killall -0 haproxy"
    interval 3
    weight -2
    fall 10
    rise 2
}

vrrp_instance VI_1 {
    state BACKUP 
    interface ens33 
    virtual_router_id 51
    priority 200
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass ceb1b3ec013d66163d6ab
    }
    virtual_ipaddress {
        192.168.31.102
    }
    track_script {
        check_haproxy
    }

}
EOF
```

master01 & master02：Keepalived 启动

```sh
# 启动keepalived
systemctl start keepalived.service
# 设置开机启动
systemctl enable keepalived.service
# 查看启动状态
systemctl status keepalived.service
# 查看网卡信息，应该由两个ip了，/32后缀是vip
ip a s ens33
# 启动后，我们查看对应的端口是否包含 16443
netstat -tunlp | grep haproxy


```

master01 & master02：Haproxy 配置

```sh
# 写入 
cat > /etc/haproxy/haproxy.cfg << EOF
#---------------------------------------------------------------------
# Global settings
#---------------------------------------------------------------------
global
    # to have these messages end up in /var/log/haproxy.log you will
    # need to:
    # 1) configure syslog to accept network log events.  This is done
    #    by adding the '-r' option to the SYSLOGD_OPTIONS in
    #    /etc/sysconfig/syslog
    # 2) configure local2 events to go to the /var/log/haproxy.log
    #   file. A line like the following can be added to
    #   /etc/sysconfig/syslog
    #
    #    local2.*                       /var/log/haproxy.log
    #
    log         127.0.0.1 local2
    
    chroot      /var/lib/haproxy
    pidfile     /var/run/haproxy.pid
    maxconn     4000
    user        haproxy
    group       haproxy
    daemon 
       
    # turn on stats unix socket
    stats socket /var/lib/haproxy/stats
#---------------------------------------------------------------------
# common defaults that all the 'listen' and 'backend' sections will
# use if not designated in their block
#---------------------------------------------------------------------  
defaults
    mode                    http
    log                     global
    option                  httplog
    option                  dontlognull
    option http-server-close
    option forwardfor       except 127.0.0.0/8
    option                  redispatch
    retries                 3
    timeout http-request    10s
    timeout queue           1m
    timeout connect         10s
    timeout client          1m
    timeout server          1m
    timeout http-keep-alive 10s
    timeout check           10s
    maxconn                 3000
#---------------------------------------------------------------------
# kubernetes apiserver frontend which proxys to the backends
#--------------------------------------------------------------------- 
frontend kubernetes-apiserver
    mode                 tcp
    bind                 *:16443
    option               tcplog
    default_backend      kubernetes-apiserver    
#---------------------------------------------------------------------
# round robin balancing between the various backends
#---------------------------------------------------------------------
backend kubernetes-apiserver
    mode        tcp
    balance     roundrobin
    server      master01.k8s.io   192.168.44.155:6443 check
    server      master02.k8s.io   192.168.44.156:6443 check
#---------------------------------------------------------------------
# collection haproxy statistics message
#---------------------------------------------------------------------
listen stats
    bind                 *:1080
    stats auth           admin:awesomePassword
    stats refresh        5s
    stats realm          HAProxy\ Statistics
    stats uri            /admin?stats
EOF
```

###  部署Kubernetes Master

master01

```sh
# 创建文件夹
mkdir /usr/local/kubernetes/manifests -p
# 到manifests目录
cd /usr/local/kubernetes/manifests/
# 新建yaml文件
vi kubeadm-config.yaml
```

```yaml
apiServer:
  certSANs:
    - master1
    - master2
    - master.k8s.io
    - 192.168.44.158
    - 192.168.44.155
    - 192.168.44.156
    - 127.0.0.1
  extraArgs:
    authorization-mode: Node,RBAC
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta1
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controlPlaneEndpoint: "master.k8s.io:16443"
controllerManager: {}
dns: 
  type: CoreDNS
etcd:
  local:    
    dataDir: /var/lib/etcd
imageRepository: registry.aliyuncs.com/google_containers
kind: ClusterConfiguration
kubernetesVersion: v1.16.3
networking: 
  dnsDomain: cluster.local  
  podSubnet: 10.244.0.0/16
  serviceSubnet: 10.1.0.0/16
scheduler: {}
```

```sh
# 初始化集群
kubeadm init --config kubeadm-config.yaml

# 配置环境变量以使用kubectl
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
# 查看节点
kubectl get nodes
# 查看pod
kubectl get pods -n kube-system
# 查看集群状态
kubectl get cs
# 查看pod
kubectl get pods -n kube-system
# 安装flannel网络
kubectl apply -f kube-flannel.yml
# 查看
kubectl get pods -n kube-system
```

master02

```sh
# 从master01复制密钥及相关文件到master02

# ssh root@192.168.44.156 mkdir -p /etc/kubernetes/pki/etcd

# scp /etc/kubernetes/admin.conf root@192.168.44.156:/etc/kubernetes
   
# scp /etc/kubernetes/pki/{ca.*,sa.*,front-proxy-ca.*} root@192.168.44.156:/etc/kubernetes/pki
   
# scp /etc/kubernetes/pki/etcd/ca.* root@192.168.44.156:/etc/kubernetes/pki/etcd
```

```sh
# master02加入集群。--control-plane 只有在添加 master 节点的时候才有
kubeadm join master.k8s.io:16443 --token jv5z7n.3y1zi95p952y9p65 \
    --discovery-token-ca-cert-hash sha256:403bca185c2f3a4791685013499e7ce58f9848e2213e27194b75a2e3293d8812 \
    --control-plane 
    
# 查看集群状态
kubectl get cs
# 查看pod
kubectl get pods -n kube-system
```

node01

```sh
# 添加节点即可
kubeadm join master.k8s.io:16443 --token jv5z7n.3y1zi95p952y9p65 \
    --discovery-token-ca-cert-hash sha256:403bca185c2f3a4791685013499e7ce58f9848e2213e27194b75a2e3293d8812
```

### 测试

从集群中创建一个 nginx 测试，http://192.168.101/

```
# 创建nginx deployment
kubectl create deployment nginx --image=nginx
# 暴露端口
kubectl expose deployment nginx --port=80 --type=NodePort
# 查看状态
kubectl get pod,svc
```

# 待整理

- 特点:
  - 轻量级:消耗资源小
  - 开源
  - 弹性伸缩
  - 负载均衡
  - IPVS
- 梳理
  - Kubernetes：构建 K8S 集群
  - 资源清单：资源  掌握资源清单的语法  编写Pod  掌握Pod的生命周期***
  - Pod控制器：掌握各种控制器的特点以及使用定义方式
  - 服务发现：掌握SVC原理及其构建方式
  - 存储：掌握多种存储类型的特点。并且能够在不同环境中选择合适的存储方案(有自己的简介)
  - 调度器：掌握调度器原理能够 根据要求把Pod定义到想要的节点
  - 安全：集群的认证 鉴权 访问控制 原理及其流程
  - HELM: Linux yum，掌握HELM原理，HELM模板自定义，HELM部署一些常用插件
  - 运维：修改Kubeadm 达到证书可用期限为10年，能够构建高可用的Kubernetes集群
- 主要组件：
  - APISERVER: 所有服务访问统一入口
  - CrontrollerManager: 维持副本期望数目
  - Scheduler: 负责介绍任务，选择合适的节点进行分配任务
  - ETCD: 键值对数据库 储存K8S集群所有重要信息(持久化)
  - Kubelet:直接跟容器引擎交互实现容器的生命周期管理，
  - Kube-proxy:负责写入规则至IPTABLES、 IPVS实现服务映射访问的
- 其他插件
  - COREDNS: 可以为集群中的svC创建一 个域名IP的对应关系解析
  - DASHBOARD: 给K8S 集群提供个B/S 结构访问体系
  - INGRESS CONTROLLER: 官方只能实现四层代理，INGRESS 可以实现七层代理
  - FEDERATION: 提供-个可以跨集群中心多K8S统一 管理功能
  - PROMETHEUS: 提供K8S集群的监控能力
  - ELK: 提供K8S集群日志统一分析介入平台



## pod

- pod：1个pod中可以有1或多个容器，pod中运行这一个特殊容器pause，其它的则被称为业务容器，业务容器之间通过pause共享网络栈、数据卷等，同一pod中的不同业务容器上的服务互相访问就无需代理地址、映射端口等等，等于都在一个localhost上了，注意自然也不能用相同的端口了，这样通信更加高效
- pod类型
  - 自主式Pod
  - 控制器管理的Pod
    - ReplicationController&ReplicaSet&Deployment
      - ReplicationController用来确保容器应用的副本数始终保持在用户定义的副本数，即如果有容器异常退出，会自动创建新的Pod来替代;而如果异常多出来的容器也会自动回收。在新版本的Kubernetes中建议使用ReplicaSet来取代Repl icat ionControlle .
      - Repl icaSet跟ReplicationController没有本质的不同，只是名字不一-样，但是ReplicaSet支持集合式的selector
      - 虽然ReplicaSet 可以独立使用，但一般还是建议 使用Deployment 来自动管理ReplicaSet，这样就无需担心跟其他机制的不兼容问题(比如ReplicaSet 不支持rolling- update 但Deployment 支持)
    - HPA (Hori zontalPodAutoScale)
      - Hori zontal Pod Autoscaling仅适用于Deployment 和ReplicaSet ，在V1版本中仅支持根据Pod的CPU利用率扩所容，在v1alpha 版本中，支持根据内存和用户自定义的metric 扩缩容
    - StatefulSet：是为了解决有状态服务的问题(对应Deployments 和Repl icaSets是为无状态服务而设
      计)，其应用场景包括:
      - 稳定的持久化存储，即Pod重新调度后还是能访问到相同的持久化数据，基于PVC来实现
      - 稳定的网络标志，即Pod重新调度后其PodName 和HostName 不变，基于Headless Service(即没有Cluster IP的Service )来实现，
      - 有序部署，有序扩展，即Pod是有顺序的，在部署或者扩展的时候要依据定义的顺序依次依次进行(即从0到N-1,在下一个Pod运行之前所有之前的Pod必须都是Running 和Ready 状态)，基于init containers 来实现
      - 有序收缩，有序删除(即从N-1到0)
    - DaemonSet确保全部(或者一些) Node. 上运行一个Pod 的副本。当有Node 加入集群时，也会为他们新增一个Pod。当有Node 从集群移除时，这些Pod也会被回收。删除DaemonSet将会删除它创建的所有Pod。使用DaemonSet 的一些典型用法:
      - 运行集群存储daemon, 例如在每个Node. 上运行glusterd、 ceph。
      - 在每个Node.上运行 日志收集daemon, 例如fluentd、logstash。
      - 在每个Node.上运行监控daemon, 例如Prometheus Node Exporter
    - Job负责批处理任务，即仅执行一次的任务， 它保证批处理任务的一个或多 个Pod成功结束。Cron Job管理基于时间的Job, 即:
      - 在给定时间点只运行一次
      - 周期性地在给定时间点运行



- 网络通信方式：Kubernetes的网络模型假定了所有Pod都在一个何以直接连通的扁平的网络空间中，这在GCE (Google Compute Engine) 里面是现成的网络模型，Kubernetes 假定这个网络已经存在。而在私有云里搭建Kubernetes集群，就不能假定这个网络已经存在了。我们需要自己实现这个网络假设，将不同节点上的Docker 容器之间的互相访问先打通，然后运行Kubernetes

  - 同一个Pod内的多个容器之间:localhost（pause的网络工作栈）
  - 各Pod 之间的通讯: Overlay Network
  - Pod与Service之间的通讯：各节点的Iptables规则

- Flannel是CoreOS 团队针对Kubernetes 设计的一个网络规划服务，简单来说，它的功能是让集群中的不同节点主机创建的Docker容器都具有全集群唯一的虚拟IP地址。而且它还能在这些IP地址之间建立-一个覆盖网络(OverlayNetwork)，通过这个覆盖网络，将数据包原封不动地传递到目标容器内

  - ETCD之Flannel 提供说明:

    - 存储管理 Flannel可分配的IP地址段资 源
    - 监控ETCD中每个Pod的实际地址，并在内存中建立维护Pod 节点路由表

  - 流程：

    - 同-一个Pod内部通讯：同 一个Pod共享同一一个网络命名空间，共享同一一个Linux协议栈

    - Pod1至Pod2 ;

      Pod1与Pod2不在同- -台主机，Pod的地址是与docker0在同- 一个网段的，但docker0网段与宿主机网卡是两个完
      全不同的IP网段，并且不同Node之间的通信只能通过宿主机的物理网卡进行。将Pod的IP和所 在Node的IP关联起来，通过
      这个关联让Pod可以互相访问
      Pod1与Pod2在同一.台机器，由Docker0 网桥直接转发请求至Pod2， 不需要经过Flannel
      演示

    - Pod至Service的网络:目前基于性能考虑，全部为iptables 维护和转发

    - Pod到外网: Pod 向外网发送请求，查找路由表，转发数据包到宿主机的网卡，宿主网卡完成路由选择后，iptables执
      行Masquerade,把源IP更改为宿主网卡的IP， 然后向外网服务器发送请求

    - 外网访问Pod: Service

## Kubernetes

- Kubernetes是Google Omega的开源版本
- 据说Google的数据中心里运行着20多亿个容器，而且Google十年前就开始使用容器技术
- 最初Google开发了一个叫Borg的系统（现在命名为Omega）来调度如此庞大数量的容器和工作负载
- 之后通过Golang重写Borg即成就Kubernetes，并将其贡献到开源社区让全世界都能受益
- Kubernetes：编排容器，优化资源利用、高可用、滚动更新、网络插件、服务发现、监控、数据管理、日志管理等

### Hello

官方提供了一个交互式教程，通过Web浏览器就能使用预先部署好的一个Kubernetes集群，快速体验Kubernetes的功能和应用场景：https://kubernetes.io/zh/docs/tutorials/kubernetes-basics/

```shell
#集群命令。minikube是官方提供的一个迷你kubernetes集群
$ minikube version
$ minikube start #启动集群
#客户端命令。kubectl是官方提供的客户端
$ kubectl version
$ kubectl cluster-info #集群信息。如：地址...
$ kubectl get nodes

$ kubectl run kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1 --port=8080 #部署应用，需要提供部署名称(kubernetes-bootcamp)和应用程序映像位置（--image指定，包括Docker Hub外部托管映像的完整存储库URL）。
$ kubectl get deployments #列出部署。内容：名字、副本数...
$ kubectl get pods #列出pods
$ kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080 #默认情况下，所有Pod只能在集群内部访问，要访问上面这个应用只能直接访问容器的8080端口，为了能够从外部访问应用，需要将容器的8080端口映射到节点的端口
$ kubectl get services #列出服务。service是对外提供的服务？可以查看到应用被映射到节点的哪个端口
$ curl 地址:32253 #service的端口是随机分配的
$ kubectl scale deployments/kubernetes-bootcamp --replicas=3 #默认情况下应用只会运行一个副本，这里将副本数增加到3个
$ kubectl get pods
$ curl 地址:32253 #将负载均衡轮询处理请求
$ kubectl scale deployments/kubernetes-bootcamp --replicas=2 #同样的方式，将有一个副本被终止
$ kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=jocatalin/kubernetes-bootcamp:v2 #升级image的版本
$ kubectl get pods #s可以观察滚动更新的过程：v1的Pod被逐个 删除，同时启动了新的v2 Pod
$ kubectl rollout undo deployments/kubernetes-bootcamp #回退版本
```



### 概念

- **Cluster**：Cluster是计算、存储和网络资源的集合，Kubernetes利用这些资源运行各种基于容器的应用
- **Master**：Master是Cluster的大脑，它的主要职责是调度，即决定将应用放在哪里运行。Master运行Linux操作系统，可以是物理机或者虚拟机。为了实现高可用，可以运行多个Master
- **Node**：Node的职责是运行容器应用。Node由Master管理，Node负责监控并汇报容器的状态，同时根据Master的要求管理容器的生命周期。 Node运行在Linux操作系统上，可以是物理机或者是虚拟机
- **Pod**：Kubernetes的最小工作单元。每个Pod包含一个或多个容器。Pod中的容器会作为一个整体被Master调度到一个Node上运行
  - Pod被引入的两个主要目的
    - 可管理性：有些容器天生就是需要紧密联系，一起工作。Pod提供了比容器 更高层次的抽象，将它们封装到一个部署单元中。Kubernetes以Pod为最小单位进行调度、扩展、共享资源、管理生命周期
    - 通信和资源共享：Pod中的所有容器使用同一个网络namespace，即相同的IP地址和Port空间。它们可以直接用localhost通信。同样的，这些容器可以共 享存储，当Kubernetes挂载volume到Pod，本质上是将volume挂载到Pod中的每一个容器
  - Pods的两种使用方式
    - 运行单一容器：one-container-per-Pod是Kubernetes最常见的模型，这种情况下，只是将单个容器简单封装成Pod。即便是只有一个容器，Kubernetes管理的也是Pod而不是直接管理容器
    - 运行多个容器：问题在于哪些容器应该放到一个Pod中？这些容器联系必须非常紧密，而且需要直接共享资源。比如Pod包含两个容器，一个是File Puller，一个是Web Server，File Puller会定期从外部的Content Manager中拉取最新的文件，将其存放在共享的volume中。Web Server从volume读取文件，响应Consumer的请求，这两个容器是紧密协作的，它们一起为Consumer提供最新的数据，同时它们也通过volume共享数据，所以放到一个Pod是合适的；是否需要将Tomcat和MySQL放到一个Pod中？Tomcat从MySQL读取数据，它们之间需要协作，但还不至于需要放到一个Pod中一起部署、一起启动、一起停止。同时它们之间是通过JDBC交换数据，并不是直接共享存储，所以放到各自的Pod中更合适
- **Controller**：Kubernetes通常不会直接创建Pod，而是通过Controller来管理Pod的。Controller中定义了Pod的部署特性，比如有几个副本、在什么样的Node上运行等。为了满足不同的业务场景，Kubernetes提供了多种Controller，包括Deployment、ReplicaSet、DaemonSet、StatefuleSet、Job等
  - **Deployment**：最常用的Controller，比如在线教程中就是通过创建Deployment来部署应用的。Deployment可以管理Pod的多个副本，并确保Pod按照期望的状态运行
  - **ReplicaSet**：实现了Pod的多副本管理。使用Deployment时会自动创建ReplicaSet，也就是说Deployment是通过ReplicaSet来管理Pod的多个副本的，我们通常不需要直接使用ReplicaSet
  - **DaemonSet**：用于每个Node最多只运行一个Pod副本的场景。正如其名称所揭示的，DaemonSet通常用于运行daemon
  - **StatefuleSet**：能够保证Pod的每个副本在整个生命周期中名称 是不变的，而其他Controller不提供这个功能。当某个Pod发生故障需要删除并重新启动时，Pod的名称会发生变化，同时StatefuleSet会保证副本按照固定的顺序启动、更新或者删除
  - **Job**：用于运行结束就删除的应用，而其他Controller中的Pod通常是长期持续运行
- **Service**：Deployment可以部署多个副本，每个Pod都有自己的IP，外界如何访问这些副本呢？通过Pod的IP吗？要知道Pod很可能会被频繁地销毁和重启，它们的IP会发生变化，用IP来访问不太现实。Kubernetes Service定义了外界访问一组特定Pod的方式。Service有自己的IP和端口，Service为Pod提供了负载均衡。Kubernetes运行容器（Pod）与访问容器（Pod）这两项任务分别由Controller和Service执行
- **Namespace**：如果有多个用户或项目组使用同一个Kubernetes Cluster，如何将 他们创建的Controller、Pod等资源分开呢？Namespace可以将一个物理的Cluster逻辑上划分成多个虚拟 Cluster，每个Cluster就是一个Namespace。不同Namespace里的资源是完全隔离的。Kubernetes默认创建了两个Namespace，通过kubectl get namespace列出
  - default：创建资源时如果不指定，将被放到这个Namespace中
  - kube-system：Kubernetes自己创建的系统资源将放到这个Namespace中

### Cluster

- 部署三个节点的Kubernetes Cluster。k8s-master是Master，k8s-node1和k8s-node2是Node
- 所有节点都需要安装Docker、kubelet、kubeadm、kubectl：https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
  - kubelet：。运行在Cluster所有节点上，负责启动Pod和容器
  - kubeadm：用来初始化集群的指令
  - kubectl是Kubernetes命令行工具。通过kubectl可以部署和管理应用，查看各种资源，创建、删除和更新各种组件



```shell
#配置kubernetes源
$ cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

$ setenforce 0 #关闭SELinux（防火墙），临时生效
$ cp /etc/selinux/config /etc/selinux/config.bk
$ sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config #关闭SELinux，永久生效：将SELinux设置为permissive模式，相当于将其禁用
$ getenforce #查看selinux状态

$ yum install kubelet kubeadm kubectl -y --disableexcludes=kubernetes

$ systemctl enable --now kubelet #设置开机启动，并现在立即启动
```



```shell
kubeadm init --apiserver-advertise-address 192.168.56.105 --pod-network-cidr=10.244.0.0/16
```

