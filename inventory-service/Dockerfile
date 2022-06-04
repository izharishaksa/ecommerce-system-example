FROM golang:1.18-alpine3.14 AS builder

WORKDIR /app
COPY . ./

RUN apk update && apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/inventory-service ./cmd/rest

FROM alpine

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

COPY --from=builder /usr/bin/inventory-service /usr/bin/inventory-service
EXPOSE 2001

CMD ["inventory-service"]