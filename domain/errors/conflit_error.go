package errors

type ConflictError struct {
	message string
}

func (e *ConflictError) Error() string {
	return e.message
}
