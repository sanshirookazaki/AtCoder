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
		countable := true
		if contain(v, r) {
			break
		}
		for j := i + 1; j < n; j++ {
			if v == list[j] {
				break
			}

			if list[j]%v == 0 {
				r = append(r, list[j])
				countable = false
				continue
			}
		}
		if countable == true {
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
