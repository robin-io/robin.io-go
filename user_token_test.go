package robin

import (
	"fmt"
	"testing"
)

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

	res, err := robin.GetUserToken("YFXOKVyKBGvHxuBaqKgDWOhE")

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
