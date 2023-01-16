

+++

title = "crypto.tool"
description = "crypto.tool"
tags = ["it","crypto"]

+++

# crypto.tool

## cfssl

> [cfssl](https://github.com/cloudflare/cfssl)；

### 安装

golang 环境则可以一键安装

```shell
$ go install github.com/cloudflare/cfssl/cmd/...@latest
$ ls $GOPATH/bin
cfssl  cfssl-bundle  cfssl-certinfo  cfssljson  cfssl-newkey  cfssl-scan
```

### CA证书

CA 证书：用于签发其它证书，因为其他证书从 CA 证书开始，一般 CA 证书即 根（root）证书

#### csr

csr：证书签名请求

`cfssl print-defaults csr` 可以打印 csr 文件模板作为参考

```shell
cfssl print-defaults csr > ca-csr.json
```

根据需要进行修改

```shell
cat > ca-csr.json <<EOF
{
    "CN": "yuanyatianchi.io",
    "hosts": [
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "Sichuan",
            "L": "Chengdu",
            "O": "yuanya",
            "OU": "tianchi"
        }
    ],
    "ca": {
      "expiry": "8760h"
    }
}
EOF
```

- hosts: 证书通信可用主机地址列表（使用证书的主机列表 ？？）
- CN: Common Name，浏览器使用该字段验证网站是否合法，一般写的是地址端口或域名。
- C：Contry，国家
- ST：State，省、州
- L：Locality，地区、城市
- O：Organization Name，组织名称、公司名称
- OU： Organization Unit Name，组织单位名称，公司名称
- ca.expiry：证书过期时间，默认有效期 8760h（1年）

> `CN`字段这里其实已经[无所谓](https://www.ethanzhang.xyz/cfssl%E4%BD%BF%E7%94%A8%E6%96%B9%E6%B3%95/#33-%E7%94%9F%E6%88%90%E6%9C%8D%E5%8A%A1%E5%99%A8%E8%AF%81%E4%B9%A6)了，从Chrome 58开始，只通过校验SAN属性验证证书的有效性。
>
> SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。
>
> SAN属性定义在`hosts`属性内。

#### 证书

生成 CA 证书和私钥，`cfssl gencert -initca ca-csr.json` 只会打印出证书和私钥，还需借助 `cfssl-json` 将 `cfssl` 生成的内容解析到文件

```shell
$ cfssl gencert -initca ca-csr.json | cfssljson -bare ca
$ ls ca*
ca.csr  ca-csr.json  ca-key.pem  ca.pem
```

- ca.csr：由 `ca-csr.json` 加密得到，生成证书和私钥真正使用的 csr 文件
- ca.pem：证书
- ca-key.pem：私钥



### 签发证书

> 此处使用我们自生成（非商业ca机构颁发）的ca证书进行签名，所以是自签名证书

#### 签发配置

签发配置：ca 证书签名时使用的配置文件

`cfssl print-defaults config` 可以打印模板作为参考

```shell
cfssl print-defaults config > ca-config.json
```

根据需要进行修改

```shell
cat > ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "87600h"
    },
    "profiles": {
      "client": {
        "usages": [
            "signing",
            "key encipherment",
            "client auth"
        ],
        "expiry": "87600h"
      },
      "server": {
        "usages": [
            "signing",
            "key encipherment",
            "server auth"
        ],
        "expiry": "87600h"
      },
      "peer": {
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

这里设置了三种证书配置

- server: 服务器启动（响应？）需要证书
- client: 客户端连接服务端需要证书
- peer: 对端通信两端都需要证书（双向验证）

下面以 etcd 集群需要的双向通信证书为例演示

#### csr

```shell
# 证书签名请求
cat > etcd-csr.json <<EOF
{
    "CN": "etcd",
    "hosts": [
      "172.20.0.111",
      "172.20.0.112",
      "172.20.0.113"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "BeiJing",
            "L": "BeiJing",
            "O": "etcd",
            "OU": "etcd"
        }
    ]
}
EOF
```

hosts 中添加网段第一个地址如 `"172.20.0.1"`，貌似可以直接监听该网段全部主机？

#### 签发证书

-ca 指定根证书，-ca-key 指定根证书密钥，-config 指定签发配置文件，-profile 指定配置域。同样需要借助 `cfssl-json` 将 `cfssl` 生成的内容解析到文件

```shell
# 签发
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=peer etcd-csr.json | cfssljson -bare etcd
```

#### 证书信息

查看指定域名的证书信息

```shell
cfssl-certinfo -domain www.baidu.com 
```

### 证书验证

```shell
cfssl-certinfo -cert ca.pem
```

## openssl

自签名证书

```shell
$ openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
-keyout example-key.pem -out example.pem \
-subj "/C=CN/ST=Beijing/L=Beijing/O=registry/OU=registry/CN=registry"
```

查看证书

```shell
openssl x509 -noout -text -in ca.pem
```
