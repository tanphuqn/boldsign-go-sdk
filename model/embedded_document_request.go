package model

type EmbeddedDocumentRequest struct {
	Title              string           `form_field:"Title"`
	Message            string           `form_field:"Message"`
	EnableSigningOrder bool             `form_field:"EnableSigningOrder"`
	RedirectUrl        string           `form_field:"RedirectUrl"`
	Signers            []DocumentSigner `form_field:"Signers"`
	Files              []DocumentFile   `form_field:"Files"`
	SendViewOption     string           `form_field:"SendViewOption"`
	ShowToolbar        bool             `form_field:"ShowToolbar"`
	// ShowNavigationButtons bool             `form_field:"ShowNavigationButtons"`
	// ShowSaveButton        bool             `form_field:"ShowSaveButton"`
	// ShowPreviewButton     bool             `form_field:"ShowPreviewButton"`
	// ShowSendButton        bool             `form_field:"ShowSendButton"`
}

func (e *EmbeddedDocumentRequest) GetTitle() string {
	if e != nil {
		return e.Title
	}
	return ""
}

func (e *EmbeddedDocumentRequest) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

func (e *EmbeddedDocumentRequest) IsEnableSigningOrder() bool {
	if e != nil {
		return e.EnableSigningOrder
	}
	return false
}

func (e *EmbeddedDocumentRequest) IsShowToolbar() bool {
	if e != nil {
		return e.ShowToolbar
	}
	return false
}

func (e *EmbeddedDocumentRequest) GetRedirectUrl() string {
	if e != nil {
		return e.RedirectUrl
	}
	return ""
}

func (e *EmbeddedDocumentRequest) GetSendViewOption() string {
	if e != nil {
		return e.SendViewOption
	}
	return ""
}

// GetFile returns File
func (e *EmbeddedDocumentRequest) GetFiles() []DocumentFile {
	if e != nil {
		return e.Files
	}
	return nil
}

// GetSignerRoles returns Signers
func (e *EmbeddedDocumentRequest) GetSigners() []DocumentSigner {
	if e != nil {
		return e.Signers
	}
	return nil
}
