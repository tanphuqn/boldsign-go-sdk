package boldsign

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

// createEmbeddedRequestUrl creates a new embedded signature with template id
func (m *Client) CreateEmbeddedRequestUrlFromTemplate(templateId string, req model.EmbeddedDocumentRequest) (*model.EmbeddedSendCreated, error) {
	bodyBuf, bodyWriter, err := m.MarshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}
	path := fmt.Sprintf("template/createEmbeddedRequestUrl?templateId=%s", templateId)
	response, err := m.post(path, bodyBuf, *bodyWriter, false)
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

// CreateEmbeddedTemplateRequestUrl creates a new embedded signature with template id
func (m *Client) CreateEmbeddedTemplateRequestUrl(req model.EmbeddedDocumentRequest) (*model.EmbeddedTemplateCreated, error) {
	bodyBuf, bodyWriter, err := m.MarshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}

	response, err := m.post("template/createEmbeddedTemplateUrl", bodyBuf, *bodyWriter, false)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.EmbeddedTemplateCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	return data, nil

}

// GetEmbeddedTemplateEditUrl creates a new embedded signature with template id
func (m *Client) GetEmbeddedTemplateEditUrl(templateId string, req model.EmbeddedDocumentRequest) (*model.EmbeddedTemplateCreated, error) {
	bodyBuf, bodyWriter, err := m.MarshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}
	path := fmt.Sprintf("template/getEmbeddedTemplateEditUrl?templateId=%s", templateId)
	response, err := m.post(path, bodyBuf, *bodyWriter, false)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data := &model.EmbeddedTemplateCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
	return data, nil

}

// DeleteTemplate
func (m *Client) DeleteTemplate(templateId string) error {
	path := fmt.Sprintf("template/delete?templateId=%s", templateId)
	response, err := m.delete(path, false)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

// UpdateTemplate
func (m *Client) UpdateTemplate(templateId string, req model.EmbeddedDocumentRequest) error {
	path := fmt.Sprintf("template/edit?templateId=%s", templateId)
	jsonData, _ := json.Marshal(req)
	response, err := m.putJson(path, jsonData, true)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	return err
}

// GetTemplate
func (m *Client) GetTemplate(templateId string) (*model.DocumentProperties, error) {
	template := &model.DocumentProperties{}
	path := fmt.Sprintf("template/properties?templateId=%s", templateId)
	response, err := m.get(path, true)
	if err != nil {
		return template, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(template)
	if err != nil {
		return template, err
	}

	return template, nil
}

func (m *Client) DownloadTemplate(documentId string, onBehalfOf string) ([]byte, error) {
	path := fmt.Sprintf("template/download?templateId=%s&onBehalfOf=%s", documentId, onBehalfOf)
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
