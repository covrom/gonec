package core

type VMSlice []VMValuer

func (x VMSlice) vmval() {}

func (x VMSlice) Interface() interface{} {
	return x
}
func (x VMSlice) Slice() VMSlice {
	return x
}

// TODO: маршаллинг!!!
