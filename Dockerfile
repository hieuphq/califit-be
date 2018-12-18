FROM golang:1.11.1-alpine3.8
WORKDIR /go/src/github.com/hieuphq/califit-be
ADD . /go/src/github.com/hieuphq/califit-be
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o califit-be ./ext/cmd/califit-be

FROM alpine:3.8
WORKDIR /
RUN apk add --no-cache ca-certificates
RUN apk add --update openssl
RUN openssl genrsa -out rsa_key 2048
RUN openssl rsa -in rsa_key -pubout > rsa_key.pub

COPY --from=0 /go/src/github.com/hieuphq/califit-be/califit-be /califit-be
COPY swagger swagger

EXPOSE 9000

CMD ["/califit-be"]

