package main

import (
	"bufio" // Package bufio menyediakan buffered I/O
	"os"    // Package os menyediakan antarmuka untuk fungsi OS
)

func main() {
	// Membuat writer baru yang akan menulis ke os.Stdout (standar output)
	writer := bufio.NewWriter(os.Stdout)
	
	// Menulis string "Hello World1\n" ke buffer
	writer.WriteString("Hello World1\n")
	// Menulis string "Hello World2\n" ke buffer
	writer.WriteString("Hello World2\n")
	// Menulis string "Hello World3\n" ke buffer
	writer.WriteString("Hello World3\n")
	// Menulis string "Hello World4\n" ke buffer
	writer.WriteString("Hello World4\n")
	
	// Memastikan semua data dalam buffer ditulis ke os.Stdout
	writer.Flush()
}