package errors

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}
