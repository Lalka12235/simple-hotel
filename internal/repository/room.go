package repository

import (
	"database/sql"
	"github.com/Lalka12235/simple-hotel.git/internal/model"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) Create(room *model.Room) (int, error) {
	query := `INSERT INTO rooms (id_room_categories, capacity, price) VALUES ($1, $2, $3) RETURNING id_room`
	var id int
	err := r.db.QueryRow(query, room.IdRoomCategories, room.Capacity, room.Price).Scan(&id)
	return id, err
}

func (r *RoomRepository) GetByID(id int) (*model.Room, error) {
	query := `SELECT id_room, id_room_categories, capacity, price FROM rooms WHERE id_room = $1`
	var room model.Room
	err := r.db.QueryRow(query, id).Scan(&room.IdRoom, &room.IdRoomCategories, &room.Capacity, &room.Price)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) Update(room *model.Room) error {
	query := `UPDATE rooms SET id_room_categories=$1, capacity=$2, price=$3 WHERE id_room=$4`
	_, err := r.db.Exec(query, room.IdRoomCategories, room.Capacity, room.Price, room.IdRoom)
	return err
}

func (r *RoomRepository) Delete(id int) error {
	query := `DELETE FROM rooms WHERE id_room = $1`
	_, err := r.db.Exec(query, id)
	return err
}