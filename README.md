# mockpi

mockpi is a tool for faking APIs. it responds to all methods and endpoints, optionally including a response body provided in the request as the `x-response-json` header.

```shell
➜ curl -i -s -X POST -H 'x-response-json:{"id":0,"name":"nate"}' -d '{"name":"nate"}' localhost:8080/users/create
HTTP/1.1 200 OK
Date: Thu, 19 Oct 2023 04:14:33 GMT
Content-Length: 22
Content-Type: text/plain; charset=utf-8

{"id":0,"name":"nate"}
```


## Run

`go run main.go`

## Build container image

```shell
  docker buildx create --use
  docker buildx build --platform linux/amd64 -t mockpi:latest . --load
```