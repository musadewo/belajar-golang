package main

import "fmt"

// factorialLoop menghitung faktorial menggunakan loop
func factorialLoop(value int) int {
    result := 1 // Inisialisasi hasil dengan 1
    // Loop dari value ke 1
    for i := value; i > 0; i-- {
        result *= i // Mengalikan result dengan i
    }
    return result // Mengembalikan hasil faktorial
}

// factorialRecursive menghitung faktorial menggunakan rekursi
func factorialRecursive(value int) int {
    // Basis kasus: jika value adalah 1, kembalikan 1
    if value == 1 {
        return 1
    }
    // Rekursi: value dikalikan dengan hasil faktorial dari (value-1)
    return value * factorialRecursive(value-1)
}

func main() {
    // Mencetak hasil faktorial dari 10 menggunakan loop
    fmt.Println(factorialLoop(10))
    // Mencetak hasil faktorial dari 10 menggunakan rekursi
    fmt.Println(factorialRecursive(10))
}