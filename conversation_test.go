package robin

import (
	"fmt"
	"testing"
)

func TestConversationCreation(t *testing.T) {
	notify := Robin{
		Secret: "NT-AygOqSqOAkTXqBoaxCvyOWarmgthOgLSFVlc",
		Tls:    true,
	}

	conv, err := notify.CreateConversation("elvis", "elvis-doc", "elvis-doc", "jesse")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conv)
}
