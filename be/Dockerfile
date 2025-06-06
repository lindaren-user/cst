FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/myapp .
VOLUME /app/data

CMD ["./myapp"]
