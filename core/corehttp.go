package core

import (
	"fmt"
	"net/http"

	"github.com/covrom/gonec/names"
)

type VMHttpRequest struct {
	r    *http.Request
	data VMValuer
}

func (x *VMHttpRequest) vmval() {}

func (x *VMHttpRequest) Interface() interface{} {
	return x.r
}

func (x *VMHttpRequest) String() string {
	// TODO: представление строкой
	return fmt.Sprintf("Запрос %s", x.r)
}

func (x *VMHttpRequest) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	// TODO: параметр() из data

	// case "получить":
	// 	return VMFuncMustParams(0, x.Получить), true
	}

	return nil, false
}

type VMHttpResponse struct {
	r    *http.Response
	w    http.ResponseWriter
	data VMValuer
}

func (x *VMHttpResponse) vmval() {}

func (x *VMHttpResponse) Interface() interface{} {
	return x.r
}

func (x *VMHttpResponse) String() string {
	return fmt.Sprintf("Ответ %s", x.r)
}

func (x *VMHttpResponse) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	// case "получить":
	// 	return VMFuncMustParams(0, x.Получить), true
	}

	return nil, false
}
