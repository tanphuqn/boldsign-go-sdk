package model

type EmbeddedSendCreated struct {
	DocumentId string `json:"documentId"`
	SendUrl    string `json:"sendUrl"`
}

type EmbeddedTemplateCreated struct {
	TemplateId string `json:"templateId"`
	CreateUrl  string `json:"createUrl"`
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
