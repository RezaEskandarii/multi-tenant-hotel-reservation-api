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
- add addDetails function to rate_code_handler
- add rate code details service
- add rate_code_details_repository
- add feat: add new entities to migration
- add new fields to rate_code_details structModel
- add rate_code_details structModel
- add file_transfer_service interface
- add command runner
- add auditChannel to handlers
- add thumbnail structModel
- add checkIn function to reservation_repository
- add CreateReservationRequest function to reservation_repository
- add seed function to country_repository
- add cache to city service
- add cache to city service
- implement cache manage functions
- implement redis cache layer
- add audit middleware
- add audit middleware
- add audit service
- add audit repository
- add audit structModel
- protect routes with JWTAuthMiddleware
- add auth middleware
- implement refreshToken function in auth_handler
- implement signin function
- add AuthHandler and implement signin function
- add FindByUsernameAndPassword function userservice
- add roomType seeder
- add seeder service
- add seeder repository
- add roomTypes seed file
- add reservation type field to reservation structModel
- add reservation repository
- add validation to reservation structModel
- add validation to reservation structModel
- generate reservationKey's hash before create
- add hash_generator
- add reservationLock structModel
- add reservation structModel
- add ratCodeHandler to serviceRegistry
- add ratCodeHandler's functions
- add rate_code_handler
- add rate_code_service
- implement rate_code_repository functions
- add rate_code_repository
- add validator to rate_code structModel
- add rate_code structModel
- add rate group handler
- add rate group repository service
- add rate group repository functions
- add rate group repository
- add rate group structModel
- add guest structModel to migrator
- implement guest handler functions
- add guest handler to service registry
- implement guests handler
- add guests handler
- add activate and deactivate functions to roomType service
- add activate and deactivate functions to roomType repository
- add guest service
- add guest repository
- add guest structModel
- add room handler
- add room room service
- add room repository
- add logger to handlers
- add room type handler
- add room type service
- add room room type repository
- add room struct fields
- add room type struct
- add logger to roomType handler
- add logger to residence type handler
- add logger to residence handler
- add logger to residence grade handler
- add logger to province handler
- add logger to currency handler
- add logger to country handler
- add logger to city handler
- add room structModel
- add residence handler
- add residence service
- add residence repository
- add validation to structs
- add validator package and add validation tags to models
- implement find all and delete function in residence_type_handler
- implement find all and delete function in residence_grade_handler
- add residence grade handler
- add ResidenceGradeService
- add residence_grade_repository
- add delete function to residence_type_repository
- implement update residence type function
- add ResidenceType handler
- add ResidenceTypeService and ResidenceTypeRepository
- add ResidenceType and ResidenceGrade struct and add to Residence struct relations
- add Residence struct
- implement logger package methods and add disitlog library
- add cors origin middleware
- add persian new translations
- add translator package to handlers
- add findByUsername function roomType handler
- add findByUsername function in userService
- add roomType handler
- add roomType service

### Fix
- singleton audit channel listener in audit_middleware
- regenerate gmo mod file
- regenerate gmo mod file
- register rate group handler in service registry

### Refactor
- refactor audit_middleware
- refactor audit_middleware
- remove logger from config module
- refactor kernel module codes
- move services to domain_service module
- refactor residence word to hotel in all part of project
- refactor handlers input
- refactor handlers input
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
- validate roomType's gender before create
- ignore validation to some structModel relations
- add pointer to residence relations
- prevent to validation residence releations
- move handlers and middlewares to api directory
- move test files to test package
- improve some codes
- refactor som docs
- beautify some codes
- remove some extra error codes
- refactor residence_grade and residence_type handler
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
- add roomType reository
- add i18n translator to city handler
- add i18n translator to city handler
- add message_keys module to hold translation message files keys
- add message_keys module to hold translation message files keys
- add localized message translator module
- add roomType structModel
- add currency handler
- add currency service
- add currency repository
- add currency structModel
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
- add country and city structModel
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
