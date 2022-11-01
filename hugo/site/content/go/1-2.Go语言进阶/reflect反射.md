---
title: "reflect反射"
date: 2022-05-06T10:26:00+08:00
draft: true
---
### 基本概念

1. 反射：支持程序在运行中，动态地访问和修改变量的类型和值

### 反射规则

- 实例、Value、Type之间可以相互转换
- reflect.Type表示实例的类型信息
- reflect.Value表示实例的值信息

#### 实例、Type互转

```go
// src/reflect/type.go
type Type interface {...}

// 实例转Type
// 若i传入的是具体类型，则返回具体类型信息
// 若i传入的是接口类型，且绑定了具体实例，则返回接口的动态类型信息
// 若i传入的是接口类型，其未绑定具体实例，则返回接口的静态类型信息
func TypeOf(i any) Type {}

// Type转实例？
```

#### 实例、Value互转

```go
// src/reflect/value.go
type Value struct {
	typ *rtype
	ptr unsafe.Pointer
	flag
} 

// 实例转Value
func ValueOf(i any) Value {}

// Value转实例
func (v Value) Interface() (i any) {}
func (v Value) Int() int64 {}
func (v Value) Float() float64 {}
func (v Value) Bool() bool {}
...
```

#### Type、Value互转

```go
// Value转Type
func (v Value) Type() Type {}

// Type转Value
func New(typ Type) Value {}
func NewAt(typ Type, p unsafe.Pointer) Value {}
func Zero(typ Type) Value {}
```

#### Value到Value

```go
// 指针类型Value到值类型Value
func (v Value) Elem() Value {}
func Indirect(v Value) Value {}
```

#### Type到Type

```go
// 指针类型Type到值类型Type
// Elem returns a type's element type.
// It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
type Type interface {
	Elem() Type
}

// 值类型Type到指针类型Type
func PtrTo(t Type) Type { return PointerTo(t) }
```

#### Value值的可修改性

```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}


func main() {
	u := User{1, "Bob", 25}
	fmt.Println(u)

	va := reflect.ValueOf(u)	// va是Value值类型，不可修改原值
	vb := reflect.ValueOf(&u)	// va是Value指针类型，可修改原值

	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet())
	//fmt.Println(vb.CanSet(), vb.FieldByName("Name").CanSet())				// vb是指针类型，FieldByName是值方法，会panic
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet())	// 通过Elem()将指针类型Value转换为值类型Value

	// 修改指针类型的Name字段
	name := "Alic"
	vname := reflect.ValueOf(name)
	vb.Elem().FieldByName("Name").Set(vname)

	fmt.Println(vb)
	fmt.Println(u)
}

//{1 Bob 25}
//false false
//false true
//&{1 Alic 25}
//{1 Alic 25}
```

