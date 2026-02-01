package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Membuka koneksi ke database PostgreSQL menggunakan driver "postgres".
	// sql.Open tidak langsung membuat koneksi, hanya mempersiapkan.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Menguji koneksi dengan Ping untuk memastikan database dapat diakses.
	// Ini akan membuat koneksi fisik jika belum ada.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Mengatur pengaturan connection pool (opsional tapi direkomendasikan).
	// SetMaxOpenConns membatasi jumlah koneksi maksimal yang terbuka.
	db.SetMaxOpenConns(25)
	// SetMaxIdleConns membatasi jumlah koneksi idle (tidak digunakan) yang disimpan.
	db.SetMaxIdleConns(5)

	// Mencetak pesan sukses ke log.
	log.Println("Database connected successfully")
	return db, nil
}
