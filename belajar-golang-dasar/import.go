package main

import (
	"belajar-golang-dasar/helper" // Mengimpor package helper yang berisi fungsi dan variabel yang akan digunakan
	"fmt"                         // Mengimpor package fmt untuk operasi input/output format
)

func main() {
	// Memanggil fungsi SayHello dari package helper dengan argumen "Dewo"
	// Fungsi ini mengembalikan string sapaan
	result := helper.SayHello("Dewo")
	
	// Mencetak hasil dari fungsi SayHello ke konsol
	fmt.Println(result)

	// Mencetak versi dari package helper
	// Variabel version diambil dari package helper
	fmt.Println("Version:", helper.version)
	
	// Mencetak nama aplikasi dari package helper
	// Variabel Application diambil dari package helper
	fmt.Println("Application:", helper.Application)
}