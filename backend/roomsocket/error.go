package roomsocket

var (
	ErrMissingRoomIDHeader = NewError("this request requires a room_id header with 1 value")
	ErrRoomIDHeaderInvalid = NewError("the provided room_id header was invalid")
	ErrRoomDoesNotExist    = NewError("the requested room does not exist")
)

var (
	// ErrUnexpected represents an error that the system did not anticipate, and does not have a
	// specific error message for.
	ErrUnexpected = NewError("unexpected error")
)

// Error represents an error within this domain.
type Error struct {
	s string
}

// NewError creates a new instance of an error.
func NewError(s string) Error {
	return Error{
		s,
	}
}

// Error returns the error message associated with this error.
func (e Error) Error() string {
	return e.s
}
