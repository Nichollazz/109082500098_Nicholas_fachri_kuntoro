# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Nicholas fachri kuntoro] - [109082500098]</p>

## Unguided 

### 1. 
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var satu, dua, tiga string
	var temp string

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal =", satu, dua, tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir =", satu, dua, tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul2/output/outputsoal1.png)
[Program tersebut menerima tiga input string dari pengguna, kemudian menampilkan urutan awalnya. Setelah itu program menukar posisi ketiga string tersebut sehingga urutannya berubah menggunakan variabel sementara.]

### 2. 
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
#### soal2.go

```go
package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if !(w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul2/output/outputsoal2.png)
[Program tersebut melakukan lima kali percobaan dengan memasukkan empat warna setiap percobaan. Program memeriksa apakah urutan warna adalah merah, kuning, hijau, ungu. Jika semua percobaan benar maka hasilnya true, jika ada yang salah maka false.]

### 3. 
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal3.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul2/output/outputsoal3.png)
[Program tersebut menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Berat diubah menjadi kilogram dan sisa gram, kemudian biaya dihitung dari tarif per kilogram dan tambahan biaya dari sisa gram, lalu ditampilkan total biayanya.]
