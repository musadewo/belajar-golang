package main

import "fmt"

// sayHelloTo adalah fungsi yang mencetak pesan sapaan dengan nama depan dan nama belakang
func sayHelloTo(firstName string, lastName string) {
	// Mencetak pesan sapaan dengan format "Hello {firstName} {lastName}"
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	// Memanggil fungsi sayHelloTo dengan parameter "Sadewo" sebagai firstName dan "Wicak" sebagai lastName
	sayHelloTo("Sadewo", "Wicak")
}