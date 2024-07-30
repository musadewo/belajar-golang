package internal

import "fmt"

// init adalah fungsi khusus yang akan dipanggil secara otomatis
// ketika paket diimpor. Ini digunakan untuk inisialisasi.
func init() {
	fmt.Println("This is internal package")
}