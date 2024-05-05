package messages

type MessageType int

const (
	Register MessageType = 0
	Login    MessageType = 1
	Confirm  MessageType = 2
	Create   MessageType = 3
	Join     MessageType = 4
)
