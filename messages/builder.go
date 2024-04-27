package messages

import "fmt"

func Hello(challenge string) string {
	return fmt.Sprintf("HELLO %s\r\n", challenge)
}

func Err(message string) string {
	return fmt.Sprintf("-ERR %s\r\n", message)
}
