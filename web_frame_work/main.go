package main

import (
	"fmt"
	"go-frameWork/web_frame_work/core"
	"log"
	// question 1: 为什么要如此引入
	// 可能的原因,在我们的项目中之所以不如此引入,是因为使用了gitlab上的地址,所以不用填写相对位置.
	//todo 搞清楚golang的引入机制是如何运作的!!!!
)

func main() {

	r := core.New()
	r.GET("/", func(c *core.Context) {

		fmt.Fprintf(c.Response, "URL.Path = %q\n", c.Request.URL)
	})

	r.GET("/hello", func(c *core.Context) {
		for k, v := range c.Request.Header {
			fmt.Fprintf(c.Response, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":9998")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
