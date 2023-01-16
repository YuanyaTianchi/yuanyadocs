

+++

title = "prometheus"
description = "prometheus"
tags = ["it", "cloud", "prometheus"]

+++



# prometheus



## cAdvisor

Kubelet 内置 cAdvisor，无需单独安装，可通过 https://127.0.0.1:10250/metrics/cadvisor 访问

```shell
# 创建 ServiceAccount
kubectl create serviceaccount cadvisor-sa -n monitoring
# 创建 ClusterRoleBinding 并绑定 ClusterRole 到 ServiceAccount，这里直接使用 cluster-admin
kubectl create clusterrolebinding cadvisor-crb --clusterrole=cluster-admin --serviceaccount=monitoring:cadvisor-sa
# 查看 secret
kubectl get secret -n monitoring
kubectl describe secret cadvisor-sa-token-<xxxxx> -n monitoring
# 通过 http 访问
curl https://127.0.0.1:10250/metrics/cadvisor -k -H "Authorization: Bearer <token>"
```



## Grafana

> [在 Kubernetes 上部署 Grafana](https://grafana.com/docs/grafana/latest/setup-grafana/installation/kubernetes/)。
>
> 
>
> [Kubernetes 部署图表工具 Grafana](http://www.mydlq.club/article/111/)。



```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: grafana-pv
  labels:
    app: grafana
spec:
  storageClassName: grafana-storage
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data/grafana"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: grafana-pvc
spec:
  selector:
    matchLabels:
      app: grafana
  storageClassName: grafana-storage
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
---
apiVersion: v1
kind: Service
metadata:
  name: grafana-svc
  labels:
    app: grafana
spec:
  type: NodePort
  ports:
  - name: http-grafana
    port: 3000
    nodePort: 30300
    targetPort: http-grafana
  selector:
    app: grafana
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grafana
  labels:
    app: grafana
spec:
  selector:
    matchLabels:
      app: grafana
  template:
    metadata:
      labels:
        app: grafana
    spec:
      securityContext:
        fsGroup: 472
        supplementalGroups:
          - 0
      initContainers: # 初始化容器，用于修改挂载的存储的文件夹归属组与归属用户
        - name: init-file
          image: busybox:latest
          imagePullPolicy: IfNotPresent
          securityContext:
            runAsUser: 0
          command: ['chown', '-R', "472:0", "/var/lib/grafana"]
          volumeMounts:
            - name: grafana-volume
              mountPath: "/var/lib/grafana"
              subPath: grafana-subpath
      containers:
        - name: grafana
          image: grafana/grafana:latest
          imagePullPolicy: IfNotPresent
          ports:
            - name: http-grafana
              containerPort: 3000
          readinessProbe: # 就绪探针
            failureThreshold: 3
            httpGet:
              path: /robots.txt
              port: 3000
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 2
          livenessProbe: # 存活探针
            failureThreshold: 3
            initialDelaySeconds: 30
            periodSeconds: 10
            successThreshold: 1
            tcpSocket:
              port: 3000
            timeoutSeconds: 1
          env: # 配置环境变量，设置 Grafana 的默认管理员用户名/密码
            - name: GF_SECURITY_ADMIN_USER
              value: "admin"
            - name: GF_SECURITY_ADMIN_PASSWORD
              value: "admin"
          resources:
            requests:
              cpu: 250m
              memory: 750Mi
          securityContext: # 容器安全策略，设置运行容器使用的归属组与用户
            runAsUser: 472
          volumeMounts:
            - name: grafana-volume
              mountPath: "/var/lib/grafana"
              subPath: grafana-subpath
      volumes:
        - name: grafana-volume
          persistentVolumeClaim:
            claimName: grafana-pvc
```

