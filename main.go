package robin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sacOO7/gowebsocket"
)

var (
	baseUrl = "https://api.robinapp.io/api/v1"
	url     = "api.robinapp.io"
	ws      = fmt.Sprintf(`ws://%s/ws`, url)
	wss     = fmt.Sprintf(`wss://%s/ws`, url)
)

func (r *Robin) GetSession(userToken string) (string, error) {
	body := map[string]string{
		"user_token": userToken,
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/session/get/auth", baseUrl), bytes.NewBuffer(byteBody))

	if err != nil {
		return "", err
	}

	req.Header.Set("x-api-key", r.Secret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("bad status code : %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var sessionResponse SessionResponse

	err = json.Unmarshal(respBody, &sessionResponse)

	if err != nil {
		return "", err
	}

	if sessionResponse.Error {
		return "", errors.New(sessionResponse.Msg)
	}

	r.Session = sessionResponse.Token

	return sessionResponse.Token, nil

}

func (r *Robin) Connect(user_token string, connected Fxn,
	connect_error, disconnected ErrFxn,
	text_recieved, ping, pong MsgFxn,
	binary ByteFxn) (*gowebsocket.Socket, error) {

	var conn_string string

	if r.Tls {
		conn_string = wss
	} else {
		conn_string = ws
	}
	socket := gowebsocket.New(fmt.Sprintf(`%s?auth=%s`, conn_string, r.Session))

	socket.OnConnected = connected
	socket.OnConnectError = connect_error
	socket.OnTextMessage = text_recieved
	socket.OnBinaryMessage = binary
	socket.OnPingReceived = ping
	socket.OnPongReceived = pong
	socket.OnDisconnected = disconnected

	socket.Connect()
	r.Conn = &socket
	return &socket, nil
}

func (r *Robin) Subscribe(channel string) error {
	msg := Message{
		Type:    0,
		Channel: channel,
		Content: nil,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	r.Conn.SendText(string(body))

	return nil
}

func (r *Robin) SendMessage(channel string, content map[string]interface{}) error {
	msg := Message{
		Type:    1,
		Channel: channel,
		Content: content,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	r.Conn.SendText(string(body))

	return nil
}

func (r *Robin) SendMessageToConversation(channel, conversation_id string, content map[string]interface{}, sender_token, sender_name string) error {
	msg := Message{
		Type:           1,
		Channel:        channel,
		Content:        content,
		ConversationId: conversation_id,
		SenderName:     sender_name,
		SenderToken:    sender_token,
	}

	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	r.Conn.SendText(string(body))

	return nil
}
