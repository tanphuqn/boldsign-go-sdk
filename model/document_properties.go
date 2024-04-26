package model

type DocumentProperties struct {
	DocumentId          string                  `json:"documentId"`
	TemplateId          string                  `json:"templateId,omitempty"`
	BrandId             string                  `json:"brandId,omitempty"`
	MessageTitle        string                  `json:"messageTitle,omitempty"`
	DocumentDescription string                  `json:"documentDescription,omitempty"`
	Status              string                  `json:"status,omitempty"`
	Files               []File                  `json:"files,omitempty"`
	SenderDetail        SenderDetail            `json:"senderDetail,omitempty"`
	SignerDetails       []DocumentSignerDetails `json:"signerDetails,omitempty"`
	CcDetails           []DocumentCcDetails     `json:"ccDetails,omitempty"`
	ReminderSettings    ReminderSettings        `json:"reminderSettings,omitempty"`
	Reassign            []DocumentReassign      `json:"reassign,omitempty"`
	DocumentHistory     []AuditTrail            `json:"documentHistory,omitempty"`
	ActivityBy          string                  `json:"activityBy,omitempty"`
	ActivityDate        int                     `json:"activityDate,omitempty"`
	ActivityAction      string                  `json:"activityAction,omitempty"`
	CreatedDate         int                     `json:"createdDate,omitempty"`
	ExpiryDays          int                     `json:"expiryDays,omitempty"`
	ExpiryDate          int                     `json:"expiryDate,omitempty"`
	EnableSigningOrder  bool                    `json:"enableSigningOrder,omitempty"`
	IsDeleted           bool                    `json:"isDeleted,omitempty"`
	RevokeMessage       string                  `json:"revokeMessage,omitempty"`
	DeclineMessage      string                  `json:"declineMessage,omitempty"`
	ApplicationId       string                  `json:"applicationId,omitempty"`
	Labels              []string                `json:"labels,omitempty"`
	DisableEmails       bool                    `json:"disableEmails,omitempty"`
	EnablePrintAndSign  bool                    `json:"enablePrintAndSign,omitempty"`
	EnableReassign      bool                    `json:"enableReassign,omitempty"`
	DisableExpiryAlert  bool                    `json:"disableExpiryAlert,omitempty"`
	HideDocumentId      bool                    `json:"hideDocumentId,omitempty"`
	Roles               []TemplateRole          `json:"roles,omitempty,omitempty"`
	Description         string                  `json:"description,omitempty"`
	DocumentTitle       string                  `json:"documentTitle,omitempty"`
	DocumentMessage     string                  `json:"documentMessage,omitempty"`
	Title               string                  `json:"title,omitempty"`
}

type File struct {
	DocumentName string `json:"documentName"`
	Order        int    `json:"order"`
	PageCount    int    `json:"pageCount"`
}

type SenderDetail struct {
	Name           string `json:"name"`
	PrivateMessage string `json:"privateMessage"`
	EmailAddress   string `json:"emailAddress"`
	IsViewed       bool   `json:"isViewed"`
}

type ReminderSettings struct {
	EnableAutoReminder bool `json:"enableAutoReminder"`
	ReminderDays       int  `json:"reminderDays"`
	ReminderCount      int  `json:"reminderCount"`
}

type DocumentReassign struct {
	SignerEmail string `json:"signerEmail"`
	Order       int    `json:"order"`
	Message     string `json:"message"`
}

type AuditTrail struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ToName    string `json:"toName"`
	ToEmail   string `json:"toEmail"`
	Ipaddress string `json:"ipaddress"`
	Action    string `json:"action"`
	Timestamp int    `json:"timestamp"`
}
