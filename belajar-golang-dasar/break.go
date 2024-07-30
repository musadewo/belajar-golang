package main

import "fmt"

func main() {
	// Perulangan dari 0 sampai 9
	for i := 0; i < 10; i++ {
		// Jika i sama dengan 5, keluar dari perulangan
		if i == 5 {
			break // Menghentikan perulangan ketika i sama dengan 5
		}
		// Cetak nilai i
		fmt.Println("Perulangan ke-", i) // Menampilkan nilai i saat ini
	}
	// Cetak pesan selesai setelah perulangan pertama
	fmt.Println("Selesai") // Menampilkan pesan bahwa perulangan pertama selesai

	// Perulangan dari 0 sampai 9
	for i := 0; i < 10; i++ {
		// Jika i adalah bilangan genap, lanjutkan ke iterasi berikutnya
		if i%2 == 0 {
			continue // Melanjutkan ke iterasi berikutnya jika i adalah bilangan genap
		}
		// Cetak nilai i
		fmt.Println("Perulangan ke-", i) // Menampilkan nilai i saat ini
	}
	// Cetak pesan selesai setelah perulangan kedua
	fmt.Println("Selesai") // Menampilkan pesan bahwa perulangan kedua selesai
}