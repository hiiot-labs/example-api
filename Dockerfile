FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o example-api .

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/example-api .
EXPOSE 8899
CMD ["./example-api"]
