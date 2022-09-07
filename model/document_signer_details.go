package model

type DocumentSignerDetails struct {
	SignerName             string                   `json:"signerName"`
	SignerRole             string                   `json:"signerRole"`
	SignerEmail            string                   `json:"signerEmail"`
	Status                 string                   `json:"status"`
	IsAuthenticationFailed bool                     `json:"isAuthenticationFailed"`
	EnableEmailOTP         bool                     `json:"enableEmailOTP"`
	IsDeliveryFailed       bool                     `json:"isDeliveryFailed"`
	IsViewed               bool                     `json:"isViewed"`
	Order                  int                      `json:"order"`
	SignerType             string                   `json:"signerType"`
	IsReassigned           bool                     `json:"isReassigned"`
	PrivateMessage         string                   `json:"privateMessage"`
	FormFields             []map[string]interface{} `json:"formFields"`
	HostEmail              string                   `json:"hostEmail"`
	HostName               string                   `json:"hostName"`
}
