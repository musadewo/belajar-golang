package main

import (
	"fmt"  // Mengimpor paket fmt untuk format I/O
	"path" // Mengimpor paket path untuk manipulasi path file
)

func main() {
	// Menggunakan fungsi Dir dari paket path untuk mendapatkan direktori dari path yang diberikan
	fmt.Println(path.Dir("hello/world.go")) // Output: "hello"

	// Menggunakan fungsi Base dari paket path untuk mendapatkan nama file dari path yang diberikan
	fmt.Println(path.Base("hello/world.go")) // Output: "world.go"

	// Menggunakan fungsi Ext dari paket path untuk mendapatkan ekstensi file dari path yang diberikan
	fmt.Println(path.Ext("hello/world.go")) // Output: ".go"

	// Menggunakan fungsi Join dari paket path untuk menggabungkan beberapa elemen path menjadi satu path
	fmt.Println(path.Join("hello", "world", "main.go")) // Output: "hello/world/main.go"
}