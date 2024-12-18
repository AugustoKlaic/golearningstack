package error

type UnhashablePasswordError struct{}

func (e *UnhashablePasswordError) Error() string {
	return "password cannot be hashed!"
}
