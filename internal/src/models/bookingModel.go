package models

// RegisterBooking
type Booking struct {
	ID        string `json:"id" gorm:"primaryKey"`
	CarModel  string `json:"car"`
	CarNumber string `json:"carnumber"`
	FuelType  string `json:"fueltype"`
	Problem   string `json:"service"`
	Time      string `json:"time"`
	Date      string `json:"date"`
	Address   string `json:"address"`
	LandMark  string `json:"landmark" `
	UserID    uint   `json:"userid"`
}
