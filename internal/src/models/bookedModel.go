package models



// After user confirm service
type Booked struct {
	BookedID string `json:"id" gorm:"primaryKey;size:100"`

	UserID  uint  `json:"user_id"`
	StaffID *uint `json:"staff_id"` 

	CarNumber string `json:"car_number"`
	Status    string `json:"status"`
	Problem   string `json:"problem"`

	Address string `json:"address"`

	PaymentStatus string  `json:"payment_status"`
	PaymentAmount float64 `json:"payment_amount"`
	PaymentMode   string  `json:"payment_mode"`

	Date         string `json:"date"` // Booking date
	ServiceStart string `json:"service_start"`
	ServiceEnd   string `json:"service_end"`
	Delivery     string `json:"delivery"`

	Staff Staff `gorm:"foreignKey:StaffID"`
}
