package gee

import (
	"fmt"
	"log"
	"net/http"
)

/* 这是本框架的入口方法 */
// 分组数据由engine保存.
// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New is the constructor of gee.Engine
func New() *Engine {

	engine := &Engine{router: newRouter()}
	// 初始化,看看删掉会出现什么问题
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
// 如果调用同一个RouterGroup的Group方法,就可以实现,分组嵌套的功能.
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 加入路由分组控制
/*
某一组路由需要相似的处理。例如：
以/post开头的路由匿名可访问。
以/admin开头的路由需要鉴权。
以/api开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权。
大部分情况下的路由分组，是以相同的前缀来区分的。因此，我们今天实现的分组控制也是以前缀来区分，并且支持分组的嵌套。例如/post是一个分组，/post/a和/post/b可以是该分组下的子分组。
*/
// prefix 为实现分组嵌套,以及记录分组 url
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFuncution // support middleware
	engine      *Engine            // all groups share a Engine instance
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

// 本函数将被用作是否为合理的处理函数
// HandlerFunc defines the request handler used by gee
type HandlerFuncution func(c *Context)

// 分组处理器由分组控制.
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFuncution) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFuncution) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFuncution) {
	group.addRoute("POST", pattern, handler)
}
