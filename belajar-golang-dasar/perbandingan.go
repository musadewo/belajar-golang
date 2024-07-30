package main

import "fmt"

func main() {
    // Deklarasi variabel a dan b dengan nilai 10 dan 20
    var a = 10
    var b = 20

    // Membandingkan nilai a dan b
    fmt.Println(a == b)  // Apakah a sama dengan b? (false)
    fmt.Println(a != b)  // Apakah a tidak sama dengan b? (true)
    fmt.Println(a > b)   // Apakah a lebih besar dari b? (false)
    fmt.Println(a < b)   // Apakah a lebih kecil dari b? (true)
    fmt.Println(a >= b)  // Apakah a lebih besar atau sama dengan b? (false)
    fmt.Println(a <= b)  // Apakah a lebih kecil atau sama dengan b? (true)

    // Deklarasi variabel name1 dan name2 dengan nilai "Agus"
    var name1 = "Agus"
    var name2 = "Agus"

    // Membandingkan nilai name1 dan name2
    var result bool = name1 == name2
    fmt.Println(result)  // Apakah name1 sama dengan name2? (true)
}