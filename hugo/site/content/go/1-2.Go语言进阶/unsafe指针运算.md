---
title: "unsafe指针运算"
date: 2022-05-06T10:26:00+08:00
draft: true
---
### Go中的三种指针类型

- 普通类型的指针
- unsafe.Pointer指针
- uintptr指针

普通类型的指针和unsafe.Pointer指针可以相互转换，因此不同普通类型的指针可以互转

unsafe.Pointer指针可以和uintptr指针相互转换，且uintptr指针可以运算，因此可以实现普通类型的指针运算

```go
package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	s := "hello"
	fmt.Printf("%p %p %v\n", &s, unsafe.Pointer(&s), uintptr(unsafe.Pointer(&s)))
}

// 0xc00004a250 0xc00004a250 824634024528
```

### 转换关系

普通指针类型 <==> unsafe.Pointer <==> uintptr

- 通过unsafe.Pointer，可以实现不同指针类型相互转换
- 通过unsafe.Pointer+uintptr，可以实现指针运算

### 修改私有成员

```go
// src/strings/reader.go
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
```

```go
package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	r := strings.NewReader("hello")
	fmt.Println(r)
    // 因为r是结构体指针类型，所有直接传r而不是&r
	p1 := (*string)(unsafe.Pointer(r))
	p2 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + unsafe.Sizeof(string(""))))
	p3 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + unsafe.Sizeof(string("")) + unsafe.Sizeof(int64(0))))
	fmt.Printf("r.s:%v\nr.i:%v\nr.prevRune:%v\n", *p1, *p2, *p3)
	*p1 = "world"
	*p2 = 2
	*p3 = 2
	fmt.Println(r)
}

//&{hello 0 -1}
//r.s:hello
//r.i:0
//r.prevRune:-1
//&{world 2 2}
```

### 获取slice长度

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sli := make([]int, 5, 10)
	fmt.Println(len(sli), cap(sli))
    // 切片应该视为结构体，而不是结构体指针，所有此处应传递&sli
	lenPtr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sli)) + uintptr(8)))
	capPtr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sli)) + uintptr(16)))
	fmt.Println(*lenPtr, *capPtr)
}

//5 10
//5 10
```

### 获取map长度

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := make(map[string]int)
	m["math"] = 90
	m["english"] = 88

	p := (**int)(unsafe.Pointer(&m))
	fmt.Println(len(m), **p)
}

//2 2
```

### 实现string和[]byte的零复制转换

```go
// src/reflect/value.go
type StringHeader struct {
	Data uintptr
	Len  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

```go
func stringToBytes(s string) []byte  {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytesToString(b []byte) string  {
	return *(*string)(unsafe.Pointer(&b))
}
```

### 