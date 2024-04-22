package model

type TemplateRole struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
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
