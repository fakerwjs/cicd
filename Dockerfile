FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o app ./cmd/app

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=3s \
  CMD wget -qO- http://127.0.0.1:8080/health || exit 1

CMD ["./app"]
