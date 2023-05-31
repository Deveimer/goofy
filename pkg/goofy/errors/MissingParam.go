package errors

import (
	"fmt"
	"strings"
)

type MissingParam struct {
	Param []string `json:"param"`
}

func (e MissingParam) Error() string {
	if len(e.Param) > 1 {
		return fmt.Sprintf("Parameters " + strings.Join(e.Param, ", ") + " are required for this request")
	} else if len(e.Param) == 1 {
		return fmt.Sprintf("Parameter " + e.Param[0] + " is required for this request")
	} else {
		return "This request is missing parameters"
	}
}
