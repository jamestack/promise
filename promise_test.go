package promise

import (
	"fmt"
	"testing"
	"time"
)

func Say(a, b string) Promise {
	fmt.Println("Say")
	return WithResolve(a, b)
}

func TestPromise_Resolve(t *testing.T) {
	Say("a", "b").Then(func(a, b string) Promise {
		p := New()
		fmt.Println("Then", a, b)
		time.AfterFunc(2*time.Second, func() {
			p.Resolve(1, 2)
		})

		//p.Resolve(3, 4)
		return p
	}).Then(func(a, b int) Promise {
		p := New()
		fmt.Println("Then", a, b)
		p.Resolve()
		p.Then(func() Promise {
			return New()
		})
		return p
	}).Catch(func(errors ...Any) {
		fmt.Println("catch:", errors)
	})

	fmt.Println("done")
	<-time.After(10*time.Second)
}

func TestPromise_Then(t *testing.T) {
	Say("a", "b").Then(func(a,b string) Promise {
		fmt.Println("Then", a,b)
		return WithResolve(1,2)
	}).Then(func(a, b int) Promise {

		return WithResolve()
	}).Catch(func(errors ...Any) {
		fmt.Println("Catch:", errors)
	})

	<-time.After(3*time.Second)
}
