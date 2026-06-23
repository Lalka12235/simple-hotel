package model
type Booking struct {
	IdBooking          int
	IdClient           int
	IdRoom             int 
	IdRoomCategories   int 
	CheckInTime        string 
	CheckOutTime       string
}