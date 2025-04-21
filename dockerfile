FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tz-telecom ./cmd/app

FROM debian:bullseye-slim

WORKDIR /app

# netcat
RUN apt-get update && apt-get install -y netcat && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/tz-telecom .
COPY --from=builder /app/internal /app/internal
COPY wait-for-it.sh .
COPY entrypoint.sh .

RUN chmod +x wait-for-it.sh entrypoint.sh

ENV PORT=8080
EXPOSE 8080

CMD ["./entrypoint.sh"]
