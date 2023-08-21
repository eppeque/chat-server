pub fn ok(message: &str) -> String {
    format!("+OK {message}\r\n")
}

pub fn err(message: &str) -> String {
    format!("+ERR {message}\r\n")
}

pub fn msg(sender_username: &str, content: &str) -> String {
    format!("+MSG {sender_username} {content}\r\n")
}
