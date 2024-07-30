package main

import "fmt"

func main() {
	// Mendeklarasikan konstanta dengan tipe data string
	const firstName string = "Agus" // Konstanta firstName bertipe string dengan nilai "Agus"
	
	// Mendeklarasikan konstanta tanpa menyebutkan tipe data (inferensi tipe)
	const lastName = "Dayat" // Konstanta lastName dengan inferensi tipe string dan nilai "Dayat"

	// Mencetak nilai konstanta ke konsol
	fmt.Println(firstName, lastName) // Mencetak "Agus Dayat" ke konsol
}