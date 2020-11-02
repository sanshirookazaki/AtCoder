package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var sum int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		sum += Sum(a, b)
	}
	fmt.Println(sum)
}

func Sum(a, b int) (r int) {
	r += (b - a + 1) * (a + b) / 2
	return r
}
