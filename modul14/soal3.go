package main

import "fmt"

func main() {

	var data []int
	var x int

	for {

		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x != 0 {

			data = append(data, x)

		} else {

			for i := 0; i < len(data)-1; i++ {

				minIdx := i

				for j := i + 1; j < len(data); j++ {

					if data[j] < data[minIdx] {
						minIdx = j
					}
				}

				data[i], data[minIdx] = data[minIdx], data[i]
			}

			n := len(data)

			if n%2 == 1 {

				fmt.Println(data[n/2])

			} else {

				median := (data[n/2] + data[n/2-1]) / 2
				fmt.Println(median)
			}
		}
	}
}