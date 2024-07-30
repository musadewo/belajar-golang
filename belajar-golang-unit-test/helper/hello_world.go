// Package helper berisi fungsi-fungsi pembantu untuk aplikasi
package helper

// Import yang digunakan:
// Tidak ada import yang digunakan dalam file ini karena fungsi hanya menggunakan tipe data bawaan Go

// HelloWorld adalah fungsi yang menerima nama sebagai input dan mengembalikan salam
// Param:
//   - name: string, nama yang akan disapa
// Return:
//   - string, salam yang berisi "Hello, " diikuti dengan nama yang diberikan
func HelloWorld(name string) string {
	// Menggabungkan string "Hello, " dengan nama yang diberikan
	// Operator + digunakan untuk menggabungkan string dalam Go
	return "Hello, " + name
}