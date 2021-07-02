package core

import (
	"encoding/json"
	"fmt"
	"go-frameWork/web_frame_work/model"
	"log"
	"net/http"
)

type engine struct {
	router *router
}

/* 框架初始化 */
func New() *engine {
	return &engine{
		router: newRouter(),
	}
}

/* 处理访问请求 */
func (e *engine) Handler(c *Context) {
	fun := "handler-->"
	handler, ok := e.router.core[c.Method+"_"+c.Url]

	if ok {
		res, err := handler(c)
		if err != nil {
			c.JSON(500, res)
		}
		c.JSON(200, res)
	} else {
		res := &model.DefaultRes{Errinfo: &model.Errinfo{Code: 404, Msg: "404 NOT FOUND"}}
		_, err := json.Marshal(res)
		if err != nil {
			log.Printf("%s json marshal fail", fun)
		}

		c.JSON(404, res)
	}
}

/* HTTP请求入口 */
/* ListenAndServe 的第二个参数要求传入一个实现了 Handler接口 的结构体   即实现了ServeHTTP(ResponseWriter, *Request) 的方法*/
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/* 封装请求 */
	c := newContext(w, r)
	fmt.Println(c.Method + "_" + c.Url)
	e.Handler(c)
}

/* 开启框架 */
func (e *engine) Run(addr string) error {
	log.Println("最终结果")
	log.Println(e.router.core)
	return http.ListenAndServe(addr, e)
}

// 请求访问方法(无分组)
func (m *engine) GET(url string, handler HandlerFuncution) {
	m.router.addRouter("GET", url, handler)
}
func (m *engine) POST(url string, handler HandlerFuncution) {
	m.router.addRouter("POST", url, handler)
}
func (m *engine) PUT(url string, handler HandlerFuncution) {
	m.router.addRouter("PUT", url, handler)
}
func (m *engine) AddRouter(method, url string, handler HandlerFuncution) {
	m.router.addRouter(method, url, handler)
}

//
