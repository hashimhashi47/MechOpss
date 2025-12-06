package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/repository"
	"MechOpss/internal/src/utils"
	"errors"
)

type AdminService struct {
	Repo repository.Repository
}

func NewAdminService(repo repository.Repository) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) AdminLogin(email, password string) (*models.Admin, string, string, error) {
	// Fetch admin
	admin, err := s.Repo.FindAdminByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("admin not found")
	}


	if password != admin.Password {
		return nil, "", "", errors.New("invalid password")
	}

	// Generate tokens
	access, err := utils.AccessToken(admin.ID, admin.Email, admin.Role)
	if err != nil {
		return nil, "", "", errors.New("failed generating access token")
	}


	refresh, err := utils.RefershToken(admin.ID, admin.Email, admin.Role)
	if err != nil {
		return nil, "", "", errors.New("failed generating refresh token")
	}

	admin.RefreshToken = refresh

	if err := s.Repo.Save(admin); err != nil {
		return nil, "", "", errors.New("failed saving refresh token")
	}

	return admin, access, refresh, nil
}

