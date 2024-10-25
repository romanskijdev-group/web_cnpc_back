package typescore

import "net"

type WEvent struct {
	Error        string  `json:"error"`
	Err          error   `json:"-"`
	Text         string  `json:"text"`
	SystemUserID *string `json:"-"`
	EventType    string  `json:"-"`
	IPAddress    *net.IP `json:"-"`
	Source       *string `json:"-"`
}
