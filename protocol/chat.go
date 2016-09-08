package protocol

import (
	"strings"
	"time"
)

type JsonDuration time.Duration

func (jd *JsonDuration) UnmarshalJSON(data []byte) (err error) {
	d, err := time.ParseDuration(strings.Trim(string(data), "\""))
	*jd = JsonDuration(d)
	return err
}

////////////////////////////////////////////////////////////server//////////////////////////////////////////////////////////////////////////////////
type ServerComMessage struct {
}

type MsgServerInfo struct {
	Topic string `json:"topic"`
	// ID of the user who originated the message
	From string `json:"from"`
	// what is being reported: "rcpt" - message received, "read" - message read, "kp" - typing notification
	What string `json:"what"`
	// Server-issued message ID being reported
	SeqId int `json:"seq,omitempty"`
}

/////////////////////////////////////////////////////////////client/////////////////////////////////////////////////////////////////////////////////
// Handshake {hi} message
type MsgClientHi struct {
	// Message Id
	Id string `json:"id,omitempty"`
	// User agent
	UserAgent string `json:"ua,omitempty"`
	// Authentication scheme
	Version string `json:"ver,omitempty"`
	// Client's unique device ID
	DeviceID string `json:"dev,omitempty"`
}

type ClientComMessage struct {
	// from: userid as string
	from      string
	timestamp time.Time
	Hi        *MsgClientHi `json:"hi"`
}
