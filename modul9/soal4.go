package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Scan(&input)

	*n = len(input)
	for i := 0; i < *n; i++ {
		t[i] = rune(input[i])
	}
}

func balikArray(t tabel, n int) tabel {
	var hasil tabel
	for i := 0; i < n; i++ {
		hasil[i] = t[n-1-i]
	}
	return hasil
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func main() {
	var tab, rev tabel
	var n int

	isiArray(&tab, &n)

	rev = balikArray(tab, n)

	fmt.Print("Reverse teks: ")
	cetakArray(rev, n)

	fmt.Print("Palindrome: ")
	fmt.Println(palindrom(tab, n))
}