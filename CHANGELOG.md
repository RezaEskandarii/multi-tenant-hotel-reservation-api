<a name="unreleased"></a>
## [Unreleased]

### Doc
- add comment
- change readme file content.
- add comment
- add comment
- add comment
- add comment
- add comment
- add comment
- add comment to some codes
- add comemnt to some codes
- add comment to some codes
- add document to some functions
- regenerate changelog file
- add some comments
- regenerate changelog file

### Feat
- protect routes with JWTAuthMiddleware
- add auth middleware
- implement refreshToken function in auth_handler
- implement signin function
- add AuthHandler and implement signin function
- add FindByUsernameAndPassword function userservice
- add user seeder
- add seeder service
- add seeder repository
- add users seed file
- add reservation type field to reservation model
- add reservation repository
- add validation to reservation model
- add validation to reservation model
- generate reservationKey's hash before create
- add hash_generator
- add reservationLock model
- add reservation model
- add ratCodeHandler to serviceRegistry
- add ratCodeHandler's functions
- add rate_code_handler
- add rate_code_service
- implement rate_code_repository functions
- add rate_code_repository
- add validator to rate_code model
- add rate_code model
- add rate group handler
- add rate group repository service
- add rate group repository functions
- add rate group repository
- add rate group model
- add guest model to migrator
- implement guest handler functions
- add guest handler to service registry
- implement guests handler
- add guests handler
- add activate and deactivate functions to user service
- add activate and deactivate functions to user repository
- add guest service
- add guest repository
- add guest model
- add room handler
- add room room service
- add room repository
- add logger to handlers
- add room type handler
- add room type service
- add room room type repository
- add room struct fields
- add room type struct
- add logger to user handler
- add logger to hotel type handler
- add logger to hotel handler
- add logger to hotel grade handler
- add logger to province handler
- add logger to currency handler
- add logger to country handler
- add logger to city handler
- add room model
- add hotel handler
- add hotel service
- add hotel repository
- add validation to structs
- add validator package and add validation tags to models
- implement find all and delete function in hotel_type_handler
- implement find all and delete function in hotel_grade_handler
- add hotel grade handler
- add hotelGradeService
- add hotel_grade_repository
- add delete function to hotel_type_repository
- implement update hotel type function
- add HotelType handler
- add hotelTypeService and hotelTypeRepository
- add HotelType and HotelGrade struct and add to hotel struct relations
- add hotel struct
- implement logger package methods and add disitlog library
- add cors origin middleware
- add persian new translations
- add translator package to handlers
- add findByUsername function user handler
- add findByUsername function in userService
- add user handler
- add user service

### Fix
- register rate group handler in service registry

### Refactor
- rename module name to reservation-api
- refactor currency_handler
- refactor country_handler
- refactor some codes
- add logger config
- add logger config
- add new models to migration models
- refactor validations in models beforeCreate hook
- change logger library with zaplogger
- tidy go mod file
- refactor some field names is tests
- change handlers input with shared struct
- refactor some codes
- refactor some codes
- improve repositories code
- validate user's gender before create
- ignore validation to some model relations
- add pointer to hotel relations
- prevent to validation hotel releations
- move handlers and middlewares to api directory
- move test files to test package
- improve some codes
- refactor som docs
- beautify some codes
- remove some extra error codes
- refactor hotel_grade and hotel_type handler
- rename some structs
- add logs file to gitignore
- refactor services dependencies

### Test
- refactor country_service_test
- add country service test


<a name="v0.1"></a>
## v0.1 - 2021-06-29
### Doc
- add document to province service functions
- add document to currency service functions
- add document to country service functions
- add document to city service functions
- add changelog file
- add document and comments
- add document

### Feat
- add user reository
- add i18n translator to city handler
- add i18n translator to city handler
- add message_keys module to hold translation message files keys
- add message_keys module to hold translation message files keys
- add localized message translator module
- add user model
- add currency handler
- add currency service
- add currency repository
- add currency model
- implement province handler
- add province service
- add province repository
- add province struct
- implement generic findAll function in repository layer
- add get cities handler in country_handler
- add get cities function in country_repository
- register city handler
- implement city handler
- implement city respository
- implement pagination in country repository
- implement update and find function in country handler
- add some middlewares to echo router
- implement country handler's create method
- implement country handler
- implement paginated_list's methods
- implement application config struct
- add country_repository
- add country and city model
- add project structure and folders

### Fix
- fix check error in create city handler

### Refactor
- change localizer directory name to translator
- remove main to cmd directory
- rename data field to records in paginates list struct
- refactor migration function
- refactor migration function
- add yaml package to go mod


[Unreleased]: https://github.com/RezaEskandarii/hotel-reservation/compare/v0.1...HEAD
