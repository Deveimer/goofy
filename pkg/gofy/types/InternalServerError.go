package types

type InternalServerError struct {
}

func (e InternalServerError) Error() string {
	return "Internal Server Error"
}
