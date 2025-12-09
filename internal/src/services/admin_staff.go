package services

import (
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"errors"
)

// register staff
func (c *AdminService) ServiceStaffRegister(model *models.Staff) (interface{}, error) {

	if err := c.Repo.FindByEmail(models.Staff{}, model.Email); err == nil {
		return nil, errors.New("already exisit")
	}

	Hash, err := utils.Hashing(model.Password)
	IdentityID := utils.RandomIDGenerate("S-ID-")
	if err != nil {
		return nil, errors.New("failed to hashpassword")
	}
	model.Password = string(Hash)
	model.IdentityCard = IdentityID

	if err := c.Repo.Insert(&model); err != nil {
		return nil, errors.New("failed to Store staff data")
	}

	return model, nil
}


// GetStaff
func (c *AdminService) ServiceGetStaff() ([]models.Staff, error) {
	var Staff []models.Staff
	if err := c.Repo.FindAllStaffsWithBookings(&Staff); err != nil {
		return nil, err
	}
	return Staff, nil
}

// updateStaff
func (c *AdminService) ServiceUpdateStaff(Input models.Staff, id string) (interface{}, error) {
	var ExistingStaff models.Staff

	if err := c.Repo.FindByID(&ExistingStaff, id); err != nil {
		return nil, errors.New("staff not found")
	}

	if Input.FirstName != "" {
		ExistingStaff.FirstName = Input.FirstName
	}

	if Input.LastName != "" {
		ExistingStaff.LastName = Input.LastName
	}

	if Input.Email != "" {
		ExistingStaff.Email = Input.Email
	}

	if Input.Password != "" {
		hash, _ := utils.Hashing(Input.Password)
		ExistingStaff.Password = string(hash)
	}

	if Input.Department != "" {
		ExistingStaff.Department = Input.Department
	}

	if err := c.Repo.Save(&ExistingStaff); err != nil {
		return nil, errors.New("failed to update the staff")
	}

	return ExistingStaff, nil
}

// block staff
func (c *AdminService) ServiceBlockStaff(id string, body bool) (models.Staff, error) {

	var staff models.Staff
	if err := c.Repo.FindByID(&staff, id); err != nil {
		return models.Staff{}, errors.New("staff not found")
	}

	staff.Block = body

	if err := c.Repo.Save(&staff); err != nil {
		return models.Staff{}, errors.New("failed to block staff")
	}

	return staff, nil
}
