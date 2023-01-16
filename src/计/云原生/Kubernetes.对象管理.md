

+++

title = "Kubernetes.对象管理"
description = "Kubernetes.对象管理"
tags = ["it", "cloud"]

+++



# Kubernetes.对象管理

## kustomize

> [kustomize](https://github.com/kubernetes-sigs/kustomize)：kubectl 1.14 官方支持的声明式对象管理工具，可以[使用 Kustomize 对 Kubernetes 对象进行声明式管理](https://kubernetes.io/zh-cn/docs/tasks/manage-kubernetes-objects/kustomization/)，详细用法和用例见 [examples](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)。



### ConfigMap 生成器

> Kustomize 可以从 [ConfigMapGenerator](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/configGeneration.md) 生成 ConfigMap 对象，并通过 [Generator Options](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/generatorOptions.md) 设置 labels 和 annotations。



因为 config & kustomize 一般为不同目录，所以在 config 目录下添加 kustomization.yaml 以生成 ConfigMap 资源，由 kustomize 目录下的 kustomization.yaml 引用即可

```shell
$ tree
.
├── configs
│   ├── config.yaml
│   └── kustomization.yaml
└── kustomize
    └── base
        └── kustomization.yaml
```

```yaml
# configs/config.yaml
author:
  name: yuanya
```

```yaml
# configs/kustomization.yaml
configMapGenerator:
  - name: yuanya-cm
    # 直接设置
    literals:
      - foo=bar
    # 指定文件：文件内容作 value，缺省 key 则文件名作 key
    files:
      - config=config.yaml
      - config.yaml
generatorOptions:
  # 禁用将内容哈希后缀附加到生成资源的名称。不禁用则名称如`name: yuanya-cm-h69799k82t`
  disableNameSuffixHash: true
  # 为生成的资源添加 labels
  labels:
    kustomize.generated.resource: somevalue
  # 为生成的资源添加 annotations
  annotations:
    annotations.only.for.generated: othervalue
```

```shell
$ kubectl kustomize configs/
apiVersion: v1
data:
  config: |
    author:
      name: yuanya
  config.yaml: |
    author:
      name: yuanya
  foo: bar
kind: ConfigMap
metadata:
  annotations:
    annotations.only.for.generated: othervalue
  labels:
    kustomize.generated.resource: somevalue
  name: yuanya-cm
```

```yaml
# kustomize/base/kustomization.yaml
resources:
  - ../../configs
```

```shell
$ kubectl kustomize kustomize/base/
apiVersion: v1
data:
  config: |
    author:
      name: yuanya
  config.yaml: |
    author:
      name: yuanya
  foo: bar
kind: ConfigMap
metadata:
  annotations:
    annotations.only.for.generated: othervalue
  labels:
    kustomize.generated.resource: somevalue
  name: yuanya-cm
```



### ConfigMap 合并 & 替代

> Kustomize 可以通过 [ConfigMap 合并 & 替代](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/combineConfigs.md#a-mixin-approach-to-management)，以实现配置的混合管理。

```shell
$ tree
.
├── configs
│   ├── config.yaml
│   └── kustomization.yaml
└── kustomize
    ├── base
    │   └── kustomization.yaml
    └── overlays
        └── dev
            └── kustomization.yaml
```

```yaml
# kustomize/overlays/dev/kustomization.yaml
resources:
  - ../../base
namePrefix: dev-
configMapGenerator:
  - name: yuanya-cm
    # merge 将合并 ConfigMap，存在重复则仅替代该重复声明；replace 将完全替代原 ConfigMap。
    behavior: merge
    literals:
      - baz=qux
generatorOptions:
  annotations:
    annotations.only.for.generated: basevalue
```

```shell
$ kubectl kustomize kustomize/overlays/dev/
apiVersion: v1
data:
  baz: qux # 新增
  config: |
    author:
      name: yuanya
  config.yaml: |
    author:
      name: yuanya
  foo: bar
kind: ConfigMap
metadata:
  annotations:
    annotations.only.for.generated: basevalue # 重复替代
  labels:
    kustomize.generated.resource: somevalue
  name: dev-yuanya-cm # dev-前缀
```

### ConfigMap 补丁

> kustomize 通过 [patchesStrategicMerge](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/configGeneration.md#establish-base-and-staging) 实现给 ConfigMap 打补丁

```shell
$ tree
.
├── configs
│   ├── config.yaml
│   └── kustomization.yaml
└── kustomize
    ├── base
    │   └── kustomization.yaml
    └── overlays
        ├── dev
        │   └── kustomization.yaml
        └── staging
            ├── kustomization.yaml
            └── patch.yaml
```

```yaml
# kustomize/overlays/staging/patch.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: yuanya-cm
data:
  foo: foobar
```

```yaml
# kustomize/overlays/staging/kustomization.yaml
resources:
  - ../../base
namePrefix: staging-
nameSuffix: -v1
commonLabels:
  variant: staging
  org: acmeCorporation
commonAnnotations:
  note: Hello, I am staging!
patchesStrategicMerge:
  - patch.yaml
```

```shell
$ kubectl kustomize kustomize/overlays/staging/
apiVersion: v1
data:
  config: |
    author:
      name: yuanya
  config.yaml: |
    author:
      name: yuanya
  foo: foobar # 配置替换
kind: ConfigMap
metadata:
  annotations:
    annotations.only.for.generated: othervalue
    note: Hello, I am staging!
  labels:
    kustomize.generated.resource: somevalue
    org: acmeCorporation
    variant: staging
  name: staging-yuanya-cm-v1
```





## helm

> helm：

从 GitHub [helm 仓库](https://github.com/helm/helm/tags) 下载对应版本压缩包

```shell
# 解压并移入已添加环境变量的某个目录，或者添加环境变量
tar -zxf /mnt/share/helm-v3.6.3-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm
rm -rf linux-amd64/
```

