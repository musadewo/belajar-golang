// Package service berisi implementasi layanan untuk kategori
package service

import (
	// entity: Berisi definisi struktur data untuk kategori
	"belajar-golang-unit-test/entity"
	// repository: Berisi interface dan mock untuk akses data kategori
	"belajar-golang-unit-test/repository"
	// testing: Package standar Go untuk unit testing
	"testing"

	// assert: Dari package testify, menyediakan fungsi-fungsi assertion untuk testing
	"github.com/stretchr/testify/assert"
	// mock: Dari package testify, menyediakan utilitas untuk membuat objek mock
	"github.com/stretchr/testify/mock"
)

// Membuat objek mock untuk CategoryRepository
var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}

// Membuat instance CategoryService dengan menggunakan mock repository
var categoryService = CategoryService{
	Repository: categoryRepository,
}

// TestCategoryService_Get menguji skenario ketika kategori tidak ditemukan
func TestCategoryService_Get(t *testing.T) {
	// Menyiapkan ekspektasi bahwa FindById dengan parameter "1" akan mengembalikan nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// Memanggil method Get pada categoryService
	category, err := categoryService.Get("1")

	// Memastikan bahwa kategori yang dikembalikan adalah nil
	assert.Nil(t, category)
	// Memastikan bahwa error tidak nil (ada error)
	assert.NotNil(t, err)
	// Memastikan pesan error sesuai dengan yang diharapkan
	assert.EqualError(t, err, "Category Not Found")
}

// TestCategoryService_GetSuccess menguji skenario ketika kategori berhasil ditemukan
func TestCategoryService_GetSuccess(t *testing.T) {
	// Membuat objek kategori untuk pengujian
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	// Menyiapkan ekspektasi bahwa FindById dengan parameter "2" akan mengembalikan kategori yang telah dibuat
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// Memanggil method Get pada categoryService
	result, err := categoryService.Get("2")

	// Memastikan bahwa tidak ada error
	assert.Nil(t, err)
	// Memastikan bahwa hasil tidak nil
	assert.NotNil(t, result)
	// Memastikan bahwa ID kategori yang dikembalikan sesuai
	assert.Equal(t, category.Id, result.Id)
	// Memastikan bahwa nama kategori yang dikembalikan sesuai
	assert.Equal(t, category.Name, result.Name)
}