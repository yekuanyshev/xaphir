package stubs

import (
	_ "embed"
	"encoding/json"

	"github.com/yekuanyshev/xaphir/internal/tui/components/chatlist/item"
)

//go:embed stubs.json
var stubsFile []byte

type Stubs struct {
	Chats []item.Chat `json:"chats"`
}

func Load() (Stubs, error) {
	var stubs Stubs

	err := json.Unmarshal(stubsFile, &stubs)
	if err != nil {
		return stubs, err
	}

	return stubs, nil
}
