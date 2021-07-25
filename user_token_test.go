package robin

import (
	"fmt"
	"testing"
)

func TestCreateUserToken(t *testing.T) {
	robin := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
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
