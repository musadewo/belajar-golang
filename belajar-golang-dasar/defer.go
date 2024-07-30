package main

import "fmt"

// logging adalah fungsi yang akan mencetak pesan ketika dipanggil
func logging() {
    fmt.Println("Selesai memanggil function")
}

// runApplication adalah fungsi yang menjalankan aplikasi
// menggunakan defer untuk memastikan logging dipanggil setelah fungsi ini selesai
func runApplication() {
    defer logging() // defer memastikan logging dipanggil terakhir sebelum runApplication selesai
    fmt.Println("Run Application") // mencetak pesan bahwa aplikasi sedang berjalan
}

func main() {
    runApplication() // memanggil fungsi runApplication dari fungsi utama
}