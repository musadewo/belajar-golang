// Package repository mendefinisikan interface dan implementasi untuk akses data kategori
package repository

// Mengimpor package entity dari proyek "belajar-golang-unit-test"
// Package entity kemungkinan berisi definisi struct Category
import "belajar-golang-unit-test/entity"

// CategoryRepository adalah interface yang mendefinisikan metode untuk mengakses data kategori
// Interface ini memungkinkan abstraksi dari implementasi konkret, memudahkan pengujian dan penggantian implementasi
type CategoryRepository interface {
	// FindById adalah metode untuk mencari kategori berdasarkan ID
	// Parameter:
	//   - id: string yang merepresentasikan ID kategori yang dicari
	// Return:
	//   - *entity.Category: pointer ke struct Category jika ditemukan, atau nil jika tidak ditemukan
	FindById(id string) *entity.Category
}