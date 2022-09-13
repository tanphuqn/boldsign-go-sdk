package model

type EmbeddedSigningLink struct {
	SignLink string `json:"signLink"`
}

func (s *EmbeddedSigningLink) GetSignLink() string {
	if s != nil {
		return s.SignLink
	}
	return ""
}
