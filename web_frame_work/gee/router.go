package gee

import (
	"log"
	"net/http"
)

/* 将router功能提取出来 */
/* 路由基础的功能,只需要构建一个可以存储所有url和处理函数(Handler)的Map 以及配套的增加方法 和 get方法(寻找目标url的执行函数的方法)*/

type router struct {
	handlers map[string]HandlerFuncution
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFuncution)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFuncution) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// 本函数将被用作是否为合理的处理函数
// HandlerFunc defines the request handler used by gee
type HandlerFuncution func(c *Context)
