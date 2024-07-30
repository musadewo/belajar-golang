package main

import (
	"fmt"
)

// endApp akan dipanggil ketika aplikasi selesai
func endApp() {
    // Menampilkan pesan bahwa aplikasi telah selesai
    fmt.Println("End App")
    // recover digunakan untuk menangkap panic dan mengembalikan nilainya
	message := recover()
    // Menampilkan pesan error yang ditangkap oleh recover
    fmt.Println("Error :", message)
}

// runApp menjalankan aplikasi dan memicu panic jika terjadi error
func runApp(error bool) {
    // defer memastikan endApp dipanggil setelah runApp selesai
    defer endApp()
    if error {
        // panic digunakan untuk menghentikan eksekusi dan menampilkan pesan error
        panic("Error")
    }
    // Menampilkan pesan bahwa aplikasi berjalan normal jika tidak ada error
    fmt.Println("App is running normally")
}

func main() {
    // Memanggil runApp dengan parameter true, sehingga akan terjadi panic
    runApp(true)
    // Memanggil runApp dengan parameter false, sehingga aplikasi berjalan normal
    runApp(false)
}