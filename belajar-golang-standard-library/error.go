package main

import (
	"errors"
	"fmt"
)

// Mendefinisikan variabel error global
var (
	ValidateError = errors.New("Validate Error") // Error untuk validasi, digunakan ketika input tidak valid
	NotFoundError = errors.New("Not Found Error") // Error untuk data tidak ditemukan, digunakan ketika data tidak ditemukan
)

// Fungsi untuk mendapatkan data berdasarkan ID
func GetById(id string) error {
	// Jika ID kosong, kembalikan error validasi
	if id == "" {
		return ValidateError // Mengembalikan error validasi jika ID kosong
	}
	// Jika ID adalah "Dewo", kembalikan error data tidak ditemukan
	if id == "Dewo" {
		return NotFoundError // Mengembalikan error data tidak ditemukan jika ID adalah "Dewo"
	}
	// Jika tidak ada error, kembalikan nil
	return nil // Mengembalikan nil jika tidak ada error
}

func main() {
	// Memanggil fungsi GetById dengan ID kosong
	err := GetById("") // Memanggil fungsi dengan ID kosong untuk memicu ValidateError
	// Jika ada error
	if err != nil {
		// Cek apakah error adalah ValidateError
		if errors.Is(err, ValidateError) {
			fmt.Println("Validate Error", err) // Menampilkan pesan error validasi
		// Cek apakah error adalah NotFoundError
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error", err) // Menampilkan pesan error data tidak ditemukan
		// Jika error tidak dikenal
		} else {
			fmt.Println("Unknown Error", err) // Menampilkan pesan error tidak dikenal
		}
	}
}