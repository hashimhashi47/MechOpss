package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"errors"
)

func (c *UserService) ServiceBookingUser(data models.Booking) (string, error) {
	ID := utils.RandomIDGenerate("BOOK")
	data.ID = ID

	if err := c.repo.Insert(&data); err != nil {
		return "", errors.New("booking failed")
	}
	return ID, nil
}



