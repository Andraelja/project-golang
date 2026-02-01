package services

import (
	"errors"
	"task-session-1/models"
	"task-session-1/repositories"
)

// ProductService adalah struct yang menyimpan dependency untuk operasi produk.
// productRepo digunakan untuk mengakses data produk di database.
// categoryRepo digunakan untuk validasi kategori saat membuat produk.
type ProductService struct {
	productRepo  *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

// NewProductService adalah konstruktor untuk membuat instance ProductService.
// Fungsi ini menerima ProductRepository dan CategoryRepository sebagai parameter.
// Mengembalikan pointer ke ProductService yang siap digunakan.
func NewProductService(
	productRepo *repositories.ProductRepository,
	categoryRepo *repositories.CategoryRepository,
) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// GetAll mengambil semua data produk dari database.
// Fungsi ini memanggil method GetAll dari ProductRepository.
// Mengembalikan slice dari Product dan error jika ada.
func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

// Create membuat produk baru setelah melakukan validasi.
// Pertama, memeriksa apakah CategoryID tidak kosong (tidak 0).
// Kemudian, memeriksa apakah kategori dengan ID tersebut ada di database.
// Jika validasi lolos, maka produk disimpan ke database.
// Mengembalikan error jika validasi gagal atau penyimpanan gagal.
func (s *ProductService) Create(data *models.Product) error {
	// Validasi: CategoryID tidak boleh kosong.
	if data.CategoryID == 0 {
		return errors.New("category cannot empty!")
	}

	// Mengambil data kategori berdasarkan ID untuk memastikan kategori ada.
	category, err := s.categoryRepo.GetByID(data.CategoryID)
	if err != nil {
		return err
	}

	// Jika kategori tidak ditemukan (nil), kembalikan error.
	if category == nil {
		return errors.New("Category not found!")
	}

	// Jika semua validasi lolos, simpan produk ke database.
	return s.productRepo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	if id <= 0 {
		return nil, errors.New("invalid product id")
	}

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s *ProductService) Update(product *models.Product) error {
	if product.CategoryID == 0 {
		return errors.New("Category cannot empty!")
	}

	// Mengambil data kategori berdasarkan ID untuk memastikan kategori ada.
	category, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return err
	}

	// Jika kategori tidak ditemukan (nil), kembalikan error.
	if category == nil {
		return errors.New("Category not found!")
	}

	return s.productRepo.Update(product)
}

// Delete menghapus kategori berdasarkan ID.
// Fungsi ini memanggil method Delete dari CategoryRepository.
// Mengembalikan error jika delete gagal.
func (s *ProductService) Delete(id int) error {
	// Ambil product
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found")
	}

	// Hapus product
	return s.productRepo.Delete(id)
}
