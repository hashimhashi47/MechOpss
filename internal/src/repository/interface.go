package repository

import "MechOpss/internal/src/models"

type Repository interface {
	Insert(model interface{}) error
	FindByEmail(model interface{}, email string) error
	FindByID(model interface{}, id string) error
	FindAll(models interface{}) error
	Save(model interface{}) error
	UpdateRefreshToken(userID interface{}, token string) error
	FindAdminByEmail(email string) (*models.Admin, error)
	Count(model interface{}) (int64, error)
	First(model interface{}) error
	LimitFind(model interface{}, limit int) error
	Delete(model interface{}, id string) error
	FindBookingByID(model interface{}, id string) error
}
