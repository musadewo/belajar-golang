package main

import "fmt"

// Fungsi Ups mengembalikan nilai tipe 'any' (interface kosong)
func Ups() any {
	return 1       // Mengembalikan integer 1
	// return true    // Tidak akan pernah dieksekusi karena return sebelumnya
	// return "Ups"   // Tidak akan pernah dieksekusi karena return sebelumnya
}

func main() {
	// Memanggil fungsi Ups dan menyimpan hasilnya dalam variabel kosong
	kosong := Ups() 
	// Mencetak nilai dari variabel kosong
	fmt.Println(kosong) 

	// Memanggil fungsi Ups dan menyimpan hasilnya dalam variabel kosong dengan tipe 'any'
	var kosong any = Ups() 
	// Mencetak nilai dari variabel kosong
	fmt.Println(kosong) 
}