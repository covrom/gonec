// +build !appengine

package os

import (
	pkg "os"
	"reflect"

	envir "github.com/covrom/gonec/env"
)

func handleAppEngine(m *envir.Env) {
	m.DefineS("Getppid", reflect.ValueOf(pkg.Getppid))
}
