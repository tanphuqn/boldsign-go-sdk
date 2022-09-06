package boldsign

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

func (m *Client) GetToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "BoldSign.Documents.All BoldSign.Templates.All")

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
				fmt.Println(accessToken)
				return accessToken, nil
			}
		}
	}
}

func (m *Client) get(path string) (*http.Response, error) {
	token, err := m.GetToken()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)

	var b bytes.Buffer
	request, _ := http.NewRequest(http.MethodGet, endpoint, &b)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (m *Client) post(path string, params *bytes.Buffer, w multipart.Writer) (*http.Response, error) {
	return m.request(http.MethodPost, path, params, w)
}

func (m *Client) request(method string, path string, params *bytes.Buffer, w multipart.Writer) (*http.Response, error) {
	token, err := m.GetToken()
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)
	request, _ := http.NewRequest(method, endpoint, params)
	fmt.Println(endpoint)
	fmt.Println(w.FormDataContentType())
	request.Header.Set("Content-Type", w.FormDataContentType())
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(resBody, &data); err != nil {
		return nil, err
	}
	fmt.Println(data)

	// if response.StatusCode >= 400 {
	// 	msg := fmt.Sprintf("boldsign request failed with status %d", response.StatusCode)
	// 	body, _ := ioutil.ReadAll(response.Body)
	// 	var data map[string]interface{}
	// 	if err := json.Unmarshal(body, &data); err != nil {
	// 		return nil, err
	// 	}

	// 	fmt.Println(data)
	// 	e := &model.ErrorResponse{}
	// 	json.NewDecoder(response.Body).Decode(e)
	// 	if e.Error != nil {
	// 		fmt.Println(err.Error())
	// 		msg = fmt.Sprintf("%s: %s", e.Error.Name, e.Error.Message)
	// 	} else {
	// 		messages := []string{}
	// 		for _, w := range e.Warnings {
	// 			messages = append(messages, fmt.Sprintf("%s: %s", w.Name, w.Message))
	// 		}
	// 		msg = strings.Join(messages, ", ")
	// 	}
	// 	fmt.Println(1)
	// 	fmt.Println(msg)
	// 	return response, errors.New(msg)
	// }
	// fmt.Println(2)
	return response, err
}

func (m *Client) nakedGet(path string) (*http.Response, error) {
	token, err := m.GetToken()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)
	var b bytes.Buffer
	request, _ := http.NewRequest(http.MethodGet, endpoint, &b)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		msg := fmt.Sprintf("boldsign request failed with status %d", response.StatusCode)
		e := &model.ErrorResponse{}
		json.NewDecoder(response.Body).Decode(e)
		if e.Error != nil {
			msg = fmt.Sprintf("%s: %s", e.Error.Name, e.Error.Message)
		} else {
			messages := []string{}
			for _, w := range e.Warnings {
				messages = append(messages, fmt.Sprintf("%s: %s", w.Name, w.Message))
			}
			msg = strings.Join(messages, ", ")
		}

		return response, errors.New(msg)
	}

	return response, err
}

func (m *Client) nakedPost(path string) (*http.Response, error) {
	token, err := m.GetToken()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s%s", m.getEndpoint(), path)
	var b bytes.Buffer
	request, _ := http.NewRequest(http.MethodPost, endpoint, &b)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := m.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (m *Client) getEndpoint() string {
	var url string
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
