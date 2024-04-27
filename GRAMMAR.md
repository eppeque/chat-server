# Grammar definitions

Here are the different types of messages that are understood by this server:

## Table of contents

- [Authentication](#authentication)
- [Room navigation](#room-navigation)
- [Message exchange](#message-exchange)
- [Status](#status)

## Authentication

These messages will authenticate the user.

### HELLO

The client receives this message directly when it connects to the server. It contains 22 random characters which will be used for the authentication by challenge. Here is the format:

```
HELLO <22_RAND_CHARS>\r\n
```

### REGISTER

This message will create an account for the user. Here is the format:

```
REGISTER <USERNAME> <BCRYPT_HASH>\r\n
```

### LOGIN

This message will start the authentication process. You just need to give the username here:

```
LOGIN <USERNAME>\r\n
```

The server will respond with a `PARAMS` message.

### PARAMS

This message will give the bcrypt salt to allow the client to build the challenge:

```
PARAMS <BCRYPT_SALT>\r\n
```

### CONFIRM

This is the final message of the authentication process. The client will send the challenge to the server. If the challenge is passed, the client is authenticated. Otherwise, the authentication fails.

Here is what the client needs to do to build a valid challenge:

1. Place the 22 random characters received in the `HELLO` message.
2. Place the bcrypt result for the following inputs: the user password and the salt given in the `PARAMS` message.
3. Hash everything with SHA3-256.

To sum things up:

```
CHALLENGE = SHA3_256(<22_RAND_CHARS>+<BCRYPT_RESULT>)
```

Here is the format of the message:

```
CONFIRM <CHALLENGE>\r\n
```

## Room navigation

> The user needs to be authenticated for this part.

### CREATE

This message creates a new chat room:

```
CREATE <ROOM_NAME>\r\n
```

The server will respond with a `+OK` message with the new room's id as the content.

### JOIN

This message allows the user to join an existing chat room:

```
JOIN <ROOM_ID>\r\n
```

### LEAVE

This message removes the user from the specified chat room:

```
LEAVE <ROOM_ID>\r\n
```

## Message exchange

> The user needs to be authenticated for this part.

### MSGS

This message allows the user to send a message in the current chat room:

```
MSGS <CONTENT>\r\n
```

### MSG

This message notifies the client that a new message has arrived:

```
MSG <AUTHOR> <CONTENT>\r\n
```

## Status

### +OK

This message notifies the client that the last action went successfully:

```
+OK <MESSAGE>\r\n
```

### -ERR

This message notifies the client that the last action didn't go well:

```
-ERR <MESSAGE>\r\n
```

---

&copy; 2024 Quentin Eppe. All Rights Reserved.
