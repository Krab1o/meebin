package repository

type ErrorType int

type Error struct {
	Type ErrorType
	Err  error
}

const (
	NotFound ErrorType = iota
	Internal
)

func (e ErrorType) String() string {
	switch e {
	case NotFound:
		return "Not found"
	case Internal:
		return "Internal Error"
	default:
		return "Unknown Error"
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}

func newError(errType ErrorType, err error) *Error {
	return &Error{
		Type: errType,
		Err:  err,
	}
}

func NewNotFoundError(err error) *Error {
	return newError(NotFound, err)
}
func NewInternalError(err error) *Error {
	return newError(Internal, err)
}
