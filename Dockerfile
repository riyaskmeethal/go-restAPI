FROM golang:1.20-alpine as builder
RUN apk update
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./main .

FROM golang:1.20-alpine
WORKDIR /app
COPY --from=builder app/main ./main
COPY config.yml ./
EXPOSE 8000
CMD ["./main"]