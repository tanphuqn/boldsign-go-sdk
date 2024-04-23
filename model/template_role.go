package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type TemplateRole struct {
	Name               string `json:"name"`
	Index              int    `json:"index"`
	DefaultSignerName  string `json:"defaultSignerName"`
	DefaultSignerEmail string `json:"defaultSignerEmail"`
	SignerOrder        int    `json:"signerOrder"`
	SignerType         string `json:"signerType"`
}

func (s *TemplateRole) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

func (s *TemplateRole) GetIndex() int {
	if s != nil {
		return s.Index
	}
	return 1
}

func (s *TemplateRole) GetDefaultSignerName() string {
	if s != nil {
		return s.DefaultSignerName
	}
	return ""
}

func (s *TemplateRole) GetDefaultSignerEmail() string {
	if s != nil {
		return s.DefaultSignerEmail
	}
	return ""
}
func (s *TemplateRole) GetSignerOrder() int {
	if s != nil {
		return s.SignerOrder
	}
	return 1
}

func (s *TemplateRole) GetSignerType() string {
	if s != nil {
		return s.SignerType
	}
	return ""
}

// Implement a custom scanner for TemplateRole slice
func (tr *TemplateRole) Scan(value interface{}) error {
	// Check if the value is nil or empty
	if value == nil {
		return nil
	}

	// Convert the JSONB data to bytes
	data, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan TemplateRole: value is not []byte")
	}

	// Unmarshal the JSON data into the TemplateRole slice
	return json.Unmarshal(data, &tr)
}

// Implement a custom valuer for TemplateRole slice
func (tr TemplateRole) Value() (driver.Value, error) {
	// Marshal the TemplateRole slice into JSON data
	return json.Marshal(tr)
}
