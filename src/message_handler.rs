use regex::Regex;

const REGISTER_REGEX: &str = r"(\+REGISTER) ([\w]+)(\r\n)";
const MSGS_REGEX: &str = r"(\+MSGS) ([\w\W\s]+)(\r\n)";

pub fn register(message: &str) -> Option<String> {
    let re = Regex::new(REGISTER_REGEX).unwrap();

    match re.captures(&message) {
        Some(groups) => Some(String::from(&groups[2])),
        None => None,
    }
}

pub fn msgs(message: &str) -> Option<String> {
    let re = Regex::new(MSGS_REGEX).unwrap();

    match re.captures(&message) {
        Some(groups) => Some(String::from(&groups[2])),
        None => None,
    }
}
