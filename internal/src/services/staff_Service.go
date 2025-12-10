package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/utils"
	"errors"
)

type StaffService struct {
	Repo repository.Repository
}

func NewStaffService(repo repository.Repository) *StaffService {
	return &StaffService{Repo: repo}
}

// staff login
func (ss *StaffService) ServiceStaffLogin(email, password string) (staff_id uint, data interface{}, AccessToken, RefreshToken string, err error) {
	var Staff models.Staff

	if err := ss.Repo.FindByEmail(&Staff, email); err != nil {
		return 0, nil, "", "", errors.New("staff not found")
	}

	if err := utils.HashCompare(Staff.Password, password); err != nil {
		return 0, nil, "", "", errors.New("invalid password")
	}

	Accesstoken, err := utils.AccessToken(Staff.ID, Staff.Email, Staff.Role)
	if err != nil {
		return 0, nil, "", "", errors.New("failed to generate accesstoken")
	}

	Refreshtoken, err := utils.RefershToken(Staff.ID, Staff.Email, Staff.Role)
	if err != nil {
		return 0, nil, "", "", errors.New("failed to generate refershtoken")
	}

	Staff.RefreshToken = Refreshtoken
	ss.Repo.Save(&Staff)

	return Staff.ID, Staff, Accesstoken, Refreshtoken, nil
}

// updated the current details
func (ss *StaffService) ServiceStaffProfileUpdate() {

}

// logout
func (ss *StaffService) ServiceStaffLogout() {
	
}
