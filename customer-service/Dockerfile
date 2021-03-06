FROM golang:1.18-alpine3.14 as builder

WORKDIR /app
COPY ./lib ./lib
COPY ./customer-service ./customer-service

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/customer-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/customer-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/customer-service /usr/bin/customer-service
EXPOSE 4002

ENV APP_NAME="customer-rest-service"
ENV APP_HTTP_PORT="4002"

ENV DB_DSN="postgres://postgres:postgres@customer-database-host:5432/customer?sslmode=disable"
ENV DB_MAX_IDLE_CONNECTIONS="20"
ENV DB_MAX_OPEN_CONNECTIONS="100"
ENV DB_MAX_IDLE_DURATION_IN_SECS="60"
ENV DB_MAX_LIFE_TIME_DURATION_IN_SECS="300"

CMD ["customer-service"]