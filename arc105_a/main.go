package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	l := []int{a, b, c, d}
	sort.Ints(l)
	if l[3] == l[0]+l[1]+l[2] {
		fmt.Println("Yes")
		return
	}

	if l[0]+l[3] == l[1]+l[2] {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
