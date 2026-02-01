package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-session-1/models"
	"task-session-1/services"
)

// ProductHandler adalah struct yang menangani request HTTP untuk produk.
// service adalah dependency ke ProductService untuk logika bisnis.
type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler adalah konstruktor untuk membuat instance ProductHandler.
// Menerima ProductService sebagai parameter dan mengembalikan pointer ke ProductHandler.
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProduct menangani request ke /api/product.
// Berdasarkan metode HTTP (GET untuk GetAll, POST untuk Create).
// Jika metode tidak didukung, kembalikan error 405 Method Not Allowed.
func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
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

// GetAll menangani GET /api/product untuk mengambil semua produk.
// Memanggil service.GetAll(), lalu encode hasil ke JSON dan kirim sebagai response.
// Jika ada error, kembalikan status 500 Internal Server Error.
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	product, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Create menangani POST /api/product untuk membuat produk baru.
// Decode JSON dari request body ke struct Product.
// Panggil service.Create(), lalu encode hasil ke JSON dan kirim sebagai response dengan status 201 Created.
// Jika ada error, kembalikan status 400 Bad Request atau 500 Internal Server Error.
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body!", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
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

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Ekstrak ID dari URL, decode JSON dari body.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "product deleted successfully",
	})
}
