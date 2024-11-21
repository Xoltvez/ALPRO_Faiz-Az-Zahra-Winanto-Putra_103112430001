package main

import ( "fmt")

func main() {

	//deklarasikan variable namaTanaman dengan tipe data string
	var namaTanaman string
	//meminta user untuk memasukan nama tanaman
	fmt.Scan(&namaTanaman)

	//gunakan switch case untuk menentukan kondisi nama tanaman
	switch namaTanaman {
	case "Nepenthes", "drosera" :
		fmt.Println("Termasuk tanaman karnivora ")
		fmt.Println("Asli Indonesia")

	case "venus", "monstera" :
		fmt.Println("Termasuk tanaman karnivora ")
		fmt.Println("Tidak Asli Indonesia")


	//menampilkan kondisi yg semua tidak terpenuhi
	default:
		fmt.Println("Tidak termasuk tanaman karnivora")
	}
}