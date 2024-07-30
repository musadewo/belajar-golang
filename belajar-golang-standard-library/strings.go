package main

import (
	"fmt"     // Mengimpor paket fmt untuk fungsi format I/O
	"strings" // Mengimpor paket strings untuk manipulasi string
)

func main() {
	// Memeriksa apakah substring "Sadewo" ada dalam string "Muhammad Sadewo"
	fmt.Println(strings.Contains("Muhammad Sadewo", "Sadewo"))

	// Memisahkan string "Muhammad Sadewo" menjadi slice berdasarkan spasi
	fmt.Println(strings.Split("Muhammad Sadewo", " "))

	// Mengubah semua karakter dalam string "Muhammad Sadewo" menjadi huruf kecil
	fmt.Println(strings.ToLower("Muhammad Sadewo"))

	// Mengubah semua karakter dalam string "Muhammad Sadewo" menjadi huruf besar
	fmt.Println(strings.ToUpper("Muhammad Sadewo"))

	// Menghapus karakter spasi di awal dan akhir string "   Muhammad Sadewo   "
	fmt.Println(strings.Trim("   Muhammad Sadewo   ", " "))

	// Menghapus semua spasi di awal dan akhir string "   Muhammad Sadewo   "
	fmt.Println(strings.TrimSpace("   Muhammad Sadewo   "))

	// Mengganti substring "Sadewo" dengan "Sadewo" dalam string "Muhammad Sadewo" (tidak ada perubahan karena substring sama)
	fmt.Println(strings.Replace("Muhammad Sadewo", "Sadewo", "Sadewo", -1))

	// Mengganti semua kemunculan substring "Sadewo" dengan "Sadewo" dalam string "Muhammad Sadewo" (tidak ada perubahan karena substring sama)
	fmt.Println(strings.ReplaceAll("Muhammad Sadewo", "Sadewo", "Sadewo"))
}