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
type HandlerFuncution func(c *Context)

func (m *router) addRouter(method, url string, handler HandlerFuncution) {

	key := method + "_" + url

	m.core[key] = handler
}
