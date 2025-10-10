FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=arm64 go build -o main ./cmd/app/.

EXPOSE 8000

CMD ["./main"]