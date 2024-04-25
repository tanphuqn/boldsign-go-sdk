package model

type EmbeddedDocumentRequest struct {
	Title                 string           `form_field:"Title"`
	Message               string           `form_field:"Message"`
	RedirectUrl           string           `form_field:"RedirectUrl"`
	Signers               []DocumentSigner `form_field:"Signers"`
	Files                 []string         `form_field:"Files"`
	EnableSigningOrder    bool             `form_field:"EnableSigningOrder"`
	SendViewOption        string           `form_field:"SendViewOption"`
	BrandId               string           `form_field:"BrandId"`
	OnBehalfOf            string           `form_field:"OnBehalfOf"`
	SendLinkValidTill     string           `form_field:"SendLinkValidTill"`
	ShowToolbar           bool             `form_field:"ShowToolbar"`
	ShowNavigationButtons bool             `form_field:"ShowNavigationButtons"`
	ShowSaveButton        bool             `form_field:"ShowSaveButton"`
	ShowPreviewButton     bool             `form_field:"ShowPreviewButton"`
	ShowSendButton        bool             `form_field:"ShowSendButton"`
	EnableReassign        bool             `form_field:"EnableReassign"`
	EnablePrintAndSign    bool             `form_field:"EnablePrintAndSign"`
	UseTextTags           bool             `form_field:"UseTextTags"`
	DisableExpiryAlert    bool             `form_field:"DisableExpiryAlert"`
	EnableEmbeddedSigning bool             `form_field:"EnableEmbeddedSigning"`
	HideDocumentId        bool             `form_field:"HideDocumentId"`
	DisableEmails         bool             `form_field:"DisableEmails"`
	ExpiryDays            int              `form_field:"ExpiryDays"`
	ReminderSettings      ReminderSettings `form_field:"ReminderSettings"`

	Description         string         `form_field:"Description"`
	DocumentTitle       string         `form_field:"DocumentTitle"`
	DocumentMessage     string         `form_field:"DocumentMessage"`
	ViewOption          string         `form_field:"ViewOption"`
	AllowNewFiles       bool           `form_field:"AllowNewFiles"`
	AllowModifyFiles    bool           `form_field:"AllowModifyFiles"`
	LinkValidTill       string         `form_field:"LinkValidTill"`
	ShowTooltip         bool           `form_field:"ShowTooltip"`
	AutoDetectFields    bool           `form_field:"autoDetectFields"`
	AllowMessageEditing bool           `form_field:"AllowMessageEditing"`
	Roles               []TemplateRole `form_field:"Roles" json:"roles,omitempty"`
	TemplateIds         []string       `form_field:"templateIds"`
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

// func (e *EmbeddedDocumentRequest) IsEnableSigningOrder() bool {
// 	if e != nil {
// 		return e.EnableSigningOrder
// 	}
// 	return false
// }

// func (e *EmbeddedDocumentRequest) IsShowToolbar() bool {
// 	if e != nil {
// 		return e.ShowToolbar
// 	}
// 	return false
// }

func (e *EmbeddedDocumentRequest) GetRedirectUrl() string {
	if e != nil {
		return e.RedirectUrl
	}
	return ""
}

// func (e *EmbeddedDocumentRequest) GetSendViewOption() string {
// 	if e != nil {
// 		return e.SendViewOption
// 	}
// 	return ""
// }

// GetFile returns File
func (e *EmbeddedDocumentRequest) GetFiles() []string {
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

func (e *EmbeddedDocumentRequest) GetRoles() []TemplateRole {
	if e != nil {
		return e.Roles
	}
	return nil
}

func (e *EmbeddedDocumentRequest) GetTemplateIds() []string {
	if e != nil {
		return e.TemplateIds
	}
	return nil
}
