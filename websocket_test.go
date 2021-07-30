package robin

import (
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"testing"
)

func TestRobin_Connect(t *testing.T) {
	robin := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	conn, err := robin.Connect()

	if err != nil {
		t.Error(err)
	}

	conn.OnConnected = connected
	conn.OnTextMessage = received

	//subscribe
	chann := robin.CreateChannel("test")
	err = robin.Subscribe(chann)

	if err != nil {
		t.Error(err)
	}

	err = robin.SendMessage(chann, map[string]interface{}{"msg":"hello"})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conn)
}

func connected(soc gowebsocket.Socket){
	fmt.Println("connected", soc)
}

func received(msg string, soc gowebsocket.Socket) {
	fmt.Println(msg)
}

func connectError(err error,soc gowebsocket.Socket) {
	fmt.Println(err, soc)
}
func disconnect(err error,soc gowebsocket.Socket) {
	fmt.Println(err, soc)
}