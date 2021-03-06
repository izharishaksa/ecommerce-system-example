FROM golang:1.18-alpine3.14 as builder

WORKDIR /app
COPY ./lib ./lib
COPY ./inventory-service ./inventory-service

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/inventory-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/inventory-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/inventory-service /usr/bin/inventory-service
EXPOSE 4001 6001

ENV APP_NAME="inventory-service"
ENV APP_HTTP_PORT="4001"

ENV DB_DSN="postgres://postgres:postgres@inventory-database-host:5432/inventory?sslmode=disable"
ENV DB_MAX_IDLE_CONNECTIONS="20"
ENV DB_MAX_OPEN_CONNECTIONS="100"
ENV DB_MAX_IDLE_DURATION_IN_SECS="60"
ENV DB_MAX_LIFE_TIME_DURATION_IN_SECS="300"

CMD ["inventory-service"]