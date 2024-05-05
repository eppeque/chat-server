package models

import (
	"errors"
	"strings"

	"github.com/eppeque/chat-server/messages"
)

type MessageHandler struct {
	state  ClientState
	auth   *AuthManager
	sender MessageSender
}

func NewHandler(sender MessageSender) *MessageHandler {
	state := New
	auth := NewAuthManager()

	sender.SendMessage(messages.Hello(auth.challenge))

	return &MessageHandler{state, auth, sender}
}

func (h *MessageHandler) HandleLine(line string) {
	mType, err := h.detectType(line)

	if err != nil {
		h.sendError(err)
		return
	}

	if h.state == New {
		switch mType {
		case messages.Register:
			h.handleRegister(line)
		case messages.Login:
			h.handleLogin(line)
		}
	} else if h.state == AuthProcess && mType == messages.Confirm {
		h.handleConfirm(line)
	} else if h.state == Authenticated {
		switch mType {
		case messages.Create:
			h.handleCreate(line)
		case messages.Join:
			h.handleJoin(line)
		}
	} else {
		res := messages.Err("The message is valid but shouldn't be sent in the current state of the client")
		h.sender.SendMessage(res)
	}
}

func (h *MessageHandler) detectType(line string) (messages.MessageType, error) {
	if strings.HasPrefix(line, "REGISTER") {
		return messages.Register, nil
	}

	if strings.HasPrefix(line, "LOGIN") {
		return messages.Login, nil
	}

	if strings.HasPrefix(line, "CONFIRM") {
		return messages.Confirm, nil
	}

	if strings.HasPrefix(line, "CREATE") {
		return messages.Create, nil
	}

	if strings.HasPrefix(line, "JOIN") {
		return messages.Join, nil
	}

	return -1, errors.New("the given line doesn't correspond to any type")
}

func (h *MessageHandler) sendError(err error) {
	res := messages.Err(err.Error())
	h.sender.SendMessage(res)
}

func (h *MessageHandler) handleRegister(line string) {
	username, salt, hash, err := messages.ScanRegister(line)

	if err != nil {
		h.sendError(err)
		return
	}

	err = h.auth.Register(username, salt, hash)

	if err != nil {
		h.sendError(err)
		return
	}

	h.state = Authenticated
	res := messages.Ok("You're registered")
	h.sender.SendMessage(res)
}

func (h *MessageHandler) handleLogin(line string) {
	username, err := messages.ScanLogin(line)

	if err != nil {
		h.sendError(err)
		return
	}

	salt, err := h.auth.Login(username)

	if err != nil {
		h.sendError(err)
		return
	}

	h.state = AuthProcess
	res := messages.Params(salt)
	h.sender.SendMessage(res)
}

func (h *MessageHandler) handleConfirm(line string) {
	challenge, err := messages.ScanConfirm(line)

	if err != nil {
		h.sendError(err)
		return
	}

	if h.auth.Confirm(challenge) {
		h.state = Authenticated
		res := messages.Ok("Successfull authentication")
		h.sender.SendMessage(res)
		return
	}

	res := messages.Err("The challenge is not correct")
	h.sender.SendMessage(res)
}

func (h *MessageHandler) handleCreate(line string) {
	room, err := messages.ScanCreate(line)

	if err != nil {
		h.sendError(err)
		return
	}

	id := Router.CreateRoom(room, h.sender)
	res := messages.Ok(id)
	h.sender.SendMessage(res)
}

func (h *MessageHandler) handleJoin(line string) {
	// TODO: Implement this method...
}
