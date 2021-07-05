package core

import (
	"fmt"
	"strings"
)

type RouterGroup struct {
	// 前缀 以此形成嵌套 的 路由组
	prefix string
	// 为了获得router
	engine      *engine
	middlewares []HandlerFuncution // support middleware
}

// 新建 一个路由组
func (rg *RouterGroup) NewGroup(pattern string) *RouterGroup {
	engine := rg.engine
	newGroup := &RouterGroup{
		prefix: rg.prefix + pattern,
		engine: engine,
	}
	engine.routerGroup = append(engine.routerGroup, newGroup)
	return newGroup
}

func (m *RouterGroup) addRouter(method, url string, handler HandlerFuncution) {
	// 去除group自带的头.
	url = strings.TrimPrefix(m.prefix, "$root") + url
	url = strings.TrimSuffix(url, "/")
	key := fmt.Sprintf("%s_%s", method, url)
	m.engine.router.core[key] = handler
}

// 请求访问方法(有分组)
func (m *RouterGroup) GET(url string, handler HandlerFuncution) {
	m.addRouter("GET", url, handler)
}
func (m *RouterGroup) POST(url string, handler HandlerFuncution) {
	m.addRouter("POST", url, handler)
}
func (m *RouterGroup) PUT(url string, handler HandlerFuncution) {
	m.addRouter("PUT", url, handler)
}
func (m *RouterGroup) AddRouter(method, url string, handler HandlerFuncution) {
	m.addRouter(method, url, handler)
}

// 为路由组添加中间件
// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFuncution) {
	group.middlewares = append(group.middlewares, middlewares...)
}

/* 规范url */
/* 主要目的是 防止出现
GET_/V1/
GET_/V1
虽然是同一个url但是匹配不到同一个url的情况
*/
func FormatUrl(url string) string {
	return strings.TrimSuffix(url, "/")
}
