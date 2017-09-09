// Package core implements core interface for gonec script.
package core

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/covrom/gonec/ast"
	envir "github.com/covrom/gonec/env"

	gonec_encoding_json "github.com/covrom/gonec/builtins/encoding/json"
	gonec_errors "github.com/covrom/gonec/builtins/errors"
	gonec_flag "github.com/covrom/gonec/builtins/flag"
	gonec_fmt "github.com/covrom/gonec/builtins/fmt"
	gonec_io "github.com/covrom/gonec/builtins/io"
	gonec_io_ioutil "github.com/covrom/gonec/builtins/io/ioutil"
	gonec_math "github.com/covrom/gonec/builtins/math"
	gonec_math_big "github.com/covrom/gonec/builtins/math/big"
	gonec_math_rand "github.com/covrom/gonec/builtins/math/rand"
	gonec_net "github.com/covrom/gonec/builtins/net"
	gonec_net_http "github.com/covrom/gonec/builtins/net/http"
	gonec_net_url "github.com/covrom/gonec/builtins/net/url"
	gonec_os "github.com/covrom/gonec/builtins/os"
	gonec_os_exec "github.com/covrom/gonec/builtins/os/exec"
	gonec_os_signal "github.com/covrom/gonec/builtins/os/signal"
	gonec_path "github.com/covrom/gonec/builtins/path"
	gonec_path_filepath "github.com/covrom/gonec/builtins/path/filepath"
	gonec_regexp "github.com/covrom/gonec/builtins/regexp"
	gonec_runtime "github.com/covrom/gonec/builtins/runtime"
	gonec_sort "github.com/covrom/gonec/builtins/sort"
	gonec_strings "github.com/covrom/gonec/builtins/strings"

	gonec_colortext "github.com/covrom/gonec/builtins/github.com/daviddengcn/go-colortext"
)

// LoadAllBuiltins is a convenience function that loads all defineSd builtins.
func LoadAllBuiltins(env *envir.Env) {
	Import(env)

	pkgs := map[string]func(env *envir.Env) *envir.Env{
		"encoding/json":                       gonec_encoding_json.Import,
		"errors":                              gonec_errors.Import,
		"flag":                                gonec_flag.Import,
		"fmt":                                 gonec_fmt.Import,
		"io":                                  gonec_io.Import,
		"io/ioutil":                           gonec_io_ioutil.Import,
		"math":                                gonec_math.Import,
		"math/big":                            gonec_math_big.Import,
		"math/rand":                           gonec_math_rand.Import,
		"net":                                 gonec_net.Import,
		"net/http":                            gonec_net_http.Import,
		"net/url":                             gonec_net_url.Import,
		"os":                                  gonec_os.Import,
		"os/exec":                             gonec_os_exec.Import,
		"os/signal":                           gonec_os_signal.Import,
		"path":                                gonec_path.Import,
		"path/filepath":                       gonec_path_filepath.Import,
		"regexp":                              gonec_regexp.Import,
		"runtime":                             gonec_runtime.Import,
		"sort":                                gonec_sort.Import,
		"strings":                             gonec_strings.Import,
		"github.com/daviddengcn/go-colortext": gonec_colortext.Import,
	}

	env.DefineS("импорт", func(s string) interface{} {
		if loader, ok := pkgs[strings.ToLower(s)]; ok {
			m := loader(env)
			return m
		}
		panic(fmt.Sprintf("Пакет '%s' не найден", s))
	})

}

/////////////////
// TttStructTest - тестовая структура для отладки работы с системными функциональными структурами
type TttStructTest struct {
	ПолеЦелоеЧисло int
	ПолеСтрока     string
}

// обратите внимание - русскоязычное название метода для структуры
func (tst *TttStructTest) ВСтроку() string {
	return fmt.Sprintf("ПолеЦелоеЧисло=%v, ПолеСтрока=%v", tst.ПолеЦелоеЧисло, tst.ПолеСтрока)
}

func (tst TttStructTest) ВСтроку2() string {
	return fmt.Sprintf("ПолеЦелоеЧисло=%v, ПолеСтрока=%v", tst.ПолеЦелоеЧисло, tst.ПолеСтрока)
}

/////////////////

// Import defineSs core language builtins - len, range, println, int64, etc.
func Import(env *envir.Env) *envir.Env {
	env.DefineS("длина", func(v interface{}) int64 {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.String {
			return int64(len([]byte(rv.String())))
		}
		if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
			panic("Аргумент должен быть строкой или массивом")
		}
		return int64(rv.Len())
	})

	env.DefineS("ключи", func(v interface{}) []string {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Map {
			panic("Аргумент должен быть структурой")
		}
		keys := []string{}
		mk := rv.MapKeys()
		for _, key := range mk {
			keys = append(keys, key.String())
		}
		// ключи потом обходим в порядке сортировки по алфавиту
		sort.Strings(keys)
		return keys
	})

	env.DefineS("диапазон", func(args ...int64) []int64 {
		if len(args) < 1 {
			panic("Отсутствуют аргументы")
		}
		if len(args) > 2 {
			panic("Должен быть только один аргумент")
		}
		var min, max int64
		if len(args) == 1 {
			min = 0
			max = args[0] - 1
		} else {
			min = args[0]
			max = args[1]
		}
		arr := []int64{}
		for i := min; i <= max; i++ {
			arr = append(arr, i)
		}
		return arr
	})

	env.DefineS("текущаядата", func() VMTime {
		return VMTime(time.Now())
	})

	env.DefineS("прошловременис", func(t VMTime) time.Duration {
		return time.Since(time.Time(t))
	})

	env.DefineS("пауза", time.Sleep)

	env.DefineS("выполнитьспустя", time.AfterFunc)

	env.DefineS("длительностьнаносекунды", time.Nanosecond)
	env.DefineS("длительностьмикросекунды", time.Microsecond)
	env.DefineS("длительностьмиллисекунды", time.Millisecond)
	env.DefineS("длительностьсекунды", time.Second)
	env.DefineS("длительностьминуты", time.Minute)
	env.DefineS("длительностьчаса", time.Hour)
	env.DefineS("длительностьдня", time.Duration(time.Hour*24))
	env.DefineS("длительность", func(s string) time.Duration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(err)
		}
		return d
	})

	env.DefineS("нрег", func(v interface{}) string {
		if b, ok := v.([]byte); ok {
			return strings.ToLower(string(b))
		}
		return strings.ToLower(fmt.Sprint(v))
	})

	env.DefineS("врег", func(v interface{}) string {
		if b, ok := v.([]byte); ok {
			return strings.ToUpper(string(b))
		}
		return strings.ToUpper(fmt.Sprint(v))
	})

	env.DefineS("формат", env.Sprintf)

	env.DefineS("кодсимвола", func(s string) rune {
		if len(s) == 0 {
			return 0
		}
		return []rune(s)[0]
	})

	env.DefineS("вбайты", func(s string) []byte {
		return []byte(s)
	})

	env.DefineS("вруны", func(s string) []rune {
		return []rune(s)
	})

	env.DefineS("типзнч", func(v interface{}) string {
		if v == nil {
			return "Неопределено"
		}
		return ast.UniqueNames.Get(env.TypeName(reflect.TypeOf(v)))
	})

	env.DefineS("присвоенозначение", func(s string) bool {
		_, err := env.Get(ast.UniqueNames.Set(s))
		return err == nil
	})

	env.DefineS("паника", func(e interface{}) {
		// os.Setenv("GONEC_DEBUG", "1")
		panic(e)
	})

	env.DefineS("вывести", env.Print)
	env.DefineS("сообщить", env.Println)
	env.DefineS("сообщитьф", env.Printf)
	env.DefineS("stdout", env.StdOut())
	env.DefineS("закрыть", func(e interface{}) {
		reflect.ValueOf(e).Close()
	})
	env.DefineS("обработатьгорутины", runtime.Gosched)

	env.DefineTypeS("целоечисло", int64(0))
	env.DefineTypeS("число", float64(0.0))
	env.DefineTypeS("булево", true)
	env.DefineTypeS("строка", "")
	env.DefineTypeS("массив", VMSlice{})
	env.DefineTypeS("структура", VMStringMap{})
	env.DefineTypeS("дата", VMTime{})

	//////////////////
	env.DefineTypeS("__функциональнаяструктуратест__", TttStructTest{})
	env.DefineS("__дамп__", env.Dump)
	/////////////////////

	return env
}
