package errors

type DbError struct {
	Err error
}

func (e DbError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return "DB Error"
}
