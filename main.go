package robin

import (
	"fmt"
	"github.com/sacOO7/gowebsocket"
)

var (
	baseUrl = "https://robbin-api.herokuapp.com/api/v1"
	url     = "robbin-api.herokuapp.com"
	ws      = fmt.Sprintf(`ws://%s/ws`, url)
	wss     = fmt.Sprintf(`wss://%s/ws`, url)
)

type Fxn func(gowebsocket.Socket)
type ErrFxn func(error, gowebsocket.Socket)
type MsgFxn func(string, gowebsocket.Socket)
type ByteFxn func([]byte, gowebsocket.Socket)




