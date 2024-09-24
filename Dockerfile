FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY scripts/wait-for-it.sh .

EXPOSE 8080

CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]
