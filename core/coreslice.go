package core

type VMSlice []VMValuer

func (x VMSlice) vmval() {}

func (x VMSlice) Interface() interface{} {
	return x
}
func (x VMSlice) Slice() VMSlice {
	return x
}

func (x VMSlice) Len() VMInt {
	return VMInt(len(x))
}

func (x VMSlice) Index(i VMValuer) VMValuer {
	if ii, ok := i.(VMInt); ok {
		return x[int(ii)]
	}
	panic("Индекс должен быть целым числом")
}

// TODO: маршаллинг!!!
