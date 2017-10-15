package gonecsvc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/covrom/gonec/bincode"
	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/core"
	"github.com/covrom/gonec/parser"
)

func NewGonecInterpreter(header core.VMServiceHeader, args []string, tmode bool) *VMGonecInterpreterService {
	v := &VMGonecInterpreterService{
		hdr:          header,
		fsArgs:       args,
		testingMode:  tmode,
		sessions:     make(map[string]*core.Env),
		lastAccess:   make(map[string]time.Time),
		lockSessions: sync.RWMutex{},
		srv:          nil,
		lasterr:      core.VMErrorServiceNotReady,
	}
	return v
}

type VMGonecInterpreterService struct {
	core.VMValueStruct
	hdr          core.VMServiceHeader
	fsArgs       []string
	testingMode  bool
	sessions     map[string]*core.Env
	lastAccess   map[string]time.Time
	lockSessions sync.RWMutex
	srv          *http.Server
	lasterr      error
}

func (x *VMGonecInterpreterService) vmval() {}

func (x *VMGonecInterpreterService) Header() core.VMServiceHeader {
	return x.hdr
}

func (x *VMGonecInterpreterService) Start() error {

	if x.srv != nil {
		return core.VMErrorServerAlreadyStarted
	}

	http.HandleFunc("/", x.handlerIndex)
	http.HandleFunc("/"+x.hdr.Path, x.handlerAPI)
	http.HandleFunc("/"+x.hdr.Path+"/src", x.handlerSource)
	http.HandleFunc("/"+x.hdr.Path+"/healthcheck", x.handlerHealth) // в таком же формате регистрируется в consul и т.п.

	//добавляем горутину на принудительное закрытие сессий через 10 мин без активности
	go func() {
		for {
			time.Sleep(time.Minute)
			x.lockSessions.Lock()
			for id, lat := range x.lastAccess {
				if time.Since(lat) >= 10*time.Minute {
					delete(x.sessions, id)
					delete(x.lastAccess, id)
					log.Println("Закрыта сессия Sid=" + id)
				}
			}
			x.lockSessions.Unlock()
		}
	}()

	x.srv = &http.Server{
		Addr:    ":" + x.hdr.Port,
		Handler: nil,
	}

	x.lasterr = nil

	// горутина сервера
	go func(c *VMGonecInterpreterService) {
		lasterr := c.srv.ListenAndServe()
		c.lockSessions.Lock()
		c.lasterr = lasterr
		c.srv = nil
		c.lockSessions.Unlock()
		log.Println("Остановлен сервер интерпретатора на порту", c.hdr.Port)
		log.Println(lasterr)
	}(x)

	log.Println("Запущен сервер интерпретатора на порту", x.hdr.Port)

	return nil
}

func (x *VMGonecInterpreterService) HealthCheck() error {
	x.lockSessions.RLock()
	defer x.lockSessions.RUnlock()

	return x.lasterr
}

func (x *VMGonecInterpreterService) Stop() error {
	x.lockSessions.Lock()
	defer x.lockSessions.Unlock()

	if x.lasterr != nil && x.srv == nil {
		return x.lasterr
	}
	if x.srv != nil {
		x.lasterr = x.srv.Close()
		return x.lasterr
	}
	return nil
}

func (x *VMGonecInterpreterService) handlerSource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer r.Body.Close()
		srcname := r.URL.Query().Get("name")
		if srcname != "" {
			switch srcname {
			case "jquery":
				w.Header().Set("Content-Type", "text/javascript")
				_, err := w.Write([]byte(jQuery))
				if err != nil {
					time.Sleep(time.Second) //анти-ddos
					http.Error(w, err.Error(), http.StatusInternalServerError)
					log.Println(err)
					return
				}
			case "ace":
				w.Header().Set("Content-Type", "text/javascript")
				_, err := w.Write([]byte(jsAce))
				if err != nil {
					time.Sleep(time.Second) //анти-ddos
					http.Error(w, err.Error(), http.StatusInternalServerError)
					log.Println(err)
					return
				}
			case "acetheme":
				w.Header().Set("Content-Type", "text/javascript")
				_, err := w.Write([]byte(jsAceTheme))
				if err != nil {
					time.Sleep(time.Second) //анти-ddos
					http.Error(w, err.Error(), http.StatusInternalServerError)
					log.Println(err)
					return
				}
			case "acelang":
				w.Header().Set("Content-Type", "text/javascript")
				_, err := w.Write([]byte(jsAceLang))
				if err != nil {
					time.Sleep(time.Second) //анти-ddos
					http.Error(w, err.Error(), http.StatusInternalServerError)
					log.Println(err)
					return
				}
			default:
				http.Error(w, "Неправильно указано имя", http.StatusBadRequest)
			}

		} else {
			http.Error(w, "Не указано имя", http.StatusBadRequest)
		}

	default:
		time.Sleep(time.Second) //анти-ddos
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
}

func (x *VMGonecInterpreterService) handlerAPI(w http.ResponseWriter, r *http.Request) {

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

		x.lockSessions.RLock()
		env, ok := x.sessions[sid]
		x.lockSessions.RUnlock()
		if !ok {

			//создаем новое окружение
			env = core.NewEnv()
			env.DefineS("аргументызапуска", core.NewVMSliceFromStrings(x.fsArgs))

			x.lockSessions.Lock()
			x.sessions[sid] = env
			x.lastAccess[sid] = time.Now()
			x.lockSessions.Unlock()
			w.Header().Set("Newsid", "true")
		} else {
			x.lockSessions.Lock()
			x.lastAccess[sid] = time.Now()
			x.lockSessions.Unlock()
		}

		w.Header().Set("Sid", sid)

		env.SetSid(sid)
		//log.Println("Сессия:",sid)

		err := x.parseAndRun(r.Body, w, env)

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

func (x *VMGonecInterpreterService) handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexPage)
}

func (x *VMGonecInterpreterService) parseAndRun(r io.Reader, w io.Writer, env *core.Env) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	parser.EnableErrorVerbose()

	sb := string(b)

	if x.testingMode {
		log.Printf("--Выполняется код-- %s\n%s\n", env.GetSid(), sb)
	}

	//замер производительности
	tstart := time.Now()
	_, bins, err := bincode.ParseSrc(sb)
	tsParse := time.Since(tstart)

	if x.testingMode {
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

	if x.testingMode {
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

func (x *VMGonecInterpreterService) handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// log.Println("Healthcheck")
}
