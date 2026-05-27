# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Nicholas fachri kuntoro] - [109082500098]</p>

## Unguided 

### 1. 
Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya
Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program
rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut
membesar menggunakan algoritma selection sort.
Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat
Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m <
1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan
rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-
masing daerah.
#### soal1.go

```go

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		var m int
		fmt.Scan(&m)

		arr := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		for j := 0; j < m-1; j++ {

			minIdx := j

			for k := j + 1; k < m; k++ {
				if arr[k] < arr[minIdx] {
					minIdx = k
				}
			}

			arr[j], arr[minIdx] = arr[minIdx], arr[j]
		}

		for j := 0; j < m; j++ {
			fmt.Print(arr[j], " ")
		}
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul14/output/outputsoal1.png)
[Program ini digunakan untuk mengurutkan nomor rumah pada setiap daerah dari kecil ke besar menggunakan metode Selection Sort. Program membaca jumlah daerah, lalu setiap data rumah disimpan ke array dan diurutkan dengan mencari nilai terkecil kemudian ditukar ke posisi depan. Setelah terurut.]

### 2. 
Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu
diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena
nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah
program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.
Format Masukan masih persis sama seperti sebelumnya.
Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
#### soal2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul14/output/outputsoal2.png)
[Program ini memisahkan bilangan ganjil dan genap. Bilangan ganjil diurutkan naik (ascending), sedangkan bilangan genap diurutkan turun (descending) menggunakan Selection Sort. Setelah selesai diurutkan, program menampilkan bilangan ganjil terlebih dahulu lalu genap.]

### 3. 
Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan
tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan
sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu
problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba
untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?
"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data
genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua
data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke
bawah."
Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah
terbaca, jika data yang dibaca saat itu adalah 0.
Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.
Keluaran adalah median yang diminta, satu data per baris.
#### soal3.go

```go
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
	
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul14/output/outputsoal3.png)
[Program ini digunakan untuk mencari nilai median dari sekumpulan data. Data disimpan ke array lalu diurutkan menggunakan Selection Sort saat input 0 dimasukkan. Jika jumlah data ganjil, median diambil dari nilai tengah, sedangkan jika genap median dihitung dari rata-rata dua nilai tengah.]
