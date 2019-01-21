package model

import "time"

// ChatEntry structure contains one entry for a chat line in the history
type ChatEntry struct {
	ID     string     // the internal id for the chat entry
	Author string     // the nick for the autor
	Date   *time.Time // the datetime when the chat was sent
	Body   []byte     // internal data for the chat entry, it usually will be a textline, but it can also can store a photo, emoticons, audio, etc.
}

// ChatHistory contains all the history for a chat
type ChatHistory []ChatEntry

// Chat structure contains the information for a chat
type Chat struct {
	ID             string      // the internal id for the chat
	AdminID        string      // the internal admin id for the chat
	Name           string      // the name for the chat
	Description    string      // provided description for the chat
	CreationDate   *time.Time  // the datetime when the chat was created
	ExpirationDate *time.Time  // the datetime when the chat will expire
	History        ChatHistory // all the history for the chat
}
