package model

type ErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	Message      string `json:"message"`
	ErrorContent string `json:"errorContent"`
}

// GetMessage returns Message
func (e *ErrorResponse) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

// GetName returns Name
func (e *ErrorResponse) GetErrorCode() int {
	return e.ErrorCode
}

// GetMessage returns Message
func (e *ErrorResponse) GetErrorContent() string {
	if e != nil {
		return e.ErrorContent
	}
	return ""
}
