package main

import "fmt"

// Man adalah struct yang memiliki satu field, yaitu Name
type Man struct {
	Name string
}

// Method Married untuk Man (value receiver) - tidak mengubah nilai asli dari Man
// Method ini hanya mengubah salinan dari Man, bukan instance aslinya
func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

// Method Married untuk Man (pointer receiver) - mengubah nilai asli dari Man
// Method ini mengubah instance asli dari Man karena menggunakan pointer receiver
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	// Membuat instance dari Man dengan nama "Dewo"
	dewo := Man{"Dewo"}
	// Menampilkan instance awal dari Man
	fmt.Println(dewo) // Output: {Dewo}

	// Memanggil method Married dengan pointer receiver
	// Ini akan mengubah nilai asli dari instance dewo
	dewo.Married()
	// Menampilkan instance dewo setelah dipanggil method Married
	fmt.Println(dewo) // Output: {Mr. Dewo}
}