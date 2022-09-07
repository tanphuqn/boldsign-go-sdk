package model

type DocumentCC struct {
	EmailAddress string `json:"emailAddress"`
}

func (s *DocumentCC) GetEmailAddress() string {
	if s != nil {
		return s.EmailAddress
	}
	return ""
}
