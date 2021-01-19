FROM golang:1.15.6-alpine

WORKDIR /src

COPY . .

RUN go build cmd/server/main.go

CMD ["./main"]
