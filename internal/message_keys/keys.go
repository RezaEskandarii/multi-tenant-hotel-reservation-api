package message_keys

// hold localization files key here
const (
	CrudMessages = "CrudMessages."
	Errors       = "Errors."
	Users        = "Users."
	hotels       = "hotels."
	Rooms        = "Rooms."

	Created = CrudMessages + "Created"
	Updated = CrudMessages + "Updated"
	Deleted = CrudMessages + "Deleted"

	NotFound            = Errors + "NotFound"
	InternalServerError = Errors + "InternalServerError"
	BadRequest          = Errors + "BadRequest"

	UsernameDuplicated = Users + "DuplicatedUsername"

	TypeHashotel    = "TypeHashotel"
	GradeHashotel   = "GradeHashotel"
	ValidationError = "ValidationError"
	GenderInvalid   = "GenderInvalid"

	HotelRepeatPostalCode = hotels + "RepeatPostalCode"

	InvalidRoomCleanStatus = Rooms + "InvalidCleanStatus"
	RoomTypeHasRoomErr     = Rooms + "RoomTypeHasRoomErr"
)
