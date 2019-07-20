#

* Docker file for ipshow

```dockerfile
FROM golang:1.8 as builder
WORKDIR /go/src/github.com/gangsta/dockerip/
ADD ./main.go /go/src/github.com/gangsta/dockerip/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./dockerip /go/src/github.com/gangsta/dockerip/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache curl
COPY --from=builder /go/src/github.com/gangsta/dockerip/dockerip /app/
RUN addgroup --gid 3034 dockerip
RUN adduser -h /app -s /bin/sh -G dockerip -u 3034 -D dockerip
RUN chown dockerip:dockerip -R /app
USER dockerip
WORKDIR /app
ENTRYPOINT ["/app/dockerip"]
```
go
go
