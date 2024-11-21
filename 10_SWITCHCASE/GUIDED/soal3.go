package main  
import "fmt"  
func main() {
    //deklarasikan variable kendaraan,durasi, tarif dengan tipe data string dan integer  
     var kendaraan string 
     var durasi int 
     var tarif int 

     fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ") 
     //meminta inputan dari user untuk memasukan kendaraan dan durasinya
     fmt.Scan(&kendaraan) 
     fmt.Print("Masukkan durasi parkir (dalam jam): ") 
     fmt.Scan(&durasi) 

     //gunakan switch case untuk menentukan tarif parkir berdasarkan jenis kendaraan dan durasinya
     switch { 
     case kendaraan == "Motor" && durasi >= 1 && durasi <= 2: 
         tarif = 7000 
     case kendaraan == "Motor" && durasi > 2: 
         tarif = 9000 
     case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2: 
         tarif = 15000 
     case kendaraan == "Mobil" && durasi > 2: 
         tarif = 20000 
     case kendaraan == "Truk" && durasi >= 1 && durasi <= 2: 
         tarif = 25000 
     case kendaraan == "Truk" && durasi > 2: 
         tarif = 35000 

     //menampilkan kondisi yg semua tidak terpenuhi
     default: 
         fmt.Println("Jenis kendaraan atau durasi parkir tidak valid") 
     } 

     //menampilkan hasil
     fmt.Printf("Tarif Parkir: Rp %d\n", tarif) 
}