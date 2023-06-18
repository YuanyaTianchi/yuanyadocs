

+++

title = "Kubernetes.operator"
description = "Title sub_title quick_start"
tags = ["techn", "computer", "tagx", "Title", "_sub_title", "__quick_start(content)"]

+++



# Kubernetes.operator

> [官网]()；





# Kubebuilder

#### 快开

```shell
# 初始化项目
kuberbuilder init --domain yuanyatianchi.io --repo yuanyatianchi.io/<project_name> --component-config

kuberbuilder edit --multigroup  true

kubebuilder create api --group <group_name>.yuanyatianchi.io --version v1 --kind <CrdName>

kubebuilder create webhook --group <group_name>.yuanyatianchi.io --version v1 --kind <CrdName> --defaulting --programmatic-validation
```



```shell
kubebuilder create webhook --group <group_name>.yuanyatianchi.io --version v1 --kind <CrdName> --defaulting --programmatic-validation
```

