package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var gcd int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		gcd = GCD(gcd, a)
	}

	fmt.Println(gcd)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
