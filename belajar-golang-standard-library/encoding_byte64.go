package main

import (
	"encoding/base64" // Mengimpor paket untuk encoding dan decoding base64
	"fmt"             // Mengimpor paket untuk format I/O
)

func main() {
	value := "Sadewo" // Mendefinisikan string yang akan diencode

	// Meng-encode string ke base64
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded) // Mencetak hasil encoding

	// Meng-decode string dari base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		// Menangani error jika terjadi kesalahan saat decoding
		fmt.Println("Error decoding: ", err.Error())
	}
	fmt.Println(string(decoded)) // Mencetak hasil decoding
}