# Installation

## Client & server
### Binaries
Please have a look at the [Releases](https://github.com/boltchat/bolt/releases) page for the most
up-to-date client and server binaries.

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

## Server
### Docker (preferred)
If you'd like to run the server in a Docker container, you should use the most up-to-date
image from [Docker Hub](https://hub.docker.com/r/boltchat/server). The following command
should get you up and running within seconds:

```bash
docker run -p 3300:3300 --tty boltchat/server:latest
```

### Docker Compose
Sample configuration can be found [here](../conf/docker/docker-compose.yml).

### Daemon
You can also run the server as a daemon. Service files can be found below:
* [`systemd` service](../conf/linux/systemd/boltchat.service)
* [`runit` service](../conf/linux/runit)
