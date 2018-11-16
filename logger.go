package main

import (
	"log"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
)

// 将一些必要的中间过程输出，方便查看调试
func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		fn(w, r, param)
		log.Printf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}
