package main

import (
	"fmt"           // Package fmt menyediakan fungsi untuk format I/O, seperti Println
	"path/filepath" // Package filepath menyediakan fungsi untuk memanipulasi path file
)

func main() {
	// Menampilkan direktori dari path yang diberikan
	fmt.Println(filepath.Dir("hello/world.go")) // Output: "hello"

	// Menampilkan nama file dari path yang diberikan
	fmt.Println(filepath.Base("hello/world.go")) // Output: "world.go"

	// Menampilkan ekstensi file dari path yang diberikan
	fmt.Println(filepath.Ext("hello/world.go")) // Output: ".go"

	// Mengecek apakah path yang diberikan adalah path absolut
	fmt.Println(filepath.IsAbs("hello/world.go")) // Output: false (karena ini adalah path relatif)

	// Mengecek apakah path yang diberikan adalah path lokal
	fmt.Println(filepath.IsLocal("hello/world.go")) // Output: true (karena ini adalah path relatif)

	// Menggabungkan beberapa elemen path menjadi satu path
	fmt.Println(filepath.Join("hello", "world", "main.go")) // Output: "hello/world/main.go"
}