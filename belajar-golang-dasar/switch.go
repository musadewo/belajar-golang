package main

import "fmt"

// Fungsi utama yang akan dieksekusi pertama kali
func main() {
    // Mendeklarasikan variabel 'name' dengan nilai "Sadewo"
    name := "Sadewo"

    // Memulai pernyataan switch untuk variabel 'name'
    switch name {
    // Jika 'name' adalah "Sadewo"
    case "Sadewo":
        fmt.Println("Halo Sadewo")
    // Jika 'name' adalah "Dewo"
    case "Dewo":
        fmt.Println("Halo Dewo")
    // Jika 'name' tidak cocok dengan case manapun
    default:
        fmt.Println("Halo, Boleh Kenal?")
    }

    // Switch dengan kondisi panjang string 'name'
    switch length := len(name); length > 5 {
    // Jika panjang string lebih dari 5
    case true:
        fmt.Println("Nama terlau panjang")
    // Jika panjang string tidak lebih dari 5
    case false:
        fmt.Println("Nama pendek sudah benar")
    }

    // Mengubah nilai 'name' menjadi "Sadewo"
    name = "Sadewo"
    // Mendapatkan panjang string 'name'
    length := len(name)
    // Switch tanpa kondisi variabel, hanya menggunakan kondisi logika
    switch {
    // Jika panjang string lebih dari 5
    case length > 5:
        fmt.Println("Kependekan")
    // Jika panjang string lebih dari 10
    case length > 10:
        fmt.Println("Kepanjangan")
    // Jika tidak ada kondisi yang terpenuhi
    default:
        fmt.Println("Nama sudah benar")
    }
}