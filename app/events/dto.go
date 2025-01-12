package events

type EventAttributes struct {
	Timestamp       int64  `json:"timestamp"`
	EventProperties []byte `json:"event_properties"`
	Datetime        string `json:"datetime"`
	Uuid            string `json:"uuid"`
}

type Response struct {
	Attributes []EventAttributes `json:"data"`
}
