package main

import "fmt"

func main() {

	// Deklarasi variabel
	var (
		password, inputPassword string
		maxLogin                int
	)

	// Isi variabel untuk password dan jumlah maximum memasukan password
	password = "frizy7"
	maxLogin = 3

	// Perulangan untuk memasukan password, yaitu maksimal 3 kali
	for i := 0; i < maxLogin; i++ {

		// Inputan untuk memasukan passoword
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&inputPassword)

		// Pecabangan untuk memeriksa apakah password yang dimasukan benar atau salah.
		if inputPassword == password {
			fmt.Print("Berhasil Login")
			break
		} else {
			fmt.Println("Password Salah")
		}

		// Percabangan lagi untuk menampilkan pesan jika upaya login melebihi batas
		if i == maxLogin-1 {
			print("Login Ditolak")
		}

	}
}