# Go Promise
Golang版的Promise实现

1. 任何返回值为"Promise"的函数均被视为"Promise函数".
2. 支持Promise链式调用.
3. 支持错误捕获.

## 示例:
```go
package main

import (
	"fmt"
	"github.com/jamestack/promise"
	"time"
)

func Hello(num int64) promise.Promise {
	p := promise.New()
	time.AfterFunc(3*time.Second, func() {
		if num%2 == 0 {
			p.Resolve(num, "yes success")
		} else {
			p.Reject("oh no fail")
		}
	})
	return p
}

func main() {
	fmt.Println("start call")

	Hello(time.Now().Unix()).Then(func(num int64, msg string) promise.Promise {
		fmt.Println("call done: ", num, msg)
		return promise.WithResolve("a", "b")
	}).Then(func(a, b string) promise.Promise {
		p := promise.New()
		time.AfterFunc(3*time.Second, func() {
			p.Resolve()
			fmt.Println("looks good.", a, b)
		})
		return p
	}).Catch(func(errs ...promise.Any) {
		fmt.Println("catch:", errs)
	})

	fmt.Println("end call")
	<-time.After(5 * time.Second)
}

```

## [English Document](https://github.com/jamestack/promise/blob/main/README.md)
