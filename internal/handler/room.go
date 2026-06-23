package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lalka12235/simple-hotel.git/internal/model"
	"github.com/Lalka12235/simple-hotel.git/internal/repository"
)

type RoomHandler struct {
	repo *repository.RoomRepository
}

func NewRoomHandler(repo *repository.RoomRepository) *RoomHandler {
	return &RoomHandler{repo: repo}
}

func (h *RoomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/api/rooms")
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
		http.Error(w, `{"error": "Неверный ID комнаты"}`, http.StatusBadRequest)
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

func (h *RoomHandler) Create(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(&room)
	if err != nil {
		http.Error(w, `{"error": "Не удалось создать комнату"}`, http.StatusInternalServerError)
		return
	}
	room.IdRoom = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	room, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, `{"error": "Комната не найдена"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var room model.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}
	room.IdRoom = id
	if err := h.repo.Update(&room); err != nil {
		http.Error(w, `{"error": "Не удалось обновить комнату"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *RoomHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, `{"error": "Не удалось удалить комнату"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}