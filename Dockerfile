FROM golang:latest

# Установите необходимые библиотеки
RUN apt-get update && apt-get install -y libsqlite3-dev gcc

WORKDIR /app
COPY go.mod ./
COPY . .

# Выполните сборку
RUN go build -o service ./cmd/service
RUN go build -o push_sender ./cmd/push_sender && go clean -cache -modcache
