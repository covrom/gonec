// +build go1.8

package sort

import (
	s "sort"

	envir "github.com/covrom/gonec/env"
)

func handleGo18(m *envir.Env) {
	m.DefineS("Slice", func(arr interface{}, less func(i, j int) bool) interface{} {
		s.Slice(arr, less)
		return arr
	})
}
