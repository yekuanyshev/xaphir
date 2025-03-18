package stubs

import (
	_ "embed"
	"encoding/json"

	"github.com/yekuanyshev/xaphir/internal/service"
)

//go:embed stubs.json
var stubsFile []byte

type Stubs struct {
	Chats []service.Chat
}

func Load() (Stubs, error) {
	var stub Stubs

	err := json.Unmarshal(stubsFile, &stub.Chats)
	if err != nil {
		return stub, err
	}

	for i := range stub.Chats {
		chat := stub.Chats[i]
		messages := chat.Messages
		for j := range messages {
			messages[j].ChatID = chat.ID

			if j == len(messages)-1 {
				stub.Chats[i].LastMessage = &messages[j]
			}
		}
	}

	return stub, nil
}
