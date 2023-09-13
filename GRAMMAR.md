# Grammar

Every client wanting to communicate with this server must respect the chat server grammar.

## Register

First of all, the client has to register to the server by giving a username and a room name with the following message:

```
+REGISTER <USERNAME> <ROOM_NAME>\r\n
```

## Send a message

Then the client is able to send a message to every other user in the same room with the following line:

```
+MSGS <MESSAGE_CONTENT>\r\n
```

## Receive a message

Here is the line the client will receive for a new message:

```
+MSG <SENDER_USERNAME> <MESSAGE_CONTENT>\r\n
```

## OK

If an action passed successfully the client will receive a message with the following format:

```
+OK <MESSAGE>\r\n
```

## Error

If an action failed the client will receive a message with the following format:

```
+ERR <MESSAGE>\r\n
```

---

&copy; 2023 Quentin Eppe. All Rights Reserved.
