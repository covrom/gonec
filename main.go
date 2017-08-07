// +build !appengine

package main

import (
	"bufio"
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

	gonec_core "github.com/covrom/gonec/builtins"
)

const version = "1.0a"
const APIPath = "/gonec"

var (
	fs   = flag.NewFlagSet(os.Args[0], 1)
	line = fs.String("e", "", "Исполнение одной строчки кода")
	v    = fs.Bool("v", false, "Версия программы")
	w    = fs.Bool("web", false, "Запустить вэб-сервер на порту 5000, если не указан параметр -p")
	port = fs.String("p", "", "Номер порта вэб-сервера")

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

	penv := os.Getenv("GONEC_WEB")
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

	env := vm.NewEnv(os.Stdout)
	env.Define("args", fsArgs)
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
			env = vm.NewEnv(w)
			env.Define("args", fsArgs)
			gonec_core.LoadAllBuiltins(env)

			lockSessions.Lock()
			sessions[sid] = env
			lastAccess[sid] = time.Now()
			lockSessions.Unlock()
			w.Header().Set("Newsid", "true")
		} else {
			lockSessions.Lock()
			lastAccess[sid] = time.Now()
			env.SetStdOut(w)
			lockSessions.Unlock()
		}

		w.Header().Set("Sid", sid)
		
		log.Println("Сессия:",sid)

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
		
		<title>Интерпретатор языка Гонец</title>

		<style type="text/css">
			#head {
				float: left;
				padding: 15px 10px;
				font-size: 20px;
				font-family: sans-serif;
			}
			input[type=button] {
				margin: 10px;
				height: 30px;
				border: 1px solid #375EAB;
				font-size: 16px;
				font-family: sans-serif;
				background: #375EAB;
				color: white;
				position: static;
				top: 1px;
				border-radius: 5px;
			}
			#wrap, #about {
				padding: 5px;
				margin: 10px;
				position: absolute;
				top: 40px;
				bottom: 25%;
				left: 0;
				right: 0;
				background: #FFD;
			}
			#code, #output, pre, .lines {
				font-family: Consolas, Roboto Mono, Menlo, monospace;
				font-size: 11pt;
			}			
			#code {
				color: black;
				background: inherit;
				width: 100%;
				height: 100%;
				margin: 0;
				outline: none;
			}
			#output {
				position: absolute;
				top: 75%;
				bottom: 0;
				left: 0;
				right: 0;
				padding: 10px;
				margin: 10px;
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
	<body>
		<div id="head" itemprop="name">Интерпретатор ГОНЕЦ `+version+`</div>
		<div id="wrap">
			<textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false" wrap="off"></textarea>
		</div>
		<input type="button" value="Выполнить" id="run">
		<br>
		<pre><div id="output"></div></pre>
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
	log.Println("Запущен сервер на порту",srv)
	log.Fatal(http.ListenAndServe(":"+srv, nil))
}

func ParseAndRun(r io.Reader, w io.Writer, env *vm.Env) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	parser.EnableErrorVerbose()

	sb:=string(b);

	log.Println("--Выполняется код--")
	log.Println(sb)

	stmts, err := parser.ParseSrc(sb)
	if err != nil {
		return err
	}
	_, err = vm.Run(stmts, env)
	log.Println("--Завершено выполнение кода--")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
