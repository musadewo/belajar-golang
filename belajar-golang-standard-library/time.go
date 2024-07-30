package main

import (
	"fmt"  // Paket untuk format I/O, seperti fungsi Println
	"time" // Paket untuk bekerja dengan waktu dan tanggal
)

func main() {
	// Mendapatkan waktu saat ini
	var now time.Time = time.Now()
	// Mencetak waktu saat ini dalam zona waktu lokal
	fmt.Println(now.Local())

	// Membuat waktu tertentu dalam zona waktu UTC
	utc := time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)
	// Mencetak waktu tersebut dalam zona waktu UTC
	fmt.Println(utc)
	// Mencetak waktu tersebut dalam zona waktu lokal
	fmt.Println(utc.Local())

	// Format waktu yang akan digunakan untuk parsing dan formatting
	formatter := "2006-01-02 15:04:05"
	// Parsing string waktu dalam format RFC3339 menjadi objek waktu
	parse, _ := time.Parse(time.RFC3339, formatter)
	// Mencetak waktu yang telah di-parse dalam zona waktu lokal
	fmt.Println(parse.Local())

	// String waktu yang akan di-parse
	value := "2020-10-10 10:10:10"
	// Parsing string waktu dengan format yang telah ditentukan
	valueTime, err := time.Parse(formatter, value)
	// Jika terjadi error saat parsing, cetak error
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		// Jika tidak ada error, cetak waktu yang telah di-parse
		fmt.Println(valueTime)
	}
	// Mencetak komponen-komponen dari waktu yang telah di-parse
	fmt.Println(valueTime.Year())   // Tahun
	fmt.Println(valueTime.Month())  // Bulan
	fmt.Println(valueTime.Day())    // Hari
	fmt.Println(valueTime.Hour())   // Jam
	fmt.Println(valueTime.Minute()) // Menit
	fmt.Println(valueTime.Second()) // Detik
}