package apierr

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e ApiError) Error() string {
	return e.Message
}

func NewAPIError(message string, status int) *ApiError {
	return &ApiError{
		Message: message,
		Status:  status,
	}
}
