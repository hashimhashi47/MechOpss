package repository

import (
	"MechOpss/internal/src/models"
	"gorm.io/gorm"
)

type SQLrepo struct {
	DB *gorm.DB
}

func Newrepo(db *gorm.DB) Repository {
	return &SQLrepo{DB: db}
}

// create
func (r *SQLrepo) Insert(model interface{}) error {
	return r.DB.Create(model).Error
}

// find email
func (r *SQLrepo) FindByEmail(model interface{}, email string) error {
	return r.DB.Where("email = ?", email).First(model).Error
}

// save
func (r *SQLrepo) Save(model interface{}) error {
	return r.DB.Save(model).Error
}

// update the database refersh token
func (r *SQLrepo) UpdateRefreshToken(userID interface{}, token string) error {
	return r.DB.Model(&models.User{}).Where("id = ?", userID).Update("refresh_token", token).Error
}

// find admin by email
func (r *SQLrepo) FindAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := r.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// count of all strcut
func (r *SQLrepo) Count(model interface{}) (int64, error) {
	var count int64
	err := r.DB.Model(model).Count(&count).Error
	return count, err
}

// take the first
func (r *SQLrepo) First(model interface{}) error {
	err := r.DB.First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// only find the data from the given struct limited by 5
func (r *SQLrepo) LimitFind(model interface{}, limit int) error {
	err := r.DB.Limit(limit).Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

// find all user
func (r *SQLrepo) FindAll(models interface{}) error {
	return r.DB.Find(models).Error
}

// find anthing by id
func (r *SQLrepo) FindByID(model interface{}, id string) error {
	return r.DB.Where("id = ?", id).First(model).Error
}

// delete permenet using id
func (r *SQLrepo) Delete(model interface{}, id string) error {
	return r.DB.Where("id = ?", id).Unscoped().Delete(model).Error
}

// get a specific booked with staff
func (r *SQLrepo) FindBookingWithStaffAndSlot(model interface{}, id string) error {
	return r.DB.Preload("Staff").Preload("Slot").Where("id = ?", id).First(model).Error
}

// get all booking with staff
func (r *SQLrepo) FindAllBookingsWithStaff(model interface{}) error {
	return r.DB.Preload("Staff").Find(model).Error
}

// get all staffs with their bookings preloaded
func (r *SQLrepo) FindAllStaffsWithBookings(model interface{}) error {
	return r.DB.Preload("Bookings").Find(model).Error
}

// get a single staff by ID with bookings
func (r *SQLrepo) FindStaffByIDWithBookings(model interface{}, id uint) error {
	return r.DB.Preload("Bookings").First(model, id).Error
}


