package api

import "github.com/datoga/chat_manager_backend/model"

// ChatManager contains all the chat managing API that will be exported by the internal services
type ChatManager interface {
	NewChat(chatname string, description string) (*model.Chat, error)
	GetChat(id string) (*model.Chat, error)
	RenewChat(id string, adminid string) error
	DeleteChat(id string, adminid string) error
}

// ChatHistoryWriter contains all the chat live API that will be exported by the internal services
type ChatHistoryWriter interface {
	WriteEntry(author string, body []byte) error
}

// ChatHistoryQuerier contains all the history chat API that will be exported by the internal services
type ChatHistoryQuerier interface {
	LastChatEntries(id string, entries int) (model.ChatHistory, error)
	QueryChatEntries(id string, entryMin int, entryMax int) (model.ChatHistory, error)
}
