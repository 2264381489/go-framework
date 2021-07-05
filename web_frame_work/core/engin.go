package core

import (
	"encoding/json"
	"go-frameWork/web_frame_work/model"
	"log"
	"net/http"
	"strings"
)

type engine struct {
	router       *router
	*RouterGroup // 提供routergroup 功能,外加初始routerGroup,其余的routerGroup都是由这个创建的
	routerGroup  []*RouterGroup
}

/* 核心 */

/* 框架初始化 */
func New() *engine {
	engine := &engine{
		router: newRouter(),
	}
	// $root 为 所有路由的根节点, 并且隐藏
	routerGroup := &RouterGroup{prefix: "$root", engine: engine}
	// 初始的routerGroup,根节点
	engine.RouterGroup = routerGroup
	engine.routerGroup = append(engine.routerGroup, routerGroup)
	// engine
	return engine
}

/* 处理访问请求 */
func (e *engine) Handler(c *Context) {
	fun := "handler-->"
	handler, ok := e.router.core[c.Method+"_"+c.Url]

	if ok {
		c.Next()
		res, err := handler(c)
		if err != nil {
			c.JSON(500, res)
		}
		c.JSON(200, res)
	} else {
		c.Next()
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
	/* 封装中间件 */
	var middlewares []HandlerFuncution
	// 遍历routergroup中的每一个路由组 确定该请求 使用的中间件
	//
	for _, group := range e.routerGroup {
		url := strings.TrimPrefix(group.prefix, "$root")
		if strings.HasPrefix(r.URL.Path, url) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	/* 封装请求 */
	c := newContext(w, r)
	c.handlers = middlewares
	e.Handler(c)
}

/* 开启框架 */
func (e *engine) Run(addr string) error {
	log.Println("最终结果")
	log.Println(e.router.core)
	return http.ListenAndServe(addr, e)
}
