package main

import "fmt"

func main() {
	// Deklarasi array dengan 3 elemen string
	var names [3]string
	// Mengisi elemen-elemen array
	names[1] = "Agus"       // Mengisi elemen kedua dengan "Agus"
	names[2] = "Kurniawan"  // Mengisi elemen ketiga dengan "Kurniawan"
	names[0] = "Khannedy"   // Mengisi elemen pertama dengan "Khannedy"

	// Mencetak elemen-elemen array
	fmt.Println(names[0], names[1], names[2]) // Mencetak semua elemen array dalam satu baris
	// fmt.Println(names[1]) // Mencetak elemen kedua (dalam komentar)
	// fmt.Println(names[2]) // Mencetak elemen ketiga (dalam komentar)

	// Deklarasi dan inisialisasi array dengan 3 elemen integer
	var values = [3]int{
		90, // Elemen pertama diisi dengan 90
		80, // Elemen kedua diisi dengan 80
		70, // Elemen ketiga diisi dengan 70
	}

	// Mencetak seluruh elemen array
	fmt.Println(values) // Mencetak semua elemen array dalam satu baris
	// Mencetak panjang array
	fmt.Println("panjang array : ", len(values)) // Mencetak panjang array (jumlah elemen)
	// Mencetak kapasitas array
	fmt.Println("kapasitas array : ", cap(values)) // Mencetak kapasitas array (jumlah elemen yang dapat ditampung)
	// Mencetak slice dari array (elemen ke-0 sampai ke-1)
	fmt.Println(values[0:2]) // Mencetak elemen pertama dan kedua dari array

	// Mengubah nilai elemen pertama array
	values[0] = 100 // Mengubah nilai elemen pertama menjadi 100
	// Mencetak seluruh elemen array setelah perubahan
	fmt.Println(values) // Mencetak semua elemen array setelah perubahan
}