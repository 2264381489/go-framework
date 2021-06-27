package utils

import "testing"

func TestConvertCommentToString(t *testing.T) {

	u := " \"Write writes the data to the connection as part of an HTTP reply.\n    //\n    // If WriteHeader has not yet been called, Write calls\n    // WriteHeader(http.StatusOK) before writing the data. If the Header\n    // does not contain a Content-Type line, Write adds a Content-Type set\n    // to the result of passing the initial 512 bytes of written data to\n    // DetectContentType. Additionally, if the total size of all written\n    // data is under a few KB and there are no Flush calls, the\n    // Content-Length header is added automatically.\n    //\n    // Depending on the HTTP protocol version and the client, calling\n    // Write or WriteHeader may prevent future reads on the\n    // Request.Body. For HTTP/1.x requests, handlers should read any\n    // needed request body data before writing the response. Once the\n    // headers have been flushed (due to either an explicit Flusher.Flush\n    // call or writing enough data to trigger a flush), the request body\n    // may be unavailable. For HTTP/2 requests, the Go HTTP server permits\n    // handlers to continue to read the request body while concurrently\n    // writing the response. However, such behavior may not be supported\n    // by all HTTP/2 clients. Handlers should read before writing if\n    // possible to maximize compatibility\""

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"第一次测试",
			args{s: u},
			" ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertCommentToString(tt.args.s); got != tt.want {
				t.Errorf("ConvertCommentToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
