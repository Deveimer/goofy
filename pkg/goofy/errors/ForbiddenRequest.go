package errors

import "fmt"

type ForbiddenRequest struct {
	URL string
}

func (f ForbiddenRequest) Error() string {
	return fmt.Sprintf("Access to %v is forbidden", f.URL)
}
