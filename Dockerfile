# Stage 1: Сборка приложения
FROM golang:1.21 AS builder

WORKDIR /app

# Копируем go.mod и go.sum для ускорения сборки при изменении зависимостей
COPY go.mod .
COPY go.sum .

# Скачиваем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY internal internal
COPY cmd cmd

# Собираем Go-приложение
RUN CGO_ENABLED=0 go build -o /app/server cmd/main.go

# Stage 2: Финальный образ
FROM alpine:latest

WORKDIR /app

# Копируем исполняемый файл из Stage 1
COPY --from=builder /app/server .

# Запускаем приложение
CMD ["./server"]
