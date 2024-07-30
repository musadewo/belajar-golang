package main

import (
	"fmt"     // Mengimpor paket fmt untuk format I/O
	"reflect" // Mengimpor paket reflect untuk refleksi runtime
)

// Definisi struct Sample dengan satu field Name
type Sample struct {
	Name string
}

// Definisi struct Person dengan beberapa field dan tag
type Person struct {
	Name    string `required:"true" max:"10"`    // Field Name dengan tag required dan max
	Address string `required:"false" max:"10"`   // Field Address dengan tag required dan max
	Email   string `required:"true" max:"10"`    // Field Email dengan tag required dan max
}

// Fungsi readField untuk membaca dan mencetak informasi field dari struct
func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)  // Mendapatkan tipe dari value menggunakan refleksi

	fmt.Println("Type Name: ", valueType.Name())  // Mencetak nama tipe
	for i := 0; i < valueType.NumField(); i++ {   // Iterasi melalui semua field dalam struct
		var structField reflect.StructField = valueType.Field(i)  // Mendapatkan informasi field
		fmt.Println(structField.Name, "with type", structField.Type)  // Mencetak nama dan tipe field
		fmt.Println("Required: ", structField.Tag.Get("required"))    // Mencetak nilai tag "required"
		fmt.Println("Max: ", structField.Tag.Get("max"))              // Mencetak nilai tag "max"
	}
}

// Fungsi main sebagai titik masuk program
func main() {
	readField(Sample{"Dewo"})  // Membaca field dari struct Sample
	readField(Person{"Dewo", "Jl. Imam Bonjol", "dewo@gmail.com"})  // Membaca field dari struct Person
}