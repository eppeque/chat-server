package msg

import "fmt"

func Err(content string) string {
	return fmt.Sprintf("+ERR %v\r\n", content)
}

func Ok(content string) string {
	return fmt.Sprintf("+OK %v\r\n", content)
}

func Msg(sender, content string) string {
	return fmt.Sprintf("+MSG %v %v\r\n", sender, content)
}
