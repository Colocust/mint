package linkedList

type Error struct {
	S string
}

func (e *Error) Error() string {
	return e.S
}

func NewError(s string) error {
	return &Error{s}
}
