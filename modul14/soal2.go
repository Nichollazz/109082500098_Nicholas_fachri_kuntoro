package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {

			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}

		for j := 0; j < len(ganjil)-1; j++ {

			minIdx := j

			for k := j + 1; k < len(ganjil); k++ {
				if ganjil[k] < ganjil[minIdx] {
					minIdx = k
				}
			}

			ganjil[j], ganjil[minIdx] = ganjil[minIdx], ganjil[j]
		}

		for j := 0; j < len(genap)-1; j++ {

			maxIdx := j

			for k := j + 1; k < len(genap); k++ {
				if genap[k] > genap[maxIdx] {
					maxIdx = k
				}
			}

			genap[j], genap[maxIdx] = genap[maxIdx], genap[j]
		}

		for j := 0; j < len(ganjil); j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < len(genap); j++ {
			fmt.Print(genap[j], " ")
		}

		fmt.Println()
	}
}