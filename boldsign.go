package boldsign

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	gomime "github.com/cubewise-code/go-mime"
	"github.com/tanphuqn/boldsign-go-sdk/model"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
var extestion = ".pdf,.png,.jpg,.docx"

const (
	baseURL             string = "https://api-eu.boldsign.com/v1-beta/"
	baseDomain          string = "https://account-eu.boldsign.com"
	FormFieldKey        string = "form_field"
	FileKey             string = "Files"
	SignersKey          string = "Signers"
	ReminderSettingsKey string = "ReminderSettings"
	ExpiryDaysKey       string = "ExpiryDays"
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
	bodyBuf, bodyWriter, err := m.MarshalMultipartEmbeddedSignatureRequest(req)
	if err != nil {
		fmt.Println("marshalMultipartEmbeddedSignatureRequest Error:", err.Error())
		return nil, err
	}

	response, err := m.post("document/createEmbeddedRequestUrl", bodyBuf, *bodyWriter)
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
	response, err := m.get(path)
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
	response, err := m.get(path)
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

func (m *Client) DownloadDocument(documentId string) ([]byte, error) {
	path := fmt.Sprintf("document/download?documentId=%s", documentId)
	response, err := m.get(path)
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

func (m *Client) MarshalMultipartEmbeddedSignatureRequest(embRequest model.EmbeddedDocumentRequest) (*bytes.Buffer, *multipart.Writer, error) {
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

					formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][SignerOrder]", SignersKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(signer.GetSignerOrder())))
				}
			case FileKey:
				for _, path := range embRequest.GetFiles() {
					// https://www.boldsign.com/help/prepare-document/what-are-the-supported-file-formats-and-file-sizes/
					ext := filepath.Ext(path)
					fmt.Println("ext2:", ext)
					if !strings.Contains(extestion, ext) {
						return nil, nil, errors.New("Error: The BoldSign e-signature application supports files in the following formats: PDF, PNG, JPG, and Docx.")
					}
					file, err := os.Open(path)
					if err != nil {
						return nil, nil, err
					}
					defer file.Close()
					fileName := filepath.Base(path)
					formField, err := m.CreateFormFileWithContentType(bodyWriter, "Files", fileName, path)
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
					if err != nil {
						return nil, nil, err
					}
				}
			}
		case reflect.Bool:
			formField, err := bodyWriter.CreateFormField(fieldTag)
			if err != nil {
				return nil, nil, err
			}
			formField.Write([]byte(m.BoolToIntString(val.Bool())))
		case reflect.Struct:
			switch fieldTag {
			case ReminderSettingsKey:
				if embRequest.ReminderSettings.ReminderCount != 0 {
					formField, err := bodyWriter.CreateFormField("ReminderSettings.ReminderCount")
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(embRequest.ReminderSettings.ReminderCount)))
				}

				if embRequest.ReminderSettings.ReminderCount != 0 {
					formField, err := bodyWriter.CreateFormField("ReminderSettings.ReminderCount")
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(embRequest.ReminderSettings.ReminderCount)))
				}
			}
		case reflect.Int:
			switch fieldTag {
			case ExpiryDaysKey:
				{
					if val.Int() != 0 {
						formField, err := bodyWriter.CreateFormField(fieldTag)
						if err != nil {
							return nil, nil, err
						}
						formField.Write([]byte(strconv.Itoa(int(val.Int()))))
					}
				}
			default:
				{
					formField, err := bodyWriter.CreateFormField(fieldTag)
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(int(val.Int()))))
				}
			}
		default:
			if val.String() != "" {
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

func (m *Client) BoolToIntString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func (m *Client) CreateFormFileWithContentType(w *multipart.Writer, fieldname, filename, path string) (io.Writer, error) {
	// Get the file extension
	ext := filepath.Ext(filename)
	// An empty string is returned if the extension is not found
	contentType := gomime.TypeByExtension(ext)
	fmt.Println("contentType:", contentType)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			m.EscapeQuotes(fieldname), m.EscapeQuotes(filename)))
	h.Set("Content-Type", contentType)

	return w.CreatePart(h)
}

func (m *Client) EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// CreateSenderIdentity creates sender identity
func (m *Client) CreateSenderIdentity(senderRequest model.SenderCreateRequest) (*model.SenderCreated, error) {
	jsonData, _ := json.Marshal(senderRequest)
	response, err := m.postJson("senderIdentities/create", jsonData)

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
	response, err := m.postJson(path, jsonData)
	defer response.Body.Close()
	return err
}

// DeleteSenderIdentity
func (m *Client) DeleteSenderIdentity(email string) error {
	path := fmt.Sprintf("senderIdentities/delete?email=%s", email)
	response, err := m.delete(path)
	defer response.Body.Close()
	return err
}

// DeleteSenderIdentity
func (m *Client) VerifySenderIdentity(email string) (model.SenderIdentityDetail, bool, error) {
	path := fmt.Sprintf("senderIdentities/list?PageSize=1&Page=1&Search=%s", email)
	response, err := m.get(path)

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
