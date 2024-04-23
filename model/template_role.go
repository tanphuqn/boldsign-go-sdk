package model

type TemplateRole struct {
	Name               string `json:"name"`
	Index              int    `json:"index"`
	DefaultSignerName  string `json:"defaultSignerName"`
	DefaultSignerEmail string `json:"defaultSignerEmail"`
	SignerOrder        int    `json:"signerOrder"`
	SignerType         string `json:"signerType"`
}

func (s *TemplateRole) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

func (s *TemplateRole) GetIndex() int {
	if s != nil {
		return s.Index
	}
	return 1
}

func (s *TemplateRole) GetDefaultSignerName() string {
	if s != nil {
		return s.DefaultSignerName
	}
	return ""
}

func (s *TemplateRole) GetDefaultSignerEmail() string {
	if s != nil {
		return s.DefaultSignerEmail
	}
	return ""
}
func (s *TemplateRole) GetSignerOrder() int {
	if s != nil {
		return s.SignerOrder
	}
	return 1
}

func (s *TemplateRole) GetSignerType() string {
	if s != nil {
		return s.SignerType
	}
	return ""
}
