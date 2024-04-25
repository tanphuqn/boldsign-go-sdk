package boldsign

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

// Send document using multiple templates
func (m *Client) MergeAndSend(req model.EmbeddedDocumentRequest) (*model.EmbeddedSendCreated, error) {
	jsonData, _ := json.Marshal(req)
	response, err := m.postJson("template/mergeAndSend", jsonData, true)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.EmbeddedSendCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CreateEmbeddedRequestUrl creates a new embedded signature
func (m *Client) CreateEmbeddedRequestUrl(req model.EmbeddedDocumentRequest) (*model.EmbeddedSendCreated, error) {
	bodyBuf, bodyWriter, err := m.MarshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}

	response, err := m.post("document/createEmbeddedRequestUrl", bodyBuf, *bodyWriter, false)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.EmbeddedSendCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	return data, nil

}

func (m *Client) GetEmbeddedSignLink(documentId string, signerEmail string, redirectUrl string) (*model.EmbeddedSigningLink, error) {
	path := fmt.Sprintf("document/getEmbeddedSignLink?documentId=%s", documentId)
	if signerEmail != "" {
		path = fmt.Sprintf("%s&signerEmail=%s", path, signerEmail)
	}
	if redirectUrl != "" {
		path = fmt.Sprintf("%s&redirectUrl=%s", path, redirectUrl)
	}
	response, err := m.get(path, false)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.EmbeddedSigningLink{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Client) GetProperties(documentId string) (*model.DocumentProperties, error) {
	path := fmt.Sprintf("document/properties?documentId=%s", documentId)
	response, err := m.get(path, false)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.DocumentProperties{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Client) DownloadDocument(documentId string, onBehalfOf string) ([]byte, error) {
	path := fmt.Sprintf("document/download?documentId=%s&onBehalfOf=%s", documentId, onBehalfOf)
	response, err := m.get(path, false)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
