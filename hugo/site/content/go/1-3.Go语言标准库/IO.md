### 理解io.Reader

```go
// 从Reader中读取数据到切片中，返回值n<=len(p)
// 若传入的切片长度为0，则不会读取任何数据，如：p := make([]byte, 0 ,10)
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	b := make([]byte, 0, 10)
	//b := make([]byte, 10)
	fmt.Println(len(b), cap(b), b)
	f,err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	//n,err := f.Read(b[len(b):cap(b)])
	n,err := f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(len(b), cap(b), b)
}

//0 10 []
//0
//0 10 []
```

