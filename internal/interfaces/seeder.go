package interfaces

type Seeder interface {
	// Seed gives the model and address of the json file from the input and
	// fills given model's fields with given json and bulk inserts given model.
	Seed(jsonFilePath string, model []interface{}, uniqueColumns map[string]interface{}) error
}
