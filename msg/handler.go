package msg

import (
	"errors"
	"regexp"
)

const (
	register string = `(\+REGISTER) ([a-zA-Z0-9]{5,15}) ([a-zA-Z0-9]+)`
	msgs     string = `(\+MSGS) ([\w\W]+)`
)

var errFmt error = errors.New("invalid format")

func Register(msg string) (username, room string, err error) {
	regex := regexp.MustCompile(register)
	groups := regex.FindStringSubmatch(msg)

	if groups == nil {
		return "", "", errFmt
	}

	username = groups[2]
	room = groups[3]
	return
}

func Msgs(msg string) (content string, err error) {
	regex := regexp.MustCompile(msgs)
	groups := regex.FindStringSubmatch(msg)

	if groups == nil {
		return "", errFmt
	}

	content = groups[2]
	return
}
