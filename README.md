![bolt.chat](./assets/banner/boltchat-banner.jpg)

> ⚡ A fast, lightweight, and secure chat protocol, client and server, written in Go.

## About
> ⚠ IMPORTANT: This project is still a work-in-progress. I strongly discourage installing this on a
> public-facing server as it could potentially harm your security and privacy. See the
> [roadmap](https://github.com/bolt-chat/bolt.chat/projects) for the progress of this project.

_bolt.chat_ is intended as a modern replacement for [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat).
I started this project because I feel like there aren't many open source chat protocols that follow modern
standards.

Not only do I think it's a great fit for an IRC replacement; it might even be suitable for a replacement of
present-day proprietary protocols and chat applications, such as Discord and Slack. _bolt.chat_ comes with
a nifty text-base user interface, but since it uses its own protocol, it's possible to build a GUI client
in, say, Electron. (please don't, use [Tauri](https://github.com/tauri-apps/tauri))

## Roadmap
The project boards for _bolt.chat_ can be found [here](https://github.com/bolt-chat/bolt.chat/projects).

## Installation
### Binaries
Please have a look at the [Releases](https://github.com/keesvv/bolt.chat/releases) page for the most
up-to-date client and server binaries.
> Unfortunately, there are currently no binaries available for download. For the time being,
> follow the steps as described in the subsection **From source**.

### Docker
If you'd like to run the server in a Docker container, follow the compilation steps below and run `mage docker:build`.

### From source
If you'd like to compile _bolt.chat_ from source, please follow the steps below:

#### Prerequisites
* Git
* Go (v1.15.6)

#### Cloning & installing
```bash
git clone git@github.com:bolt-chat/bolt.chat.git
cd bolt.chat
go get github.com/magefile/mage
go install github.com/magefile/mage
```

#### Building
Run `mage` to see all available targets.

## Quick start
### Server
`// TODO`

### Client
`// TODO`

## Author
[Kees van Voorthuizen (@keesvv)](https://github.com/keesvv)

## Logo
The lightning bolt used in the logo of this project is courtesy of [Icons8](https://icons8.com/icons/set/lightning-bolt--v1).

## License
This project is licensed under the [GPLv3](./LICENSE).
