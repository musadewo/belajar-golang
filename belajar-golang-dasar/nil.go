package main

import "fmt"

// NewMap membuat peta baru dengan satu entri jika nama diberikan, atau mengembalikan nil jika nama kosong
func NewMap(name string) map[string]string {
	if name == "" {
		// Jika nama kosong, kembalikan nil
		return nil
	} else {
		// Jika nama tidak kosong, kembalikan peta dengan entri "name"
		return map[string]string{
			"name": name, // Entri peta dengan kunci "name" dan nilai dari parameter name
		}
	}
}

func main() {
	// Panggil fungsi NewMap dengan argumen "Dewo"
	data := NewMap("Dewo")
	if data == nil {
		// Jika data nil, cetak "Data is nil"
		fmt.Println("Data is nil")
	} else {
		// Jika data tidak nil, cetak nilai dari entri "name"
		fmt.Println(data["name"]) // Mengakses nilai dari peta dengan kunci "name"
	}
}