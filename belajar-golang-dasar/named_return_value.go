package main

import "fmt"

// getCompleteName mengembalikan nama lengkap yang terdiri dari firstName, middleName, dan lastName
func getCompleteName() (firstName, middleName, lastName string) {
	// Menginisialisasi nilai dari firstName, middleName, dan lastName
	firstName = "Muhammad"  // Inisialisasi firstName dengan nilai "Muhammad"
	middleName = "Sadewo"   // Inisialisasi middleName dengan nilai "Sadewo"
	lastName = "Wicaksono"  // Inisialisasi lastName dengan nilai "Wicaksono"

	// Mengembalikan nilai-nilai tersebut
	return firstName, middleName, lastName  // Mengembalikan firstName, middleName, dan lastName
}

func main() {
	// Memanggil fungsi getCompleteName dan menyimpan hasilnya ke dalam variabel
	firstName, middleName, lastName := getCompleteName()  // Memanggil fungsi dan menyimpan hasil ke variabel firstName, middleName, lastName
	
	// Mencetak nama lengkap ke konsol
	fmt.Println(firstName, middleName, lastName)  // Mencetak nilai-nilai variabel ke konsol

	// Memanggil fungsi getCompleteName lagi dan menyimpan hasilnya ke dalam variabel a, b, c
	a, b, c := getCompleteName()  // Memanggil fungsi lagi dan menyimpan hasil ke variabel a, b, c
	
	// Mencetak nama lengkap ke konsol menggunakan variabel a, b, c
	fmt.Println(a, b, c)  // Mencetak nilai-nilai variabel a, b, c ke konsol
}