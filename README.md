![@observerly:nocturnal](https://user-images.githubusercontent.com/84131395/174099920-eda972ff-2493-4eea-b572-e31394cd0d36.jpg)

Nocturnal 🌑 is observerly's Gin Gonic API for Lunar and Solar advanced scheduling, that utilises [Dusk](https://github.com/observerly/dusk).

## API Development

### Project Requirements

- [Docker](https://www.docker.com/).
- [Docker Compose](https://docs.docker.com/compose/install/).
- [Go](https://go.dev/dl/) version 1.17+ (or later).

### Local Development

The Nocturnal development stack can be built with the following `docker` `compose` command:

```console
$ docker compose -f local.yml build
```

The Nocturnal development stack can then be started with:

```console
$ docker compose -f local.yml up
```

or simply by utilising the following go command:

```console
go run main.go
```

### Testing

The Nocturnal development stack can be tested with the following command:

```console
GIN_MODE=release bash -c "go test ./... -race -coverprofile=coverage.txt -covermode=atomic -v"
```
