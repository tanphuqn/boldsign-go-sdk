package model

type EmbeddedSendCreated struct {
	DocumentId string `json:"documentId"`
	SendUrl    string `json:"sendUrl"`
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
