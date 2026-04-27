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