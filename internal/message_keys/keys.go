package message_keys

// hold localization files key here
const (
	crudMessages = "CrudMessages."
	errors       = "Errors."
	users        = "Users."
	hotels       = "Hotels."
	rooms        = "Rooms."
	reservation  = "Reservation."

	Created = crudMessages + "Created"
	Updated = crudMessages + "Updated"
	Deleted = crudMessages + "Deleted"

	NotFound            = errors + "NotFound"
	InternalServerError = errors + "InternalServerError"
	BadRequest          = errors + "BadRequest"

	UsernameDuplicated = users + "DuplicatedUsername"

	TypeHashotel    = "TypeHashotel"
	GradeHashotel   = "GradeHashotel"
	ValidationError = "ValidationError"
	GenderInvalid   = "GenderInvalid"

	HotelRepeatPostalCode = hotels + "RepeatPostalCode"

	InvalidRoomCleanStatus    = rooms + "InvalidCleanStatus"
	RoomTypeHasRoomErr        = rooms + "RoomTypeHasRoomErr"
	RoomHasReservationRequest = rooms + "HasReservationRequest"

	InvalidReservationRequestKey = reservation + "InvalidReservationRequestKey"
	EmptySharerError             = reservation + "EmptySharerError"
	ReservationConflictError     = reservation + "ReservationConflictError"
)
