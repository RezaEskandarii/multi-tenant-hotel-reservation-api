package repositories

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"reservation-api/internal/utils"
)

type SeederRepository struct {
	DB *gorm.DB
}

func NewSeederRepository(db *gorm.DB) *SeederRepository {
	return &SeederRepository{DB: db}
}

// Seed gives the model and address of the json file from the input and
// fills given model's fields with given json and bulk inserts given model.
func (r *SeederRepository) Seed(jsonFilePath string, model []interface{}) error {

	if err := castJsonFileToModel(jsonFilePath, model); err != nil {
		if err := r.DB.CreateInBatches(model, len(model)).Error; err != nil {
			return err
		}
		return err
	}
	return nil
}

func castJsonFileToModel(path string, model []interface{}) error {
	if utils.FileExists(path) {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		} else {
			return json.Unmarshal(file, &model)
		}
	} else {
		return nil
	}
}
