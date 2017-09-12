package core

type VMSlice []interface{}

func (x VMSlice) vmval() {}

func (x VMSlice) Interface() interface{} {
	return ([]interface{})(x)
}
func (x VMSlice) Slice() VMSlice {
	return x
}
