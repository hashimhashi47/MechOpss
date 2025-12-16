package services

import (
	"MechOpss/internal/src/dto"
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

// indirect conncetion with repository using dependency injection
type UserService struct {
	repo repository.Repository
}

// constructor
func NewUserservices(r repository.Repository) *UserService {
	return &UserService{repo: r}
}

// signup
func (r *UserService) Signup(data *models.User) (interface{}, error) {

	var existing models.User
	if err := r.repo.FindByEmail(&existing, data.Email); err == nil {
		return nil, errors.New("already exisit")
	}

	hash, _ := utils.Hashing(data.Password)
	data.Password = string(hash)

	if err := r.repo.Insert(data); err != nil {
		return nil, errors.New(err.Error())
	}

	return data, nil
}

// login
func (r *UserService) Login(data *models.User) (interface{}, uint, string, error) {

	var user models.User

	if err := r.repo.FindByEmail(&user, data.Email); err != nil {
		return nil, 0, "", errors.New("user not found")
	}

	if user.Block && user.Role == "user" {
		return nil, 0, "", errors.New("your account is blocked")
	}

	if err := utils.HashCompare(user.Password, data.Password); err != nil {
		return nil, 0, "", errors.New("invalid password")
	}

	access, _ := utils.AccessToken(user.ID, user.Email, user.Role)
	refresh, _ := utils.RefershToken(user.ID, user.Email, user.Role)

	user.RefreshToken = refresh
	r.repo.Save(&user)

	resp := gin.H{
		"access":  access,
		"refresh": refresh,
		"userid":  user.ID,
		"user":    user,
	}

	return resp, user.ID, access, nil
}

// logout for admin,user,staff
func (s *UserService) Logout(userID interface{}) error {

	err := s.repo.UpdateRefreshToken(models.User{}, userID, "")
	if err != nil {
		return errors.New("failed to clear refresh token")
	}
	return nil
}

// user booking
func (c *UserService) ServiceBookingUser(data models.Booking) (string, error) {
	ID := utils.RandomIDGenerate("BOOK")
	data.ID = ID

	if err := c.repo.Insert(&data); err != nil {
		return "", errors.New("booking failed")
	}
	return ID, nil
}

// get the whole booking status of that user
func (s *UserService) ServiceGetBookingstatus(id string) (Details interface{}, err error) {
	var user models.User

	if err := s.repo.FindWithPreload(&user, "Bookings", id); err != nil {
		return nil, errors.New("failed to find the user")
	}

	var detail []dto.UserBooking
	for _, v := range user.Bookings {
		detail = append(detail, dto.UserBooking{
			ID:         v.ID,
			CarModel:   v.CarModel,
			CarNumber:  v.CarNumber,
			UserStatus: v.UserStatus,
			Message:    v.Message,
		})
	}

	return detail, nil
}

// get the whole booking status of that user
func (s *UserService) ServiceGetBookingstatusID(Userid, BookingID string) (interface{}, error) {
	var user models.User

	if err := s.repo.FindWithPreload(&user, "Bookings", Userid); err != nil {
		return nil, errors.New("failed to find the user")
	}

	var BookingDetails dto.UserBooking
	found := false
	for _, v := range user.Bookings {
		if v.ID == BookingID {
			BookingDetails.ID = v.ID
			BookingDetails.CarModel = v.CarModel
			BookingDetails.CarNumber = v.CarModel
			BookingDetails.UserStatus = v.UserStatus
			BookingDetails.Message = v.Message
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("booked id does not belong to this user")
	}

	return BookingDetails, nil
}

// get user bookedds
func (s *UserService) ServiceGetBookeds(id string) ([]dto.UserBookeds, error) {
	var user models.User
	if err := s.repo.FindWithPreload(&user, "Booked", id); err != nil {
		return nil, errors.New("failed to find the user")
	}

	var bookedsDetails []dto.UserBookeds

	for _, v := range user.Booked {
		bookedsDetails = append(bookedsDetails, dto.UserBookeds{
			ID:            v.ID,
			Date:          v.Date,
			CarModel:      v.CarModel,
			CarNumber:     v.CarNumber,
			UserID:        *v.UserID,
			Status:        v.Status,
			Message:       v.Message,
			ServiceStart:  v.ServiceStart,
			ServiceEnd:    v.ServiceEnd,
			Description:   v.Description,
			PaymentAmount: v.PaymentAmount,
			PaymentMode:   v.PaymentMode,
			PaymentStatus: v.PaymentStatus,
			Delivery:      v.Delivery,
		})
	}
	return bookedsDetails, nil
}

// get bookeds by id
func (s *UserService) ServiceGetBookedsID(Userid, BookingID string) (interface{}, error) {
	var user models.User
	if err := s.repo.FindWithPreload(&user, "Booked", Userid); err != nil {
		return nil, errors.New("failed to find the user")
	}

	if len(user.Booked) == 0 {
		return nil, errors.New("no bookings found for this user")
	}

	var bookedsDetails dto.UserBookeds

	found := false
	for _, v := range user.Booked {
		if v.ID == BookingID {
			bookedsDetails.Date = v.Date
			bookedsDetails.CarModel = v.CarModel
			bookedsDetails.CarNumber = v.CarNumber
			bookedsDetails.UserID = *v.UserID
			bookedsDetails.Status = v.Status
			bookedsDetails.Message = v.Message
			bookedsDetails.ServiceStart = v.ServiceStart
			bookedsDetails.ServiceEnd = v.ServiceEnd
			bookedsDetails.Description = v.Description
			bookedsDetails.PaymentAmount = v.PaymentAmount
			bookedsDetails.PaymentMode = v.PaymentMode
			bookedsDetails.PaymentStatus = v.PaymentStatus
			bookedsDetails.Delivery = v.Delivery
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("booked id does not belong to this user")
	}
	return bookedsDetails, nil
}

// user payments can show
func (s *UserService) ServiceGetPayments(id string) (data []dto.Userpayments, errr error) {
	var user models.User

	if err := s.repo.FindWithPreload(&user, "Booked", id); err != nil {
		return nil, errors.New("failed to find user")
	}

	var payments []dto.Userpayments
	for _, v := range user.Booked {
		payments = append(payments, dto.Userpayments{
			ID:            v.ID,
			PaymentStatus: v.PaymentStatus,
			PaymentAmount: v.PaymentAmount,
			PaymentMode:   v.PaymentMode,
		})
	}

	return payments, nil
}

// pay the service
func (s *UserService) ServicePayTheService(BookedID, Userid string, amount float64) (interface{}, error) {
	var bookeds models.Bookeds
	var user models.User

	if err := s.repo.FindWithPreload(&user,"Booked" ,Userid); err != nil {
		return nil, errors.New("failed to find the user")
	}
	found := false
	for _, v := range user.Booked {
		if v.ID == BookedID {
			found = true
			if err := s.repo.FindByID(&bookeds, BookedID); err != nil {
				return nil, errors.New("failed to find booking")
			}
			break
		}
	}

	if !found {
		return nil, errors.New("booking not linked to user")
	}

	if amount != bookeds.PaymentAmount {
		return nil, errors.New("amount mismatched")
	}

	bookeds.PaymentMode = "upi"
	bookeds.PaymentAmount = amount
	bookeds.PaymentStatus = "paid"

	if err := s.repo.Save(&bookeds); err != nil {
		return nil, errors.New("failed to save payment")
	}

	payment := dto.Userpayments{
		ID:            bookeds.ID,
		PaymentStatus: bookeds.PaymentStatus,
		PaymentAmount: bookeds.PaymentAmount,
		PaymentMode:   bookeds.PaymentMode,
	}

	return payment, nil
}
