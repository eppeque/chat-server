use std::{
    io::{prelude::*, BufReader, BufWriter},
    net::TcpStream,
    sync::{Arc, Mutex},
};

use crate::{broadcast::BroadcastManager, message_builder, message_handler};

pub struct Client {
    stream: TcpStream,
    id: usize,
}

impl Client {
    pub fn new(stream: TcpStream, id: usize) -> Self {
        Self { stream, id }
    }

    pub fn handle_connection(&self, bm: &Arc<Mutex<BroadcastManager>>) {
        let mut reader = BufReader::new(&self.stream);
        let mut username: Option<String> = None;

        loop {
            let message = match self.read_message(&mut reader) {
                Some(message) => message,
                None => break,
            };

            print!("[{}] - {message}", self.stream.peer_addr().unwrap());

            match &username {
                Some(username) => self.handle_msgs(&message, username, bm),
                None => username = self.handle_register(&message),
            }
        }
    }

    fn read_message(&self, reader: &mut BufReader<&TcpStream>) -> Option<String> {
        let mut message = String::new();
        match reader.read_line(&mut message) {
            Ok(len) => {
                if len != 0 {
                    Some(message)
                } else {
                    None
                }
            }
            Err(_) => None,
        }
    }

    fn handle_register(&self, message: &str) -> Option<String> {
        match message_handler::register(&message) {
            Some(new_username) => {
                let ok = message_builder::ok("You're now registered");
                self.send_message(&ok);
                Some(String::from(new_username))
            }
            None => {
                let error = message_builder::err("You're not registered");
                self.send_message(&error);
                None
            }
        }
    }

    fn handle_msgs(&self, message: &str, username: &str, bm: &Arc<Mutex<BroadcastManager>>) {
        match message_handler::msgs(&message) {
            Some(content) => {
                let message = message_builder::msg(username, &content);
                bm.lock().unwrap().broadcast_message(self, &message);
            }
            None => {
                let error = message_builder::err("Invalid format");
                self.send_message(&error);
            }
        }
    }

    pub fn send_message(&self, message: &str) {
        let mut writer = BufWriter::new(&self.stream);
        writer.write_all(message.as_bytes()).unwrap();
        writer.flush().unwrap();
    }

    pub fn id(&self) -> usize {
        self.id
    }
}
