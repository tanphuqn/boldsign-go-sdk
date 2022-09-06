package model

type ErrorResponse struct {
	Error    *Error    `json:"error"`
	Warnings []Warning `json:"warnings"`
}

type Error struct {
	Message string `json:"error_msg"`
	Name    string `json:"error_name"`
}

// GetMessage returns Message
func (e *Error) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

// GetName returns Name
func (e *Error) GetName() string {
	if e != nil {
		return e.Name
	}
	return ""
}

// GetError returns Error
func (er *ErrorResponse) GetError() *Error {
	if er != nil {
		return er.Error
	}
	return nil
}

// GetWarnings returns Warnings
func (er *ErrorResponse) GetWarnings() []Warning {
	if er != nil {
		return er.Warnings
	}
	return nil
}
