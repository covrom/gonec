package core

type VMStringMap map[string]interface{}

func (x VMStringMap) vmval() {}

func (x VMStringMap) Interface() interface{} {
	return (map[string]interface{})(x)
}

func (x VMStringMap) StringMap() VMStringMap {
	return x
}
