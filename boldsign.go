package boldsign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

const (
	baseURL               string = "https://api-eu.boldsign.com/v1/"
	baseDomain            string = "https://account-eu.boldsign.com"
	FormFieldKey          string = "form_field"
	FileKey               string = "Files"
	SignersKey            string = "Signers"
	TitleKey              string = "Title"
	RedirectUrlKey        string = "RedirectUrl"
	MessageKey            string = "Message"
	SendViewOptionKey     string = "SendViewOption"
	ShowToolbarKey        string = "ShowToolbar"
	EnableSigningOrderKey string = "EnableSigningOrder"
)

const (
	PreparePage int64 = 0
	FillingPage int64 = 1
)

// Client contains APIKey and optional http.client
type Client struct {
	Secret     string
	ClientID   string
	BaseURL    string
	BaseDomain string
	HTTPClient *http.Client
}

// CreateEmbeddedRequestUrl creates a new embedded signature with template id
func (m *Client) CreateEmbeddedRequestUrl(req model.EmbeddedDocumentRequest) (*model.EmbeddedSendCreated, error) {
	bodyBuf, bodyWriter, err := m.marshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}

	response, err := m.post("document/createEmbeddedRequestUrl", bodyBuf, *bodyWriter)
	if err != nil {
		return nil, err
	}

	data := &model.EmbeddedSendCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *Client) GetProperties(documentId string) (*model.DocumentProperties, error) {
	path := fmt.Sprintf("https://api-eu.boldsign.com/v1/document/properties?documentId=%s", documentId)
	response, err := m.get(path)
	if err != nil {
		return nil, err
	}
	data := &model.DocumentProperties{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Client) marshalMultipartEmbeddedSignatureRequest(embRequest model.EmbeddedDocumentRequest) (*bytes.Buffer, *multipart.Writer, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	structType := reflect.TypeOf(embRequest)
	val := reflect.ValueOf(embRequest)

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		f := valueField.Interface()
		val := reflect.ValueOf(f)
		field := structType.Field(i)
		fieldTag := field.Tag.Get(FormFieldKey)
		switch val.Kind() {
		case reflect.Slice:
			switch fieldTag {
			case SignersKey:
				for i, signer := range embRequest.GetSigners() {
					formField, err := bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][EmailAddress]", SignersKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(signer.GetEmailAddress()))

					formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][Name]", SignersKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(signer.GetName()))
					// fmt.Println(fmt.Sprintf("%s[%v][Name]", SignersKey, i), "=", signer.GetName())

					formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][SignerOrder]", SignersKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(signer.GetSignerOrder())))

					// fmt.Println(fmt.Sprintf("%s[%v][SignerOrder]", SignersKey, i), "=", signer.GetSignerOrder())
				}
			case FileKey:
				for _, path := range embRequest.GetFiles() {
					file, _ := os.Open(path)
					formField, err := bodyWriter.CreateFormFile("Files", file.Name())
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
					// fmt.Println("Files=", file.Name())
				}
			}
		case reflect.Bool:
			formField, err := bodyWriter.CreateFormField(fieldTag)
			if err != nil {
				return nil, nil, err
			}
			// fmt.Println(fieldTag, "=", m.boolToIntString(val.Bool()))
			formField.Write([]byte(m.boolToIntString(val.Bool())))
		default:
			if val.String() != "" {
				// fmt.Println(fieldTag, "=", val.String())
				formField, err := bodyWriter.CreateFormField(fieldTag)
				if err != nil {
					return nil, nil, err
				}
				formField.Write([]byte(val.String()))
			}
		}
	}

	bodyWriter.Close()
	return bodyBuf, bodyWriter, nil
}

func (m *Client) boolToIntString(value bool) string {
	if value == true {
		return "true"
	}
	return "false"
}
