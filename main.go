// +build !appengine

package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
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

	"github.com/covrom/gonec/parser"
	"github.com/covrom/gonec/vm"
	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/go-isatty"

	_ "net/http/pprof"
	
	gonec_core "github.com/covrom/gonec/builtins"

)

const version = "1.6a"
const APIPath = "/gonec"

var (
	fs          = flag.NewFlagSet(os.Args[0], 1)
	line        = fs.String("e", "", "Исполнение одной строчки кода")
	testingMode = fs.Bool("t", false, "Режим вывода отладочной информации")
	v           = fs.Bool("v", false, "Версия программы")
	w           = fs.Bool("web", false, "Запустить вэб-сервер на порту 5000, если не указан параметр -p")
	port        = fs.String("p", "", "Номер порта вэб-сервера")

	istty = isatty.IsTerminal(os.Stdout.Fd())

	fsArgs []string

	sessions     = map[string]*vm.Env{}
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

	interactive := fs.NArg() == 0 && *line == ""
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

	env := vm.NewEnv()
	env.DefineS("аргументызапуска", fsArgs)
	gonec_core.LoadAllBuiltins(env)

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

		stmts, err := parser.ParseSrc(code)

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

		following = false
		code = ""
		// v := vm.NilValue

		if err == nil {
			// v, err = vm.Run(stmts, env)
			_, err = vm.Run(stmts, env)
		}
		if err != nil {
			colortext(ct.Red, false, func() {
				if e, ok := err.(*vm.Error); ok {
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

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

func mustGenerateRandomString(s int) string {
	b, _ := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)
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
			sid = mustGenerateRandomString(32)
		}

		lockSessions.RLock()
		env, ok := sessions[sid]
		lockSessions.RUnlock()
		if !ok {

			//создаем новое окружение
			env = vm.NewEnv()
			env.DefineS("аргументызапуска", fsArgs)
			gonec_core.LoadAllBuiltins(env)

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
	fmt.Fprint(w, `<!doctype html>
	<html lang="ru">
	<head>	
		<meta charset="utf-8">
		<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge" /><![endif]-->
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
		
		<title>Интерпретатор Гонец</title>

		<style type="text/css">
			#head {
				float: left;
				height: 45px;
				display:flex;
				align-items:center;
			}
			.header {
				color: #7F6C5F;
				font-size: 20px;
				font-family: sans-serif;
			}
			input[type=button] {
				margin: 10px;
				height: 30px;
				border: 1px solid #FF7822;
				font-size: 16px;
				font-family: sans-serif;
				background: #FF7822;
				color: #FFECDF;
				position: static;
				top: 1px;
				border-radius: 5px;
			}
			#wrap, #about {
				margin: 10px;
				position: absolute;
				top: 45px;
				bottom: 25%;
				left: 0;
				right: 0;
				background: #FFECDF;
			}
			#wrapout {
				margin: 10px;
				position: absolute;
				top: 75%;
				bottom: 0;
				left: 0;
				right: 0;
				background: #7F6C5F;
				border: none;
			}
			#code, #output, pre, .lines {
				font-family: Consolas, Roboto Mono, Menlo, monospace;
				font-size: 11pt;
			}			
			#code, #output {
				border-width: 0;
				background: inherit;
				width: 100%;
				height: 100%;
				margin: 0;
				outline: none;
			}
			#code {
				color: black;
			}
			#output {
				color: white;
			}
			#output .system, #output .loading {
				color: #999;
			}
			#output .stderr, #output .error {
				color: #900;
			}
			
		</style>

		<script type='text/javascript'>
			$(document).ready(function() {
				$('#code').attr('wrap', 'off');
				$('#run').click(function(){
					var body = $("textarea#code").val();
					$.ajax('/gonec', {
						type: 'POST',
						data: body,
						processData : false,
						dataType: 'text',
						cache: false,
						beforeSend: function(xhr){
							xhr.overrideMimeType("text/plain");
							xhr.setRequestHeader('Sid', $("#sid").val());
						},
						success: function(data, textStatus, request) {
							$("#output").text(data);
							$("#sid").val(request.getResponseHeader('Sid'));
						},
						error: function(xhr, status, error) {
							$("#output").text(xhr.responseText);
						}
					});
				});
			});
		</script>
	</head>
	<body bgcolor=#CCBDB3>
		<div id="head" itemprop="name">	
		<a href="https://github.com/covrom/gonec/wiki" style="text-decoration: none">
		<img id="headimg" alt="Интерпретатор ГОНЕЦ" src="data:image/png;base64,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"/>
		</a>
		<input type="button" value="Выполнить" id="run">
		<span class="header"><b>v`+version+`</b></span></div>
		<div id="wrap">
			<textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off"></textarea>
		</div>
		<div id="wrapout">
		<textarea id="output" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off" readonly></textarea>
		</div>
		<input type="hidden" id="sid" name="sid" value="">
	</body>
	`)
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

func ParseAndRun(r io.Reader, w io.Writer, env *vm.Env) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	parser.EnableErrorVerbose()

	sb := string(b)

	ls := fmt.Sprintf("--Выполняется код-- %s\n%s\n", env.GetSid(), sb)

	//замер производительности
	tstart := time.Now()
	stmts, err := parser.ParseSrc(sb)
	tsParse := time.Since(tstart)

	if err != nil {
		return err
	}

	var rb bytes.Buffer
	env.SetStdOut(&rb)

	tstart = time.Now()
	_, err = vm.Run(stmts, env)
	tsRun := time.Since(tstart)

	if err != nil {
		if e, ok := err.(*vm.Error); ok {
			env.Printf("Ошибка исполнения:%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
		} else if e, ok := err.(*parser.Error); ok {
			env.Printf("Ошибка в коде:%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
		} else {
			env.Println(err)
		}
	}

	if *testingMode {
		env.Printf("Время компиляции: %v\n", tsParse)
		env.Printf("Время исполнения: %v\n", tsRun)
	}

	log.Printf("%s--Результат выполнения кода--\n%s\n", ls, rb.String())

	_, err = w.Write(rb.Bytes())

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
