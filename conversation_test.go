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


func TestRobin_GetConversationMessages(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	messages, err := notify.GetConversationMessages("610041ac411c882b47d633db")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(messages)
}