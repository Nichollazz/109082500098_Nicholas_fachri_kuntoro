# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
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

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(combination(b, d))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul3/output/outputsoal1.png)
[Program ini menghitung permutasi dan kombinasi menggunakan fungsi faktorial. Input berupa empat bilangan (a, b, c, d), lalu menghasilkan permutasi dan kombinasi dari a terhadap c, serta kombinasi dari b terhadap d.]

### 2. 
Diberikan tiga buah fungsi matematika yaitu f (x) = x
2
, g (x) = x − 2 dan h (x) = x +
Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul3/output/outputsoal2.png)
[Program ini menghitung komposisi dari tiga fungsi f(x)=x², g(x)=x−2, dan h(x)=x+1. Input tiga bilangan (a, b, c), kemudian menampilkan hasil fogoh(a), gohof(b), dan hofog(c).]

### 3. 
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func dalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 := dalamLingkaran(cx1, cy1, r1, x, y)
	in2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul3/output/outputsoal3.png)
[Program ini menentukan posisi sebuah titik terhadap dua lingkaran dengan menghitung jarak titik ke pusat lingkaran. Output menunjukkan apakah titik berada di dalam salah satu, keduanya, atau di luar kedua lingkaran.]
