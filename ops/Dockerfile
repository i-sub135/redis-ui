FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o redis-ui ./main.go

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/redis-ui .
COPY --from=builder /app/source/templates ./source/templates
COPY --from=builder /app/web ./web
COPY --from=builder /app/config.yaml .
EXPOSE 8080
CMD ["./redis-ui"]
