package main

import "fmt"

// HasName adalah interface yang memiliki satu metode GetName
type HasName interface {
	// GetName mengembalikan nama sebagai string
	GetName() string
}

// SayHello menerima parameter yang mengimplementasikan interface HasName
// dan mencetak pesan sapaan menggunakan nama yang didapat dari metode GetName
func SayHello(hasName HasName) {
	// Mencetak pesan sapaan dengan nama yang didapat dari metode GetName
	fmt.Println("Hello", hasName.GetName())
}

// Person adalah struct yang memiliki satu field Name
type Person struct {
	// Name adalah nama dari person
	Name string
}

// GetName mengimplementasikan metode GetName dari interface HasName untuk struct Person
func (person Person) GetName() string {
	// Mengembalikan nama dari person
	return person.Name
}

// Animal adalah struct yang memiliki satu field Name
type Animal struct {
	// Name adalah nama dari animal
	Name string
}

// GetName mengimplementasikan metode GetName dari interface HasName untuk struct Animal
func (animal Animal) GetName() string {
	// Mengembalikan nama dari animal
	return animal.Name
}

func main() {
	// Membuat instance dari struct Person dengan nama "Dewo"
	person := Person{Name: "Dewo"}
	// Memanggil fungsi SayHello dengan parameter person
	SayHello(person)

	// Membuat instance dari struct Animal dengan nama "Cat"
	animal := Animal{Name: "Cat"}
	// Memanggil fungsi SayHello dengan parameter animal
	SayHello(animal)
}