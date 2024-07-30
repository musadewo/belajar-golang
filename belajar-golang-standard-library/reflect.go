package main

import (
	"fmt"     // Mengimpor paket fmt untuk format I/O
	"reflect" // Mengimpor paket reflect untuk refleksi runtime
)

// Definisi struct Sample dengan satu field Name
type Sample struct {
	Name string
}

// Definisi struct Person dengan tiga field: Name, Address, dan Email
type Person struct {
	Name, Address, Email string
}

// Fungsi readField menerima parameter value dengan tipe any (interface kosong)
// Fungsi ini menggunakan refleksi untuk membaca dan mencetak nama dan tipe field dari struct yang diberikan
func readField(value any) {
	// Mendapatkan tipe dari value menggunakan refleksi
	var valueType reflect.Type = reflect.TypeOf(value)

	// Mencetak nama tipe dari value
	fmt.Println("Type Name: ", valueType.Name())
	
	// Loop melalui semua field dari struct
	for i := 0; i < valueType.NumField(); i++ {
		// Mendapatkan informasi field menggunakan refleksi
		var valueField reflect.StructField = valueType.Field(i)
		// Mencetak nama field dan tipe field
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

// Fungsi main adalah titik masuk dari program
func main() {
	// Memanggil fungsi readField dengan instance dari Sample
	readField(Sample{"Dewo"})
	// Memanggil fungsi readField dengan instance dari Person
	readField(Person{"Dewo", "Jl. Imam Bonjol", "dewo@gmail.com"})
}