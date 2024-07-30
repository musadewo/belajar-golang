package main

import "fmt"

func main() {
    // Mencetak hasil penjumlahan 10 + 10
    fmt.Println(10 + 10) // Output: 20

    // Deklarasi variabel a, b, dan z dengan nilai awal
    var a = 10
    var b = 20
    var z = 5

    // Operasi aritmatika dengan variabel a, b, dan z
    var c = a + b + z // Penjumlahan: 10 + 20 + 5 = 35
    var d = a - b - z // Pengurangan: 10 - 20 - 5 = -15
    var e = a * b + z // Perkalian dan penjumlahan: 10 * 20 + 5 = 205
    var f = a / b / z // Pembagian: 10 / 20 / 5 = 0 (karena integer division)

    // Mencetak hasil operasi aritmatika
    fmt.Println(c) // Output: 35
    fmt.Println(d) // Output: -15
    fmt.Println(e) // Output: 205
    fmt.Println(f) // Output: 0

    // Deklarasi variabel i dan operasi penugasan
    var i = 10
    i += 10 // Menambahkan 10 ke i: 10 + 10 = 20
    fmt.Println(i) // Output: 20

    i -= 5 // Mengurangi 5 dari i: 20 - 5 = 15
    fmt.Println(i) // Output: 15

    // Deklarasi variabel j dan operasi inkremen/dekrement
    var j = 10
    j++ // Menambah 1 ke j: 10 + 1 = 11
    fmt.Println(j) // Output: 11

    j-- // Mengurangi 1 dari j: 11 - 1 = 10
    fmt.Println(j) // Output: 10

    j++ // Menambah 1 ke j lagi: 10 + 1 = 11
    fmt.Println(j) // Output: 11
}