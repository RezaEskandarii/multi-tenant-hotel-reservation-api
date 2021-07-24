package message_keys

// hold localization files key here
const (
	CrudMessages = "CrudMessages."
	Errors       = "Errors."
	Users        = "Users."

	Created = CrudMessages + "Created"
	Updated = CrudMessages + "Updated"
	Deleted = CrudMessages + "Deleted"

	NotFound            = Errors + "NotFound"
	InternalServerError = Errors + "InternalServerError"
	BadRequest          = Errors + "BadRequest"

	UsernameDuplicated = Users + "DuplicatedUsername"

	TypeHasResidence = "TypeHasResidence"
)
