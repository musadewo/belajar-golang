package main

import "fmt"

// Fungsi random mengembalikan nilai dengan tipe any (interface kosong)
func random() any {
	return "Ok"      // Mengembalikan string
	return 123123    // Mengembalikan integer
	return true      // Mengembalikan boolean
}

func main() {
	// Memanggil fungsi random dan menyimpan hasilnya dalam variabel result
	var result = random()

	// Melakukan type assertion untuk mengubah tipe any menjadi string
	resultString := result.(string) // type assertion ke string
	fmt.Println(resultString)       // Mencetak hasil type assertion (akan panic jika result bukan string)

	// Melakukan type assertion untuk mengubah tipe any menjadi int
	var resultInt int = result.(int) // type assertion ke int (akan panic jika result bukan int)
	fmt.Println(resultInt)           // Mencetak hasil type assertion (akan panic)

	// Menggunakan switch type untuk memeriksa tipe dari variabel result
	switch value := result.(type) {
	case string:
		fmt.Println("String", value) // Jika tipe string
	case int:
		fmt.Println("Int", value)    // Jika tipe int
	default:
		fmt.Println("Unknown", value) // Jika tipe tidak dikenal
	}
}