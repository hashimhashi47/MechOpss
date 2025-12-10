package models

// After user confirm service
type Bookeds struct {
	ID            string  `json:"id" gorm:"primaryKey;type:varchar(50)"`
	CarModel      string  `json:"car"`
	ServiceStart  string  `json:"service_start"`
	ServiceEnd    string  `json:"service_end"`
	Delivery      string  `json:"delivery"`
	PaymentStatus string  `json:"payment_status"`
	PaymentAmount float64 `json:"payment_amount"`
	PaymentMode   string  `json:"payment_mode"`
	Description   string  `json:"description"`
	Message       string  `json:"message"`

	UserID    uint   `json:"user_id"`
	StaffID   *uint  `json:"staff_id"`
	CarNumber string `json:"car_number"`
	Status    string `json:"status"`
	Address   string `json:"address"`
	Date      string `json:"date"`

	SlotID *uint `json:"slot_id"`

	Slot  Slot  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Staff Staff `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
