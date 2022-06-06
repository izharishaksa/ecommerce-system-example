FROM golang:1.18-alpine3.14 as builder

WORKDIR /app
COPY ./lib ./lib
COPY ./order-service ./order-service

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/order-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/order-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/order-service /usr/bin/order-service

CMD ["order-service"]
