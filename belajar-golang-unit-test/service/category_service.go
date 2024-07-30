// Package service berisi implementasi logika bisnis untuk kategori
package service

// Import yang diperlukan:
import (
	// entity: berisi definisi struktur data Category
	"belajar-golang-unit-test/entity"
	// repository: berisi interface untuk akses data kategori
	"belajar-golang-unit-test/repository"
	// errors: digunakan untuk membuat dan menangani error
	"errors"
)

// CategoryService adalah struct yang menyediakan layanan terkait kategori
type CategoryService struct {
	// Repository adalah dependency injection untuk akses data kategori
	Repository repository.CategoryRepository
}

// Get adalah method untuk mengambil data kategori berdasarkan ID
func (service CategoryService) Get(id string) (*entity.Category, error) {
	// Mencari kategori berdasarkan ID menggunakan repository
	category := service.Repository.FindById(id)
	
	// Jika kategori tidak ditemukan, kembalikan error
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		// Jika kategori ditemukan, kembalikan data kategori
		return category, nil
	}
}