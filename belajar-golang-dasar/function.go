package main

import "fmt"

// sayHello mencetak "Hello" ke konsol
func sayHello() {
    // fmt.Println adalah fungsi dari paket fmt yang digunakan untuk mencetak teks ke konsol
    fmt.Println("Hello")
}

// main adalah fungsi utama yang memanggil sayHello
func main() {
    // Memanggil fungsi sayHello untuk mencetak "Hello" ke konsol
    sayHello()
}