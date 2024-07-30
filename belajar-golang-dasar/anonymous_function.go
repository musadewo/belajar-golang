package main

import "fmt"

// Blacklist adalah tipe data fungsi yang menerima string dan mengembalikan boolean
type Blacklist func(string) bool

// registerUser adalah fungsi yang menerima nama dan fungsi blacklist
func registerUser(name string, blacklist Blacklist) {
	// Memeriksa apakah nama ada dalam blacklist
	if blacklist(name) {
		// Jika nama ada dalam blacklist, cetak pesan bahwa user diblokir
		fmt.Println("You are blocked", name)
	} else {
		// Jika nama tidak ada dalam blacklist, cetak pesan selamat datang
		fmt.Println("Welcome", name)
	}
}

func main() {
	// Mendaftarkan user "Eko" dengan fungsi blacklist anonim
	registerUser("Eko", func(name string) bool {
		// Fungsi anonim yang memblokir nama "admin"
		return name == "admin"
	})

	// Mendefinisikan fungsi blacklist yang memblokir nama "admin"
	blacklist := func(name string) bool {
		// Fungsi yang memblokir nama "admin"
		return name == "admin"
	}

	// Mendaftarkan user "admin" dan "eko" dengan fungsi blacklist yang telah didefinisikan
	registerUser("admin", blacklist) // User "admin" akan diblokir
	registerUser("eko", blacklist)   // User "eko" tidak akan diblokir
}