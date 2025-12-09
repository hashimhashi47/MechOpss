package services

import (
	"MechOpss/internal/src/constants"
	"MechOpss/internal/src/models"
	"errors"
	"strconv"
	"time"
)


//get all bookings 
func (s *AdminService) ServiceGetAllBooking() ([]models.Booking, error) {
	var Booking []models.Booking

	if err := s.Repo.FindAll(&Booking); err != nil {
		return nil, errors.New("booking not found")
	}

	var visibleBookings []models.Booking
	for i := range Booking {
		if !Booking[i].VisibleBooking {
			visibleBookings = append(visibleBookings, Booking[i])
		}
	}
	return visibleBookings, nil
}

//approve the bookings 
func (c *AdminService) ServiceApproveBooking(id string) (models.Bookeds, error) {
	var existingbooking models.Booking
	if err := c.Repo.FindByID(&existingbooking, id); err != nil {
		return models.Bookeds{}, err
	}

	booked := models.Bookeds{
		ID:          existingbooking.ID,
		UserID:      existingbooking.UserID,
		CarNumber:   existingbooking.CarNumber,
		Description: existingbooking.Problem,
		Address:     existingbooking.Address,
		Date:        existingbooking.Date,
		Status:      constants.APPROVED,
		StaffID:     nil,
		Message:     constants.APPROVEDMSG,
	}

	if err := c.Repo.Insert(&booked); err != nil {
		return models.Bookeds{}, errors.New("booking failed")
	}

	existingbooking.VisibleBooking = true
	existingbooking.UserStatus = constants.ACCEPTED
	existingbooking.Message = constants.ACCEPTORASSIGNEDBOOKING

	if err := c.Repo.Save(&existingbooking); err != nil {
		return models.Bookeds{}, errors.New("unavilable to remove from booked")
	}
	return booked, nil
}


//assign the bookings to booked and slots
func (c *AdminService) ServiceAssignBooking(bookingID string, staffIDstr string) (interface{}, error) {

	var existingbooking models.Booking
	if err := c.Repo.FindByID(&existingbooking, bookingID); err != nil {
		return models.Bookeds{}, err
	}

	staffUint, err := strconv.ParseUint(staffIDstr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid staff_id")
	}
	staffID := uint(staffUint)

	//first added to the slot 
	Slot := models.Slot{
		CarModel:    existingbooking.CarModel,
		CarNumber:   existingbooking.CarNumber,
		Time:        time.Now(),
		Description: existingbooking.Problem,
		StaffID:     &staffID,
		Status:      constants.ASSIGNEDSLOT,
	}

	if err := c.Repo.Insert(&Slot); err != nil {
		return nil, errors.New("failed to add on salot")
	}

	//then it added to the booked
	booked := models.Bookeds{
		ID:          existingbooking.ID,
		CarModel:    existingbooking.CarModel,
		UserID:      existingbooking.UserID,
		CarNumber:   existingbooking.CarNumber,
		Description: existingbooking.Problem,
		Address:     existingbooking.Address,
		Date:        existingbooking.Date,
		Status:      constants.ASSIGNEDSLOT,
		StaffID:     &staffID,
		Message:     constants.ASSIGNEDSLOTMSG,
		SlotID:      &Slot.ID,
	}


	if err := c.Repo.Insert(&booked); err != nil {
		return models.Bookeds{}, errors.New("booking failed")
	}

	//remove the visibility of booking from the interface and keep the data ok database for users.
	existingbooking.VisibleBooking = true
	existingbooking.UserStatus = constants.APPROVED
	existingbooking.Message = constants.ACCEPTORASSIGNEDBOOKING

	if err := c.Repo.Save(&existingbooking); err != nil {
		return nil, errors.New("failed to store booking")
	}

	var updated models.Bookeds
	if err := c.Repo.FindBookingWithStaffAndSlot(&updated, bookingID); err != nil {
		return nil, errors.New("failed to load staff data")
	}

	return updated, nil
}


//reject the booking it will remov ethe visibility stiil on the database
func (c *AdminService) ServiceRejectBooking(id string) (interface{}, error) {
	var booking models.Booking
	if err := c.Repo.FindByID(&booking, id); err != nil {
		return models.Booking{}, errors.New("failed to find the booking")
	}

	booking.VisibleBooking = true
	booking.UserStatus = constants.REJECTEDMSG
	if err := c.Repo.Save(&booking); err != nil {
		return nil, errors.New("failed to save the data on database")
	}
	return booking, nil
}
