package error

import "fmt"

type MessageNotFoundError struct {
	Id int
}

func (e *MessageNotFoundError) Error() string {
	return fmt.Sprintf("message with Id %d not found", e.Id)
}
