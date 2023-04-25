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

	conv, err := notify.CreateConversation("elvis", "YFXOKVyKBGvHxuBaqKgDWOhE", "jesse")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conv)
}

func TestGroupConversationCreation(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	conv, err := notify.CreateGroupConversation("Elvis & sons",
		UserToken{UserToken: "YFXOKVyKBGvHxuBaqKgDWOhE"},
		[]UserToken{{UserToken: "YFXOKVyKBGvHxuBaqKgDWOhE"}})

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

func TestRobin_SearchConversation(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	messages, err := notify.SearchConversation("610041ac411c882b47d633db", "hi")

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(messages)
	}
}

func TestRobin_DeleteMessage(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	err := notify.DeleteMessage("60c000df26dcd315e219b0f3")

	if err != nil {
		t.Error(err)
	}
}

func TestRobin_AssignGroupModerator(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	conversation, err := notify.AssignGroupModerator("YFXOKVyKBGvHxuBaqKgDWOhE", "6103ee6628e71d0daf8dcd03")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conversation)
}

func TestRobin_AddGroupParticipants(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	conversation, err := notify.AddGroupParticipants("6103ee6628e71d0daf8dcd03",
		[]UserToken{
			{
				UserToken: "YFXOKVyKBGvHxuBaqKgDWOhE",
			},
		})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conversation)
}

func TestRobin_RemoveGroupParticipant(t *testing.T) {
	notify := Robin{
		Secret: "NT-QuNtKolpzoWLahimkIjGAllEcJwGrymaVxQX",
		Tls:    true,
	}

	conversation, err := notify.RemoveGroupParticipant("YFXOKVyKBGvHxuBaqKgDWOhE", "6103ee6628e71d0daf8dcd03")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conversation)
}
