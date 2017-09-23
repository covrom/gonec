package bincode

import (
	"github.com/covrom/gonec/core"
)

// Регистры виртуальной машины

type VMRegs struct {
	Env          *core.Env
	Reg          []core.VMValuer // регистры значений
	Labels       []int           // [label]=index в BinCode
	TryLabel     []int           // последний элемент - это метка на текущий обработчик CATCH
	TryRegErr    []int           // последний элемент - это регистр с ошибкой текущего обработчика
	ForBreaks    []int           // последний элемент - это метка для break
	ForContinues []int           // последний элемент - это метка для continue
	// ReturnTo     []int           // стек возвратов по RET
}

func (v *VMRegs) FreeFromReg(reg int) {
	// освобождаем память, начиная с reg, для сборщика мусора
	// v.Reg = v.Reg[:reg]
	if reg < len(v.Reg) {
		cl := make([]core.VMValuer, len(v.Reg)-reg)
		copy(v.Reg[reg:], cl)
	}

	// for i := reg; i < len(v.Reg); i++ {
	// 	v.Reg[i] = nil
	// }
}

func (v *VMRegs) PushTry(reg, label int) {
	v.TryRegErr = append(v.TryRegErr, reg)
	v.TryLabel = append(v.TryLabel, label)
}

func (v *VMRegs) TopTryLabel() int {
	l := len(v.TryLabel)
	if l == 0 {
		return -1
	}
	return v.TryLabel[l-1]
}

func (v *VMRegs) PopTry() (reg int, label int) {
	l := len(v.TryLabel)
	if l == 0 {
		return -1, -1
	}
	reg = v.TryRegErr[l-1]
	v.TryRegErr = v.TryRegErr[0 : l-1]
	label = v.TryLabel[l-1]
	v.TryLabel = v.TryLabel[0 : l-1]
	return
}

func (v *VMRegs) PushBreak(label int) {
	v.ForBreaks = append(v.ForBreaks, label)
}

func (v *VMRegs) TopBreak() int {
	l := len(v.ForBreaks)
	if l == 0 {
		return -1
	}
	return v.ForBreaks[l-1]
}

func (v *VMRegs) PopBreak() (label int) {
	l := len(v.ForBreaks)
	if l == 0 {
		return -1
	}
	label = v.ForBreaks[l-1]
	v.ForBreaks = v.ForBreaks[0 : l-1]
	return
}

func (v *VMRegs) PushContinue(label int) {
	v.ForContinues = append(v.ForContinues, label)
}

func (v *VMRegs) TopContinue() int {
	l := len(v.ForContinues)
	if l == 0 {
		return -1
	}
	return v.ForContinues[l-1]
}

func (v *VMRegs) PopContinue() (label int) {
	l := len(v.ForContinues)
	if l == 0 {
		return -1
	}
	label = v.ForContinues[l-1]
	v.ForBreaks = v.ForContinues[0 : l-1]
	return
}
