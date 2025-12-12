package repository

import "MechOpss/internal/src/models"

type Repository interface {
	Insert(model interface{}) error
	FindByEmail(model interface{}, email string) error
	FindByID(model interface{}, id string) error
	FindBy(model interface{}, id string, find string) error
	FindAll(models interface{}) error
	Save(model interface{}) error
	UpdateRefreshToken(model, userID interface{}, token string) error
	FindAdminByEmail(email string) (*models.Admin, error)
	Count(model interface{}) (int64, error)
	First(model interface{}) error
	LimitFind(model interface{}, limit int) error
	Delete(model interface{}, id string) error
	FindWithTwoPreload(model interface{}, first, second, id string) error
	FindWithPreload(model interface{}, Preload, id string) error
	FindAllBookingsWithStaff(model interface{}) error
	FindAllStaffsWithBookings(model interface{}) error
	FindStaffByIDWithBookings(model interface{}, id uint) error
}
