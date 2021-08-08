package robin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"strings"
)

func (r *Robin) Connect() (*gowebsocket.Socket, error) {

	if len(r.UserToken) == 0 {
		return nil, errors.New("UserToken cannot be empty")
	}

	var connString string

	if r.Tls {
		connString = wss
	} else {
		connString = ws
	}
	socket := gowebsocket.New(fmt.Sprintf(`%s/%s`, connString, r.Secret))

	//socket.OnConnected = connected
	//socket.OnConnectError = connectError
	//socket.OnTextMessage = textReceived
	//socket.OnBinaryMessage = binary
	//socket.OnPingReceived = ping
	//socket.OnPongReceived = pong
	//socket.OnDisconnected = disconnected

	socket.Connect()
	r.Conn = &socket
	return &socket, nil
}

func (r *Robin) Subscribe(channel Channel) error {
	msg := Message{
		Type:    0,
		Channel: channel.PublicName,
		Content: nil,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	r.Conn.SendText(string(body))

	return nil
}

func (r *Robin) SendMessage(channel Channel, content map[string]interface{}) error {
	msg := Message{
		Type:    1,
		Channel: channel.PublicName,
		Content: content,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	r.Conn.SendText(string(body))

	return nil
}
func (r *Robin) SendConversationMessage(channel Channel, conversationId string, content map[string]interface{}) error {
	msg := Message{
		Type:           1,
		Channel:        channel.PublicName,
		Content:        content,
		ConversationId: conversationId,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	r.Conn.SendText(string(body))

	return nil
}

func (r *Robin) CreateChannel(name string) Channel {
	channel := Channel{
		Name: name,
		PublicName: r.Secret + "-" + strings.ReplaceAll(name, " ", ""),
	}
	return channel
}