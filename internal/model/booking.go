package model
import "time"
type Booking struct {
	IdBooking          int
	IdClient           int
	IdRoom             int 
	IdRoomCategories   int 
	CheckInTime        time.Time 
	CheckOutTime       time.Time
}