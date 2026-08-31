# --- Base Stage ---
FROM golang:1.26.5-alpine AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# --- Builder Stage (คอมไพล์เพื่อ prod) ---
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./cmd/server/main.go

# --- Production Runner Stage ---
FROM alpine:latest AS runner
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/server /app/server

EXPOSE 8080
CMD ["/app/server"]