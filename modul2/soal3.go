package main

import "fmt"

func main() {
	var gram int
	var kg, sisa int
	var biayaKg, biayaSisa int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&gram)

	kg = gram / 1000
	sisa = gram % 1000

	biayaKg = kg * 10000

	if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total := biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}
