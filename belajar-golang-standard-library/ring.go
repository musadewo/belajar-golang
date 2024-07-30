package main

import (
	"container/ring" // Mengimpor paket untuk struktur data ring
	"fmt"            // Mengimpor paket untuk format I/O
	"strconv"        // Mengimpor paket untuk konversi string dan angka
)

func main() {
	// Membuat ring baru dengan ukuran 5
	data := ring.New(5)

	// Mengisi ring secara manual
	data.Value = "Value 1" // Mengisi node pertama dengan "Value 1"
	data = data.Next()     // Pindah ke node berikutnya
	data.Value = "Value 2" // Mengisi node kedua dengan "Value 2"
	data = data.Next()     // Pindah ke node berikutnya
	data.Value = "Value 3" // Mengisi node ketiga dengan "Value 3"
	data = data.Next()     // Pindah ke node berikutnya
	data.Value = "Value 4" // Mengisi node keempat dengan "Value 4"
	data = data.Next()     // Pindah ke node berikutnya
	data.Value = "Value 5" // Mengisi node kelima dengan "Value 5"

	// Menampilkan semua nilai dalam ring
	data.Do(func(value any) {
		fmt.Println(value) // Mencetak setiap nilai dalam ring
	})

	// Mengisi ring secara otomatis menggunakan strconv.FormatInt
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value ke-" + strconv.FormatInt(int64(i), 10) // Mengisi node dengan nilai yang diformat menggunakan FormatInt
		data = data.Next() // Pindah ke node berikutnya
	}

	// Menampilkan semua nilai dalam ring
	data.Do(func(value any) {
		fmt.Println(value) // Mencetak setiap nilai dalam ring
	})

	// Mengisi ring secara otomatis menggunakan strconv.Itoa
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value ke-" + strconv.Itoa(i) // Mengisi node dengan nilai yang diformat menggunakan Itoa
		data = data.Next() // Pindah ke node berikutnya
	}

	// Menampilkan semua nilai dalam ring
	data.Do(func(value any) {
		fmt.Println(value) // Mencetak setiap nilai dalam ring
	})
}