package services

import (
	"MechOpss/internal/src/models"
	"errors"
	"strconv"
)

func (s *AdminService) ServiceGetBooked() ([]models.Booked, error) {
	var Booked []models.Booked
	if err := s.Repo.FindAll(&Booked); err != nil {
		return nil, errors.New("booked is empty")
	}

	return Booked, nil
}

func (s *AdminService) AssignStaffService(bookingID string, staffIDstr string) (interface{}, error) {

	staffUint, err := strconv.ParseUint(staffIDstr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid staff_id")
	}
	staffID := uint(staffUint)

	var booked models.Booked
	if err := s.Repo.FindBookingByID(&booked, bookingID); err != nil {
		return nil, errors.New("booking not found")
	}

	booked.StaffID = &staffID

	if err := s.Repo.Save(&booked); err != nil {
		return nil, errors.New("failed to assign staff")
	}

	return booked, nil
}

func (s *AdminService) ServiceUpadteBooked(id string, Input models.Booked) (interface{}, error) {
	var Booked models.Booked

	if err := s.Repo.FindBookingByID(&Booked, id); err != nil {
		return nil, errors.New("booking not found")
	}
	if Input.Status != "" {
		Booked.Status = Input.Status
	}

	if Input.Problem != "" {
		Booked.Problem = Input.Problem
	}

	if Input.StaffID != nil {
		Booked.StaffID = Input.StaffID
	}

	if Input.PaymentAmount != 0 {
		Booked.PaymentAmount = Input.PaymentAmount
	}

	if Input.PaymentMode != "" {
		Booked.PaymentMode = Input.PaymentMode
	}
	if Input.ServiceStart != "" {
		Booked.ServiceStart = Input.ServiceStart
	}
	if Input.ServiceEnd != "" {
		Booked.ServiceEnd = Input.ServiceEnd
	}
	if Input.Delivery != "" {
		Booked.Delivery = Input.Delivery
	}

	if err := s.Repo.Save(Booked); err != nil {
		return nil, errors.New("failed to update")
	}

	return Booked, nil
}
