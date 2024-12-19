package error

import "fmt"

type UserNotFoundError struct {
	username string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with name %s not found", e.username)
}
