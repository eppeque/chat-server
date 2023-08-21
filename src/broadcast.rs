use std::sync::Arc;

use crate::client::Client;

pub struct BroadcastManager {
    clients: Vec<Arc<Client>>,
}

impl BroadcastManager {
    pub fn new() -> Self {
        Self {
            clients: Vec::new(),
        }
    }

    pub fn register_client(&mut self, client: Arc<Client>) {
        self.clients.push(client);
    }

    pub fn remove_client(&mut self, id: usize) {
        if let Some(index) = self.clients.iter().position(|client| client.id() == id) {
            self.clients.remove(index);
        }
    }

    pub fn len(&self) -> usize {
        self.clients.len()
    }

    pub fn broadcast_message(&self, sender: &Client, message: &str) {
        self.clients
            .iter()
            .filter(|client| client.id() != sender.id())
            .for_each(|client| client.send_message(message));
    }
}
