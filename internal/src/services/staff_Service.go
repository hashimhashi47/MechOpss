package services

import (
	"MechOpss/internal/src/dto"
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

	if Staff.Block {
		return 0, nil, "", "", errors.New("blocked or not able to login this staff")
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

func (ss *StaffService) ServiceCheckStaffBookeds(id string) ([]dto.BookedsStaff, error) {
	var bookeds []models.Bookeds

	if err := ss.Repo.FindBy(&bookeds, id, "staff_id = ?"); err != nil {
		return nil, err
	}

	var data []dto.BookedsStaff

	for _, b := range bookeds {
		data = append(data, dto.BookedsStaff{
			ID:          b.ID,
			StaffID:     b.StaffID,
			CarModel:    b.CarModel,
			Description: b.Description,
			CarNumber:   b.CarNumber,
			Status:      b.Status,
			Address:     b.Address,
			Date:        b.Date,
			SlotID:      b.SlotID,
		})
	}

	return data, nil
}

func (ss *StaffService) ServiceUpdateStatus(status, description string, id string) (data interface{}, err error) {
	var bookeds models.Bookeds
	if err := ss.Repo.FindByID(&bookeds, id); err != nil {
		return nil, errors.New("failed to find the bookeds")
	}

	bookeds.Status = status
	bookeds.Description = description

	if err := ss.Repo.Save(&bookeds); err != nil {
		return nil, errors.New("failed update the bookeds")
	}

	return bookeds, nil
}

func (ss *StaffService) ServiceGetSlots(id string) ([]models.Slot, error) {
	var slots []models.Slot
	if err := ss.Repo.FindBy(&slots, id, "staff_id = ?"); err != nil {
		return nil, errors.New("failed to find slots")
	}

	return slots, nil
}
