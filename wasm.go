// +build wasm

package wasmtools

import "syscall/js"

type V struct {
	js js.Value
}

func (v *V) String() string {
	return v.js.String()
}

func (v *V) Int() int {
	return v.js.Int()
}

func (v *V) Bool() bool {
	return v.js.Bool()
}

func (v *V) Get(key string) *V {
	return &V{
		v.js.Get(key),
	}
}

func (v *V) Set(key string, value interface{}) *V {
	v.js.Set(key, value)
	return v
}

func Global() *V {
	return &V{
		js.Global(),
	}
}
