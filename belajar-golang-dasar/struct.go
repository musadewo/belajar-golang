package main

import "fmt"

// Definisi struct Customer dengan field Name, Address, dan Age
type Customer struct {
	Name, Address string // Field Name dan Address bertipe string
	Age           int    // Field Age bertipe int
}

// Method sayHello untuk struct Customer
func (customer Customer) sayHello(name string) {
	// Mencetak pesan sapaan dengan nama yang diberikan dan nama customer
	fmt.Println("Hello,", name, "My Name is", customer.Name)
}

func main() {
	// Deklarasi variabel dewo dengan tipe Customer
	var dewo Customer
	// Mencetak nilai default dari dewo (semua field kosong atau nol)
	fmt.Println(dewo)

	// Deklarasi dan inisialisasi variabel customer1 dengan tipe Customer
	var customer1 Customer
	customer1.Name = "John Doe"           // Mengisi field Name dengan "John Doe"
	customer1.Address = "Jl. Imam Bonjol" // Mengisi field Address dengan "Jl. Imam Bonjol"
	customer1.Age = 30                    // Mengisi field Age dengan 30

	// Mencetak seluruh field dari customer1
	fmt.Println(customer1)
	// Mencetak field Name dari customer1
	fmt.Println(customer1.Name)
	// Mencetak field Address dari customer1
	fmt.Println(customer1.Address)
	// Mencetak field Age dari customer1
	fmt.Println(customer1.Age)

	// Deklarasi dan inisialisasi variabel joko dengan tipe Customer menggunakan literal struct
	joko := Customer{
		Name:    "Joko",            // Mengisi field Name dengan "Joko"
		Address: "Jl. Imam Bonjol", // Mengisi field Address dengan "Jl. Imam Bonjol"
		Age:     30,                // Mengisi field Age dengan 30
	}
	// Mencetak seluruh field dari joko
	fmt.Println(joko)

	// Memanggil method sayHello dari struct Customer dengan parameter "agus"
	dewo.sayHello("agus")
}