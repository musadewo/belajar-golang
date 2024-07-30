package main

import "fmt"

func main() {
    // Mencetak "Hello, World!" ke konsol
    fmt.Println("Hello, World!")
    
    // Mencetak "Hello, World!2" ke konsol
    fmt.Println("Hello, World!2")
    
    // Mencetak "Hello, World3!" ke konsol
    fmt.Println("Hello, World3!")
    
    // Mencetak panjang dari string "Hello, World!" ke konsol
    // Fungsi len() mengembalikan jumlah karakter dalam string
    fmt.Println(len("Hello, World!"))
    
    // Mencetak karakter kedua dari string "Hello, World!" ke konsol
    // Mengakses karakter kedua (indeks 1) dari string
    fmt.Println("Hello, World!"[1])
    
    // Mencetak karakter keempat dari string "Hello, World!" ke konsol
    // Mengakses karakter keempat (indeks 3) dari string
    fmt.Println("Hello, World!"[3])
    
    // Mencetak substring dari karakter pertama hingga kelima dari string "Hello, World!" ke konsol
    // Mengambil substring dari indeks 0 hingga 4 (karakter pertama hingga kelima)
    fmt.Println("Hello, World!"[0:5])
}