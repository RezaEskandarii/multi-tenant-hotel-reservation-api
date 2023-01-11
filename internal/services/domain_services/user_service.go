package domain_services

import (
	"context"
	"fmt"
	"github.com/andskur/argon2-hashing"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
	"reservation-api/internal_errors"
)

type UserService struct {
	Repository *repositories.UserRepository
}

// NewUserService returns new UserService
func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{Repository: r}
}

// Create creates new User.
func (s *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {

	usr, err := s.FindByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	if usr != nil {
		return nil, internal_errors.DuplicatedUser
	}

	hash, err := argon2.GenerateFromPassword([]byte(user.Password), argon2.DefaultParams)

	if err != nil {
		return nil, err
	}

	user.Password = fmt.Sprintf("%s", hash)

	return s.Repository.Create(ctx, user)
}

// Update updates User.
func (s *UserService) Update(ctx context.Context, user *models.User) (*models.User, error) {

	return s.Repository.Update(ctx, user)
}

// Find returns User and if it does not find the User, it returns nil.
func (s *UserService) Find(ctx context.Context, id uint64) (*models.User, error) {

	return s.Repository.Find(ctx, id)
}

// FindByUsername returns User by username and if it does not find the User, it returns nil.
func (s *UserService) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	return s.Repository.FindByUsername(ctx, username)
}

// Delete removes user permanently;
func (s *UserService) Delete(ctx context.Context, id uint64) error {

	return s.Repository.Delete(ctx, id)
}

// FindAll returns paginated list of cities.
func (s *UserService) FindAll(ctx context.Context, filter *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	return s.Repository.FindAll(ctx, filter)
}

// Activate find user by id and activate.
func (s *UserService) Activate(ctx context.Context, id uint64) (*models.User, error) {

	return s.Repository.Activate(ctx, id)
}

// Deactivate find and deactivates user by user id.
func (s *UserService) Deactivate(ctx context.Context, id uint64) (*models.User, error) {

	return s.Repository.Deactivate(ctx, id)
}

// Seed seed given json file to multi_tenancy_database.
func (s *UserService) Seed(ctx context.Context, jsonFilePath string) error {
	return s.Repository.Seed(ctx, jsonFilePath)
}

// FindByUsernameAndPassword finds user by username and password.
func (s *UserService) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*models.User, error) {
	return s.Repository.FindByUsernameAndPassword(ctx, username, password)
}
