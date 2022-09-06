package boldsign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"

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
	params, writer, err := m.marshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}

	response, err := m.post("document/createEmbeddedRequestUrl", params, *writer)
	if err != nil {
		fmt.Println("post Error:", err.Error())
		return nil, err
	}

	return m.parseSignatureRequestResponse(response)
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
					err := bodyWriter.WriteField(fmt.Sprintf("%s[%v][EmailAddress]", SignersKey, i), signer.GetEmailAddress())
					if err != nil {
						return nil, nil, err
					}
					fmt.Println(fmt.Sprintf("%s[%v][EmailAddress]", SignersKey, i), "=", signer.GetEmailAddress())

					err = bodyWriter.WriteField(fmt.Sprintf("%s[%v][Name]", SignersKey, i), signer.GetName())
					if err != nil {
						return nil, nil, err
					}
					fmt.Println(fmt.Sprintf("%s[%v][Name]", SignersKey, i), "=", signer.GetName())
				}
			case FileKey:
				for _, path := range embRequest.GetFiles() {
					file, _ := os.Open(path.FilePath)

					formField, err := bodyWriter.CreateFormFile("Files", file.Name())

					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
					fmt.Println("file=", file.Name())
				}
			}
		case reflect.Bool:
			err := bodyWriter.WriteField(fieldTag, m.boolToIntString(val.Bool()))
			if err != nil {
				return nil, nil, err
			}
			fmt.Println(fieldTag, "=", m.boolToIntString(val.Bool()))
		default:
			if val.String() != "" {
				fmt.Println(fieldTag, "=", val.String())
				err := bodyWriter.WriteField(fieldTag, val.String())
				if err != nil {
					return nil, nil, err
				}
			}
		}
	}

	bodyWriter.Close()
	return bodyBuf, bodyWriter, nil
}

func (m *Client) parseSignatureRequestResponse(response *http.Response) (*model.EmbeddedSendCreated, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(response.Status)
	fmt.Println(string(body))

	sigRequestResponse := &model.EmbeddedSendCreated{}

	err = json.NewDecoder(response.Body).Decode(sigRequestResponse)

	return sigRequestResponse, err
}

func (m *Client) boolToIntString(value bool) string {
	if value == true {
		return "true"
	}
	return "false"
}
