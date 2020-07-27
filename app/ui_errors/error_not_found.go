package ui_errors

type ErrorNotFond struct {
	msg string
}

func (e *ErrorNotFond) Error() string {
	return e.msg
}

func NewErrorNotFound(s string) error {
	return &ErrorNotFond{msg: s}
}
