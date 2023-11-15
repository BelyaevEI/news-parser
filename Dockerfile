# Используем официальный образ Golang как базовый образ
FROM golang:1.21 AS build

# Установка рабочей директории
WORKDIR /src

# Копируем зависимости
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копируем исходный код проекта в контейнер
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Создание конечного образа
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированный бинарный файл из предыдущего этапа
COPY --from=build /src/main .

# Определение переменных окружения, если необходимо
# ENV TELEGRAM_BOT_TOKEN=YOUR_BOT_TOKEN
# ENV OTHER_ENV_VARIABLE=VALUE

# Запуск бота
CMD ["./main"]
