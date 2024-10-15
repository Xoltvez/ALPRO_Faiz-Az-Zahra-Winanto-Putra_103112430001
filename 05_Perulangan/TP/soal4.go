package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Inisialisasi angka acak antara 1 hingga 100
	rand.Seed(time.Now().UnixNano())
	angkaTersembunyi := rand.Intn(100) + 1

	// Variabel untuk menyimpan input tebakan pengguna
	var tebakanPengguna int

	// Menjalankan permainan dengan maksimal 5 percobaan
	for kesempatan := 5; kesempatan > 0; kesempatan-- {
		fmt.Printf("Kamu punya %d kesempatan tersisa. Coba tebak angkanya: ", kesempatan)
		fmt.Scan(&tebakanPengguna)

		// Memeriksa apakah tebakan pengguna terlalu kecil, besar, atau benar
		if tebakanPengguna < angkaTersembunyi {
			fmt.Println("Tebakan kamu terlalu rendah!")
			fmt.Println("===========================")
		} else if tebakanPengguna > angkaTersembunyi {
			fmt.Println("Tebakan kamu terlalu tinggi!")
			fmt.Println("===========================")
		} else {
			fmt.Printf("Selamat! Tebakanmu tepat. Angka yang benar adalah %d.\n", angkaTersembunyi)
			break
		}

		// Jika tidak ada kesempatan tersisa, tampilkan pesan ini
		if kesempatan == 1 {
			fmt.Printf("Sayang sekali! Kesempatanmu habis. Angka yang benar adalah %d.\n", angkaTersembunyi)
		}
	}
}
