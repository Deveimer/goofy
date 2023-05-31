package types

type Response struct {
	Data interface{} `json:"data" xml:"data"`
	Meta interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
}
