package robin

import (
	"fmt"
	"testing"
)

func TestConversationCreation(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	conv, err := notify.CreateConversation("elvis", "YFXOKVyKBGvHxuBaqKgDWOhE", "YFXOKVyKBGvHxuBaqKgDWOhE", "jesse")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conv)
}
