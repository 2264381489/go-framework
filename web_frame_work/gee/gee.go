package gee

import (
	"fmt"
	"net/http"
)

/* 这是本框架的入口方法 */

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 加入路由，将方法以method和url的组合为key，执行函数为value生成一个map
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFuncution) {

	engine.router.addRoute(method, pattern, handler)
}

//添加GET方法
// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFuncution) {
	engine.addRoute("GET", pattern, handler)
}

//添加post方法
// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFuncution) {
	engine.addRoute("POST", pattern, handler)
}

//服务启动
// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	fmt.Println("server is Runing")
	return http.ListenAndServe(addr, engine)
}

//实现的处理器。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
