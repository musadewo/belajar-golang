package main

import "fmt"

// validateError adalah struct yang digunakan untuk kesalahan validasi
type validateError struct {
	Message string // Pesan kesalahan
}

// Error mengimplementasikan interface error untuk validateError
func (v *validateError) Error() string {
	return v.Message // Mengembalikan pesan kesalahan
}

// notFoundError adalah struct yang digunakan untuk kesalahan data tidak ditemukan
type notFoundError struct {
	Message string // Pesan kesalahan
}

// Error mengimplementasikan interface error untuk notFoundError
func (n *notFoundError) Error() string {
	return n.Message // Mengembalikan pesan kesalahan
}

// SaveData menyimpan data berdasarkan id dan mengembalikan error jika ada masalah
func SaveData(id string, data any) error {
	if id == "" {
		// Mengembalikan validateError jika id kosong
		return &validateError{Message: "Id tidak boleh kosong"}
	}
	if id != "Dewo" {
		// Mengembalikan notFoundError jika id tidak sama dengan "Dewo"
		return &notFoundError{Message: "Data tidak ditemukan"}
	}
	// Mengembalikan nil jika tidak ada error
	return nil
}

func main() {
	// Memanggil SaveData dengan id "Dewo" dan data nil
	err := SaveData("Dewo", nil)
	if err != nil {
		// Memeriksa jenis error dan mencetak pesan yang sesuai
		if validateError, ok := err.(*validateError); ok {
			// Jika error adalah validateError, cetak pesan kesalahan validasi
			fmt.Println("Validate Error: ", validateError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			// Jika error adalah notFoundError, cetak pesan kesalahan data tidak ditemukan
			fmt.Println("Not Found Error: ", notFoundError.Error())
		} else {
			// Jika error tidak diketahui, cetak pesan error tidak diketahui
			fmt.Println("Error tidak diketauhi", err.Error())
		}

		// Menggunakan switch untuk memeriksa jenis error dan mencetak pesan yang sesuai
		switch finalError := err.(type) {
		case *validateError:
			// Jika error adalah validateError, cetak pesan kesalahan validasi
			fmt.Println("Validate Error: ", finalError.Error())
		case *notFoundError:
			// Jika error adalah notFoundError, cetak pesan kesalahan data tidak ditemukan
			fmt.Println("Not Found Error: ", finalError.Error())
		default:
			// Jika error tidak diketahui, cetak pesan error tidak diketahui
			fmt.Println("Error tidak diketauhi", finalError.Error())
		}
	} else {
		// Mencetak pesan sukses jika tidak ada error
		fmt.Println("Sukses")
	}
}