package services

import (
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

func (s *UserService) ServiceUpdateuser() {

}

func (s *UserService) Logout(userID interface{}) error {

	err := s.repo.UpdateRefreshToken(models.User{}, userID, "")
	if err != nil {
		return errors.New("failed to clear refresh token")
	}
	return nil
}
