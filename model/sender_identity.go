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
