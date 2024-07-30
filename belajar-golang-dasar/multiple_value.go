package main

import "fmt"

// getFullName mengembalikan dua nilai string: nama depan dan nama belakang
func getFullName() (string, string) {
	// Mengembalikan nama depan "Sadewo" dan nama belakang "Wicak"
	return "Sadewo", "Wicak"
}

func main() {
	// Memanggil getFullName dan menyimpan hasilnya ke dalam firstName dan lastName
	firstName, lastName := getFullName()
	// Mencetak nama depan dan nama belakang
	fmt.Println(firstName, lastName)

	// Memanggil getFullName lagi, tetapi hanya menyimpan nama depan
	firstName, _ = getFullName()
	// Mencetak nama depan saja
	fmt.Println(firstName)
}