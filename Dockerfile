FROM golang:1.21 AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .
# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Создание конечного образа
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированный бинарный файл из предыдущего этапа
COPY --from=build /src/main .

CMD ["./main"]  