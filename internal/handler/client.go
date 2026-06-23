package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lalka12235/simple-hotel.git/internal/model"
	"github.com/Lalka12235/simple-hotel.git/internal/repository"
)

type ClientHandler struct {
	repo *repository.ClientRepository
}

func NewClientHandler(repo *repository.ClientRepository) *ClientHandler {
	return &ClientHandler{repo: repo}
}

// Главный роутер для /api/clients и /api/clients/
func (h *ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Простейший парсинг ID из URL вида /api/clients/5
	path := strings.TrimPrefix(r.URL.Path, "/api/clients")
	path = strings.TrimPrefix(path, "/")

	if path == "" {
		// Маршруты без ID: /api/clients
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		}
		return
	}

	// Маршруты с ID: /api/clients/{id}
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, `{"error": "Неверный ID клиента"}`, http.StatusBadRequest)
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

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c model.Client
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}

	id, err := h.repo.Create(&c)
	if err != nil {
		http.Error(w, `{"error": "Не удалось создать клиента"}`, http.StatusInternalServerError)
		return
	}

	c.IdClient = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *ClientHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	client, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, `{"error": "Клиент не найден"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(client)
}

func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var c model.Client
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}

	c.IdClient = id // Гарантируем, что ID берется из URL
	if err := h.repo.Update(&c); err != nil {
		http.Error(w, `{"error": "Не удалось обновить данные"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, `{"error": "Не удалось удалить клиента"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}