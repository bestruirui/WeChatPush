FROM golang:alpine as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM alpine
COPY --from=builder /app/main /app/main
CMD  /app/main 