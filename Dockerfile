FROM golang:1.17.0-alpine as base

FROM base as dev
WORKDIR /app
EXPOSE 8080

RUN apk update && apk add libc-dev && apk add gcc && apk add make

COPY ./go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Install CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint.sh

RUN chmod +rx /entrypoint.sh
ENTRYPOINT ["sh", "/entrypoint.sh"]
