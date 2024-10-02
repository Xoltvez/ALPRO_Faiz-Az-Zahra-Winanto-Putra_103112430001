package main

import "fmt"


//kode untuk pengkodisian jika tahun kabisat habis dibagi 4 adalah benar dan tahun kabisat != (tidak habis dibabgi 100) adalah salah
func yearlaps(kabisat int) bool {
	if kabisat % 4 == 0 && kabisat % 100 != 0 {
		return true
	
	}
	return false
		
	}


func main(){

	var(
		tahun int
	)

	//kode untuk meminta inputan dari user
	fmt.Println("Masukan tahun :")
	fmt.Scan(&tahun)

	//untuk menampilkan hasil operasi tahun kabisat atau bukan
	fmt.Println(yearlaps(tahun))


}