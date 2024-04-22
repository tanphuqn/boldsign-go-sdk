package boldsign

import (
	"encoding/json"
	"fmt"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

// CreateSenderIdentity creates sender identity
func (m *Client) CreateSenderIdentity(senderRequest model.SenderCreateRequest) (*model.SenderCreated, error) {
	jsonData, _ := json.Marshal(senderRequest)
	response, err := m.postJson("senderIdentities/create", jsonData, true)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.SenderCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateSenderIdentity
func (m *Client) UpdateSenderIdentity(email string, senderRequest model.SenderUpdateRequest) error {
	path := fmt.Sprintf("senderIdentities/update?email=%s", email)
	jsonData, _ := json.Marshal(senderRequest)
	response, err := m.postJson(path, jsonData, true)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	return err
}

// DeleteSenderIdentity
func (m *Client) DeleteSenderIdentity(email string) error {
	path := fmt.Sprintf("senderIdentities/delete?email=%s", email)
	response, err := m.delete(path, true)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	return err
}

// DeleteSenderIdentity
func (m *Client) VerifySenderIdentity(email string) (model.SenderIdentityDetail, bool, error) {
	path := fmt.Sprintf("senderIdentities/list?PageSize=1&Page=1&Search=%s", email)
	response, err := m.get(path, true)

	if err != nil {
		return model.SenderIdentityDetail{}, false, err
	}

	defer response.Body.Close()
	data := &model.SenderIdentitiesResponse{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return model.SenderIdentityDetail{}, false, err
	}

	if len(data.Result) == 0 {
		return model.SenderIdentityDetail{}, false, nil
	}

	return data.Result[0], data.Result[0].IsVerified(), nil
}
