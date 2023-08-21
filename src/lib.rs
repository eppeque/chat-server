mod broadcast;
mod client;
mod message_builder;
mod message_handler;

use std::{
    io,
    net::{IpAddr, TcpListener},
    sync::{Arc, Mutex},
    thread,
};

use broadcast::BroadcastManager;
use client::Client;

pub struct Config {
    ip_addr: IpAddr,
    port: usize,
}

impl Config {
    pub fn build(mut args: impl Iterator<Item = String>) -> Result<Self, &'static str> {
        args.next();

        let ip_addr: IpAddr = match args.next() {
            Some(arg) => match arg.parse() {
                Ok(ip_addr) => ip_addr,
                Err(_) => return Err("The IP address is not valid"),
            },
            None => return Err("Didn't get an IP address"),
        };

        let port: usize = match args.next() {
            Some(arg) => match arg.parse() {
                Ok(port) => port,
                Err(_) => return Err("The port number is not valid"),
            },
            None => return Err("Didn't get a port number"),
        };

        Ok(Self { ip_addr, port })
    }
}

pub fn run(config: Config) -> io::Result<()> {
    let addr = format!("{}:{}", config.ip_addr.to_string(), config.port);
    let listener = TcpListener::bind(&addr)?;
    let manager = BroadcastManager::new();

    println!("Server running at {addr}...");
    listen(listener, manager)?;

    Ok(())
}

fn listen(listener: TcpListener, manager: BroadcastManager) -> io::Result<()> {
    let manager = Arc::new(Mutex::new(manager));
    let mut id_counter = 0;

    for stream in listener.incoming() {
        let client = Arc::new(Client::new(stream?, id_counter));
        manager.lock().unwrap().register_client(client.clone());
        let manager = Arc::clone(&manager);

        create_thread(client, manager);

        id_counter += 1;
    }

    Ok(())
}

fn create_thread(client: Arc<Client>, manager: Arc<Mutex<BroadcastManager>>) {
    thread::spawn(move || {
        println!("[INFO] - Thread created");

        client.handle_connection(&manager);

        let mut manager = manager.lock().unwrap();
        manager.remove_client(client.id());
        println!(
            "[INFO] - Thread terminated. {} thread(s) running...",
            manager.len()
        );
    });
}
