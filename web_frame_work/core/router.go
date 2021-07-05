package core

/* 框架中只有一个路由 */
type router struct {
	// 核心
	core map[string]HandlerFuncution
}

func newRouter() *router {
	return &router{
		core: make(map[string]HandlerFuncution),
	}
}

//
type HandlerFuncution func(c *Context) (interface{}, error)

func (m *router) addRouter(method, url string, handler HandlerFuncution) {

	key := method + "_" + url

	m.core[key] = handler
}

/* 访问请求 */

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
