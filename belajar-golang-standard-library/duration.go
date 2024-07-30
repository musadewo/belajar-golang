package main

import (
	"fmt"  // Package fmt menyediakan fungsi format I/O seperti Println
	"time" // Package time menyediakan fungsi dan tipe untuk mengelola waktu
)

func main() {
	// Mendeklarasikan variabel duration1 dengan tipe time.Duration dan menginisialisasinya dengan 100 detik
	var duration1 time.Duration = 100 * time.Second
	// Mencetak nilai dari duration1
	fmt.Println(duration1)

	// Mendeklarasikan variabel duration2 dengan tipe time.Duration dan menginisialisasinya dengan 10 milidetik
	duration2 := 10 * time.Millisecond
	// Mencetak nilai dari duration2
	fmt.Println(duration2)

	// Menghitung selisih antara duration1 dan duration2, hasilnya disimpan dalam variabel duration3
	duration3 := duration1 - duration2
	// Mencetak nilai dari duration3
	fmt.Println(duration3)
}