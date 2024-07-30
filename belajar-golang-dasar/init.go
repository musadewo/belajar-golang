package main

import (
	// Mengimpor package database untuk menggunakan fungsi-fungsi yang ada di dalamnya
	"belajar-golang-dasar/database"
	// Mengimpor package internal tanpa menggunakan secara langsung (Blank Identifier)
	// Ini biasanya dilakukan untuk memastikan package diinisialisasi walaupun tidak digunakan secara langsung
	_ "belajar-golang-dasar/internal"
	// Mengimpor package fmt untuk format I/O seperti mencetak ke layar
	"fmt"
)

func main() {
	// Memanggil fungsi GetConnection dari package database
	// Fungsi ini mungkin mengembalikan string yang berisi informasi koneksi ke database
	// Hasil dari fungsi ini kemudian dicetak ke layar menggunakan fmt.Println
	fmt.Println(database.GetConnection())
}