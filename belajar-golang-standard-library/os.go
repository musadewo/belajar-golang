package main

import (
	"fmt" // Mengimpor paket fmt untuk format I/O
	"os"  // Mengimpor paket os untuk berinteraksi dengan sistem operasi
)

func main() {
	args := os.Args // Mengambil argumen baris perintah yang diberikan saat menjalankan program
	for _, arg := range args { // Loop melalui setiap argumen
		fmt.Println(arg) // Mencetak setiap argumen ke output standar
	}

	hostname, err := os.Hostname() // Mendapatkan nama host dari sistem
	if err == nil { // Jika tidak ada error saat mendapatkan nama host
		fmt.Println(hostname) // Mencetak nama host ke output standar
	} else { // Jika ada error saat mendapatkan nama host
		fmt.Println("Error:", err) // Mencetak pesan error ke output standar
	}
}