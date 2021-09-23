<a name="unreleased"></a>
## [Unreleased]

### Doc
- regenerate changelog file
- add some comments
- regenerate changelog file

### Feat
- add logger to user handler
- add logger to residence type handler
- add logger to residence handler
- add logger to residence grade handler
- add logger to province handler
- add logger to currency handler
- add logger to country handler
- add logger to city handler
- add model model
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
- add findByUsername function user handler
- add findByUsername function in userService
- add user handler
- add user service

### Refactor
- improve repositories code
- validate user's gender before create
- ignore validation to some model relations
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
