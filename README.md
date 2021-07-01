# Go Promise
A Promise library for Go

1. any value returned as "Promise" is "PromiseFunc".
2. support promise chain call.
3. support err catch.

## example:
```go
package main

import (
    "github/jamestack/promise"
    "time"
    "fmt"
)

func Hello(num int) promise.Promise {
    p := promise.New()
    time.AfterFunc(3*time.Second, func(){
        if num % 2 == 0 {
            p.Resolve(num, "yes success")
        }else {
            p.Reject("oh no fail")
        }
    })
    return p
}

func main()  {
    fmt.Println("start call")

    Hello(time.Now().Unix()).Then(func(num int, msg string) promise.Promise {
        fmt.Println("call done: ",num, msg)
        return promise.WithResolve("a", "b")
    }).Then(func(a, b string) promise.Promise {
        p := promise.New()
        time.AfterFunc(3*time.Second, func(){
        	p.Resolve()
            fmt.Println("looks good.", a, b)
        })
        return p
    }).Catch(func(errs ...promise.Any){
    	fmt.Println("catch:", errs)
    })

    fmt.Println("end call")
    <-time.After(5*time.Second)
}

```