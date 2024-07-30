package main

import "fmt"

func main() {
	// Inisialisasi slice 'name' dengan beberapa nama
	name := []string{"Agus", "Kurniawan", "Khannedy", "Budi", "Joko"}
	
	// Membuat slice baru 'slice1' yang berisi elemen dari index 2 sampai 3
	slice1 := name[2:4]
	fmt.Println(slice1) // Output: [Khannedy Budi]
	fmt.Println(len(slice1)) // Panjang slice1
	fmt.Println(cap(slice1)) // Kapasitas slice1

	// Menambahkan elemen baru ke 'slice1' dan menyimpannya di 'slice2'
	slice2 := append(slice1, "Kurniawan")
	fmt.Println(slice2) // Output: [Khannedy Budi Kurniawan]
	fmt.Println(len(slice2)) // Panjang slice2
	fmt.Println(cap(slice2)) // Kapasitas slice2

	// Menambahkan elemen baru ke 'slice2' dan menyimpannya di 'slice3'
	slice3 := append(slice2, "Khannedy")
	fmt.Println(slice3) // Output: [Khannedy Budi Kurniawan Khannedy]
	fmt.Println(len(slice3)) // Panjang slice3
	fmt.Println(cap(slice3)) // Kapasitas slice3

	// Menambahkan elemen baru ke 'slice3' dan menyimpannya di 'slice4'
	slice4 := append(slice3, "Budi")
	fmt.Println(slice4) // Output: [Khannedy Budi Kurniawan Khannedy Budi]
	fmt.Println(len(slice4)) // Panjang slice4
	fmt.Println(cap(slice4)) // Kapasitas slice4

	// Inisialisasi slice 'days' dengan nama-nama hari dalam seminggu
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	
	// Membuat slice baru 'daySlice1' yang berisi elemen dari index 5 sampai akhir
	daySlice1 := days[5:]
	fmt.Println(daySlice1) // Output: [Sabtu Minggu]

	// Mengubah elemen pertama dan kedua dari 'daySlice1'
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // Output: [Sabtu Baru Minggu Baru]
	fmt.Println(days)      // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	// Menambahkan elemen baru ke 'daySlice1' dan menyimpannya di 'daySlice2'
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1) // Output: [Sabtu Lama Minggu Baru]
	fmt.Println(daySlice2) // Output: [Sabtu Lama Minggu Baru Libur Baru]
	fmt.Println(days)      // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
}