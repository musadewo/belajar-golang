package main

import "fmt"

func main() {
    // Mendeklarasikan variabel ujian dengan nilai 80
    var ujian = 80
    // Mendeklarasikan variabel absen dengan nilai 60
    var absen = 60

    // Mengecek apakah nilai ujian >= 80
    var lulusUjian = ujian >= 80
    // Mengecek apakah nilai absen >= 80
    var lulusAbsen = absen >= 80

    // Mengecek apakah lulus ujian dan lulus absen
    var lulus = lulusUjian && lulusAbsen
    // Mencetak hasil kelulusan
    fmt.Println(lulus) // Output: false karena lulusUjian true dan lulusAbsen false
}