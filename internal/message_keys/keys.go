package message_keys

// hold localization files key here
const (
	CrudMessages = "CrudMessages."
	Errors       = "Errors."

	Created = CrudMessages + "Created"
	Updated = CrudMessages + "Updated"
	Deleted = CrudMessages + ".Deleted"

	NotFound            = Errors + "NotFound"
	InternalServerError = Errors + "InternalServerError"
	BadRequest          = Errors + "BadRequest"
)
