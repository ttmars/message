---
title: "内存管理与GC"
date: 2022-05-06T10:26:00+08:00
draft: true
---
### 堆栈区别

- 堆在运行过程中动态分配内存，需要通过GC回收，分配速度慢
- 栈内存自动释放，分配速度快

### 逃逸分析

```go
package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

func main() {
	x := foo()
	fmt.Println(*x)
}
```

```shell
$ go build -gcflags '-m -l'  main.go
# command-line-arguments
.\main.go:6:2: moved to heap: t
.\main.go:12:13: ... argument does not escape
.\main.go:12:14: *x escapes to heap
```

参数说明：

- -gcflags 启动编译器支持的额外标志
- -m 输出编译器的优化细节
- -N 禁用优化
- -l 禁用内联

[https://stackoverflow.com/questions/62589743/whats-go-cmd-option-gcflags-all-possible-values](https://stackoverflow.com/questions/62589743/whats-go-cmd-option-gcflags-all-possible-values)

[https://pkg.go.dev/cmd/compile](https://pkg.go.dev/cmd/compile)

### 内存对齐

数据类型的大小和对齐边界

1. 类型大小小于8，对齐边界等于类型大小
2. 类型大小大于等于8，对齐边界等于8
3. 例外：complex64大小为8，对齐边界为4

```go
// 获取类型大小
func Sizeof(x ArbitraryType) uintptr

// 获取结构体某个字段偏移
func Offsetof(x ArbitraryType) uintptr

// 获取类型对齐边界
func Alignof(x ArbitraryType) uintptr
```

```go
package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	a int8
	b int64
	c int8
	d int32
	e int16
}

type s2 struct {
	a int8
	c int8
	e int16
	d int32
	b int64
}

type s3 struct {
	A int32 // 4
	B []int32 // 24
	C string // 16
	D bool // 1
}

func main()  {
	var v1 s1
	var v2 s2
	var v3 s3
	fmt.Println("s1:",unsafe.Sizeof(v1))
	fmt.Println("s2:",unsafe.Sizeof(v2))
	fmt.Println("s3:",unsafe.Sizeof(v3))
}

//s1: 32
//s2: 16
//s3: 56
```

