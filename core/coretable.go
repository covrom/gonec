package core

// ТаблицаЗначений

type VMTableColumn struct {
	VMMetaObj

	cols *VMTableColumns
	name string
}

func NewVMTableColumn(vtcs *VMTableColumns) *VMTableColumn {
	vtc := &VMTableColumn{
		cols: vtcs,
	}
	vtc.VMInit(vtc)
	vtc.VMRegister()
	return vtc
}

func (vtc *VMTableColumn) VMRegister() {
	// vt.VMRegisterField("ПолеСтрока", &vt.ПолеСтрока)
	// vt.VMRegisterMethod("Имя", vt.Имя)
}

type VMTableColumns struct {
	VMMetaObj

	table *VMTable
	cols  []*VMTableColumn
}

func NewVMTableColumns(vt *VMTable) *VMTableColumns {
	vtcs := &VMTableColumns{
		table: vt,
	}
	vtcs.VMInit(vtcs)
	vtcs.VMRegister()
	return vtcs
}

func (vtcs *VMTableColumns) VMRegister() {
	vtcs.cols = make([]*VMTableColumn, 0, 8)
	// vt.VMRegisterField("ПолеСтрока", &vt.ПолеСтрока)
	// vt.VMRegisterMethod("Имя", vt.Имя)
}

type VMTableLine struct {
	VMMetaObj

	table *VMTable
	line  VMSlice
}

func NewVMTableLine(vt *VMTable) *VMTableLine {
	vtl := &VMTableLine{
		table: vt,
	}
	vtl.VMInit(vtl)
	vtl.VMRegister()
	return vtl
}

func (vtl *VMTableLine) VMRegister() {
	// vt.VMRegisterField("ПолеСтрока", &vt.ПолеСтрока)
	// vt.VMRegisterMethod("Имя", vt.Имя)
}

type VMTable struct {
	VMMetaObj

	cols  *VMTableColumns
	lines []*VMTableLine
}

func (vt *VMTable) VMRegister() {
	vt.cols = NewVMTableColumns(vt)

	vt.lines = make([]*VMTableLine, 20)
	vt.VMRegisterField("колонки", vt.cols)

	// vt.VMRegisterMethod("колонки", vt.Колонки)
	// vt.VMRegisterField("ПолеСтрока", &vt.ПолеСтрока)
}

func (vt *VMTable) Slice() VMSlice {
	rm := make(VMSlice, len(vt.lines))
	for i, v := range vt.lines {
		rm[i] = v
	}
	return rm
}

func (vt *VMTable) Length() VMInt {
	return VMInt(len(vt.lines))
}

func (vt *VMTable) IndexVal(idx VMValuer) VMValuer {
	if i, ok := idx.(VMInt); ok {
		return vt.lines[int(i)]
	}
	panic("индекс должен быть числом")
}
