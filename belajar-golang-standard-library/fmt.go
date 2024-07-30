package main

import "fmt"

// Fungsi utama yang akan dieksekusi saat program dijalankan
func main() {
    // Mendeklarasikan variabel firstname dengan nilai "M"
    // Tipe data string digunakan untuk menyimpan teks
    firstname := "M"
    
    // Mendeklarasikan variabel lastname dengan nilai "Dewo"
    // Tipe data string digunakan untuk menyimpan teks
    lastname := "Dewo"

    // Mencetak string dengan menggunakan fmt.Println
    // fmt.Println menambahkan spasi otomatis antara argumen dan menambahkan newline di akhir
    // Argumen pertama adalah string "Hello '"
    // Argumen kedua adalah variabel firstname yang nilainya "M"
    // Argumen ketiga adalah variabel lastname yang nilainya "Dewo"
    // Argumen keempat adalah string "'"
    fmt.Println("Hello '", firstname, lastname, "'")

    // Mencetak string dengan menggunakan fmt.Printf
    // fmt.Printf memungkinkan format string yang lebih fleksibel
    // '%s' adalah placeholder untuk string
    // '\n' menambahkan karakter newline (baris baru) di akhir string
    // Argumen pertama adalah format string "Hello '%s %s' \n"
    // Argumen kedua adalah variabel firstname yang nilainya "M"
    // Argumen ketiga adalah variabel lastname yang nilainya "Dewo"
    fmt.Printf("Hello '%s %s' \n", firstname, lastname)
}