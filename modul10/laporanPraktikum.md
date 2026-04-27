# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Nicholas fachri kuntoro] - [109082500098]</p>

## Unguided 

### 1. 
Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.
#### soal1.go

```go

package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var arr [1000]float64

    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    min := arr[0]
    max := arr[0]

    for i := 1; i < n; i++ {
        if arr[i] < min {
            min = arr[i]
        }
        if arr[i] > max {
            max = arr[i]
        }
    }

    fmt.Printf("Minimum: %.2f\n", min)
    fmt.Printf("Maximum: %.2f\n", max)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul10/output/outputsoal1.png)
[Program ini untuk mencari nilai minimum dan maksimum dari data berat dengan membandingkan tiap elemen.]

### 2. 
Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### soal2.go

```go
 package main
import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    var ikan [1000]float64

    for i := 0; i < x; i++ {
        fmt.Scan(&ikan[i])
    }

    var totalWadah []float64
    var sum float64 = 0
    count := 0

    for i := 0; i < x; i++ {
        sum += ikan[i]
        count++

        if count == y {
            totalWadah = append(totalWadah, sum)
            sum = 0
            count = 0
        }
    }

    if count > 0 {
        totalWadah = append(totalWadah, sum)
    }

    for _, v := range totalWadah {
        fmt.Printf("%.2f ", v)
    }

    fmt.Println()

    var total float64 = 0
    for _, v := range totalWadah {
        total += v
    }

    rata := total / float64(len(totalWadah))
    fmt.Printf("Rata-rata: %.2f\n", rata)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul10/output/outputsoal2.png)
[Program ini mengelompokkan data ke dalam wadah, menghitung total tiap wadah, lalu mencari rata-ratanya.]

### 3. 
Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go
package main
import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
    *bMin = arr[0]
    *bMax = arr[0]

    for i := 1; i < n; i++ {
        if arr[i] < *bMin {
            *bMin = arr[i]
        }
        if arr[i] > *bMax {
            *bMax = arr[i]
        }
    }
}

func rerata(arr arrBalita, n int) float64 {
    var sum float64 = 0
    for i := 0; i < n; i++ {
        sum += arr[i]
    }
    return sum / float64(n)
}

func main() {
    var n int
    fmt.Scan(&n)

    var data arrBalita

    for i := 0; i < n; i++ {
        fmt.Scan(&data[i])
    }

    var min, max float64

    hitungMinMax(data, n, &min, &max)
    avg := rerata(data, n)

    fmt.Printf("Berat minimum: %.2f kg\n", min)
    fmt.Printf("Berat maksimum: %.2f kg\n", max)
    fmt.Printf("Rata-rata: %.2f kg\n", avg)
}
	
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Nichollazz/109082500098_Nicholas_fachri_kuntoro/blob/main/modul10/output/outputsoal3.png)
[Program ini untuk menghitung minimum, maksimum, dan rata-rata menggunakan fungsi terpisah.]
