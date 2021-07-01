package snow

import (
	"errors"
	"reflect"
)

type Any interface {}

type Promise interface {
	Then(fn PromiseFunc) Promise
	Catch(fn ...ErrFunc)
	Resolve(args ...Any) Promise
	Reject(errors ...Any) Promise
}

var ErrNotAValidPromiseFunc = errors.New("NotAValidPromiseFunc")
var ErrPromiseFuncArgsNotMatch = errors.New("PromiseFuncArgsNotMatch")

// 任何返回值为Promise的函数均被视为合法的Promise函数
type PromiseFunc interface {}      // eg. func(args ...Any) Promise
type ErrFunc func(errors ...Any)

type promise struct {
	then     []PromiseFunc
	catch    []ErrFunc
	resolves *[]Any
	rejects  *[]Any
	np Promise
}

func (p *promise) Resolve(args ...Any) Promise {
	if p.rejects != nil || p.resolves != nil {
		return p
	}

	p.resolves = &args

	if len(p.then) > 0 {
		p.run(args...)
	}

	return p
}

func (p *promise) Reject(errors ...Any) Promise {
	if p.rejects != nil || p.resolves != nil {
		return p
	}

	p.rejects = &errors

	if len(p.catch) > 0 {
		for _,f := range p.catch {
			f(*p.rejects...)
		}
	}
	return p
}

func (p *promise) Catch(fn ...ErrFunc) {
	if fn == nil {
		return
	}

	if p.np != nil {
		p.np.Catch(fn...)
		return
	}

	if p.rejects != nil {
		for _,f := range fn {
			f(*p.rejects...)
		}
	}else {
		p.catch = append(p.catch, fn...)
	}
}

func (p *promise) Then(fn PromiseFunc) Promise {
	if fn == nil {
		return p
	}

	// PromiseFunc类型检查
	vf := reflect.ValueOf(fn)
	vt := vf.Type()
	if vt.NumOut() != 1 || vt.Out(0).Name() != "Promise" {
		panic(ErrNotAValidPromiseFunc)
	}

	if p.np != nil {
		p.np.Then(fn)
		return p
	}

	p.then = append(p.then, fn)

	if p.rejects == nil && p.resolves != nil {
		p.run(*p.resolves...)
	}

	return p
}

func (p *promise) run(args ...Any) {
	var fn PromiseFunc

	p.resolves = nil

	if len(p.then) > 0 {
		fn = p.then[0]
		p.then = p.then[1:]
	}else {
		return
	}

	vf := reflect.ValueOf(fn)
	if vf.Type().NumIn() != len(args) {
		panic(ErrPromiseFuncArgsNotMatch)
	}

	var req []reflect.Value
	for _,a := range args {
		req = append(req, reflect.ValueOf(a))
	}

	res := vf.Call(req)
	np,ok := res[0].Interface().(Promise)
	if !ok {
		panic(ErrNotAValidPromiseFunc)
	}

	if len(p.then) == 0  {
		p.np = np
	}else {
		fn = p.then[0]
		p.then = p.then[1:]

		np.Then(fn)
		if len(p.catch) > 0 {
			np.Catch(p.catch...)
		}
	}
}

func New() Promise {
	p := &promise{}
	return p
}

func WithResolve(args ...Any) Promise {
	return New().Resolve(args...)
}

func WithReject(errors ...Any) Promise {
	return New().Reject(errors...)
}
