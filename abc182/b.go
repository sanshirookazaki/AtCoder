package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	listMap := map[int]int{}
	for i := 0; i < n; i++ {
		var a int
		var list []int
		fmt.Scan(&a)
		d := divlist(a)
		list = append(list, d...)
		for _, v := range list {
			listMap[v] += 1
		}
	}
	fmt.Println(max(listMap))
}

func divlist(n int) []int {
	var list []int
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			list = append(list, i)
		}
	}
	return list
}

func max(m map[int]int) int {
	var value, result int
	for i, v := range m {
		if value <= v {
			result = i
			value = v
		}
	}
	return result
}
