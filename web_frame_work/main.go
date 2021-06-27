package main

import (
	"fmt"
	// question 1: 为什么要如此引入
	// 可能的原因,在我们的项目中之所以不如此引入,是因为使用了gitlab上的地址,所以不用填写相对位置.
	//todo 搞清楚golang的引入机制是如何运作的!!!!
	"go-frameWork/web_frame_work/gee"
	"net/http"
)

func main() {

	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
