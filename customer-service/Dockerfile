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

CMD ["customer-service"]