package service

import (
	"time"
)

type Service struct {
	chats    []Chat
	chatsMap map[int64]Chat
}

func NewService(
	chats []Chat,
) *Service {
	chatsMap := make(map[int64]Chat)

	for _, chat := range chats {
		chatsMap[chat.ID] = chat
	}

	return &Service{
		chats:    chats,
		chatsMap: chatsMap,
	}
}

func (s *Service) ListChats() ([]Chat, error) {
	return s.chats, nil
}

func (s *Service) GetChat(chatID int64) (Chat, error) {
	chat, ok := s.chatsMap[chatID]
	if !ok {
		return Chat{}, ErrChatNotFound
	}

	return chat, nil
}

func (s *Service) SendMessage(chatID int64, content string) error {
	chat, ok := s.chatsMap[chatID]
	if !ok {
		return ErrChatNotFound
	}

	message := ChatMessage{
		ID:       0,
		UserID:   0,
		ChatID:   chatID,
		Content:  content,
		SentTime: time.Now(),
	}

	chat.Messages = append(chat.Messages, message)

	s.chatsMap[chatID] = chat

	return nil
}
