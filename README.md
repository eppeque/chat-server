# Chat Server

![Go version badge](https://img.shields.io/badge/Go-1.22.2-blue?logo=go)

This is a robust and encrypted cloud messaging server written in [Go](https://go.dev).

## Grammar

In order to communicate with this server, any client must respect [the defined grammar](./GRAMMAR.md).

## Getting Started

### 1. Download the source code

You clone this repository and download the source code with the following command:

```bash
git clone https://github.com/eppeque/chat-server
```

### 2. Run & build

You can run the development version of the server with the following command:

```bash
go run .
```

And you can build an optimized executable with the following command:

```bash
go build
```

## Usage

Here is how the executable works:

```bash
chat-server -p 3000
```

This command will start the chat server and listen on port `3000` instead of the default `8080` port.

_Enjoy!_
