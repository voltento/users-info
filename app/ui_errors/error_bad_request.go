package ui_errors

type ErrorBadRequest struct {
	msg string
}

func (e *ErrorBadRequest) Error() string {
	return e.msg
}

func NewErrorBadRequest(s string) error {
	return &ErrorBadRequest{msg: s}
}
