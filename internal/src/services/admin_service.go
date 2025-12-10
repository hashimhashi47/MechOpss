package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/utils"
	"errors"
	"fmt"
)

// indirect conncetion with repository using dependency injection
type AdminService struct {
	Repo repository.Repository
}

// constructor
func NewAdminService(repo repository.Repository) *AdminService {
	return &AdminService{Repo: repo}
}

// admin login
func (s *AdminService) AdminLogin(email, password string) (*models.Admin, string, string, error) {
	admin, err := s.Repo.FindAdminByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("admin not found")
	}
	fmt.Println("✅", admin)

	if password != admin.Password {
		return nil, "", "", errors.New("invalid password")
	}

	access, err := utils.AccessToken(admin.ID, admin.Email, admin.Role)
	if err != nil {
		return nil, "", "", errors.New("failed generating access token")
	}

	refresh, err := utils.RefershToken(admin.ID, admin.Email, admin.Role)
	if err != nil {
		return nil, "", "", errors.New("failed generating refresh token")
	}

	admin.RefreshToken = refresh

	if err := s.Repo.Save(&admin); err != nil {
		return nil, "", "", errors.New("failed saving refresh token")
	}

	return admin, access, refresh, nil
}

//admin logout with remove the refershtoken

func (s *AdminService) ServiceAdminLogout(id string) error {
	var admin models.Admin
	if err := s.Repo.FindByID(&admin, id); err != nil {
		return errors.New("failed to find admin")
	}
	admin.RefreshToken = ""

	if err := s.Repo.Save(&admin); err != nil {
		return errors.New("failed remove refersh token")
	}
	return nil
}
