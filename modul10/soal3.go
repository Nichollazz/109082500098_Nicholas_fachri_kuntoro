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