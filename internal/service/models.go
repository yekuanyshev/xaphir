package service

import (
	"time"

	"github.com/yekuanyshev/xaphir/internal/tui/components/models"
)

type (
	User struct {
		ID       int64
		Username string
	}

	Chat struct {
		ID          int64
		CreatedAt   time.Time
		Member      User
		LastMessage *ChatMessage
		Messages    []ChatMessage
	}

	ChatMessage struct {
		ID       int64
		UserID   int64
		ChatID   int64
		Content  string
		SentTime time.Time
	}
)

func (c Chat) ToComponentModel() models.Chat {
	var lastMessage *models.ChatMessage
	if c.LastMessage != nil {
		lastMessage = &models.ChatMessage{
			Content:  c.LastMessage.Content,
			SendTime: c.LastMessage.SentTime,
			IsFromMe: false,                       // todo: after authorization module
			Status:   models.MessageStatusUnknown, // todo: after message status feature
		}
	}

	return models.Chat{
		ID:          c.ID,
		Username:    c.Member.Username,
		LastMessage: lastMessage,
	}
}

func (cm ChatMessage) ToComponentModel() models.ChatMessage {
	return models.ChatMessage{
		Content:  cm.Content,
		SendTime: cm.SentTime,
		IsFromMe: false,
		Status:   models.MessageStatusUnknown,
	}
}
