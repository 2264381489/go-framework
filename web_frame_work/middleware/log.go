package middleware

import (
	"fmt"
	"go-frameWork/web_frame_work/core"
	"log"
	"time"
)

func Logger() core.HandlerFuncution {
	return func(c *core.Context) (interface{}, error) {
		r := c.Request
		// Start timer
		t := time.Now()
		fmt.Println()
		fmt.Printf("req method %s url: %s ", r.Method, r.URL.Path)
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
		return nil, nil
	}
}
