package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lalka12235/simple-hotel.git/internal/model"
	"github.com/Lalka12235/simple-hotel.git/internal/repository"
)

type RoomCategoryHandler struct {
	repo *repository.RoomCategoryRepository
}

func NewRoomCategoryHandler(repo *repository.RoomCategoryRepository) *RoomCategoryHandler {
	return &RoomCategoryHandler{repo: repo}
}

func (h *RoomCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/api/categories")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		if r.Method == http.MethodPost {
			h.Create(w, r)
			return
		}
		http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, `{"error": "Неверный ID категории"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r, id)
	case http.MethodPut:
		h.Update(w, r, id)
	case http.MethodDelete:
		h.Delete(w, r, id)
	default:
		http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
	}
}

func (h *RoomCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cat model.RoomCategories
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(&cat)
	if err != nil {
		http.Error(w, `{"error": "Не удалось создать категорию"}`, http.StatusInternalServerError)
		return
	}
	cat.IdRoomCategories = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cat)
}

func (h *RoomCategoryHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	cat, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, `{"error": "Категория не найдена"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cat)
}

func (h *RoomCategoryHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var cat model.RoomCategories
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}
	cat.IdRoomCategories = id
	if err := h.repo.Update(&cat); err != nil {
		http.Error(w, `{"error": "Не удалось обновить категорию"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *RoomCategoryHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, `{"error": "Не удалось удалить категорию"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}