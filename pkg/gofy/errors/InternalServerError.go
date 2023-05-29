package errors

type InternalServerError struct {
}

func (e InternalServerError) Error() string {
	return "Internal Server Error"
}
