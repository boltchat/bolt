# Protocol Specification
> ⚠ Everything contained in this document is subject to change until the first stable release (v1.0.0).

_bolt.chat_ exchanges information by sending JSON-marshalled data over a TCP connection. The exchanged information is often in the form of [Events](##Events).

## Events
Events represent any kind of activity happening in a server. They can be emitted by both the client and the server. Emitted events may not be limited to one connection only: some events are broadcasted by the server to all connected clients or to a selection of connected clients.

The most basic form of an event looks like this:
```json
{
  "e": {
    "t": "msg",
    "c": 1611670138,
    "r": 1611670139
  },
  ...
}
```

| key | type        | desc                 | can send |
|-----|-------------|----------------------|----------|
| `e` | `EventMeta` | Event metadata       | both     |
| `t` | `string`    | Type identifier      | both     |
| `c` | `int64`     | Creation date (Unix) | both     |
| `r` | `int64`     | Receipt date (Unix)  | server   |

The `...` indicates additional data that is accompanied with the event.
