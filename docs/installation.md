# Installation

### Binaries
Please have a look at the [Releases](https://github.com/boltchat/bolt/releases) page for the most
up-to-date client and server binaries.
> Unfortunately, there are currently no binaries available for download. For the time being,
> follow the steps as described in the subsection **From source**.

### Docker
If you'd like to run the server in a Docker container, follow the compilation steps below and run `mage docker:build`.

### From source
If you'd like to compile _Bolt_ from source, please follow the steps below:

#### Prerequisites
* Git
* Go (v1.15.6)

#### Cloning & installing
```bash
git clone git@github.com:boltchat/bolt.git
cd bolt
go get github.com/magefile/mage
go install github.com/magefile/mage
```

#### Building
Run `mage` to see all available targets.
