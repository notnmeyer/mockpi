# mockpi

mockpi is a tool for faking APIs. it responds to all methods and endpoints, optionally including a response body provided in the request as the `x-response-json` header.

```shell
➜ curl -s -X POST -H 'x-response-json:{"id":0,"name":"nate"}' -d '{"name
":"nate"}' localhost:8080/users/create| jq .
{
  "id": 0,
  "name": "nate"
}
```


## Run

`go run main.go`

## Build container image

```shell
  docker buildx create --use

  # local
  docker buildx build --platform linux/amd64 -t ghcr.io/mockpi:latest . --load

  # push
  docker buildx build --platform linux/amd64 -t ghcr.io/mockpi:latest . --push
```