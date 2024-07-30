package main

import (
	"fmt"  // Mengimpor paket fmt untuk format I/O
	"math" // Mengimpor paket math untuk fungsi matematika
)

func main() {
	// Menggunakan fungsi Round dari paket math untuk membulatkan angka ke nilai terdekat
	fmt.Println(math.Round(1.60)) // Output: 2

	// Menggunakan fungsi Floor dari paket math untuk membulatkan angka ke bawah
	fmt.Println(math.Floor(1.60)) // Output: 1

	// Menggunakan fungsi Ceil dari paket math untuk membulatkan angka ke atas
	fmt.Println(math.Ceil(1.40)) // Output: 2

	// Menggunakan fungsi Max dari paket math untuk mendapatkan nilai maksimum dari dua angka
	fmt.Println(math.Max(10, 11)) // Output: 11

	// Menggunakan fungsi Min dari paket math untuk mendapatkan nilai minimum dari dua angka
	fmt.Println(math.Min(10, 11)) // Output: 10
}