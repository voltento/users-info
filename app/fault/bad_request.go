package fault

type BadRequest struct {
	msg string
}

func (e *BadRequest) Error() string {
	return e.msg
}

func NewBadRequest(s string) error {
	return &BadRequest{msg: s}
}
