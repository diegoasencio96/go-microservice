package error

type (
	APIError struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)
