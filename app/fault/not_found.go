package fault

type NotFond struct {
	msg string
}

func (e *NotFond) Error() string {
	return e.msg
}

func NewNotFound(s string) error {
	return &NotFond{msg: s}
}
