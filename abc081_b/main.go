package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)
	var l []int

	for i := 0; i < n; i++ {
		var i int
		fmt.Scan(&i)
		l = append(l, i)
	}

L:
	for {
		for i, v := range l {
			if v%2 == 1 {
				break L
			}

			l[i] = v / 2
		}

		c++
	}

	fmt.Println(c)

}
