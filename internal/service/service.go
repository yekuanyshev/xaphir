package service

import (
	"slices"
	"time"
)

type Service struct {
	// ids of chats sorted by last message's sent time
	ids      []int64
	chatsMap map[int64]Chat
}

func NewService(
	chats []Chat,
) *Service {
	ids := make([]int64, 0, len(chats))
	chatsMap := make(map[int64]Chat, len(chats))

	for _, chat := range chats {
		chatsMap[chat.ID] = chat
		ids = append(ids, chat.ID)
	}

	slices.SortFunc(ids, func(i, j int64) int {
		chat1 := chatsMap[i]
		chat2 := chatsMap[j]
		return -chat1.CompareByLastMessageSentTime(chat2)
	})

	return &Service{
		ids:      ids,
		chatsMap: chatsMap,
	}
}

func (s *Service) ListChats() ([]Chat, error) {
	var chats []Chat
	for _, id := range s.ids {
		chats = append(chats, s.chatsMap[id])
	}

	return chats, nil
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
	chat.LastMessage = &message

	s.chatsMap[chatID] = chat
	s.sortIDs(chatID)

	return nil
}

func (s *Service) sortIDs(chatID int64) {
	if len(s.ids) > 0 {
		if s.ids[0] == chatID {
			return
		}
	}

	s.ids = slices.DeleteFunc(s.ids, func(id int64) bool {
		return id == chatID
	})
	s.ids = append([]int64{chatID}, s.ids...)
}
