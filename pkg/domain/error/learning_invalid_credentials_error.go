package error

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "username or password invalid"
}
