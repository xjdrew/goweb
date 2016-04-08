package core

type C struct {
	Env map[interface{}]interface{}
}

func NewC() *C {
	return &C{make(map[interface{}]interface{})}
}
