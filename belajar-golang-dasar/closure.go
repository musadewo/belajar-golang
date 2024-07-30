package main

import "fmt"

func main() {
    // Inisialisasi variabel counter dengan nilai awal 0
    counter := 0
    
    // Mendefinisikan fungsi anonim (closure) yang akan menambah nilai counter
    increment := func() {
        fmt.Println("Increment") // Mencetak "Increment" ke konsol untuk menunjukkan bahwa fungsi dipanggil
        counter++                // Menambah nilai counter sebesar 1
    }
    
    // Memanggil fungsi increment dua kali
    increment() // Panggilan pertama, counter akan menjadi 1
    increment() // Panggilan kedua, counter akan menjadi 2
    
    // Mencetak nilai counter ke konsol
    fmt.Println(counter) // Output: 2
}