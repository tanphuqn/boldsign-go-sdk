package model

type DocumentCcDetails struct {
	EmailAddress string `json:"emailAddress"`
	UserId       string `json:"userId"`
	IsViewed     bool   `json:"isViewed"`
}

func (s *DocumentCcDetails) GetEmailAddress() string {
	if s != nil {
		return s.EmailAddress
	}
	return ""
}
