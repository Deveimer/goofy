package errors

import "fmt"

type FileNotFound struct {
	FileName string `json:"fileName"`
	Path     string `json:"path"`
}

func (f FileNotFound) Error() string {
	return fmt.Sprintf("File %v not found at location %v", f.FileName, f.Path)
}
