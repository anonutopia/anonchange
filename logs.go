package main

import (
	"log"
	"runtime"
)

func mylog() {
	pc, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] in %s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)
}
