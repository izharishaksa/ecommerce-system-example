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

CMD ["inventory-service"]