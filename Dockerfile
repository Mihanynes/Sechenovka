FROM golang:alpine

WORKDIR /app
ADD go.mod .
COPY . .
RUN go build -o main main.go

ENTRYPOINT ["./main"]