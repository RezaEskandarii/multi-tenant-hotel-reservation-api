package repositories

import (
	"github.com/andskur/argon2-hashing"
	"gorm.io/gorm"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/utils"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {

	if tx := r.DB.Create(&user); tx.Error != nil {

		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {

	if tx := r.DB.Updates(&user); tx.Error != nil {

		return nil, tx.Error
	}

	return user, nil
}

func (r *UserRepository) Find(id uint64) (*models.User, error) {

	model := models.User{}
	if tx := r.DB.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {

	model := models.User{Username: username}
	if tx := r.DB.Where(model).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}
func (r *UserRepository) FindByUsernameAndPassword(username string, password string) (*models.User, error) {

	model := models.User{}
	if tx := r.DB.Where("username=?", username).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	err := argon2.CompareHashAndPassword([]byte(model.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *UserRepository) Deactivate(id uint64) (*models.User, error) {

	user := models.User{}

	query := r.DB.Model(&models.User{}).Where("id=?", id).Find(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	user.IsActive = false

	if tx := r.DB.Model(&models.User{}).Updates(user); tx.Error != nil {

		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Activate(id uint64) (*models.User, error) {

	user := models.User{}

	query := r.DB.Model(&models.User{}).Where("id=?", id).Find(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	user.IsActive = true

	if tx := r.DB.Model(&models.User{}).Updates(user); tx.Error != nil {

		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) Delete(id uint64) error {
	if tx := r.DB.Model(&models.User{}).Where("id=?", id).Delete(&models.User{}); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *UserRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedList, error) {

	return paginatedList(&models.User{}, r.DB, input)
}

func (r *UserRepository) Seed(jsonFilePath string) error {

	users := make([]models.User, 0)
	if err := utils.CastJsonFileToStruct(jsonFilePath, &users); err == nil {
		for _, user := range users {
			var count int64 = 0
			if err := r.DB.Model(models.User{}).Where("username", user.Username).Count(&count).Error; err != nil {
				return err
			} else {
				if count == 0 {
					if err := r.DB.Create(&user).Error; err != nil {
						return err
					}
				}
			}
		}
	} else {
		return err
	}
	return nil
}
