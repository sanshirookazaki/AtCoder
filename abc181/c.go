package main

import (
	"fmt"
)

type P struct {
	x float64
	y float64
}

func main() {
	var n int
	fmt.Scan(&n)
	var Ps []*P
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		Ps = append(Ps, &P{float64(a), float64(b)})
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				var dx, dy []float64
				if Ps[i].x <= Ps[j].x && Ps[i].x <= Ps[k].x {
					dx = []float64{Ps[i].x, Ps[j].x, Ps[k].x}
					dy = []float64{Ps[i].y, Ps[j].y, Ps[k].y}
				} else if Ps[j].x <= Ps[i].x && Ps[j].x <= Ps[k].x {
					dx = []float64{Ps[j].x, Ps[i].x, Ps[k].x}
					dy = []float64{Ps[j].y, Ps[i].y, Ps[k].y}
				} else if Ps[k].x <= Ps[i].x && Ps[k].x <= Ps[j].x {
					dx = []float64{Ps[k].x, Ps[i].x, Ps[j].x}
					dy = []float64{Ps[k].y, Ps[i].y, Ps[j].y}
				} else {
					dx = []float64{Ps[k].x, Ps[i].x, Ps[j].x}
					dy = []float64{Ps[k].y, Ps[i].y, Ps[j].y}
				}
				dx1 := (dx[1] - dx[0])
				dy1 := (dy[1] - dy[0])
				dx2 := (dx[2] - dx[0])
				dy2 := (dy[2] - dy[0])
				if dx1*dy2 == dx2*dy1 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")

}
