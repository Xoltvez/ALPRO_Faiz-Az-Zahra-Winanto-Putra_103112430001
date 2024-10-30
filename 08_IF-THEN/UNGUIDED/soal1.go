package main

import "fmt"

func main() {
    //variabel untuk menyimpan data
    var jumlahOrang int

    //meminta inputan dari user untuk memasukan jumlahOrang
    fmt.Print("Masukkan jumlah orang yang akan melakukan touring: ")
    fmt.Scanln(&jumlahOrang)

	//percabangan jika jumlahorang bisa habis dibagi 2 tidak sama dengan 0 maka menampilkan "Jumlah motor yang diperlukan: + 1"
	if jumlahOrang%2 != 0 {
        jumlahOrang = (jumlahOrang / 2) + 1

    //
    } else {
        jumlahOrang = (jumlahOrang / 2)
    }
    


    //menampilkan output
    jumlahMotor := (jumlahOrang)
    fmt.Printf("Jumlah motor yang diperlukan: %d\n", jumlahMotor)
}