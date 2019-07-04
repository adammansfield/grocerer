FROM golang:1.10 AS build
WORKDIR /go/src
COPY internal/go ./go
COPY internal/main.go .

RUN mkdir -p /go/src/github.com/adammansfield/grocerer/pkg/ourgrocer
COPY pkg/ourgrocer /go/src/github.com/adammansfield/grocerer/pkg/ourgrocer

ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go build -a -installsuffix cgo -o openapi .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch AS runtime
COPY --from=build /go/src/openapi .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
