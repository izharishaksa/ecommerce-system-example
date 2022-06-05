FROM golang:1.18-alpine3.14 as builder

WORKDIR /app
COPY ./lib ./lib
COPY ./order-service ./order-service

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/order-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/order-service ./cmd/rest

FROM alpine

COPY --from=builder /usr/bin/order-service /usr/bin/order-service
EXPOSE 4003

ENV APP_NAME="order-service"
ENV APP_HTTP_PORT="4003"

ENV DB_DSN="postgres://postgres:postgres@order-database-host:5432/order?sslmode=disable"
ENV DB_MAX_IDLE_CONNECTIONS="20"
ENV DB_MAX_OPEN_CONNECTIONS="100"
ENV DB_MAX_IDLE_DURATION_IN_SECS="60"
ENV DB_MAX_LIFE_TIME_DURATION_IN_SECS="300"

CMD ["order-service"]