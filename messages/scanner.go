package messages

import (
	"errors"
	"regexp"
)

func ScanRegister(line string) (string, string, string, error) {
	re := regexp.MustCompile(RxRegister)
	groups := re.FindStringSubmatch(line)

	if groups == nil {
		return "", "", "", errors.New("invalid format")
	}

	return groups[1], groups[2], groups[3], nil
}

func ScanLogin(line string) (string, error) {
	re := regexp.MustCompile(RxLogin)
	groups := re.FindStringSubmatch(line)

	if groups == nil {
		return "", errors.New("invalid format")
	}

	return groups[1], nil
}
