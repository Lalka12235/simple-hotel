package repository

import (
	"database/sql"
	"github.com/Lalka12235/simple-hotel.git/internal/model"
)

type RoomCategoryRepository struct {
	db *sql.DB
}

func NewRoomCategoryRepository(db *sql.DB) *RoomCategoryRepository {
	return &RoomCategoryRepository{db: db}
}

func (r *RoomCategoryRepository) Create(cat *model.RoomCategories) (int, error) {
	query := `INSERT INTO room_categories (class_name) VALUES ($1) RETURNING id_room_categories`
	var id int
	err := r.db.QueryRow(query, cat.ClassName).Scan(&id)
	return id, err
}

func (r *RoomCategoryRepository) GetByID(id int) (*model.RoomCategories, error) {
	query := `SELECT id_room_categories, class_name FROM room_categories WHERE id_room_categories = $1`
	var cat model.RoomCategories
	err := r.db.QueryRow(query, id).Scan(&cat.IdRoomCategories, &cat.ClassName)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *RoomCategoryRepository) Update(cat *model.RoomCategories) error {
	query := `UPDATE room_categories SET class_name = $1 WHERE id_room_categories = $2`
	_, err := r.db.Exec(query, cat.ClassName, cat.IdRoomCategories)
	return err
}

func (r *RoomCategoryRepository) Delete(id int) error {
	query := `DELETE FROM room_categories WHERE id_room_categories = $1`
	_, err := r.db.Exec(query, id)
	return err
}