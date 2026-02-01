package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"task-session-1/database"
	"task-session-1/handlers"
	"task-session-1/repositories"
	"task-session-1/services"
)

// Config adalah struct yang menyimpan konfigurasi aplikasi.
// Port adalah port tempat server akan berjalan.
// DBConn adalah string koneksi ke database PostgreSQL.
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

// main adalah fungsi utama yang dijalankan saat aplikasi dimulai.
// Fungsi ini melakukan inisialisasi, validasi konfigurasi, dan menjalankan server HTTP.
func main() {
	// Memuat file .env jika ada. Jika tidak ada, menggunakan environment variables sistem.
	// Ini memungkinkan konfigurasi fleksibel antara development dan production.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Mengaktifkan pembacaan otomatis dari environment variables menggunakan viper.
	viper.AutomaticEnv()

	// Membuat instance Config dan mengisi dengan nilai dari environment variables.
	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Validasi bahwa konfigurasi PORT dan DB_CONN tidak kosong.
	// Jika kosong, aplikasi akan berhenti karena tidak bisa berjalan tanpa konfigurasi ini.
	if config.Port == "" {
		log.Fatal("PORT is required")
	}
	if config.DBConn == "" {
		log.Fatal("DB_CONN is required")
	}

	// Inisialisasi koneksi database menggunakan fungsi InitDB dari package database.
	// Jika gagal, aplikasi akan berhenti karena tidak bisa mengakses database.
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	// Pastikan koneksi database ditutup saat aplikasi selesai.
	defer db.Close()

	// Membuat alamat server berdasarkan port yang dikonfigurasi.
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	// Inisialisasi komponen untuk kategori: repository, service, dan handler.
	// Repository mengakses database, service menangani logika bisnis, handler menangani HTTP requests.
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Inisialisasi komponen untuk produk: repository, service, dan handler.
	// Product service membutuhkan category repository juga untuk validasi.
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo, categoryRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Mendaftarkan handler untuk endpoint HTTP.
	// /api/category untuk operasi umum kategori (GET semua, POST buat baru).
	http.HandleFunc("/api/category", categoryHandler.HandleCategory)
	// /api/category/ untuk operasi berdasarkan ID (GET, PUT, DELETE).
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)
	// /api/product untuk operasi produk (GET semua, POST buat baru).
	http.HandleFunc("/api/product", productHandler.HandleProduct)
	// /api/product/ untuk operasi berdasarkan ID (GET, PUT, DELETE).
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// Menjalankan server HTTP di alamat yang ditentukan.
	// Jika gagal, aplikasi akan berhenti dengan pesan error.
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("gagal running server:", err)
	}
}
