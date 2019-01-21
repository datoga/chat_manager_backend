package mongodb

import (
	"time"

	"github.com/datoga/chat_manager_backend/model"
)

// ExpirationDays is the number of days for the duration of a chat
const ExpirationDays = 30

// ChatManagerMongoDB is the struct containing data access
type ChatManagerMongoDB struct {
}

// NewChatManagerMongoDB is the constructor
func NewChatManagerMongoDB() *ChatManagerMongoDB {
	return &ChatManagerMongoDB{}
}

// NewChat creation on the database
func (manager *ChatManagerMongoDB) NewChat(chatname string, description string) (*model.Chat, error) {
	now := time.Now()

	expirationDate := now.Add(time.Duration(ExpirationDays) * time.Hour * 24)

	adminID, err := manager.getNewAdminIDFromDatabase()

	if err != nil {
		return nil, err
	}

	chat := model.Chat{
		AdminID:        adminID,
		CreationDate:   &now,
		Description:    description,
		ExpirationDate: &expirationDate,
		History:        model.ChatHistory{},
	}

	return &chat, nil
}

func (manager *ChatManagerMongoDB) getNewAdminIDFromDatabase() (string, error) {
	return "1", nil
}
