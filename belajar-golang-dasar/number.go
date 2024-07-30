package main

import (
	"fmt"  // Mengimpor paket fmt untuk fungsi format I/O seperti Println
	"math" // Mengimpor paket math untuk fungsi matematika seperti Pow
)

func main() {
	// Mencetak string "Satu = " diikuti dengan angka 1
	fmt.Println("Satu = ", 1)
	
	// Mencetak string "Dua = " diikuti dengan angka 2
	fmt.Println("Dua = ", 2)
	
	// Mencetak string "Tiga koma lima = " diikuti dengan angka desimal 3.5
	fmt.Println("Tiga koma lima = ", 3.5)
	
	// Menghitung 16 pangkat 16 menggunakan fungsi Pow dari paket math
	result := math.Pow(16, 16)
	
	// Mencetak string "16 pangkat 16 = " diikuti dengan hasil perhitungan 16 pangkat 16
	fmt.Println("16 pangkat 16 = ", result)
}