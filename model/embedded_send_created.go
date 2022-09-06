package model

type EmbeddedSendCreated struct {
	DocumentId string `json:"document_id"`
	SendUrl    string `json:"send_url"`
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
