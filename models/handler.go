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
	} else {
		res := messages.Err("Not implemented yet")
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
	// TODO: Implement
}
