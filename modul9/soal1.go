package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	center Titik
	r      int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func diDalam(c Lingkaran, p Titik) bool {
	return jarak(c.center, p) <= float64(c.r)
}

func main() {
	var c1, c2 Lingkaran
	var p Titik

	fmt.Scan(&c1.center.x, &c1.center.y, &c1.r)
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	in1 := diDalam(c1, p)
	in2 := diDalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}