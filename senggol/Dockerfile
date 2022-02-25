FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o senggol-run cmd/main.go

CMD ["./senggol-run"]