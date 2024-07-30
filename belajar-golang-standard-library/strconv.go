package main

import (
	"fmt"     // Paket untuk format I/O
	"strconv" // Paket untuk konversi string ke tipe data lain dan sebaliknya
)

func main() {
	// Mencoba mengkonversi string "SALAH" ke boolean
	result, err := strconv.ParseBool("SALAH")
	if err != nil {
		// Jika terjadi error, cetak pesan error
		fmt.Println("Error: ", err)
	} else {
		// Jika tidak ada error, cetak hasil konversi
		fmt.Println(result)
	}

	// Mencoba mengkonversi string "1000" ke integer
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		// Jika terjadi error, cetak pesan error
		fmt.Println("Error: ", err)
	} else {
		// Jika tidak ada error, cetak hasil konversi
		fmt.Println(resultInt)
	}

	// Mengkonversi integer 999 ke string biner
	binary := strconv.FormatInt(999, 2)
	// Cetak hasil konversi ke biner
	fmt.Println(binary)

	// Mengkonversi integer 999 ke string
	var stringInt string = strconv.Itoa(999)
	// Cetak hasil konversi ke string
	fmt.Println(stringInt)
}