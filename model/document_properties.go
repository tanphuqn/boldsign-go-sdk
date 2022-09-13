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
	CcDetails           []DocumentCcDetails     `json:"ccDetails"`
	ReminderSettings    ReminderSettings        `json:"reminderSettings"`
	Reassign            []DocumentReassign      `json:"reassign"`
	DocumentHistory     []AuditTrail            `json:"documentHistory"`
	ActivityBy          string                  `json:"activityBy"`
	ActivityDate        int                     `json:"activityDate"`
	ActivityAction      string                  `json:"activityAction"`
	CreatedDate         int                     `json:"createdDate"`
	ExpiryDays          int                     `json:"expiryDays"`
	ExpiryDate          int                     `json:"expiryDate"`
	EnableSigningOrder  bool                    `json:"enableSigningOrder"`
	IsDeleted           bool                    `json:"isDeleted"`
	RevokeMessage       string                  `json:"revokeMessage"`
	DeclineMessage      string                  `json:"declineMessage"`
	ApplicationId       string                  `json:"applicationId"`
	Labels              []string                `json:"labels"`
	DisableEmails       bool                    `json:"disableEmails"`
	EnablePrintAndSign  bool                    `json:"enablePrintAndSign"`
	EnableReassign      bool                    `json:"enableReassign"`
	DisableExpiryAlert  bool                    `json:"disableExpiryAlert"`
	HideDocumentId      bool                    `json:"hideDocumentId"`
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
