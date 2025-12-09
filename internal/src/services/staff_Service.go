package services

import "MechOpss/internal/src/repository"

type StaffService struct {
	Repo repository.Repository
}

func NewStaffService(repo repository.Repository) *StaffService {
	return &StaffService{Repo: repo}
}

//staff login
func (ss *StaffService) ServiceLogin() {
	
}

//updated the current details
func (ss *StaffService) ServiceProfileUpdate() {

}

//logout 
func (ss *StaffService) ServiceLogout() {

}
