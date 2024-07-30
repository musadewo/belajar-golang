package main

import (
	"fmt"    // Paket fmt digunakan untuk format I/O, seperti mencetak ke konsol
	"regexp" // Paket regexp digunakan untuk bekerja dengan regular expressions
)

func main() {
	// Membuat pointer ke objek Regexp yang dihasilkan dari kompilasi pola regex `d([a-z])o`
	// Pola `d([a-z])o` akan mencocokkan string yang dimulai dengan 'd', diikuti oleh satu huruf kecil, dan diakhiri dengan 'o'
	var regex *regexp.Regexp = regexp.MustCompile(`d([a-z])o`)

	// Mencetak hasil pencocokan string "dwo" dengan pola regex
	// Output: true, karena "dwo" cocok dengan pola `d([a-z])o`
	fmt.Println(regex.MatchString("dwo"))
	
	// Mencetak hasil pencocokan string "dio" dengan pola regex
	// Output: true, karena "dio" cocok dengan pola `d([a-z])o`
	fmt.Println(regex.MatchString("dio"))
	
	// Mencetak hasil pencocokan string "dewo" dengan pola regex
	// Output: false, karena "dewo" tidak cocok dengan pola `d([a-z])o`
	fmt.Println(regex.MatchString("dewo"))

	// Mencetak semua string yang cocok dengan pola regex dalam string "dwo dio dui dio"
	// Output: ["dwo", "dio", "dio"], karena "dwo", "dio", dan "dio" adalah string yang cocok dengan pola `d([a-z])o`
	fmt.Println(regex.FindAllString("dwo dio dui dio", -1))
}