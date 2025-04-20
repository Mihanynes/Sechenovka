FROM golang:latest

# Установите необходимые библиотеки
RUN apt-get update && apt-get install -y libsqlite3-dev gcc

WORKDIR /app
COPY go.mod ./
COPY . .

# Выполните сборку
RUN go build -o ${GOPATH}/bin/service ./cmd/service
RUN go build -o ${GOPATH}/bin/push_sender ./cmd/push_dender && go clean -cache -modcache
