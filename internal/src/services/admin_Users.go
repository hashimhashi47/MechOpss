package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"errors"
)

// get all users service
func (s *AdminService) GetUsers() ([]models.User, error) {
	var Users []models.User

	if err := s.Repo.FindAll(&Users); err != nil {
		return nil, errors.New("user not found")
	}
	return Users, nil
}

// update user
func (s *AdminService) ServiceUpdateUser(Input models.User, id string) (interface{}, error) {
	var user models.User

	if err := s.Repo.FindByID(&user, id); err != nil {
		return models.User{}, errors.New("user not found")
	}

	if Input.FirstName != "" {
		user.FirstName = Input.FirstName
	}
	if Input.Email != "" {
		user.Email = Input.Email
	}
	if Input.Phone != "" {
		user.Phone = Input.Phone
	}
	if Input.Password != "" {
		hash, _ := utils.Hashing(Input.Password)
		user.Password = string(hash)
	}

	if err := s.Repo.Save(&user); err != nil {
		return nil, errors.New("cannot update user")
	}

	return user, nil
}


// delete user 
func (s *AdminService) ServiceDeleteuser(model interface{}, id string) error {
	if err := s.Repo.Delete(model, id); err != nil {
		return errors.New("user not found")
	}
	return nil
}


//block user
func (s *AdminService) ServiceBlockUser(id string , Body bool) (models.User, error) {

	var user models.User
	if err := s.Repo.FindByID(&user, id); err != nil {
		return models.User{}, errors.New("user not found")
	}
	user.Block = Body

	if err := s.Repo.Save(&user); err != nil {
		return models.User{}, errors.New("not able to block this user")
	}

	return user, nil
}
