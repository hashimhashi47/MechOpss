package controllers

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/services"
	"MechOpss/internal/src/utils"
	"MechOpss/internal/src/utils/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// indirect conncetion with service using dependency injection
type UserController struct {
	Service *services.UserService
}

// constructor
func NewController(s *services.UserService) *UserController {
	return &UserController{Service: s}
}

// register
func (s *UserController) RegisterAuth(c *gin.Context) {
	//validation
	var Input struct {
		FirstName string `json:"firstname" binding:"required"`
		Lastname  string `json:"lastname" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Phone     string `json:"phone" binding:"required,min=10"`
		Password  string `json:"password" binding:"required,min=6,max=16"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		errorMsg := utils.ErrorMessage(constants.BADREQUEST, err)
		c.JSON(http.StatusBadRequest, gin.H{"Binding erorr": errorMsg})
		return
	}

	CreateUser := models.User{
		FirstName: Input.FirstName,
		Lastname:  Input.Lastname,
		Email:     Input.Email,
		Password:  Input.Password,
		Phone:     Input.Phone,
		Role:      constants.User,
	}

	Response, err := s.Service.Signup(&CreateUser)
	if err != nil {
		errorMsg := utils.ErrorMessage(constants.BADREQUEST, err)
		c.JSON(http.StatusBadRequest, gin.H{"Unable to signup user": errorMsg})
		return
	}

	Success := utils.SuccessResponse(Response)
	c.JSON(http.StatusOK, gin.H{"Signup successfully": Success})

}

// login
func (s *UserController) LoginAuth(c *gin.Context) {

	var Input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	user := models.User{
		Email:    Input.Email,
		Password: Input.Password,
	}

	resp, userid, AccessToken, err := s.Service.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.SetCookie("Token", AccessToken, 7*24*3600, "/", "localhost", false, true)
	c.Set("id", userid)
	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(resp)})
}

// logout for user/admin/staff
func (s *UserController) UserLogout(c *gin.Context) {
	id, _ := c.Get("id")
	userID := id.(uint)

	c.SetCookie("Token", "", -1, "/", "localhost", false, true)
	if err := s.Service.Logout(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse("User logout")})
}

// user can see the booking status
func (s *UserController) GetBookingStatus(c *gin.Context) {
	id, _ := c.Get("id")
	UserID := fmt.Sprintf("%v", id)

	data, err := s.Service.ServiceGetBookingstatus(UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(data)})
}

// get booking status by id
func (s *UserController) GetBookingstatusID(c *gin.Context) {
	bookingID := c.Param("id")
	id, _ := c.Get("id")
	UserID := fmt.Sprintf("%v", id)

	data, err := s.Service.ServiceGetBookingstatusID(UserID, bookingID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(data)})
}

//get booked status

func (s *UserController) GetBookedsStatus(c *gin.Context) {
	id, _ := c.Get("id")
	UserID := fmt.Sprintf("%v", id)
	data, err := s.Service.ServiceGetBookeds(UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(data)})
}

// get booking status by id
func (s *UserController) GetBookedstatusID(c *gin.Context) {
	bookedID := c.Param("id")
	id, _ := c.Get("id")
	UserID := fmt.Sprintf("%v", id)

	data, err := s.Service.ServiceGetBookedsID(UserID, bookedID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": utils.ErrorMessage(constants.BADREQUEST, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": utils.SuccessResponse(data)})
}
