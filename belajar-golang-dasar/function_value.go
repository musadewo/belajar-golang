package main

import "fmt"

// getGoodBye mengembalikan string "Good Bye" yang diikuti dengan nama yang diberikan
func getGoodBye(name string) string {
	// Menggabungkan string "Good Bye " dengan parameter name
	return "Good Bye " + name
}

func main() {
	// Mendeklarasikan variabel goodBye yang merujuk ke fungsi getGoodBye
	goodBye := getGoodBye
	// Mendeklarasikan variabel contoh yang juga merujuk ke fungsi getGoodBye
	contoh := getGoodBye

	// Memanggil fungsi goodBye dengan argumen "John" dan mencetak hasilnya
	fmt.Println(goodBye("John"))
	// Memanggil fungsi goodBye dengan argumen "Doe" dan mencetak hasilnya
	fmt.Println(goodBye("Doe"))

	// Memanggil fungsi contoh dengan argumen "John" dan mencetak hasilnya
	fmt.Println(contoh("John"))
	// Memanggil fungsi contoh dengan argumen "Doe" dan mencetak hasilnya
	fmt.Println(contoh("Doe"))
}