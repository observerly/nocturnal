# Nocturnal 🌑

Nocturnal is observerly's Gin Gonic API for Lunar and Solar advanced scheduling, that utilises [Dusk](https://github.com/observerly/dusk).

## API Development

### Project Requirements

- [Docker](https://www.docker.com/).
- [Docker Compose](https://docs.docker.com/compose/install/).
- [Go](https://go.dev/dl/) version 1.17+ (or later).

### Local Development

The Nocturnal development stack can be built with the following `docker` `compose` command:

```console
$ docker compose build"
```

The Nocturnal development stack can then be started with:

```console
$ docker compose up"
```

or simply by utilising the following go command:

```console
go run main.go
```