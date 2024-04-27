package messages

type MessageType int

const (
	Register MessageType = 0
	Login    MessageType = 1
	Confirm  MessageType = 2
)
