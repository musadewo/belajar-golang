package main

import (
	"container/list" // Mengimpor paket list dari container
	"fmt"            // Mengimpor paket fmt untuk format I/O
)

func main() {
	// Membuat list baru
	var data *list.List = list.New()

	// Menambahkan elemen ke dalam list
	data.PushBack("M")         // Menambahkan elemen "M" ke akhir list
	data.PushBack("Sadewo")    // Menambahkan elemen "Sadewo" ke akhir list
	data.PushBack("Kurniawan") // Menambahkan elemen "Kurniawan" ke akhir list

	// Manual
	// Mengambil elemen pertama dari list
	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Mencetak nilai elemen pertama ("M")

	// Mengambil elemen berikutnya dari elemen pertama
	next := head.Next()
	fmt.Println(next.Value) // Mencetak nilai elemen kedua ("Sadewo")

	// Mengambil elemen berikutnya dari elemen kedua
	next = next.Next()
	fmt.Println(next.Value) // Mencetak nilai elemen ketiga ("Kurniawan")

	// Automatis
	// Iterasi melalui semua elemen dalam list
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // Mencetak nilai setiap elemen dalam list
	}
}