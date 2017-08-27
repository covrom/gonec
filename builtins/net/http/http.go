// +build !appengine

// Package http implements http interface for anko script.
package http

import (
	"errors"
	h "net/http"
	"reflect"

	envir "github.com/covrom/gonec/env"
)

type Client struct {
	c *h.Client
}

func (c *Client) Get(args ...reflect.Value) (reflect.Value, error) {
	if len(args) < 1 {
		return envir.NilValue, errors.New("Missing arguments")
	}
	if len(args) > 1 {
		return envir.NilValue, errors.New("Too many arguments")
	}
	if args[0].Kind() != reflect.String {
		return envir.NilValue, errors.New("Argument should be string")
	}
	res, err := h.Get(args[0].String())
	return reflect.ValueOf(res), err
}

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("http")
	m.DefineS("DefaultClient", h.DefaultClient)
	m.DefineS("NewServeMux", h.NewServeMux)
	m.DefineS("Handle", h.Handle)
	m.DefineS("HandleFunc", h.HandleFunc)
	m.DefineS("ListenAndServe", h.ListenAndServe)
	return m
}
