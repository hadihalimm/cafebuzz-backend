# Build stage
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cafebuzz-backend ./cmd/main.go

# Runtime stage
FROM scratch

WORKDIR /app

COPY --from=builder /app/cafebuzz-backend .

CMD ["./cafebuzz-backend"]

EXPOSE 8080
