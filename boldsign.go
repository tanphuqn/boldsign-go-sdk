package boldsign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

const (
	baseURL             string = "https://api-eu.boldsign.com/v1/"
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
	// bodyBytes, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// bodyString := string(bodyBytes)
	// fmt.Printf("%+v\n", bodyString)
	data := &model.EmbeddedSendCreated{}
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", data)
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
					file, err := os.Open(path)
					if err != nil {
						panic(err)
					}
					defer file.Close()
					// buffer := make([]byte, 512)
					// _, err = file.Read(buffer)
					// if err != nil {
					// 	fmt.Println(err)
					// }
					// fmt.Println(http.DetectContentType(buffer))
					contentType := "application/pdf"
					formField, err := m.CreateFormFileWithContentType(bodyWriter, "Files", file.Name(), contentType)
					if err != nil {
						return nil, nil, err
					}
					_, err = io.Copy(formField, file)
					if err != nil {
						fmt.Println(err)
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
			// case ReminderSettingsKey:
			// 	formField, err := bodyWriter.CreateFormField("ReminderSettings.ReminderCount")
			// 	if err != nil {
			// 		return nil, nil, err
			// 	}
			// 	formField.Write([]byte(strconv.Itoa(embRequest.ReminderSettings.ReminderCount)))
			// 	formField, err = bodyWriter.CreateFormField("ReminderSettings.ReminderCount")
			// 	if err != nil {
			// 		return nil, nil, err
			// 	}
			// 	formField.Write([]byte(strconv.Itoa(embRequest.ReminderSettings.ReminderCount)))

			}
		case reflect.Int:
			switch fieldTag {
			case ExpiryDaysKey:
				{
					fmt.Println(val.Int())
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

func (m *Client) CreateFormFileWithContentType(w *multipart.Writer, fieldname, filename, contentType string) (io.Writer, error) {
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
