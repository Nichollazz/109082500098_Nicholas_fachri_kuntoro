package main

import "fmt"

func main() {
	var klubA, klubB string

	fmt.Scan(&klubA, &klubB)

	var skorA, skorB int
	hasil := []string{}

	for {
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	for i, h := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, h)
	}
	fmt.Println("Pertandingan selesai")
}