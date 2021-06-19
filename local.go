// +build !wasm

package wasmtools

type V struct {
}

func (v *V) String() string {
	return ""
}

func (v *V) Int() int {
	return 0
}

func (v *V) Bool() bool {
	return false
}

func (v *V) Get(key string) *V {
	return &V{}
}

func (v *V) Set(key string, value interface{}) *V {
	return v
}

func Global() *V {
	return &V{}
}
