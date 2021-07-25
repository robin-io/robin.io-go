package robin

import (
	"github.com/sacOO7/gowebsocket"
	"time"
)

type Robin struct {
	Secret string
	Tls    bool
	Conn   *gowebsocket.Socket
}

type UserToken struct {
	MetaData map[string]interface{} `json:"meta_data"`
}

type Response struct {
	Error    bool        `json:"error"`
	Msg      string      `json:"msg"`
	Data     UserTokenResponse `json:"data"`
}

type UserTokenResponse struct {
	Id            string                 `json:"id"`
	Conversations []interface{}          `json:"conversations"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	UserToken     string                 `json:"user_token"`
	MetaData      map[string]interface{} `json:"meta_data"`
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
