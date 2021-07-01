package utils

import "testing"

func TestConvertCommentToString(t *testing.T) {

	u :=
		"// Copyright 2014 The Go Authors. All rights reserved.\n// Use of this source code is governed by a BSD-style\n// license that can be found in the LICENSE file.\n\n// Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.\n//\n// Incoming requests to a server should create a Context, and outgoing\n// calls to servers should accept a Context. The chain of function\n// calls between them must propagate the Context, optionally replacing\n// it with a derived Context created using WithCancel, WithDeadline,\n// WithTimeout, or WithValue. When a Context is canceled, all\n// Contexts derived from it are also canceled.\n//\n// The WithCancel, WithDeadline, and WithTimeout functions take a\n// Context (the parent) and return a derived Context (the child) and a\n// CancelFunc. Calling the CancelFunc cancels the child and its\n// children, removes the parent's reference to the child, and stops\n// any associated timers. Failing to call the CancelFunc leaks the\n// child and its children until the parent is canceled or the timer\n// fires. The go vet tool checks that CancelFuncs are used on all\n// control-flow paths.\n//\n// Programs that use Contexts should follow these rules to keep interfaces\n// consistent across packages and enable static analysis tools to check context\n// propagation:\n//\n// Do not store Contexts inside a struct type; instead, pass a Context\n// explicitly to each function that needs it. The Context should be the first\n// parameter, typically named ctx:\n//\n// \tfunc DoSomething(ctx context.Context, arg Arg) error {\n// \t\t// ... use ctx ...\n// \t}\n//\n// Do not pass a nil Context, even if a function permits it. Pass context.TODO\n// if you are unsure about which Context to use.\n//\n// Use context Values only for request-scoped data that transits processes and\n// APIs, not for passing optional parameters to functions.\n//\n// The same Context may be passed to functions running in different goroutines;\n// Contexts are safe for simultaneous use by multiple goroutines.\n//\n// See https://blog.golang.org/context for example code for a server that uses\n// Contexts."
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
