package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-session-1/models"
	"task-session-1/services"
)

// CategoryHandler adalah struct yang menangani request HTTP untuk kategori.
// service adalah dependency ke CategoryService untuk logika bisnis.
type CategoryHandler struct {
	service *services.CategoryService
}

// NewCategoryHandler adalah konstruktor untuk membuat instance CategoryHandler.
// Menerima CategoryService sebagai parameter dan mengembalikan pointer ke CategoryHandler.
func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategory menangani request ke /api/category.
// Berdasarkan metode HTTP (GET untuk GetAll, POST untuk Create).
// Jika metode tidak didukung, kembalikan error 405 Method Not Allowed.
func (h *CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetAll menangani GET /api/category untuk mengambil semua kategori.
// Memanggil service.GetAll(), lalu encode hasil ke JSON dan kirim sebagai response.
// Jika ada error, kembalikan status 500 Internal Server Error.
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	category, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// Create menangani POST /api/category untuk membuat kategori baru.
// Decode JSON dari request body ke struct Category.
// Panggil service.Create(), lalu encode hasil ke JSON dan kirim sebagai response dengan status 201 Created.
// Jika ada error, kembalikan status 400 Bad Request.
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// HandleCategoryByID menangani request ke /api/category/{id}.
// Berdasarkan metode HTTP (GET untuk GetByID, PUT untuk Update, DELETE untuk Delete).
// Jika metode tidak didukung, kembalikan error 405 Method Not Allowed.
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetByID menangani GET /api/category/{id} untuk mengambil kategori berdasarkan ID.
// Ekstrak ID dari URL path, konversi ke int.
// Panggil service.GetByID(), lalu encode hasil ke JSON.
// Jika ID invalid, kembalikan status 400 Bad Request.
// Jika kategori tidak ditemukan, kembalikan status 404 Not Found.
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// Update menangani PUT /api/category/{id} untuk memperbarui kategori.
// Ekstrak ID dari URL, decode JSON dari body.
// Set ID ke category, panggil service.Update(), lalu encode hasil ke JSON.
// Jika ada error, kembalikan status 400 Bad Request.
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// Delete menangani DELETE /api/category/{id} untuk menghapus kategori.
// Ekstrak ID dari URL, panggil service.Delete().
// Jika berhasil, kirim response JSON dengan pesan sukses.
// Jika ada error, kembalikan status 500 Internal Server Error.
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "category deleted successfully",
	})
}
