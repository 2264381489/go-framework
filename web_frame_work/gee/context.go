package gee

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

/* 封装上下文
Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载
*/
/* 一个Context需要什么?
1.作为HTTP框架,就需要request和response.在golang中,
	1.1 request必须是http包中提供的request结构体,这在于http包可以自动解析前端发来的HTTP请求.
	1.2 response则是http包中的ResponseWriter,可以写入数据
2.为了区分前端的HTTP请求路径,需要Path字段
3.为了保证返回的成功与否,要有状态码.
	3.1 因为HTTP的response中 返回体和状态码是两个部分
4.还有HTTP请求的方法.路径里面包含不了方法.
*/
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

/* 需要的方法,1.初始化 2.获取Req中内容 3.获得请求体中内容 4.写入的需求 5.转传承HTML的需求 */
/* 1.初始化  ps:初始化方法,不被加入reciever */
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,          //涉及到的知识: https://golang.org/pkg/net/http/#Request
		Path:   req.URL.Path, //涉及到的知识: https://golang.org/pkg/net/url/#URL
		Method: req.Method,
		// StatueCode应该不必初始化
	}
}

// 获取request中的内容.
// PostForm contains the parsed form data from PATCH, POST
// or PUT body parameters.
// This field is only available after ParseForm is called.
// The HTTP client ignores PostForm and uses Body instead.
// PostForm是在使用PATCH, POST 和 PUT 的时候的请求体(表单),只有在调用了ParseForm方法之后才能使用.
// 其余的? 使用Body这个子弹
func (c *Context) PostForm(key string) string {
	// fromvalue会自动调用ParseForm. https://golang.org/pkg/net/http/#Request.FormValue
	return c.Req.FormValue(key)
}
func (c *Context) PostFormMuli(key string) (string, error) {
	err := c.Req.ParseForm()
	if err != nil {
		return "", errors.New("加载失败")
	}
	return c.Req.Form.Get(key), nil
}

//解析url 如将 https://example.org/?a=1&a=2&b=&=3&&&& 取出里面的请求参数组成一个map ["a"]=[1,2]因为有两个a ["b"]= 不可识别的自动跳过 [""]=3 注意 &=3
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
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

// 写入 响应码
func (c *Context) Status(code int) {
	c.StatusCode = code
	/* WriteHeader发送一个HTTP响应头，其中包括所提供的
	   状态代码。

	   如果WriteHeader没有被显式调用，第一次调用Write
	    将触发一个隐式的 WriteHeader(http.StatusOK)。
	    因此对WriteHeader的显式调用主要用于
	    发送错误代码。

	    提供的代码必须是有效的HTTP 1xx-5xx状态代码。
	   只能写一个头信息。Go目前并不
	   支持发送用户定义的 1xx 信息头。
	   100-continue 响应头除外。
	   服务器在读取Request.Body时自动发送的100-continue响应头。 */
	c.Writer.WriteHeader(code)
}

// 写入header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//向返回体中写入string
/* 如果WriteHeader还没有被调用，Write会在写入数据前调用
   WriteHeader(http.StatusOK)，然后再写入数据。如果Header
   不包含Content-Type行，Write会将Content-Type集添加到
   到将最初512字节的写入数据传递给
   DetectContentType。此外，如果所有写入的数据的总大小
   数据的总大小低于几KB，并且没有调用Flush，那么
   Content-Length头被自动添加。

   根据HTTP协议版本和客户端的不同，调用
   Write或WriteHeader可能会阻止未来对
   Request.Body。对于HTTP/1.x请求，处理程序应该在写入响应之前读取任何
   所需的请求正文数据，然后再写入响应。一旦
   头信息被刷新（由于明确的Flusher.Flush
   调用或写入足够的数据来触发刷新），请求正文
   可能无法使用。对于HTTP/2请求，Go HTTP服务器允许
   处理程序继续读取请求正文，同时
   写入响应。然而，这种行为可能不被所有的HTTP/2客户端支持
   所有的HTTP/2客户端都支持这种行为。处理程序应该先读后写，如果
   可能的话，以最大限度地提高兼容性" */
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 返回json格式
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//返回体写入
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

//返回HTMl
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
