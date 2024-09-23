# Используйте более раннюю версию Go, если это необходимо
FROM golang:1.22 AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Компилируем приложение с включенным CGO
RUN CGO_ENABLED=1 GOOS=linux go build -o myapp .

# Создаем финальный образ
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/myapp .

# Устанавливаем необходимые библиотеки для работы с cgo
RUN apk add --no-cache sqlite-dev

# Запускаем приложение
CMD ["./app/myapp"]