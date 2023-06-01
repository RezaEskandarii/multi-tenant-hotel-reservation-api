package repositories

import (
	"context"
	"fmt"
	"github.com/andskur/argon2-hashing"
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/tenant_resolver"
	"reservation-api/internal/utils/file_utils"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type UserRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
	Db         gorm.DB
}

func NewUserRepository(r *tenant_database_resolver.TenantDatabaseResolver) *UserRepository {
	return &UserRepository{
		DbResolver: r,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	user.TenantId = tenant_resolver.GetCurrentTenant(ctx)

	if tx := db.Create(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	user.TenantId = tenant_resolver.GetCurrentTenant(ctx)

	if tx := db.Updates(&user); tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Find(ctx context.Context, id uint64) (*models.User, error) {

	user := models.User{}
	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Where("id=?", id).Find(&user); tx.Error != nil {
		return nil, tx.Error
	}

	if user.Id == 0 {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	user := models.User{Username: username}

	if tx := db.Where(user).Find(&user); tx.Error != nil {
		return nil, tx.Error
	}

	if user.Id == 0 {
		return nil, nil
	}

	return &user, nil
}
func (r *UserRepository) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*models.User, error) {

	user := models.User{}
	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Where("username=?", username).Find(&user); tx.Error != nil {
		return nil, tx.Error
	}

	if user.Id == 0 {
		return nil, nil
	}

	err := argon2.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Deactivate(ctx context.Context, id uint64) (*models.User, error) {

	user := models.User{}
	db := r.DbResolver.GetTenantDB(ctx)

	query := db.Model(&models.User{}).Where("id=?", id).Find(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	user.IsActive = false

	if tx := db.Model(&models.User{}).Updates(user); tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Activate(ctx context.Context, id uint64) (*models.User, error) {

	user := models.User{}
	db := r.DbResolver.GetTenantDB(ctx)

	query := db.Model(&models.User{}).Where("id=?", id).Find(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	user.IsActive = true
	if tx := db.Model(&models.User{}).Updates(user); tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {

	db := r.DbResolver.GetTenantDB(ctx)

	if tx := db.Model(&models.User{}).Where("id=?", id).Delete(&models.User{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	db := r.DbResolver.GetTenantDB(ctx)
	return paginatedList(&models.User{}, db, input)
}

func (r *UserRepository) HashPassword(password string) (string, error) {
	params := argon2.DefaultParams
	hash, err := argon2.GenerateFromPassword([]byte(password), params)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return fmt.Sprintf("%s", hash), nil
}

func (r *UserRepository) Seed(ctx context.Context, jsonFilePath string) error {

	db := r.DbResolver.GetTenantDB(ctx)

	users := make([]models.User, 0)
	if err := file_utils.CastJsonFileToStruct(jsonFilePath, &users); err != nil {
		return err
	}

	for _, user := range users {
		var existingUser models.User
		if err := db.Where("username", user.Username).FirstOrCreate(&existingUser, &user).Error; err != nil {
			return err
		}
		if existingUser.Id == 0 {
			hashedPassword, err := r.HashPassword(existingUser.Password)
			if err != nil {
				return err
			}
			existingUser.Password = hashedPassword

			if err := db.Create(&existingUser).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
