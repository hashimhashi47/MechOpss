package dto

type BookedsStaff struct {
	ID          string
	StaffID     *uint  `json:"staff_id"`
	CarModel    string `json:"car"`
	Description string `json:"description"`
	CarNumber   string `json:"car_number"`
	Status      string `json:"status"`
	Address     string `json:"address"`
	Date        string `json:"date"`
	SlotID      *uint  `json:"slot_id"`
}

type UserBooking struct {
	ID         string `json:"id"`
	CarModel   string `json:"car"`
	CarNumber  string `json:"carnumber"`
	UserStatus string `json:"userstatus"`
	Message    string `json:"message"`
}

type UserBookeds struct {
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
	CarNumber string `json:"car_number"`
	Status    string `json:"status"`
	Date      string `json:"date"`
}

type Userpayments struct {
	ID            string  `json:"id" gorm:"primaryKey;type:varchar(50)"`
	PaymentStatus string  `json:"payment_status"`
	PaymentAmount float64 `json:"payment_amount"`
	PaymentMode   string  `json:"payment_mode"`
}
