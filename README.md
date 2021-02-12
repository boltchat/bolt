![boltchat](https://raw.githubusercontent.com/boltchat/branding/main/svg/bolt-banner.svg)
> ⚡ A fast, lightweight, and secure chat protocol, client and server, written in Go.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/boltchat/bolt/Deploy?label=deploy)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/boltchat/bolt/Test?label=test)
![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/boltchat/bolt/develop)

## About
> ⚠ IMPORTANT: This project is still a work-in-progress. I strongly discourage installing this on a
> public-facing server as it could potentially harm your security and privacy. See the
> [roadmap](https://github.com/boltchat/bolt/projects) for the progress of this project.

_Bolt_ is intended as a modern replacement for [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat).
I started this project because I feel like there aren't many open source chat protocols that follow modern
standards.

Not only do I think it's a great fit for an IRC replacement; it might even be suitable for a replacement of
present-day proprietary protocols and chat applications, such as Discord and Slack. _Bolt_ comes with
a nifty text-based user interface, but since it uses its own protocol, it's possible to build a GUI client
in, say, Electron. (please don't, use [Tauri](https://github.com/tauri-apps/tauri))

## Roadmap
The project boards for _Bolt_ can be found [here](https://github.com/boltchat/bolt/projects).

## References
* [Installation guide](./docs/installation.md)
* [Quick start guide](./docs/quick-start.md)
* [Protocol Specification](./docs/protocol-spec.md)

## Author
[Kees van Voorthuizen (@keesvv)](https://github.com/keesvv)

## Logo
The lightning bolt used in the logo of this project is courtesy of [Icons8](https://icons8.com/icons/set/lightning-bolt--v1).

## License
This project is licensed under the [Apache License 2.0](./LICENSE).
