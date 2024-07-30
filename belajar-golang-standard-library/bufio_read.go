package main

import (
	"bufio"   // Mengimpor paket bufio untuk buffered I/O
	"fmt"     // Mengimpor paket fmt untuk format I/O
	"io"      // Mengimpor paket io untuk I/O dasar
	"strings" // Mengimpor paket strings untuk manipulasi string
)

func main() {
	// Membuat pembaca string baru dengan input "Hello World"
	input := strings.NewReader("Hello World")
	
	// Membuat buffered reader baru dari input
	reader := bufio.NewReader(input)

	// Loop untuk membaca setiap baris dari reader
	for {
		// Membaca satu baris dari reader
		line, _, err := reader.ReadLine()
		
		// Jika mencapai akhir file (EOF), keluar dari loop
		if err == io.EOF {
			break
		}
		
		// Mencetak baris yang dibaca sebagai string
		fmt.Println(string(line))
	}
}