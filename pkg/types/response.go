package types

import "fmt"

type Response struct {
	StatusCode int         `json:"code"`
	Status     string      `json:"status"`
	Code       string      `json:"-"`
	Reason     string      `json:"reason"`
	ResourceID string      `json:"resourceId,omitempty"`
	Detail     interface{} `json:"detail,omitempty"`
	Path       string      `json:"path,omitempty"`
	RootCauses []RootCause `json:"rootCauses,omitempty"`
}

type RootCause map[string]interface{}

func (k Response) Error() string {
	if e, ok := k.Detail.(error); ok {
		return fmt.Sprintf("%v : %v ", k.Reason, e)
	}

	return k.Reason
}

type Error string

func (e Error) Error() string {
	return string(e)
}
