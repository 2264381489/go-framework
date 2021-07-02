package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	// 参数
	params map[string][]string
}

/* 初始化context */
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	url := FormatUrl(r.URL.Path)
	return &Context{Response: w, Request: r, Method: r.Method, Url: url}

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

// 解析 post 参数
func (c *Context) PostForm(key string) string {
	// fromvalue会自动调用ParseForm. https://golang.org/pkg/net/http/#Request.FormValue
	return c.Request.FormValue(key)
}

// 获取某个 post 参数的值
func (c *Context) PostFormMuli(key string) (string, error) {
	err := c.Request.ParseForm()
	if err != nil {
		return "", errors.New("加载失败")
	}
	return c.Request.Form.Get(key), nil
}

// 返回全部 post 参数
func (c *Context) PostFormValues() map[string][]string {
	c.params = c.Request.Form
	return c.Request.Form
}

// 解析 get 参数
//解析url 如将 https://example.org/?a=1&a=2&b=&=3&&&& 取出里面的请求参数组成一个map ["a"]=[1,2]因为有两个a ["b"]= 不可识别的自动跳过 [""]=3 注意 &=3
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// 严格版解析,可以发现url中的错误
func (c *Context) ParseQuery(key string) (queryValue string, err error) {
	v, err := url.ParseQuery(key) // https://golang.org/pkg/net/url/#ParseQuery
	if err != nil {
		err = errors.New("url解析失败")
		return
	}
	queryValue = v.Get(key)
	return
}
