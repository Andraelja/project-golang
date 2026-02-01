package handlers

import (
	"encoding/json"
	"net/http"
	"project-golang/services"
)

type RoleHandler struct {
	services *services.RoleService
}

func NewRoleHandler(service services.RoleService) *RoleHandler {
	return &RoleHandler{services: &service}
}

func (h *RoleHandler) HandleRole(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	category, err := h.services.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
