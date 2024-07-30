package main

import "fmt"

func main() {

	// Contoh IF ELSE
	name := "John A" // Inisialisasi variabel name dengan nilai "John A"

	if name == "John" { // Jika nilai name adalah "John"
		fmt.Println("Hello John") // Cetak "Hello John"
	} else { // Jika nilai name bukan "John"
		fmt.Println("Bukan John") // Cetak "Bukan John"
	}

	// Contoh ELSE IF
	name = "John A" // Mengubah nilai variabel name menjadi "John A"

	if name == "John" { // Jika nilai name adalah "John"
		fmt.Println("Hello John") // Cetak "Hello John"
	} else if name == "John A" { // Jika nilai name adalah "John A"
		fmt.Println("Hello John A") // Cetak "Hello John A"
	} else { // Jika nilai name bukan "John" atau "John A"
		fmt.Println("Bukan John") // Cetak "Bukan John"
	}

	// Contoh Short Statement
	name = "Dewo" // Mengubah nilai variabel name menjadi "Dewo"

	// Menggunakan short statement untuk mendapatkan panjang string
	if length := len(name); length > 5 { // Jika panjang string name lebih dari 5 karakter
		fmt.Println("Nama terlalu panjang") // Cetak "Nama terlalu panjang"
	} else { // Jika panjang string name 5 karakter atau kurang
		fmt.Println("Nama pendek sudah benar") // Cetak "Nama pendek sudah benar"
	}
}