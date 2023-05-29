package errors

import "fmt"

type EntityNotFound struct {
	Entity string `json:"entity"`
	ID     string `json:"id,omitempty"`
}

func (e EntityNotFound) Error() string {
	return fmt.Sprintf("No '%v' found for Id: '%v'", e.Entity, e.ID)
}
