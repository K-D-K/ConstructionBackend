package error

type APIError struct {
	message string
}

func (error *APIError) Error() string {
	return error.message
}

func ThrowAPIError(message string) *APIError {
	return &APIError{
		message: message,
	}
}
