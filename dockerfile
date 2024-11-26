FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

# RUN go build -o myApp .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o myApp .

FROM alpine:3.20
WORKDIR /server
COPY --from=builder /app/myApp .
COPY --from=builder /app/templates ./templates

EXPOSE 3000

CMD ["./myApp"]