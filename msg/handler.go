package msg

import (
	"errors"
	"regexp"
)

const (
	register string = `(\+REGISTER) ([a-zA-Z0-9]{5,15}) ([a-zA-Z0-9]+)`
)

func Register(msg string) (username, room string, err error) {
	regex := regexp.MustCompile(register)
	groups := regex.FindStringSubmatch(msg)

	if groups == nil {
		return "", "", errors.New("invalid format")
	}

	username = groups[2]
	room = groups[3]
	return
}
