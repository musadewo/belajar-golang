package main

import "fmt"

// Definisikan struct Address dengan field City, Province, dan Country
type Address struct {
	City, Province, Country string
}

func main() {
	// Buat instance dari Address
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	
	// Salin address1 ke address2 (shallow copy)
	address2 := address1

	// Ubah field City dari address2
	address2.City = "Bandung"

	// Cetak address1 dan address2
	fmt.Println(address1) // Output: {Subang Jawa Barat Indonesia}
	fmt.Println(address2) // Output: {Bandung Jawa Barat Indonesia}
	fmt.Println(address2.City) // Output: Bandung

	// Buat pointer ke address1
	address3 := &address1
	
	// Ubah field City dari address1 melalui address3
	address3.City = "Jakarta"

	// Cetak address1, address2, dan address3
	fmt.Println(address1) // Output: {Jakarta Jawa Barat Indonesia}
	fmt.Println(address2) // Output: {Bandung Jawa Barat Indonesia}
	fmt.Println(address3) // Output: &{Jakarta Jawa Barat Indonesia}

	// Buat pointer lain ke address1 dan ubah field Country
	address4 := &address1
	address4.Country = "USA"
	
	// Cetak address1, address2, address3, dan address4
	fmt.Println(address1) // Output: {Jakarta Jawa Barat USA}
	fmt.Println(address2) // Output: {Bandung Jawa Barat Indonesia}
	fmt.Println(address3) // Output: &{Jakarta Jawa Barat USA}
	fmt.Println(address4) // Output: &{Jakarta Jawa Barat USA}

	// Buat instance baru dari Address
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// Buat pointer ke address1
	var address2 *Address = &address1 // Pointer ke address1

	// Ubah field City dari address1 melalui address2
	address2.City = "Bandung"

	// Cetak address1 dan address2
	fmt.Println(address1) // Output: {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // Output: &{Bandung Jawa Barat Indonesia}
}