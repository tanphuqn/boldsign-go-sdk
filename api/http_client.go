package boldsign

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	gomime "github.com/cubewise-code/go-mime"
	"github.com/tanphuqn/boldsign-go-sdk/model"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
var extestion = ".pdf,.png,.jpg,.docx"

const (
	baseURL             string = "https://api-eu.boldsign.com/v1/"
	baseURLBeta         string = "https://api-eu.boldsign.com/v1-beta/"
	baseDomain          string = "https://account-eu.boldsign.com"
	FormFieldKey        string = "form_field"
	FileKey             string = "Files"
	SignersKey          string = "Signers"
	RolesKey            string = "Roles"
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

func (m *Client) GetToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "BoldSign.Documents.All BoldSign.Templates.All BoldSign.SenderIdentity.Read BoldSign.SenderIdentity.Create BoldSign.SenderIdentity.Write BoldSign.SenderIdentity.Delete")

	apiUrl := m.getDomain()
	resource := "/connect/token"
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	if req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())); err != nil {
		return "", err
	} else {
		str := fmt.Sprintf("%s:%s", m.ClientID, m.Secret)
		encodedAuth := base64.StdEncoding.EncodeToString([]byte(str))
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedAuth))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if resp, err := m.getHTTPClient().Do(req); err != nil {
			return "", err
		} else {
			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				return "", err
			} else {
				var data map[string]interface{}
				if err := json.Unmarshal(body, &data); err != nil {
					return "", err
				}
				accessToken := data["access_token"].(string)
				// fmt.Println(accessToken)
				return accessToken, nil
			}
		}
	}
}

func (m *Client) get(path string, isBeta bool) (*http.Response, error) {
	token, err := m.GetToken()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(isBeta), path)

	var b bytes.Buffer
	request, _ := http.NewRequest(http.MethodGet, endpoint, &b)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return response, m.getError(response)
	}

	return response, err
}

func (m *Client) post(path string, params *bytes.Buffer, w multipart.Writer, isBeta bool) (*http.Response, error) {
	return m.request(http.MethodPost, path, params, w, isBeta)
}

func (m *Client) delete(path string, isBeta bool) (*http.Response, error) {
	return m.requestJson(http.MethodDelete, path, nil, isBeta)
}

func (m *Client) postJson(path string, jsonData []byte, isBeta bool) (*http.Response, error) {
	return m.requestJson(http.MethodPost, path, jsonData, isBeta)
}

func (m *Client) putJson(path string, jsonData []byte, isBeta bool) (*http.Response, error) {
	return m.requestJson(http.MethodPut, path, jsonData, isBeta)
}

func (m *Client) requestJson(method string, path string, jsonData []byte, isBeta bool) (*http.Response, error) {
	token, err := m.GetToken()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(isBeta), path)
	request, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return response, m.getError(response)
	}

	return response, err
}

func (m *Client) request(method string, path string, params *bytes.Buffer, w multipart.Writer, isBeta bool) (*http.Response, error) {
	token, err := m.GetToken()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(isBeta), path)
	fmt.Println(endpoint, "endpoint")
	request, _ := http.NewRequest(method, endpoint, params)
	request.Header.Add("Content-Type", w.FormDataContentType())
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return response, m.getError(response)
	}

	return response, err
}

func (m *Client) getError(response *http.Response) error {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	return errors.New(fmt.Sprintf(`Error: %d`, response.StatusCode) + "-" + bodyString)
}

func (m *Client) getEndpoint(isBeta bool) string {
	var url string
	if isBeta {
		return baseURLBeta
	}

	if m.BaseURL != "" {
		url = m.BaseURL
	} else {
		url = baseURL
	}
	return url
}

func (m *Client) getDomain() string {
	var url string
	if m.BaseURL != "" {
		url = m.BaseDomain
	} else {
		url = baseDomain
	}
	return url
}

func (m *Client) getHTTPClient() *http.Client {
	var httpClient *http.Client
	if m.HTTPClient != nil {
		httpClient = m.HTTPClient
	} else {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}
	return httpClient
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
			case RolesKey:
				for i, role := range embRequest.GetRoles() {
					formField, err := bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][name]", RolesKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(role.GetName()))

					formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][index]", RolesKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(role.GetIndex())))
					if role.DefaultSignerEmail != "" {
						formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][defaultSignerName]", RolesKey, i))
						if err != nil {
							return nil, nil, err
						}
						formField.Write([]byte(role.GetDefaultSignerName()))
					}
					if role.DefaultSignerEmail != "" {
						formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][defaultSignerEmail]", RolesKey, i))
						if err != nil {
							return nil, nil, err
						}
						formField.Write([]byte(role.GetDefaultSignerEmail()))
					}
					formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][signerOrder]", RolesKey, i))
					if err != nil {
						return nil, nil, err
					}
					formField.Write([]byte(strconv.Itoa(role.GetSignerOrder())))

					if role.SignerType != "" {
						formField, err = bodyWriter.CreateFormField(fmt.Sprintf("%s[%v][signerType]", RolesKey, i))
						if err != nil {
							return nil, nil, err
						}
						formField.Write([]byte(role.GetSignerType()))
					}
				}
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
						return nil, nil, errors.New("error: The BoldSign e-signature application supports files in the following formats: PDF, PNG, JPG, and Docx.")
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
