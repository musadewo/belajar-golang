package main

import (
	"fmt"  // Mengimpor paket fmt untuk format I/O seperti Println
	"sort" // Mengimpor paket sort untuk melakukan pengurutan slice
)

// Mendefinisikan struct User dengan dua field: Name dan Age
type User struct {
	Name string
	Age  int
}

// Mendefinisikan tipe UserSlice sebagai slice dari User
type UserSlice []User

// Metode Len untuk mendapatkan panjang dari UserSlice
func (s UserSlice) Len() int {
	return len(s)
}

// Metode Less untuk membandingkan dua elemen dalam UserSlice berdasarkan Age
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// Metode Swap untuk menukar dua elemen dalam UserSlice
func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// Membuat slice dari User dengan beberapa data
	users := []User{
		{"M", 30},
		{"Sadewo", 20},
		{"Wicaksono", 25},
	}

	// Mengurutkan slice users menggunakan sort.Sort dan UserSlice
	sort.Sort(UserSlice(users))

	// Mencetak hasil pengurutan
	fmt.Println(users)
}