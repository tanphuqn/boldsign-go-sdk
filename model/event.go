package model

type Event struct {
	Event *Event              `json:"event"`
	Data  *DocumentProperties `json:"data"`
}

type EventData struct {
	Id        string `json:"id"`
	EventType string `json:"eventType"`
}
