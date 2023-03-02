![@observerly:nocturnal](https://user-images.githubusercontent.com/84131395/174099920-eda972ff-2493-4eea-b572-e31394cd0d36.jpg)

Nocturnal 🌑 is observerly's Gin Gonic API for Lunar and Solar advanced scheduling, that utilises [Dusk](https://github.com/observerly/dusk), available at https://nocturnal.observerly.com as a highly performant public API. 

The Nocturnal API is designed to be stateless, and to return a response based on the current time, and the parameters passed to the API.

Please note, the specifics of the Nocturnal API is currently in active development, and is subject to without warning.

## Usage

### Introduction

As mentioned, the Nocturnal API is designed to be stateless. That is, a request to the API will return a response based on the current time, and the parameters passed to the API. 

This means that the API can be used to understanding the position of the Moon, the Sun, or any transiting body (e.g., a Star) at any time in the past, present, or future.

### Schema

The Nocturnal base URL for the API will be {HOST}/api/v2/ where {HOST} is the host name of the API, e.g., https://nocturnal.observerly.com.

There is a deprecate version of the API, but as of 14.02.2023 we do not recommend its use. The deprecate version of the API is available at {HOST}/api/v1/ where {HOST} is the host name of the API, e.g., https://nocturnal.observerly.com.

The Nocturnal API has a standardised API JSON response format, which adheres to the [JSON API](https://jsonapi.org/) specification as well as the schema defined in the [OpenAPI](https://swagger.io/specification/) specification.

For Sun, Moon and Transit endpoints, the API JSON response adheres to the following schema:

```json
{
  "observer": {
    
  },
  "rise": {

  },
  "maximum": {

  },
  "set": {

  },
}
```

For Twilight endpoints, the API JSON response adheres to the following schema:

```json
{
  "observer": {
    
  },
  "civil": {

  },
  "nautical": {

  },
  "astronomical": {

  },
}
```

### API Endpoints

The Nocturnal API has the following endpoints:

- [GET /api/v2/sun](#get-apiv2sun)

- [GET /api/v2/moon](#get-apiv2moon)

- [GET /api/v2/transit](#get-apiv2transit)

- [GET /api/v2/twilight](#get-apiv2twilight)

## API Development

### Project Requirements

We welcome any and all contributions or raised issues to this repository.

To contribute to this project, you will need the following tools installed on your local machine:

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

or simply by utilising the following combined command:

```console
docker compose -f local.yml up --build
```

### Testing

The Nocturnal development stack can be tested with the following command:

```console
GIN_MODE=release bash -c "go test ./... -race -coverprofile=coverage.txt -covermode=atomic -v"
```