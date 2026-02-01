package services

import (
	"task-session-1/models"
	"task-session-1/repositories"
)

// CategoryService adalah struct yang menyimpan dependency untuk operasi kategori.
// repo adalah pointer ke CategoryRepository yang digunakan untuk mengakses database.
type CategoryService struct {
	repo *repositories.CategoryRepository
}

// NewCategoryService adalah konstruktor untuk membuat instance CategoryService.
// Fungsi ini menerima CategoryRepository sebagai parameter.
// Mengembalikan pointer ke CategoryService yang siap digunakan.
func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAll mengambil semua data kategori dari database.
// Fungsi ini memanggil method GetAll dari CategoryRepository.
// Mengembalikan slice dari Category dan error jika ada.
func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

// Create membuat kategori baru.
// Fungsi ini memanggil method Create dari CategoryRepository.
// Mengembalikan error jika penyimpanan gagal.
func (s *CategoryService) Create(data *models.Category) error {
	return s.repo.Create(data)
}

// GetByID mengambil satu kategori berdasarkan ID.
// Fungsi ini memanggil method GetByID dari CategoryRepository.
// Mengembalikan pointer ke Category dan error jika ada.
func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

// Update memperbarui data kategori.
// Fungsi ini memanggil method Update dari CategoryRepository.
// Mengembalikan error jika update gagal.
func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

// Delete menghapus kategori berdasarkan ID.
// Fungsi ini memanggil method Delete dari CategoryRepository.
// Mengembalikan error jika delete gagal.
func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
