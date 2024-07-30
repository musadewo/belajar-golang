package main

import "fmt"

// sayHelloWithFilter adalah fungsi yang menerima nama dan filter sebagai parameter
// name: string yang akan difilter
// filter: fungsi yang menerima string dan mengembalikan string
func sayHelloWithFilter(name string, filter func(string) string) {
    // Mencetak "Hello, " diikuti dengan hasil filter dari nama
    fmt.Println("Hello, " + filter(name))
}

// spamFilter adalah fungsi yang memfilter kata "Anjing"
// name: string yang akan diperiksa
// return: jika name adalah "Anjing", kembalikan "...", jika tidak, kembalikan name
func spamFilter(name string) string {
    if name == "Anjing" {
        // Jika nama adalah "Anjing", kembalikan "..."
        return "..."
    } else {
        // Jika tidak, kembalikan nama asli
        return name
    }
}

func main() {
    // Memanggil sayHelloWithFilter dengan nama "Eko" dan filter spamFilter
    // Output: "Hello, Eko"
    sayHelloWithFilter("Eko", spamFilter)

    // Menyimpan spamFilter ke dalam variabel filter
    filter := spamFilter
    // Memanggil sayHelloWithFilter dengan nama "Anjing" dan filter yang disimpan
    // Output: "Hello, ..."
    sayHelloWithFilter("Anjing", filter)
}