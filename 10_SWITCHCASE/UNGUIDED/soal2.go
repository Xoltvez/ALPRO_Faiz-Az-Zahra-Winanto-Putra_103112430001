package main

import "fmt"

func main(){

	//dekralasikan variable jenis kendaraan, durasi parkir, dan tarif perjam dengan tipe data string dan integer
	var jenisKendaraan string
	var durasiParkir int 
	var tarifPerjam int

	//meminta inputan dari user untuk memasukan jenis kendaraan dan durasi parkir
	fmt.Print("Masukan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&jenisKendaraan)

	fmt.Print("Masukan durasi parkir (dalam jam): ")	
	fmt.Scan(&durasiParkir)

	//gunakan switch case untuk menentukan tarif parkir berdasarkan jenis kendaraan yg terpenuhi
	switch jenisKendaraan {
	case "Motor":
		tarifPerjam = 2000
	case "Mobil":
		tarifPerjam = 5000
	case "Truk":
		tarifPerjam = 8000

	 //menampilkan kondisi yg semua tidak terpenuhi
	default:
		fmt.Println("Jenis kendaraan tidak valid")
	}	

	//rumus untuk total biaya = tarif perjam * durasi parkir
	totalBiaya := tarifPerjam * durasiParkir

	//menampilkan total biaya
	fmt.Println("Total biaya parkir:", totalBiaya)


}

