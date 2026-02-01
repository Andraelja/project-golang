# Penjelasan Detail Kode Proyek Task-Session-1

Proyek ini adalah aplikasi REST API sederhana yang dibangun menggunakan bahasa pemrograman Go (Golang). Aplikasi ini mengelola data kategori (category) dan produk (product) dengan menggunakan database PostgreSQL. Struktur proyek mengikuti pola arsitektur yang umum digunakan dalam pengembangan aplikasi web, yaitu dengan pemisahan layer: **Models**, **Repositories**, **Services**, **Handlers**, dan **Database**.

## Struktur Proyek

Berikut adalah struktur direktori dan file dalam proyek ini:

```
task-session-1/
├── go.mod                 # File modul Go yang mendefinisikan dependensi
├── go.sum                 # File checksum untuk dependensi
├── main.go                # File utama aplikasi (entry point)
├── task-session-1.exe     # File executable hasil kompilasi (untuk Windows)
├── database/
│   └── database.go        # Konfigurasi dan inisialisasi koneksi database
├── handlers/
│   ├── category_handler.go # Handler untuk endpoint kategori
│   └── product_handler.go  # Handler untuk endpoint produk
├── models/
│   ├── category.go        # Model data untuk kategori
│   └── product.go         # Model data untuk produk
├── repositories/
│   ├── category_repository.go # Repository untuk operasi database kategori
│   └── product_repository.go  # Repository untuk operasi database produk
└── services/
    ├── category_service.go    # Service layer untuk logika bisnis kategori
    └── product_service.go     # Service layer untuk logika bisnis produk
```

## Dependensi Utama

Berdasarkan file `go.mod`, proyek ini menggunakan dependensi berikut:

- **github.com/joho/godotenv v1.5.1**: Untuk memuat file `.env` yang berisi variabel lingkungan.
- **github.com/lib/pq v1.10.9**: Driver PostgreSQL untuk Go.
- **github.com/spf13/viper v1.21.0**: Untuk membaca konfigurasi dari variabel lingkungan atau file konfigurasi.

Dependensi lainnya adalah dependensi tidak langsung (indirect) yang diperlukan oleh library utama.

## Penjelasan Detail Setiap File

### 1. main.go

File ini adalah titik masuk (entry point) dari aplikasi. Fungsi utamanya adalah:

- **Memuat file .env**: Menggunakan `godotenv.Load()` untuk memuat variabel lingkungan dari file `.env`. Jika file tidak ditemukan, aplikasi akan menggunakan variabel lingkungan sistem.
- **Membaca konfigurasi**: Menggunakan `viper.AutomaticEnv()` untuk membaca variabel lingkungan seperti `PORT` dan `DB_CONN`.
- **Validasi konfigurasi**: Memastikan bahwa `PORT` dan `DB_CONN` tidak kosong. Jika kosong, aplikasi akan berhenti dengan pesan error.
- **Inisialisasi database**: Memanggil fungsi `database.InitDB()` untuk membuat koneksi ke database PostgreSQL.
- **Inisialisasi komponen**: Membuat instance dari repository, service, dan handler untuk kategori dan produk.
- **Routing**: Mendaftarkan endpoint HTTP menggunakan `http.HandleFunc()`:
  - `/api/category`: Untuk operasi CRUD kategori (GET untuk semua, POST untuk membuat baru).
  - `/api/category/{id}`: Untuk operasi berdasarkan ID (GET, PUT, DELETE).
  - `/api/product`: Untuk operasi produk (GET untuk semua, POST untuk membuat baru).
- **Menjalankan server**: Menggunakan `http.ListenAndServe()` untuk menjalankan server HTTP di alamat yang ditentukan.

### 2. database/database.go

File ini bertanggung jawab untuk inisialisasi koneksi database PostgreSQL.

- **Fungsi InitDB**: Menerima string koneksi database sebagai parameter.
- **Membuka koneksi**: Menggunakan `sql.Open("postgres", connectionString)` untuk membuka koneksi ke database.
- **Menguji koneksi**: Menggunakan `db.Ping()` untuk memastikan koneksi berhasil.
- **Konfigurasi pool koneksi**: Mengatur `SetMaxOpenConns(25)` dan `SetMaxIdleConns(5)` untuk mengoptimalkan penggunaan koneksi.
- **Logging**: Mencetak pesan "Database connected successfully" jika koneksi berhasil.

### 3. models/category.go

File ini mendefinisikan struktur data untuk kategori.

- **Struct Category**: Berisi field:
  - `ID` (int): ID unik kategori.
  - `Name` (string): Nama kategori.
  - `Description` (string): Deskripsi kategori.
- **Tag JSON**: Setiap field memiliki tag `json` untuk serialisasi/deserialisasi JSON.

### 4. models/product.go

File ini mendefinisikan struktur data untuk produk.

- **Struct Product**: Berisi field:
  - `ID` (int): ID unik produk.
  - `Name` (string): Nama produk.
  - `Price` (int): Harga produk.
  - `Stock` (int): Stok produk.
  - `CategoryID` (int): ID kategori yang terkait.
  - `Category` (*Category): Relasi ke model Category (opsional, ditandai dengan `omitempty`).
- **Tag JSON**: Semua field memiliki tag `json` untuk serialisasi JSON.

### 5. repositories/category_repository.go

File ini mengimplementasikan layer repository untuk operasi database kategori. Menggunakan pola Repository untuk mengabstraksi akses data.

- **Struct CategoryRepository**: Berisi field `db *sql.DB` untuk koneksi database.
- **Fungsi NewCategoryRepository**: Konstruktor untuk membuat instance repository.
- **Metode GetAll**: Mengambil semua kategori dari database dengan query `SELECT id, name, description FROM category`.
- **Metode Create**: Menyisipkan kategori baru dan mengembalikan ID yang dihasilkan.
- **Metode GetByID**: Mengambil kategori berdasarkan ID. Mengembalikan `nil` jika tidak ditemukan.
- **Metode Update**: Memperbarui data kategori berdasarkan ID.
- **Metode Delete**: Menghapus kategori berdasarkan ID. Mengembalikan error jika kategori tidak ditemukan.

### 6. repositories/product_repository.go

File ini mengimplementasikan layer repository untuk operasi database produk.

- **Struct ProductRepository**: Berisi field `db *sql.DB`.
- **Fungsi NewProductRepository**: Konstruktor.
- **Metode GetAll**: Mengambil semua produk dengan JOIN ke tabel category untuk mendapatkan data kategori terkait. Namun, ada kesalahan dalam query: `stock FROM product p WHERE id=$1` seharusnya `p.stock`.
- **Metode Create**: Menyisipkan produk baru.
- **Metode GetByID**: Mengambil produk berdasarkan ID. Ada kesalahan sintaks dalam query: `stock FROM product p WHERE id=$1` dan `Scan` tidak sesuai dengan field yang dipilih.

**Catatan**: Ada beberapa kesalahan dalam kode ini yang perlu diperbaiki, seperti query SQL yang tidak lengkap dan pemindaian field yang salah.

### 7. services/category_service.go

File ini mengimplementasikan layer service untuk logika bisnis kategori. Service layer bertindak sebagai perantara antara handler dan repository.

- **Struct CategoryService**: Berisi field `repo *repositories.CategoryRepository`.
- **Fungsi NewCategoryService**: Konstruktor.
- **Metode**: Semua metode (GetAll, Create, GetByID, Update, Delete) hanya meneruskan panggilan ke repository tanpa logika tambahan.

### 8. services/product_service.go

File ini mengimplementasikan layer service untuk logika bisnis produk.

- **Struct ProductService**: Berisi field `productRepo` dan `categoryRepo` untuk akses ke kedua repository.
- **Fungsi NewProductService**: Konstruktor yang menerima kedua repository.
- **Metode GetAll**: Meneruskan panggilan ke repository.
- **Metode Create**: Memvalidasi bahwa `CategoryID` tidak kosong, memeriksa apakah kategori ada dengan memanggil `categoryRepo.GetByID()`, dan baru kemudian menyisipkan produk.

### 9. handlers/category_handler.go

File ini menangani request HTTP untuk endpoint kategori.

- **Struct CategoryHandler**: Berisi field `service *services.CategoryService`.
- **Fungsi NewCategoryHandler**: Konstruktor.
- **Metode HandleCategory**: Menangani request ke `/api/category` berdasarkan metode HTTP (GET untuk GetAll, POST untuk Create).
- **Metode HandleCategoryByID**: Menangani request ke `/api/category/{id}` (GET, PUT, DELETE).
- **Metode lainnya**: GetAll, Create, GetByID, Update, Delete – masing-masing menangani logika spesifik, termasuk parsing JSON, validasi, dan response HTTP.

### 10. handlers/product_handler.go

File ini menangani request HTTP untuk endpoint produk.

- **Struct ProductHandler**: Berisi field `service *services.ProductService`.
- **Fungsi NewProductHandler**: Konstruktor.
- **Metode HandleProduct**: Menangani request ke `/api/product` (GET untuk GetAll, POST untuk Create).
- **Metode GetAll dan Create**: Mirip dengan handler kategori, menangani parsing JSON dan response.

## Cara Menjalankan Aplikasi

1. **Persiapan Database**: Pastikan PostgreSQL sudah terinstall dan buat database baru. Buat tabel `category` dan `product` sesuai dengan model.

2. **Konfigurasi Environment**: Buat file `.env` di root proyek dengan isi:
   ```
   PORT=8080
   DB_CONN=postgres://username:password@localhost/dbname?sslmode=disable
   ```

3. **Install Dependensi**: Jalankan `go mod tidy` untuk menginstall dependensi.

4. **Jalankan Aplikasi**: Gunakan `go run main.go` atau kompilasi dengan `go build` dan jalankan executable.

5. **Test Endpoint**: Gunakan tools seperti Postman atau curl untuk test API.

## Kesimpulan

Proyek ini mendemonstrasikan implementasi REST API sederhana dengan arsitektur layered (handler -> service -> repository -> database). Meskipun ada beberapa kesalahan kode yang perlu diperbaiki (terutama di `product_repository.go`), struktur keseluruhan sudah baik dan dapat dikembangkan lebih lanjut. Penggunaan Viper untuk konfigurasi dan godotenv untuk environment variable membuat aplikasi lebih fleksibel dan aman.
