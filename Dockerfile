FROM golang:1.21.4

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]