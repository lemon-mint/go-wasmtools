// +build wasm

package wasmtools

import (
	"syscall/js"
)

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

func (v *V) IsNull() bool {
	return v.js.IsNull()
}

func (v *V) IsUndefined() bool {
	return v.js.IsUndefined()
}

func (v *V) Call(key string, args ...interface{}) *V {
	p := make([]interface{}, len(args))
	for i := range args {
		switch args[i].(type) {
		case *V:
			p[i] = args[i].(*V).js
		case V:
			p[i] = args[i].(V).js
		default:
			p[i] = args[i]
		}

	}
	return &V{
		v.js.Call(key, args...),
	}
}

func Global() *V {
	return &V{
		js.Global(),
	}
}
