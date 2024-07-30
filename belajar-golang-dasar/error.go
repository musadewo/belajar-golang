package main

import (
	"errors"
	"fmt"
)

// Fungsi Pembagian membagi dua angka dan mengembalikan hasilnya atau error jika pembagi adalah 0
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		// Mengembalikan error jika pembagi adalah 0
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	// Mengembalikan hasil pembagian dan nil untuk error
	return nilai / pembagi, nil
}

func main() {
	// Contoh penggunaan fungsi Pembagian dengan pembagi 0
	hasil, err := Pembagian(10, 0)
	if err != nil {
		// Menampilkan pesan error jika terjadi error
		fmt.Println("Error: ", err)
	}
	// Menampilkan hasil pembagian, meskipun ada error, hasil tetap ditampilkan
	fmt.Println("Hasil: ", hasil)

	// Contoh penggunaan fungsi Pembagian dengan pembagi yang valid
	hasil, err = Pembagian(10, 2)
	if err == nil {
		// Menampilkan hasil jika tidak ada error
		fmt.Println("Hasil: ", hasil)
	} else {
		// Menampilkan pesan error jika terjadi error
		fmt.Println("Error: ", err)
	}
}