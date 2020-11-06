package main

import (
	"fmt"
)

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	var list []int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		list = append(list, a)
	}

	var a int
	for {
		if !containe(x-a, list) {
			fmt.Println(x - a)
			return
		}

		if !containe(x+a, list) {
			fmt.Println(x + a)
			return
		}

		a++
	}

}

func containe(a int, b []int) bool {
	for _, v := range b {
		if v == a {
			return true
		}
	}
	return false
}
