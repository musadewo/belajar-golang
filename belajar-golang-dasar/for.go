package main

import "fmt"

func main() {
    // Inisialisasi variabel counter dengan nilai 1
    counter := 1

    // Perulangan while-like menggunakan for
    for counter <= 3 {
        // Mencetak nilai counter saat ini
        fmt.Println("perulangan ke ", counter)
        // Increment counter
        counter++
    }
    // Mencetak pesan selesai setelah perulangan selesai
    fmt.Println("Selesai")

    // Perulangan for dengan kondisi dan increment
    for i := 0; i < 3; i++ {
        // Mencetak nilai i saat ini
        fmt.Println("Perulangan ke-", i)
    }
    // Mencetak pesan selesai setelah perulangan selesai
    fmt.Println("Selesai")

    // Perulangan for dengan deklarasi variabel di dalam for
    for counter := 1; counter <= 3; counter++ {
        // Mencetak nilai counter saat ini
        fmt.Println("Perulangan ke-", counter)
    }
    // Mencetak pesan selesai setelah perulangan selesai
    fmt.Println("Selesai")

    // Inisialisasi array name dengan beberapa elemen string
    name := []string{"Sadewo", "Dewo", "Wo"}
    // Perulangan for untuk mengakses elemen array berdasarkan indeks
    for i := 0; i < len(name); i++ {
        // Mencetak elemen array pada indeks i
        fmt.Println(name[i])
    }
    // Mencetak pesan selesai setelah perulangan selesai
    fmt.Println("Selesai")

    // Perulangan for dengan range untuk mengakses elemen array
    for index, value := range name {
        // Mencetak indeks dan nilai elemen array
        fmt.Println("Nama ke-", index, "adalah", value)
    }
    // Mencetak pesan selesai setelah perulangan selesai
    fmt.Println("Selesai")

    // Perulangan for dengan range tanpa menggunakan index
    for _, value := range name {
        // Mencetak nilai elemen array tanpa mencetak indeks
        fmt.Println("Nama ", value)
    }
    // Mencetak pesan selesai setelah perulangan selesai
    fmt.Println("Selesai")
}