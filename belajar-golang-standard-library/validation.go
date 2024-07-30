package main

import (
	"fmt"     // Package untuk format I/O
	"reflect" // Package untuk refleksi runtime
)

// Definisi struct Sample dengan satu field
type Sample struct {
	Name string
}

// Definisi struct Person dengan beberapa field yang memiliki tag
type Person struct {
	Name    string `required:"true" max:"10"`  // Field dengan tag 'required' dan 'max'
	Address string `required:"true" max:"10"`  // Field dengan tag 'required' dan 'max'
	Email   string `required:"true" max:"10"`  // Field dengan tag 'required' dan 'max'
}

// Fungsi untuk membaca dan mencetak informasi field dari struct menggunakan refleksi
func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)  // Mendapatkan tipe dari value

	fmt.Println("Type Name: ", valueType.Name())  // Mencetak nama tipe
	for i := 0; i < valueType.NumField(); i++ {   // Iterasi melalui semua field dalam struct
		var structField reflect.StructField = valueType.Field(i)  // Mendapatkan informasi field
		fmt.Println(structField.Name, "with type", structField.Type)  // Mencetak nama dan tipe field
		fmt.Println("Required: ", structField.Tag.Get("required"))    // Mencetak nilai tag 'required'
		fmt.Println("Max: ", structField.Tag.Get("max"))              // Mencetak nilai tag 'max'
	}
}

// Fungsi untuk memvalidasi apakah semua field yang ditandai 'required' tidak kosong
func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)  // Mendapatkan tipe dari value
	for i := 0; i < t.NumField(); i++ {  // Iterasi melalui semua field dalam struct
		f := t.Field(i)  // Mendapatkan informasi field
		if f.Tag.Get("required") == "true" {  // Memeriksa apakah field memiliki tag 'required'
			data := reflect.ValueOf(value).Field(i).Interface()  // Mendapatkan nilai field
			result = data != ""  // Memeriksa apakah nilai field tidak kosong
			if result == false {  // Jika ada field yang kosong, kembalikan false
				return result
			}
		}
	}
	return result  // Kembalikan true jika semua field 'required' tidak kosong
}

func main() {
	person := Person{"Dewo", "Jl. Imam Bonjol", "dewo@gmail.com"}  // Membuat instance dari struct Person
	fmt.Println(IsValid(person))  // Memvalidasi instance dan mencetak hasilnya
}