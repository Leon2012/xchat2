package protocol

import (
	"net/http"
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

type MsgServerCtrl struct {
	Id     string      `json:"id,omitempty"`
	Topic  string      `json:"topic,omitempty"`
	Params interface{} `json:"params,omitempty"`

	Code      int       `json:"code"`
	Text      string    `json:"text,omitempty"`
	Timestamp time.Time `json:"ts"`
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

type MsgServerData struct {
	Topic string `json:"topic"`
	// ID of the user who originated the message as {pub}, could be empty if sent by the system
	From      string      `json:"from,omitempty"`
	Timestamp time.Time   `json:"ts"`
	SeqId     int         `json:"seq"`
	Content   interface{} `json:"content"`
}

type ServerComMessage struct {
	Ctrl *MsgServerCtrl `json:"ctrl,omitempty"`
	Info *MsgServerInfo `json:"info,omitempty"`
	Data *MsgServerData `json:"data,omitempty"`
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
	From      int64
	Timestamp time.Time
	Hi        *MsgClientHi `json:"hi"`
}

/////////////////////////////////////////////////////////////errors////////////////////////////////////////////////////////////////////////////////////

func NoErrShutdown(ts time.Time) *ServerComMessage {
	msg := &ServerComMessage{Ctrl: &MsgServerCtrl{
		Code:      http.StatusResetContent, // 205
		Text:      "server shutdown",
		Timestamp: ts}}
	return msg
}

//3xx error

// 4xx error
func ErrMalformed(id, topic string, ts time.Time) *ServerComMessage {
	msg := &ServerComMessage{Ctrl: &MsgServerCtrl{
		Id:        id,
		Code:      http.StatusBadRequest, // 400
		Text:      "malformed",
		Topic:     topic,
		Timestamp: ts}}
	return msg
}
