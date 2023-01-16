

+++

title = "Title"
description = "Description"
tags = ["it","tag","tag"]

+++



# Title



## sync.atomic

#### atomic.AddUint32减法操作

可以通过先声明 uint32 变量，传参时再在参数前添加 `-` 号实现

```go
func main() {
	var (
		a uint32
		b uint32 = 1
	)
	fmt.Println(atomic.AddUint32(&a, -b))
}
```

```shell
$ go run main.go
4294967295
```

