// Package repository berisi implementasi mock untuk repository kategori
package repository

import (
	// Import entity dari package belajar-golang-unit-test
	"belajar-golang-unit-test/entity"

	// Import mock dari package github.com/stretchr/testify/mock
	// Digunakan untuk membuat objek mock dalam pengujian
	"github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock adalah struct yang mengimplementasikan mock untuk CategoryRepository
type CategoryRepositoryMock struct {
	// Mock adalah objek dari testify/mock yang digunakan untuk menyimulasikan perilaku
	Mock mock.Mock
}

// FindById adalah metode mock untuk mencari kategori berdasarkan ID
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	// Called() digunakan untuk merekam pemanggilan metode dengan parameter tertentu
	arguments := repository.Mock.Called(id)
	
	// Memeriksa apakah nilai yang dikembalikan adalah nil
	if arguments.Get(0) == nil {
		return nil
	} else {
		// Jika tidak nil, konversi nilai yang dikembalikan ke tipe entity.Category
		category := arguments.Get(0).(entity.Category)
		// Mengembalikan pointer ke objek Category
		return &category
	}
}