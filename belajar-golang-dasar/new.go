package main

import "fmt"

// Address adalah struct yang merepresentasikan alamat dengan kota, provinsi, dan negara
type Address struct {
	City, Province, Country string
}

func main() {
	// Membuat pointer ke Address baru menggunakan literal struct
	var alamat1 *Address = &Address{}
	// Membuat pointer ke Address baru menggunakan fungsi new
	var alamat2 *Address = new(Address)
	// Membuat pointer alamat3 yang menunjuk ke alamat yang sama dengan alamat1
	var alamat3 *Address = alamat1

	// Mengubah nilai Country dari alamat yang ditunjuk oleh alamat2
	alamat2.Country = "Indonesia"

	// Mencetak nilai dari alamat1 (pointer)
	fmt.Println(alamat1) // Output: &{  Indonesia}
	// Mencetak nilai dari alamat2 (pointer)
	fmt.Println(alamat2) // Output: &{  Indonesia}
	// Mencetak nilai dari alamat3 (pointer)
	fmt.Println(alamat3) // Output: &{  Indonesia}
	// Mencetak nilai dari alamat1 (nilai yang ditunjuk oleh pointer)
	fmt.Println(*alamat1) // Output: {  Indonesia}
}