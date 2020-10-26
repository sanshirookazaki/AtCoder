package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; 40 > i; i++ {
		for j := 1; 30 > j; j++ {
			a := int(math.Pow(3, float64(i)))
			b := int(math.Pow(5, float64(j)))
			if a+b == n {
				fmt.Println(i, j)
				return
			}
		}
	}

	fmt.Println(-1)

}
