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