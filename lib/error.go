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
