package main

import (
	"fmt"
)

func main() {
	// Mendeklarasikan dan menginisialisasi peta (map) dengan data
	person := map[string]string{
		"name": "John", // Menambahkan pasangan kunci-nilai "name": "John"
		"age":  "30",   // Menambahkan pasangan kunci-nilai "age": "30"
	}
	// Mencetak seluruh peta
	fmt.Println(person) // Output: map[name:John age:30]

	// Mencetak nilai dari kunci "name"
	fmt.Println(person["name"]) // Output: John

	// Mencetak nilai dari kunci "age"
	fmt.Println(person["age"]) // Output: 30

	// Membuat peta kosong menggunakan make
	book := make(map[string]string) // Inisialisasi peta kosong

	// Menambahkan pasangan kunci-nilai ke peta
	book["title"] = "Belajar Golang" // Menambahkan pasangan kunci-nilai "title": "Belajar Golang"
	book["author"] = "John Doe"      // Menambahkan pasangan kunci-nilai "author": "John Doe"
	book["ups"] = "Salah"            // Menambahkan pasangan kunci-nilai "ups": "Salah"

	// Mencetak seluruh peta
	fmt.Println(book) // Output: map[title:Belajar Golang author:John Doe ups:Salah]

	// Menghapus pasangan kunci-nilai dengan kunci "ups"
	delete(book, "ups") // Menghapus pasangan kunci-nilai dengan kunci "ups"

	// Mencetak peta setelah penghapusan
	fmt.Println(book) // Output: map[title:Belajar Golang author:John Doe]
}