package service

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	chats, expectedIDs, expectedChatsMap := generateChats()
	srv := NewService(chats)
	assert.Equal(t, expectedIDs, srv.ids)
	assert.Equal(t, expectedChatsMap, srv.chatsMap)

	chats, expectedIDs, expectedChatsMap = generateChatsDesc()
	srv = NewService(chats)
	assert.Equal(t, expectedIDs, srv.ids)
	assert.Equal(t, expectedChatsMap, srv.chatsMap)

	chats, expectedIDs, expectedChatsMap = generateChatsGaps()
	srv = NewService(chats)
	assert.Equal(t, expectedIDs, srv.ids)
	assert.Equal(t, expectedChatsMap, srv.chatsMap)

	chats, expectedIDs, expectedChatsMap = generateChatsGapsDesc()
	srv = NewService(chats)
	assert.Equal(t, expectedIDs, srv.ids)
	assert.Equal(t, expectedChatsMap, srv.chatsMap)
}

func TestListChats(t *testing.T) {
	chats, expectedIDs, expectedChatsMap := generateChats()
	got, err := NewService(chats).ListChats()
	assert.NoError(t, err)
	assert.Equal(t, sortedChats(expectedIDs, expectedChatsMap), got)

	chats, expectedIDs, expectedChatsMap = generateChatsDesc()
	got, err = NewService(chats).ListChats()
	assert.NoError(t, err)
	assert.Equal(t, sortedChats(expectedIDs, expectedChatsMap), got)

	chats, expectedIDs, expectedChatsMap = generateChatsGaps()
	got, err = NewService(chats).ListChats()
	assert.NoError(t, err)
	assert.Equal(t, sortedChats(expectedIDs, expectedChatsMap), got)

	chats, expectedIDs, expectedChatsMap = generateChatsGapsDesc()
	got, err = NewService(chats).ListChats()
	assert.NoError(t, err)
	assert.Equal(t, sortedChats(expectedIDs, expectedChatsMap), got)
}

func TestGetChat(t *testing.T) {
	chats, expectedIDs, expectedChatsMap := generateChats()
	srv := NewService(chats)

	for _, id := range expectedIDs {
		expectedChat := expectedChatsMap[id]
		got, err := srv.GetChat(id)
		assert.NoError(t, err)
		assert.Equal(t, expectedChat, got)
	}

	_, err := srv.GetChat(-1)
	assert.ErrorIs(t, err, ErrChatNotFound)
}

func TestSendMessage(t *testing.T) {
	chats, expectedIDs, _ := generateChats()
	srv := NewService(chats)

	for i, id := range expectedIDs {
		content := strconv.Itoa(i)
		err := srv.SendMessage(id, content)
		assert.NoError(t, err)

		chat, err := srv.GetChat(id)
		assert.NoError(t, err)

		assert.Equal(t, content, chat.Messages[len(chat.Messages)-1].Content)

		chats, err := srv.ListChats()
		assert.NoError(t, err)
		assert.Equal(t, chat, chats[0])
	}

	err := srv.SendMessage(-1, "")
	assert.ErrorIs(t, err, ErrChatNotFound)
}

func sortedChats(ids []int64, chatsMap map[int64]Chat) []Chat {
	var chats []Chat

	for _, id := range ids {
		chats = append(chats, chatsMap[id])
	}

	return chats
}

func generateChats() (chats []Chat, ids []int64, m map[int64]Chat) {
	m = make(map[int64]Chat)

	for i := range 10 {
		id := int64(i)
		chat := Chat{
			ID: id,
			LastMessage: &ChatMessage{
				SentTime: time.Now(),
			},
		}
		chats = append(chats, chat)
		ids = append([]int64{id}, ids...)
		m[id] = chat
	}

	return
}

func generateChatsDesc() (chats []Chat, ids []int64, m map[int64]Chat) {
	m = make(map[int64]Chat)

	now := time.Now()
	for i := range 10 {
		id := int64(i)
		chat := Chat{
			ID: id,
			LastMessage: &ChatMessage{
				SentTime: now,
			},
		}
		chats = append(chats, chat)
		ids = append(ids, id)
		m[id] = chat
		now = now.Add(-time.Second)
	}

	return
}

func generateChatsGaps() (chats []Chat, ids []int64, m map[int64]Chat) {
	m = make(map[int64]Chat)
	now := time.Now()

	for i := range 10 {
		var lastMessage *ChatMessage
		random := rand.Intn(2)
		id := int64(i)
		if random == 0 {
			lastMessage = &ChatMessage{
				SentTime: now,
			}
			ids = append([]int64{id}, ids...)
		} else {
			ids = append(ids, id)
		}

		chat := Chat{
			ID:          id,
			LastMessage: lastMessage,
		}
		chats = append(chats, chat)
		m[id] = chat
		now = now.Add(time.Second)
	}

	return
}

func generateChatsGapsDesc() (chats []Chat, ids []int64, m map[int64]Chat) {
	m = make(map[int64]Chat)
	now := time.Now()

	var emptyIDs []int64

	for i := range 10 {
		var lastMessage *ChatMessage
		random := rand.Intn(2)
		id := int64(i)
		if random == 0 {
			lastMessage = &ChatMessage{
				SentTime: now,
			}
			ids = append(ids, id)
		} else {
			emptyIDs = append(emptyIDs, id)
		}

		chat := Chat{
			ID:          id,
			LastMessage: lastMessage,
		}
		chats = append(chats, chat)
		m[id] = chat
		now = now.Add(-time.Second)
	}

	ids = append(ids, emptyIDs...)

	return
}
