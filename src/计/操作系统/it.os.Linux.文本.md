+++

title = "Linux.文本"
description = "Linux.文本"
tags = ["it","os","Linux"]

+++

# Linux.文本



## grep

`grep`：

作用：打印匹配模式的行

语法

- `grep [OPTIONS] PATTERN [FILE...]`
- `grep [OPTIONS] -e PATTERN ... [FILE...]`
- `grep [OPTIONS] -f FILE ... [FILE...]`

```shell

```



## sed

http://c.biancheng.net/view/4028.html





### s替换

格式：`sed '[address]s/<pattern>/<replacement>/<flags>' <filename>`

```shell
# 替换一次
sed -i 's/<pattern>/<replacement>/'
# 替换所有
sed -i 's/<pattern>/<replacement>/g'
```

### i插入

```shell
# 在匹配行前面，添加一行内容。
sed -i '/<pattern>/i<replacement>'

# 在最后一行前面，添加一行内容。$表示最后
sed -i '$i内容' file
```

### a插入

```shell
# 在匹配行后面，添加一行内容。
sed -i '/<pattern>/a<replacement>'
# 在最后一行后面，添加一行内容。$表示最后。
sed -i '$a内容' file
```

### 匹配行并在前后插入内容

```shell
# 替换所有
sed -i '/pattern/ireplacement/g' file
```



## 重定向符号

[重定向符号](https://blog.csdn.net/hellozpc/article/details/46721811)：

- 标准输入输出重定向：即指定命令产生的 stdout（标准输出信息）、stdin（标准输入信息） 写入哪个文件
  - `>`：输出重定向到一个文件或设备，覆盖原来的文件。
    - `goland.sh >/dev/null` 表示将 goland 运行产生的 stdout 的输出到 /dev/null 中
  - `>!`：输出重定向到一个文件或设备 强制覆盖原来的文件
  - `>>`：输出重定向到一个文件或设备 追加原来的文件
  - `<`：输入重定向到一个程序
- 标准错误重定向：即指定命令产生的 stderr（标准错误信息） 写入哪个文件
  - `2>`：将一个标准错误输出重定向到一个文件或设备 覆盖原来的文件 b-shell
    - `goland.sh >/dev/null` 表示将 goland 运行产生的 stdout 的输出到 /dev/null 中
  - `2>>`：将一个标准错误输出重定向到一个文件或设备 追加到原来的文件
  - `2>&1`：将一个标准错误输出重定向到标准输出 注释:1 可能就是代表 标准输出
  - `>&`：将一个标准错误输出重定向到一个文件或设备 覆盖原来的文件 c-shell
  - `|&`：将一个标准错误 管道 输送 到另一个命令作为输入

重定向符号说明

| 符号    | 说明                                                         |
| ------- | ------------------------------------------------------------ |
| <       | 输入符号：将 file 中的内容作为输入                           |
| <<      | 输入符号：从键盘输入多行内容，遇到自定义的标识符（一般使用 EOF）时结束 |
| > file  | 输出符号：将输入内容**覆盖**输出到 file 中                   |
| >> file | 输出符号：将输入内容**追加**输出到 file 中                   |

结合 cat 使用

```bash
# 将命令下方键盘输入的内容作为命令的输入，遇到 EOF(end of file) 时结束，并输出到文件 demo01 中
$ cat > demo01 << EOF
> hello
> EOF
$ cat demo01
hello

# EOF 并不是固定的，可以使用任意自定义标识
$ cat > demo01 << END
> hello
> END
$ cat demo01
hello

# 使用 `>` 将覆盖文件内容
$ cat > demo01 << EOF
> asdasd
> EOF
$ cat demo01
asdasd

# 使用 `>>` 则可以追加内容
$ cat >> demo01 << EOF
> hello
> EOF
$ cat demo01
asdasd
hello

# 将文件 demo01 的内容作为命令的输入，覆盖输出到文件 demo02 中。追加使用 >> 即可
$ cat > demo02 < demo01
$ cat demo02
asdasd
hello
```

