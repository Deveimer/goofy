package errors

import (
	"fmt"
	"strings"
)

type MultipleErrors struct {
	StatusCode int     `json:"-" xml:"-"`
	Errors     []error `json:"errors" xml:"errors"`
}

func (m MultipleErrors) Error() string {
	var result string

	for _, v := range m.Errors {
		result += fmt.Sprintf("%s\n", v)
	}

	return strings.TrimSuffix(result, "\n")
}
