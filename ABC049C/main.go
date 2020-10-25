package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	t := []string{"dream", "dreamer", "erase", "eraser"}
	for {
		r := hasSuffix(s, t)
		if hasSuffix(s, t) != "" {
			s = strings.TrimSuffix(s, r)
		} else {
			break
		}
	}

	if s == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func hasSuffix(s string, l []string) (r string) {
	for _, v := range l {
		if strings.HasSuffix(s, v) {
			return v
		}
	}
	return ""
}
