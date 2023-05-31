package response

type Responder interface {
	Response(data interface{}, err error)
}
