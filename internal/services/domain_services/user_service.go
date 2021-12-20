package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type UserService struct {
	Repository *repositories.UserRepository
}

// NewUserService returns new UserService
func NewUserService() *UserService {
	return &UserService{}
}

// Create creates new User.
func (s *UserService) Create(user *models.User) (*models.User, error) {

	return s.Repository.Create(user)
}

// Update updates User.
func (s *UserService) Update(user *models.User) (*models.User, error) {

	return s.Repository.Update(user)
}

// Find returns User and if it does not find the User, it returns nil.
func (s *UserService) Find(id uint64) (*models.User, error) {

	return s.Repository.Find(id)
}

// FindByUsername returns User by username and if it does not find the User, it returns nil.
func (s *UserService) FindByUsername(username string) (*models.User, error) {

	return s.Repository.FindByUsername(username)
}

// Delete removes user permanently;
func (s *UserService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}

// FindAll returns paginated list of cities.
func (s *UserService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Activate find user by id and activate.
func (s *UserService) Activate(id uint64) (*models.User, error) {

	return s.Repository.Activate(id)
}

// Deactivate find and deactivates user by user id.
func (s *UserService) Deactivate(id uint64) (*models.User, error) {

	return s.Repository.Deactivate(id)
}

// Seed seed given json file to database.
func (s *UserService) Seed(jsonFilePath string) error {
	return s.Repository.Seed(jsonFilePath)
}

// FindByUsernameAndPassword finds user by username and password.
func (s *UserService) FindByUsernameAndPassword(username string, password string) (*models.User, error) {
	return s.Repository.FindByUsernameAndPassword(username, password)
}
