FROM golang:1.23

# Установите необходимые библиотеки
RUN apt-get update && apt-get install -y libsqlite3-dev gcc

WORKDIR /app
# Копируем только файлы зависимостей сначала
COPY go.mod go.sum ./

# Скачиваем зависимости (кэшируется отдельно)
RUN go mod download

COPY . .

# Выполните сборку
RUN go build -o service ./cmd/service
RUN go build -o telegram_producer ./cmd/telegram_producer
RUN go build -o telegram_consumer ./cmd/telegram_consumer
RUN go build -o push_sender ./cmd/push_sender && go clean -cache -modcache
