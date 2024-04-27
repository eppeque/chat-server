package models

type ClientState int

const (
	New           ClientState = 0
	AuthProcess   ClientState = 1
	Authenticated ClientState = 2
	Joined        ClientState = 3
)
