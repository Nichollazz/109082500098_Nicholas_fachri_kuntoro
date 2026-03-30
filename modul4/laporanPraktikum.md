# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Nicholas fachri kuntoro] - [109082500098]</p>

## Unguided 

### 1. 
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan
menggunakan persamaan berikut!
#### soal1.go

```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutation(a, c, &p1)
	combination(a, c, &c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}



```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul4/output/outputsoal1.png)
[Program ini menghitung permutasi dan kombinasi dari input menggunakan prosedur faktorial, lalu menampilkan hasil untuk pasangan (a, c) dan (b, d).]

### 2. 
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah
8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.
Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan
dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang
diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil
diselesaikan.
#### soal2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul4/output/outputsoal2.png)
[Program ini menentukan pemenang lomba dengan menghitung jumlah soal yang diselesaikan dan total waktunya, pemenang adalah yang menyelesaikan soal paling banyak, atau jika sama, yang waktunya paling kecil.]
