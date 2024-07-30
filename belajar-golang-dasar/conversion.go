package main

import "fmt"

func main() {
	// Mendeklarasikan variabel nilai32 dengan tipe int32 dan nilai 32768
	var nilai32 int32 = 32768
	
	// Mengkonversi nilai32 ke tipe int64
	var nilai64 int64 = int64(nilai32)
	
	// Mengkonversi nilai32 ke tipe int16 (akan terjadi overflow karena nilai32 lebih besar dari kapasitas int16)
	var nilai16 int16 = int16(nilai32)

	// Mencetak nilai dari variabel nilai32, nilai64, dan nilai16
	// nilai32: 32768
	// nilai64: 32768
	// nilai16: -32768 (overflow terjadi karena kapasitas int16 hanya sampai 32767)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Mendeklarasikan variabel name dengan nilai "xgus"
	var name = "xgus"
	
	// Mengambil karakter pertama dari string name dan mengkonversinya ke tipe uint8
	// name[0] mengambil karakter pertama 'x' yang memiliki nilai ASCII 120
	var e uint8 = name[0]
	
	// Mengkonversi nilai e ke string
	// string(e) mengkonversi nilai ASCII 120 kembali ke karakter 'x'
	var eString string = string(e)

	// Mencetak nilai dari variabel name, e, dan eString
	// name: "xgus"
	// e: 120 (nilai ASCII dari 'x')
	// eString: "x" (karakter yang sesuai dengan nilai ASCII 120)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}