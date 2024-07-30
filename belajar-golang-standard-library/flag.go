package main

import (
	"flag" // Mengimpor paket flag untuk parsing argumen baris perintah
	"fmt"  // Mengimpor paket fmt untuk output format
)

func main() {
	// Mendefinisikan flag baris perintah untuk username, default "root", dengan deskripsi
	var username *string = flag.String("username", "root", "Put your database username")
	
	// Mendefinisikan flag baris perintah untuk password, default "root", dengan deskripsi
	var password *string = flag.String("password", "root", "Put your database password")
	
	// Mendefinisikan flag baris perintah untuk host, default "localhost", dengan deskripsi
	var host *string = flag.String("host", "localhost", "Put your database host")
	
	// Mendefinisikan flag baris perintah untuk port, default 0, dengan deskripsi
	var port *int = flag.Int("port", 0, "Put your database port")
	
	// Mem-parsing semua flag yang telah didefinisikan dari argumen baris perintah
	flag.Parse()

	// Mencetak nilai dari flag yang telah di-parsing
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
}