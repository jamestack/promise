# Go Promise
A Promise library for Go

1. any value returned as "Promise" is "PromiseFunc".
2. support promise chain call.
3. support err catch.

## Example:
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

## [中文文档](https://github.com/jamestack/promise/blob/main/README.zh_cn.md)
