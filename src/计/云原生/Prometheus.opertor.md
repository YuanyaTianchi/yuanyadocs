

+++

title = "prometheus.operator"
description = "prometheus.operator"
tags = ["it", "cloud", "prometheus"]

+++



# prometheus.operator



## Prometheus 管理

[参考](https://yunlzheng.gitbook.io/prometheus-book/part-iii-prometheus-shi-zhan/operator/what-is-prometheus-operator)

- Prometheus：声明式创建和管理Prometheus Server实例；
- ServiceMonitor：负责声明式的管理监控配置；
- PrometheusRule：负责声明式的管理告警配置；
- Alertmanager：声明式的创建和管理Alertmanager实例。

### CRD创建

通过 [bundle.yaml](https://github.com/prometheus-operator/prometheus-operator/blob/v0.52.1/bundle.yaml) 创建 Prometheus、ServiceMonitor、PrometheusRule、Alertmanager 等 CRD，也会部署所需的 kubernetes RBAC 资源

```shell
# 从 GitHub 获取部署 yaml，copy 一份并替换其中 namespace
git clone git@github.com:prometheus-operator/prometheus-operator.git
cp prometheus-operator/bundle.yaml prometheus-operator-bundle.yaml
sed -i 's/namespace: default/namespace: monitoring/g' prometheus-operator-bundle.yaml
# 检查
$ cat prometheus-operator-bundle.yaml | grep 'namespace: monitoring'
  namespace: monitoring
  namespace: monitoring
  namespace: monitoring
  namespace: monitoring

# 创建 monitoring 命名空间并部署 prometheus-operator
$ kubectl create namespace monitoring
$ kubectl -n monitoring apply -f prometheus-operator-bundle.yaml
The CustomResourceDefinition "prometheuses.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes
# 发现有一行错误，因为 annotations 太长导致 CRD 创建失败，参考 https://github.com/prometheus-operator/prometheus-operator/issues/4355 再次通过 kubectl create 创建资源
$ kubectl create -n monitoring -f prometheus-operator-bundle.yaml
customresourcedefinition.apiextensions.k8s.io/prometheuses.monitoring.coreos.com created
# 检查
$ kubectl get pod -n monitoring
NAME                                   READY   STATUS    RESTARTS   AGE
prometheus-operator-56d5887898-ghn54   1/1     Running   0          9m15s
```

### Prometheus初窥

创建 Prometheus 实例

```yaml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: inst
  namespace: monitoring
spec:
  resources:
    requests:
      memory: 400Mi
```

```shell
kubectl create -f prometheus-inst.yaml
# 查看
kubectl -n monitoring get statefulsets
kubectl -n monitoring get pods
# 通过 ort-forward 暴露 Prometheus 实例中的服务端口
kubectl -n monitoring port-forward statefulsets/prometheus-inst 9090:9090
```

访问 http://localhost:9090/，http://localhost:9090/config 查看 Prometheus 页面。



### example-app

参照官方实例 [example-app](https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/getting-started.md#related-resources) 创建 ServiceMonitor

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: example-app
  template:
    metadata:
      labels:
        app: example-app
    spec:
      containers:
      - name: example-app
        image: fabxc/instrumented_app
        ports:
        - name: web
          containerPort: 8080
---
kind: Service
apiVersion: v1
metadata:
  name: example-app
  labels:
    app: example-app
spec:
  selector:
    app: example-app
  ports:
  - name: web
    port: 8080
---

```

```shell
kubectl apply -f example-app.yaml
# 暴露端口
kubectl port-forward deployments/example-app 8080:8080
```

查看指标：http://localhost:8080/metrics

### ServiceMonitor

默认情况下 ServiceMonitor 和监控对象必须是在相同 Namespace 下，通过 namespaceSelector 使其可以跨命名空间

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: example-app
  namespace: monitoring
  labels:
    team: frontend
spec:
  namespaceSelector:
    matchNames:
    - default
  selector:
    matchLabels:
      app: example-app
  endpoints:
  - port: web
```

如果希望ServiceMonitor可以关联任意命名空间下的标签，则通过以下方式定义

```yaml
spec:
  namespaceSelector:
    any: true
```

```shell
kubectl create -f example-app-service-monitor.yaml
kubectl get ServiceMonitor -n monitoring
```

#### basicAuth

如果监控的Target对象启用了BasicAuth认证，那在定义ServiceMonitor对象时，可以使用endpoints配置中定义basicAuth如下所示

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: example-app
  namespace: monitoring
  labels:
    team: frontend
spec:
  namespaceSelector:
    matchNames:
    - default
  selector:
    matchLabels:
      app: example-app
  endpoints:
  - basicAuth:
      password:
        name: basic-auth
        key: password
      username:
        name: basic-auth
        key: user
    port: web
```

其中basicAuth中关联了名为basic-auth的Secret对象，用户需要手动将认证信息保存到Secret中

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: basic-auth
data:
  password: dG9vcg== # base64编码后的密码
  user: YWRtaW4= # base64编码后的用户名
type: Opaque
```



### Prometheus 关联 ServiceMonitor

Prometheus 使用 serviceMonitorSelector [关联 ServiceMonitor](https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/getting-started.md#include-servicemonitors)，

```yaml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: inst
  namespace: monitoring
spec:
  serviceMonitorSelector:
    matchLabels:
      team: frontend
  resources:
    requests:
      memory: 400Mi
```

将对 Prometheus 的变更应用到集群中

```shell
kubectl -n monitoring apply -f prometheus-inst.yaml
kubectl -n monitoring port-forward statefulsets/prometheus-inst 9090:9090
```

查看：http://localhost:9090/config，可以发现 Prometheus 实例配置中自动包含了一条名为monitoring/example-app/0的Job配置

```yaml
scrape_configs:
- job_name: serviceMonitor/monitoring/example-app/0
  honor_timestamps: true
  scrape_interval: 30s
  scrape_timeout: 10s
  metrics_path: /metrics
  scheme: http
  follow_redirects: true
  relabel_configs:
  #...
```

虽然配置有了，但是 http://localhost:9090/targets 并没包含任何的监控对象，查看 Prometheus 的 Pod 实例日志，可以看到如下信息

```shell
level=error component=k8s_client_runtime func=ErrorDepth msg="pkg/mod/k8s.io/client-go@v0.22.4/tools/cache/reflector.go:167: Failed to watch *v1.Pod: failed to list *v1.Pod: pods is forbidden: User \"system:serviceaccount:monitoring:default\" cannot list resource \"pods\" in API group \"\" in the namespace \"default\""
```

### ServiceAccount

默认创建的Prometheus实例使用的是monitoring命名空间下的default账号，该账号并没有权限能够获取default命名空间下的任何资源信息。因此需要在 Monitoring 命名空间下为创建一个 [ServiceAccount](https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/getting-started.md#enable-rbac-rules-for-prometheus-pods)，并且为该账号赋予相应的集群访问权限。

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
  namespace: monitoring
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
  - nodes
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources:
  - configmaps
  verbs: ["get"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: prometheus
  namespace: monitoring
```



### Prometheus

```yaml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: inst
  namespace: monitoring
spec:
  serviceAccountName: prometheus # 关联 ServiceAccount
  serviceMonitorSelector: # 关联 ServiceMonitor
    matchLabels:
      team: frontend
  resources:
    requests:
      memory: 400Mi
```

```shell
kubectl apply -f prometheus-inst.yaml
kubectl -n monitoring port-forward statefulsets/prometheus-inst 9090:9090
```

查看：http://localhost:9090/targets。



## PrometheusRule 管理

Prometheus 原生管理方式需要手动创建 Prometheus 的告警文件，并通过在 Prometheus 配置中声明式的加载

Prometheus Operator 模式中，告警规则变成一个通过 Kubernetes API 声明式创建的资源 PrometheusRule

官方[例子](https://github.com/prometheus-operator/prometheus-operator/blob/v0.52.1/example/thanos/prometheus-rule.yaml)如下

```yaml
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
  labels:
    prometheus: example
    role: alert-rules
  name: prometheus-example-rules
spec:
  groups:
  - name: ./example.rules
    rules:
    - alert: ExampleAlert
      expr: vector(1)
```

```shell
kubectl -n monitoring create -f example-rule.yaml
```

通过在Prometheus中使用ruleSelector通过选择需要关联的PrometheusRule即可

```yaml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: inst
  namespace: monitoring
spec:
  serviceAccountName: prometheus # 关联 ServiceAccount
  serviceMonitorSelector: # 关联 ServiceMonitor
    matchLabels:
      team: frontend
  ruleSelector: # 关联 PrometheusRule
    matchLabels:
      role: alert-rules
      prometheus: example
  resources:
    requests:
      memory: 400Mi
```

```shell
kubectl apply -f prometheus-inst.yaml
kubectl -n monitoring port-forward statefulsets/prometheus-inst 9090:9090
```

查看告警规则：http://localhost:9090/rules。



## Alertmanager 管理



```yaml
apiVersion: monitoring.coreos.com/v1
kind: Alertmanager
metadata:
  name: inst
  namespace: monitoring
spec:
  replicas: 2
```

当replicas大于1时，Prometheus Operator会自动通过集群的方式创建Alertmanager

```shell
kubectl -n monitoring create -f alertmanager-inst.yaml
```



（此间不需要也可以访问，直接就是 Runing 状态）此时 Alertmanager 的 Pod 实例会一直处于 ContainerCreating 的状态，因为 Prometheus Operator 通过 Statefulset 的方式创建的 Alertmanager 实例，在默认情况下，会通过 `alertmanager-{ALERTMANAGER_NAME}` 的命名规则去查找 Secret 配置并以文件挂载的方式，将 Secret 的内容作为配置文件挂载到 Alertmanager 实例当中。因此，这里还需要为 Alertmanager 创建相应的配置内容

```yaml
global:
  resolve_timeout: 5m
route:
  group_by: ['job']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 12h
  receiver: 'webhook'
receivers:
- name: 'webhook'
  webhook_configs:
  - url: 'http://alertmanagerwh:30500/'
```

保存为 alertmanager.yaml 并创建 Secret

```shell
kubectl -n monitoring create secret generic alertmanager-inst --from-file=alertmanager.yaml
```

（此间不需要也可以访问，直接就是 Runing 状态）



```shell
kubectl -n monitoring port-forward statefulsets/alertmanager-inst 9093:9093
```

查看集群状态：http://localhost:9093/#/status



Prometheus 绑定 Alertmanager

```yaml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: inst
  namespace: monitoring
spec:
  serviceAccountName: prometheus # 关联 ServiceAccount
  serviceMonitorSelector: # 关联 ServiceMonitor
    matchLabels:
      team: frontend
  ruleSelector: # 关联 PrometheusRule
    matchLabels:
      role: alert-rules
      prometheus: example
  alerting: # 关联 Alertmanager
    alertmanagers:
    - name: alertmanager-example
      namespace: monitoring
      port: web
  resources:
    requests:
      memory: 400Mi
```

```shell
kubectl apply -f prometheus-inst.yaml
kubectl -n monitoring port-forward statefulsets/prometheus-inst 9090:9090
```

查看配置：http://localhost:9090/rules。可以看到如下信息

```yaml
alerting:
  alert_relabel_configs:
  - separator: ;
    regex: prometheus_replica
    replacement: $1
    action: labeldrop
  alertmanagers:
  - follow_redirects: true
    scheme: http
    path_prefix: /
    timeout: 10s
    api_version: v2
    relabel_configs:
```

## 自定义配置

https://yunlzheng.gitbook.io/prometheus-book/part-iii-prometheus-shi-zhan/operator/use-custom-configuration-in-operator

对于用户而言，可能还是希望能够手动管理Prometheus配置文件，而非通过Prometheus Operator自动完成，因为Prometheus Operator对于Job的配置只适用于在Kubernetes中部署和管理的应用程序，如果希望使用Prometheus监控一些其他的资源，例如AWS或者其他平台中的基础设施或者应用，这些并不在Prometheus Operator的能力范围之内，就需要手动管理 Prometheus 配置

为了能够在通过Prometheus Operator创建的Prometheus实例中使用自定义配置文件，我们只能创建一个不包含任何与配置文件内容相关的Prometheus实例

```yaml
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: inst-cc
  namespace: monitoring
spec:
  serviceAccountName: prometheus
  resources:
    requests:
      memory: 400Mi
```

将以上内容保存到prometheus-inst-cc.yaml文件中，并且通过kubectl创建:

```shell
kubectl -n monitoring create -f prometheus-inst-cc.yaml
```

如果查看新建Prometheus的Pod实例YAML定义，我们可以看到Pod中会包含一个volume配置：

```yaml
volumes:
  - name: config
    secret:
      defaultMode: 420
      secretName: prometheus-inst-cc
```

Prometheus的配置文件实际上是保存在名为`prometheus-<name-of-prometheus-object>`的Secret中，当用户创建的Prometheus中关联ServiceMonitor这类会影响配置文件内容的定义时，Promethues Operator会自动管理。而如果Prometheus定义中不包含任何与配置相关的定义，那么Secret的管理权限就落到了用户自己手中。

通过修改prometheus-inst-cc的内容，从而可以让用户可以使用自定义的Prometheus配置文件，作为示例，我们创建一个prometheus.yaml文件并添加以下内容：

```yaml
global:
  scrape_interval: 10s
  scrape_timeout: 10s
  evaluation_interval: 10s
```

生成文件内容的base64编码后的内容：

```shell
$ cat prometheus.yaml | base64
Z2xvYmFsOgogIHNjcmFwZV9pbnRlcnZhbDogMTBzCiAgc2NyYXBlX3RpbWVvdXQ6IDEwcwogIGV2YWx1YXRpb25faW50ZXJ2YWw6IDEwcw==
```

修改名为prometheus-inst-cc的Secret内容，如下所示：

```shell
$ kubectl -n monitoring edit secret prometheus-inst-cc
# ...
data:
  prometheus.yaml: "Z2xvYmFsOgogIHNjcmFwZV9pbnRlcnZhbDogMTBzCiAgc2NyYXBlX3RpbWVvdXQ6IDEwcwogIGV2YWx1YXRpb25faW50ZXJ2YWw6IDEwcw=="
```

通过port-forward在本地访问新建的Prometheus实例，观察配置文件变化即可：http://localhost:9091/config

```shell
kubectl -n monitoring port-forward statefulsets/prometheus-inst-cc 9091:9090
```



