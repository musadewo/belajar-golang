package main

import "fmt"

func main() {
    // Deklarasi variabel dengan tipe data string dan inisialisasi nilai
    var firstName string = "John" // Deklarasi variabel firstName dengan tipe data string dan nilai awal "John"
    var lastName string           // Deklarasi variabel lastName dengan tipe data string tanpa nilai awal
    lastName = "Doe"              // Inisialisasi nilai variabel lastName dengan "Doe"
    fmt.Println(firstName)        // Mencetak nilai variabel firstName ke konsol
    fmt.Println(lastName)         // Mencetak nilai variabel lastName ke konsol
    fmt.Println(firstName, lastName) // Mencetak nilai variabel firstName dan lastName ke konsol

    // Deklarasi variabel dengan tipe data string dan inisialisasi nilai
    var name string               // Deklarasi variabel name dengan tipe data string tanpa nilai awal
    name = "Agus"                 // Inisialisasi nilai variabel name dengan "Agus"
    fmt.Printf(name)              // Mencetak nilai variabel name ke konsol tanpa karakter newline

    // Cara singkat deklarasi dan inisialisasi variabel (banyak digunakan)
    name := "Dayat"               // Deklarasi dan inisialisasi variabel name dengan nilai "Dayat" menggunakan short declaration
    fmt.Println(name)             // Mencetak nilai variabel name ke konsol

    // Deklarasi beberapa variabel sekaligus dengan inisialisasi nilai
    var (
        firstName  = "Agus"       // Deklarasi variabel firstName dengan nilai "Agus"
        middleName = "Oya"        // Deklarasi variabel middleName dengan nilai "Oya"
        lastName   = "Dayat"      // Deklarasi variabel lastName dengan nilai "Dayat"
    )
    fmt.Println(firstName, middleName, lastName) // Mencetak nilai variabel firstName, middleName, dan lastName ke konsol
}