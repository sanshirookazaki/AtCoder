package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	c := strings.Count(s, "1")
	fmt.Println(c)
}
