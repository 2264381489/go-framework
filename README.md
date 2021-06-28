# go-framework
go的自编框架
之所以要有框架，是因为go语言提供的包仅仅提供了对访问的url进行处理的功能。

看到如下代码你可能有些头大。

但是实际上代码中只执行了一个简单的逻辑和java Spring框架中@requestMapping的基础作用一致，匹配访问的URL，并选择一个合适的处理器。

The documentation for ServeMux explains how patterns are matched处理的具体逻辑在这里。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/hello",helloHandler)
	//打了：9999 会给一个缺省的localhost 不打没有
	log.Fatal(http.ListenAndServe(":9999",nil))
}

// The documentation for ServeMux explains how patterns are matched. 有空看一下handler的实现
// http.ResponseWriter 的作用是向返回体写入内容
func handler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
```

这里要熟知几个概念。

`http.ListenAndServe(":9999",nil)`

go语言依靠这个方法开启一个服务器进程。

当执行这个语句的时候，这个go程序就变成了一个不会自动退出（手动或者panic才可以退出）可以接收HTTP请求的服务器程序。

 `http.ResponseWriter`

http.ResponseWriter 的作用是向返回体写入内容。

```go
func handler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
```

以这个方法为例，方法执行时将访问路径直接写入到返回体中。

访问localhost：9999/ 返回体中返回的是 URL.Path = / 

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4b247510-f160-4cc4-aaa5-7a38dc973c2a/Untitled.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4b247510-f160-4cc4-aaa5-7a38dc973c2a/Untitled.png)
展示成功结果

## 拦截HTTP请求

当我们了解了http包的运行机制之后，应该清楚，实际上框架就是对语言原本功能的一个封装。

GO也提供了这个功能，在`server.ListenAndServe()` 提供了这个功能。

接受一个handler去处理，handler是接口和我们上文用的handler方法实现的是同一个接口。

```go
// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case the DefaultServeMux is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

当我们将 **实现了一个handler接口的结构体（实现了ServeHTTP方法） 传入**ListenAndServe的第二个参数之后。

```go
func main() {
// handler ， helloHandler 代码已经省略，在前文中有方法实现。
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	mvc:=new(MVC)
	//打了：9999 会给一个缺省的localhost 不打没有
	log.Fatal(http.ListenAndServe(":9999", mvc))
}

// 实现一个ServerHTTp接口实现对于URL访问的拦截

type MVC struct{}

func (mvc *MVC) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "MVC URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Println(w, "Support by MVC")
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Println(w, "MVC start")
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}

}
```

我们在访问“/" 和 ”/hello"的时候返回体的内容为

```go
GET /
MVC URL.Path = "/"
GET /hello
Support by MVC
Header["Connection"] = ["keep-alive"]
Support by MVC
Header["User-Agent"] = ["PostmanRuntime/7.28.0"]
Support by MVC
Header["Accept"] = ["*/*"]
Support by MVC
Header["Cache-Control"] = ["no-cache"]
Support by MVC
Header["Postman-Token"] = ["20b34de4-9ff6-440d-a24d-b204c0c7f3d2"]
Support by MVC
Header["Accept-Encoding"] = ["gzip, deflate, br"]
```

但是注意我们的main函数，虽然依然有`http.HandleFunc("/", handler)` 但是明显没有执行，执行的是MVC结构体实现的ServeHTTP函数。

这就是拦截功能，用我们自己实现的handler去匹配url和寻找处理函数。

而不用GOhttp包自带的功能。

```go
func main() {
// handler ， helloHandler 代码已经省略，在前文中有方法实现。
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	mvc:=new(MVC)
	//打了：9999 会给一个缺省的localhost 不打没有
	log.Fatal(http.ListenAndServe(":9999", mvc))
}
```

> 我们定义了一个空的结构体Engine，实现了方法ServeHTTP。这个方法有2个参数，第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
在 main 函数中，我们给 ListenAndServe 方法的第二个参数传入了刚才创建的engine实例。至此，我们走出了实现Web框架的第一步，即，将所有的HTTP请求转向了我们自己的处理逻辑。还记得吗，在实现Engine之前，我们调用 http.HandleFunc 实现了路由和Handler的映射，也就是只能针对具体的路由写处理逻辑。比如/hello。但是在实现Engine之后，我们拦截了所有的HTTP请求，拥有了统一的控制入口。在这里我们可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，例如日志、异常处理等。
代码的运行结果与之前的是一致的。

## 框架实现

GO语言中函数也可作为一种type被使用

```go
package gee

import (
	"fmt"
	"net/http"
)
// 本函数将被用作是否为合理的处理函数
// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
// 加入路由，将方法以method和url的组合为key，执行函数为value生成一个map
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}
//添加GET方法
// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
//添加post方法
// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}
//服务启动
// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//实现的处理器。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
```