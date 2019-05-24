package model

// Error 自定义error.
type Error struct {
	ErrMsg string
}

// NewError .
func NewError(msg string) *Error {
	return &Error{ErrMsg: msg}
}

func (err *Error) Error() string {
	return err.ErrMsg
}