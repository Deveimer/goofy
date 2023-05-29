package errors

import "fmt"

type DbStoreNotInitialized struct {
	DBName string
	Reason string
}

func (d DbStoreNotInitialized) Error() string {
	return fmt.Sprintf("couldn't initialize %v, %v", d.DBName, d.Reason)
}
