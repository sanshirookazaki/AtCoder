package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a == b {
			c++
			if c == 3 {
				fmt.Println("Yes")
				return
			}
		} else {
			c = 0
		}
	}
	fmt.Println("No")
}
