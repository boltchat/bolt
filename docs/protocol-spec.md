# Protocol Specification
> ⚠ Everything contained in this document is subject to change until the first stable release (v1.0.0). Parts denoted with a '⚠' symbol have a certainty of being changed in the near future.

> **Legend:**
>
> `-->` client-to-server (send)
>
> `<--` server-to-client (receive)
> 
> `<-->` client-to-server and server-to-client (send and receive)

_Bolt_ exchanges information by sending JSON-marshalled data over a TCP connection. The exchanged information is often in the form of [Events](#events).

## Events
Events represent any kind of activity happening in a server. They can be emitted by both the client and the server. Emitted events may not be limited to one connection only: some events are broadcasted by the server to all connected clients or to a selection of connected clients.

The most basic form of an event looks like this:

`<-->`
```json
{
  "e": {
    "t": "msg",
    "c": 1611670138
  },
  "d": {
    ...
  }
}
```

| key | type                      | desc                 |
|-----|---------------------------|----------------------|
| `e` | [`EventMeta`](#eventmeta) | Event metadata       |
| `d` | [`EventData`](#eventdata) | Event data           |

The `...` indicates additional data that is accompanied with the event, depending on the type of event.

### `EventMeta`
Basic metadata that is accompanied with each event.

| key | type        | desc                 |
|-----|-------------|----------------------|
| `t` | `string`    | Type identifier      |
| `c` | `int64`     | Creation date (Unix) |

### `EventData`
See below for an overview of data types.

## Messages
Messages represent chat messages. They can be sent by the client only. The server is responsible for delivering these messages to the intended recipients.

A message looks like this:

`-->`
```json
{
  "body": "This is a message.",
  "sig": "-----BEGIN PGP SIGNATURE----- (...)",
  "user": {
    "nick": "keesvv"
  }
}
```
> ⚠ `user` will soon be removed from message events emitted by clients.

`<--`
```json
{
  "body": "This is a message.",
  "fprint": "131e12c7087e576743cb6c26eaf3f4d4ee6305a9",
  "user": {
    "nick": "keesvv"
  }
}
```

| key       | type     | desc                         |
|-----------|----------|------------------------------|
| `body`    | `string` | Message content              |
| `sig`     | `string` | ASCII-armored PGP signature  |
| `fprint`  | `string` | Public key fingerprint (hex) |
| `user`    | `User`   | User sending the message     |
