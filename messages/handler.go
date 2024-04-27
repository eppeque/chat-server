package messages

import (
	"errors"
	"strings"
)

func DetectType(line string) (MessageType, error) {
	if strings.HasPrefix(line, "REGISTER") {
		return Register, nil
	}

	if strings.HasPrefix(line, "LOGIN") {
		return Login, nil
	}

	if strings.HasPrefix(line, "CONFIRM") {
		return Confirm, nil
	}

	return -1, errors.New("the given line doesn't correspond to any type")
}
