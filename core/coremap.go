package core

type VMStringMap map[string]VMValuer

func (x VMStringMap) vmval() {}

func (x VMStringMap) Interface() interface{} {
	return x
}

func (x VMStringMap) StringMap() VMStringMap {
	return x
}

// TODO: маршаллинг!!!
