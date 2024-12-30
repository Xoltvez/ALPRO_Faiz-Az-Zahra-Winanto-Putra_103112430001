package main

import "fmt"

func main() {
    var p, r, hari int

   
    fmt.Print("Masukkan jumlah halaman novel (p): ")
    fmt.Scan(&p)
    fmt.Print("Masukkan jumlah halaman maksimal yang bisa dibaca per hari (r): ")
    fmt.Scan(&r)

   
    hari = p / r
    if p%r > 0 {
        hari = hari + 1
    }

    // Output hasil
    fmt.Printf("%d hari\n", hari)
}