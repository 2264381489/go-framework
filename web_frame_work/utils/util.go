package utils

import (
	"fmt"
	"strings"
)

func ConvertCommentToString(s string) string {
	if strings.Contains(s, "//") {
		s = strings.ReplaceAll(s, "//", "")
	}
	if strings.Contains(s, "/*") {
		s = strings.ReplaceAll(s, "/*", "")
		s = strings.ReplaceAll(s, "*/", "")
	}
	fmt.Println(s)
	return s
}
