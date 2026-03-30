package main

import "fmt"


func hitungSkor(waktu [8]int, soal *int, total *int) {
	*soal = 0
	*total = 0

	for i := 0; i < 8; i++ {
		if waktu[i] < 301 {
			*soal++
			*total += waktu[i]
		}
	}
}

func main() {
	var nama string

	maxSoal := -1
	minWaktu := 999999
	pemenang := ""

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var waktu [8]int
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, total int
		hitungSkor(waktu, &soal, &total)

		if soal > maxSoal || (soal == maxSoal && total < minWaktu) {
			maxSoal = soal
			minWaktu = total
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}