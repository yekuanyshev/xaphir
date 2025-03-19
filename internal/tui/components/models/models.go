package models

import (
	"time"

	"github.com/yekuanyshev/xaphir/pkg/utils"
)

const (
	MessageStatusUnknown MessageStatus = iota
	MessageStatusSent
	MessageStatusRead
)

type (
	MessageStatus int

	Chat struct {
		ID          int64
		Username    string
		LastMessage *ChatMessage
	}

	ChatMessage struct {
		Content  string
		SentTime time.Time
		IsFromMe bool
		Status   MessageStatus
	}
)

func (cm ChatMessage) IsStatusUnknown() bool { return cm.Status == MessageStatusUnknown }
func (cm ChatMessage) IsStatusSent() bool    { return cm.Status == MessageStatusSent }
func (cm ChatMessage) IsStatusRead() bool    { return cm.Status == MessageStatusRead }

func (cm ChatMessage) FormatSentTime() string {
	if utils.InCurrentDay(cm.SentTime) {
		return cm.SentTime.Format("15:04")
	}

	if utils.InCurrentWeekRange(cm.SentTime) {
		return cm.SentTime.Format("Mon")
	}

	return cm.SentTime.Format("02.01.2006")
}
