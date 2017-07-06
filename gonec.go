package gonec

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type interpreter struct {
	sync.RWMutex
	connCount int64
	query     []byte
}

// Interpreter возвращает новый интерпретатор
func Interpreter() interpreter {
	return interpreter{}
}

// handlerMain обрабатывает входящие запросы к интерпретатору через POST-запросы
// тело запроса - это код для интерпретации
func (i *interpreter) handlerMain(w http.ResponseWriter, r *http.Request) {

	i.RLock()
	overconn := i.connCount > 5 //лимит количества одновременно работающих интерпретаторов
	i.RUnlock()
	if overconn{
		http.Error(w, "Слишком много запросов обрабатывается в данный момент", http.StatusForbidden)
		return
	}
	i.Lock()
	i.connCount++
	i.Unlock()

	if r.ContentLength > 1<<26 {
		http.Error(w, "Слишком большой запрос", http.StatusForbidden)
		return
	}
	switch r.Method {

	case http.MethodPost:

		defer r.Body.Close()
		//интерпретируется код и возвращается вывод как простой текст
		w.Header().Set("Content-Type", "text/plain")
		err := i.ParseAndRun(r.Body, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
}

// Run запускает микросервис интерпретатора по адресу и порту
func (i *interpreter) Run(srv string) {
	http.HandleFunc("/", i.handlerMain)
	log.Fatal(http.ListenAndServe(srv, nil))
}

// ParseAndRun разбирает запрос и запускает интерпретацию
func (i *interpreter) ParseAndRun(r io.Reader, w io.Writer) (err error) {
	//если r==nil, то разбирается старый запрос, иначе заполняется новый
	if r != nil {
		i.query, err = ioutil.ReadAll(r)
		if err != nil {
			return
		}
	}

	//TODO: синхронно запускается код модуля, но он может создавать вэб-сервера и горутины, которые будут работать и после возврата

	return nil
}
