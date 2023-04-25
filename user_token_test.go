package robin

import (
	"errors"
	"fmt"
	"testing"
)

func TestRobin_GetSession(t *testing.T) {
	robin := Robin{
		Secret: "NT-UAzQwycFjXwvGfeciRyVumTWjfUFCImrRFQH",
		Tls:    true,
	}

	session, err := robin.GetSession("CwJUzaISPWhkoEseoLteRLDn")

	if err != nil {
		t.Error(err)
	}

	if len(session) == 0 {
		t.Error(errors.New("invalid session length, can not be 0"))
	}

	if len(robin.Session) == 0 {
		t.Error(errors.New("invalid session length, can not be 0"))
	}
}

func TestRobin_CreateUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-qBsdCDfPFYQkAKcfxMeNgSXvYTmqakOBVYRr",
		Tls:    true,
	}

	token, err := robin.CreateUserToken(UserToken{MetaData: map[string]interface{}{
		"name": "elvis",
	}})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

func TestRobin_GetUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-qBsdCDfPFYQkAKcfxMeNgSXvYTmqakOBVYRr",
		Tls:    true,
	}

	res, err := robin.GetUserToken()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestRobin_SyncUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	res, err := robin.SyncUserToken(UserToken{
		UserToken: "YFXOKVyKBGvHxuBaqKgDWOhE",
		MetaData: map[string]interface{}{
			"email": "elvis@acumen.com.ng",
		},
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestRobin_UpdateDisplayPhoto(t *testing.T) {
	robin := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	res, err := robin.UpdateDisplayPhoto("https://s3.us-east-2.amazonaws.com/robinapp.io/IMG_7159.jpeg")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestRobin_CheckUserTokenOnlineStatus(t *testing.T) {
	robin := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	res, err := robin.CheckUserTokenOnlineStatus("FefXITDgAeTVrghcOHiimDVB", "wwLpVXwNoZVitOPCOgrRZYBA", "kSSY6DKoMqyNQNoA")

	if err != nil {
		t.Error(err)
	}

	fmt.Println("RESULT", res)
}
