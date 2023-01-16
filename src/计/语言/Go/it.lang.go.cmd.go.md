

+++

title = "go.cmd.go"
description = "go.cmd.go"
tags = ["it","lang","go"]

+++

# go.cmd.go



https://golang.org/cmd/go/



作用：管理 Go 源代码的工具

```shell
# 直接输入 go 获取帮助信息
$ go
...
```

语法：`go <command> [arguments]`

命令（常用）

- `build`：编译包和依赖

```shell
# go help <command>
go help build
# go help <topic>
go help 
```



### help

作用：打印帮助信息到标准输出

语法

- 



### bug

### build

作用：用于编译指定的源码文件或代码包以及它们的依赖包，将在当前文件夹下生成可执行文件

```shell
# 直接输入 go tool compile 获取帮助信息
$ go tool compile
...
```

语法：`go tool compile [options] file.go...`

参数（常用）

- `-S`：汇编列表打印到标准输出（仅代码）。可以看到汇编代码。
- `-S -S`：汇编列表打印到标准输出（代码 & 数据）。
- `-N`：禁用优化。golang 编译过程中可能会进行一些优化，会使汇编代码与源代码过程不完全一致。
- `-l`：禁用内联。

`go build`：不后跟任何代码包，将试图编译当前目录所对应的代码包

`go build yuanyatianchi.io/cmd`：指定代码包编译，没有 相对/绝对路径 需要保证代码包在 `GOPATH` 下

`go build ./cmd/main.go ./pkg/demo.go `：指定多文件编译，且必须有且仅有一个包含 main 函数的文件



`-a`：编译所有依赖的代码包（如果已经编译过一次，默认仅重新编译变化的代码包）

`-v`：将打印被编译到的代码包

`-n`：将打印编译期间用到的其它命令，但并不真正执行它们。

`-x`：将打印编译期间所用到的其它命令，但不会取消这些命令的执行。

`-work`：将打印编译时生成的临时工作目录，并在编译结束时保留它（编译结束时，默认会删除该目录）

`-p n`：指定编译过程中执行各任务的并发量，默认值为 CPU 的逻辑核数。在`darwin/arm`平台下默认值值为 1。



```shell
# `-gcflags`：。-S 即 go tool compile 的 -S 参数，将汇编列表打印到标准输出（仅代码）
$ go build -gcflags -S
go build -gcflags -S add.go
# command-line-arguments
"".add STEXT nosplit size=4 args=0x10 locals=0x0 funcid=0x0
        0x0000 00000 (D:\it\go\goproject\yuanyatianchi\cmd\demo\add.go:3)       TEXT    "".add(SB), NOSPLIT|ABIInternal, $0-16
        0x0000 00000 (D:\it\go\goproject\yuanyatianchi\cmd\demo\add.go:3)       FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (D:\it\go\goproject\yuanyatianchi\cmd\demo\add.go:3)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (D:\it\go\goproject\yuanyatianchi\cmd\demo\add.go:3)       FUNCDATA        $5, "".add.arginfo1(SB)
        0x0000 00000 (D:\it\go\goproject\yuanyatianchi\cmd\demo\add.go:4)       ADDQ    BX, AX
        0x0003 00003 (D:\it\go\goproject\yuanyatianchi\cmd\demo\add.go:4)       RET
        0x0000 48 01 d8 c3                                      H...
...
```



##### -race

开启竞态条件的检测。不过此标记目前仅在`linux/amd64`、`freebsd/amd64`、`darwin/amd64`和`windows/amd64`平台下受到支持。

准备一个存在数据竞争的程序

```go
package main

import "fmt"

func main() {
	a := make([]int, 1)

	go func() {
		a[0] = 1
	}()

	fmt.Println("a[0] =", a[0])
}
```

```shell
# go build、go run、go test 均可以使用 -race 参数
$ go build -race -o main main.go && ./main
a[0] = 0
==================
WARNING: DATA RACE
Write at 0x00c0000b4008 by goroutine 7:
  main.main.func1()
      /root/go/path/yuanyatianchi.io/go/tool/build/main.go:9 +0x44

Previous read at 0x00c0000b4008 by main goroutine:
  main.main()
      /root/go/path/yuanyatianchi.io/go/tool/build/main.go:12 +0xdb

Goroutine 7 (running) created at:
  main.main()
      /root/go/path/yuanyatianchi.io/go/tool/build/main.go:8 +0xad
==================
Found 1 data race(s)
```

经 race 检测，在同一个内存地址上，goroutine 7 和 main goroutine 存在读写竞争，且打印了 goroutine 7 创建的源码位置

##### -tags

用于指定在实际编译期间，需要受理的编译标签（也可被称为编译约束）的列表。这些编译标签，一般会作为源码文件开始处的注释的一部分，例如，在`$GOROOT/src/os/file_posix.go`开始处的注释为：

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris windows
```

最后一行注释`// +build ...`即包含了与编译标签有关的内容。大家可以查看代码包`go/build`的文档已获得更多的关于编译标签的信息。

```sh
go build -tags darwin
```





##### other

明：

- `-asmflags`

此标记可以后跟另外一些标记，如`-D`、`-I`、`-S`等。这些后跟的标记用于控制Go语言编译器编译汇编语言文件时的行为。

- `-buildmode`

此标记用于指定编译模式，使用方式如`-buildmode=default`（这等同于默认情况下的设置）。此标记支持的编译模式目前有6种。借此，我们可以控制编译器在编译完成后生成静态链接库（即.a文件，也就是我们之前说的归档文件）、动态链接库（即.so文件）或/和可执行文件（在Windows下是.exe文件）。

- `-compiler`

此标记用于指定当前使用的编译器的名称。其值可以为`gc`或`gccgo`。其中，gc编译器即为Go语言自带的编辑器，而gccgo编译器则为GCC提供的Go语言编译器。而GCC则是GNU项目出品的编译器套件。GNU是一个众所周知的自由软件项目。在开源软件界不应该有人不知道它。好吧，如果你确实不知道它，赶紧去google吧。

- `-gccgoflags`

此标记用于指定需要传递给gccgo编译器或链接器的标记的列表。

- `-gcflags`

此标记用于指定需要传递给`go tool compile`命令的标记的列表。

- `-installsuffix`

为了使当前的输出目录与默认的编译输出目录分离，可以使用这个标记。此标记的值会作为结果文件的父目录名称的后缀。其实，如果使用了`-race`标记，这个标记会被自动追加且其值会为`race`。如果我们同时使用了`-race`标记和`-installsuffix`，那么在`-installsuffix`标记的值的后面会再被追加`_race`，并以此来作为实际使用的后缀。

- `-ldflags`

此标记用于指定需要传递给`go tool link`命令的标记的列表。

- `-linkshared`

此标记用于与`-buildmode=shared`一同使用。后者会使作为编译目标的非`main`代码包都被合并到一个动态链接库文件中，而前者则会在此之上进行链接操作。

- `-pkgdir`

使用此标记可以指定一个目录。编译器会只从该目录中加载代码包的归档文件，并会把编译可能会生成的代码包归档文件放置在该目录下。

- `-toolexec`

此标记可以让我们去自定义在编译期间使用一些Go语言自带工具（如`vet`、`asm`等）的方式。



#### 交叉编译执行

cmd/hello.go

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

Makefile

```makefile
.PHONY: build build-hello.amd build-hello.arm clean

all: build-hello.amd build-hello.arm

build-hello.amd64: ./cmd/hello.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/hello.amd64 ./cmd/hello.go

build-hello.arm64: ./cmd/hello.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/hello.arm64 ./cmd/hello.go

clean:
	rm -rf ./bin
```

编译 & 执行

```sh
# 编译 x86 可执行文件 & 交叉编译 arm 可执行文件
$ make all
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/hello.amd64 ./cmd/hello.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/hello.arm64 ./cmd/hello.go

# x86-64 文件可以执行
$ file ./bin/hello.amd64
./bin/hello.amd: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
$ ./bin/hello.amd64
hello

# ARM aarch64 文件不可以执行
$ file ./bin/hello.arm64
./bin/hello.arm: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, not stripped
$ ./bin/hello.arm64
-bash: ./bin/hello.arm: 无法执行二进制文件: 可执行文件格式错误
```

通过 qemu-aarch64-static （该程序使用内核 binfmt_misc 功能）模拟 aarch 环境

```shell
# 下载 qemu-aarch64-static
$ curl -L -O https://github.com/multiarch/qemu-user-static/releases/download/v5.1.0-8/qemu-aarch64-static.tar.gz
$ tar -zxf qemu-aarch64-static.tar.gz
$ mv qemu-aarch64-static /usr/bin/qemu-aarch64-static

# 使用 qemu-aarch64-static 执行 hello.aarch64
$ qemu-aarch64-static ./bin/hello.arm64
hello
```



### clean

执行`go clean`命令会删除掉执行其它命令时产生的一些文件和目录，包括：

1. 在使用`go build`命令时在当前代码包下生成的与包名同名或者与Go源码文件同名的可执行文件。在Windows下，则是与包名同名或者Go源码文件同名且带有“.exe”后缀的文件。
2. 在执行`go test`命令并加入`-c`标记时在当前代码包下生成的以包名加“.test”后缀为名的文件。在Windows下，则是以包名加“.test.exe”后缀为名的文件。我们会在后面专门介绍`go test`命令。
3. 如果执行`go clean`命令时带有标记`-i`，则会同时删除安装当前代码包时所产生的结果文件。如果当前代码包中只包含库源码文件，则结果文件指的就是在工作区的pkg目录的相应目录下的归档文件。如果当前代码包中只包含一个命令源码文件，则结果文件指的就是在工作区的bin目录下的可执行文件。
4. 还有一些目录和文件是在编译Go或C源码文件时留在相应目录中的。包括：“_obj”和“_test”目录，名称为“_testmain.go”、“test.out”、“build.out”或“a.out”的文件，名称以“.5”、“.6”、“.8”、“.a”、“.o”或“.so”为后缀的文件。这些目录和文件是在执行`go build`命令时生成在临时目录中的。如果你忘记了这个临时目录是怎么回事儿，可以再回顾一下前面关于`go build`命令的介绍。临时目录的名称以`go-build`为前缀。
5. 如果执行`go clean`命令时带有标记`-r`，则还包括当前代码包的所有依赖包的上述目录和文件。

我们再以goc2p项目的`logging`为例。为了能够反复体现每个标记的作用，我们会使用标记`n`。使用标记`-n`会让命令在执行过程中打印用到的系统命令，但不会真正执行它们。如果想既打印命令又执行命令则需使用标记`-x`。现在我们来试用一下`go clean`命令：

```bash
hc@ubt:~/golang/goc2p/src$ go clean -x logging   
cd /home/hc/golang/goc2p/src/logging
rm -f logging logging.exe logging.test logging.test.exe
```

现在，我们加上标记`-i`：

```bash
hc@ubt:~/golang/goc2p/src$ go clean -x -i logging   
cd /home/hc/golang/goc2p/src/logging
rm -f logging logging.exe logging.test logging.test.exe
rm -f /home/hc/golang/goc2p/pkg/linux_386/logging.a
```

如果再加上标记`-r`又会打印出哪些命令呢？请读者自己试一试吧。



### doc

`go doc`命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

**插播：**所谓Go语言的程序实体，是指变量、常量、函数、结构体以及接口。而程序实体的标识符即是代表它们的名称。标识符又分非限定标识符和限定标识符。其中，限定标识符一般用于表示某个代码包中的程序实体或者某个结构体类型中的方法或字段。例如，标准库代码包`io`中的名为`EOF`的变量用限定标识符表示即`io.EOF`。又例如，如果我有一个`sync.WaitGroup`类型的变量`wg`并且想调用它的`Add`方法，那么可以这样写`wg.Add()`。其中，`wg.Add`就是一个限定标识符，而后面的`()`则代表了调用操作。

下面说明怎样使用`go doc`命令。先来看一下`go doc`命令课结束的标记。

*表0-5 `go doc`命令的标记说明*

| 标记名称 | 标记描述                                                     |
| :------- | :----------------------------------------------------------- |
| -c       | 加入此标记后会使`go doc`命令区分参数中字母的大小写。默认情况下，命令是大小写不敏感的。 |
| -cmd     | 加入此标记后会使`go doc`命令同时打印出`main`包中的可导出的程序实体（其名称的首字母大写）的文档。默认情况下，这部分文档是不会被打印出来的。 |
| -u       | 加入此标记后会使`go doc`命令同时打印出不可导出的程序实体（其名称的首字母小写）的文档。默认情况下，这部分文档是不会被打印出来的。 |

这几个标记的意图都非常简单和明确，大家可以根据实际情况选用。

`go doc`命令可以后跟一个或两个参数。当然，我们也可以不附加任务参数。如果不附加参数，那么`go doc`命令会试图打印出当前目录所代表的代码包的文档及其中的包级程序实体的列表。

例如，我要在goc2p项目的`loadgen`代码包所在目录中运行`go doc`命令的话，那么就会是这样：

```bash
hc@ubt:~/golang/goc2p/src/loadgen$ go doc
package loadgen // import "loadgen"

func NewGenerator(
    caller lib.Caller,
    timeoutNs time.Duration,
    lps uint32,
    durationNs time.Duration,
    resultCh chan *lib.CallResult) (lib.Generator, error)
```

如果你需要指定代码包或程序实体，那么就需要在`go doc`命令后附上参数了。例如，只要我本地的goc2p项目的所在目录存在于GOPATH环境变量中，我就可以在任意目录中敲入`go doc loadgen`。如此得到的输出一定是与上面那个示例一致的。

看过`loadgen`代码包中源码的读者会知道，其中只有一个可导出的程序实体，即`NewGenerator`函数。这也是上述示例中如此输出的原因。该代码包中的结构体类型`myGenerator`是不可导出，但是我们只需附加`-u`标记便可查看它的文档了：

```bash
hc@ubt:~$ go doc -u loadgen.myGenerator
type myGenerator struct {
    caller      lib.Caller           // 调用器。
    timeoutNs   time.Duration        // 处理超时时间，单位：纳秒。
    lps         uint32               // 每秒载荷量。
    durationNs  time.Duration        // 负载持续时间，单位：纳秒。
    concurrency uint32               // 并发量。
    tickets     lib.GoTickets        // Goroutine票池。
    stopSign    chan byte            // 停止信号的传递通道。
    cancelSign  byte                 // 取消发送后续结果的信号。
    endSign     chan uint64          // 完结信号的传递通道，同时被用于传递调用执行计数。
    callCount   uint64               // 调用执行计数。
    status      lib.GenStatus        // 状态。
    resultCh    chan *lib.CallResult // 调用结果通道。
}

    载荷发生器的实现。

func (gen *myGenerator) Start()
func (gen *myGenerator) Status() lib.GenStatus
func (gen *myGenerator) Stop() (uint64, bool)
func (gen *myGenerator) asyncCall()
func (gen *myGenerator) genLoad(throttle <-chan time.Time)
func (gen *myGenerator) handleStopSign(callCount uint64)
func (gen *myGenerator) init() error
func (gen *myGenerator) interact(rawReq *lib.RawReq) *lib.RawResp
func (gen *myGenerator) sendResult(result *lib.CallResult) bool
```

如此一来，`loadgen.myGenerator`类型的文档、字段和方法都尽收眼底。注意，这里我们使用到了限定标识符。下面再进一步，如果你只想查看`loadgen.myGenerator`类型的`init`方法的文档，那么只要续写这个限定标识符就可以了，像这样：

```bash
hc@ubt:~$ go doc -u loadgen.myGenerator.init
func (gen *myGenerator) init() error

    初始化载荷发生器。
```

注意，结构体类型中的字段的文档是无法被单独打印的。另外，`go doc`命令根据参数查找代码包或程序实体的顺序是：先Go语言根目录（即GOROOT所环境变量指定的那个目录）后工作区目录（即GOPATH环境变量包含的那些目录）。并且，在前者或后者中，`go doc`命令的查找顺序遵循字典序。因此，如果某个工作区目录中的代码包与标准库中的包重名了，那么它是无法被打印出来的。`go doc`命令只会打印出第一个匹配的代码包或程序实体的文档。

我们在前面说过，`go doc`命令还可以接受两个参数。这是一种更加精细的指定代码包或程序实体的方式。一个显著的区别是，如果你想打印标准库代码包`net/http`中的结构体类型`Request`的文档，那么可以这样敲入`go doc`命令：

```bash
go doc http.Request
```

注意，这里并没有写入`net/http`代码包的导入路径，而只是写入了其中的最后一个元素`http`。但是如果你把`http.Request`拆成两个参数（即`http Request`）的话，命令程序就会什么也查不到了。因为这与前一种用法的解析方式是不一样的。正确的做法是，当你指定两个参数时，作为第一个参数的代码包名称必须是完整的导入路径，即：在敲入命令`go doc net/http Request`后，你会得到想要的结果。

最后，在给定两个参数时，`go doc`会打印出所有匹配的文档，而不是像给定一个参数时那样只打印出第一个匹配的文档。这对于查找只有大小写不同的多个方法（如`New`和`new`）的文档来说非常有用。

##### godoc

命令`godoc`是一个很强大的工具，同样用于展示指定代码包的文档。在Go语言的1.5版本中，它是一个内置的标准命令。

该命令有两种模式可供选择。如果在执行命令时不加入`-http`标记，则该命令就以命令行模式运行。在打印纯文本格式的文档到标准输出后，命令执行就结束了。比如，我们用命令行模式查看代码包fmt的文档：

```bash
hc@ubt:~$ godoc fmt
```

为了节省篇幅，我们在这里略去了文档查询结果。读者可以自己运行一下上述命令。在该命令被执行之后，我们就可以看到编排整齐有序的文档内容了。这包括代码包`fmt`及其中所有可导出的包级程序实体的声明、文档和例子。

有时候我们只是想查看某一个函数或者结构体类型的文档，那么我们可以将这个函数或者结构体的名称加入命令的后面，像这样：

```bash
hc@ubt:~$ godoc fmt Printf
```

或者：

```bash
hc@ubt:~$ godoc os File
```

如果我们想同时查看一个代码包中的几个函数的文档，则仅需将函数或者结构体名称追加到命令后面。比如我们要查看代码包`fmt`中函数`Printf`和函数`Println`的文档：

```bash
hc@ubt:~$ godoc fmt Printf Println
```

如果我们不但想在文档中查看可导出的程序实体的声明，还想看到它们的源码，那么我们可以在执行`godoc`命令的时候加入标记`-src`，比如这样：

```bash
hc@ubt:~$ godoc -src fmt Printf
```

Go语言为程序使用示例代码设立了专有的规则。我们在这里暂不讨论这个规则的细节。只需要知道正因为有了这个专有规则，使得`godoc`命令可以根据这些规则提取相应的示例代码并把它们加入到对应的文档中。如果我们想在查看代码包`net`中的结构体类型`Listener`的文档的同时查看关于它的示例代码，那么我们只需要在执行命令时加入标记`-ex`。使用方法如下：

```bash
hc@ubt:~$ godoc -ex net/http FileServer
```

注意，我们在使用`godoc`命令时，只能把代码包和程序实体的标识符拆成两个参数。也就是说，`godoc`命令不支持前文所述的`go doc`命令的单参数用法。

在实际的Go语言环境中，我们可能会遇到一个命令源码文件所产生的可执行文件与代码包重名的情况。比如，这里介绍的标准命令`go`和官方代码包`go`。现在我们要明确的告诉`godoc`命令要查看可执行文件go的文档，我们需要在名称前加入`cmd/`前缀：

```bash
hc@ubt:~$ godoc cmd/go
```

另外，如果我们想查看HTML格式的文档，就需要加入标记`-html`。当然，这样在命令行模式下的查看效果是很差的。但是，如果仔细查看的话，可以在其中找到一些相应源码的链接地址。

一般情况下，`godoc`命令会去Go语言根目录和环境变量GOPATH包含的工作区目录中查找代码包。我们可以通过加入标记`-goroot`来制定一个Go语言根目录。这个被指定的Go语言根目录仅被用于当次命令的执行。示例如下：

```bash
hc@ubt:~$ godoc -goroot="/usr/local/go" fmt
```

现在让我们来看看另外一种模式。如果我们在执行命令时加上`-http`标记则会启用另一模式。这种模式被叫做Web服务器模式，它以Web页面的形式提供Go语言文档。

我们使用如下命令启动这个文档Web服务器：

```bash
hc@ubt:~/golang/goc2p$ godoc -http=:6060
```

标记`-http`的值`:6060`表示启动的Web服务器使用本机的6060端口。之后，我们就可以通过在网络浏览器的地址栏中输入[http://localhost:6060](http://localhost:6060/)来查看以网页方式展现的Go文档了。

![本机的Go文档Web服务首页](https://wiki.jikexueyuan.com/project/go-command-tutorial/images/0-1.png)

*图0-1 本机的Go文档Web服务首页*

这与[Go语言官方站点](http://golang.org/)的Web服务页面如出一辙。这使得我们在不方便访问Go语言官方站点的情况下也可以查看Go语言文档。并且，更便利的是，通过本机的Go文档Web服务，我们还可以查看所有本机工作区下的代码的文档。比如，goc2p项目中的代码包`pkgtool`的页面如下图：

![goc2p项目中的pkgtool包的Go文档页面](https://wiki.jikexueyuan.com/project/go-command-tutorial/images/0-2.png)

*图0-2 goc2p项目中的pkgtool包的Go文档页面*

现在，我们在本机开启Go文档Web服务器，端口为9090。命令如下:

```bash
hc@ubt:~$ godoc -http=:9090 -index
```

注意，要使用`-index`标记开启搜索索引。这个索引会在服务器启动时创建并维护。如果不加入此标记，那么无论在Web页面还是命令行终端中都是无法进行查询操作的。

索引中提供了标识符和全文本搜索信息（通过正则表达式为可搜索性提供支持）。全文本搜索结果显示条目的最大数量可以通过标记`-maxresults`提供。标记`-maxresults`默认值是10000。如果不想提供如此多的结果条目，可以设置小一些的值。甚至，如果不想提供全文本搜索结果，可以将标记`-maxresults`的值设置为0，这样服务器就只会创建标识符索引，而根本不会创建全文本搜索索引了。标识符索引即为对程序实体名称的索引。

正因为在使用了`-index`标记的情况下文档服务器会在启动时创建索引，所以在文档服务器启动之后还不能立即提供搜索服务，需要稍等片刻。在索引为被创建完毕之前，我们的搜索操作都会得到提示信息“Indexing in progress: result may be inaccurate”。

如果我们在本机用`godoc`命令启动了Go文档Web服务器，且IP地址为192.168.1.4、端口为9090，那么我们就可以在另一个命令行终端甚至另一台能够与本机联通的计算机中通过如下命令进行查询了。查询命令如下：

```bash
    hc@ubt:~$ godoc -q -server="192.168.1.4:9090" Listener
```

命令的最后为要查询的内容，可以是任何你想搜索的字符串，而不仅限于代码包、函数或者结构体的名称。

标记`-q`开启了远程查询的功能。而标记`-server="192.168.1.4:9090"`则指明了远程文档服务器的IP地址和端口号。实际上，如果不指明远程查询服务器的地址，那么该命令会自行将地址“:6060”和“golang.org”作为远程查询服务器的地址。这两个地址即是默认的本机文档Web站点地址和官方的文档Web站点地址。所以执行如下命令我们也可以查询到标准库的信息：

```bash
hc@ubt:~$ godoc -q fmt
```

命令`godoc`还有很多可用的标记，但在通常情况下并不常用。读者如果有兴趣，可以在命令行环境下敲入`godoc`并查看其文档。

至于怎样才能写出优秀的代码包文档，我在《Go并发编程实战》的5.2节中做了详细说明。



### env

令`go env`用于打印Go语言的环境信息。其中的一些信息我们在之前已经多次提及，但是却没有进行详细的说明。在本小节，我们会对这些信息进行深入介绍。我们先来看一看`go env`命令情况下都会打印出哪些Go语言通用环境信息。

*表0-25 `go env`命令可打印出的Go语言通用环境信息*

| 名称        | 说明                                     |
| :---------- | :--------------------------------------- |
| CGO_ENABLED | 指明cgo工具是否可用的标识。              |
| GOARCH      | 程序构建环境的目标计算架构。             |
| GOBIN       | 存放可执行文件的目录的绝对路径。         |
| GOCHAR      | 程序构建环境的目标计算架构的单字符标识。 |
| GOEXE       | 可执行文件的后缀。                       |
| GOHOSTARCH  | 程序运行环境的目标计算架构。             |
| GOOS        | 程序构建环境的目标操作系统。             |
| GOHOSTOS    | 程序运行环境的目标操作系统。             |
| GOPATH      | 工作区目录的绝对路径。                   |
| GORACE      | 用于数据竞争检测的相关选项。             |
| GOROOT      | Go语言的安装目录的绝对路径。             |
| GOTOOLDIR   | Go工具目录的绝对路径。                   |

下面我们对这些环境信息进行逐一说明。

**CGO_ENABLED**

通过上一小节的介绍，相信读者对cgo工具已经很熟悉了。我们提到过，标准go命令可以自动的使用cgo工具对导入了代码包C的代码包和源码文件进行处理。这里所说的“自动”并不是绝对的。因为当环境变量CGO_ENABLED被设置为0时，标准go命令就不能处理导入了代码包C的代码包和源码文件了。请看下面的示例：

```
hc@ubt:~/golang/goc2p/src/basic/cgo$ export CGO_ENABLED=0
hc@ubt:~/golang/goc2p/src/basic/cgo$ go build -x
WORK=/tmp/go-build775234613
```

我们临时把环境变量CGO_ENABLED的值设置为0，然后执行`go build`命令并加入了标记`-x`。标记`-x`会让命令程序将运行期间所有实际执行的命令都打印到标准输出。但是，在执行命令之后没有任何命令被打印出来。这说明对代码包`basic/cgo`的构建操作并没有被执行。这是因为，构建这个代码包需要用到cgo工具，但cgo工具已经被禁用了。下面，我们再来运行调用了代码包`basic/cgo`中函数的命令源码文件cgo_demo.go。也就是说，命令源码文件cgo_demo.go间接的导入了代码包`C`。还记得吗？这个命令源码文件被存放在goc2p项目的代码包`basic/cgo`中。示例如下：

```
hc@ubt:~/golang/goc2p/src/basic/cgo$ export CGO_ENABLED=0
hc@ubt:~/golang/goc2p/src/basic/cgo$ go run -work cgo_demo.go
WORK=/tmp/go-build856581210
# command-line-arguments
./cgo_demo.go:4: can't find import: "basic/cgo/lib"
```

在上面的示例中，我们在执行`go run`命令时加入了两个标记——`-a`和`-work`。标记`-a`会使命令程序强行重新构建所有的代码包（包括涉及到的标准库），即使它们已经是最新的了。标记`-work`会使命令程序将临时工作目录的绝对路径打印到标准输出。命令程序输出的错误信息显示，命令程序没有找到代码包`basic/cgo`。其原因是由于代码包`basic/cgo`无法被构建。所以，命令程序在临时工作目录和工作区中都找不到代码包basic/cgo对应的归档文件cgo.a。如果我们使用命令`ll /tmp/go-build856581210`查看临时工作目录，也找不到名为basic的目录。

不过，如果我们在环境变量CGO_ENABLED的值为1的情况下生成代码包`basic/cgo`对应的归档文件cgo.a，那么无论我们之后怎样改变环境变量CGO_ENABLED的值也都可以正确的运行命令源码文件cgo_demo.go。即使我们在执行`go run`命令时加入标记`-a`也是如此。因为命令程序依然可以在工作区中找到之前在我们执行`go install`命令时生成的归档文件cgo.a。示例如下：

```
hc@ubt:~/golang/goc2p/src/basic/cgo$ export CGO_ENABLED=1
hc@ubt:~/golang/goc2p/src/basic/cgo$ go install ../basic/cgo
hc@ubt:~/golang/goc2p/src/basic/cgo$ export CGO_ENABLED=0
hc@ubt:~/golang/goc2p/src/basic/cgo$ go run -a -work cgo_demo.go
WORK=/tmp/go-build130612063
The square root of 2.330000 is 1.526434.
ABC
CFunction1() is called.
GoFunction1() is called.
```

由此可知，只要我们事先成功安装了引用了代码包C的代码包，即生成了对应的代码包归档文件，即使cgo工具在之后被禁用，也不会影响到其它Go语言代码对该代码包的使用。当然，命令程序首先会到临时工作目录中寻找需要的代码包归档文件。

关于cgo工具还有一点需要特别注意，即：当存在交叉编译的情况时，cgo工具一定是不可用的。在标准go命令的上下文环境中，交叉编译意味着程序构建环境的目标计算架构的标识与程序运行环境的目标计算架构的标识不同，或者程序构建环境的目标操作系统的标识与程序运行环境的目标操作系统的标识不同。在这里，我们可以粗略认为交叉编译就是在当前的计算架构和操作系统下编译和构建Go语言代码并生成针对于其他计算架构或/和操作系统的编译结果文件和可执行文件。

**GOARCH**

GOARCH的值的含义是程序构建环境的目标计算架构的标识，也就是程序在构建或安装时所对应的计算架构的名称。在默认情况下，它会与程序运行环境的目标计算架构一致。即它的值会与GOHOSTARCH的值是相同。但如果我们显式的设置了环境变量GOARCH，则它的值就会是这个环境变量的值。

**GOBIN**

GOBIN的值为存放可执行文件的目录的绝对路径。它的值来自环境变量GOBIN。在我们使用`go tool install`命令安装命令源码文件时生成的可执行文件会存放于这个目录中。

**GOCHAR**

GOCHAR的值是程序构建环境的目标计算架构的单字符标识。它的值会根据GOARCH的值来设置。当GOARCH的值为386时，GOCHAR的值就是8。当GOARCH的值为amd64时GOCHAR的值就是6。而GOCHAR的值5与GOARCH的值arm相对应。

GOCHAR主要有两个用途，列举如下：

1. Go语言官方的平台相关的工具的名称会以它的值为前缀。的名称会以GOCHAR的值为前缀。比如，在amd64计算架构下，用于编译Go语言代码的编译器的名称是6g，链接器的名称是6l。用于编译C语言代码的编译器的名称是6c。而用于编译汇编语言代码的编译器的名称为6a。
2. Go语言的官方编译器生成的结果文件会以GOCHAR的值作为扩展名。Go语言的官方编译器6g在对命令源码文件编译之后会把结果文件*go*.6存放到临时工作目录的相应位置中。

**GOEXE**

GOEXE的值会被作为可执行文件的后缀。它的值与GOOS的值存在一定关系，即只有GOOS的值为“windows”时GOEXE的值才会是“.exe”，否则其值就为空字符串“”。这与在各个操作系统下的可执行文件的默认后缀是一致的。

**GOHOSTARCH**

GOHOSTARCH的值的含义是程序运行环境的目标计算架构的标识，也就是程序在运行时所在的计算机系统的计算架构的名称。在通常情况下，它的值不需要被显式的设置。因为用来安装Go语言的二进制分发文件和MSI（Microsoft软件安装）软件包文件都是平台相关的。所以，对于不同计算架构的Go语言环境来说，它都会是一个常量。

**GOHOSTOS**

GOHOSTOS的值的含义是程序运行环境的目标操作系统的标识，也即程序在运行时所在的计算机系统的操作系统的名称。与GOHOSTARCH类似，它的值在不同的操作系统下是固定不变的，同样不需要显式的设置。

**GOPATH**

这个环境信息我们在之前已经提到过很多次。它的值指明了Go语言工作区目录的绝对路径。我们需要显式的设置环境变量GOPATH。如果有多个工作区，那么多个工作区的绝对路径之间需要用分隔符分隔。在windows操作系统下，这个分隔符为“;”。在其它操作系统下，这个分隔符为“:”。注意，GOPATH的值不能与GOROOT的值相同。

**GORACE**

GORACE的值包含了用于数据竞争检测的相关选项。数据竞争是在并发程序中最常见和最难调试的一类bug。数据竞争会发生在多个Goroutine争相访问相同的变量且至少有一个Goroutine中的程序在对这个变量进行写操作的情况下。

数据竞争检测需要被显式的开启。还记得标记`-race`吗？我们可以通过在执行一些标准go命令时加入这个标记来开启数据竞争检测。在这种情况下，GORACE的值就会被使用到了。支持`-race`标记的标准go命令包括：`go test`命令、`go run`命令、`go build`命令和`go install`命令。

GORACE的值形如“option1=val1 option2=val2”，即：选项名称与选项值之间以等号“=”分隔，多个选项之间以空格“ ”分隔。数据竞争检测的选项包括log_path、exitcode、strip_path_prefix和history_size。为了设置GORACE的值，我们需要设置环境变量GORACE。或者，我们也可以在执行go命令时临时设置它，像这样：

```
hc@ubt:~/golang/goc2p/src/cnet/ctcp$ GORACE="log_path=/home/hc/golang/goc2p /race/report strip_path_prefix=home/hc/golang/goc2p/" go test -race
```

关于数据竞争检测的更多细节我们将会在本书的第三部分予以说明。

**GOROOT**

GOROOT会是我们在安装Go语言时第一个碰到Go语言环境变量。它的值指明了Go语言的安装目录的绝对路径。但是，只有在非默认情况下我们才需要显式的设置环境变量GOROOT。这里所说的默认情况是指：在Windows操作系统下我们把Go语言安装到c:\Go目录下，或者在其它操作系统下我们把Go语言安装到/usr/local/go目录下。另外，当我们不是通过二进制分发包来安装Go语言的时候，也不需要设置环境变量GOROOT的值。比如，在Windows操作系统下，我们可以使用MSI软件包文件来安装Go语言。

**GOTOOLDIR**

GOTOOLDIR的值指明了Go工具目录的绝对路径。根据GOROOT、GOHOSTOS和GOHOSTARCH来设置。其值为$GOROOT/pkg/tool/$GOOS_$GOARCH。关于这个目录，我们在之前也提到过多次。

除了上面介绍的这些通用的Go语言环境信息，还两个针对于非Plan 9操作系统的环境信息。它们是CC和GOGCCFLAGS。环境信息CC的值是操作系统默认的C语言编译器的命令名称。环境信息GOGCCFLAGS的值则是Go语言在使用操作系统的默认C语言编译器对C语言代码进行编译时加入的参数。

如果我们要有针对性的查看上述的一个或多个环境信息，可以在`go env`命令的后面加入它们的名字并执行之。在`go env`命令和环境信息名称之间需要用空格分隔，多个环境信息名称之间也需要用空格分隔。示例如下：

```
hc@ubt:~$ go env GOARCH GOCHAR3868
```

上例的`go env`命令的输出信息中，每一行对一个环境信息的值，且其顺序与我们输入的环境信息名称的顺序一致。比如，386为环境信息GOARCH，而8则是环境信息GOCHAR的值。

`go env`命令能够让我们对当前的Go语言环境进行简要的了解。通过它，我们也可以对当前安装的Go语言的环境设置进行简单的检查。



### fix

命令`go fix`会把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。这里所说的版本即Go语言的版本。代码包的所有Go语言源码文件不包括其子代码包（如果有的话）中的文件。修正操作包括把对旧程序调用的代码更换为对新程序调用的代码、把旧的语法更换为新的语法，等等。

这个工具其实非常有用。在编程语言的升级和演进的过程中，难免会对过时的和不够优秀的语法及标准库进行改进。这样的改进对于编程语言的向后兼容性是个挑战。我们在前面提到过向后兼容这个词。简单来说，向后兼容性就是指新版本的编程语言程序能够正确识别和解析用该编程语言的旧版本编写的程序和软件，以及在新版本的编程语言的运行时环境中能够运行用该编程语言的旧版本编写的程序和软件。对于Go语言来说，语法的改变和标准库的变更都会使得用旧版本编写的程序无法在新版本环境中编译通过。这就等于破坏了Go语言的向后兼容性。对于一个编程语言、程序库或基础软件来说，向后兼容性是非常重要的。但有时候为了让软件更加优秀，软件的开发者或维护者不得不在向后兼容性上做出一些妥协。这是一个在多方利益之间进行权衡的结果。本小节所讲述的工具正是Go语言的创造者们为了不让这种妥协给语言使用者带来困扰和额外的工作量而编写的自动化修正工具。这也充分体现了Go语言的软件工程哲学。下面让我们来详细了解它们的使用方法和内部机理。

命令`go fix`其实是命令`go tool fix`的简单封装。这甚至比`go fmt`命令对`gofmt`命令的封装更简单。像其它的Go命令一样，`go fix`命令会先对作为参数的代码包导入路径进行验证，以确保它是正确有效的。像在本小节开始处描述的那样，`go fix`命令会把有效代码包中的所有Go语言源码文件作为多个参数传递给`go tool fix`命令。实际上，`go fix`命令本身不接受任何标记，它会把加入的所有标记都原样传递给`go tool fix`命令。`go tool fix`命令可接受的标记如下表。

*表0-15 `go tool fix`命令的标记说明*

| 标记名称 | 标记描述                                                     |
| :------- | :----------------------------------------------------------- |
| -diff    | 不将修正后的内容写入文件，而只打印修正前后的内容的对比信息到标准输出。 |
| -r       | 只对目标源码文件做有限的修正操作。该标记的值即为允许的修正操作的名称。多个名称之间用英文半角逗号分隔。 |
| -force   | 使用此标记后，即使源码文件中的代码已经与Go语言的最新版本相匹配了，也会强行执行指定的修正操作。该标记的值就是需要强行执行的修正操作的名称，多个名称之间用英文半角逗号分隔。 |

在默认情况下，```go tool fix```命令程序会在目标源码文件上执行所有的修正操作。多个修正操作的执行会按照每个修正操作中标示的操作建立日期以从早到晚的顺序进行。我们可以通过执行```go tool fix -?```来查看```go tool fix```命令的使用说明以及当前支持的修正操作。 与本书对应的Go语言版本的```go tool fix```命令目前只支持两个修正操作。一个是与标准库代码包```go/printer```中的结构体类型```Config```的初始化代码相关的修正操作，另一个是与标准库代码包`net`中的结构体类型```IPAddr```、```UDPAddr```和```TCPAddr```的初始化代码相关的修正操作。从修正操作的数量来看，自第一个正式版发布以来，Go语言的向后兼容性还是很好的。从Go语言官网上的说明也可以获知，在Go语言的第二个大版本（Go 2.x）出现之前，它会一直良好的向后兼容性。 值得一提的是，上述的修正操作都是依靠Go语言的标准库代码包```go```及其子包中提供的功能来完成的。实际上，```go tool fix```命令程序在执行修正操作之前，需要先将目标源码文件中的内容解析为一个抽象语法树实例。这一功能其实就是由代码包```go/parser```提供的。而在这个抽象语法树实例中的各个元素的结构体类型的定义以及检测、访问和修改它们的方法则由代码包```go/ast```提供。有兴趣的读者可以阅读这些代码包中的代码。这对于深入理解Go语言对代码的静态处理过程是非常有好处的。 回到正题。与```gofmt```命令相同，```go tool fix```命令也有交互模式。我们同样可以通过执行不带任何参数的命令来进入到这个模式。但是与```gofmt```命令不同的是，我们在```go tool fix```命令的交互模式中输入的代码必须是完整的，即必须要符合Go语言源码文件的代码组织形式。当我们输入了不完整的代码片段时，命令程序将显示错误提示信息并退出。示例如下： hc@ubt:~$ go tool fix -r='netipv6zone' a := &net.TCPAddr{ip4, 8080} standard input:1:1: expected 'package', found 'IDENT' a 相对于上面的示例，我们必须要这样输入源码才能获得正常的结果： hc@ubt:~$ go tool fix -r='netipv6zone' package main import ( "fmt" "net" ) func main() { addr := net.TCPAddr{"127.0.0.1", 8080} fmt.Printf("TCP Addr: %s\n", addr) } standard input: fixed netipv6zone package main import ( "fmt" "net" ) func main() { addr := net.TCPAddr{IP: "127.0.0.1", Port: 8080} fmt.Printf("TCP Addr: %s\n", addr) } 上述示例的输出结果中有这样一行提示信息：“standard input: fixed netipv6zone”。其中，“standard input”表明源码是从标准输入而不是源码文件中获取的，而“fixed netipv6zone”则表示名为netipv6zone的修正操作发现输入的源码中有需要修正的地方，并且已经修正完毕。另外，我们还可以看到，输出结果中的代码已经经过了格式化。



### fmt

### generate

### get

```shell
go get github.com/go-sql-driver/mysql #命令可以下载依赖包
go get github.com/go-sql-driver/mysql@v1.4.1 #指定下载版本，或升级到指定的版本号version
go get -u github.com/go-sql-driver/mysql@v1.4.1 #将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
go get -u=patch #将会升级到最新的修订版本
```

```go
hc@ubt:~$ go get github.com/hyper-carrot/go_lib/logging
```

命令`go get`可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。在上面这个示例中，我们从著名的代码托管站点Github上下载了一个项目（或称代码包），并安装到了环境变量GOPATH中包含的第一个工作区中。与此同时，我们也知道了这个代码包的导入路径就是github.com/hyper-carrot/go_lib/logging。

一般情况下，为了分离自己与第三方的代码，我们会设置两个或更多的工作区。我们现在有一个目录路径为/home/hc/golang/lib的工作区，并且它是环境变量GOPATH值中的第一个目录路径。注意，环境变量GOPATH中包含的路径不能与环境变量GOROOT的值重复。好了，如果我们使用`go get`命令下载和安装代码包，那么这些代码包都会被安装在上面这个工作区中。我们暂且把这个工作区叫做Lib工作区。在我们运行`go get github.com/hyper-carrot/go_lib/logging`之后，这个代码包就应该会被保存在Lib工作的src目录下，并且已经被安装妥当，如下所示：

```go
/home/hc/golang/lib:
    bin/
    pkg/
        linux_386/
            github.com/
            hyper-carrot/
        go_lib/
            logging.a  
    src/
        github.com/
        hyper-carrot/
        go_lib/
            logging/
    ...
```

另一方面，如果我们想把一个项目上传到Github网站（或其他代码托管网站）上并被其他人使用的话，那么我们就应该把这个项目当做一个代码包来看待。其实我们在之前已经提到过原因，`go get`命令会将项目下的所有子目录和源码文件存放到第一个工作区的src目录下，而src目录下的所有子目录都会是某个代码包导入路径的一部分或者全部。也就是说，我们应该直接在项目目录下存放子代码包和源码文件，并且直接存放在项目目录下的源码文件所声明的包名应该与该项目名相同（除非它是命令源码文件）。这样做可以让其他人使用`go get`命令从Github站点上下载你的项目之后直接就能使用它。

实际上，像goc2p项目这样直接以项目根目录的路径作为工作区路径的做法是不被推荐的。之所以这样做主要是想让读者更容易的理解Go语言的工程结构和工作区概念，也可以让读者看到另一种项目结构。当然，如果你的项目使用了[gb](https://github.com/constabulary/gb)这样的工具那就是另外一回事了。这样的项目的根目录就应该被视为一个工作区（但是你不必把它加入到GOPATH环境变量中）。它应该由`git clone`下载到Go语言工作区之外的某处，而不是使用`go get`命令。

**远程导入路径分析**

实际上，`go get`命令所做的动作也被叫做代码包远程导入，而传递给该命令的作为代码包导入路径的那个参数又被叫做代码包远程导入路径。

`go get`命令不仅可以从像Github这样著名的代码托管站点上下载代码包，还可以从任何命令支持的代码版本控制系统（英文为Version Control System，简称为VCS）检出代码包。任何代码托管站点都是通过某个或某些代码版本控制系统来提供代码上传下载服务的。所以，更严格地讲，`go get`命令所做的是从代码版本控制系统的远程仓库中检出/更新代码包并对其进行编译和安装。

该命令所支持的VCS的信息如下表：

*表0-2 `go get`命令支持的VCS*

| 名称       | 主命令 | 说明                                                         |
| :--------- | :----- | :----------------------------------------------------------- |
| Mercurial  | hg     | Mercurial是一种轻量级分布式版本控制系统，采用Python语言实现，易于学习和使用，扩展性强。 |
| Git        | git    | Git最开始是Linux Torvalds为了帮助管理 Linux 内核开发而开发的一个开源的分布式版本控制软件。但现在已被广泛使用。它是被用来进行有效、高速的各种规模项目的版本管理。 |
| Subversion | svn    | Subversion是一个版本控制系统，也是第一个将分支概念和功能纳入到版本控制模型的系统。但相对于Git和Mercurial而言，它只算是传统版本控制系统的一员。 |
| Bazaar     | bzr    | Bazaar是一个开源的分布式版本控制系统。但相比而言，用它来作为VCS的项目并不多。 |

`go get`命令在检出代码包之前必须要知道代码包远程导入路径所对应的版本控制系统和远程仓库的URL。

如果该代码包在本地工作区中已经存在，则会直接通过分析其路径来确定这几项信息。`go get`命令支持的几个版本控制系统都有一个共同点，那就是会在检出的项目目录中存放一个元数据目录，名称为“.”前缀加其主命令名。例如，Git会在检出的项目目录中加入一个名为“.git”的子目录。所以，这样就很容易判定代码包所用的版本控制系统。另外，又由于代码包已经存在，我们只需通过代码版本控制系统的更新命令来更新代码包，因此也就不需要知道其远程仓库的URL了。对于已存在于本地工作区的代码包，除非要求强行更新代码包，否则`go get`命令不会进行重复下载。如果想要强行更新代码包，可以在执行`go get`命令时加入`-u`标记。这一标记会稍后介绍。

如果本地工作区中不存在该代码包，那么就只能通过对代码包远程导入路径进行分析来获取相关信息了。首先，`go get`命令会对代码包远程导入路径进行静态分析。为了使分析过程更加方便快捷，`go get`命令程序中已经预置了几个著名代码托管网站的信息。如下表：

*表0-3 预置的代码托管站点的信息*

| 名称                        | 主域名          | 支持的VCS                  | 代码包远程导入路径示例                                       |
| :-------------------------- | :-------------- | :------------------------- | :----------------------------------------------------------- |
| Bitbucket                   | bitbucket.org   | Git, Mercurial             | bitbucket.org/user/project bitbucket.org/user/project/sub/directory |
| GitHub                      | github.com      | Git                        | github.com/user/project github.com/user/project/sub/directory |
| Google Code Project Hosting | code.google.com | Git, Mercurial, Subversion | code.google.com/p/project code.google.com/p/project/sub/directory code.google.com/p/project.subrepository code.google.com/p/project.subrepository/sub/directory |
| Launchpad                   | launchpad.net   | Bazaar                     | launchpad.net/project launchpad.net/project/series launchpad.net/project/series/sub/directory launchpad.net/~user/project/branch launchpad.net/~user/project/branch/sub/directory |
| IBM DevOps Services         | hub.jazz.net    | Git                        | hub.jazz.net/git/user/project hub.jazz.net/git/user/project/sub/directory |

一般情况下，代码包远程导入路径中的第一个元素就是代码托管网站的主域名。在静态分析的时候，`go get`命令会将代码包远程导入路径与预置的代码托管站点的主域名进行匹配。如果匹配成功，则在对代码包远程导入路径的初步检查后返回正常的返回值或错误信息。如果匹配不成功，则会再对代码包远程导入路径进行动态分析。至于动态分析的过程，我就不在这里详细展开了。

如果对代码包远程导入路径的静态分析或/和动态分析成功并获取到对应的版本控制系统和远程仓库URL，那么`go get`命令就会进行代码包检出或更新的操作。随后，`go get`命令会在必要时以同样的方式检出或更新这个代码包的所有依赖包。

**自定义代码包远程导入路径**

如果你想把你编写的（被托管在不同的代码托管网站上的）代码包的远程导入路径统一起来，或者不希望让你的代码包中夹杂某个代码托管网站的域名，那么你可以选择自定义你的代码包远程导入路径。这种自定义的实现手段叫做“导入注释”。导入注释的写法示例如下：

```go
package analyzer // import "hypermind.cn/talon/analyzer"
```

代码包`analyzer`实际上属于我的一个网络爬虫项目。这个项目的代码被托管在了Github网站上。它的网址是：https://github.com/hyper-carrot/talon。如果用标准的导入路径来下载`analyzer`代码包的话，命令应该这样写`go get github.com/hyper-carrot/talon/analyzer`。不过，如果我们像上面的示例那样在该代码包中的一个源码文件中加入导入注释的话，这样下载它就行不通了。我们来看一看这个导入注释。

导入注释的写法如同一条代码包导入语句。不同的是，它出现在了单行注释符`//`的右边，因此Go语言编译器会忽略掉它。另外，它必须出现在源码文件的第一行语句（也就是代码包声明语句）的右边。只有符合上述这两个位置条件的导入注释才是有效的。再来看其中的引号部分。被双引号包裹的应该是一个符合导入路径语法规则的字符串。其中，`hypermind.cn`是我自己的一个域名。实际上，这也是用来替换掉我想隐去的代码托管网站域名及部分路径（这里是`github.com/hyper-carrot`）的那部分。在`hypermind.cn`右边的依次是我的项目的名称以及要下载的那个代码包的相对路径。这些与其标准导入路径中的内容都是一致的。为了清晰起见，我们再来做下对比。

```go
github.com/hyper-carrot/talon/analyzer // 标准的导入路径
hypermind.cn           /talon/analyzer // 导入注释中的导入路径                   
```

你想用你自己的域名替换掉标准导入路径中的哪部分由你自己说了算。不过一般情况下，被替换的部分包括代码托管网站的域名以及你在那里的用户ID就可以了。这足以达到我们最开始说的那两个目的。

虽然我们在talon项目中的所有代码包中都加入了类似的导入注释，但是我们依然无法通过`go get hypermind.cn/talon/analyzer`命令来下载这个代码包。因为域名`hypermind.cn`所指向的网站并没有加入相应的处理逻辑。具体的实现步骤应该是这样的：

1. 编写一个可处理HTTP请求的程序。这里无所谓用什么编程语言去实现。当然，我推荐你用Go语言去做。

2. 将这个处理程序与`hypermind.cn/talon`这个路径关联在一起，并总是在作为响应的HTML文档的头中写入下面这行内容：

   ```html
   <meta name="go-import" content="hypermind.cn/talon git https://github.com/hyper-carrot/talon">
   ```

   hypermind.cn/talon/analyzer熟悉HTML的读者都应该知道，这行内容会被视为HTML文档的元数据。它实际上`go get`命令的文档中要求的写法。它的模式是这样的：

```html
<meta name="go-import" content="import-prefix vcs repo-root">
```

实际上，`content`属性中的`import-prefix`的位置上应该填入我们自定义的远程代码包导入路径的前缀。这个前缀应该与我们的处理程序关联的那个路径相一致。而`vsc`显然应该代表与版本控制系统有关的标识。还记得表0-2中的主命令列吗？这里的填入内容就应该该列中的某一项。在这里，由于talon项目使用的是Git，所以这里应该填入`git`。至于`repo-root`，它应该是与该处理程序关联的路径对应的Github网站的URL。在这里，这个路径是`hypermind.cn/talon`，那么这个URL就应该是`https://github.com/hyper-carrot/talon`。后者也是talon项目的实际网址。

好了，在我们做好上述处理程序之后，`go get hypermind.cn/talon/analyzer`命令的执行结果就会是正确的。`analyzer`代码包及其依赖包中的代码会被下载到GOPATH环境变量中的第一个工作区目录的src子目录中，然后被编译并安装。

注意，具体的代码包源码存放路径会是/home/hc/golang/lib/src/hypermind.cn/talon/analyzer。也就是说，存放路径（包括代码包源码文件以及相应的归档文件的存放路径）会遵循导入注释中的路径（这里是`hypermind.cn/talon/analyzer`），而不是原始的导入路径（这里是`github.com/hyper-carrot/talon/analyzer`）。另外，我们只需在talon项目的每个代码包中的某一个源码文件中加入导入注释，但这些导入注释中的路径都必须是一致的。在这之后，我们就只能使用`hypermind.cn/talon/`作为talon项目中的代码包的导入路径前缀了。一个反例如下：

```go
hc@ubt:~$ go get github.com/hyper-carrot/talon/analyzer
package github.com/hyper-carrot/talon/analyzer: code in directory /home/hc/golang/lib/src/github.com/hyper-carrot/talon/analyzer expects import "hypermind.cn/talon/analyzer"
```

与自定义的代码包远程导入路径有关的内容我们就介绍到这里。从中我们也可以看出，Go语言为了让使用者的项目与代码托管网站隔离所作出的努力。只要你有自己的网站和一个不错的域名，这就很容易搞定并且非常值得。这会在你的代码包的使用者面前强化你的品牌，而不是某个代码托管网站的。当然，使你的代码包导入路径整齐划一是最直接的好处。

OK，言归正传，我下面继续关注`go get`这个命令本身。

**命令特有标记**

`go get`命令可以接受所有可用于`go build`命令和`go install`命令的标记。这是因为`go get`命令的内部步骤中完全包含了编译和安装这两个动作。另外，`go get`命令还有一些特有的标记，如下表所示：

*表0-4 `go get`命令的特有标记说明*

| 标记名称  | 标记描述                                                     |
| :-------- | :----------------------------------------------------------- |
| -d        | 让命令程序只执行下载动作，而不执行安装动作。                 |
| -f        | 仅在使用`-u`标记时才有效。该标记会让命令程序忽略掉对已下载代码包的导入路径的检查。如果下载并安装的代码包所属的项目是你从别人那里Fork过来的，那么这样做就尤为重要了。 |
| -fix      | 让命令程序在下载代码包后先执行修正动作，而后再进行编译和安装。 |
| -insecure | 允许命令程序使用非安全的scheme（如HTTP）去下载指定的代码包。如果你用的代码仓库（如公司内部的Gitlab）没有HTTPS支持，可以添加此标记。请在确定安全的情况下使用它。 |
| -t        | 让命令程序同时下载并安装指定的代码包中的测试源码文件中依赖的代码包。 |
| -u        | 让命令利用网络来更新已有代码包及其依赖包。默认情况下，该命令只会从网络上下载本地不存在的代码包，而不会更新已有的代码包。 |

为了更好的理解这几个特有标记，我们先清除Lib工作区的src目录和pkg目录中的所有子目录和文件。现在我们使用带有`-d`标记的`go get`命令来下载同样的代码包：

```go
hc@ubt:~$ go get -d github.com/hyper-carrot/go_lib/logging
```

现在，让我们再来看一下Lib工作区的目录结构：

```go
/home/hc/golang/lib:
    bin/
    pkg/
    src/
        github.com/
        hyper-carrot/
        go_lib/
            logging/
    ...
```

我们可以看到，`go get`命令只将代码包下载到了Lib工作区的src目录，而没有进行后续的编译和安装动作。这个加入`-d`标记的结果。

再来看`-fix`标记。我们知道，绝大多数计算机编程语言在进行升级和演进过程中，不可能保证100%的向后兼容（Backward Compatibility）。在计算机世界中，向后兼容是指在一个程序或者代码库在更新到较新的版本后，用旧的版本程序创建的软件和系统仍能被正常操作或使用，或在旧版本的代码库的基础上编写的程序仍能正常编译运行的能力。Go语言的开发者们已想到了这点，并提供了官方的代码升级工具——`fix`。`fix`工具可以修复因Go语言规范变更而造成的语法级别的错误。关于`fix`工具，我们将放在本节的稍后位置予以说明。

假设我们本机安装的Go语言版本是1.5，但我们的程序需要用到一个很早之前用Go语言的0.9版本开发的代码包。那么我们在使用`go get`命令的时候可以加入`-fix`标记。这个标记的作用是在检出代码包之后，先对该代码包中不符合Go语言1.5版本的语言规范的语法进行修正，然后再下载它的依赖包，最后再对它们进行编译和安装。

标记`-u`的意图和执行的动作都比较简单。我们在执行`go get`命令时加入`-u`标记就意味着，如果在本地工作区中已存在相关的代码包，那么就是用对应的代码版本控制系统的更新命令更新它，并进行编译和安装。这相当于强行更新指定的代码包及其依赖包。我们来看如下示例：

```go
hc@ubt:~$ go get -v github.com/hyper-carrot/go_lib/logging 
```

因为我们在之前已经检出并安装了这个代码包，所以我们执行上面这条命令后什么也没发生。还记得加入标记`-v`标记意味着会打印出被构建的代码包的名字吗？现在我们使用标记`-u`来强行更新代码包：

```go
hc@ubt:~$ go get -v -u  github.com/hyper-carrot/go_lib/logging
github.com/hyper-carrot/go_lib (download)
```

其中，“(download)”后缀意味着命令从远程仓库检出或更新了该行显示的代码包。如果我们要查看附带`-u`的`go get`命令到底做了些什么，还可以加上一个`-x`标记，以打印出用到的命令。读者可以自己试用一下它。

**智能的下载**

命令`go get`还有一个很值得称道的功能。在使用它检出或更新代码包之后，它会寻找与本地已安装Go语言的版本号相对应的标签（tag）或分支（branch）。比如，本机安装Go语言的版本是1.x，那么`go get`命令会在该代码包的远程仓库中寻找名为“go1”的标签或者分支。如果找到指定的标签或者分支，则将本地代码包的版本切换到此标签或者分支。如果没有找到指定的标签或者分支，则将本地代码包的版本切换到主干的最新版本。

前面我们说在执行`go get`命令时也可以加入`-x`标记，这样可以看到`go get`命令执行过程中所使用的所有命令。不知道读者是否已经自己尝试了。下面我们还是以代码包`github.com/hyper-carrot/go_lib`为例，并且通过之前示例中的命令的执行此代码包已经被检出到本地。这时我们再次更新这个代码包：

```go
hc@ubt:~$ go get -v -u -x github.com/hyper-carrot/go_lib
github.com/hyper-carrot/go_lib (download)
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git fetch
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git show-ref
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git checkout origin/master
WORK=/tmp/go-build034263530
```

在上述示例中，`go get`命令通过`git fetch`命令将所有远程分支更新到本地，而后有用`git show-ref`命令列出本地和远程仓库中记录的代码包的所有分支和标签。最后，当确定没有名为“go1”的标签或者分支后，`go get`命令使用`git checkout origin/master`命令将代码包的版本切换到主干的最新版本。下面，我们在本地增加一个名为“go1”的标签，看看`go get`命令的执行过程又会发生什么改变：

```go
hc@ubt:~$ cd ~/golang/lib/src/github.com/hyper-carrot/go_lib
hc@ubt:~/golang/lib/src/github.com/hyper-carrot/go_lib$ git tag go1
hc@ubt:~$ go get -v -u -x github.com/hyper-carrot/go_lib
github.com/hyper-carrot/go_lib (download)
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git fetch
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git show-ref
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git show-ref tags/go1 origin/go1
cd /home/hc/golang/lib/src/github.com/hyper-carrot/go_lib
git checkout tags/go1
WORK=/tmp/go-build636338114
```

将这两个示例进行对比，我们会很容易发现它们之间的区别。第二个示例的命令执行过程中使用`git show-ref`查看所有分支和标签，当发现有匹配的信息又通过`git show-ref tags/go1 origin/go1`命令进行精确查找，在确认无误后将本地代码包的版本切换到标签“go1”之上。

命令`go get`的这一功能是非常有用的。我们的代码在直接或间接依赖某些同时针对多个Go语言版本开发的代码包时，可以自动的检出其正确的版本。也可以说，`go get`命令内置了一定的代码包多版本依赖管理的功能。

到这里，我向大家介绍了`go get`命令的使用方式。`go get`命令与之前介绍的两个命令一样，是我们编写Go语言程序、构建Go语言项目时必不可少的辅助工具。

### install

命令`go install`用于编译并安装指定的代码包及它们的依赖包。当指定的代码包的依赖包还没有被编译和安装时，该命令会先去处理依赖包。与`go build`命令一样，传给`go install`命令的代码包参数应该以导入路径的形式提供。并且，`go build`命令的绝大多数标记也都可以用于`go install`命令。实际上，`go install`命令只比`go build`命令多做了一件事，即：安装编译后的结果文件到指定目录。

在对`go install`命令进行详细说明之前，让我们先回顾一下goc2p的目录结构。为了节省篇幅，我在这里隐藏了代码包中的源码文件。如下：

```go
$HOME/golang/goc2p:
    bin/
    pkg/
    src/
        cnet/
        logging/
        helper/
            ds/
        pkgtool/
```

我们看到，goc2p项目中有三个子目录，分别是bin目录、pkg目录和src目录。现在只有src目录中包含了一些目录，而其他两个目录都是空的。

现在，我们来看看安装代码包的规则。

**安装代码包**

如果`go install`命令后跟的代码包中仅包含库源码文件，那么`go install`命令会把编译后的结果文件保存在源码文件所在工作区的pkg目录下。对于仅包含库源码文件的代码包来说，这个结果文件就是对应的代码包归档文件（也叫静态链接库文件，名称以.a结尾）。相比之下，我们在使用`go build`命令对仅包含库源码文件的代码包进行编译时，是不会在当前工作区的src目录以及pkg目录下产生任何结果文件的。结果文件会出于编译的目的被生成在临时目录中，但并不会使当前工作区目录产生任何变化。

如果我们在执行`go install`命令时不后跟任何代码包参数，那么命令将试图编译当前目录所对应的代码包。比如，我们现在要安装代码包`pkgtool`：

```go
hc@ubt:~/golang/goc2p/src/pkgtool$ go install -v -work
WORK=D:\cygwin\tmp\go-build758586887
pkgtool
```

我们在前面说过，执行`go install`命令后会对指定代码包先编译再安装。其中，编译代码包使用了与`go build`命令相同的程序。所以，执行`go install`命令后也会首先建立一个名称以go-build为前缀的临时目录。如果我们想强行重新安装指定代码包及其依赖包，那么就需要加入标记`-a`:

```go
hc@ubt:~/golang/goc2p/src/pkgtool$ go install -a -v -work
WORK=/tmp/go-build014992994
runtime
errors
sync/atomic
unicode
unicode/utf8
sort
sync
io
syscall
strings
bytes
bufio
time
os
path/filepath
pkgtool
```

可以看到，代码包`pkgtool`仅仅依赖了Go语言标准库中的代码包。

现在我们再来查看一下goc2p项目目录：

```go
$HOME/golang/goc2p:
    bin/
    pkg/
        linux_386/
            pkgtool.a
        src/
```

现在pkg目录中多了一个子目录。读过0.0节的读者应该已经知道，linux*386被叫做平台相关目录。它的名字可以由`${GOOS}*${GOARCH}`来得到。其中，`${GOOS}`和`${GOARCH}`分别是当前操作系统中的环境变量GOOS和GOARCH的值。如果它们不存在，那么Go语言就会使用其内部的预定值。上述示例在计算架构为386且操作系统为Linux的计算机上运行。所以，这里的平台相关目录即为linux_386。我们还看到，在goc2p项目中的平台相关目录下存在一个文件，名称是pkgtool.a。这就是代码包`pkgtool`的归档文件，文件名称是由代码包名称与“.a”后缀组合而来的。

实际上，代码包的归档文件并不都会被直接保存在pkg目录的平台相关目录下，还可能被保存在这个平台相关目录的子目录下。 下面我们来安装`cnet/ctcp`包：

```go
hc@ubt:~/golang/goc2p/src/pkgtool$ go install -a -v -work ../cnet/ctcp
WORK=/tmp/go-build083178213
runtime
errors
sync/atomic
unicode
unicode/utf8
math
sync
sort
io
syscall
internal/singleflight
bytes
strings
strconv
bufio
math/rand
time
reflect
os
fmt
log
runtime/cgo
logging
net
cnet/ctcp
```

请注意，我们是在代码包`pkgtool`对应的目录下安装`cnet/ctcp`包的。我们使用了一个目录相对路径。

实际上，这种提供代码包位置的方式被叫做本地代码包路径方式，也是被所有Go命令接受的一种方式，这包括之前已经介绍过的`go build`命令。但是需要注意的是，本地代码包路径只能以目录相对路径的形式呈现，而不能使用目录绝对路径。请看下面的示例：

```go
hc@ubt:~/golang/goc2p/src/cnet/ctcp$ go install -v -work ~/golang/goc2p/src/cnet/ctcp
can't load package: package /home/hc/golang/goc2p/src/cnet/ctcp: import "/home/hc/golang/goc2p/src/cnet/ctcp": cannot import absolute path
```

从上述示例中的命令提示信息我们可知，以目录绝对路径的形式提供代码包位置是不会被Go命令认可的。

这是由于Go认为本地代码包路径的表示只能以“./”或“../”开始，再或者直接为“.”或“..”，而代码包的代码导入路径又不允许以“/”开始。所以，这种用绝对路径表示代码包位置的方式也就不能被支持了。

上述规则适用于所有Go命令。读者可以自己尝试一下，比如在执行`go build`命令时分别以代码包导入路径、目录相对路径和目录绝对路径的形式提供代码包位置，并查看执行结果。

我们已经通过上面的示例强行的重新安装了`cnet/ctcp`包及其依赖包。现在我们再来看一下goc2p的项目目录：

```go
$HOME/golang/goc2p:
    bin/
    pkg/
        linux_386/
            /cnet
                ctcp.a
            logging.a
            pkgtool.a
    src/
```

我们发现在pkg目录的平台相关目录下多了一个名为cnet的目录，而在这个目录下的就是名为ctcp.a的代码包归档文件。由此我们可知，代码包归档文件的存放目录的相对路径（相对于当前工作区的pkg目录的平台相关目录）即为代码包导入路径除去最后一个元素后的路径。而代码包归档文件的名称即为代码包导入路径中的最后一个元素再加“.a”后缀。再举一个例子，如果代码包导入路径为x/y/z，则它的归档文件存放路径的相对路径即为x/y/，而这个归档文件的名称即为z.a。

回顾代码包`pkgtool`的归档文件的存放路径。因为它的导入路径中只有一个元素，所以其归档文件就被直接存放到了goc2p项目的pkg目录的平台相关目录下了。

此外，我们还发现pkg目录的平台相关目录下还有一个名为logging.a的文件。很显然，我们并没有显式的安装代码包`logging`。这是因为`go install`命令在安装指定的代码包之前，会先去安装指定代码包的依赖包。当依赖包被正确安装后，指定的代码包的安装才会开始。由于代码包`cnet/ctcp`依赖于goc2p项目（即当前工作区）中的代码包`logging`，所以当代码包`logging`被成功安装之后，代码包`cnet/ctcp`才会被安装。

还有一个问题：上述的安装过程涉及到了那么多代码包，那为什么goc2p项目的pkg目录中只包含该项目中代码包的归档文件呢？实际上，`go install`命令会把标准库中的代码包的归档文件存放到Go语言安装目录的pkg子目录中，而把指定代码包依赖的第三方项目的代码包的归档文件存放到当前工作区的pkg目录下。这样就实现了Go语言标准库代码包的归档文件与用户代码包的归档文件，以及处在不同工作区的用户代码包的归档文件之间的分离。

**安装命令源码文件**

除了安装代码包之外，`go install`命令还可以安装命令源码文件。为了看到安装命令源码文件是goc2p项目目录的变化，我们先把该目录还原到原始状态，即清除bin子目录和pkg子目录下的所有目录和文件。然后，我们来安装代码包`helper/ds`下的命令源码文件showds.go，如下：

```go
hc@ubt:~/golang/goc2p/src$ go install helper/ds/showds.go
go install: no install location for .go files listed on command line (GOBIN not set)
```

这次我们没能成功安装。该Go命令认为目录/home/hc/golang/goc2p/src/helper/ds不在环境GOPATH中。我们可以通过Linux的`echo`命令来查看一下环境变量GOPATH的值：

```go
hc@ubt:~/golang/goc2p/src$ echo $GOPATH
/home/hc/golang/lib:/home/hc/golang/goc2p
```

环境变量GOPATH的值中确实包含了goc2p项目的根目录。这到底是怎么回事呢？

我通过查看Go命令的源码文件找到了其根本原因。在上一小节我们提到过，在环境变量GOPATH中包含多个工作区目录路径时，我们需要在编译命令源码文件前先对环境变量GOBIN进行设置。实际上，这个环境变量所指的目录路径就是命令程序生成的结果文件的存放目录。`go install`命令会把相应的可执行文件放置到这个目录中。

由于命令`go build`在编译库源码文件后不会产生任何结果文件，所以自然也不用会在意结果文件的存放目录。在该命令编译单一的命令源码文件或者包含一个命令源码文件和多个库源码文件时，在结果文件存放目录无效的情况下会将结果文件（也就是可执行文件）存放到执行该命令时所在的目录下。因此，即使环境变量GOBIN的值无效，我们在执行`go build`命令时也不会见到这个错误提示信息。

然而，`go install`命令中一个很重要的步骤就是将结果文件（归档文件或者可执行文件）存放到相应的目录中。所以，`go install`命令在安装命令源码文件时，如果环境变量GOBIN的值无效，则它会在最后检查结果文件存放目录的时候发现这一问题，并打印与上述示例所示内容类似的错误提示信息，最后直接退出。

这个错误提示信息在我们安装多个库源码文件时也有可能遇到。示例如下：

```go
hc@ubt:~/golang/goc2p/src/pkgtool$ go install envir.go fpath.go ipath.go pnode.go util.go
go install: no install location for .go files listed on command line (GOBIN not set)
```

而且，在我们为环境变量GOBIN设置了正确的值之后，这个错误提示信息仍然会出现。这是因为，只有在安装命令源码文件的时候，命令程序才会将环境变量GOBIN的值作为结果文件的存放目录。而在安装库源码文件时，在命令程序内部的代表结果文件存放目录路径的那个变量不会被赋值。最后，命令程序会发现它依然是个无效的空值。所以，命令程序会同样返回一个关于“无安装位置”的错误。这就引出一个结论，我们只能使用安装代码包的方式来安装库源码文件，而不能在`go install`命令罗列并安装它们。另外，`go install`命令目前无法接受标记`-o`以自定义结果文件的存放位置。这也从侧面说明了`go install`命令不支持针对库源码文件的安装操作。

至此，我们对怎样用`go install`命令来安装代码包以及命令源码文件进行了说明。如果你已经熟知了`go build`命令，那么理解这些内容应该不在话下。



### list

`go list`命令的作用是列出指定的代码包的信息。与其他命令相同，我们需要以代码包导入路径的方式给定代码包。被给定的代码包可以有多个。这些代码包对应的目录中必须直接保存有Go语言源码文件，其子目录中的文件不算在内。否则，代码包将被看做是不完整的。现在我们来试用一下：

```bash
hc@ubt:~$ go list cnet/ctcp pkgtoolcnet/ctcppkgtool
```

我们看到，在不加任何标记的情况下，命令的结果信息中只包含了我们指定的代码包的导入路径。我们刚刚提到，作为参数的代码包必须是完整的代码包。例如：

```bash
hc@ubt:~$ go list cnet pkgtool
can't load package: package cnet: no buildable Go source files in /home/hc/golang/goc2p/src/cnet/
pkgtool
```

这时，`go list`命令报告了一个错误——代码包`cnet`对应的目录下没有Go源码文件。但是命令还是把代码包pkgtool的导入路径打印出来了。然而，当我们在执行`go list`命令并加入标记`-e`时，即使参数中包含有不完整的代码包，命令也不会提示错误。示例如下：

```bash
hc@ubt:~$ go list -e cnet pkgtool
cnet
pkgtool
```

标记`-e`的作用是以容错模式加载和分析指定的代码包。在这种情况下，命令程序如果在加载或分析的过程中遇到错误只会在内部记录一下，而不会直接把错误信息打印出来。我们为了看到错误信息可以使用`-json`标记。这个标记的作用是把代码包的结构体实例用JSON的样式打印出来。

这里解释一下，JSON的全称是Javascript Object Notation。它一种轻量级的承载数据的格式。JSON的优势在于语法简单、短小精悍，且非常易于处理。JSON还是一种纯文本格式，独立于编程语言。正因为如此，得到了绝大多数编程语言和浏览器的支持，应用非常广泛。Go语言当然也不例外，在它的标准库中有专门用于处理和转换JSON格式的数据的代码包`encoding/json`。关于JSON格式的具体内容，读者可以去它的[官方网站](http://www.json.org/)查看说明。

在了解了这些基本概念之后，我们来试用一下`-json`标记。示例如下：

```bash
hc@ubt:~$ go list -e -json cnet
    {
            "Dir": "/home/hc/golang/goc2p/src/cnet",
            "ImportPath": "cnet",
            "Stale": true,
            "Root": "/home/hc/golang/goc2p",
            "Incomplete": true,
            "Error": {
                    "ImportStack": [
                            "cnet"
                    ],
                    "Pos": "",
                    "Err": "no Go source files in /home/hc/golang/goc2p/src/cnet"
            }
    }
```

在上述JSON格式的代码包信息中，对于结构体中的字段的显示是不完整的。因为命令程序认为我们指定`cnet`就是不完整的。在名为`Error`的字段中，我们可以看到具体说明。`Error`字段的内容其实也是一个结构体。在JSON格式下，这种嵌套的结构体被完美的展现了出来。`Error`字段所指代的结构体实例的`Err`字段说明了`cnet`不完整的原因。这与我们在没有使用`-e`标记的情况下所打印出来的错误提示信息是一致的。我们再来看`Incomplete`字段。它的值为`true`。这同样说明`cnet`是一个不完整的代码包。

实际上，在从这个代码包结构体实例到JSON格式文本的转换过程中，所有的值为其类型的空值的字段都已经被忽略了。

现在我们使用带`-json`标记的`go list`命令列出代码包`cnet/ctcp`的信息：

```bash
hc@ubt:~$ go list -json cnet/ctcp
{
    "Dir": "/home/hc/golang/github/goc2p/src/cnet/ctcp",
    "ImportPath": "cnet/ctcp",
    "Name": "ctcp",
    "Target": "/home/hc/golang/github/goc2p/pkg/darwin_amd64/cnet/ctcp.a",
    "Stale": true,
    "Root": "/home/hc/golang/github/goc2p",
    "GoFiles": [
        "base.go",
        "tcp.go"
    ],
    "Imports": [
        "bufio",
        "bytes",
        "errors",
        "logging",
        "net",
        "sync",
        "time"
    ],
    "Deps": [
        "bufio",
        "bytes",
        "errors",
        "fmt",
        "internal/singleflight",
        "io",
        "log",
        "logging",
        "math",
        "math/rand",
        "net",
        "os",
        "reflect",
        "runtime",
        "runtime/cgo",
        "sort",
        "strconv",
        "strings",
        "sync",
        "sync/atomic",
        "syscall",
        "time",
        "unicode",
        "unicode/utf8",
        "unsafe"
    ],
    "TestGoFiles": [
        "tcp_test.go"
    ],
    "TestImports": [
        "bytes",
        "fmt",
        "net",
        "runtime",
        "strings",
        "sync",
        "testing",
        "time"
    ]
}
```

由于`cnet/ctcp`包是一个完整有效的代码包，所以我们不使用`-e`标记也是没有问题的。在上面打印的`cnet/ctcp`包的信息中没有`Incomplete`字段。这是因为完整的代码包中的`Incomplete`字段的其类型的空值`false`。它已经在转换过程中被忽略掉了。另外，在`cnet/ctcp`包的信息中我们看到了很多其它的字段。现在我就来看看在Go命令程序中的代码包结构体都有哪些公开的字段。如下表。

表0-7 代码包结构体中的基本字段

| 字段名称      | 字段类型         | 字段描述                                       |
| :------------ | :--------------- | :--------------------------------------------- |
| Dir           | 字符串（string） | 代码包对应的目录。                             |
| ImportPath    | 字符串（string） | 代码包的导入路径。                             |
| ImportComment | 字符串（string） | 代码包声明语句右边的用于自定义导入路径的注释。 |
| Name          | 字符串（string） | 代码包的名称。                                 |
| Doc           | 字符串（string） | 代码包的文档字符串。                           |
| Target        | 字符串（string） | 代码包的安装路径。                             |
| Shlib         | 字符串（string） | 包含该代码包的共享库（shared library）的名称。 |
| Goroot        | 布尔（bool）     | 该代码包是否在Go语言安装目录下。               |
| Standard      | 布尔（bool）     | 该代码包是否属于标准库的一部分。               |
| Stale         | 布尔（bool）     | 该代码包的最新代码是否未被安装。               |
| Root          | 字符串（string） | 该代码包所属的工作区或Go安装目录的路径。       |

表0-8 代码包结构体中与源码文件有关的字段

| 字段名称       | 字段类型               | 字段描述                                                     |
| :------------- | :--------------------- | :----------------------------------------------------------- |
| GoFiles        | 字符串切片（[]string） | Go源码文件的列表。不包含导入了代码包“C”的源码文件和测试源码文件。 |
| CgoFiles       | 字符串切片（[]string） | 导入了代码包“C”的源码文件的列表。                            |
| IgnoredGoFiles | 字符串切片（[]string） | 忽略编译的源码文件的列表。                                   |
| CFiles         | 字符串切片（[]string） | 名称中有“.c”后缀的源码文件的列表。                           |
| CXXFiles       | 字符串切片（[]string） | 名称中有“.cc”、“.cxx”或“.cpp”后缀的源码文件的列表。          |
| MFiles         | 字符串切片（[]string） | 名称中“.m”后缀的源码文件的列表。                             |
| HFiles         | 字符串切片（[]string） | 名称中有“.h”后缀的源码文件的列表。                           |
| SFiles         | 字符串切片（[]string） | 名称中有“.s”后缀的源码文件的列表。                           |
| SwigFiles      | 字符串切片（[]string） | 名称中有“.swig”后缀的文件的列表。                            |
| SwigCXXFiles   | 字符串切片（[]string） | 名称中有“.swigcxx”后缀的文件的列表。                         |
| SysoFiles      | 字符串切片（[]string） | 名称中有“.syso”后缀的文件的列表。这些文件是需要被加入到归档文件中的。 |

表0-9 代码包结构体中与Cgo指令有关的字段

| 字段名称     | 字段类型               | 字段描述                                   |
| :----------- | :--------------------- | :----------------------------------------- |
| CgoCFLAGS    | 字符串切片（[]string） | 需要传递给C编译器的标记的列表。针对Cgo。   |
| CgoCPPFLAGS  | 字符串切片（[]string） | 需要传递给C预处理器的标记的列表。针对Cgo。 |
| CgoCXXFLAGS  | 字符串切片（[]string） | 需要传递给C++编译器的标记的列表。针对Cgo。 |
| CgoLDFLAGS   | 字符串切片（[]string） | 需要传递给链接器的标记的列表。针对Cgo。    |
| CgoPkgConfig | 字符串切片（[]string） | pkg-config的名称的列表。针对Cgo。          |

表0-10 代码包结构体中与依赖信息有关的字段

| 字段名称 | 字段类型               | 字段描述                                               |
| :------- | :--------------------- | :----------------------------------------------------- |
| Imports  | 字符串切片（[]string） | 该代码包中的源码文件显式导入的依赖包的导入路径的列表。 |
| Deps     | 字符串切片（[]string） | 所有的依赖包（包括间接依赖）的导入路径的列表。         |

表0-11 代码包结构体中与错误信息有关的字段

| 字段名称   | 字段类型            | 字段描述                                                     |
| :--------- | :------------------ | :----------------------------------------------------------- |
| Incomplete | 布尔（bool）        | 代码包是否是完整的，也即在载入或分析代码包及其依赖包时是否有错误发生。 |
| Error      | *PackageError类型   | 载入或分析代码包时发生的错误。                               |
| DepsErrors | []*PackageError类型 | 载入或分析依赖包时发生的错误。                               |

表0-12 代码包结构体中与测试源码文件有关的字段

| 字段名称     | 字段类型               | 字段描述                                                     |
| :----------- | :--------------------- | :----------------------------------------------------------- |
| TestGoFiles  | 字符串切片（[]string） | 代码包中的测试源码文件的名称列表。                           |
| TestImports  | 字符串切片（[]string） | 代码包中的测试源码文件显示导入的依赖包的导入路径的列表。     |
| XTestGoFiles | 字符串切片（[]string） | 代码包中的外部测试源码文件的名称列表。                       |
| XTestImports | 字符串切片（[]string） | 代码包中的外部测试源码文件显示导入的依赖包的导入路径的列表。 |

代码包结构体中定义的字段很多，但有些时候我们只需要查看其中的一些字段。那要怎么做呢？标记`-f`可以满足这个需求。比如这样：

```bash
hc@ubt:~$ go list -f {{.ImportPath}} cnet/ctcp
cnet/ctcp
```

实际上，`-f`标记的默认值就是`{{.ImportPath}}`。这也正是我们在使用不加任何标记的`go list`命令时依然能看到指定代码包的导入路径的原因了。

标记`-f`的值需要满足标准库的代码包``text/template`中定义的语法。比如，`{{.S}}`代表根结构体的`S`字段的值。在`go list`命令的场景下，这个根结构体就是指定的代码包所对应的结构体。如果`S`字段的值也是一个结构体的话，那么`{{.S.F}}`就代表根结构体的`S`字段的值中的`F`字段的值。如果我们要查看`cnet/ctcp`包中的命令源码文件和库源码文件的列表，可以这样使用`-f`标记：

```bash
hc@ubt:~$ go list -f {{.GoFiles}} cnet/ctcp
[base.go tcp.go]
```

如果我们想查看不完整的代码包`cnet`的错误提示信息，还可以这样：

```bash
hc@ubt:~$ go list -e -f {{.Error.Err}} cnet
no buildable Go source files in /home/hc/golang/goc2p/src/cnet
```

我们还可以利用代码包`text/template`中定义的强大语法让`go list`命令输出定制化更高的代码包信息。比如：

~~~bash
hc@ubt:~$ go list -e -f 'The package {{.ImportPath}} is {{if .Incomplete}}incomplete!{{else}}complete.{{end}}' cnet
The package cnet is incomplete!

```bash 
hc@ubt:~$ go list -f 'The imports of package {{.ImportPath}} is [{{join .Imports ", "}}].' cnet/ctcp
The imports of package cnet/ctcp is [bufio, bytes, errors, logging, net, sync, time].
~~~

其中，`join`是命令程序在`text/template`包原有语法之上自定义的语法，在底层使用标准库代码包`strings`中的`Join`函数。关于更多的语法规则，请读者查看代码包`text/template`的相关文档。

另外，`-tags`标记也可以被`go list`接受。它与我们在讲`go build`命令时提到的`-tags`标记是一致的。读者可以查看代码包`go/build``的文档以了解细节。

`go list`命令很有用。它可以为我们提供指定代码包的更深层次的信息。这些信息往往是我们无法从源码文件中直观看到的。

### mod

Go语言从v1.5开始开始引入`vendor`模式，如果项目目录下有vendor目录，那么go工具链会优先使用`vendor`内的包进行编译、测试等。godep是一个通过vender模式实现的Go语言的第三方依赖管理工具，类似的还有由社区维护准官方包管理工具dep

`go module`是Go1.11版本后官方推出的版本管理工具，从Go1.13版本开始，`go module`作为Go默认的依赖管理工具

| 命令     | 说明                                                         |
| -------- | ------------------------------------------------------------ |
| download | download modules to local cache(下载依赖包，默认为$GOPATH/pkg/mod) |
| edit     | edit go.mod from tools or scripts（编辑go.mod)               |
| graph    | print module requirement graph (打印模块依赖图)              |
| init     | initialize new module in current directory（初始化当前目录为mod） |
| tidy     | add missing and remove unused modules(拉取缺少的模块，移除不用的模块) |
| vendor   | make vendored copy of dependencies(将依赖复制到vendor下)     |
| verify   | verify dependencies have expected content (验证依赖是否正确） |
| why      | explain why packages or modules are needed(解释为什么需要依赖) |



```shell
go env -w GO111MODULE=on #on启用模块支持；off禁用模块支持；auto当项目在$GOPATH/src外，且项目根目录有go.mod文件时，开启模块支持。on时就无所谓在不在gopath中创建项目了

go mod init <模块名> #初始化模块，将在当前目录下生成go.mod文件

go mod edit -droprequire=<模块名> #移除模块
go help mod edit #查看go mod edit用法
```

###### go.mod

```go
module github.com/yuanya/package1 //定义本模块名，"github.com/<花名>/"作前缀防止模块名冲突

go 1.14  //golang版本

require "github.com/yuanya/package2" v0.0.0 //定义依赖包及版本
replace "github.com/yuanya/package2" => "../package2" //以相对路径导入其它模块
```

go mod支持语义化版本号，比如`go get foo@v1.2.3`，也可以跟git的分支或tag，比如`go get foo@master`，当然也可以跟git提交哈希，比如`go get foo@e3702bed2`。关于依赖的版本支持以下几种格式：

```go
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/vmihailenco/msgpack.v2 v2.9.1
gopkg.in/yaml.v2 <=v2.2.1
github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
latest
```

在国内访问golang.org/x上的包都需要翻墙，可以在go.mod中使用replace替换成github上对应的库

```go
replace (
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
)
```



### run

在《Go并发编程实战》的第二章中，我介绍了Go源码文件的分类。Go源码文件包括：命令源码文件、库源码文件和测试源码文件。其中，命令源码文件总应该属于`main`代码包，且在其中有无参数声明、无结果声明的main函数。单个命令源码文件可以被单独编译，也可以被单独安装（可能需要设置环境变量GOBIN）。当然，命令源码文件也可以被单独运行。我们想要运行命令源码文件就需要使用命令`go run`。

`go run`命令可以编译并运行命令源码文件。由于它其中包含了编译动作，因此它也可以接受所有可用于`go build`命令的标记。除了标记之外，`go run`命令只接受Go源码文件作为参数，而不接受代码包。与`go build`命令和`go install`命令一样，`go run`命令也不允许多个命令源码文件作为参数，即使它们在同一个代码包中也是如此。而原因也是一致的，多个命令源码文件会都有main函数声明。

如果命令源码文件可以接受参数，那么在使用`go run`命令运行它的时候就可以把它的参数放在它的文件名后面，像这样：

```bash
hc@ubt:~/golang/goc2p/src/helper/ds$ go run showds.go -p ~/golang/goc2p
```

在上面的示例中，我们使用`go run`命令运行命令源码文件showds.go。这个命令源码文件可以接受一个名称为“p”的参数。我们用“-p”这种形式表示“p”是一个参数名而不是参数值。它与源码文件名之间需要用空格隔开。参数值会放在参数名的后面，两者成对出现。它们之间也要用空格隔开。如果有第二个参数，那么第二个参数的参数名与第一个参数的参数值之间也要有一个空格。以此类推。

`go run`命令只能接受一个命令源码文件以及若干个库源码文件（必须同属于`main`包）作为文件参数，且不能接受测试源码文件。它在执行时会检查源码文件的类型。如果参数中有多个或者没有命令源码文件，那么`go run`命令就只会打印错误提示信息并退出，而不会继续执行。

在通过参数检查后，`go run`命令会将编译参数中的命令源码文件，并把编译后的可执行文件存放到临时工作目录中。

**编译和运行过程**

为了更直观的体现出`go run`命令中的操作步骤，我们在执行命令时加入标记`-n`，用于打印相关命令而不实际执行。现在让我们来模拟运行goc2p项目中的代码包helper/ds的命令源码文件showds.go。示例如下：

```bash
hc@ubt:~/golang/goc2p/src/helper/ds$ go run -n showds.go

#
# command-line-arguments
#

mkdir -p $WORK/command-line-arguments/_obj/
mkdir -p $WORK/command-line-arguments/_obj/exe/
cd /home/hc/golang/goc2p/src/helper/ds
/usr/local/go1.5/pkg/tool/linux_amd64/compile -o $WORK/command-line-arguments.a -trimpath $WORK -p main -complete -buildid df49387da030ad0d3bebef3f046d4013f8cb08d3 -D _/home/hc/golang/goc2p/src/helper/ds -I $WORK -pack ./showds.go
cd .
/usr/local/go1.5/pkg/tool/linux_amd64/link -o $WORK/command-line-arguments/_obj/exe/showds -L $WORK -w -extld=clang -buildmode=exe -buildid=df49387da030ad0d3bebef3f046d4013f8cb08d3 $WORK/command-line-arguments.a
$WORK/command-line-arguments/_obj/exe/showds
```

在上面的示例中并没有显示针对命令源码文件showds.go的依赖包进行编译和运行的相关打印信息。这是因为该源码文件的所有依赖包已经在之前被编译过了。

现在，我们来逐行解释这些被打印出来的信息。

以前缀“#”开始的是注释信息。我们看到信息中有三行注释信息，并在中间行出现了内容“command-line-arguments”。我们在讲`go build`命令的时候说过，编译命令在分析参数的时候如果发现第一个参数是Go源码文件而不是代码包时，会在内部生成一个名为“command-line-arguments”的虚拟代码包。所以这里的注释信息就是要告诉我们下面的几行信息是关于虚拟代码包“command-line-arguments”的。

打印信息中的“$WORK”表示临时工作目录的绝对路径。为了存放对虚拟代码包“command-line-arguments”的编译结果，命令在临时工作目录中创建了名为command-line-arguments的子目录，并在其下又创建了_obj子目录和_obj/exe子目录。

然后，命令程序使用Go语言工具目录`compile`命令对命令源码文件showds.go进行了编译，并把结果文件存放到了$WORK目录下，名为command-line-arguments.a。其中，`compile`是Go语言自带的编程工具。

在编译成功之后，命令程序使用链接命令`link`生成最终的可执行文件，并将其存于$WORK/command-line-arguments/_obj/exe/目录中。打印信息中的最后一行表示，命令运行了生成的可执行文件。

通过对这些打印出来的命令的解读，我们了解了临时工作目录的用途以和内容。

在上面的示例中，我们只是让`go run`命令打印出运行命令源码文件showds.go过程中需要执行的命令，而没有真正运行它。如果我们想真正运行命令源码文件showds.go并且想知道临时工作目录的位置，就需要去掉标记`-n`并且加上标记`-work`。当然，如果依然想看到过程中执行的命令，可以加上标记`-x`。如果读者已经看过之前我们对`go build`命令的介绍，就应该知道标记`-x`与标记`-n`一样会打印出过程执行的命令，但不同的这些命令会被真正的执行。调整这些标记之后的命令就像这样：

```bash
hc@ubt:~/golang/goc2p/src/helper/ds$ go run -x -work showds.go
```

当命令真正执行后，临时工作目录中就会出现实实在在的内容了，像这样：

```bash
/tmp/go-build604555989:
  command-line-arguments/
    _obj/
      exe/
        showds
  command-line-arguments.a
```

由于上述命令中包含了`-work`标记，所以我们可以从其输出中找到实际的工作目录（这里是/tmp/go-build604555989）。有意思的是，我们恰恰可以通过运行命令源码文件showds.go来查看这个临时工作目录的目录树：

```bash
hc@ubt:~/golang/goc2p/src/helper/ds$ go run showds.go -p /tmp/go-build604555989
```

读者可以自己试一试。

我们在前面介绍过，命令源码文件如果可以接受参数，则可以在执行`go run`命令运行这个命令源码文件时把参数名和参数值成对的追加在后面。实际上，如果在命令后追加参数，那么在最后执行生成的可执行文件的时候也会追加一致的参数。例如，如果这样执行命令：

```bash
hc@ubt:~/golang/goc2p/src/helper/ds$ go run -n showds.go -p ~/golang/goc2p
```

那么带`-x`或`-n`标记的命令程序打印的最后一个命令就是：

```bash
$WORK/command-line-arguments/_obj/exe/showds -p /home/hc/golang/goc2p
```

可见，`go run`命令会把追加到命令源码文件后面的参数原封不动的传给对应的可执行文件。

以上简要展示了一个命令源码文件从编译到运行的全过程。请记住，`go run`命令包含了两个动作：编译命令源码文件和运行对应的可执行文件。

### test

`go test`命令用于对Go语言编写的程序进行测试。这种测试是以代码包为单位的。当然，这还需要测试源码文件的帮助。关于怎样编写并写好Go程序测试代码，我们会在本章的第二节加以详述。在这里，我们只讨论怎样使用命令启动测试。

`go test`命令会自动测试每一个指定的代码包。当然，前提是指定的代码包中存在测试源码文件。关于测试源码文件方面的知识，在我的图书《Go并发编程实战》中有详细介绍。测试源码文件是名称以“_test.go”为后缀的、内含若干测试函数的源码文件。测试函数一般是以“Test”为名称前缀并有一个类型为“testing.T”的参数声明的函数.

现在，我们来测试goc2p项目中的几个代码包。在使用`go test`命令时指定代码包的方式与其他命令无异——使用代码包导入路径。如果需要测试多个代码包，则需要在它们的导入路径之间加入空格以示分隔。示例如下：

```bash
hc@ubt:~$ go test basic cnet/ctcp pkgtool
ok      basic   0.012s
ok      cnet/ctcp   2.014s
ok      pkgtool 0.014s
```

`go test`命令在执行完所有的代码包中的测试文件之后，会以代码包为单位打印出测试概要信息。在上面的示例中，对应三个代码包的三行信息的第一列都是“ok”。这说明它们都通过了测试。每行的第三列显示运行相应测试所用的时间，以秒为单位。我们还可以在代码包目录下运行不加任何参数的运行`go test`命令。其作用和结果与上面的示例是一样的。

另外，我们还可以指定测试源码文件来进行测试。这样的话，`go test`命令只会执行指定文件中的测试，像这样：

```bash
    hc@ubt:~/golang/goc2p/src/pkgtool$ go test envir_test.go
# command-line-arguments
./envir_test.go:25: undefined: GetGoroot
./envir_test.go:40: undefined: GetAllGopath
./envir_test.go:81: undefined: GetSrcDirs
./envir_test.go:83: undefined: GetAllGopath
./envir_test.go:90: undefined: GetGoroot
FAIL    command-line-arguments [build failed]
```

我们看到，与指定源码文件进行编译或运行一样，命令程序会为指定的源码文件生成一个虚拟代码包——“command-line-arguments”。但是，测试并没有通过。但其原因并不是测试失败，而是编译失败。对于运行这次测试的命令程序来说，测试源码文件envir_test.go是属于代码包“command-line-arguments”的。并且，这个测试源码文件中使用了库源码文件envir.go中的函数。但是，它却没有显示导入这个库源码文件所属的代码包。这显然会引起编译错误。如果想解决这个问题，我们还需要在执行命令时加入这个测试源码文件所测试的那个源码文件。示例如下：

```bash
hc@ubt:~/golang/goc2p/src/pkgtool$ go test envir_test.go envir.go
ok      command-line-arguments  0.010s
```

现在，我们故意使代码包`pkgtool`中的某个测试失败。现在我们再来运行测试：

```bash
hc@ubt:~$ go test basic cnet/ctcp pkgtool
ok      basic   0.010s
ok      cnet/ctcp       2.015s
--- FAIL: TestGetSrcDirs (0.00 seconds)
        envir_test.go:85: Error: The src dir '/usr/local/go/src/pkg' is incorrect.
FAIL
FAIL    pkgtool 0.009s
```

我们通过以上示例中的概要信息获知，测试源码文件中envir_test.go的测试函数`TestGetSrcDirs`中的测试失败了。在包含测试失败的测试源码文件名的那一行信息中，紧跟测试源码文件名的用冒号分隔的数字是错误信息所处的行号，在行号后面用冒号分隔的是错误信息。这个错误信息的内容是用户自行编写的。另外，概要信息的最后一行以“FAIL”为前缀。这表明针对代码包pkgtool的测试未通过。未通过的原因在前面的信息中已有描述。

一般情况下，我们会把测试源码文件与被测试的源码文件放在同一个代码包中。并且，这些源码文件中声明的包名也都是相同的。除此之外我们还有一种选择，那就是测试源码文件中声明的包名可以是所属包名再加“_test”后缀。我们把这种测试源码文件叫做包外测试源码文件。不过，包外测试源码文件存在一个弊端，那就是在它们的测试函数中无法测试被测源码文件中的包级私有的程序实体，比如包级私有的变量、函数和结构体类型。这是因为这两者的所属代码包是不相同的。所以，我们一般很少会编写包外测试源码文件。

**关于标记**

`go test`命令的标记处理部分是庞大且繁杂的，以至于使Go语言的开发者们不得不把这一部分的逻辑从`go test`命令程序主体中分离出来并建立单独的源码文件。因为`go test`命令中包含了编译动作，所以它可以接受可用于`go build`命令的所有标记。另外，它还有很多特有的标记。这些标记的用于控制命令本身的动作，有的用于控制和设置测试的过程和环境，还有的用于生成更详细的测试结果和统计信息。

可用于`go test`命令的几个比较常用的标记是`-c`、`-i`和`-o`。这两个就是用于控制`go test`命令本身的动作的标记。详见下表。

表0-6 `go test`命令的标记说明

| 标记名称 | 标记描述                                                     |
| :------- | :----------------------------------------------------------- |
| -c       | 生成用于运行测试的可执行文件，但不执行它。这个可执行文件会被命名为“pkg.test”，其中的“pkg”即为被测试代码包的导入路径的最后一个元素的名称。 |
| -i       | 安装/重新安装运行测试所需的依赖包，但不编译和运行测试代码。  |
| -o       | 指定用于运行测试的可执行文件的名称。追加该标记不会影响测试代码的运行，除非同时追加了标记`-c`或`-i`。 |

上述这几个标记可以搭配使用。搭配使用的目的可以是让`go test`命令既安装依赖包又编译测试代码，但不运行测试。也就是说，让命令程序跑一遍运行测试之前的所有流程。这可以测试一下测试过程。注意，在加入`-c`标记后，命令程序会把用于运行测试的可执行文件存放到当前目录下。

除此之外，`go test`命令还有很多功效各异的标记。但是由于这些标记的复杂性，我们需要结合测试源码文件进行详细的讲解。所以我们在这里略过不讲。如果读者想了解相关详情，请参看《Go并发编程实战》的第5章



#### -bench

性能测试



### tool

### version

### vet

https://golang.org/cmd/vet/

用于检查Go语言源码中静态错误的简单工具。开发中 idea 可以自动检测这类问题并给与一定提示，但是仍然需要在 makefile 中使用 vet 以避免我们粗心导致的错误

```go
package main

import "fmt"

func main() {
   fmt.Printf("%s", 3)
}
```

```shell
$ go vet ./...
cmd\main.go:6:2: Printf format %s has arg 3 of wrong type int
```



命令`go vet`是一个用于检查Go语言源码中静态错误的简单工具。与大多数Go命令一样，`go vet`命令可以接受`-n`标记和`-x`标记。`-n`标记用于只打印流程中执行的命令而不真正执行它们。`-x`标记也用于打印流程中执行的命令，但不会取消这些命令的执行。示例如下：

```
hc@ubt:~$ go vet -n pkgtool
/usr/local/go/pkg/tool/linux_386/vet golang/goc2p/src/pkgtool/envir.go golang/goc2p/src/pkgtool/envir_test.go golang/goc2p/src/pkgtool/fpath.go golang/goc2p/src/pkgtool/ipath.go golang/goc2p/src/pkgtool/pnode.go golang/goc2p/src/pkgtool/util.go golang/goc2p/src/pkgtool/util_test.go
```

`go vet`命令的参数既可以是代码包的导入路径，也可以是Go语言源码文件的绝对路径或相对路径。但是，这两种参数不能混用。也就是说，`go vet`命令的参数要么是一个或多个代码包导入路径，要么是一个或多个Go语言源码文件的路径。

`go vet`命令是`go tool vet`命令的简单封装。它会首先载入和分析指定的代码包，并把指定代码包中的所有Go语言源码文件和以“.s”结尾的文件的相对路径作为参数传递给`go tool vet`命令。其中，以“.s”结尾的文件是汇编语言的源码文件。如果`go vet`命令的参数是Go语言源码文件的路径，则会直接将这些参数传递给`go tool vet`命令。

如果我们直接使用`go tool vet`命令，则其参数可以传递任意目录的路径，或者任何Go语言源码文件和汇编语言源码文件的路径。路径可以是绝对的也可以是相对的。

实际上，`vet`属于Go语言自带的特殊工具，也是比较底层的命令之一。Go语言自带的特殊工具的存放路径是$GOROOT/pkg/tool/$GOOS*$GOARCH/，我们暂且称之为Go工具目录。我们再来复习一下，环境变量GOROOT的值即Go语言的安装目录，环境变量GOOS的值代表程序构建环境的目标操作系统的标识，而环境变量$GOARCH的值则为程序构建环境的目标计算架构。另外，名为$GOOS*$GOARCH的目录被叫做平台相关目录。Go语言允许我们通过执行`go tool`命令来运行这些特殊工具。在Linux 32bit的环境下，我们的Go语言安装目录是/usr/local/go/。因此，`go tool vet`命令指向的就是被存放在/usr/local/go/pkg/tool/linux_386目录下的名为`vet`的工具。

`go tool vet`命令的作用是检查Go语言源代码并且报告可疑的代码编写问题。比如，在调用`Printf`函数时没有传入格式化字符串，以及某些不标准的方法签名，等等。该命令使用试探性的手法检查错误，因此并不能保证报告的问题确实需要解决。但是，它确实能够找到一些编译器没有捕捉到的错误。

`go tool vet`命令程序在被执行后会首先解析标记并检查标记值。`go tool vet`命令支持的所有标记如下表。

*表0-16 `go tool vet`命令的标记说明*

| 标记名称            | 标记描述                                                     |
| :------------------ | :----------------------------------------------------------- |
| -all                | 进行全部检查。如果有其他检查标记被设置，则命令程序会将此值变为false。默认值为true。 |
| -asmdecl            | 对汇编语言的源码文件进行检查。默认值为false。                |
| -assign             | 检查赋值语句。默认值为false。                                |
| -atomic             | 检查代码中对代码包sync/atomic的使用是否正确。默认值为false。 |
| -buildtags          | 检查编译标签的有效性。默认值为false。                        |
| -composites         | 检查复合结构实例的初始化代码。默认值为false。                |
| -compositeWhiteList | 是否使用复合结构检查的白名单。仅供测试使用。默认值为true。   |
| -methods            | 检查那些拥有标准命名的方法的签名。默认值为false。            |
| -printf             | 检查代码中对打印函数的使用是否正确。默认值为false。          |
| -printfuncs         | 需要检查的代码中使用的打印函数的名称的列表，多个函数名称之间用英文半角逗号分隔。默认值为空字符串。 |
| -rangeloops         | 检查代码中对在```range```语句块中迭代赋值的变量的使用是否正确。默认值为false。 |
| -structtags         | 检查结构体类型的字段的标签的格式是否标准。默认值为false。    |
| -unreachable        | 查找并报告不可到达的代码。默认值为false。                    |

在阅读上面表格中的内容之后，读者可能对这些标签的具体作用及其对命令程序检查步骤的具体影响还很模糊。不过没关系，我们下面就会对它们进行逐一的说明。

**-all标记**

如果标记`-all`有效（标记值不为`false`），那么命令程序会对目标文件进行所有已知的检查。实际上，标记`-all`的默认值就是`true`。也就是说，在执行`go tool vet`命令且不加任何标记的情况下，命令程序会对目标文件进行全面的检查。但是，只要有一个另外的标记（`-compositeWhiteList`和`-printfuncs`这两个标记除外）有效，命令程序就会把标记`-all`设置为false，并只会进行与有效的标记对应的检查。

**-assign标记**

如果标记`-assign`有效（标记值不为`false`），则命令程序会对目标文件中的赋值语句进行自赋值操作检查。什么叫做自赋值呢？简单来说，就是将一个值或者实例赋值给它本身。像这样：

```
var s1 string = "S1"
s1 = s1 // 自赋值
```

或者

```
s1, s2 := "S1", "S2"
s2, s1 = s2, s1 // 自赋值
```

检查程序会同时遍历等号两边的变量或者值。在抽象语法树的语境中，它们都被叫做表达式节点。检查程序会检查等号两边对应的表达式是否相同。判断的依据是这两个表达式节点的字符串形式是否相同。在当前的场景下，这种相同意味着它们的变量名是相同的。如前面的示例。

有两种情况是可以忽略自赋值检查的。一种情况是短变量声明语句。根据Go语言的语法规则，当我们在函数中要在声明局部变量的同时对其赋值，就可以使用`:=`形式的变量赋值语句。这也就意味着`:=`左边的变量名称在当前的上下文环境中应该还未曾出现过（否则不能通过编译）。因此，在这种赋值语句中不可能出现自赋值的情况，忽略对它的检查也是合理的。另一种情况是等号左右两边的表达式个数不相等的变量赋值语句。如果在等号的右边是对某个函数或方法的调用，就会造成这种情况。比如：

```
file, err := os.Open(wp)
```

很显然，这个赋值语句肯定不是自赋值语句。因此，不需要对此种情况进行检查。如果等号右边并不是对函数或方法调用的表达式，并且等号两边的表达式数量也不相等，那么势必会在编译时引发错误，也不必检查。

**-atomic标记**

如果标记`-atomic`有效（标记值不为`false`），则命令程序会对目标文件中的使用代码包`sync/atomic`进行原子赋值的语句进行检查。原子赋值语句像这样：

```
var i32 int32
i32 = 0
newi32 := atomic.AddInt32(&i32, 3)
fmt.Printf("i32: %d, newi32: %d.\n", i32, newi32)
```

函数`AddInt32`会原子性的将变量`i32`的值加`3`，并返回这个新值。因此上面示例的打印结果是：

```
i32: 3, newi32: 3
```

在代码包`sync/atomic`中，与`AddInt32`类似的函数还有`AddInt64`、`AddUint32`、`AddUint64`和`AddUintptr`。检查程序会对上述这些函数的使用方式进行检查。检查的关注点在破坏原子性的使用方式上。比如：

```
i32 = 1
i32 = atomic.AddInt32(&i32, 3)
_, i32 = 5, atomic.AddInt32(&i32, 3)
i32, _ = atomic.AddInt32(&i32, 1), 5 
```

上面示例中的后三行赋值语句都属于原子赋值语句，但它们都破坏了原子赋值的原子性。以第二行的赋值语句为例，等号左边的`atomic.AddInt32(&i32, 3)`的作用是原子性的将变量`i32`的值增加`3`。但该语句又将函数的结果值赋值给变量`i32`，这个二次赋值属于对变量`i32`的重复赋值，也使原本拥有原子性的赋值操作被拆分为了两个步骤的非原子操作。如果在对变量`i32`的第一次原子赋值和第二次非原子的重复赋值之间又有另一个程序对变量`i32`进行了原子赋值，那么当前程序中的这个第二次赋值就破坏了那两次原子赋值本应有的顺序性。因为，在另一个程序对变量`i32`进行原子赋值后，当前程序中的第二次赋值又将变量`i32`的值设置回了之前的值。这显然是不对的。所以，上面示例中的第二行代码应该改为：

```
atomic.AddInt32(&i32, 3)
```

并且，对第三行和第四行的代码也应该有类似的修改。检查程序如果在目标文件中查找到像上面示例的第二、三、四行那样的语句，就会打印出相应的错误信息。

另外，上面所说的导致原子性被破坏的重复赋值语句还有一些类似的形式。比如：

```
i32p := &i32
*i32p = atomic.AddUint64(i32p, 1)
```

这与之前的示例中的代码的含义几乎是一样。另外还有：

```
var counter struct{ N uint32 }
counter.N = atomic.AddUint64(&counter.N, 1) 
```

和

```
ns := []uint32{10, 20}
ns[0] = atomic.AddUint32(&ns[0], 1)
nps := []*uint32{&ns[0], &ns[1]}
*nps[0] = atomic.AddUint32(nps[0], 1)
```

在最近的这两个示例中，虽然破坏原子性的重复赋值操作因结构体类型或者数组类型的介入显得并不那么直观了，但依然会被检查程序发现并及时打印错误信息。

顺便提一句，对于原子赋值语句和普通赋值语句，检查程序都会忽略掉对等号两边的表达式的个数不相等的赋值语句的检查。

**-buildtags标记**

前文已提到，如果标记`-buildtags`有效（标记值不为`false`），那么命令程序会对目标文件中的编译标签（如果有的话）的格式进行检查。什么叫做条件编译？在实际场景中，有些源码文件中包含了平台相关的代码。我们希望只在某些特定平台下才编译它们。这种有选择的编译方法就被叫做条件编译。在Go语言中，条件编译的配置就是通过编译标签来完成的。编译器需要依据源码文件中编译标签的内容来决定是否编译当前文件。编译标签可必须出现在任何源码文件（比如扩展名为“.go”，“.h”，“.c”，“.s”等的源码文件) 的头部的单行注释中，并且在其后面需要有空行。

至于编译标签的具体写法，我们就不在此赘述了。读者可以参看Go语言官方的相关文档。我们在这里只简单罗列一下`-buildtags`有效时命令程序对编译标签的检查内容：

1. 若编译标签前导符“+build”后没有紧随空格，则打印格式错误信息。
2. 若编译标签所在行与第一个多行注释或代码行之间没有空行，则打印错误信息。
3. 若在某个单一参数的前面有两个英文叹号“!!”，则打印错误信息。
4. 若单个参数包含字母、数字、“_”和“.”以外的字符，则打印错误信息。
5. 若出现在文件头部单行注释中的编译标签前导符“+build”未紧随在单行注释前导符“//”之后，则打印错误信息。

如果一个在文件头部的单行注释中的编译标签通过了上述的这些检查，则说明它的格式是正确无误的。由于只有在文件头部的单行注释中编译标签才会被编译器认可，所以检查程序只会查找和检查源码文件中的第一个多行注释或代码行之前的内容。

**-composites标记和-compositeWhiteList标记**

如果标记`-composites`有效（标记值不为`false`），则命令程序会对目标文件中的复合字面量进行检查。请看如下示例：

```
type counter struct {
    name   string
    number int
}
...
c := counter{name: "c1", number: 0}
```

在上面的示例中，代码`counter{name: "c1", number: 0}`是对结构体类型`counter`的初始化。如果复合字面量中涉及到的类型不在当前代码包内部且未在所属文件中被导入，那么检查程序不但会打印错误信息还会将退出代码设置为1，并且取消后续的检查。退出代码为1意味着检查程序已经报告了一个或多个问题。这个问题比仅仅引起错误信息报告的问题更加严重。

在通过上述检查的前提下，如果复合字面量中包含了对结构体类型的字段的赋值但却没有指明字段名，像这样：

```
var v = flag.Flag{
    "Name",
    "Usage",
    nil, // Value
    "DefValue",
}
```

那么检查程序也会打印错误信息，以提示在复合字面量中包含有未指明的字段赋值。

这有一个例外，那就是当标记`-compositeWhiteList`有效（标记值不为`false`）的时候。只要类型在白名单中，即使其初始化语句中含有未指明的字段赋值也不会被提示。这是出于什么考虑呢？先来看下面的示例：

```
type sliceType []string
...
st1 := sliceType{"1", "2", "3"}
```

上面示例中的`sliceType{"1", "2", "3"}`也属于复合字面量。但是它初始化的类型实际上是一个切片值，只不过这个切片值被别名化并被包装为了另一个类型而已。在这种情况下，复合字面量中的赋值不需要指明字段，事实上这样的类型也不包含任何字段。白名单中所包含的类型都是这种情况。它们是在标准库中的包装了切片值的类型。它们不需要被检查，因为这种情况是合理的。

在默认情况下，标记`-compositeWhiteList`是有效的。也就是说，检查程序不会对它们的初始化代码进行检查，除非我们在执行`go tool vet`命令时显示的将`-compositeWhiteList`标记的值设置为false。

**-methods标记**

如果标记`-methods`有效（标记值不为`false`），则命令程序会对目标文件中的方法定义进行规范性的进行检查。这里所说的规范性是狭义的。

在检查程序内部存有一个规范化方法字典。这个字典的键用来表示方法的名称，而字典的元素则用来描述方法应有的参数和结果的类型。在该字典中列出的都是Go语言标准库中使用最广泛的接口类型的方法。这些方法的名字都非常通用。它们中的大多数都是它们所属接口类型的唯一方法。我们在第4章中提到过，Go语言中的接口类型实现方式是非侵入式的。只要结构体类型实现了某一个接口类型中的所有方法，就可以说这个结构体类型是该接口类型的一个实现。这种判断方式被称为动态接口检查。它只在运行时进行。如果我们想让一个结构体类型成为某一个接口类型的实现，但又写错了要实现的接口类型中的方法的签名，那么也不会引发编译器报错。这里所说的方法签名包括方法的参数声明列表和结果声明列表。虽然动态接口检查失败时并不会报错，但是它却会间接的引发其它错误。而这些被间接引发的错误只会在运行时发生。示例如下：

```
type MySeeker struct {
    // 忽略字段定义
}

func (self *MySeeker) Seek(whence int, offset int64) (ret int64, err error) { 
    // 想实现接口类型io.Seeker中的唯一方法，但是却把参数的顺序写颠倒了。
    // 忽略实现代码
}

func NewMySeeker io.Seeker {
    return &MySeeker{/* 忽略字段初始化 */} // 这里会引发一个运行时错误。
                                           //由于MySeeker的Seek方法的签名写错了，所以MySeeker不是io.Seeker的实现。
}
```

这种运行时错误看起来会比较诡异，并且错误排查也会相对困难，所以应该尽量避免。`-methods`标记所对应的检查就是为了达到这个目的。检查程序在发现目标文件中某个方法的名字被包含在规范化方法字典中但其签名与对应的描述不对应的时候，就会打印错误信息并设置退出代码为1。

我在这里附上在规范化方法字典中列出的方法的信息：

*表0-17 规范化方法字典中列出的方法*

| 方法名称      | 参数类型                | 结果类型               | 所属接口         | 唯一方法 |
| :------------ | :---------------------- | :--------------------- | :--------------- | :------- |
| Format        | "fmt.State", "rune"     |                        | fmt.Formatter    | 是       |
| GobDecode     | "[]byte"                | "error"                | gob.GobDecoder   | 是       |
| GobEncode     |                         | "[]byte", "error"      | gob.GobEncoder   | 是       |
| MarshalJSON   |                         | "[]byte", "error"      | json.Marshaler   | 是       |
| Peek          | "int"                   | "[]byte", "error"      | image.reader     | 否       |
| ReadByte      | "int"                   | "[]byte", "error"      | io.ByteReader    | 是       |
| ReadFrom      | "io.Reader"             | "int64", "error"       | io.ReaderFrom    | 是       |
| ReadRune      |                         | "rune", "int", "error" | io.RuneReader    | 是       |
| Scan          | "fmt.ScanState", "rune" | "error"                | fmt.Scanner      | 是       |
| Seek          | "int64", "int"          | "int64", "error"       | io.Seeker        | 是       |
| UnmarshalJSON | "[]byte"                | "error"                | json.Unmarshaler | 是       |
| UnreadByte    |                         | "error"                | io.ByteScanner   | 否       |
| UnreadRune    |                         | "error"                | io.RuneScanner   | 否       |
| WriteByte     | "byte"                  | "error"                | io.ByteWriter    | 是       |
| WriteTo       | "io.Writer"             | "int64", "error"       | io.WriterTo      | 是       |

**-printf标记和-printfuncs标记**

标记`-printf`旨在目标文件中检查各种打印函数使用的正确性。而标记`-printfuncs`及其值则用于明确指出需要检查的打印函数。`-printfuncs`标记的默认值为空字符串。也就是说，若不明确指出检查目标则检查所有打印函数。可被检查的打印函数如下表：

*表0-18 格式化字符串中动词的格式要求*

| 函数全小写名称 | 支持格式化 | 可自定义输出 | 自带换行 |
| :------------- | :--------- | :----------- | :------- |
| error          | 否         | 否           | 是       |
| fatal          | 否         | 否           | 是       |
| fprint         | 否         | 是           | 否       |
| fprintln       | 否         | 是           | 是       |
| panic          | 否         | 否           | 否       |
| panicln        | 否         | 否           | 是       |
| print          | 否         | 否           | 否       |
| println        | 否         | 否           | 是       |
| sprint         | 否         | 否           | 否       |
| sprintln       | 否         | 否           | 是       |
| errorf         | 是         | 否           | 否       |
| fatalf         | 是         | 否           | 否       |
| fprintf        | 是         | 是           | 否       |
| panicf         | 是         | 否           | 否       |
| printf         | 是         | 否           | 否       |
| sprintf        | 是         | 是           | 否       |

以字符串格式化功能来区分，打印函数可以分为可打印格式化字符串的打印函数（以下简称格式化打印函数）和非格式化打印函数。对于格式化打印函数来说，其第一个参数必是格式化表达式，也可被称为模板字符串。而其余参数应该为需要被填入模板字符串的变量。像这样：

```
fmt.Printf("Hello, %s!\n", "Harry") 
// 会输出：Hello, Harry!
```

而非格式化打印函数的参数则是一个或多个要打印的内容。比如：

```
fmt.Println("Hello,", "Harry!") 
// 会输出：Hello, Harry!
```

以指定输出目的地功能区分，打印函数可以被分为可自定义输出目的地的的打印函数（以下简称自定义输出打印函数）和标准输出打印函数。对于自定义输出打印函数来说，其第一个函数必是其打印的输出目的地。比如：

```
fmt.Fprintf(os.Stdout, "Hello, %s!\n", "Harry")
// 会在标准输出设备上输出：Hello, Harry!
```

上面示例中的函数`fmt.Fprintf`既能够让我们自定义打印的输出目的地，又能够格式化字符串。此类打印函数的第一个参数的类型应为`io.Writer`接口类型。只要某个类型实现了该接口类型中的所有方法，就可以作为函数`Fprintf`的第一个参数。例如，我们还可以使用代码包`bytes`中的结构体`Buffer`来接收打印函数打印的内容。像这样：

```
var buff bytes.Buffer
fmt.Fprintf(&buff, "Hello, %s!\n", "Harry")
fmt.Print("Buffer content:", buff.String())
// 会在标准输出设备上输出：Buffer content: Hello, Harry!
```

而标准输出打印函数则只能将打印内容到标准输出设备上。就像函数`fmt.Printf`和`fmt.Println`所做的那样。

检查程序会首先关注打印函数的参数数量。如果参数数量不足，则可以认为在当前调用打印函数的语句中并不会出现用法错误。所以，检查程序会忽略对它的检查。检查程序中对打印函数的最小参数是这样定义的：对于可以自定义输出的打印函数来说，最小参数数量为2，其它打印函数的最小参数数量为1。如果打印函数的实际参数数量小于对应的最小参数数量，就会被判定为参数数量不足。

对于格式化打印函数，检查程序会进行如下检查：

1. 如果格式化字符串无法被转换为基本字面量（标识符以及用于表示int类型值、float类型值、char类型值、string类型值的字面量等），则检查程序会忽略剩余的检查。如果`-v`标记有效，则会在忽略检查前打印错误信息。另外，格式化打印函数的格式化字符串必须是字符串类型的。因此，如果对应位置上的参数的类型不是字符串类型，那么检查程序会立即打印错误信息，并设置退出代码为1。实际上，这个问题已经可以引起一个编译错误了。

2. 如果格式化字符串中不包含动词（verbs），而格式化字符串后又有多余的参数，则检查程序会立即打印错误信息，并设置退出代码为1，且忽略后续检查。我现在举个例子。我们拿之前的一个示例作为基础，即：

   fmt.Printf("Hello, %s!\n", "Harry")

在这个示例中，格式化字符串中的“%s”就是我们所说的动词，“%”就是动词的前导符。它相当于一个需要被填的空。一般情况下，在格式化字符串中被填的空的数量应该与后续参数的数量相同。但是可以出现在格式化字符串中没有动词并且在格式化字符串之后没有额外参数的情况。在这种情况下，该格式化打印函数就相当于一个非格式化打印函数。例如，下面这个语句会导致此步检查不通过：

```
fmt.Printf("Hello!\n", "Harry") 
```

1. 检查程序还会检查动词的格式。这部分检查会非常严格。检查程序对于格式化字符串中动词的格式要求如表0-19。表中对每个动词只进行了简要的说明。读者可以查看标准库代码包`fmt`的文档以了解关于它们的详细信息。命令程序会按照表5-19中的要求对格式化及其后续参数进行检查。如上表所示，这部分检查分为两步骤。第一个步骤是检查格式化字符串中的动词上是否附加了不合法的标记，第二个步骤是检查格式化字符串中的动词与后续对应的参数的类型是否匹配。只要检查出问题，检查程序就会打印出错误信息并且设置退出代码为1。
2. 如果格式化字符串中的动词不被支持，则检查程序同样会打印错误信息后，并设置退出代码为1。

*表0-19 格式化字符串中动词的格式要求*

| 动词 | 合法的附加标记               | 允许的参数类型    | 简要说明                                              |
| :--- | :--------------------------- | :---------------- | :---------------------------------------------------- |
| b    | “ ”，“-”，“+”，“.”和“0”      | int或float        | 用于二进制表示法。                                    |
| c    | “-”                          | rune或int         | 用于单个字符的Unicode表示法。                         |
| d    | “ ”，“-”，“+”，“.”和“0”      | int               | 用于十进制表示法。                                    |
| e    | “ ”，“-”，“+”，“.”和“0”      | float             | 用于科学记数法。                                      |
| E    | “ ”，“-”，“+”，“.”和“0”      | float             | 用于科学记数法。                                      |
| f    | “ ”，“-”，“+”，“.”和“0”      | float             | 用于控制浮点数精度。                                  |
| F    | “ ”，“-”，“+”，“.”和“0”      | float             | 用于控制浮点数精度。                                  |
| g    | “ ”，“-”，“+”，“.”和“0”      | float             | 用于压缩浮点数输出。                                  |
| G    | “ ”，“-”，“+”，“.”和“0”      | float             | 用于动态选择浮点数输出格式。                          |
| o    | “ ”，“-”，“+”，“.”，“0”和“#” | int               | 用于八进制表示法。                                    |
| p    | “-”和“#”                     | pointer           | 用于表示指针地址。                                    |
| q    | “ ”，“-”，“+”，“.”，“0”和“#” | rune，int或string | 用于生成带双引号的字符串形式的内容。                  |
| s    | “ ”，“-”，“+”，“.”和“0”      | rune，int或string | 用于生成字符串形式的内容。                            |
| t    | “-”                          | bool              | 用于生成与布尔类型对应的字符串值。（“true”或“false”） |
| T    | “-”                          | 任何类型          | 用于用Go语法表示任何值的类型。                        |
| U    | “-”和“#”                     | rune或int         | 用于针对Unicode的表示法。                             |
| v    | “”，“-”，“+”，“.”，“0”和“#”  | 任何类型          | 以默认格式格式化任何值。                              |
| x    | “”，“-”，“+”，“.”，“0”和“#”  | rune，int或string | 以十六进制、全小写的形式格式化每个字节。              |
| X    | “”，“-”，“+”，“.”，“0”和“#”  | rune，int或string | 以十六进制、全大写的形式格式化每个字节。              |

对于非格式化打印函数，检查程序会进行如下检查：

1. 如果打印函数不是可以自定义输出的打印函数，那么其第一个参数就不能是标准输出`os.Stdout`或者标准错误输出`os.Stderr`。否则，检查程序将打印错误信息并设置退出代码为1。这主要是为了防止程序编写人员的笔误。比如，他们可能会把函数`fmt.Println`当作函数`fmt.Printf`来用。

2. 如果打印函数是不自带换行的，比如`fmt.Printf`和`fmt.Print`，则它必须只少有一个参数。否则，检查程序将打印错误信息并设置退出代码为1。像这样的调用打印函数的语句是没有任何意义的。并且，如果这个打印函数还是一个格式化打印函数，那么这还会引起一个编译错误。需要注意的是，函数名称为`Error`的方法不会在被检查之列。比如，标准库代码包`testing`中的结构体类型`T`和`B`的方法`Error`。这是因为它们可能实现了接口类型`Error`。这个接口类型中唯一的方法`Error`无需任何参数。

3. 如果第一个参数的值为字符串类型的字面量且带有格式化字符串中才应该有的动词的前导符“%”，则检查程序会打印错误信息并设置退出代码为1。因为非格式化打印函数中不应该出现格式化字符串。

4. 如果打印函数是自带换行的，那么在打印内容的末尾就不应该有换行符“\n”。否则，检查程序会打印错误信息并设置退出代码为1。换句话说，检查程序认为程序中如果出现这样的代码：

   fmt.Println("Hello!\n")

常常是由于程序编写人员的笔误。实际上，事实确实如此。如果我们确实想连续输入多个换行，应该这样写：

```
fmt.Println("Hello!")fmt.Println()
```

至此，我们详细介绍了`go tool vet`命令中的检查程序对打印函数的所有步骤和内容。打印函数的功能非常简单，但是`go tool vet`命令对它的检查却很细致。从中我们可以领会到一些关于打印函数的最佳实践。

**-rangeloops标记**

如果标记`-rangeloop`有效（标记值不为`false`），那么命令程序会对使用`range`进行迭代的`for`代码块进行检查。我们之前提到过，使用`for`语句需要注意两点：

1. 不要在`go`代码块中处理在迭代过程中被赋予值的迭代变量。比如：

   mySlice := []string{"A", "B", "C"} for index, value := range mySlice { go func() { fmt.Printf("Index: %d, Value: %s\n", index, value) }() }

在Go语言的并发编程模型中，并没有线程的概念，但却有一个特有的概念——Goroutine。Goroutine也可被称为Go例程或简称为Go程。关于Goroutine的详细介绍在第6章和第7章。我们现在只需要知道它是一个可以被并发执行的代码块。

1. 不要在`defer`语句的延迟函数中处理在迭代过程中被赋予值的迭代变量。比如：

   myDict := make(map[string]int) myDict["A"] = 1 myDict["B"] = 2 myDict["C"] = 3 for key, value := range myDict { defer func() { fmt.Printf("Key: %s, Value: %d\n", key, value) }() }

其实，上述两点所关注的问题是相同的，那就是不要在可能被延迟处理的代码块中直接使用迭代变量。`go`代码块和`defer`代码块都有这样的特质。这是因为等到go函数（跟在`go`关键字之后的那个函数）或延迟函数真正被执行的时候，这些迭代变量的值可能已经不是我们想要的值了。

另一方面，当检查程序发现在带有`range`子句的`for`代码块中迭代出的数据并没有赋值给标识符所代表的变量时，则会忽略对这一代码块的检查。比如像这样的代码：

```
func nonIdentRange(slc []string) {
    l := len(slc)
    temp := make([]string, l)
    l--
    for _, temp[l] = range slc {
        // 忽略了使用切片值temp的代码。
        if l > 0 {
            l--
        }
    }
}
```

就不会受到检查程序的关注。另外，当被迭代的对象的大小为`0`时，`for`代码块也不会被检查。

据此，我们知道如果在可能被延迟处理的代码块中直接使用迭代中的临时变量，那么就可能会造成与编程人员意图不相符的结果。如果由此问题使程序的最终结果出现偏差甚至使程序报错的话，那么看起来就会非常诡异。这种隐晦的错误在排查时也是非常困难的。这种不正确的代码编写方式应该彻底被避免。这也是检查程序对迭代代码块进行检查的最终目的。如果检查程序发现了上述的不正确的代码编写方式，就会打印出错误信息以提醒编程人员。

**-structtags标记**

如果标记``-structtags`有效（标记值不为`false```），那么命令程序会对结构体类型的字段的标签进行检查。我们先来看下面的代码：

```
type Person struct {
    XMLName xml.Name    `xml:"person"`
    Id          int     `xml:"id,attr"`
    FirstName   string  `xml:"name>first"`
    LastName    string  `xml:"name>last"`
    Age         int     `xml:"age"`
    Height      float32 `xml:"height,omitempty"`
    Married     bool
    Address
    Comment     string  `xml:",comment"`
}
```

在上面的例子中，在结构体类型的字段声明后面的那些字符串形式的内容就是结构体类型的字段的标签。对于Go语言本身来说，结构体类型的字段标签就是注释，它们是可选的，且会被Go语言的运行时系统忽略。但是，这些标签可以通过标准库代码包`reflect`中的程序访问到。因此，不同的代码包中的程序可能会赋予这些结构体类型的字段标签以不同的含义。比如上面例子中的结构体类型的字段标签就对代码包`encoding/xml`中的程序非常有用处。

严格来讲，结构体类型的字段的标签应该满足如下要求：

1. 标签应该包含键和值，且它们之间要用英文冒号分隔。
2. 标签的键应该不包含空格、引号或冒号。
3. 标签的值应该被英文双引号包含。
4. 如果标签内容符合了第3条，那么标签的全部内容应该被反引号“`”包含。否则它需要被双引号包含。
5. 标签可以包含多个键值对，其它们之间要用空格“ ”分隔。例如：`key:"value" _gofix:"_magic"`

检查程序首先会对结构体类型的字段标签的内容做去引号处理，也就是把最外面的双引号或者反引号去除。如果去除失败，则检查程序会打印错误信息并设置退出代码为1，同时忽略后续检查。如果去引号处理成功，检查程序则会根据前面的规则对标签的内容进行检查。如果检查出问题，检查程序同样会打印出错误信息并设置退出代码为1。

**-unreachable标记**

如果标记``-unreachable`有效（标记值不为`false```），那么命令程序会在函数或方法定义中查找死代码。死代码就是永远不会被访问到的代码。例如：

```
func deadCode1() int {
    print(1)
    return 2
    println() // 这里存在死代码
}
```

在上面示例中，函数`deadCode1`中的最后一行调用打印函数的语句就是死代码。检查程序如果在函数或方法中找到死代码，则会打印错误信息以提醒编码人员。我们把这段代码放到命令源码文件deadcode_demo.go中，并在main函数中调用它。现在，如果我们编译这个命令源码文件会马上看到一个编译错误：“missing return at end of function”。显然，这个错误侧面的提醒了我们，在这个函数中存在死代码。实际上，我们在修正这个问题之前它根本就不可能被运行，所以也就不存在任何隐患。但是，如果在这个函数不需要结果的情况下又会如何呢？我们稍微改造一下上面这个函数：

```
func deadCode1() {
    print(1)
    return
    println() // 这里存在死代码
}
```

好了，我们现在把函数`deadcode1`的声明中的结果声明和函数中`return`语句后的数字都去掉了。不幸的是，当我们再次编译文件时没有看到任何报错。但是，这里确实存在死代码。在这种情况下，编译器并不能帮助我们找到问题，而`go tool vet`命令却可以。

```
hc@ubt:~$ go tool vet deadcode_demo.go
deadcode_demo.go:10: unreachable code
```

`go tool vet`命令中的检查程序对于死代码的判定有几个依据，如下：

1. 在这里，我们把`return`语句、`goto`语句、`break`语句、`continue`语句和`panic`函数调用语句都叫做流程中断语句。如果在当前函数、方法或流程控制代码块的分支中的流程中断语句的后面还存在其他语句或代码块，比如：

   func deadCode2() { print(1) panic(2) println() // 这里存在死代码 }

   或

   func deadCode3() { L: { print(1) goto L } println() // 这里存在死代码 }

   或

   func deadCode4() { print(1) return { // 这里存在死代码 } }

则后面的语句或代码块就会被判定为死代码。但检查程序仅会在错误提示信息中包含第一行死代码的位置。

1. 如果带有`else`的`if`代码块中的每一个分支的最后一条语句均为流程中断语句，则在此流程控制代码块后的代码都被判定为死代码。比如：

   func deadCode5(x int) { print(1) if x == 1 { panic(2) } else { return } println() // 这里存在死代码 }

注意，只要其中一个分支不包含流程中断语句，就不能判定后面的代码为死代码。像这样：

```
func deadCode5(x int) {
    print(1)
    if x == 1 {
        panic(2)
    } else if x == 2 {
        return
    } 
    println() // 这里并不是死代码
}
```

1. 如果在一个没有显式中断条件或中断语句的`for`代码块后面还存在其它语句，则这些语句将会被判定为死代码。比如：

   func deadCode6() { for { for { break } } println() // 这里存在死代码 }

或

```
func deadCode7() {
    for {
        for {
        }
        break // 这里存在死代码
    }
    println()
}
```

而我们对这两个函数稍加改造后，就会消除`go tool vet`命令发出的死代码告警。如下：

```
func deadCode6() {
    x := 1
    for x == 1 {
        for {
            break
        }
    }
    println() // 这里存在死代码
}
```

以及

```
func deadCode7() {
    x := 1
    for {
        for x == 1 {
        }
        break // 这里存在死代码
    }
    println()
}
```

我们只是加了一个显式的中断条件就能够使之通过死代码检查。但是，请注意！这两个函数中在被改造后仍然都包含死循环代码！这说明检查程序并不对中断条件的逻辑进行检查。

1. 如果`select`代码块的所有`case`中的最后一条语句均为流程中断语句（`break`语句除外），那么在`select`代码块后面的语句都会被判定为死代码。比如：

   func deadCode8(c chan int) { print(1) select { case <-c: print(2) panic(3) } println() // 这里存在死代码 }

或

```
func deadCode9(c chan int) {
L:
    print(1)
    select {
    case <-c:
        print(2)
        panic(3)
    case c <- 1:
        print(4)
        goto L
    }
    println() // 这里存在死代码
}
```

另外，在空的`select`语句块之后的代码也会被认为是死代码。比如：

```
func deadCode10() {
    print(1)
    select {}
    println() // 这里存在死代码
}
```

或

```
func deadCode11(c chan int) {
    print(1)
    select {
    case <-c:
        print(2)
        panic(3)
    default:
        select {}
    }
    println() // 这里存在死代码
}
```

上面这两个示例中的语句`select {}`都会引发一个运行时错误：“fatal error: all goroutines are asleep - deadlock!”。这就是死锁！关于这个错误的详细说明在第7章。

1. 如果`switch`代码块的所有`case`和`default case`中的最后一条语句均为流程中断语句（除了`break`语句），那么在`switch`代码块后面的语句都会被判定为死代码。比如：

   func deadCode14(x int) { print(1) switch x { case 1: print(2) panic(3) default: return } println(4) // 这里存在死代码 }

我们知道，关键字`fallthrough`可以使流程从`switch`代码块中的一个`case`转移到下一个`case`或`default case`。死代码也可能由此产生。例如：

```
func deadCode15(x int) {
    print(1)
    switch x {
    case 1:
        print(2)
        fallthrough
    default:
        return
    }
    println(3) // 这里存在死代码
}
```

在上面的示例中，第一个case总会把流程转移到第二个case，而第二个case中的最后一条语句为return语句，所以流程永远不会转移到语句`println(3)`上。因此，`println(3)`语句会被判定为死代码。如果我们把`fallthrough`语句去掉，那么就可以消除这个死代码判定。实际上，只要某一个`case`或者`default case`中的最后一条语句是break语句，就不会有死代码的存在。当然，这个`break`语句本身不能是死代码。另外，与`select`代码块不同的是，空的`switch`代码块并不会使它后面的代码成为死代码。

综上所述，死代码的判定虽然看似比较复杂，但其实还是有原则可循的。我们应该在编码过程中就避免编写可能会造成死代码的代码。如果我们实在不确定死代码是否存在，也可以使用`go tool vet`命令来检查。不过，需要提醒读者的是，不存在死代码并不意味着不存在造成死循环的代码。当然，造成死循环的代码也并不一定就是错误的代码。但我们仍然需要对此保持警觉。

**-asmdecl标记**

如果标记``-asmdecl`有效（标记值不为`false```），那么命令程序会对汇编语言的源码文件进行检查。对汇编语言源码文件及相应编写规则的解读已经超出了本书的范围，所以我们并不在这里对此项检查进行描述。如果读者有兴趣的话，可以查看此项检查的程序的源码文件asmdecl.go。它在Go语言安装目录的子目录src/cmd/vet下。

至此，我们对`go vet`命令和`go tool vet`命令进行了全面详细的介绍。之所以花费如此大的篇幅来介绍这两个命令，不仅仅是为了介绍此命令的使用方法，更是因为此命令程序的检查工作涉及到了很多我们在编写Go语言代码时需要避免的“坑”。由此我们也可以知晓应该怎样正确的编写Go语言代码。同时，我们也应该在开发Go语言程序的过程中经常使用`go tool vet`命来检查代码

