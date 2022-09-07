package model

type ErrorResponse struct {
	Errors string `json:"errors"`
	Type   string `json:"type"`
	Title  string `json:"title"`
}
