

+++

title = "JetBrains"
description = "JetBrains"
tags = ["techn","compute","lang","ide"]

+++

# JetBrains

> [官网](https://www.jetbrains.com/)；



## goland

下载并解压：https://www.jetbrains.com/go/download/other.html

```sh
# goland 别名，使 nohup，并 stdout 重定向到 /dev/null
alias goland="nohup ~/compute/lang/go/goland/bin/goland.sh >/dev/null & 2>&1"'
```



### 常用设置


1. go

2. go mod：开启即可

3. proto import

   1. Goland插件Protocol Buffer Editor：Configure → Setting → Protocol Buffers → 取消勾选Configure automatically → 添加所需目录
      1. 一般是当前project或module的上一级目录，以使生成的go文件依然保持正确的包导入路径，且前缀统一更具可读性；D:/it/go/GoProject/pkg/mod
      2. 如果用到go mod导入的内容，还需要添加go mod包所在目录。D:/it/go/GoProject/pkg/mod
   2. 如果是`protoc`命令工具：则添加 `-I` 参数指定目录

4. 编码：Configure → Setting → Editor → File Encodings

5. 全选为UTF-8，with NO BOM

   2. Transparent native-to-ascii conversion自动转换ASCII编码：建议勾选。可能意思是，在文件中输入文字时他会自动的转换为Unicode编码，然后在idea中发开文件时他会自动转回文字来显示

6. 注解生效激活：Configure → Setting → Compiler → Annotation Processors

   1. Enable annotation processing

7. 文件过滤：Configure → Setting → Editor → File Types

   1. ActionScript → Ignore files and folders：添加 *.idea;\*.iml; 进去

8. 字体：Configure → Setting → Editor → Font

9. 主题：Configure → Setting → Editor → Color Scheme

10. 背景图：Ctrl+Shift+A，搜索set background Image

11. 参数名提示：Ctrl+Shift+A，搜索show parameter name hints

12. 文档模板注释：

    ```java
       /* 类文档注释模板配置
       1.file - settings - editor - file and data templates
       2.includs - file header - 右空白框内复制下面类文档注释模板
       3.新建类时自动生成文档注释
       */
       /** 类文档模板
       @title: ${NAME}
       @projectName ${PROJECT_NAME}
       @description: TODO
       @author ${USER}
       @date ${DATE}${TIME}
        */
       public class DocumentAnnotationTemplate {
           /* 方法文档注释模板配置
           1.file - settings - editor - live templates
           2.右+号 - template group - 名aaa(在最前比较方便)
           3.选中aaa - live template - 名ann - 描述ann - 空白框内复制下面方法注释模板 - define - everywhere
           3.新建方法后在方法上面输入ann - 按tab补全
           */
           /** 方法模板
           @description: TODO
           @param ${tags}
           @return ${return_type}
           @throws
           @author ${USER}
           @date ${DATE}${TIME}
            */
           public void function() { }
    }
    ```

13. evaluate expression：https://www.cnblogs.com/mrmoo/p/9942605.html



### 重置试用

方式1：删除家目录下的试用配置文件夹

```shell
rm -rf ~/.config/JetBrains/GoLand20xx.x/eval/
# win下为~\AppData\Local\JetBrains\GoLand2021.2
```

方式2：安装插件 Eval Reset，点击 help->Eval Reset->reset 即可无限[重置试用](https://zhile.io/2020/11/18/jetbrains-eval-reset-deprecated.html)。适用于 2021.2.3 及以前，2021.2.4 开始必须联网登录账号才能使用。可以勾选 Auto reset before per restart

方式3：破解。破解前打开 goland，选择试用；下载破解文件 jetbrains-agent.jar 2020版本，将 jetbrains-agent.jar 拖入goland窗口，会提示restart，点击restart即可完成破解。适用于2021.1.2及其以前（现在好像2021.1.2已经不行了）

