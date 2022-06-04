package lib

type ErrBadRequest struct {
	Message string
}

func NewErrBadRequest(message string) *ErrBadRequest {
	return &ErrBadRequest{Message: message}
}

func (e ErrBadRequest) Error() string {
	return e.Message
}

type ErrNotFound struct {
	Message string
}

func NewErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{Message: message}
}

func (e ErrNotFound) Error() string {
	return e.Message
}
