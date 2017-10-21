package core

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/covrom/gonec/names"
)

// VMHttpRequest запрос к http серверу
type VMHttpRequest struct {
	r    *http.Request
	data VMValuer
	body []byte
}

func (x *VMHttpRequest) vmval() {}

func (x *VMHttpRequest) Interface() interface{} {
	return x.r
}

func (x *VMHttpRequest) String() string {
	return fmt.Sprintf("Запрос %s %s %s", x.r.RemoteAddr, x.r.Method, x.r.RequestURI)
}

func (x *VMHttpRequest) GetHeader(name VMString) VMString {
	return VMString(x.r.Header.Get(string(name)))
}

func (x *VMHttpRequest) SetHeader(name, val VMString) {
	x.r.Header.Set(string(name), string(val))
}

func (x *VMHttpRequest) Close() {
	if x.r != nil {
		if x.r.Body != nil {
			x.r.Body.Close()
		}
	}
	x.r = nil
	x.body = nil
}

func (x *VMHttpRequest) ReadBody() (b VMString, err error) {
	if x.body != nil {
		return VMString(x.body), nil
	}
	x.body, err = ioutil.ReadAll(x.r.Body)
	if x.r.Body != nil {
		x.r.Body.Close()
	}
	if err != nil {
		return VMString(""), err
	}
	return VMString(x.body), nil
}

func (x *VMHttpRequest) Path() VMString {
	return VMString(x.r.URL.Path)
}

func (x *VMHttpRequest) RemoteAddr() VMString {
	return VMString(x.r.RemoteAddr)
}

func (x *VMHttpRequest) Fragment() VMString {
	return VMString(x.r.URL.Fragment)
}

func (x *VMHttpRequest) GetParam(name VMString) VMString {
	return VMString(x.r.FormValue(string(name)))
}

func (x *VMHttpRequest) Method() VMString {
	return VMString(x.r.Method)
}

// RequestAsVMStringMap возвращает структуру в формате JSON
// {
//  "Адрес":"IP адрес корреспондента",
//  "Путь":"/root",
//  "Фрагмент":"после#",
//  "Параметры":{"Имя":Значение,...},
//  "ПараметрыФормы":{"Имя":Значение,...},
//  "Метод":Метод,
//  "Заголовки":{"Имя":Значение,...},
//  "Тело":"Строка"
// }
func (x *VMHttpRequest) RequestAsVMStringMap() (VMStringMap, error) {

	rmap := make(VMStringMap)

	err := x.r.ParseMultipartForm(32 << 20)
	// if err != nil {
	// 	return rmap, err
	// }

	rmap["Тело"], err = x.ReadBody()
	if err != nil {
		return rmap, err
	}

	rmap["Адрес"] = x.RemoteAddr()
	rmap["Путь"] = x.Path()
	rmap["Фрагмент"] = x.Fragment()
	// rmap["Данные"] = x.data
	rmap["Метод"] = x.Method()

	m1 := make(VMStringMap)
	for k, v := range x.r.Header {
		if len(v) > 0 {
			m1[k] = VMString(v[0])
		}
	}
	rmap["Заголовки"] = m1

	m2 := make(VMStringMap)
	for k, v := range x.r.Form {
		if len(v) > 0 {
			m2[k] = VMString(v[0])
		}
	}
	rmap["Параметры"] = m2

	m3 := make(VMStringMap)
	for k, v := range x.r.PostForm {
		if len(v) > 0 {
			m3[k] = VMString(v[0])
		}
	}
	rmap["ПараметрыФормы"] = m3

	return rmap, nil
}

func (x *VMHttpRequest) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	case "метод":
		return VMFuncMustParams(0, x.Метод), true
	case "заголовок":
		return VMFuncMustParams(1, x.Заголовок), true
	case "установитьзаголовок":
		return VMFuncMustParams(2, x.УстановитьЗаголовок), true
	case "тело":
		return VMFuncMustParams(0, x.Тело), true
	case "путь":
		return VMFuncMustParams(0, x.Путь), true
	case "адрес":
		return VMFuncMustParams(0, x.Адрес), true
	case "фрагмент":
		return VMFuncMustParams(0, x.Фрагмент), true
	case "параметр":
		return VMFuncMustParams(1, x.Параметр), true
	case "данные":
		return VMFuncMustParams(0, x.Данные), true
	case "сообщение":
		return VMFuncMustParams(0, x.Сообщение), true
	}

	return nil, false
}

func (x *VMHttpRequest) Метод(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(x.Method())
	return nil
}

func (x *VMHttpRequest) Заголовок(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	s, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	rets.Append(x.GetHeader(s))
	return nil
}

func (x *VMHttpRequest) УстановитьЗаголовок(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	k, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	v, ok := args[1].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	x.SetHeader(k, v)
	return nil
}

func (x *VMHttpRequest) Тело(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	s, _ := x.ReadBody()
	rets.Append(VMString(s))
	return nil
}

func (x *VMHttpRequest) Путь(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(x.Path())
	return nil
}

func (x *VMHttpRequest) Адрес(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(x.RemoteAddr())
	return nil
}

func (x *VMHttpRequest) Фрагмент(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(x.Fragment())
	return nil
}

func (x *VMHttpRequest) Параметр(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	s, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	rets.Append(x.GetParam(s))
	return nil
}

func (x *VMHttpRequest) Данные(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(x.data)
	return nil
}

func (x *VMHttpRequest) Сообщение(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, err := x.RequestAsVMStringMap()
	if err != nil {
		return err
	}
	rets.Append(v)
	return nil
}

// VMHttpResponse ответ от http сервера
type VMHttpResponse struct {
	r    *http.Response
	w    http.ResponseWriter
	body []byte
	data VMValuer
}

func (x *VMHttpResponse) vmval() {}

func (x *VMHttpResponse) Interface() interface{} {
	return x.r
}

func (x *VMHttpResponse) String() string {
	return fmt.Sprintf("%s %s", x.r.Status, x.body)
}

func (x *VMHttpResponse) Close() {
	if x.r != nil {
		if x.r.Body != nil {
			x.r.Body.Close()
		}
	}
	x.r = nil
	x.body = nil
}

func (x *VMHttpResponse) ReadBody() (b VMString, err error) {
	if x.body != nil {
		return VMString(x.body), nil
	}
	x.body, err = ioutil.ReadAll(x.r.Body)
	if x.r.Body != nil {
		x.r.Body.Close()
	}
	if err != nil {
		return VMString(""), err
	}
	return VMString(x.body), nil
}

func (x *VMHttpResponse) Send(status VMInt, b VMString, h VMStringMap) error {
	hdrs := x.w.Header()
	for k, v := range h {
		vv, ok := v.(VMStringer)
		if !ok {
			return VMErrorNeedString
		}
		hdrs.Add(k, vv.String())
	}

	x.w.WriteHeader(int(status))

	fmt.Fprintln(x.w, b)
	return nil
}

func (x *VMHttpResponse) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	case "отправить":
		return VMFuncMustParams(1, x.Отправить), true
	case "тело":
		return VMFuncMustParams(0, x.Тело), true
	}

	return nil, false
}

func (x *VMHttpResponse) Отправить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	if x.w == nil || x.w == http.ResponseWriter(nil) {
		return VMErrorHTTPResponseMethod
	}

	vsm, ok := args[0].(VMStringMap)
	if !ok {
		return VMErrorNeedMap
	}

	var h VMStringMap
	if v, ok := vsm["Заголовки"]; ok {
		if h, ok = v.(VMStringMap); !ok {
			return VMErrorNeedMap
		}
	}

	var sts VMInt
	if v, ok := vsm["Статус"]; ok {
		if sts, ok = v.(VMInt); !ok {
			return VMErrorNeedInt
		}
	} else {
		sts = http.StatusOK
	}

	var b VMString
	if v, ok := vsm["Тело"]; ok {
		if b, ok = v.(VMString); !ok {
			return VMErrorNeedString
		}
	}

	return x.Send(sts, b, h)
}

func (x *VMHttpResponse) Тело(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	s, _ := x.ReadBody()
	rets.Append(VMString(s))
	return nil
}
