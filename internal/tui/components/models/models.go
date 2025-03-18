package models

import (
	"time"
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
		LastMessage string
	}

	ChatMessage struct {
		Content  string
		SendTime time.Time
		IsFromMe bool
		Status   MessageStatus
	}
)

func (cm ChatMessage) IsStatusUnknown() bool { return cm.Status == MessageStatusUnknown }
func (cm ChatMessage) IsStatusSent() bool    { return cm.Status == MessageStatusSent }
func (cm ChatMessage) IsStatusRead() bool    { return cm.Status == MessageStatusRead }
