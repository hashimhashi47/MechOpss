package services

import (
	"MechOpss/internal/src/models"
)

// dashhboard
// user count service
func (s *AdminService) UsersCount() (int64, error) {
	return s.Repo.Count(&models.User{})
}

// staff count service
func (s *AdminService) StaffCount() (int64, error) {
	return s.Repo.Count(&models.Staff{})
}

// booking count service
func (s *AdminService) BookingCount() (int64, error) {
	return s.Repo.Count(&models.Booking{})
}

// recent booking service
func (s *AdminService) RecentBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	err := s.Repo.LimitFind(&bookings, 2)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

// next booking service
func (s *AdminService) NextBooking() (models.Booking, error) {
	var booking models.Booking
	err := s.Repo.First(&booking)
	if err != nil {
		return models.Booking{}, nil
	}
	return booking, nil
}
