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
COPY --from=builder /app/version .
RUN mkdir -p data
ENV APP_MODE=release
ENV APP_PORT=8080
EXPOSE 8080
CMD ["./redis-ui"]
