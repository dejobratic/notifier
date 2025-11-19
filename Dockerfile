FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o notifier .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/notifier /app/notifier
EXPOSE 8080
ENTRYPOINT ["/app/notifier"]