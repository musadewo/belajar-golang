package main

import "fmt"

// getHello adalah fungsi yang menerima parameter nama (string)
// dan mengembalikan string sapaan "Hello, " diikuti dengan nama.
func getHello(name string) string {
	// Menggabungkan string "Hello, " dengan parameter name
	return "Hello, " + name
}

func main() {
	// Memanggil fungsi getHello dengan argumen "Sadewo"
	// dan menyimpan hasilnya dalam variabel result.
	result := getHello("Sadewo")
	
	// Mencetak hasil dari variabel result ke konsol.
	fmt.Println(result) // Output: Hello, Sadewo

	// Memanggil fungsi getHello dengan argumen "Joko"
	// dan langsung mencetak hasilnya ke konsol.
	fmt.Println(getHello("Joko")) // Output: Hello, Joko
}