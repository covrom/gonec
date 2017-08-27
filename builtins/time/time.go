// Package time implements time interface for anko script.
package time

import (
	t "time"

	envir "github.com/covrom/gonec/env"
)

func Import(env *envir.Env) *envir.Env {
	m := env.NewPackage("time")
	m.DefineS("After", t.After)
	m.DefineS("Sleep", t.Sleep)
	m.DefineS("Tick", t.Tick)
	m.DefineS("Since", t.Since)
	m.DefineS("FixedZone", t.FixedZone)
	m.DefineS("LoadLocation", t.LoadLocation)
	m.DefineS("NewTicker", t.NewTicker)
	m.DefineS("Date", t.Date)
	m.DefineS("Now", t.Now)
	m.DefineS("Parse", t.Parse)
	m.DefineS("ParseDuration", t.ParseDuration)
	m.DefineS("ParseInLocation", t.ParseInLocation)
	m.DefineS("Unix", t.Unix)
	m.DefineS("AfterFunc", t.AfterFunc)
	m.DefineS("NewTimer", t.NewTimer)
	m.DefineS("Nanosecond", t.Nanosecond)
	m.DefineS("Microsecond", t.Microsecond)
	m.DefineS("Millisecond", t.Millisecond)
	m.DefineS("Second", t.Second)
	m.DefineS("Minute", t.Minute)
	m.DefineS("Hour", t.Hour)
	return m
}
