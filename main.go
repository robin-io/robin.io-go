package go_notify

import (
	"encoding/json"
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

var (
	base_url = "https://not.acumendev.xyz/api/v1"
	url      = "not.acumendev.xyz"
	ws       = fmt.Sprintf(`ws://%s/ws`, url)
	wss      = fmt.Sprintf(`wss://%s/ws`, url)
)

type Notify struct {
	Secret string
	Tls    bool
	Conn   *gowebsocket.Socket
}

type Response struct {
	Error    bool        `json:"error"`
	Document Document    `json:"document"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
}

type Document struct {
	ID            string        `json:"_id"`
	Conversations []interface{} `json:"conversations"`
	Data          interface{}   `json:"data"`
	Name          string        `json:"name"`
}

type Message struct {
	Type           int                    `json:"type"`
	Channel        string                 `json:"channel"`
	Content        map[string]interface{} `json:"content"`
	ConversationId string                 `json:"conversation_id"`
}

type Fxn func(gowebsocket.Socket)
type ErrFxn func(error, gowebsocket.Socket)
type MsgFxn func(string, gowebsocket.Socket)
type ByteFxn func([]byte, gowebsocket.Socket)

func (n *Notify) Connect(connected Fxn,
	connect_error, disconnected ErrFxn,
	text_recieved, ping, pong MsgFxn,
	binary ByteFxn) (*gowebsocket.Socket, error) {

	var conn_string string

	if n.Tls {
		conn_string = wss
	} else {
		conn_string = ws
	}
	socket := gowebsocket.New(fmt.Sprintf(`%s/%s`, conn_string, n.Secret))

	socket.OnConnected = connected
	socket.OnConnectError = connect_error
	socket.OnTextMessage = text_recieved
	socket.OnBinaryMessage = binary
	socket.OnPingReceived = ping
	socket.OnPongReceived = pong
	socket.OnDisconnected = disconnected

	socket.Connect()
	n.Conn = &socket
	return &socket, nil
}

func (n *Notify) Subscribe(channel string) error {
	msg := Message{
		Type:    0,
		Channel: channel,
		Content: nil,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	n.Conn.SendText(string(body))

	return nil
}

func (n *Notify) SendMessage(channel string, content map[string]interface{}) error {
	msg := Message{
		Type:    1,
		Channel: channel,
		Content: content,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	n.Conn.SendText(string(body))

	return nil

}
func (n *Notify) SendConversationMessage(channel, conversation_id string, content map[string]interface{}) error {
	msg := Message{
		Type:           1,
		Channel:        channel,
		Content:        content,
		ConversationId: conversation_id,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	n.Conn.SendText(string(body))

	return nil

}
