package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils/constants"
	"errors"
	"strconv"
	"time"
)

// get all booked
func (s *AdminService) ServiceGetBooked() ([]models.Bookeds, error) {
	var booked []models.Bookeds
	if err := s.Repo.FindAllBookingsWithStaff(&booked); err != nil {
		return nil, errors.New("booked is empty")
	}
	return booked, nil
}

// assiagn staff
func (s *AdminService) AssignStaffService(bookingID string, staffIDstr string) (interface{}, error) {

	staffUint, err := strconv.ParseUint(staffIDstr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid staff_id")
	}
	staffID := uint(staffUint)

	var booked models.Bookeds
	if err := s.Repo.FindByID(&booked, bookingID); err != nil {
		return nil, errors.New("booking not found")
	}

	booked.StaffID = &staffID

	if err := s.Repo.Save(&booked); err != nil {
		return nil, errors.New("failed to assign staff")
	}

	var updated models.Bookeds
	if err := s.Repo.FindWithTwoPreload(&updated, constants.PRELOADSTAFF, constants.PRELOADSLOT, bookingID); err != nil {
		return nil, errors.New("failed to load staff data")
	}

	return updated, nil
}

// update the booked data
func (s *AdminService) ServiceUpadteBooked(id string, Input models.Bookeds) (interface{}, error) {

	var Booked models.Bookeds
	var Slot models.Slot

	if err := s.Repo.FindWithTwoPreload(&Booked, constants.PRELOADSTAFF, constants.PRELOADSLOT, id); err != nil {
		return nil, errors.New("booking not found")
	}

	//add to slot the same details what is updated
	if Booked.SlotID != nil {

		SlotID := strconv.Itoa(int(Booked.Slot.ID))
		if err := s.Repo.FindByID(&Slot, SlotID); err != nil {
			return nil, errors.New("failed to find the slot it")
		}

		if Input.Status != "" {
			Slot.Status = Input.Status
		}

		if Input.Description != "" {
			Slot.Description = Input.Description
		}

		if Input.StaffID != nil {
			Slot.StaffID = Input.StaffID
			Slot.StaffName = Booked.Staff.FirstName + " " + Booked.Staff.LastName
		}

		if err := s.Repo.Save(&Slot); err != nil {
			return nil, errors.New("failed to update")
		}
	}

	//updating the bookeds what are up need to be updated
	if Input.Status != "" {
		Booked.Status = Input.Status
	}

	if Input.Description != "" {
		Booked.Description = Input.Description
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

	if err := s.Repo.Save(&Booked); err != nil {
		return nil, errors.New("failed to update")
	}

	return Booked, nil
}

// add slot from booked to slots
func (s *AdminService) ServiceAddSlot(id string) (interface{}, error) {
	
	var Bookeds models.Bookeds
	if err := s.Repo.FindWithTwoPreload(&Bookeds, constants.PRELOADSTAFF, constants.PRELOADSLOT, id); err != nil {
		return models.Slot{}, errors.New("failed to find bookeds")
	}

	if Bookeds.StaffID == nil {
		return nil, errors.New("select the staff")
	}

	if Bookeds.SlotID != nil {
		return nil, errors.New("already on the slot")
	}

	var slots []models.Slot
	var Slot *models.Slot
	if err := s.Repo.FindAll(&slots); err != nil {
		return nil, errors.New("not avilable ")
	}
	//loop all the slot and check is empty or not
	for i := range slots {
		if slots[i].Status == "EMPTY" {
			Slot = &slots[i]
			break
		}
	}

	if Slot == nil {
		return nil, errors.New("no empty slots available")
	}
	//addd the details to the empty slot
	now := time.Now()
	Slot.CarModel = Bookeds.CarModel
	Slot.Time = &now
	Slot.CarNumber = Bookeds.CarNumber
	Slot.StaffID = Bookeds.StaffID
	Slot.StaffName = Bookeds.Staff.FirstName
	Slot.Status = Bookeds.Status

	// Save the updated slot
	if err := s.Repo.Save(Slot); err != nil {
		return nil, errors.New("failed to update slot")
	}

	Bookeds.SlotID = &Slot.ID
	if err := s.Repo.Save(&Bookeds); err != nil {
		return nil, errors.New("failed to update booked slot reference")
	}

	return Slot, nil
}
