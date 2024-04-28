# Chat Server

![Go version badge](https://img.shields.io/badge/Go-1.22.2-blue?logo=go)

This is a robust and encrypted cloud messaging server written in [Go](https://go.dev).

## Grammar

In order to communicate with this server, any client must respect [the defined grammar](./GRAMMAR.md).

## Getting Started

### 1. Prerequisites

Since this server is written in Go, you'll need the [Go toolchain](https://go.dev/dl) to build and run this project.

### 2. Build

You can download the source code and build the project with a single command:

```bash
go install https://github.com/eppeque/chat-server
```

### 3. Run

Run the installed executable with the following command:

```bash
chat-server
```

## Usage

Here is how the executable works:

```bash
chat-server -p 3000
```

This command will start the chat server and listen on port `3000` instead of the default `8080` port.

_Enjoy!_
