

+++

title = "Title"
description = "Description"
tags = ["it","tag","tag"]

+++



# runtime

## 环境变量

[控制 Go 程序的运行时行为的环境变量](https://pkg.go.dev/runtime#hdr-Environment_Variables)：GOGC、GODEBUG、GOMAXPROCS、GORACE、GOTRACEBACK。

GOGC：设置初始垃圾收集目标百分比。当新分配的数据与上一次收集后剩余的实时数据的比率达到此百分比时，将触发收集。默认值为 GOGC=100，设置 GOGC=off 完全禁用垃圾收集器。运行时/调试包的 [SetGCPercent](https://pkg.go.dev/runtime/debug#SetGCPercent) 函数允许在运行时更改此百分比。

GODEBUG：控制运行时中的调试变量。它是一个以逗号分隔的 name=val 对列表，用于设置这些命名变量。



### GODEBUG

#### gctrace

gctrace：设置 gctrace=1 会导致垃圾收集器在每个收集时发出一行标准错误，总结收集的内存量和暂停的（时间）长度。此行的格式可能会更改，目前它是：`gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P`。其中字段（说明）如下：

- **gc #**：GC 编号，每次 GC 时自增
- **@#s**：自程序启动以来的时间（单位秒）
- **#%**：自程序启动以来在 GC 中花费的时间百分比
- **#+...+#**：GC 各阶段的挂钟/CPU时间。即 mark 阶段用时 + mutator assist/dedicated&fractional/idle 用时 + markTermination阶段用时
- **#->#-># MB**：GC开始时、GC结束时和活堆时的堆大小MB。开始 mark 阶段前的 heap_live 大小 -> 开始 markTermination 阶段前 heap_live 大小 -> 被标记对象的大小
- **# MB goal**：目标堆大小。表示下一次出发 GC 回收的阈值大小
- **# P**：使用的处理器数量。本次 GC 一共涉及多少个 P

这些阶段是 stop-the-world (STW) 扫描终止、并发标记和扫描以及 STW 标记终止。标记/扫描的 CPU 时间分为辅助时间（根据分配执行的 GC）、后台 GC 时间和空闲 GC 时间。
如果该行以"(forced)"结尾，则此 GC 是由 runtime.GC() 调用强制执行的

```go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}
```

```shell
$ GODEBUG=gctrace=1 go run main.go
gc 1 @0.008s 3%: 0.050+0.90+0.028 ms clock, 0.10+0.43/0.021/0+0.056 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
gc 2 @0.016s 2%: 0.009+0.44+0.030 ms clock, 0.018+0.20/0.082/0.21+0.061 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
gc 3 @0.029s 2%: 0.068+0.38+0.001 ms clock, 0.13+0.30/0/0+0.002 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
gc 4 @0.037s 2%: 0.014+1.0+0.023 ms clock, 0.029+0.46/0.018/0+0.046 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
# command-line-arguments
gc 1 @0.002s 13%: 0.006+1.6+0.037 ms clock, 0.012+0.33/0.83/0+0.075 ms cpu, 4->4->3 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
# command-line-arguments
gc 1 @0.000s 30%: 0.061+0.81+0.020 ms clock, 0.12+0.75/0.016/0+0.040 ms cpu, 4->4->3 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
gc 2 @0.002s 15%: 0.016+1.0+0.001 ms clock, 0.032+0.056/0.022/1.0+0.002 ms cpu, 6->6->6 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 2 P
```

#### schedtrace & scheddetail

```go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}
```

```shell
$ go build -o schedtrace main.go
$ GODEBUG=schedtrace=1000 ./schedtrace
SCHED 0ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [4 4]
SCHED 1003ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=3 [2 3]
SCHED 2005ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=4 [0 4]
SCHED 3019ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=3 [1 4]
SCHED 4029ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=5 [0 3]
SCHED 5030ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=4 [2 2]
SCHED 6041ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=6 [2 0]
......
```

以 `SCHED 2005ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=4 [0 4]` 为例，解释如下：

- sched：每一行都代表调度器的调试信息，后面提示的毫秒数表示启动到现在的运行时间，输出的时间间隔受 `schedtrace` 的值影响
- gomaxprocs：当前的 CPU 核心数（GOMAXPROCS 的当前值）。
- idleprocs：空闲的处理器数量，后面的数字表示当前的空闲数量。
- threads：OS 线程数量，后面的数字表示当前正在运行的线程数量。
- spinningthreads：自旋状态的 OS 线程数量。
- idlethreads：空闲的线程数量。
- runqueue：全局队列中中的 Goroutine 数量，而后面的 [0 0 1 1] 则分别代表这 4 个 P 的本地队列正在运行的 Goroutine 数量。



**scheddetail**: 设置 schedtrace=X 和 scheddetail=1 致使调度器每 X 毫秒发出详细的多行信息，描述调度器 SCHED、处理器P、线程M、goroutine G 的状态。

```shell
$ GODEBUG=schedtrace=1000,scheddetail=1 ./schedtrace
SCHED 1004ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=4 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
  P0: status=1 schedtick=40 syscalltick=0 m=3 runqsize=1 gfreecnt=0 timerslen=0
  P1: status=1 schedtick=45 syscalltick=0 m=0 runqsize=3 gfreecnt=0 timerslen=0
  M3: p=0 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=false blocked=false lockedg=-1
  M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 spinning=false blocked=false lockedg=-1
  M0: p=1 curg=19 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=false lockedg=-1
  G1: status=4(semacquire) m=-1 lockedm=-1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=4(GC scavenge wait) m=-1 lockedm=-1
  G17: status=1() m=-1 lockedm=-1
  G18: status=1() m=-1 lockedm=-1
  G19: status=2() m=0 lockedm=-1
  G20: status=1() m=-1 lockedm=-1
  G21: status=1() m=-1 lockedm=-1
  G22: status=1() m=-1 lockedm=-1
  G23: status=1() m=-1 lockedm=-1
  G24: status=1() m=-1 lockedm=-1
  G25: status=1() m=-1 lockedm=-1
  G26: status=1() m=-1 lockedm=-1
```

详细解释见[链接](https://eddycjy.gitbook.io/golang/di-9-ke-gong-ju/godebug-sched#scheddetail)。



## lib

### FuncForPC

[FuncForPC](https://pkg.go.dev/runtime#FuncForPC) 返回一个 *[Func](https://pkg.go.dev/runtime#Func)，描述给定程序计数器（pc）地址的函数。如果 pc 由于内联而表示多个函数，则它返回描述内层函数的 *Func，但带有最外层函数的条目。

用法见下面 Calller

### Caller

[Caller](https://pkg.go.dev/runtime#Caller)：入参中的 skip 是要上升的堆栈帧数



```go
func main() {
	for i := 0; i < 5; i++ {
		PrintCaller(i)
	}
}

func PrintCaller(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("Can not get caller when skip=%v, maybe it's past the original call\n", skip)
		return
	}
	fmt.Printf("funcName=%v, fileName=%v, funcLine=%v\n", runtime.FuncForPC(pc).Name(), file, line)
}
```

可见随 skip 增大，获取到的是更上一层的 caller

```shell
$ go run main.go
funcName=main.PrintCaller, fileName=/root/go/path/yuanyatianchi.io/go/runtime/lib/main.go, funcLine=15
funcName=main.main, fileName=/root/go/path/yuanyatianchi.io/go/runtime/lib/main.go, funcLine=10
funcName=runtime.main, fileName=/root/go/go/src/runtime/proc.go, funcLine=250
funcName=runtime.goexit, fileName=/root/go/go/src/runtime/asm_amd64.s, funcLine=1571
Can not get caller when skip=4, maybe it's past the original call
```

### Callers

[Callers](https://pkg.go.dev/runtime#Callers)：入参中的 skip 是要上升的堆栈帧数

```go
func main() {
	for i := 0; i < 5; i++ {
		PrintCallers(i)
	}
}

func PrintCallers(skip int) {
	pcs := make([]uintptr, 10)
	count := runtime.Callers(skip, pcs)
	fmt.Printf("skip=%v\tcount=%v\t", skip, count)
	for i := 0; i < count; i++ {
		fmt.Printf("fileName=%v\t", runtime.FuncForPC(pcs[i]).Name())
	}
	fmt.Println()
}
```

```shell
$ go run main.go
skip=0  count=5 fileName=runtime.Callers        fileName=runtime.Callers        fileName=main.main      fileName=runtime.main   fileName=runtime.goexit
skip=1  count=4 fileName=runtime.Callers        fileName=main.main      fileName=runtime.main   fileName=runtime.goexit
skip=2  count=3 fileName=main.main      fileName=runtime.main   fileName=runtime.goexit
skip=3  count=2 fileName=runtime.main   fileName=runtime.goexit
skip=4  count=1 fileName=runtime.goexit
```

### CallersFrames

[CallersFrames](https://pkg.go.dev/runtime#CallersFrames) 获取 Callers 返回的 PC 值的一部分，并准备返回函数/文件/行信息。在完成 Frames 之前不要更改切片。

```go
func main() {
	PrintCallersFrames(0)
}

func PrintCallersFrames(skip int) {
	pcs := make([]uintptr, 10)
	runtime.Callers(skip, pcs)

	frames := runtime.CallersFrames(pcs)

	for {
		frame, _ := frames.Next()
		if frame.PC == 0 {
			break
		}
		fmt.Println(frame.Function, frame.File, frame.Line)
	}
}
```

```shell
$ go run main.go
runtime.Callers /root/go/go/src/runtime/extern.go 235
main.PrintCallersFrames /root/go/path/yuanyatianchi.io/go/runtime/lib/main.go 34
main.main /root/go/path/yuanyatianchi.io/go/runtime/lib/main.go 9
runtime.main /root/go/go/src/runtime/proc.go 250
runtime.goexit /root/go/go/src/runtime/asm_amd64.s 1571
```

