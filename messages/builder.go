package messages

import "fmt"

func Hello(challenge string) string {
	return fmt.Sprintf("HELLO %s\r\n", challenge)
}

func Params(salt string) string {
	return fmt.Sprintf("PARAMS %s\r\n", salt)
}

func Err(message string) string {
	return fmt.Sprintf("-ERR %s\r\n", message)
}

func Ok(message string) string {
	return fmt.Sprintf("+OK %s\r\n", message)
}
