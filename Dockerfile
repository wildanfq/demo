FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app .

FROM alpine:3.21

WORKDIR /app

RUN adduser -D -u 10001 appuser && \
    chown -R appuser:appuser /app
USER appuser

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
