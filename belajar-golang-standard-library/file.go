package main

import (
	"bufio" // Package bufio menyediakan buffered I/O. Ini digunakan untuk membaca file baris demi baris.
	"fmt"   // Package fmt mengimplementasikan format I/O dengan fungsi yang mirip dengan printf dan scanf di C.
	"io"    // Package io menyediakan fungsi dasar untuk I/O, termasuk io.EOF yang digunakan untuk mendeteksi akhir file.
	"log"   // Package log menyediakan logging sederhana. Digunakan untuk mencatat kesalahan yang terjadi.
	"os"    // Package os menyediakan antarmuka untuk fungsi OS, seperti membuka dan menutup file.
)

// createNewFile membuat file baru dengan nama dan pesan yang diberikan.
// Jika file sudah ada, isinya akan ditimpa.
func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644) // Membuka atau membuat file dengan izin menulis.
	if err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat membuka file.
	}
	defer file.Close() // Menutup file setelah fungsi selesai.
	file.WriteString(message) // Menulis pesan ke file.
	return nil // Mengembalikan nil jika tidak ada error.
}

// readFile membaca isi file dengan nama yang diberikan dan mengembalikan isinya sebagai string.
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644) // Membuka file dengan izin baca.
	if err != nil {
		return "", err // Mengembalikan error jika terjadi kesalahan saat membuka file.
	}
	defer file.Close() // Menutup file setelah fungsi selesai.
	reader := bufio.NewReader(file) // Membuat buffered reader untuk membaca file.
	var message string
	for {
		line, _, err := reader.ReadLine() // Membaca file baris demi baris.
		if err == io.EOF {
			break // Keluar dari loop jika mencapai akhir file.
		}
		message += string(line) + "\n" // Menambahkan baris yang dibaca ke variabel message.
	}
	return message, nil // Mengembalikan isi file sebagai string.
}

// addToFile menambahkan pesan ke file dengan nama yang diberikan.
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0644) // Membuka file dengan izin menulis dan menambahkan.
	if err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat membuka file.
	}
	defer file.Close() // Menutup file setelah fungsi selesai.
	file.WriteString(message) // Menambahkan pesan ke file.
	return nil // Mengembalikan nil jika tidak ada error.
}

func main() {
	createNewFile("sample.log", "ini file sample") // Membuat file baru dengan pesan awal.
	result, err := readFile("sample.log") // Membaca isi file.
	if err != nil {
		log.Fatal(err) // Mencatat dan menghentikan program jika terjadi kesalahan.
	}
	fmt.Println(result) // Mencetak isi file ke konsol.
	addToFile("sample.log", "\nTambah Teks Baru") // Menambahkan teks baru ke file.
}