package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/covrom/gonec/bincode"
	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/core"
	"github.com/covrom/gonec/parser"
	"github.com/covrom/gonec/services/gonecsvc"
	"github.com/covrom/gonec/version"
	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/go-isatty"
	uuid "github.com/satori/go.uuid"

	_ "net/http/pprof"
)

var (
	fs          = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	line        = fs.String("e", "", "Исполнение одной строчки кода")
	compile     = fs.Bool("c", false, "Компиляция в файл .gnx")
	testingMode = fs.Bool("t", false, "Режим вывода отладочной информации")
	toconsul    = fs.Bool("consul", false, "Зарегистрировать микросервис интерпретатора в Consul")
	// stackvm     = fs.Bool("stack", false, "Старая стековая виртуальная машина версии 1.8b")
	v    = fs.Bool("v", false, "Версия программы")
	w    = fs.Bool("web", false, "Запустить вэб-сервер на порту 5000, если не указан параметр -p")
	port = fs.String("p", "", "Номер порта вэб-сервера")

	istty = isatty.IsTerminal(os.Stdout.Fd())

	fsArgs []string
)

func colortext(color ct.Color, bright bool, f func()) {
	if istty {
		ct.ChangeColor(color, bright, ct.None, false)
		f()
		ct.ResetColor()
	} else {
		f()
	}
}

func main() {

	fs.Parse(os.Args[1:])
	if *v {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	var (
		code      string
		b         []byte
		reader    *bufio.Reader
		following bool
		source    string
	)

	interactive := fs.NArg() == 0 && *line == "" && !*compile
	fsArgs = fs.Args()

	ext := ""
	if *toconsul {
		ext = "consul"
	}

	// если есть PORT в переменных окружения - сразу стартуем сервис
	// это нужно для развертывания в Docker контейнере
	penv := os.Getenv("PORT")
	if penv != "" {
		Run(penv, ext)
		return
	}

	// если есть -web в ключах запуска (и, возможно, -p порт) - сразу стартуем сервис
	// это нужно для развертывания в Docker контейнере
	if *w {
		if *port == "" {
			*port = "5000"
		}
		Run(*port, ext)
		return
	}

	// иначе - запуск из командной строки

	if interactive {
		reader = bufio.NewReader(os.Stdin)
		source = "typein"
		os.Args = append([]string{os.Args[0]}, fs.Args()...)
	} else {
		if *line != "" {
			b = []byte(*line)
			source = "argument"
		} else {
			var err error
			b, err = ioutil.ReadFile(fs.Arg(0))
			if err != nil {
				colortext(ct.Red, false, func() {
					fmt.Fprintln(os.Stderr, err)
				})
				os.Exit(1)
			}
			fsArgs = fs.Args()[1:]
			source = filepath.Clean(fs.Arg(0))
		}
		os.Args = fs.Args()
	}

	env := core.NewEnv()
	env.DefineS("аргументызапуска", core.NewVMSliceFromStrings(fsArgs))

	for {
		if interactive {
			colortext(ct.Green, true, func() {
				if following {
					fmt.Print("  ")
				} else {
					fmt.Print("> ")
				}
			})
			var err error
			b, _, err = reader.ReadLine()
			if err != nil {
				break
			}
			if len(b) == 0 {
				continue
			}
			if code != "" {
				code += "\n"
			}
			code += string(b)
		} else {
			code = string(b)
		}

		parser.EnableErrorVerbose()

		var (
			bins binstmt.BinCode
			// stmts          ast.Stmts
			err            error
			tstart         time.Time
			tsParse, tsRun time.Duration
		)

		tstart = time.Now()

		isGNX := strings.HasSuffix(strings.ToLower(source), ".gnx")
		// если это скомпилированный файл, то сразу его выполняем
		if isGNX {
			bbuf := bytes.NewBuffer(b)
			// stmts = nil
			bins, err = binstmt.ReadBinCode(bbuf)
			tsParse = time.Since(tstart)
			if err != nil {
				log.Fatal(err)
			}
			if *testingMode {
				log.Printf("--Выполняется скомпилированный код-- \n%s\n", bins.String())
			}
		} else {
			if *testingMode {
				log.Printf("--Выполняется код--\n%s\n", code)
			}
			//замер производительности
			_, bins, err = bincode.ParseSrc(code)
			tsParse = time.Since(tstart)

			if *testingMode {
				log.Printf("--Скомпилирован код-- \n%s\n", bins.String())
			}

			if interactive {
				if e, ok := err.(*parser.Error); ok {
					es := e.Error()
					if strings.HasPrefix(es, "syntax error: unexpected") {
						if strings.HasPrefix(es, "syntax error: unexpected $end,") {
							following = true
							continue
						}
					} else {
						if e.Pos.Column == len(b) && !e.Fatal {
							println(e.Error())
							following = true
							continue
						}
						if e.Error() == "unexpected EOF" {
							following = true
							continue
						}
					}
				}
			}
		}

		if *compile {
			srcname := fs.Arg(0)
			if srcname != "" && !isGNX {
				if strings.HasSuffix(strings.ToLower(srcname), ".gnc") {
					srcname = srcname[:len(srcname)-4]
				}
				compilename := srcname + ".gnx"
				fo, err := os.Create(compilename)
				if err != nil {
					log.Fatal(err)
				}
				defer func() {
					if err := fo.Close(); err != nil {
						log.Fatal(err)
					}
				}()
				if err := binstmt.WriteBinCode(fo, bins); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal("Не указано имя файла с исходным кодом на языке Гонец")
			}
			break
		}

		following = false
		code = ""

		//замер производительности
		tstart = time.Now()
		if *testingMode {
			log.Println("--Результат выполнения кода--")
		}

		// v := vm.NilValue

		if err == nil {
			// v, err = vm.Run(stmts, env)
			// if *stackvm && stmts != nil {
			// 	_, err = vm.Run(stmts, env)
			// } else {
			_, err = bincode.Run(bins, env)
			// }
		}

		tsRun = time.Since(tstart)

		if *testingMode {
			env.Printf("Время компиляции: %v\n", tsParse)
			env.Printf("Время исполнения: %v\n", tsRun)
		}

		if err != nil {
			colortext(ct.Red, false, func() {
				if e, ok := err.(*binstmt.Error); ok {
					fmt.Fprintf(os.Stderr, "%s:%d:%d %s\n", source, e.Pos.Line, e.Pos.Column, err)
				} else if e, ok := err.(*parser.Error); ok {
					if e.Filename != "" {
						source = e.Filename
					}
					fmt.Fprintf(os.Stderr, "%s:%d:%d %s\n", source, e.Pos.Line, e.Pos.Column, err)
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			})

			if interactive {
				continue
			} else {
				os.Exit(1)
			}
		} else {
			if interactive {
				// colortext(ct.Black, true, func() {
				// 	if v == vm.NilValue || !v.IsValid() {
				// 		fmt.Println("nil")
				// 	} else {
				// 		s, ok := v.Interface().(fmt.Stringer)
				// 		if v.Kind() != reflect.String && ok {
				// 			fmt.Println(s)
				// 		} else {
				// 			fmt.Printf("%#v\n", v.Interface())
				// 		}
				// 	}
				// })

			} else {
				break
			}
		}
	}
}

// Run запускает микросервис интерпретатора на порту
func Run(port string, ext string) {

	// создаем сервис
	svc := gonecsvc.NewGonecInterpreter(
		core.VMServiceHeader{
			ID:       uuid.NewV4().String(),
			Path:     "gonec",
			Name:     "Интерпретатор Гонец",
			Port:     port,
			External: ext,
		}, fsArgs, *testingMode)

	// регистрируем
	err := core.VMMainServiceBus.Register(svc)
	if err != nil {
		log.Println(err)
	}

	// запускаем все сервисы
	core.VMMainServiceBus.Run()

	// ждем окончания работы всех сервисов
	core.VMMainServiceBus.WaitForAll()

	// дерегистрируем сервис
	err = core.VMMainServiceBus.Deregister(svc)
	if err != nil {
		log.Println(err)
	}

}
