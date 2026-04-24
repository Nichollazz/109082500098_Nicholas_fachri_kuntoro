package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Semua: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(n)
	fmt.Println("Rata-rata:", mean)

	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}
	fmt.Println("Std Dev:", math.Sqrt(total/float64(n)))

	var x, freq int
	fmt.Scan(&x)
	for _, v := range arr {
		if v == x {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}