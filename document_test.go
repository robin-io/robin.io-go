package go_notify

import (
	"fmt"
	"testing"
)

func TestDocumentCreation(t *testing.T) {
	notify := Notify{
		Secret: "NT-AygOqSqOAkTXqBoaxCvyOWarmgthOgLSFVlc",
		Tls:    true,
	}

	doc, err := notify.CreateDocument("elvis-doc")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(doc)
}
