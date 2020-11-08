package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	var list []int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		list = append(list, a)
	}

	sort.Ints(list)
	var c int
	var r []int

	for i, v := range list {
		if contain(v, r) {
			continue
		}

		f := true
		for j := i + 1; j < n; j++ {
			if list[j] == v {
				r = append(r, list[j])
				f = false
			}

			if list[j]%v == 0 {
				r = append(r, list[j])
			}
		}
		if f == true {
			c++
		}
	}
	fmt.Println(c)
}

func contain(i int, l []int) bool {
	for _, v := range l {
		if i == v {
			return true
		}
	}
	return false
}
