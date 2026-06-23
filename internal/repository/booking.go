package repository

import (
	"database/sql"
	"github.com/Lalka12235/simple-hotel.git/internal/model"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Create(b *model.Booking) (int, error) {
	query := `INSERT INTO bookings (id_client, id_room, id_room_categories, check_in_time, check_out_time) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id_booking`
	var id int
	err := r.db.QueryRow(query, b.IdClient, b.IdRoom, b.IdRoomCategories, b.CheckInTime, b.CheckOutTime).Scan(&id)
	return id, err
}

func (r *BookingRepository) GetByID(id int) (*model.Booking, error) {
	query := `SELECT id_booking, id_client, id_room, id_room_categories, check_in_time, check_out_time FROM bookings WHERE id_booking = $1`
	var b model.Booking
	err := r.db.QueryRow(query, id).Scan(&b.IdBooking, &b.IdClient, &b.IdRoom, &b.IdRoomCategories, &b.CheckInTime, &b.CheckOutTime)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookingRepository) Update(b *model.Booking) error {
	query := `UPDATE bookings SET id_client=$1, id_room=$2, id_room_categories=$3, check_in_time=$4, check_out_time=$5 WHERE id_booking=$6`
	_, err := r.db.Exec(query, b.IdClient, b.IdRoom, b.IdRoomCategories, b.CheckInTime, b.CheckOutTime, b.IdBooking)
	return err
}

func (r *BookingRepository) Delete(id int) error {
	query := `DELETE FROM bookings WHERE id_booking = $1`
	_, err := r.db.Exec(query, id)
	return err
}