package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 一个包含了从HTTP请求发送到离开的所有
// 要想被包外引用,必须使用
type Context struct {
	// 响应
	Response http.ResponseWriter
	// 请求
	Request *http.Request
	// url
	Url string
	// 方法
	Method string
	// 响应码
	StatusCode int
}

/* 初始化context */
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{Response: w, Request: r, Method: r.Method, Url: r.URL.Path}

}

// 写入 响应码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Response.WriteHeader(code)
}

// 写入header
func (c *Context) SetHeader(key string, value string) {
	c.Response.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Response.Write([]byte(fmt.Sprintf(format, values...)))
}

// 返回json格式
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)

	encoder := json.NewEncoder(c.Response)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}

//返回体写入
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Response.Write(data)
}

//返回HTMl
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Response.Write([]byte(html))
}
