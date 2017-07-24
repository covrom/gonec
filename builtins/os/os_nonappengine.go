// +build !appengine

package os

import (
	"github.com/covrom/gonec/vm"
	pkg "os"
	"reflect"
)

func handleAppEngine(m *vm.Env) {
	m.Define("Getppid", reflect.ValueOf(pkg.Getppid))
}
