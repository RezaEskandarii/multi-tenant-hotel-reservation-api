package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
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

// Delete removes user permanently;
func (s *UserService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}

// FindAll returns paginates list of cities.
func (s *UserService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}
