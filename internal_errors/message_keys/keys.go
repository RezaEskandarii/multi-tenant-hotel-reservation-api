package message_keys

// hold localization files key here
const (
	/*********************** root keys ************************/
	crudMessages = "CrudMessages."
	errors       = "internal_errors."
	users        = "Users."
	hotels       = "Hotels."
	rooms        = "Rooms."
	reservation  = "Reservation."
	/************************************************************/
	Created = crudMessages + "Created"
	Updated = crudMessages + "Updated"
	Deleted = crudMessages + "Deleted"
	/************************************************************/
	NotFound            = errors + "NotFound"
	InternalServerError = errors + "InternalServerError"
	BadRequest          = errors + "BadRequest"
	/************************************************************/
	UsernameDuplicated = users + "DuplicatedUsername"
	UserNotFound       = users + "UserNotFound"
	UserIsDeActive     = users + "UserIsNotActive"
	/************************************************************/
	TypeHashotel    = "TypeHashotel"
	GradeHashotel   = "GradeHashotel"
	ValidationError = "ValidationError"
	GenderInvalid   = "GenderInvalid"
	/************************************************************/
	HotelRepeatPostalCode = hotels + "RepeatPostalCode"
	/************************************************************/
	InvalidRoomCleanStatus    = rooms + "InvalidCleanStatus"
	RoomTypeHasRoomErr        = rooms + "RoomTypeHasRoomErr"
	RoomHasReservationRequest = rooms + "HasReservationRequest"
	/************************************************************/
	InvalidReservationRequestKey      = reservation + "InvalidReservationRequestKey"
	EmptySharerError                  = reservation + "EmptySharerError"
	ReservationConflictError          = reservation + "ReservationConflictError"
	ImpossibleReservationLatDateError = reservation + "ImpossibleReservationLatDateError"
	CheckOutDateEmptyError            = reservation + "CheckOutDateEmptyError"
	CheckInDateEmptyError             = reservation + "CheckInDateEmptyError"
)
