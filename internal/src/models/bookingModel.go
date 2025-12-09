package models

// RegisterBooking
type Booking struct {
	ID             string `json:"id" gorm:"primaryKey;type:varchar(50)"`
	CarModel       string `json:"car"`
	CarNumber      string `json:"carnumber"`
	FuelType       string `json:"fueltype"`
	Problem        string `json:"service"`
	Time           string `json:"time"`
	Date           string `json:"date"`
	Address        string `json:"address"`
	LandMark       string `json:"landmark"`
	UserID         uint   `json:"userid"`
	UserStatus     string `json:"userstatus"`
	Message        string `json:"message"`
	VisibleBooking bool   `gorm:"default:false"`
}
