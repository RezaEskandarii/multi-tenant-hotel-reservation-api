package message_keys

// hold localization files key here
const (
	CrudMessages = "CrudMessages."
	Errors       = "Errors."
	Users        = "Users."
	Residences   = "Residences."
	Rooms        = "Rooms."

	Created = CrudMessages + "Created"
	Updated = CrudMessages + "Updated"
	Deleted = CrudMessages + "Deleted"

	NotFound            = Errors + "NotFound"
	InternalServerError = Errors + "InternalServerError"
	BadRequest          = Errors + "BadRequest"

	UsernameDuplicated = Users + "DuplicatedUsername"

	TypeHasResidence  = "TypeHasResidence"
	GradeHasResidence = "GradeHasResidence"
	ValidationError   = "ValidationError"
	GenderInvalid     = "GenderInvalid"

	ResidenceRepeatPostalCode = Residences + "RepeatPostalCode"

	InvalidRoomCleanStatus = Rooms + "InvalidCleanStatus"
	RoomTypeHasRoomErr     = Rooms + "RoomTypeHasRoomErr"
)
