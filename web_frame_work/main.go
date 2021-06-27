package main

import (
	"fmt"
	"go-frameWork/web_frame_work/gee"
	"log"
	// question 1: 为什么要如此引入
	// 可能的原因,在我们的项目中之所以不如此引入,是因为使用了gitlab上的地址,所以不用填写相对位置.
	//todo 搞清楚golang的引入机制是如何运作的!!!!
)

func main() {

	r := gee.New()
	r.GET("/", func(c *gee.Context) {

		fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Req.URL)
	})

	r.GET("/hello", func(c *gee.Context) {
		for k, v := range c.Req.Header {
			fmt.Fprintf(c.Writer, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":9998")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
