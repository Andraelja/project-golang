package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"project-golang/database"
	"project-golang/handlers"
	"project-golang/repositories"
	"project-golang/services"
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

	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(*roleService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, roleRepo)
	userHandler := handlers.NewUserHandler(userService)

	// category
	http.HandleFunc("/api/category", categoryHandler.HandleCategory)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	// product
	http.HandleFunc("/api/product", productHandler.HandleProduct)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// role
	http.HandleFunc("/api/role", roleHandler.HandleRole)
	http.HandleFunc("/api/role/", roleHandler.HandleRoleByID)

	// user
	http.HandleFunc("/api/user", userHandler.HandleUser)
	http.HandleFunc("/api/user/", userHandler.HandleUserByID)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("gagal running server:", err)
	}
}
