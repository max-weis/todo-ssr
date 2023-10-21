FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o todo-app ./cmd/todo

FROM debian:12.2-slim
COPY --from=builder /app/todo-app /todo-app
ENTRYPOINT ["/todo-app"]
