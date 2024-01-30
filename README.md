# mockpi

mockpi is a tool for faking APIs. it responds to all methods and endpoints, returning the JSON response body provided via the `x-response-json` request header.

```shell
➜ curl -i -s -X POST -H 'x-response-json:{"id":0,"name":"nate"}' -d '{"name":"nate"}' localhost:8080/users/create
HTTP/1.1 200 OK
Date: Thu, 19 Oct 2023 04:14:33 GMT
Content-Length: 22
Content-Type: application/json; charset=utf-8

{"id":0,"name":"nate"}
```

## Run

* `go run main.go`
* `docker run docker.io/notnmeyer/mockpi:latest`

## Available headers

You can customize Mockpi's response with these headers:

| header | description | validation |
|--------|-------------|------------|
| x-response-json | the JSON response body to respond with | must parse as JSON |
| x-response-code | the response code to respond with | must parse into a number 100-599 (inclusive) |

## Error responses

If the validation for either header isn't met, a 400 Bad Request is returned with an error encoded in JSON, `{"error": "some message"}`.

## Build container image

```shell
  docker buildx create --use
  docker buildx build --platform linux/amd64 -t mockpi:latest . --load
```
