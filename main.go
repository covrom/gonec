// +build !appengine

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/covrom/gonec/bincode"
	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/core"
	"github.com/covrom/gonec/parser"
	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/go-isatty"

	_ "net/http/pprof"
)

const version = "3.2a"
const APIPath = "/gonec"

var (
	fs          = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	line        = fs.String("e", "", "Исполнение одной строчки кода")
	compile     = fs.Bool("c", false, "Компиляция в файл .gnx")
	testingMode = fs.Bool("t", false, "Режим вывода отладочной информации")
	// stackvm     = fs.Bool("stack", false, "Старая стековая виртуальная машина версии 1.8b")
	v    = fs.Bool("v", false, "Версия программы")
	w    = fs.Bool("web", false, "Запустить вэб-сервер на порту 5000, если не указан параметр -p")
	port = fs.String("p", "", "Номер порта вэб-сервера")

	istty = isatty.IsTerminal(os.Stdout.Fd())

	fsArgs []string

	sessions     = map[string]*core.Env{}
	lastAccess   = map[string]time.Time{}
	lockSessions = sync.RWMutex{}
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
		fmt.Println(version)
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

	penv := os.Getenv("PORT")
	if penv != "" {
		Run(penv)
		return
	}
	if *w {
		if *port == "" {
			*port = "5000"
		}
		Run(*port)
		return
	}

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

func handlerAPI(w http.ResponseWriter, r *http.Request) {

	if r.ContentLength > 1<<26 {
		time.Sleep(time.Second) //анти-ddos
		http.Error(w, "Слишком большой запрос", http.StatusForbidden)
		return
	}

	switch r.Method {

	case http.MethodPost:

		defer r.Body.Close()

		//интерпретируется код и возвращается вывод как простой текст
		w.Header().Set("Content-Type", "text/plain")

		sid := r.Header.Get("Sid")
		if sid == "" {
			sid = core.MustGenerateRandomString(22)
		}

		lockSessions.RLock()
		env, ok := sessions[sid]
		lockSessions.RUnlock()
		if !ok {

			//создаем новое окружение
			env = core.NewEnv()
			env.DefineS("аргументызапуска", core.NewVMSliceFromStrings(fsArgs))

			lockSessions.Lock()
			sessions[sid] = env
			lastAccess[sid] = time.Now()
			lockSessions.Unlock()
			w.Header().Set("Newsid", "true")
		} else {
			lockSessions.Lock()
			lastAccess[sid] = time.Now()
			lockSessions.Unlock()
		}

		w.Header().Set("Sid", sid)

		env.SetSid(sid)
		//log.Println("Сессия:",sid)

		err := ParseAndRun(r.Body, w, env)

		if err != nil {
			time.Sleep(time.Second) //анти-ddos
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		time.Sleep(time.Second) //анти-ddos
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexPage)
}

// Run запускает микросервис интерпретатора по адресу и порту
func Run(srv string) {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc(APIPath, handlerAPI)
	//добавляем горутину на принудительное закрытие сессий через 10 мин без активности
	go func() {
		for {
			time.Sleep(time.Minute)
			lockSessions.Lock()
			for id, lat := range lastAccess {
				if time.Since(lat) >= 10*time.Minute {
					delete(sessions, id)
					delete(lastAccess, id)
					log.Println("Закрыта сессия Sid=" + id)
				}
			}
			lockSessions.Unlock()
		}
	}()
	log.Println("Запущен сервер на порту", srv)
	log.Fatal(http.ListenAndServe(":"+srv, nil))
}

func ParseAndRun(r io.Reader, w io.Writer, env *core.Env) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	parser.EnableErrorVerbose()

	sb := string(b)

	if *testingMode {
		log.Printf("--Выполняется код-- %s\n%s\n", env.GetSid(), sb)
	}

	//замер производительности
	tstart := time.Now()
	_, bins, err := bincode.ParseSrc(sb)
	tsParse := time.Since(tstart)

	if *testingMode {
		log.Printf("--Скомпилирован код-- \n%s\n", bins.String())
	}

	if err != nil {
		return err
	}

	var rb bytes.Buffer
	env.SetStdOut(&rb)

	tstart = time.Now()
	// if *stackvm {
	// 	_, err = vm.Run(stmts, env)
	// } else {
	_, err = bincode.Run(bins, env)
	// }
	tsRun := time.Since(tstart)

	if err != nil {
		if e, ok := err.(*binstmt.Error); ok {
			env.Printf("Ошибка исполнения: %s\n", e)
		} else if e, ok := err.(*parser.Error); ok {
			env.Printf("Ошибка в коде: %s\n", e)
		} else {
			env.Println(err)
		}
	}

	if *testingMode {
		env.Printf("Время компиляции: %v\n", tsParse)
		env.Printf("Время исполнения: %v\n", tsRun)
		log.Printf("--Результат выполнения кода--\n%s\n", rb.String())
	}

	_, err = w.Write(rb.Bytes())

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
