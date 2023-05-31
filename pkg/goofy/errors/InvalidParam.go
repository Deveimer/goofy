package errors

import (
	"fmt"
	"strings"
)

type InvalidParam struct {
	Param []string
}

func (e InvalidParam) Error() string {
	if len(e.Param) > 1 {
		return fmt.Sprintf("Incorrect value for parameters: " + strings.Join(e.Param, ", "))
	} else if len(e.Param) == 1 {
		return fmt.Sprintf("Incorrect value for parameter: " + e.Param[0])
	} else {
		return "This request has invalid parameters"
	}
}
