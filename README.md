# bearer-auth-api
REST API server using gin-gonic to check "Authorization: Bearer" header.

[![TravisCI Status](https://travis-ci.org/nmatsui/bearer-auth-api.svg?branch=master)](https://travis-ci.org/nmatsui/bearer-auth-api)
[![DockerHub Status](https://dockerbuildbadges.quelltext.eu/status.svg?organization=nmatsui&repository=bearer-auth-api)](https://hub.docker.com/r/nmatsui/bearer-auth-api/builds/)

## Description
This REST API Server receives any path and any methods, and checks Bearer Token ("Authorization: Bearer *TOKEN*" Request Header).

1. If no Bearer Token, respond `401 Unauhtorized` always.
1. If Token does not be found in `AUTH_TOKENS` JSON which is given from the environment variable, respond `401 Unauthorized`.
1. If Token is found, but path does not be allowed, respond `403 Forbidden`
1. If Token is found and path is allowed, respond `200 OK`

This REST API is assumed to use with [Ambassador](https://www.getambassador.io/) on [Kubernetes](https://www.getambassador.io/).

## `AUTH_TOKENS` JSON template

```text
{
  "<token1>": ["<allowed path regex>", "<allowed path regex>"...],
  "<token2>": [...],
  ...
}
```

example)

```json
{
  "Znda7iglaqdoltsp7kDl60TvkkszcEGU": ["^/path1/.*$", "^/path2/path2-2/.*$"],
  "fANtLRTszYAayjtmLFllSHBrt2zRyoqV": ["^/path2/.*$"]
}
```


## Run as Docker container

1. Pull container [nmatsui/bearer-auth-api](https://hub.docker.com/r/nmatsui/bearer-auth-api/) from DockerHub

    ```bash
    $ docker pull nmatsui/bearer-auth-api
    ```
1. Run Container
    * If you want to change exposed port, set `LISTEN_PORT` environment variable.

    ```bash
    $ docker run -d -p 3000:3000 nmatsui/bearer-auth-api
    ```

## Build from source code

1. go get

    ```bash
    $ go get -u github.com/nmatsui/bearer-auth-api
    ```
1. go install

    ```bash
    $ go install github.com/nmatsui/bearer-auth-api
    ```

## License

[Apache License 2.0](/LICENSE)

## Copyright
Copyright (c) 2018 Nobuyuki Matsui <nobuyuki.matsui@gmail.com>
