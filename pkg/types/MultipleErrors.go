package types

import (
	"fmt"
	"strings"
)

type MultipleErrors struct {
	StatusCode int     `json:"-" xml:"-"`
	Errors     []error `json:"types" xml:"types"`
}

func (m MultipleErrors) Error() string {
	var result string

	for _, v := range m.Errors {
		result += fmt.Sprintf("%s\n", v)
	}

	return strings.TrimSuffix(result, "\n")
}
