package model

type Warning struct {
	Message string `json:"warning_msg"`
	Name    string `json:"warning_name"`
}

// GetMessage returns Message
func (w *Warning) GetMessage() string {
	if w != nil {
		return w.Message
	}
	return ""
}

// GetName returns Name
func (w *Warning) GetName() string {
	if w != nil {
		return w.Name
	}
	return ""
}
