FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main ./
COPY --from=builder /app/static ./static
EXPOSE 3000
CMD ["./main"]
