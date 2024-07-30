package main

import (
	"encoding/csv" // Package untuk encoding dan decoding file CSV
	"os"           // Package untuk berinteraksi dengan sistem operasi, seperti membuka file
)

func main() {
	// Membuat writer CSV yang akan menulis ke output standar (biasanya terminal)
	writter := csv.NewWriter(os.Stdout)

	// Menulis baris pertama ke CSV
	_ = writter.Write([]string{"Muhammad", "Sadewo", "Wicaksono"})
	// Menulis baris kedua ke CSV
	_ = writter.Write([]string{"Joko", "Kurniawan", "Khannedy"})
	// Menulis baris ketiga ke CSV
	_ = writter.Write([]string{"Budi", "Joko", "Budi"})
	
	// Memastikan semua data yang ditulis dikirim ke output
	writter.Flush()
}