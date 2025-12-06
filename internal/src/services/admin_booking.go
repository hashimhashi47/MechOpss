package services

import (
	"MechOpss/internal/src/constants"
	"MechOpss/internal/src/models"
	"errors"
)

func (s *AdminService) ServiceGetAllBooking() ([]models.Booking, error) {
	var Booking []models.Booking

	if err := s.Repo.FindAll(&Booking); err != nil {
		return nil, errors.New("booking not found")
	}
	return Booking, nil
}

func (c *AdminService) ServiceApproveBooking(id string) (models.Booked, error) {
	var existingbooking models.Booking
	if err := c.Repo.FindByID(&existingbooking, id); err != nil {
		return models.Booked{}, err
	}

	booked := models.Booked{
		BookedID:  existingbooking.ID,
		UserID:    existingbooking.UserID,
		CarNumber: existingbooking.CarNumber,
		Problem:   existingbooking.Problem,
		Address:   existingbooking.Address,
		Date:      existingbooking.Date,
		Status:    constants.APPROVED,
		StaffID:   nil,
	}

	if err := c.Repo.Insert(&booked); err != nil {
		return models.Booked{}, errors.New("booking failed")
	}

	if err := c.Repo.Delete(existingbooking, id); err != nil {
		return models.Booked{}, errors.New("remove this service from booking failed")
	}

	return booked, nil
}
