package robin

import (
	"time"

	"github.com/sacOO7/gowebsocket"
)

type Robin struct {
	Secret string
	Tls    bool
	Conn   *gowebsocket.Socket
}

type UserToken struct {
	UserToken string                 `json:"user_token"`
	MetaData  map[string]interface{} `json:"meta_data"`
}

type Response struct {
	Error bool              `json:"error"`
	Msg   string            `json:"msg"`
	Data  UserTokenResponse `json:"data"`
}

type ConversationResponse struct {
	Error            bool                     `json:"error"`
	Msg              string                   `json:"msg"`
	ConversationData ConversationResponseData `json:"data"`
}

type MessageResponse struct {
	Error       bool                  `json:"error"`
	Msg         string                `json:"msg"`
	MessageData []MessageResponseData `json:"data"`
}

type ConversationResponseData struct {
	Id            string        `json:"_id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Data          interface{}   `json:"data"`
	IsGroup       bool          `json:"is_group"`
	Moderator     Participant   `json:"moderator"`
	Name          string        `json:"name"`
	Participants  []Participant `json:"participants"`
	ReceiverName  string        `json:"receiver_name"`
	ReceiverToken string        `json:"receiver_token"`
	SenderName    string        `json:"sender_name"`
	SenderToken   string        `json:"sender_token"`
}

type Participant struct {
	UserToken   string                 `json:"user_token"`
	IsModerator bool                   `json:"is_moderator"`
	MetaData    map[string]interface{} `json:"meta_data"`
}

type UserTokenResponse struct {
	Id            string                     `json:"_id"`
	Conversations []ConversationResponseData `json:"conversations"`
	CreatedAt     time.Time                  `json:"created_at"`
	UpdatedAt     time.Time                  `json:"updated_at"`
	UserToken     string                     `json:"user_token"`
	MetaData      map[string]interface{}     `json:"meta_data"`
	DisplayPhoto  string                     `json:"display_photo"`
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

type MessageResponseData struct {
	Id             string                 `json:"_id"`
	TimeStamp      time.Time              `json:"timestamp"`
	OwnerId        string                 `json:"owner_id"`
	OwnerName      string                 `json:"owner_name"`
	OwnerEmail     string                 `json:"owner_email"`
	Channel        string                 `json:"channel"`
	Content        map[string]interface{} `json:"content"`
	AppId          string                 `json:"app_id"`
	AppName        string                 `json:"app_name"`
	Tokens         []string               `json:"tokens"`
	Title          string                 `json:"title"`
	SubTitle       string                 `json:"subtitle"`
	Message        interface{}            `json:"message"`
	Platform       float64                `json:"platform"`
	Data           map[string]interface{} `json:"data"`
	Image          string                 `json:"image"`
	Topic          string                 `json:"topic"`
	Priority       string                 `json:"priority"`
	ConversationId string                 `json:"conversation_id,omitempty"`
	Notify         bool                   `json:"notify"`
	ReceiverId     string                 `json:"receiver_id"`
	IsRead         bool                   `json:"is_read"`
	IsDeleted      bool                   `json:"is_deleted"`
}

type UserTokenStatusResponse struct {
	Error bool              `json:"error"`
	Data  map[string]string `json:"data"`
	Msg   string            `json:"msg"`
}
