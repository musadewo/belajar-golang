package main

import "fmt"

// sumAll adalah fungsi variadic yang menerima sejumlah argumen integer
func sumAll(numbers ...int) int {
	total := 0
	// Loop melalui setiap angka dalam 'numbers' dan tambahkan ke 'total'
	for _, number := range numbers {
		total += number // Tambahkan nilai 'number' ke 'total'
	}
	// Kembalikan total dari semua angka
	return total // Mengembalikan hasil penjumlahan semua angka
}

func main() {
	// Panggil sumAll dengan empat argumen dan cetak hasilnya
	result := sumAll(10, 10, 10, 10)
	fmt.Println(result) // Output: 40
	
	// Panggil sumAll dengan tujuh argumen dan cetak hasilnya
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10)) // Output: 70

	// Buat slice dari integer dan panggil sumAll dengan menggunakan slice tersebut
	numbers := []int{10, 10, 10, 10, 10, 10, 10}
	result = sumAll(numbers...) // Menggunakan slice sebagai argumen variadic
	// Cetak hasilnya
	fmt.Println(result) // Output: 70
}