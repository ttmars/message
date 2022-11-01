---
title: "GPM调度模型"
date: 2022-05-06T10:26:00+08:00
draft: true
---
### 例1

[https://www.bilibili.com/video/BV19b4y1i74w](https://www.bilibili.com/video/BV19b4y1i74w)

goroutine调度优先级

- runnext中的g
- 本地队列runq中的g，大小为256
- 全局队列runq中的g

```go
package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(3)

	go func(n int) {
		println(n)
		wg.Done()
	}(1)

	go func(n int) {
		println(n)
		wg.Done()
	}(2)

	go func(n int) {
		println(n)
		wg.Done()
	}(3)

	wg.Wait()
}

//3
//1
//2
```

### 例2

```go
package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	for i:=1;i<=257;i++{
		wg.Add(1)
		go func(n int) {
			println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
```

