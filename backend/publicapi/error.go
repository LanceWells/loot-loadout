package publicapi

var (
	ErrHubNotFound      = NewError("the requested websocket does not exist")
	ErrHubNameInvalid   = NewError("the requested websocket name is invalid")
	ErrHubAlreadyExists = NewError("the requested websocket already exists")
)

type Error struct {
	s string
}

func NewError(s string) Error {
	return Error{
		s,
	}
}

func (e Error) Error() string {
	return e.s
}
