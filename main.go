package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)


/*
定义root，即是restfulapi的格式
资源地址和动作*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"BookIndex", "GET", "/books", BookIndex},
		Route{"Bookshow", "GET", "/books/:isdn", BookShow},
		Route{"Bookshow", "POST", "/books", BookCreate},
	}
	return routes
}

//创建router，给所有route匹配函数
func NewRouter(routes Routes) *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = Logger(handle)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}


func main() {
	router := NewRouter(AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
