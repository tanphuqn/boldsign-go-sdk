package model

type DocumentSigner struct {
	Name         string
	EmailAddress string
	// PrivateMessage     string                   `field:"PrivateMessage"`
	// AuthenticationCode string                   `field:"AuthenticationCode"`
	// SignerOrder int `json:"SignerOrder"`
	// EnableEmailOTP     bool                     `field:"EnableEmailOTP"`
	// SignerType         string                   `field:"SignerType"`
	// FormFields         []map[string]interface{} `field:"FormFields"`
	// HostEmail          string                   `field:"HostEmail"`
	// Language           string                   `field:"Language"`
}

func (s *DocumentSigner) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

func (s *DocumentSigner) GetEmailAddress() string {
	if s != nil {
		return s.EmailAddress
	}
	return ""
}

// func (s *DocumentSigner) GetSignerOrder() int {
// 	if s != nil {
// 		return s.SignerOrder
// 	}
// 	return 0
// }
