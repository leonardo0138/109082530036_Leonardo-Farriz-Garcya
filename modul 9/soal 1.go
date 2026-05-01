package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}
type Lingkaran struct {
	pusat Titik
	r int
}

func jarak(p, q Titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func didalam(c Lingkaran, p Titik) bool {
	return jarak(p, c.pusat) <= float64(c.r)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	d1 := didalam(l1, t)
	d2 := didalam(l2, t)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 && !d2 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if !d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}