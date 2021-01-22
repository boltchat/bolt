![bolt.chat](./assets/logo/boltchat-logo.jpg)

> ⚡ A fast, lightweight, and secure chat protocol, client and server, written in Go.

## About
_bolt.chat_ is intended as a modern replacement for [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat).
I started this project because I feel like there aren't many open source chat protocols that follow modern
standards.

Not only do I think it's a great fit for an IRC replacement; it might even be suitable for a replacement of
present-day proprietary protocols and chat applications, such as Discord and Slack. _bolt.chat_ comes with
a nifty text-base user interface, but since it uses its own protocol, it's possible to build a GUI client
in, say, Electron (please don't, use [Tauri](https://github.com/tauri-apps/tauri))

## Installation
Please have a look at the [Releases](https://github.com/keesvv/bolt.chat/releases) page for the most up-to-date client and server binaries.

If you'd like to compile _bolt.chat_ from source, please follow the steps below:

```bash
git clone git@github.com:keesvv/bolt.chat.git
cd bolt.chat
go get github.com/magefile/mage
go install github.com/magefile/mage
mage
```

## Author
[Kees van Voorthuizen](https://github.com/keesvv)

## Logo
The lightning bolt used in the logo of this project is courtesy of [Icons8](https://icons8.com/icons/set/lightning-bolt--v1).

## License
This project is licensed under the [GPLv3](./LICENSE).