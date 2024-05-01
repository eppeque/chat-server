package models

type MessageSender interface {
	SendMessage(message string)
}
