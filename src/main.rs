use std::{env, process};

use chat_server::Config;

fn main() {
    let config = Config::build(env::args()).unwrap_or_else(|err| {
        eprintln!("{err}");
        process::exit(1);
    });

    if let Err(e) = chat_server::run(config) {
        eprintln!("Server error: {e}");
        process::exit(1);
    }
}
