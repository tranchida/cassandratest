# Build stage
FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/internal ./internal
EXPOSE 8080
CMD ["./app"]
