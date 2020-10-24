package main

import "fmt"

func main() {
	var i, j int
	_, _ = fmt.Scan(&i, &j)
	if i*j%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
