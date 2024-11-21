package main


import "fmt"


func main(){
	//deklarasi variable kadarPh dan keterangan dengan tipe data float64 dan string
	var kadarPh float64
	var keterangan string

	//meminta inputan dari user untuk memasukan kadar ph
	fmt.Print("Masukan kadar ph: ")
	fmt.Scan(&kadarPh)


	//gunakan switch case untuk menentukan kondisi kadar ph yg terpenuhi
	switch {
	case kadarPh < 0 || kadarPh > 14:
		keterangan = "Nilai pH tidak valid. Nilai pH harus antara 0 dan 14"
	case kadarPh >= 6.5 && kadarPh <= 8.6:
		keterangan = "Air layak minum"
	//menampilkan kondisi yg semua tidak terpenuhi	
	default:
		keterangan = "Air tidak layak minum"
	}
	//menampilkan hasil
	fmt.Println(keterangan)
}