package main

import "fmt"

// Mendefinisikan tipe bentukan untuk karyawan
type Karyawan struct {
    Nama      string
    GajiPokok float64
    Tunjangan float64
    Potongan  float64
    TotalGaji float64
}

// Fungsi untuk menghitung total gaji karyawan
func hitungTotalGaji(karyawan *Karyawan) {
    karyawan.TotalGaji = karyawan.GajiPokok + karyawan.Tunjangan - karyawan.Potongan
}

func main() {
    //variable untuk menyimpan data
    var nama string
    var gajiPokok, tunjangan, potongan float64


    // Masukan: nama, gaji pokok, tunjangan, dan potongan
    fmt.Print("Masukkan nama karyawan: ")
    fmt.Scanln(&nama)
    fmt.Print("Masukkan gaji pokok: ")
    fmt.Scanln(&gajiPokok)
    fmt.Print("Masukkan tunjangan: ")
    fmt.Scanln(&tunjangan)
    fmt.Print("Masukkan potongan: ")
    fmt.Scanln(&potongan)

    // Membuat objek karyawan
    karyawan := Karyawan{
        Nama:      nama,
        GajiPokok: gajiPokok,
        Tunjangan: tunjangan,
        Potongan:  potongan,
    }

    // Menghitung total gaji
    hitungTotalGaji(&karyawan)

    // Menampilkan total gaji karyawan
    fmt.Printf("\nTotal gaji untuk %s adalah: %.2f\n", karyawan.Nama, karyawan.TotalGaji)
}
