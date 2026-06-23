package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lalka12235/simple-hotel.git/internal/model"
	"github.com/Lalka12235/simple-hotel.git/internal/repository"
)

type BookingHandler struct {
	repo *repository.BookingRepository
}

func NewBookingHandler(repo *repository.BookingRepository) *BookingHandler {
	return &BookingHandler{repo: repo}
}

func (h *BookingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/api/bookings")
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
		http.Error(w, `{"error": "Неверный ID бронирования"}`, http.StatusBadRequest)
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

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var b model.Booking
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(&b)
	if err != nil {
		http.Error(w, `{"error": "Не удалось создать бронирование"}`, http.StatusInternalServerError)
		return
	}
	b.IdBooking = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func (h *BookingHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	b, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, `{"error": "Бронирование не найдено"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(b)
}

func (h *BookingHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var b model.Booking
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, `{"error": "Невалидный JSON"}`, http.StatusBadRequest)
		return
	}
	b.IdBooking = id
	if err := h.repo.Update(&b); err != nil {
		http.Error(w, `{"error": "Не удалось обновить бронирование"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *BookingHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, `{"error": "Не удалось удалить бронирование"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}