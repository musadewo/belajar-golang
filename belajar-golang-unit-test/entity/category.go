// Package entity mendefinisikan struktur data dan tipe yang digunakan dalam aplikasi
package entity

// Import yang mungkin digunakan (tidak terlihat dalam potongan kode ini):
// import (
//     "github.com/google/uuid" // Jika Id menggunakan UUID
// )

// Category merepresentasikan struktur data untuk kategori dalam aplikasi
type Category struct {
    Id   string // Identifier unik untuk kategori, bisa berupa UUID atau string lainnya
    Name string // Nama kategori
}

// Metode-metode yang mungkin ditambahkan (tidak ada dalam potongan kode saat ini):
// func (c *Category) Validate() error
// func (c *Category) BeforeCreate()
// func NewCategory(name string) *Category