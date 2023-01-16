# 存储

## Volume



## PV & PVC

> PV：[PersistentVolume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)，持久卷。
>
> PVC：[PersistentVolumeClaim](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims)。
>
> 用法参考[Pod配置PVC存储](https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/)。

### 本地

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nginx-pv
  labels:
    app: nginx # pcv label
spec:
  storageClassName: nginx-storage # pv & pvc 关联字段
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data/nginx" # 本地路径
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nginx-pvc
spec:
  selector:
    matchLabels:
      app: nginx # pv 匹配 pcv label
  storageClassName: nginx-storage # pv & pvc 关联字段
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  volumes:
    - name: nginx-volume
      persistentVolumeClaim:
        claimName: nginx-pvc # volume 关联 pvc
  containers:
    - name: nginx
      image: nginx:latest
      ports:
        - name: http-nginx
          containerPort: 80
      volumeMounts:
        - name: nginx-volume
          mountPath: "/usr/share/nginx"
```

### 网络

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nginx-pv
  labels:
    app: nginx
spec:
  storageClassName: nginx
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  nfs:
    server: 192.168.x.x   # nfs 服务器地址
    path: /nfs/nginx     # nfs 服务器目录
```

### 限制 PVC 容量

> 参考[限制存储使用量](https://kubernetes.io/zh-cn/docs/tasks/administer-cluster/limit-storage-consumption/)。