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

func (v *V) IsNull() bool {
	return true
}

func (v *V) IsUndefined() bool {
	return true
}

func (v *V) Call(key string, args ...interface{}) *V {
	return &V{}
}

func Global() *V {
	return &V{}
}
