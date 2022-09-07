package model

type DocumentProperties struct {
	DocumentId          string                  `json:"documentId"`
	BrandId             string                  `json:"brandId"`
	MessageTitle        string                  `json:"messageTitle"`
	DocumentDescription string                  `json:"documentDescription"`
	Status              string                  `json:"status"`
	Files               []File                  `json:"files"`
	SenderDetail        SenderDetail            `json:"senderDetail"`
	SignerDetails       []DocumentSignerDetails `json:"signerDetails"`
	ccDetails           []DocumentCcDetails     `json:"ccDetails"`
	reminderSettings    ReminderSettings        `json:"reminderSettings"`
	reassign            DocumentReassign        `json:"reassign"`
	documentHistory     AuditTrail              `json:"documentHistory"`
	activityBy          string                  `json:"activityBy"`
	activityDate        int                     `json:"activityDate"`
	activityAction      string                  `json:"activityAction"`
	createdDate         int                     `json:"createdDate"`
	expiryDays          int                     `json:"expiryDays"`
	expiryDate          int                     `json:"expiryDate"`
	enableSigningOrder  bool                    `json:"enableSigningOrder"`
	isDeleted           bool                    `json:"isDeleted"`
	revokeMessage       string                  `json:"revokeMessage"`
	declineMessage      string                  `json:"declineMessage"`
	applicationId       string                  `json:"applicationId"`
	labels              []string                `json:"labels"`
	disableEmails       bool                    `json:"disableEmails"`
	enablePrintAndSign  bool                    `json:"enablePrintAndSign"`
	enableReassign      bool                    `json:"enableReassign"`
	disableExpiryAlert  bool                    `json:"disableExpiryAlert"`
	hideDocumentId      bool                    `json:"hideDocumentId"`
}

type File struct {
	DocumentName string `json:"documentName"`
	Order        int    `json:"order"`
	PageCount    string `json:"pageCount"`
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
