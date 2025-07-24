# Use newer Go version that matches go.mod requirement
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o blog-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/blog-api .

EXPOSE 8080

CMD ["./blog-api"]
