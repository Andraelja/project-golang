package repositories

import (
	"database/sql"
	"errors"
	"task-session-1/models"
)

// CategoryRepository adalah struct yang menyimpan koneksi database untuk operasi kategori.
// db adalah pointer ke koneksi database PostgreSQL.
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository adalah konstruktor untuk membuat instance CategoryRepository.
// Menerima koneksi database sebagai parameter dan mengembalikan pointer ke CategoryRepository.
func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAll mengambil semua data kategori dari database.
// Query SQL mengambil semua kolom dari tabel category.
// Mengembalikan slice dari Category dan error jika ada.
func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	// Query SQL untuk mengambil semua kategori.
	query := "SELECT id, name, description FROM category"
	// Menjalankan query dan mendapatkan rows.
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Pastikan rows ditutup setelah selesai.
	defer rows.Close()

	// Membuat slice kosong untuk menyimpan hasil.
	category := make([]models.Category, 0)
	// Iterasi setiap row hasil query.
	for rows.Next() {
		var p models.Category
		// Scan data dari row ke struct Category.
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return nil, err
		}
		// Tambahkan kategori ke slice.
		category = append(category, p)
	}

	return category, nil
}

// Create menyisipkan kategori baru ke database.
// Menggunakan INSERT dengan RETURNING id untuk mendapatkan ID yang dihasilkan.
// Mengembalikan error jika penyisipan gagal.
func (repo *CategoryRepository) Create(category *models.Category) error {
	// Query INSERT untuk menyisipkan kategori baru.
	query := "INSERT INTO category (name, description) VALUES ($1, $2) RETURNING id"
	// Menjalankan query dan scan ID yang dihasilkan ke category.ID.
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	return err
}

// GetByID mengambil satu kategori berdasarkan ID.
// Jika kategori tidak ditemukan, mengembalikan nil tanpa error.
// Mengembalikan pointer ke Category dan error jika ada.
func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	// Query SELECT untuk mengambil kategori berdasarkan ID.
	query := "SELECT id, name, description FROM category WHERE id = $1"

	var p models.Category
	// Menjalankan query dan scan hasil ke struct Category.
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description)
	// Jika tidak ada row ditemukan, kembalikan nil.
	if err == sql.ErrNoRows {
		return nil, nil
	}
	// Jika ada error lain, kembalikan error.
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Update memperbarui data kategori berdasarkan ID.
// Menggunakan UPDATE dan memeriksa apakah ada baris yang terpengaruh.
// Jika tidak ada baris yang terpengaruh, berarti kategori tidak ditemukan.
// Mengembalikan error jika update gagal.
func (repo *CategoryRepository) Update(category *models.Category) error {
	// Query UPDATE untuk memperbarui kategori.
	query := "UPDATE category SET name = $1, description = $2 WHERE id = $3"
	// Menjalankan query dan mendapatkan result.
	result, err := repo.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	// Memeriksa jumlah baris yang terpengaruh.
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Jika tidak ada baris yang terpengaruh, kategori tidak ditemukan.
	if rows == 0 {
		return errors.New("Category not found!")
	}

	return nil
}

// Delete menghapus kategori berdasarkan ID.
// Menggunakan DELETE dan memeriksa apakah ada baris yang terpengaruh.
// Jika tidak ada baris yang terpengaruh, berarti kategori tidak ditemukan.
// Mengembalikan error jika delete gagal.
func (repo *CategoryRepository) Delete(id int) error {
	// Query DELETE untuk menghapus kategori.
	query := "DELETE FROM category WHERE id = $1"
	// Menjalankan query dan mendapatkan result.
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	// Memeriksa jumlah baris yang terpengaruh.
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Jika tidak ada baris yang terpengaruh, kategori tidak ditemukan.
	if rows == 0 {
		return errors.New("kategori tidak ditemukan")
	}

	return err
}
