package model

type SenderCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SenderUpdateRequest struct {
	Name string `json:"name"`
}

type SenderCreated struct {
	ID string `json:"senderIdentityId"`
}

type SenderIdentityDetail struct {
	Status       string `json:"status"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	CreatedBy    string `json:"createdBy"`
	ApprovedDate string `json:"approvedDate"`
}

func (s *SenderIdentityDetail) IsVerified() bool {
	if s != nil {
		return s.Status == "Verified"
	}
	return false
}

type SenderIdentities []SenderIdentityDetail
type SenderIdentitiesResponse struct {
	Result      SenderIdentities       `json:"result"`
	PageDetails map[string]interface{} `json:"pageDetails"`
}
