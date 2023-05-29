package errors

type EntityAlreadyExists struct{}

func (e EntityAlreadyExists) Error() string {
	return "entity already exists"
}
