package model

type EmbeddedSendCreated struct {
	DocumentId string `json:"documentId,omitempty"`
	SendUrl    string `json:"sendUrl,omitempty"`
}

type EmbeddedTemplateCreated struct {
	TemplateId string `json:"templateId,omitempty"`
	CreateUrl  string `json:"createUrl,omitempty"`
	EditUrl    string `json:"editUrl,omitempty"`
}

// GetDocumentId returns DocumentId
func (e *EmbeddedSendCreated) GetDocumentId() string {
	if e != nil {
		return e.DocumentId
	}
	return ""
}

// GetSendUrl returns SendUrl
func (e *EmbeddedSendCreated) GetSendUrl() string {
	if e != nil {
		return e.SendUrl
	}
	return ""
}

// GetTemplateId returns TemplateId
func (e *EmbeddedTemplateCreated) GetTemplateId() string {
	if e != nil {
		return e.TemplateId
	}
	return ""
}

// GetCreateUrl returns CreateUrl
func (e *EmbeddedTemplateCreated) GetCreateUrl() string {
	if e != nil {
		return e.CreateUrl
	}
	return ""
}
