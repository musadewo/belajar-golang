package main

import (
	"fmt"    // Package fmt menyediakan fungsi untuk format I/O, seperti Println
	"slices" // Package slices menyediakan fungsi-fungsi untuk manipulasi slice
)

func main() {
	// Mendefinisikan slice of string dengan nama 'names'
	names := []string{"Muhammad", "Sadewo", "Wicaksono"}
	
	// Mendefinisikan slice of int dengan nama 'values'
	values := []int{90, 80, 70, 60, 50}
	
	// Menggunakan fungsi Min dari package slices untuk mendapatkan nilai minimum dari slice 'values'
	fmt.Println(slices.Min(values))
	
	// Menggunakan fungsi Max dari package slices untuk mendapatkan nilai maksimum dari slice 'values'
	fmt.Println(slices.Max(values))
	
	// Menggunakan fungsi Contains dari package slices untuk memeriksa apakah 'names' mengandung "Sadewo"
	fmt.Println(slices.Contains(names, "Sadewo"))
	
	// Menggunakan fungsi Index dari package slices untuk mendapatkan indeks dari "Sadewo" dalam slice 'names'
	fmt.Println(slices.Index(names, "Sadewo"))
}