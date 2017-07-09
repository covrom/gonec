package gonec

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_interpreter_ParseAndRun(t *testing.T) {
	ti1 := Interpreter()
	in1 := strings.NewReader(`
	x = 5;
	y = 10;
	foobar = 838383;
	
	bad =

	`)
	w := &bytes.Buffer{}
	err := ti1.ParseAndRun(in1, w)
	fmt.Println(w.String())
	if err != nil {
		fmt.Println(err.Error())
	}
}
