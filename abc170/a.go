package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	list := []int{a, b, c, d, e}
	for i, v := range list {
		if v == 0 && i == 0 {
			fmt.Println(list[i+1] - 1)
		} else if v == 0 {
			fmt.Println(list[i-1] + 1)
		}
	}
}
