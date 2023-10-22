FROM docker.io/golang:1.21 as builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY *.go .
COPY go.mod .
RUN go mod tidy
RUN go build -o dist/mockpi ./main.go

FROM scratch
COPY --from=builder /build/dist/mockpi /usr/local/bin/
CMD ["/usr/local/bin/mockpi"]
