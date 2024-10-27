package main

import "fmt"

func main() {
    //variable untuk menyimpan data
    var semester, eprt int

    //meminta inputan user untuk memasukan semester dan eprt
    fmt.Print("Masukkan semester: ")
    fmt.Scan(&semester)

    fmt.Print("Masukkan eprt: ")
    fmt.Scan(&eprt)

    //aturan logika
    lulusCumlaude := semester <= 8 && eprt >= 500

    //percabangan untuk menampilkan output
    if lulusCumlaude {
        fmt.Printf("Mahasiswa lulus cumlaude dengan kuliah selama %d semester dan EPrT %d\n", semester, eprt)
    } else {
        fmt.Println("Mahasiswa tidak lulus cumlaude")
    }
}
