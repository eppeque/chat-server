package models

type MessageSender interface {
	Id() int

	SendMessage(message string)
}
